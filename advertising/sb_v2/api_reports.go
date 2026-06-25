package sb_v2

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/bytedance/sonic"
)

// ReportsAPIService 封装 SB legacy v2(HSA) 活动同步报表:创建 / 查询状态 / 下载。
type ReportsAPIService service

type ApiRequestReportRequest struct {
	ctx        context.Context
	ApiService *ReportsAPIService
	content    *RequestReportContent
}

func (r ApiRequestReportRequest) Content(content RequestReportContent) ApiRequestReportRequest {
	r.content = &content
	return r
}

func (r ApiRequestReportRequest) Execute() (*ReportResponse, *http.Response, error) {
	return r.ApiService.RequestReportExecute(r)
}

// RequestReport 创建一份 SB v2 活动报表(POST /v2/hsa/campaigns/report)。
func (a *ReportsAPIService) RequestReport(ctx context.Context) ApiRequestReportRequest {
	return ApiRequestReportRequest{ApiService: a, ctx: ctx}
}

func (a *ReportsAPIService) RequestReportExecute(r ApiRequestReportRequest) (*ReportResponse, *http.Response, error) {
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.RequestReport")
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	if r.content == nil {
		return nil, nil, reportError("content is required and must be specified")
	}
	localVarPath := localBasePath + "/v2/hsa/campaigns/report"
	localVarHeaderParams := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, http.MethodPost, r.content, localVarHeaderParams, url.Values{}, url.Values{}, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, body, err := a.client.doRead(req)
	if err != nil {
		return nil, resp, err
	}
	var out ReportResponse
	if len(body) > 0 {
		_ = sonic.Unmarshal(body, &out)
	}
	// 部分站点不在响应体回 reportId,而是放在 Location 头里,取末段兜底。
	if out.ReportId == "" && resp != nil {
		if loc := resp.Header.Get("Location"); loc != "" {
			out.ReportId = lastURLSegment(loc)
		}
	}
	if out.ReportId == "" {
		return nil, resp, &GenericOpenAPIError{body: body, error: "SB v2 report response missing reportId"}
	}
	return &out, resp, nil
}

type ApiGetReportRequest struct {
	ctx        context.Context
	ApiService *ReportsAPIService
	reportId   string
}

func (r ApiGetReportRequest) Execute() (*ReportResponse, *http.Response, error) {
	return r.ApiService.GetReportExecute(r)
}

// GetReport 查询报表状态(GET /v2/reports/{reportId});就绪后 location/url 给出下载地址。
func (a *ReportsAPIService) GetReport(ctx context.Context, reportId string) ApiGetReportRequest {
	return ApiGetReportRequest{ApiService: a, ctx: ctx, reportId: reportId}
}

func (a *ReportsAPIService) GetReportExecute(r ApiGetReportRequest) (*ReportResponse, *http.Response, error) {
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.GetReport")
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	localVarPath := localBasePath + "/v2/reports/" + url.PathEscape(r.reportId)
	localVarHeaderParams := map[string]string{"Accept": "application/json"}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, http.MethodGet, nil, localVarHeaderParams, url.Values{}, url.Values{}, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, body, err := a.client.doRead(req)
	if err != nil {
		return nil, resp, err
	}
	var out ReportResponse
	if len(body) > 0 {
		_ = sonic.Unmarshal(body, &out)
	}
	if out.Status == "" {
		out.Status = "UNKNOWN"
	}
	return &out, resp, nil
}

type ApiDownloadReportRequest struct {
	ctx        context.Context
	ApiService *ReportsAPIService
	reportURL  string
}

func (r ApiDownloadReportRequest) Execute() ([]byte, error) {
	return r.ApiService.downloadReport(r.ctx, r.reportURL, 0)
}

// DownloadReport 下载报表内容并返回解压后的原始字节(行级业务字段由调用方按需反序列化)。
// reportURL 取自 GetReport 的 location/url。
func (a *ReportsAPIService) DownloadReport(ctx context.Context, reportURL string) ApiDownloadReportRequest {
	return ApiDownloadReportRequest{ApiService: a, ctx: ctx, reportURL: reportURL}
}

