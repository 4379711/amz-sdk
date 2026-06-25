package sb_v3

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

// MutationResult 承载 SB legacy keyword / negativeKeyword 写接口(create/update)的原始返回。
// 这些接口在不同站点会返回数组、{keywords:[...]} 包装或带 errors 的对象,形态不稳定,
// 故保留原文并按需提取,不强行映射成固定结构。
type MutationResult struct {
	Body []byte
}

// FirstID 递归查找首个名为 field 的标量值(如 keywordId)并以字符串返回,找不到返回空串。
// 解析全程走 json.Number,避免大整数 id 经 float64 退化成科学计数法而丢精度。
func (m MutationResult) FirstID(field string) string {
	v, ok := m.decode()
	if !ok {
		return ""
	}
	return firstFieldString(v, field)
}

// ErrorMessage 在返回体含非空 errors/error 时提取可读错误说明;无错误信息时返回空串,
// 调用方据此判断写操作是否成功。
func (m MutationResult) ErrorMessage() string {
	v, ok := m.decode()
	if !ok {
		return ""
	}
	if !hasNonEmptyField(v, "errors") && !hasNonEmptyField(v, "error") {
		return ""
	}
	for _, key := range []string{"message", "reason", "details", "code"} {
		if msg := firstFieldString(v, key); msg != "" {
			return msg
		}
	}
	return ""
}

func (m MutationResult) decode() (any, bool) {
	if len(bytes.TrimSpace(m.Body)) == 0 {
		return nil, false
	}
	dec := json.NewDecoder(bytes.NewReader(m.Body))
	dec.UseNumber()
	var v any
	if err := dec.Decode(&v); err != nil {
		return nil, false
	}
	return v, true
}

func firstFieldString(v any, field string) string {
	switch obj := v.(type) {
	case map[string]any:
		if raw, ok := obj[field]; ok {
			if s := scalarString(raw); s != "" {
				return s
			}
		}
		for _, child := range obj {
			if s := firstFieldString(child, field); s != "" {
				return s
			}
		}
	case []any:
		for _, child := range obj {
			if s := firstFieldString(child, field); s != "" {
				return s
			}
		}
	}
	return ""
}

func hasNonEmptyField(v any, field string) bool {
	switch obj := v.(type) {
	case map[string]any:
		if raw, ok := obj[field]; ok && isNonEmptyValue(raw) {
			return true
		}
		for _, child := range obj {
			if hasNonEmptyField(child, field) {
				return true
			}
		}
	case []any:
		for _, child := range obj {
			if hasNonEmptyField(child, field) {
				return true
			}
		}
	}
	return false
}

func isNonEmptyValue(v any) bool {
	switch val := v.(type) {
	case nil:
		return false
	case []any:
		return len(val) > 0
	case map[string]any:
		return len(val) > 0
	case string:
		return strings.TrimSpace(val) != ""
	default:
		return true
	}
}

func scalarString(v any) string {
	switch val := v.(type) {
	case string:
		return strings.TrimSpace(val)
	case json.Number:
		return val.String()
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(val)
	}
	return ""
}
