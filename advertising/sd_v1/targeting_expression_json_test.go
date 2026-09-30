package sd_v1_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/4379711/amz-sdk/advertising/sd_v1"
	"github.com/4379711/amz-sdk/pkg"
	"github.com/bytedance/sonic"
)

type expressionJSON interface {
	UnmarshalJSON([]byte) error
	MarshalJSON() ([]byte, error)
	GetActualInstance() interface{}
}

var expressionFactories = []struct {
	name string
	new  func() expressionJSON
}{
	{"targeting", func() expressionJSON { return &sd_v1.TargetingExpressionInner{} }},
	{"v32", func() expressionJSON { return &sd_v1.SDTargetExpressionV32{} }},
}

func assertExpressionJSON(t *testing.T, actual []byte, expected string) {
	t.Helper()
	decode := func(raw []byte) any {
		decoder := json.NewDecoder(bytes.NewReader(raw))
		decoder.UseNumber()
		var value any
		if err := decoder.Decode(&value); err != nil {
			t.Fatalf("decode %s: %v", raw, err)
		}
		return value
	}
	if !reflect.DeepEqual(decode(actual), decode([]byte(expected))) {
		t.Fatalf("JSON mismatch:\n got %s\nwant %s", actual, expected)
	}
}

func marshalExpression(t *testing.T, value any) []byte {
	t.Helper()
	data, err := sonic.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	standard, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	assertExpressionJSON(t, standard, string(data))
	return data
}

func TestTargetingExpressionScalarEdits(t *testing.T) {
	for _, factory := range expressionFactories {
		t.Run(factory.name, func(t *testing.T) {
			for _, predicateType := range []string{"asinSameAs", "contentCategorySameAs", "futurePredicate"} {
				t.Run(predicateType, func(t *testing.T) {
					expr := factory.new()
					raw := `{"type":"` + predicateType + `","value":"old","extension":{"id":9007199254740993},"nullable":null}`
					if err := sonic.Unmarshal([]byte(raw), expr); err != nil {
						t.Fatal(err)
					}
					assertExpressionJSON(t, marshalExpression(t, expr), raw)
					expr.GetActualInstance().(interface{ SetValue(string) }).SetValue("new")
					expected := strings.Replace(raw, `"old"`, `"new"`, 1)
					assertExpressionJSON(t, marshalExpression(t, expr), expected)
					assertExpressionJSON(t, marshalExpression(t, expr), expected)
					expr.GetActualInstance().(interface{ SetValue(string) }).SetValue("old")
					assertExpressionJSON(t, marshalExpression(t, expr), raw)
				})
			}
		})
	}
}

func TestTargetingExpressionSelectsVariantAndClearsOptionalFields(t *testing.T) {
	var legacy sd_v1.TargetingExpressionInner
	if err := sonic.Unmarshal([]byte(`{"type":"exactProduct","value":"old","eventType":"views","extension":true}`), &legacy); err != nil {
		t.Fatal(err)
	}
	if legacy.TargetingPredicateLegacy == nil || legacy.ContentTargetingPredicate != nil {
		t.Fatalf("wrong legacy variant: %T", legacy.GetActualInstance())
	}
	legacy.TargetingPredicateLegacy.Value = nil
	assertExpressionJSON(t, marshalExpression(t, legacy), `{"type":"exactProduct","eventType":"views","extension":true}`)
	legacy.TargetingPredicateLegacy.EventType = nil
	assertExpressionJSON(t, marshalExpression(t, legacy), `{"type":"exactProduct","extension":true}`)

	var scalar sd_v1.SDTargetExpressionV32
	if err := sonic.Unmarshal([]byte(`{"type":"asinSameAs","value":"old","extension":true}`), &scalar); err != nil {
		t.Fatal(err)
	}
	if scalar.SDTargetingPredicateV31 == nil || scalar.SDContentTargetingPredicateV31 != nil {
		t.Fatalf("wrong scalar variant: %T", scalar.GetActualInstance())
	}
	scalar.SDTargetingPredicateV31.Value = nil
	assertExpressionJSON(t, marshalExpression(t, scalar), `{"type":"asinSameAs","extension":true}`)

	var content sd_v1.SDTargetExpressionV32
	if err := sonic.Unmarshal([]byte(`{"type":"contentCategorySameAs","value":"old"}`), &content); err != nil {
		t.Fatal(err)
	}
	if content.SDContentTargetingPredicateV31 == nil {
		t.Fatalf("wrong content variant: %T", content.GetActualInstance())
	}
}

