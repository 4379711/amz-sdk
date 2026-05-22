package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalBudgetRulesRecommendationEventRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalBudgetRulesRecommendationEventRequest{}

// GlobalBudgetRulesRecommendationEventRequest struct for GlobalBudgetRulesRecommendationEventRequest
type GlobalBudgetRulesRecommendationEventRequest struct {
	// The campaign identifier.
	CampaignId string `json:"campaignId"`
}

type _GlobalBudgetRulesRecommendationEventRequest GlobalBudgetRulesRecommendationEventRequest

// NewGlobalBudgetRulesRecommendationEventRequest instantiates a new GlobalBudgetRulesRecommendationEventRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalBudgetRulesRecommendationEventRequest(campaignId string) *GlobalBudgetRulesRecommendationEventRequest {
	this := GlobalBudgetRulesRecommendationEventRequest{}
	this.CampaignId = campaignId
	return &this
}

// NewGlobalBudgetRulesRecommendationEventRequestWithDefaults instantiates a new GlobalBudgetRulesRecommendationEventRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalBudgetRulesRecommendationEventRequestWithDefaults() *GlobalBudgetRulesRecommendationEventRequest {
	this := GlobalBudgetRulesRecommendationEventRequest{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *GlobalBudgetRulesRecommendationEventRequest) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRulesRecommendationEventRequest) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *GlobalBudgetRulesRecommendationEventRequest) SetCampaignId(v string) {
	o.CampaignId = v
}

func (o GlobalBudgetRulesRecommendationEventRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	return toSerialize, nil
}

type NullableGlobalBudgetRulesRecommendationEventRequest struct {
	value *GlobalBudgetRulesRecommendationEventRequest
	isSet bool
}

func (v NullableGlobalBudgetRulesRecommendationEventRequest) Get() *GlobalBudgetRulesRecommendationEventRequest {
	return v.value
}

func (v *NullableGlobalBudgetRulesRecommendationEventRequest) Set(val *GlobalBudgetRulesRecommendationEventRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalBudgetRulesRecommendationEventRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalBudgetRulesRecommendationEventRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalBudgetRulesRecommendationEventRequest(val *GlobalBudgetRulesRecommendationEventRequest) *NullableGlobalBudgetRulesRecommendationEventRequest {
	return &NullableGlobalBudgetRulesRecommendationEventRequest{value: val, isSet: true}
}

func (v NullableGlobalBudgetRulesRecommendationEventRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalBudgetRulesRecommendationEventRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
