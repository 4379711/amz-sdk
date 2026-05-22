package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CampaignRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignRecommendation{}

// CampaignRecommendation This object contains a set of recommendations for a campaign across bid, budget, targeting.
type CampaignRecommendation struct {
	SevenDaysEstimatedOpportunities *SevenDaysEstimatedOpportunities `json:"sevenDaysEstimatedOpportunities,omitempty"`
	BiddingStrategyRecommendation   *BiddingStrategyRecommendation   `json:"biddingStrategyRecommendation,omitempty"`
	// The identifier of the campaign.
	CampaignId                       *string                           `json:"campaignId,omitempty"`
	KeywordTargetingRecommendations  []KeywordTargetingRecommendation  `json:"keywordTargetingRecommendations,omitempty"`
	BudgetRecommendation             *BudgetRecommendation             `json:"budgetRecommendation,omitempty"`
	TargetingGroupBidRecommendations []TargetingGroupBidRecommendation `json:"targetingGroupBidRecommendations,omitempty"`
}

// NewCampaignRecommendation instantiates a new CampaignRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignRecommendation() *CampaignRecommendation {
	this := CampaignRecommendation{}
	return &this
}

// NewCampaignRecommendationWithDefaults instantiates a new CampaignRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignRecommendationWithDefaults() *CampaignRecommendation {
	this := CampaignRecommendation{}
	return &this
}

// GetSevenDaysEstimatedOpportunities returns the SevenDaysEstimatedOpportunities field value if set, zero value otherwise.
func (o *CampaignRecommendation) GetSevenDaysEstimatedOpportunities() SevenDaysEstimatedOpportunities {
	if o == nil || IsNil(o.SevenDaysEstimatedOpportunities) {
		var ret SevenDaysEstimatedOpportunities
		return ret
	}
	return *o.SevenDaysEstimatedOpportunities
}

// GetSevenDaysEstimatedOpportunitiesOk returns a tuple with the SevenDaysEstimatedOpportunities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRecommendation) GetSevenDaysEstimatedOpportunitiesOk() (*SevenDaysEstimatedOpportunities, bool) {
	if o == nil || IsNil(o.SevenDaysEstimatedOpportunities) {
		return nil, false
	}
	return o.SevenDaysEstimatedOpportunities, true
}

// HasSevenDaysEstimatedOpportunities returns a boolean if a field has been set.
func (o *CampaignRecommendation) HasSevenDaysEstimatedOpportunities() bool {
	if o != nil && !IsNil(o.SevenDaysEstimatedOpportunities) {
		return true
	}

	return false
}

// SetSevenDaysEstimatedOpportunities gets a reference to the given SevenDaysEstimatedOpportunities and assigns it to the SevenDaysEstimatedOpportunities field.
func (o *CampaignRecommendation) SetSevenDaysEstimatedOpportunities(v SevenDaysEstimatedOpportunities) {
	o.SevenDaysEstimatedOpportunities = &v
}

// GetBiddingStrategyRecommendation returns the BiddingStrategyRecommendation field value if set, zero value otherwise.
func (o *CampaignRecommendation) GetBiddingStrategyRecommendation() BiddingStrategyRecommendation {
	if o == nil || IsNil(o.BiddingStrategyRecommendation) {
		var ret BiddingStrategyRecommendation
		return ret
	}
	return *o.BiddingStrategyRecommendation
}

// GetBiddingStrategyRecommendationOk returns a tuple with the BiddingStrategyRecommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRecommendation) GetBiddingStrategyRecommendationOk() (*BiddingStrategyRecommendation, bool) {
	if o == nil || IsNil(o.BiddingStrategyRecommendation) {
		return nil, false
	}
	return o.BiddingStrategyRecommendation, true
}

// HasBiddingStrategyRecommendation returns a boolean if a field has been set.
func (o *CampaignRecommendation) HasBiddingStrategyRecommendation() bool {
	if o != nil && !IsNil(o.BiddingStrategyRecommendation) {
		return true
	}

	return false
}

// SetBiddingStrategyRecommendation gets a reference to the given BiddingStrategyRecommendation and assigns it to the BiddingStrategyRecommendation field.
func (o *CampaignRecommendation) SetBiddingStrategyRecommendation(v BiddingStrategyRecommendation) {
	o.BiddingStrategyRecommendation = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *CampaignRecommendation) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRecommendation) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *CampaignRecommendation) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *CampaignRecommendation) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetKeywordTargetingRecommendations returns the KeywordTargetingRecommendations field value if set, zero value otherwise.
func (o *CampaignRecommendation) GetKeywordTargetingRecommendations() []KeywordTargetingRecommendation {
	if o == nil || IsNil(o.KeywordTargetingRecommendations) {
		var ret []KeywordTargetingRecommendation
		return ret
	}
	return o.KeywordTargetingRecommendations
}