func TestTargetingExpressionNestedEditsAndReordering(t *testing.T) {
	const raw = `{"type":"views","value":[{"type":"exactProduct","firstExtension":{"id":9007199254740993}},{"type":"lookback","value":"30","secondExtension":"keep"}],"outerExtension":true}`
	for _, factory := range expressionFactories {
		t.Run(factory.name, func(t *testing.T) {
			expr := factory.new()
			if err := sonic.Unmarshal([]byte(raw), expr); err != nil {
				t.Fatal(err)
			}
			assertExpressionJSON(t, marshalExpression(t, expr), raw)
			switch nested := expr.GetActualInstance().(type) {
			case *sd_v1.TargetingPredicateNested:
				nested.Value[0], nested.Value[1] = nested.Value[1], nested.Value[0]
				nested.Value[0].SetValue("14")
				nested.Value = append(nested.Value, sd_v1.TargetingPredicateBase{Type: pkg.Ptr("futureChild"), Value: pkg.Ptr("new")})
			case *sd_v1.SDTargetingPredicateNestedV31:
				nested.Value[0], nested.Value[1] = nested.Value[1], nested.Value[0]
				nested.Value[0].SetValue("14")
				nested.Value = append(nested.Value, sd_v1.SDTargetingPredicateBaseV31{Type: "futureChild", Value: pkg.Ptr("new")})
			default:
				t.Fatalf("wrong nested variant: %T", nested)
			}
			assertExpressionJSON(t, marshalExpression(t, expr), `{"type":"views","value":[{"type":"lookback","value":"14","secondExtension":"keep"},{"type":"exactProduct","firstExtension":{"id":9007199254740993}},{"type":"futureChild","value":"new"}],"outerExtension":true}`)
			switch nested := expr.GetActualInstance().(type) {
			case *sd_v1.TargetingPredicateNested:
				nested.Value[0].Value = nil
				nested.Value = nested.Value[:1]
			case *sd_v1.SDTargetingPredicateNestedV31:
				nested.Value[0].Value = nil
				nested.Value = nested.Value[:1]
			}
			assertExpressionJSON(t, marshalExpression(t, expr), `{"type":"views","value":[{"type":"lookback","secondExtension":"keep"}],"outerExtension":true}`)
		})
	}
}

func TestTargetingExpressionReplacementAndReuse(t *testing.T) {
	var expr sd_v1.TargetingExpressionInner
	if err := sonic.Unmarshal([]byte(`{"type":"exactProduct","value":"old","eventType":"views","extension":true}`), &expr); err != nil {
		t.Fatal(err)
	}
	expr.TargetingPredicateLegacy = nil
	expr.TargetingPredicateNested = &sd_v1.TargetingPredicateNested{Type: pkg.Ptr("views"), Value: []sd_v1.TargetingPredicateBase{{Type: pkg.Ptr("lookback"), Value: pkg.Ptr("30")}}}
	assertExpressionJSON(t, marshalExpression(t, expr), `{"type":"views","value":[{"type":"lookback","value":"30"}],"extension":true}`)
	if err := sonic.Unmarshal([]byte(`{"type":"asinSameAs","value":"replacement"}`), &expr); err != nil {
		t.Fatal(err)
	}
	if expr.TargetingPredicateNested != nil || expr.TargetingPredicate == nil {
		t.Fatal("decoding retained an obsolete branch")
	}
	assertExpressionJSON(t, marshalExpression(t, expr), `{"type":"asinSameAs","value":"replacement"}`)
	before := marshalExpression(t, expr)
	if err := sonic.Unmarshal([]byte(`{"type":"asinSameAs","value":123}`), &expr); err == nil {
		t.Fatal("numeric value must fail")
	}
	assertExpressionJSON(t, marshalExpression(t, expr), string(before))
}

