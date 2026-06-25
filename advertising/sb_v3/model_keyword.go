package sb_v3

import (
	"encoding/json"
	"strings"

	"github.com/bytedance/sonic"
)

// StringID 兼容 SB legacy 接口里 id 字段可能以 JSON 数字或字符串两种形态出现,统一以字符串承载;
// 解析时直接取字面量,避免大整数 id 经数值类型退化成科学计数法。
type StringID string

func (id *StringID) UnmarshalJSON(src []byte) error {
	value := strings.TrimSpace(string(src))
	if value == "" || value == "null" {
		*id = ""
		return nil
	}
	if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
		var s string
		if err := sonic.Unmarshal(src, &s); err != nil {
			return err
		}
		*id = StringID(s)
		return nil
	}
	*id = StringID(value)
	return nil
}

func (id StringID) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(string(id))
}

func (id StringID) String() string {
	return string(id)
}

// Keyword 是 GET /sb/keywords 的列表项。SB legacy 响应字段形态不稳,保留 ExtendedData 原文供上层按需解析。
type Keyword struct {
	CampaignId   StringID        `json:"campaignId"`
	AdGroupId    StringID        `json:"adGroupId"`
	KeywordId    StringID        `json:"keywordId"`
	KeywordText  string          `json:"keywordText"`
	MatchType    string          `json:"matchType"`
	State        string          `json:"state"`
	Bid          float64         `json:"bid"`
	ExtendedData json.RawMessage `json:"extendedData"`
}

// CreateKeyword 是 POST /sb/keywords 的单条入参。
type CreateKeyword struct {
	CampaignId  string   `json:"campaignId"`
	AdGroupId   string   `json:"adGroupId"`
	KeywordText string   `json:"keywordText"`
	MatchType   string   `json:"matchType"`
	Bid         *float64 `json:"bid,omitempty"`
}

// UpdateKeyword 是 PUT /sb/keywords 的单条入参。
type UpdateKeyword struct {
	KeywordId string   `json:"keywordId"`
	State     string   `json:"state,omitempty"`
	Bid       *float64 `json:"bid,omitempty"`
}

// keywordListResponse 兼容 list 接口返回裸数组或 {keywords:[...]} 两种形态。
type keywordListResponse struct {
	Keywords []Keyword `json:"keywords"`
}

func decodeKeywords(body []byte) ([]Keyword, error) {
	trimmed := strings.TrimSpace(string(body))
	if trimmed == "" {
		return nil, nil
	}
	if strings.HasPrefix(trimmed, "[") {
		var direct []Keyword
		if err := sonic.Unmarshal(body, &direct); err != nil {
			return nil, err
		}
		return direct, nil
	}
	var wrapped keywordListResponse
	if err := sonic.Unmarshal(body, &wrapped); err != nil {
		return nil, err
	}
	return wrapped.Keywords, nil
}
