package sb_v2

import (
	"strings"

	"github.com/bytedance/sonic"
)

// RequestReportContent 是 POST /v2/hsa/campaigns/report 请求体。
// ReportDate 用 yyyyMMdd;Metrics 在请求体里序列化为逗号分隔字符串。
type RequestReportContent struct {
	ReportDate string
	Metrics    []string
}

func (c RequestReportContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(map[string]string{
		"reportDate": c.ReportDate,
		"metrics":    strings.Join(c.Metrics, ","),
	})
}

// ReportResponse 兼容创建与状态查询两个接口的返回:
// 创建主要给出 reportId;状态查询给出 status/statusDetails,报表就绪后给出 location 或 url。
type ReportResponse struct {
	ReportId      string `json:"reportId"`
	Status        string `json:"status"`
	StatusDetails string `json:"statusDetails"`
	Location      string `json:"location"`
	Url           string `json:"url"`
}
