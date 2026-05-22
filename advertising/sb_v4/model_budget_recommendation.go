package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetRecommendation{}

// BudgetRecommendation Budget recomendation for campagins.
type BudgetRecommendation struct {
	// The identifier of a campaign.
	CampaignId string `json:"campaignId"`
	// Recommended budget for the campaign.
	SuggestedBudget float64 `json:"suggestedBudget"`
	// Correlate the recommendation to the campaign index in the request. Zero-based.
	Index                        float32                      `json:"index"`
	SevenDaysMissedOpportunities SevenDaysMissedOpportunities `json:"sevenDaysMissedOpportunities"`
}

type _BudgetRecommendation BudgetRecommendation

// NewBudgetRecommendation instantiates a new BudgetRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetRecommendation(campaignId string, suggestedBudget float64, index float32, sevenDaysMissedOpportunities SevenDaysMissedOpportunities) *BudgetRecommendation {
	this := BudgetRecommendation{}
	this.CampaignId = campaignId
	this.SuggestedBudget = suggestedBudget
	this.Index = index
	this.SevenDaysMissedOpportunities = sevenDaysMissedOpportunities
	return &this
}

// NewBudgetRecommendationWithDefaults instantiates a new BudgetRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetRecommendationWithDefaults() *BudgetRecommendation {
	this := BudgetRecommendation{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *BudgetRecommendation) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *BudgetRecommendation) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *BudgetRecommendation) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetSuggestedBudget returns the SuggestedBudget field value
func (o *BudgetRecommendation) GetSuggestedBudget() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SuggestedBudget
}

// GetSuggestedBudgetOk returns a tuple with the SuggestedBudget field value
// and a boolean to check if the value has been set.
func (o *BudgetRecommendation) GetSuggestedBudgetOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuggestedBudget, true
}

// SetSuggestedBudget sets field value
func (o *BudgetRecommendation) SetSuggestedBudget(v float64) {
	o.SuggestedBudget = v
}

// GetIndex returns the Index field value
func (o *BudgetRecommendation) GetIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *BudgetRecommendation) GetIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *BudgetRecommendation) SetIndex(v float32) {
	o.Index = v
}

// GetSevenDaysMissedOpportunities returns the SevenDaysMissedOpportunities field value
func (o *BudgetRecommendation) GetSevenDaysMissedOpportunities() SevenDaysMissedOpportunities {
	if o == nil {
		var ret SevenDaysMissedOpportunities
		return ret
	}

	return o.SevenDaysMissedOpportunities
}

// GetSevenDaysMissedOpportunitiesOk returns a tuple with the SevenDaysMissedOpportunities field value
// and a boolean to check if the value has been set.
func (o *BudgetRecommendation) GetSevenDaysMissedOpportunitiesOk() (*SevenDaysMissedOpportunities, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SevenDaysMissedOpportunities, true
}

// SetSevenDaysMissedOpportunities sets field value
func (o *BudgetRecommendation) SetSevenDaysMissedOpportunities(v SevenDaysMissedOpportunities) {
	o.SevenDaysMissedOpportunities = v
}

func (o BudgetRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["suggestedBudget"] = o.SuggestedBudget
	toSerialize["index"] = o.Index
	toSerialize["sevenDaysMissedOpportunities"] = o.SevenDaysMissedOpportunities
	return toSerialize, nil
}

type NullableBudgetRecommendation struct {
	value *BudgetRecommendation
	isSet bool
}

func (v NullableBudgetRecommendation) Get() *BudgetRecommendation {
	return v.value
}

func (v *NullableBudgetRecommendation) Set(val *BudgetRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetRecommendation(val *BudgetRecommendation) *NullableBudgetRecommendation {
	return &NullableBudgetRecommendation{value: val, isSet: true}
}

func (v NullableBudgetRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