func TestTargetingExpressionNewValuesAndClearingNestedValue(t *testing.T) {
	newScalar := sd_v1.TargetingPredicateAsTargetingExpressionInner(&sd_v1.TargetingPredicate{
		Type: pkg.Ptr("asinSameAs"), Value: pkg.Ptr("new"),
	})
	assertExpressionJSON(t, marshalExpression(t, newScalar), `{"type":"asinSameAs","value":"new"}`)
	newV32 := sd_v1.SDContentTargetingPredicateV31AsSDTargetExpressionV32(sd_v1.NewSDContentTargetingPredicateV31("contentCategorySameAs", "new"))
	assertExpressionJSON(t, marshalExpression(t, newV32), `{"type":"contentCategorySameAs","value":"new"}`)

	const raw = `{"type":"views","value":[{"type":"lookback","value":"30","extension":true}],"outerExtension":true}`
	var optional sd_v1.TargetingExpressionInner
	if err := sonic.Unmarshal([]byte(raw), &optional); err != nil {
		t.Fatal(err)
	}
	copyOfChild := optional.TargetingPredicateNested.Value[0]
	copyOfChild.SetValue("14")
	assertExpressionJSON(t, marshalExpression(t, copyOfChild), `{"type":"lookback","value":"14","extension":true}`)
	assertExpressionJSON(t, marshalExpression(t, optional), raw)
	optional.TargetingPredicateNested.Value = nil
	assertExpressionJSON(t, marshalExpression(t, optional), `{"type":"views","outerExtension":true}`)
	if err := sonic.Unmarshal([]byte(`{"type":"views","value":[],"extension":true}`), &optional); err != nil {
		t.Fatal(err)
	}
	assertExpressionJSON(t, marshalExpression(t, optional), `{"type":"views","value":[],"extension":true}`)
	optional.TargetingPredicateNested.Value = nil
	assertExpressionJSON(t, marshalExpression(t, optional), `{"type":"views","extension":true}`)

	var required sd_v1.SDTargetExpressionV32
	if err := sonic.Unmarshal([]byte(raw), &required); err != nil {
		t.Fatal(err)
	}
	copyOfV31Child := required.SDTargetingPredicateNestedV31.Value[0]
	copyOfV31Child.SetValue("14")
	assertExpressionJSON(t, marshalExpression(t, copyOfV31Child), `{"type":"lookback","value":"14","extension":true}`)
	assertExpressionJSON(t, marshalExpression(t, required), raw)
	required.SDTargetingPredicateNestedV31.Value = nil
	assertExpressionJSON(t, marshalExpression(t, required), `{"type":"views","value":null,"outerExtension":true}`)
}

func TestTargetingExpressionNullAndMalformedJSON(t *testing.T) {
	for _, factory := range expressionFactories {
		t.Run(factory.name, func(t *testing.T) {
			expr := factory.new()
			assertExpressionJSON(t, marshalExpression(t, expr), `null`)
			if err := sonic.Unmarshal([]byte(`null`), expr); err != nil {
				t.Fatal(err)
			}
			assertExpressionJSON(t, marshalExpression(t, expr), `null`)
			if err := sonic.Unmarshal([]byte(`{}`), expr); err != nil {
				t.Fatal(err)
			}
			assertExpressionJSON(t, marshalExpression(t, expr), `{}`)
			expr.GetActualInstance().(interface{ SetType(string) }).SetType("asinSameAs")
			expr.GetActualInstance().(interface{ SetValue(string) }).SetValue("new")
			assertExpressionJSON(t, marshalExpression(t, expr), `{"type":"asinSameAs","value":"new"}`)
			for _, raw := range []string{`[]`, `true`, `"scalar"`, `{"type":12}`, `{"type":"asinSameAs","value":{}}`} {
				if err := sonic.Unmarshal([]byte(raw), expr); err == nil {
					t.Fatalf("expected error for %s", raw)
				}
			}
		})
	}
	duplicate := sd_v1.SDTargetExpressionV32{
		SDTargetingPredicateV31:        &sd_v1.SDTargetingPredicateV31{Type: "asinSameAs"},
		SDContentTargetingPredicateV31: &sd_v1.SDContentTargetingPredicateV31{Type: "contentCategorySameAs", Value: "category"},
	}
	if _, err := sonic.Marshal(duplicate); err == nil {
		t.Fatal("multiple populated variants must fail")
	}
}

