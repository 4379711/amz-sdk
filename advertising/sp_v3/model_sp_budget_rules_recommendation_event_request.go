package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SPBudgetRulesRecommendationEventRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SPBudgetRulesRecommendationEventRequest{}

// SPBudgetRulesRecommendationEventRequest struct for SPBudgetRulesRecommendationEventRequest
type SPBudgetRulesRecommendationEventRequest struct {
	// The campaign identifier.
	CampaignId string `json:"campaignId"`
}

type _SPBudgetRulesRecommendationEventRequest SPBudgetRulesRecommendationEventRequest

// NewSPBudgetRulesRecommendationEventRequest instantiates a new SPBudgetRulesRecommendationEventRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSPBudgetRulesRecommendationEventRequest(campaignId string) *SPBudgetRulesRecommendationEventRequest {
	this := SPBudgetRulesRecommendationEventRequest{}
	this.CampaignId = campaignId
	return &this
}

// NewSPBudgetRulesRecommendationEventRequestWithDefaults instantiates a new SPBudgetRulesRecommendationEventRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSPBudgetRulesRecommendationEventRequestWithDefaults() *SPBudgetRulesRecommendationEventRequest {
	this := SPBudgetRulesRecommendationEventRequest{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *SPBudgetRulesRecommendationEventRequest) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SPBudgetRulesRecommendationEventRequest) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SPBudgetRulesRecommendationEventRequest) SetCampaignId(v string) {
	o.CampaignId = v
}

func (o SPBudgetRulesRecommendationEventRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	return toSerialize, nil
}

type NullableSPBudgetRulesRecommendationEventRequest struct {
	value *SPBudgetRulesRecommendationEventRequest
	isSet bool
}

func (v NullableSPBudgetRulesRecommendationEventRequest) Get() *SPBudgetRulesRecommendationEventRequest {
	return v.value
}

func (v *NullableSPBudgetRulesRecommendationEventRequest) Set(val *SPBudgetRulesRecommendationEventRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSPBudgetRulesRecommendationEventRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSPBudgetRulesRecommendationEventRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPBudgetRulesRecommendationEventRequest(val *SPBudgetRulesRecommendationEventRequest) *NullableSPBudgetRulesRecommendationEventRequest {
	return &NullableSPBudgetRulesRecommendationEventRequest{value: val, isSet: true}
}

func (v NullableSPBudgetRulesRecommendationEventRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPBudgetRulesRecommendationEventRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
