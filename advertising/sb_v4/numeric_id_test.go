package sb_v4

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/4379711/amz-sdk/pkg"
)

type numericIDTransport func(*http.Request) (*http.Response, error)

func (f numericIDTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestBudgetRuleCampaignIDPaths(t *testing.T) {
	for _, id := range []int64{16777217, 36027908138365, 9223372036854775807} {
		t.Run(strconv.FormatInt(id, 10), func(t *testing.T) {
			calls := 0
			cfg := &pkg.Configuration{
				Servers: pkg.ServerConfigurations{{URL: "https://example.invalid"}},
				HTTPClient: &http.Client{Transport: numericIDTransport(func(r *http.Request) (*http.Response, error) {
					calls++
					want := "/sb/campaigns/" + strconv.FormatInt(id, 10) + "/budgetRules"
					if r.Method == http.MethodDelete {
						want += "/rule-id"
					}
					if r.URL.EscapedPath() != want {
						t.Errorf("request path = %q, want %q", r.URL.EscapedPath(), want)
					}
					return &http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": {"application/json"}}, Body: io.NopCloser(strings.NewReader("{}")), Request: r}, nil
				})},
			}
			api := NewAPIClient(cfg).BudgetRulesAPI
			_, _, err := api.CreateAssociatedBudgetRulesForSBCampaigns(context.Background(), id).
				CreateAssociatedBudgetRulesRequest(CreateAssociatedBudgetRulesRequest{BudgetRuleIds: []string{"rule-id"}}).Execute()
			if err != nil {
				t.Fatal(err)
			}
			if _, _, err := api.ListAssociatedBudgetRulesForSBCampaigns(context.Background(), id).Execute(); err != nil {
				t.Fatal(err)
			}
			if _, _, err := api.DisassociateAssociatedBudgetRuleForSBCampaigns(context.Background(), id, "rule-id").Execute(); err != nil {
				t.Fatal(err)
			}
			if calls != 3 {
				t.Fatalf("requests = %d, want 3", calls)
			}
		})
	}
}
