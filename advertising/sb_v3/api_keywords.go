package sb_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// KeywordsAPIService 封装 SB legacy 关键词资源 /sb/keywords(list/create/update)。
type KeywordsAPIService service

// keywordMediaTypeV3 是 SB 关键词资源的版本化媒体类型,Content-Type 与 Accept 均使用它。
const keywordMediaTypeV3 = "application/vnd.sbkeyword.v3+json"

type ApiListKeywordsRequest struct {
	ctx              context.Context
	ApiService       *KeywordsAPIService
	startIndex       *int32
	count            *int32
	campaignIdFilter *string
	adGroupIdFilter  *string
	stateFilter      *string
}

func (r ApiListKeywordsRequest) StartIndex(startIndex int32) ApiListKeywordsRequest {
	r.startIndex = &startIndex
	return r
}

func (r ApiListKeywordsRequest) Count(count int32) ApiListKeywordsRequest {
	r.count = &count
	return r
}

func (r ApiListKeywordsRequest) CampaignIdFilter(campaignIdFilter string) ApiListKeywordsRequest {
	r.campaignIdFilter = &campaignIdFilter
	return r
}

func (r ApiListKeywordsRequest) AdGroupIdFilter(adGroupIdFilter string) ApiListKeywordsRequest {
	r.adGroupIdFilter = &adGroupIdFilter
	return r
}

func (r ApiListKeywordsRequest) StateFilter(stateFilter string) ApiListKeywordsRequest {
	r.stateFilter = &stateFilter
	return r
}

func (r ApiListKeywordsRequest) Execute() ([]Keyword, *http.Response, error) {
	return r.ApiService.ListKeywordsExecute(r)
}

// ListKeywords 分页拉取 SB 关键词;分页用 startIndex/count(SB legacy 无 nextToken)。
func (a *KeywordsAPIService) ListKeywords(ctx context.Context) ApiListKeywordsRequest {
	return ApiListKeywordsRequest{ApiService: a, ctx: ctx}
}

func (a *KeywordsAPIService) ListKeywordsExecute(r ApiListKeywordsRequest) ([]Keyword, *http.Response, error) {
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeywordsAPIService.ListKeywords")
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	localVarPath := localBasePath + "/sb/keywords"

	localVarHeaderParams := map[string]string{"Accept": keywordMediaTypeV3 + ", application/json"}
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
	keywords, err := decodeKeywords(body)
	if err != nil {
		return nil, resp, &GenericOpenAPIError{body: body, error: err.Error()}
	}
	return keywords, resp, nil
}

type ApiCreateKeywordsRequest struct {
	ctx        context.Context
	ApiService *KeywordsAPIService
	keywords   []CreateKeyword
}

func (r ApiCreateKeywordsRequest) Keywords(keywords []CreateKeyword) ApiCreateKeywordsRequest {
	r.keywords = keywords
	return r
}

func (r ApiCreateKeywordsRequest) Execute() (*MutationResult, *http.Response, error) {
	return r.ApiService.CreateKeywordsExecute(r)
}

func (a *KeywordsAPIService) CreateKeywords(ctx context.Context) ApiCreateKeywordsRequest {
	return ApiCreateKeywordsRequest{ApiService: a, ctx: ctx}
}

func (a *KeywordsAPIService) CreateKeywordsExecute(r ApiCreateKeywordsRequest) (*MutationResult, *http.Response, error) {
	return a.writeKeywords(r.ctx, "KeywordsAPIService.CreateKeywords", http.MethodPost, r.keywords)
}

type ApiUpdateKeywordsRequest struct {
	ctx        context.Context
	ApiService *KeywordsAPIService
	keywords   []UpdateKeyword
}

func (r ApiUpdateKeywordsRequest) Keywords(keywords []UpdateKeyword) ApiUpdateKeywordsRequest {
	r.keywords = keywords
	return r
}

func (r ApiUpdateKeywordsRequest) Execute() (*MutationResult, *http.Response, error) {
	return r.ApiService.UpdateKeywordsExecute(r)
}

func (a *KeywordsAPIService) UpdateKeywords(ctx context.Context) ApiUpdateKeywordsRequest {
	return ApiUpdateKeywordsRequest{ApiService: a, ctx: ctx}
}

func (a *KeywordsAPIService) UpdateKeywordsExecute(r ApiUpdateKeywordsRequest) (*MutationResult, *http.Response, error) {
	return a.writeKeywords(r.ctx, "KeywordsAPIService.UpdateKeywords", http.MethodPut, r.keywords)
}

func (a *KeywordsAPIService) writeKeywords(ctx context.Context, opID, method string, body any) (*MutationResult, *http.Response, error) {
	localBasePath, err := a.client.cfg.ServerURLWithContext(ctx, opID)
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	if body == nil {
		return nil, nil, reportError("keywords is required and must be specified")
	}
	localVarPath := localBasePath + "/sb/keywords"
	localVarHeaderParams := map[string]string{
		"Content-Type": keywordMediaTypeV3,
		"Accept":       keywordMediaTypeV3 + ", application/json",
	}
	req, err := a.client.prepareRequest(ctx, localVarPath, method, body, localVarHeaderParams, url.Values{}, url.Values{}, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, raw, err := a.client.doRead(req)
	if err != nil {
		return nil, resp, err
	}
	return &MutationResult{Body: raw}, resp, nil
}

// doRead 执行请求、读出 body 并把 Body 复位为可重复读;HTTP>=400 时返回 GenericOpenAPIError 并携带原始 body。
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
