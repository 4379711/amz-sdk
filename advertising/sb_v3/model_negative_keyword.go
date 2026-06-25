package sb_v3

import (
	"encoding/json"
	"strings"

	"github.com/bytedance/sonic"
)

// NegativeKeyword 是 GET /sb/negativeKeywords 的列表项。id 字段统一用 StringID 兼容数字/字符串两种形态。
type NegativeKeyword struct {
	CampaignId   StringID        `json:"campaignId"`
	AdGroupId    StringID        `json:"adGroupId"`
	KeywordId    StringID        `json:"keywordId"`
	KeywordText  string          `json:"keywordText"`
	MatchType    string          `json:"matchType"`
	State        string          `json:"state"`
	ExtendedData json.RawMessage `json:"extendedData"`
}

// CreateNegativeKeyword 是 POST /sb/negativeKeywords 的单条入参;AdGroupId 为空表示活动级否定词。
type CreateNegativeKeyword struct {
	CampaignId  string `json:"campaignId"`
	AdGroupId   string `json:"adGroupId,omitempty"`
	KeywordText string `json:"keywordText"`
	MatchType   string `json:"matchType"`
	State       string `json:"state,omitempty"`
}

// negativeKeywordListResponse 兼容 list 接口返回裸数组或 {negativeKeywords:[...]} 两种形态。
type negativeKeywordListResponse struct {
	NegativeKeywords []NegativeKeyword `json:"negativeKeywords"`
}

func decodeNegativeKeywords(body []byte) ([]NegativeKeyword, error) {
	trimmed := strings.TrimSpace(string(body))
	if trimmed == "" {
		return nil, nil
	}
	if strings.HasPrefix(trimmed, "[") {
		var direct []NegativeKeyword
		if err := sonic.Unmarshal(body, &direct); err != nil {
			return nil, err
		}
		return direct, nil
	}
	var wrapped negativeKeywordListResponse
	if err := sonic.Unmarshal(body, &wrapped); err != nil {
		return nil, err
	}
	return wrapped.NegativeKeywords, nil
}
