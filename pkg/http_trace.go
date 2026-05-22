package pkg

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync/atomic"
	"time"
)

const httpTraceBodyLimit = 4096

var httpTraceEnabled atomic.Bool

// SetHTTPTraceEnabled 控制 SDK 是否记录 HTTP 请求与响应摘要。
func SetHTTPTraceEnabled(enabled bool) {
	httpTraceEnabled.Store(enabled)
}

func HTTPTraceEnabled() bool {
	return httpTraceEnabled.Load()
}

func wrapHTTPClientForTrace(client *http.Client) *http.Client {
	if client == nil {
		return nil
	}
	if _, ok := client.Transport.(*httpTraceTransport); ok {
		return client
	}
	cp := *client
	rt := cp.Transport
	if rt == nil {
		rt = http.DefaultTransport
	}
	cp.Transport = &httpTraceTransport{rt: rt}
	return &cp
}

type httpTraceTransport struct {
	rt http.RoundTripper
}

func (t *httpTraceTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if !HTTPTraceEnabled() {
		return t.rt.RoundTrip(req)
	}

	begin := time.Now()
	logHTTPRequest(req)
	resp, err := t.rt.RoundTrip(req)
	logHTTPResponse(req, resp, err, time.Since(begin))
	return resp, err
}

func logHTTPRequest(req *http.Request) {
	if req == nil {
		return
	}
	fields := map[string]any{
		"method":  req.Method,
		"url":     sanitizeURL(req.URL),
		"headers": sanitizeHeader(req.Header),
	}
	body, truncated, unavailable := readRequestBodyHead(req)
	if body != "" {
		fields["body_head"] = body
	}
	if truncated {
		fields["body_truncated"] = true
	}
	if unavailable {
		fields["body_unavailable"] = true
	}
	DefaultLogger.Info("sdk http request", loggerFields(fields)...)
}

func logHTTPResponse(req *http.Request, resp *http.Response, err error, duration time.Duration) {
	fields := map[string]any{
		"duration_ms": duration.Milliseconds(),
	}
	if req != nil {
		fields["method"] = req.Method
		fields["url"] = sanitizeURL(req.URL)
	}
	if err != nil {
		fields["error"] = err.Error()
		DefaultLogger.Warn("sdk http response failed", loggerFields(fields)...)
		return
	}
	if resp == nil {
		DefaultLogger.Warn("sdk http response empty", loggerFields(fields)...)
		return
	}
	fields["status"] = resp.StatusCode
	fields["headers"] = sanitizeHeader(resp.Header)
	body, truncated, readErr := readResponseBodyHead(resp)
	if body != "" {
		fields["body_head"] = body
	}
	if truncated {
		fields["body_truncated"] = true
	}
	if readErr != nil {
		fields["body_read_error"] = readErr.Error()
	}
	DefaultLogger.Info("sdk http response", loggerFields(fields)...)
}

func readRequestBodyHead(req *http.Request) (string, bool, bool) {
	if req == nil || req.Body == nil {
		return "", false, false
	}
	if req.GetBody == nil {
		return "", false, true
	}
	body, err := req.GetBody()
	if err != nil {
		return "", false, true
	}
	defer body.Close()
	return readBodyHead(body)
}

func readResponseBodyHead(resp *http.Response) (string, bool, error) {
	if resp == nil || resp.Body == nil {
		return "", false, nil
	}
	original := resp.Body
	limited := io.LimitReader(original, httpTraceBodyLimit+1)
	buf, err := io.ReadAll(limited)
	if err != nil {
		resp.Body = &multiReadCloser{
			Reader: io.MultiReader(bytes.NewReader(buf), original),
			Closer: original,
		}
		return "", false, err
	}
	truncated := len(buf) > httpTraceBodyLimit
	logBuf := buf
	if truncated {
		logBuf = buf[:httpTraceBodyLimit]
	}
	resp.Body = &multiReadCloser{
		Reader: io.MultiReader(bytes.NewReader(buf), original),
		Closer: original,
	}
	return sanitizeBody(logBuf), truncated, nil
}

func readBodyHead(r io.Reader) (string, bool, bool) {
	buf, err := io.ReadAll(io.LimitReader(r, httpTraceBodyLimit+1))
	if err != nil {
		return "", false, true
	}
	truncated := len(buf) > httpTraceBodyLimit
	if truncated {
		buf = buf[:httpTraceBodyLimit]
	}
	return sanitizeBody(buf), truncated, false
}

type multiReadCloser struct {
	io.Reader
	io.Closer
}

func sanitizeURL(u *url.URL) string {
	if u == nil {
		return ""
	}
	cp := *u
	q := cp.Query()
	for key := range q {
		if isSensitiveKey(key) {
			q.Set(key, "<redacted>")
		}
	}
	cp.RawQuery = q.Encode()
	return cp.String()
}

func sanitizeHeader(header http.Header) map[string]string {
	if len(header) == 0 {
		return nil
	}
	result := make(map[string]string, len(header))
	for key, values := range header {
		if isSensitiveKey(key) {
			result[key] = "<redacted>"
			continue
		}
		result[key] = strings.Join(values, ",")
	}
	return result
}

func sanitizeBody(body []byte) string {
	s := string(body)
	replacer := strings.NewReplacer("\n", `\n`, "\r", `\r`)
	return replacer.Replace(s)
}

func isSensitiveKey(key string) bool {
	k := strings.ToLower(key)
	return strings.Contains(k, "authorization") ||
		strings.Contains(k, "token") ||
		strings.Contains(k, "secret") ||
		strings.Contains(k, "password") ||
		strings.Contains(k, "signature") ||
		strings.Contains(k, "credential") ||
		strings.Contains(k, "cookie") ||
		strings.Contains(k, "key")
}

func loggerFields(fields map[string]any) []any {
	if len(fields) == 0 {
		return nil
	}
	keys := make([]string, 0, len(fields))
	for key := range fields {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	result := make([]any, 0, len(keys)*2)
	for _, key := range keys {
		result = append(result, key, fields[key])
	}
	return result
}