// GetKeywordTargetingRecommendationsOk returns a tuple with the KeywordTargetingRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRecommendation) GetKeywordTargetingRecommendationsOk() ([]KeywordTargetingRecommendation, bool) {
	if o == nil || IsNil(o.KeywordTargetingRecommendations) {
		return nil, false
	}
	return o.KeywordTargetingRecommendations, true
}

// HasKeywordTargetingRecommendations returns a boolean if a field has been set.
func (o *CampaignRecommendation) HasKeywordTargetingRecommendations() bool {
	if o != nil && !IsNil(o.KeywordTargetingRecommendations) {
		return true
	}

	return false
}

// SetKeywordTargetingRecommendations gets a reference to the given []KeywordTargetingRecommendation and assigns it to the KeywordTargetingRecommendations field.
func (o *CampaignRecommendation) SetKeywordTargetingRecommendations(v []KeywordTargetingRecommendation) {
	o.KeywordTargetingRecommendations = v
}

// GetBudgetRecommendation returns the BudgetRecommendation field value if set, zero value otherwise.
func (o *CampaignRecommendation) GetBudgetRecommendation() BudgetRecommendation {
	if o == nil || IsNil(o.BudgetRecommendation) {
		var ret BudgetRecommendation
		return ret
	}
	return *o.BudgetRecommendation
}

// GetBudgetRecommendationOk returns a tuple with the BudgetRecommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRecommendation) GetBudgetRecommendationOk() (*BudgetRecommendation, bool) {
	if o == nil || IsNil(o.BudgetRecommendation) {
		return nil, false
	}
	return o.BudgetRecommendation, true
}

// HasBudgetRecommendation returns a boolean if a field has been set.
func (o *CampaignRecommendation) HasBudgetRecommendation() bool {
	if o != nil && !IsNil(o.BudgetRecommendation) {
		return true
	}

	return false
}

// SetBudgetRecommendation gets a reference to the given BudgetRecommendation and assigns it to the BudgetRecommendation field.
func (o *CampaignRecommendation) SetBudgetRecommendation(v BudgetRecommendation) {
	o.BudgetRecommendation = &v
}

// GetTargetingGroupBidRecommendations returns the TargetingGroupBidRecommendations field value if set, zero value otherwise.
func (o *CampaignRecommendation) GetTargetingGroupBidRecommendations() []TargetingGroupBidRecommendation {
	if o == nil || IsNil(o.TargetingGroupBidRecommendations) {
		var ret []TargetingGroupBidRecommendation
		return ret
	}
	return o.TargetingGroupBidRecommendations
}

// GetTargetingGroupBidRecommendationsOk returns a tuple with the TargetingGroupBidRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRecommendation) GetTargetingGroupBidRecommendationsOk() ([]TargetingGroupBidRecommendation, bool) {
	if o == nil || IsNil(o.TargetingGroupBidRecommendations) {
		return nil, false
	}
	return o.TargetingGroupBidRecommendations, true
}

// HasTargetingGroupBidRecommendations returns a boolean if a field has been set.
func (o *CampaignRecommendation) HasTargetingGroupBidRecommendations() bool {
	if o != nil && !IsNil(o.TargetingGroupBidRecommendations) {
		return true
	}

	return false
}

// SetTargetingGroupBidRecommendations gets a reference to the given []TargetingGroupBidRecommendation and assigns it to the TargetingGroupBidRecommendations field.
func (o *CampaignRecommendation) SetTargetingGroupBidRecommendations(v []TargetingGroupBidRecommendation) {
	o.TargetingGroupBidRecommendations = v
}

func (o CampaignRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SevenDaysEstimatedOpportunities) {
		toSerialize["sevenDaysEstimatedOpportunities"] = o.SevenDaysEstimatedOpportunities
	}
	if !IsNil(o.BiddingStrategyRecommendation) {
		toSerialize["biddingStrategyRecommendation"] = o.BiddingStrategyRecommendation
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	if !IsNil(o.KeywordTargetingRecommendations) {
		toSerialize["keywordTargetingRecommendations"] = o.KeywordTargetingRecommendations
	}
	if !IsNil(o.BudgetRecommendation) {
		toSerialize["budgetRecommendation"] = o.BudgetRecommendation
	}
	if !IsNil(o.TargetingGroupBidRecommendations) {
		toSerialize["targetingGroupBidRecommendations"] = o.TargetingGroupBidRecommendations
	}
	return toSerialize, nil
}

type NullableCampaignRecommendation struct {
	value *CampaignRecommendation
	isSet bool
}

func (v NullableCampaignRecommendation) Get() *CampaignRecommendation {
	return v.value
}

func (v *NullableCampaignRecommendation) Set(val *CampaignRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignRecommendation(val *CampaignRecommendation) *NullableCampaignRecommendation {
	return &NullableCampaignRecommendation{value: val, isSet: true}
}

func (v NullableCampaignRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCampaignRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
