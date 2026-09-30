package sp_v3

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/4379711/amz-sdk/pkg"
	"github.com/bytedance/sonic"
)

func TestRuleCriteriaVariants(t *testing.T) {
	for _, tc := range []struct {
		name, input string
		isRange     bool
	}{
		{"range", `{"minValue":0,"maxValue":10}`, true},
		{"zero_range", `{"minValue":0,"maxValue":0}`, true},
		{"value", `{"comparisonOperator":"GREATER_THAN","value":0}`, false},
		{"unknown_operator", `{"comparisonOperator":"FUTURE_OPERATOR","value":2}`, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var criteria OptimizationRulesAPIRuleCriteria
			if err := sonic.Unmarshal([]byte(tc.input), &criteria); err != nil {
				t.Fatal(err)
			}
			if (criteria.OptimizationRulesAPIRangeTypeRuleCriteria != nil) != tc.isRange ||
				(criteria.OptimizationRulesAPIValueTypeRuleCriteria != nil) == tc.isRange {
				t.Fatalf("wrong criteria variant: %+v", criteria)
			}
			encoded, err := sonic.Marshal(criteria)
			if err != nil {
				t.Fatal(err)
			}
			var got, want map[string]json.RawMessage
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal([]byte(tc.input), &want); err != nil {
				t.Fatal(err)
			}
			if len(got) != len(want) {
				t.Fatalf("round trip = %s, want %s", encoded, tc.input)
			}
			for key, value := range want {
				if string(got[key]) != string(value) {
					t.Errorf("round trip %s = %s, want %s", key, got[key], value)
				}
			}
		})
	}
}

func TestRuleCriteriaRejectsInvalidShapes(t *testing.T) {
	for _, input := range []string{
		`{}`, `null`, `[]`, `1`, `true`, `"value"`,
		`{"minValue":0}`, `{"maxValue":0}`, `{"value":0}`, `{"comparisonOperator":"GREATER_THAN"}`,
		`{"minValue":null,"maxValue":1}`, `{"minValue":0,"maxValue":null}`,
		`{"comparisonOperator":null,"value":0}`, `{"comparisonOperator":"GREATER_THAN","value":null}`,
		`{"minValue":0,"maxValue":1,"comparisonOperator":"GREATER_THAN","value":0}`,
		`{"minValue":0,"maxValue":1,"value":null}`,
		`{"minValue":"0","maxValue":1}`, `{"comparisonOperator":123,"value":0}`,
	} {
		t.Run(input, func(t *testing.T) {
			criteria := OptimizationRulesAPIRangeTypeRuleCriteriaAsOptimizationRulesAPIRuleCriteria(
				NewOptimizationRulesAPIRangeTypeRuleCriteria(1, 2))
			if err := sonic.Unmarshal([]byte(input), &criteria); err == nil {
				t.Fatalf("accepted invalid criteria %s", input)
			}
			if criteria.GetActualInstance() != nil {
				t.Fatalf("kept stale criteria after error: %+v", criteria)
			}
		})
	}
}

func TestRuleCriteriaReuseAndOptionalParent(t *testing.T) {
	var criteria OptimizationRulesAPIRuleCriteria
	for _, input := range []string{`{"minValue":0,"maxValue":1}`, `{"comparisonOperator":"FUTURE_OPERATOR","value":0}`} {
		if err := sonic.Unmarshal([]byte(input), &criteria); err != nil {
			t.Fatal(err)
		}
	}
	if criteria.OptimizationRulesAPIRangeTypeRuleCriteria != nil || criteria.OptimizationRulesAPIValueTypeRuleCriteria == nil {
		t.Fatalf("reused destination retained the previous variant: %+v", criteria)
	}
	if got := criteria.OptimizationRulesAPIValueTypeRuleCriteria.ComparisonOperator; got != "FUTURE_OPERATOR" || got.IsValid() {
		t.Fatalf("unknown operator handling changed: %q", got)
	}
	for _, input := range []string{`{}`, `{"criteria":null}`} {
		var condition OptimizationRulesAPIRuleCondition
		if err := sonic.Unmarshal([]byte(input), &condition); err != nil {
			t.Fatalf("optional criteria in %s: %v", input, err)
		}
		if condition.Criteria != nil {
			t.Fatalf("optional criteria in %s is non-nil", input)
		}
	}
}

func TestSearchOptimizationRulesDecodesCriteria(t *testing.T) {
	response := `{"optimizationRules":[{"conditions":[{}, {"criteria":{"minValue":0,"maxValue":10}}, {"criteria":{"comparisonOperator":"FUTURE_OPERATOR","value":0}}]}]}`
	cfg := &pkg.Configuration{
		Servers: pkg.ServerConfigurations{{URL: "https://example.invalid"}},
		HTTPClient: &http.Client{Transport: numericIDTransport(func(r *http.Request) (*http.Response, error) {
			if r.URL.Path != "/sp/rules/optimization/search" {
				t.Errorf("unexpected path %q", r.URL.Path)
			}
			return &http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": {"application/vnd.spoptimizationrules.v1+json"}}, Body: io.NopCloser(strings.NewReader(response)), Request: r}, nil
		})},
	}
	result, _, err := NewAPIClient(cfg).OptimizationRulesAPI.SearchOptimizationRules(context.Background()).
		OptimizationRulesAPISearchOptimizationRulesRequest(OptimizationRulesAPISearchOptimizationRulesRequest{}).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if result == nil || len(result.OptimizationRules) != 1 || len(result.OptimizationRules[0].Conditions) != 3 {
		t.Fatalf("unexpected result: %+v", result)
	}
	conditions := result.OptimizationRules[0].Conditions
	if conditions[0].Criteria != nil || conditions[1].Criteria.OptimizationRulesAPIRangeTypeRuleCriteria == nil || conditions[2].Criteria.OptimizationRulesAPIValueTypeRuleCriteria == nil {
		t.Fatalf("incorrect criteria variants in response: %+v", conditions)
	}
}
