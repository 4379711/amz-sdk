package sb_v3

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

// NegativeKeywordsAPIService 封装 SB legacy 否定关键词资源 /sb/negativeKeywords(list/create)。
type NegativeKeywordsAPIService service

// negativeKeywordMediaTypeV3 是 SB 否定关键词资源的版本化媒体类型。
const negativeKeywordMediaTypeV3 = "application/vnd.sbnegativekeyword.v3+json"

type ApiListNegativeKeywordsRequest struct {
	ctx              context.Context
	ApiService       *NegativeKeywordsAPIService
	startIndex       *int32
	count            *int32
	campaignIdFilter *string
	adGroupIdFilter  *string
	stateFilter      *string
}

func (r ApiListNegativeKeywordsRequest) StartIndex(startIndex int32) ApiListNegativeKeywordsRequest {
	r.startIndex = &startIndex
	return r
}

func (r ApiListNegativeKeywordsRequest) Count(count int32) ApiListNegativeKeywordsRequest {
	r.count = &count
	return r
}

func (r ApiListNegativeKeywordsRequest) CampaignIdFilter(campaignIdFilter string) ApiListNegativeKeywordsRequest {
	r.campaignIdFilter = &campaignIdFilter
	return r
}

func (r ApiListNegativeKeywordsRequest) AdGroupIdFilter(adGroupIdFilter string) ApiListNegativeKeywordsRequest {
	r.adGroupIdFilter = &adGroupIdFilter
	return r
}

func (r ApiListNegativeKeywordsRequest) StateFilter(stateFilter string) ApiListNegativeKeywordsRequest {
	r.stateFilter = &stateFilter
	return r
}

func (r ApiListNegativeKeywordsRequest) Execute() ([]NegativeKeyword, *http.Response, error) {
	return r.ApiService.ListNegativeKeywordsExecute(r)
}

func (a *NegativeKeywordsAPIService) ListNegativeKeywords(ctx context.Context) ApiListNegativeKeywordsRequest {
	return ApiListNegativeKeywordsRequest{ApiService: a, ctx: ctx}
}

func (a *NegativeKeywordsAPIService) ListNegativeKeywordsExecute(r ApiListNegativeKeywordsRequest) ([]NegativeKeyword, *http.Response, error) {
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegativeKeywordsAPIService.ListNegativeKeywords")
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	localVarPath := localBasePath + "/sb/negativeKeywords"

	localVarHeaderParams := map[string]string{"Accept": negativeKeywordMediaTypeV3 + ", application/json"}
	localVarQueryParams := url.Values{}
	if r.startIndex != nil {
		localVarQueryParams.Set("startIndex", strconv.FormatInt(int64(*r.startIndex), 10))
	}
	if r.count != nil {
		localVarQueryParams.Set("count", strconv.FormatInt(int64(*r.count), 10))
	}
	if r.campaignIdFilter != nil {
		localVarQueryParams.Set("campaignIdFilter", *r.campaignIdFilter)
	}
	if r.adGroupIdFilter != nil {
		localVarQueryParams.Set("adGroupIdFilter", *r.adGroupIdFilter)
	}
	if r.stateFilter != nil {
		localVarQueryParams.Set("stateFilter", *r.stateFilter)
	}

	req, err := a.client.prepareRequest(r.ctx, localVarPath, http.MethodGet, nil, localVarHeaderParams, localVarQueryParams, url.Values{}, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, body, err := a.client.doRead(req)
	if err != nil {
		return nil, resp, err
	}
	items, err := decodeNegativeKeywords(body)
	if err != nil {
		return nil, resp, &GenericOpenAPIError{body: body, error: err.Error()}
	}
	return items, resp, nil
}

type ApiCreateNegativeKeywordsRequest struct {
	ctx          context.Context
	ApiService   *NegativeKeywordsAPIService
	negativeKeys []CreateNegativeKeyword
}

func (r ApiCreateNegativeKeywordsRequest) NegativeKeywords(negativeKeywords []CreateNegativeKeyword) ApiCreateNegativeKeywordsRequest {
	r.negativeKeys = negativeKeywords
	return r
}

func (r ApiCreateNegativeKeywordsRequest) Execute() (*MutationResult, *http.Response, error) {
	return r.ApiService.CreateNegativeKeywordsExecute(r)
}

func (a *NegativeKeywordsAPIService) CreateNegativeKeywords(ctx context.Context) ApiCreateNegativeKeywordsRequest {
	return ApiCreateNegativeKeywordsRequest{ApiService: a, ctx: ctx}
}

func (a *NegativeKeywordsAPIService) CreateNegativeKeywordsExecute(r ApiCreateNegativeKeywordsRequest) (*MutationResult, *http.Response, error) {
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegativeKeywordsAPIService.CreateNegativeKeywords")
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	if r.negativeKeys == nil {
		return nil, nil, reportError("negativeKeywords is required and must be specified")
	}
	localVarPath := localBasePath + "/sb/negativeKeywords"
	localVarHeaderParams := map[string]string{
		"Content-Type": negativeKeywordMediaTypeV3,
		"Accept":       negativeKeywordMediaTypeV3 + ", application/json",
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, http.MethodPost, r.negativeKeys, localVarHeaderParams, url.Values{}, url.Values{}, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, raw, err := a.client.doRead(req)
	if err != nil {
		return nil, resp, err
	}
	return &MutationResult{Body: raw}, resp, nil
}
