package sp_v3

import "github.com/bytedance/sonic"

// checks if the GlobalBudgetRecommendationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalBudgetRecommendationRequest{}

// GlobalBudgetRecommendationRequest struct for GlobalBudgetRecommendationRequest
type GlobalBudgetRecommendationRequest struct {
	// List of campaign Ids
	CampaignIds []string `json:"campaignIds"`
}

type _GlobalBudgetRecommendationRequest GlobalBudgetRecommendationRequest

// NewGlobalBudgetRecommendationRequest instantiates a new GlobalBudgetRecommendationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalBudgetRecommendationRequest(campaignIds []string) *GlobalBudgetRecommendationRequest {
	this := GlobalBudgetRecommendationRequest{}
	this.CampaignIds = campaignIds
	return &this
}

// NewGlobalBudgetRecommendationRequestWithDefaults instantiates a new GlobalBudgetRecommendationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalBudgetRecommendationRequestWithDefaults() *GlobalBudgetRecommendationRequest {
	this := GlobalBudgetRecommendationRequest{}
	return &this
}

// GetCampaignIds returns the CampaignIds field value
func (o *GlobalBudgetRecommendationRequest) GetCampaignIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CampaignIds
}

// GetCampaignIdsOk returns a tuple with the CampaignIds field value
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRecommendationRequest) GetCampaignIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CampaignIds, true
}

// SetCampaignIds sets field value
func (o *GlobalBudgetRecommendationRequest) SetCampaignIds(v []string) {
	o.CampaignIds = v
}

func (o GlobalBudgetRecommendationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignIds"] = o.CampaignIds
	return toSerialize, nil
}

type NullableGlobalBudgetRecommendationRequest struct {
	value *GlobalBudgetRecommendationRequest
	isSet bool
}

func (v NullableGlobalBudgetRecommendationRequest) Get() *GlobalBudgetRecommendationRequest {
	return v.value
}

func (v *NullableGlobalBudgetRecommendationRequest) Set(val *GlobalBudgetRecommendationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalBudgetRecommendationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalBudgetRecommendationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalBudgetRecommendationRequest(val *GlobalBudgetRecommendationRequest) *NullableGlobalBudgetRecommendationRequest {
	return &NullableGlobalBudgetRecommendationRequest{value: val, isSet: true}
}

func (v NullableGlobalBudgetRecommendationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalBudgetRecommendationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
