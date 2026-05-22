package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDBudgetRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDBudgetRecommendation{}

// SDBudgetRecommendation struct for SDBudgetRecommendation
type SDBudgetRecommendation struct {
	// Correlate the recommendation to the campaign index in the request. Zero-based.
	Index int32 `json:"index"`
	// Campaign id.
	CampaignId string `json:"campaignId"`
	// Recommended budget for the campaign. This will be in local currency.
	SuggestedBudget              float32                        `json:"suggestedBudget"`
	SevenDaysMissedOpportunities SDSevenDaysMissedOpportunities `json:"sevenDaysMissedOpportunities"`
}

type _SDBudgetRecommendation SDBudgetRecommendation

// NewSDBudgetRecommendation instantiates a new SDBudgetRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDBudgetRecommendation(index int32, campaignId string, suggestedBudget float32, sevenDaysMissedOpportunities SDSevenDaysMissedOpportunities) *SDBudgetRecommendation {
	this := SDBudgetRecommendation{}
	this.Index = index
	this.CampaignId = campaignId
	this.SuggestedBudget = suggestedBudget
	this.SevenDaysMissedOpportunities = sevenDaysMissedOpportunities
	return &this
}

// NewSDBudgetRecommendationWithDefaults instantiates a new SDBudgetRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDBudgetRecommendationWithDefaults() *SDBudgetRecommendation {
	this := SDBudgetRecommendation{}
	return &this
}

// GetIndex returns the Index field value
func (o *SDBudgetRecommendation) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SDBudgetRecommendation) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SDBudgetRecommendation) SetIndex(v int32) {
	o.Index = v
}

// GetCampaignId returns the CampaignId field value
func (o *SDBudgetRecommendation) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SDBudgetRecommendation) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SDBudgetRecommendation) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetSuggestedBudget returns the SuggestedBudget field value
func (o *SDBudgetRecommendation) GetSuggestedBudget() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SuggestedBudget
}

// GetSuggestedBudgetOk returns a tuple with the SuggestedBudget field value
// and a boolean to check if the value has been set.
func (o *SDBudgetRecommendation) GetSuggestedBudgetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuggestedBudget, true
}

// SetSuggestedBudget sets field value
func (o *SDBudgetRecommendation) SetSuggestedBudget(v float32) {
	o.SuggestedBudget = v
}

// GetSevenDaysMissedOpportunities returns the SevenDaysMissedOpportunities field value
func (o *SDBudgetRecommendation) GetSevenDaysMissedOpportunities() SDSevenDaysMissedOpportunities {
	if o == nil {
		var ret SDSevenDaysMissedOpportunities
		return ret
	}

	return o.SevenDaysMissedOpportunities
}

// GetSevenDaysMissedOpportunitiesOk returns a tuple with the SevenDaysMissedOpportunities field value
// and a boolean to check if the value has been set.
func (o *SDBudgetRecommendation) GetSevenDaysMissedOpportunitiesOk() (*SDSevenDaysMissedOpportunities, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SevenDaysMissedOpportunities, true
}

// SetSevenDaysMissedOpportunities sets field value
func (o *SDBudgetRecommendation) SetSevenDaysMissedOpportunities(v SDSevenDaysMissedOpportunities) {
	o.SevenDaysMissedOpportunities = v
}

func (o SDBudgetRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["suggestedBudget"] = o.SuggestedBudget
	toSerialize["sevenDaysMissedOpportunities"] = o.SevenDaysMissedOpportunities
	return toSerialize, nil
}

type NullableSDBudgetRecommendation struct {
	value *SDBudgetRecommendation
	isSet bool
}

func (v NullableSDBudgetRecommendation) Get() *SDBudgetRecommendation {
	return v.value
}

func (v *NullableSDBudgetRecommendation) Set(val *SDBudgetRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableSDBudgetRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableSDBudgetRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDBudgetRecommendation(val *SDBudgetRecommendation) *NullableSDBudgetRecommendation {
	return &NullableSDBudgetRecommendation{value: val, isSet: true}
}

func (v NullableSDBudgetRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDBudgetRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