// downloadReport 跟随至多 3 跳重定向;presigned S3 链接改用不带鉴权的默认 client,
// 否则 S3 会因多余的 Authorization 头拒绝;响应为 gzip(magic 1f 8b)时自动解压。
func (a *ReportsAPIService) downloadReport(ctx context.Context, reportURL string, redirectDepth int) ([]byte, error) {
	if redirectDepth > 3 {
		return nil, reportError("too many SB v2 report redirects")
	}
	// location 可能是相对 endpoint 的路径,补全为绝对 URL 再请求。
	if strings.HasPrefix(reportURL, "/") {
		if base, baseErr := a.client.cfg.ServerURLWithContext(ctx, "ReportsAPIService.DownloadReport"); baseErr == nil {
			reportURL = strings.TrimRight(base, "/") + reportURL
		}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reportURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	client := a.client.cfg.HTTPClient
	if isPresignedReportURL(reportURL) {
		client = http.DefaultClient
	}
	resp, err := doReportRequest(client, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		next := resp.Header.Get("Location")
		if next == "" {
			return nil, reportError("SB v2 report redirect missing Location")
		}
		return a.downloadReport(ctx, resolveReportURL(reportURL, next), redirectDepth+1)
	}
	if resp.StatusCode >= 300 {
		return nil, &GenericOpenAPIError{body: data, error: resp.Status}
	}
	if len(data) >= 2 && data[0] == 0x1f && data[1] == 0x8b {
		gz, err := gzip.NewReader(bytes.NewReader(data))
		if err != nil {
			return nil, err
		}
		defer gz.Close()
		return io.ReadAll(gz)
	}
	return data, nil
}

// doRead 执行请求、读出 body 并把 Body 复位为可重复读;HTTP>=400 返回 GenericOpenAPIError 携带原始 body。
func (c *APIClient) doRead(req *http.Request) (*http.Response, []byte, error) {
	resp, err := c.callAPI(req)
	if err != nil || resp == nil {
		return resp, nil, err
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	resp.Body = io.NopCloser(bytes.NewBuffer(body))
	if err != nil {
		return resp, nil, err
	}
	if resp.StatusCode >= 400 {
		return resp, body, &GenericOpenAPIError{body: body, error: resp.Status}
	}
	return resp, body, nil
}

func lastURLSegment(raw string) string {
	if u, err := url.Parse(raw); err == nil {
		raw = u.Path
	}
	raw = strings.TrimRight(raw, "/")
	if i := strings.LastIndex(raw, "/"); i >= 0 {
		return raw[i+1:]
	}
	return raw
}

// isPresignedReportURL 判定是否 S3 预签名直链(带签名参数或指向 s3 amazonaws 域)。
func isPresignedReportURL(raw string) bool {
	u, err := url.Parse(raw)
	if err != nil {
		return false
	}
	if u.Query().Get("X-Amz-Algorithm") != "" || u.Query().Get("Signature") != "" {
		return true
	}
	host := strings.ToLower(u.Host)
	return strings.Contains(host, "amazonaws.com") && strings.Contains(host, "s3")
}

func resolveReportURL(baseURL, nextURL string) string {
	next, err := url.Parse(nextURL)
	if err != nil || next.IsAbs() {
		return nextURL
	}
	base, err := url.Parse(baseURL)
	if err != nil {
		return nextURL
	}
	return base.ResolveReference(next).String()
}

// doReportRequest 对带鉴权的 client 关闭自动重定向(手动跟随,便于切换 presigned client);
// 默认 client 直接执行。
func doReportRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	if client == http.DefaultClient {
		return client.Do(req)
	}
	noRedirect := *client
	noRedirect.CheckRedirect = func(*http.Request, []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return noRedirect.Do(req)
}
