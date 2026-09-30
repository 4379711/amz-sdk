package sd_v1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/4379711/amz-sdk/pkg"
	"github.com/bytedance/sonic"
)

func TestNumericEntityIDsRoundTrip(t *testing.T) {
	properties := CreativeProperties{HeadlineCreativeProperties: &HeadlineCreativeProperties{Headline: pkg.Ptr("Example")}}
	for _, tc := range []struct {
		name   string
		create func(int64) interface{}
		fields []string
	}{
		{"campaign", func(int64) interface{} { return &CampaignResponseEx{} }, []string{"CampaignId"}},
		{"ad_group", func(int64) interface{} { return &AdGroupResponseEx{} }, []string{"AdGroupId", "CampaignId"}},
		{"product_ad_ex", func(int64) interface{} { return &ProductAdResponseEx{} }, []string{"AdId", "AdGroupId", "CampaignId"}},
		{"product_ad", func(int64) interface{} { return &ProductAdResponse{} }, []string{"AdId"}},
		{"target", func(int64) interface{} { return &TargetingClauseEx{} }, []string{"TargetId", "AdGroupId", "CampaignId"}},
		{"negative_target", func(int64) interface{} { return &NegativeTargetingClauseEx{} }, []string{"TargetId", "AdGroupId"}},
		{"creative", func(id int64) interface{} { return NewCreative(id, id, "IMAGE", properties, "APPROVED") }, []string{"CreativeId"}},
		{"creative_update", func(id int64) interface{} { return NewCreativeUpdate(id, properties) }, []string{"CreativeId"}},
		{"creative_response", func(int64) interface{} { return &CreativeResponse{} }, []string{"CreativeId"}},
		{"create_creative", func(id int64) interface{} { return NewCreateCreative(id, properties) }, []string{"AdGroupId"}},
		{"creative_moderation", func(id int64) interface{} {
			return NewCreativeModeration(id, "IMAGE", "APPROVED", time.Unix(0, 0).UTC(), nil)
		}, []string{"CreativeId"}},
	} {
		for _, id := range []int64{16777217, 36027908138365, 9223372036854775807} {
			t.Run(tc.name+"/"+strconv.FormatInt(id, 10), func(t *testing.T) {
				model := tc.create(id)
				input := make(map[string]int64)
				for _, field := range tc.fields {
					input[strings.ToLower(field[:1])+field[1:]] = id
				}
				data, err := json.Marshal(input)
				if err != nil {
					t.Fatal(err)
				}
				if err := sonic.Unmarshal(data, model); err != nil {
					t.Fatal(err)
				}
				for _, field := range tc.fields {
					value := reflect.ValueOf(model)
					got := value.MethodByName("Get" + field).Call(nil)[0]
					if got.Kind() != reflect.Int64 || got.Int() != id {
						t.Fatalf("%s = %v (%s), want int64 %d", field, got, got.Kind(), id)
					}
					value.MethodByName("Set" + field).Call([]reflect.Value{reflect.ValueOf(id)})
					ok := value.MethodByName("Get" + field + "Ok").Call(nil)
					if !ok[1].Bool() || ok[0].Elem().Int() != id {
						t.Fatalf("%s accessor lost ID %d", field, id)
					}
				}
				data, err = sonic.Marshal(model)
				if err != nil {
					t.Fatal(err)
				}
				var output map[string]json.RawMessage
				if err := json.Unmarshal(data, &output); err != nil {
					t.Fatal(err)
				}
				for key := range input {
					if string(output[key]) != strconv.FormatInt(id, 10) {
						t.Errorf("JSON %s = %s, want decimal number %d", key, output[key], id)
					}
				}
			})
		}
	}
}

type numericIDTransport func(*http.Request) (*http.Response, error)

func (f numericIDTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestListProductAdsExPreservesIDs(t *testing.T) {
	cfg := &pkg.Configuration{
		Servers: pkg.ServerConfigurations{{URL: "https://example.invalid"}},
		HTTPClient: &http.Client{Transport: numericIDTransport(func(r *http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": {"application/json"}}, Body: io.NopCloser(strings.NewReader(`[{"adId":36027908138365,"adGroupId":16777217,"campaignId":36027908138365}]`)), Request: r}, nil
		})},
	}
	ads, _, err := NewAPIClient(cfg).ProductAdsAPI.ListProductAdsEx(context.Background()).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if len(ads) != 1 || ads[0].GetAdId() != 36027908138365 || ads[0].GetAdGroupId() != 16777217 || ads[0].GetCampaignId() != 36027908138365 {
		t.Fatalf("response IDs changed: %+v", ads)
	}
}
