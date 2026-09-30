package sp_v3

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUpdateCampaignsNullableBody(t *testing.T) {
	unset := NewSponsoredProductsUpdateCampaign("unset")
	null := NewSponsoredProductsUpdateCampaign("null")
	null.SetPortfolioIdNil()
	null.SetEndDateNil()
	value := NewSponsoredProductsUpdateCampaign("value")
	value.SetPortfolioId("123456")
	value.SetEndDate("2026-12-31")
	valueUnset := NewSponsoredProductsUpdateCampaign("value-unset")
	valueUnset.SetPortfolioId("123456")
	valueUnset.SetEndDate("2026-12-31")
	valueUnset.UnsetPortfolioId()
	valueUnset.UnsetEndDate()
	nullUnset := NewSponsoredProductsUpdateCampaign("null-unset")
	nullUnset.SetPortfolioIdNil()
	nullUnset.SetEndDateNil()
	nullUnset.UnsetPortfolioId()
	nullUnset.UnsetEndDate()
	portfolioOnly := NewSponsoredProductsUpdateCampaign("portfolio-only")
	portfolioOnly.SetPortfolioId("123456")
	endDateOnly := NewSponsoredProductsUpdateCampaign("end-date-only")
	endDateOnly.SetEndDate("2026-12-31")
	stateOnly := NewSponsoredProductsUpdateCampaign("state-only")
	stateOnly.SetState(SPONSOREDPRODUCTSCREATEORUPDATEENTITYSTATE_PAUSED)
	nameOnly := NewSponsoredProductsUpdateCampaign("name-only")
	nameOnly.SetName("Renamed campaign")
	budgetOnly := NewSponsoredProductsUpdateCampaign("budget-only")
	budgetOnly.SetBudget(*NewSponsoredProductsCreateOrUpdateBudget(SPONSOREDPRODUCTSCREATEORUPDATEBUDGETTYPE_DAILY, 25))

	body := NewSponsoredProductsUpdateSponsoredProductsCampaignsRequestContent([]SponsoredProductsUpdateCampaign{
		*unset, *null, *value, *valueUnset, *nullUnset, *portfolioOnly, *endDateOnly, *stateOnly, *nameOnly, *budgetOnly,
	})
	assertUpdateRequestBody(t, body, "application/vnd.spCampaign.v3+json", `{"campaigns":[
		{"campaignId":"unset"},
		{"campaignId":"null","portfolioId":null,"endDate":null},
		{"campaignId":"value","portfolioId":"123456","endDate":"2026-12-31"},
		{"campaignId":"value-unset"},
		{"campaignId":"null-unset"},
		{"campaignId":"portfolio-only","portfolioId":"123456"},
		{"campaignId":"end-date-only","endDate":"2026-12-31"},
		{"campaignId":"state-only","state":"PAUSED"},
		{"campaignId":"name-only","name":"Renamed campaign"},
		{"campaignId":"budget-only","budget":{"budgetType":"DAILY","budget":25}}
	]}`)
}

func TestUpdateKeywordsNullableBody(t *testing.T) {
	unset := NewSponsoredProductsUpdateKeyword("unset")
	null := NewSponsoredProductsUpdateKeyword("null")
	null.SetBidNil()
	value := NewSponsoredProductsUpdateKeyword("value")
	value.SetBid(1.25)
	zero := NewSponsoredProductsUpdateKeyword("zero")
	zero.SetBid(0)
	valueUnset := NewSponsoredProductsUpdateKeyword("value-unset")
	valueUnset.SetBid(1.25)
	valueUnset.UnsetBid()
	nullUnset := NewSponsoredProductsUpdateKeyword("null-unset")
	nullUnset.SetBidNil()
	nullUnset.UnsetBid()
	stateOnly := NewSponsoredProductsUpdateKeyword("state-only")
	stateOnly.SetState(SPONSOREDPRODUCTSCREATEORUPDATEENTITYSTATE_PAUSED)

	body := NewSponsoredProductsUpdateSponsoredProductsKeywordsRequestContent([]SponsoredProductsUpdateKeyword{
		*unset, *null, *value, *zero, *valueUnset, *nullUnset, *stateOnly,
	})
	assertUpdateRequestBody(t, body, "application/vnd.spKeyword.v3+json", `{"keywords":[
		{"keywordId":"unset"},
		{"keywordId":"null","bid":null},
		{"keywordId":"value","bid":1.25},
		{"keywordId":"zero","bid":0},
		{"keywordId":"value-unset"},
		{"keywordId":"null-unset"},
		{"keywordId":"state-only","state":"PAUSED"}
	]}`)
}

func TestUpdateTargetingClausesNullableBody(t *testing.T) {
	unset := NewSponsoredProductsUpdateTargetingClause("unset")
	null := NewSponsoredProductsUpdateTargetingClause("null")
	null.SetBidNil()
	value := NewSponsoredProductsUpdateTargetingClause("value")
	value.SetBid(1.25)
	zero := NewSponsoredProductsUpdateTargetingClause("zero")
	zero.SetBid(0)
	valueUnset := NewSponsoredProductsUpdateTargetingClause("value-unset")
	valueUnset.SetBid(1.25)
	valueUnset.UnsetBid()
	nullUnset := NewSponsoredProductsUpdateTargetingClause("null-unset")
	nullUnset.SetBidNil()
	nullUnset.UnsetBid()
	stateOnly := NewSponsoredProductsUpdateTargetingClause("state-only")
	stateOnly.SetState(SPONSOREDPRODUCTSCREATEORUPDATEENTITYSTATE_PAUSED)

	body := NewSponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent([]SponsoredProductsUpdateTargetingClause{
		*unset, *null, *value, *zero, *valueUnset, *nullUnset, *stateOnly,
	})
	assertUpdateRequestBody(t, body, "application/vnd.spTargetingClause.v3+json", `{"targetingClauses":[
		{"targetId":"unset"},
		{"targetId":"null","bid":null},
		{"targetId":"value","bid":1.25},
		{"targetId":"zero","bid":0},
		{"targetId":"value-unset"},
		{"targetId":"null-unset"},
		{"targetId":"state-only","state":"PAUSED"}
	]}`)
}

func assertUpdateRequestBody(t *testing.T, body interface{}, contentType, wantJSON string) {
	t.Helper()
	// Exercise the same Sonic encoder as prepareRequest, using the real batch models.
	encoded, err := setBody(body, contentType)
	if err != nil {
		t.Fatal(err)
	}
	// Decode into generic JSON values so absent keys, null, and zero remain distinct.
	var got, want interface{}
	if err := json.Unmarshal(encoded.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(wantJSON), &want); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("request body = %s\nwant %s", encoded, wantJSON)
	}
}
