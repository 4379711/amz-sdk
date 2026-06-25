package sb_v3

import (
	"context"
	"net/http"
	"net/url"

	"github.com/bytedance/sonic"
)

// RecommendationsAPIService 封装 SB legacy 竞价推荐资源 /sb/recommendations/bids。
type RecommendationsAPIService service

// bidsRecommendationMediaTypeV3 是竞价推荐接口的版本化响应媒体类型,仅用于 Accept 指定响应版本;
// 请求体走普通 application/json(见 GetBidRecommendationsExecute)。
const bidsRecommendationMediaTypeV3 = "application/vnd.sbbidsrecommendation.v3+json"

type ApiGetBidRecommendationsRequest struct {
	ctx        context.Context
	ApiService *RecommendationsAPIService
	content    *GetBidRecommendationsRequestContent
}

func (r ApiGetBidRecommendationsRequest) GetBidRecommendationsRequestContent(content GetBidRecommendationsRequestContent) ApiGetBidRecommendationsRequest {
	r.content = &content
	return r
}

func (r ApiGetBidRecommendationsRequest) Execute() (*GetBidRecommendationsResponseContent, *http.Response, error) {
	return r.ApiService.GetBidRecommendationsExecute(r)
}

// GetBidRecommendations 获取关键词/投放/主题的竞价推荐;按请求体内放置的目标类型决定返回哪类推荐桶。
func (a *RecommendationsAPIService) GetBidRecommendations(ctx context.Context) ApiGetBidRecommendationsRequest {
	return ApiGetBidRecommendationsRequest{ApiService: a, ctx: ctx}
}

func (a *RecommendationsAPIService) GetBidRecommendationsExecute(r ApiGetBidRecommendationsRequest) (*GetBidRecommendationsResponseContent, *http.Response, error) {
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecommendationsAPIService.GetBidRecommendations")
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	if r.content == nil {
		return nil, nil, reportError("getBidRecommendationsRequestContent is required and must be specified")
	}
	localVarPath := localBasePath + "/sb/recommendations/bids"
	// 请求体是普通 JSON,Content-Type 必须为 application/json;版本化媒体类型只放 Accept。
	// 两者写反会被服务端以 415 "Cannot consume content type" 拒绝。
	localVarHeaderParams := map[string]string{
		"Content-Type": "application/json",
		"Accept":       bidsRecommendationMediaTypeV3,
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, http.MethodPost, r.content, localVarHeaderParams, url.Values{}, url.Values{}, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, body, err := a.client.doRead(req)
	if err != nil {
		return nil, resp, err
	}
	var out GetBidRecommendationsResponseContent
	if err := sonic.Unmarshal(body, &out); err != nil {
		return nil, resp, &GenericOpenAPIError{body: body, error: err.Error()}
	}
	return &out, resp, nil
}