type targetingRoundTripper func(*http.Request) (*http.Response, error)

func (f targetingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func targetingTestClient(t *testing.T, expectedPath string, check func([]byte)) *sd_v1.APIClient {
	t.Helper()
	return sd_v1.NewAPIClient(&pkg.Configuration{
		Servers: pkg.ServerConfigurations{{URL: "https://example.invalid"}},
		HTTPClient: &http.Client{Transport: targetingRoundTripper(func(req *http.Request) (*http.Response, error) {
			if req.Method != http.MethodPost || req.URL.Path != expectedPath {
				t.Fatalf("unexpected request: %s %s", req.Method, req.URL)
			}
			body, err := io.ReadAll(req.Body)
			if err != nil {
				t.Fatal(err)
			}
			check(body)
			return &http.Response{StatusCode: http.StatusOK, Status: "200 OK", Header: http.Header{"Content-Type": {"application/json"}}, Body: io.NopCloser(strings.NewReader(`{}`)), Request: req}, nil
		})},
	})
}

func TestTargetingExpressionEditsReachPublicRequests(t *testing.T) {
	t.Run("bid recommendations", func(t *testing.T) {
		var request sd_v1.SDTargetingBidRecommendationsRequestV34
		const raw = `{"bidOptimization":"reach","costType":"vcpm","targetingClauses":[{"targetingClause":{"expressionType":"manual","expression":[{"type":"contentCategorySameAs","value":"amzn1.iab-content.325","extension":true}]}}]}`
		if err := sonic.Unmarshal([]byte(raw), &request); err != nil {
			t.Fatal(err)
		}
		request.TargetingClauses[0].TargetingClause.Expression[0].SDContentTargetingPredicateV31.SetValue("amzn1.iab-content.641")
		called := false
		client := targetingTestClient(t, "/sd/targets/bid/recommendations", func(body []byte) {
			called = true
			assertExpressionJSON(t, body, strings.Replace(raw, ".325", ".641", 1))
		})
		if _, _, err := client.BidRecommendationsAPI.GetTargetBidRecommendations(context.Background()).SDTargetingBidRecommendationsRequestV34(request).Execute(); err != nil {
			t.Fatal(err)
		}
		if !called {
			t.Fatal("request was not sent")
		}
	})
	t.Run("forecast", func(t *testing.T) {
		var request sd_v1.SDForecastRequest
		const raw = `{"campaign":{"startDate":"20261001","costType":"vcpm"},"adGroup":{"bidOptimization":"reach","defaultBid":2},"productAds":[{"sku":"test-sku"}],"targetingClauses":[{"expressionType":"manual","expression":[{"type":"contentCategorySameAs","value":"amzn1.iab-content.325","extension":true}]}]}`
		if err := sonic.Unmarshal([]byte(raw), &request); err != nil {
			t.Fatal(err)
		}
		request.TargetingClauses[0].Expression[0].ContentTargetingPredicate.SetValue("amzn1.iab-content.641")
		called := false
		client := targetingTestClient(t, "/sd/forecasts", func(body []byte) {
			called = true
			var fields map[string]json.RawMessage
			if err := json.Unmarshal(body, &fields); err != nil {
				t.Fatal(err)
			}
			assertExpressionJSON(t, fields["targetingClauses"], `[{"bid":null,"expressionType":"manual","expression":[{"type":"contentCategorySameAs","value":"amzn1.iab-content.641","extension":true}]}]`)
		})
		if _, _, err := client.ForecastsAPI.CreateSDForecast(context.Background()).SDForecastRequest(request).Execute(); err != nil {
			t.Fatal(err)
		}
		if !called {
			t.Fatal("request was not sent")
		}
	})
}
