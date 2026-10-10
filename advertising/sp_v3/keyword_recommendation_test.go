package sp_v3

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/4379711/amz-sdk/pkg"
)

// Captured from a live v5 KEYWORDS_FOR_ASINS response (US marketplace).
const rankedKeywordV5Response = `{"keywordTargetList":[{"keyword":"steamer for clothes","bidInfo":[
	{"matchType":"BROAD","theme":"CONVERSION_OPPORTUNITIES","rank":1,"bid":206,"suggestedBid":{"rangeStart":146,"rangeMedian":206,"rangeEnd":266}},
	{"matchType":"PHRASE","theme":"CONVERSION_OPPORTUNITIES","rank":1,"bid":206,"suggestedBid":{"rangeStart":145,"rangeMedian":206,"rangeEnd":272}},
	{"matchType":"EXACT","theme":"CONVERSION_OPPORTUNITIES","rank":1,"bid":191,"suggestedBid":{"rangeStart":132,"rangeMedian":191,"rangeEnd":216}}
],"translation":null,"userSelectedKeyword":false,"searchTermImpressionRank":3,"searchTermImpressionShare":6.59,"recId":"e83444d7-0e18-4357-9202-1a6f54bcf092"}],"impactMetrics":null}`

func TestGetRankedKeywordRecommendationV5(t *testing.T) {
	const mediaType = "application/vnd.spkeywordsrecommendation.v5+json"
	cfg := &pkg.Configuration{
		Servers: pkg.ServerConfigurations{{URL: "https://example.invalid"}},
		HTTPClient: &http.Client{Transport: numericIDTransport(func(r *http.Request) (*http.Response, error) {
			if r.URL.Path != "/sp/targets/keywords/recommendations" {
				t.Errorf("unexpected path %q", r.URL.Path)
			}
			if got := r.Header.Get("Content-Type"); got != mediaType {
				t.Errorf("Content-Type = %q, want %q", got, mediaType)
			}
			if got := r.Header.Get("Accept"); got != mediaType {
				t.Errorf("Accept = %q, want %q", got, mediaType)
			}
			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			var got, want interface{}
			if err := json.Unmarshal(body, &got); err != nil {
				t.Fatal(err)
			}
			_ = json.Unmarshal([]byte(`{"asins":["B0CSCNYKR3"],"recommendationType":"KEYWORDS_FOR_ASINS","maxRecommendations":5}`), &want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("request body = %s", body)
			}
			return &http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": {mediaType}}, Body: io.NopCloser(strings.NewReader(rankedKeywordV5Response)), Request: r}, nil
		})},
	}
	max := float32(5)
	req := RankedKeywordTargetsForAsinsRequest{Asins: []string{"B0CSCNYKR3"}, RecommendationType: "KEYWORDS_FOR_ASINS", MaxRecommendations: &max}
	result, _, err := NewAPIClient(cfg).KeywordTargetsAPI.GetRankedKeywordRecommendation(context.Background()).
		GetRankedKeywordRecommendationRequest(RankedKeywordTargetsForAsinsRequestAsGetRankedKeywordRecommendationRequest1(&req)).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if result == nil || len(result.KeywordTargetList) != 1 || len(result.KeywordTargetList[0].BidInfo) != 3 {
		t.Fatalf("unexpected result: %+v", result)
	}
	exact := result.KeywordTargetList[0].BidInfo[2]
	if exact.GetMatchType() != "EXACT" || exact.GetTheme() != "CONVERSION_OPPORTUNITIES" || exact.GetBid() != 191 {
		t.Fatalf("unexpected bid info: %+v", exact)
	}
	s := exact.SuggestedBid
	if s.GetRangeStart() != 132 || s.GetRangeMedian() != 191 || s.GetRangeEnd() != 216 {
		t.Fatalf("unexpected suggested bid: %+v", s)
	}
}
