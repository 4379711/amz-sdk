package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetRecommendationForExistingCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetRecommendationForExistingCampaign{}

// BudgetRecommendationForExistingCampaign struct for BudgetRecommendationForExistingCampaign
type BudgetRecommendationForExistingCampaign struct {
	// encrypted campaignId
	CampaignId string `json:"campaignId"`
	// recommended budget for the campaign.
	SuggestedBudget float32 `json:"suggestedBudget"`
	// Correlate the recommendation to the campaign index in the request. Zero-based
	Index                        int32                        `json:"index"`
	SevenDaysMissedOpportunities SevenDaysMissedOpportunities `json:"sevenDaysMissedOpportunities"`
	BudgetRuleRecommendation     BudgetRuleRecommendation     `json:"budgetRuleRecommendation"`
}

type _BudgetRecommendationForExistingCampaign BudgetRecommendationForExistingCampaign

// NewBudgetRecommendationForExistingCampaign instantiates a new BudgetRecommendationForExistingCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetRecommendationForExistingCampaign(campaignId string, suggestedBudget float32, index int32, sevenDaysMissedOpportunities SevenDaysMissedOpportunities, budgetRuleRecommendation BudgetRuleRecommendation) *BudgetRecommendationForExistingCampaign {
	this := BudgetRecommendationForExistingCampaign{}
	this.CampaignId = campaignId
	this.SuggestedBudget = suggestedBudget
	this.Index = index
	this.SevenDaysMissedOpportunities = sevenDaysMissedOpportunities
	this.BudgetRuleRecommendation = budgetRuleRecommendation
	return &this
}

// NewBudgetRecommendationForExistingCampaignWithDefaults instantiates a new BudgetRecommendationForExistingCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetRecommendationForExistingCampaignWithDefaults() *BudgetRecommendationForExistingCampaign {
	this := BudgetRecommendationForExistingCampaign{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *BudgetRecommendationForExistingCampaign) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *BudgetRecommendationForExistingCampaign) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *BudgetRecommendationForExistingCampaign) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetSuggestedBudget returns the SuggestedBudget field value
func (o *BudgetRecommendationForExistingCampaign) GetSuggestedBudget() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SuggestedBudget
}

// GetSuggestedBudgetOk returns a tuple with the SuggestedBudget field value
// and a boolean to check if the value has been set.
func (o *BudgetRecommendationForExistingCampaign) GetSuggestedBudgetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuggestedBudget, true
}

// SetSuggestedBudget sets field value
func (o *BudgetRecommendationForExistingCampaign) SetSuggestedBudget(v float32) {
	o.SuggestedBudget = v
}

// GetIndex returns the Index field value
func (o *BudgetRecommendationForExistingCampaign) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *BudgetRecommendationForExistingCampaign) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *BudgetRecommendationForExistingCampaign) SetIndex(v int32) {
	o.Index = v
}

// GetSevenDaysMissedOpportunities returns the SevenDaysMissedOpportunities field value
func (o *BudgetRecommendationForExistingCampaign) GetSevenDaysMissedOpportunities() SevenDaysMissedOpportunities {
	if o == nil {
		var ret SevenDaysMissedOpportunities
		return ret
	}

	return o.SevenDaysMissedOpportunities
}

// GetSevenDaysMissedOpportunitiesOk returns a tuple with the SevenDaysMissedOpportunities field value
// and a boolean to check if the value has been set.
func (o *BudgetRecommendationForExistingCampaign) GetSevenDaysMissedOpportunitiesOk() (*SevenDaysMissedOpportunities, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SevenDaysMissedOpportunities, true
}

// SetSevenDaysMissedOpportunities sets field value
func (o *BudgetRecommendationForExistingCampaign) SetSevenDaysMissedOpportunities(v SevenDaysMissedOpportunities) {
	o.SevenDaysMissedOpportunities = v
}

// GetBudgetRuleRecommendation returns the BudgetRuleRecommendation field value
func (o *BudgetRecommendationForExistingCampaign) GetBudgetRuleRecommendation() BudgetRuleRecommendation {
	if o == nil {
		var ret BudgetRuleRecommendation
		return ret
	}

	return o.BudgetRuleRecommendation
}

// GetBudgetRuleRecommendationOk returns a tuple with the BudgetRuleRecommendation field value
// and a boolean to check if the value has been set.
func (o *BudgetRecommendationForExistingCampaign) GetBudgetRuleRecommendationOk() (*BudgetRuleRecommendation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetRuleRecommendation, true
}

// SetBudgetRuleRecommendation sets field value
func (o *BudgetRecommendationForExistingCampaign) SetBudgetRuleRecommendation(v BudgetRuleRecommendation) {
	o.BudgetRuleRecommendation = v
}

func (o BudgetRecommendationForExistingCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["suggestedBudget"] = o.SuggestedBudget
	toSerialize["index"] = o.Index
	toSerialize["sevenDaysMissedOpportunities"] = o.SevenDaysMissedOpportunities
	toSerialize["budgetRuleRecommendation"] = o.BudgetRuleRecommendation
	return toSerialize, nil
}

type NullableBudgetRecommendationForExistingCampaign struct {
	value *BudgetRecommendationForExistingCampaign
	isSet bool
}

func (v NullableBudgetRecommendationForExistingCampaign) Get() *BudgetRecommendationForExistingCampaign {
	return v.value
}

func (v *NullableBudgetRecommendationForExistingCampaign) Set(val *BudgetRecommendationForExistingCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetRecommendationForExistingCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetRecommendationForExistingCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetRecommendationForExistingCampaign(val *BudgetRecommendationForExistingCampaign) *NullableBudgetRecommendationForExistingCampaign {
	return &NullableBudgetRecommendationForExistingCampaign{value: val, isSet: true}
}

func (v NullableBudgetRecommendationForExistingCampaign) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetRecommendationForExistingCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
