package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetRecommendationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetRecommendationRequest{}

// BudgetRecommendationRequest struct for BudgetRecommendationRequest
type BudgetRecommendationRequest struct {
	// List of campaigns.
	CampaignIds []string `json:"campaignIds"`
}

type _BudgetRecommendationRequest BudgetRecommendationRequest

// NewBudgetRecommendationRequest instantiates a new BudgetRecommendationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetRecommendationRequest(campaignIds []string) *BudgetRecommendationRequest {
	this := BudgetRecommendationRequest{}
	this.CampaignIds = campaignIds
	return &this
}

// NewBudgetRecommendationRequestWithDefaults instantiates a new BudgetRecommendationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetRecommendationRequestWithDefaults() *BudgetRecommendationRequest {
	this := BudgetRecommendationRequest{}
	return &this
}

// GetCampaignIds returns the CampaignIds field value
func (o *BudgetRecommendationRequest) GetCampaignIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CampaignIds
}

// GetCampaignIdsOk returns a tuple with the CampaignIds field value
// and a boolean to check if the value has been set.
func (o *BudgetRecommendationRequest) GetCampaignIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CampaignIds, true
}

// SetCampaignIds sets field value
func (o *BudgetRecommendationRequest) SetCampaignIds(v []string) {
	o.CampaignIds = v
}

func (o BudgetRecommendationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignIds"] = o.CampaignIds
	return toSerialize, nil
}

type NullableBudgetRecommendationRequest struct {
	value *BudgetRecommendationRequest
	isSet bool
}

func (v NullableBudgetRecommendationRequest) Get() *BudgetRecommendationRequest {
	return v.value
}

func (v *NullableBudgetRecommendationRequest) Set(val *BudgetRecommendationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetRecommendationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetRecommendationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetRecommendationRequest(val *BudgetRecommendationRequest) *NullableBudgetRecommendationRequest {
	return &NullableBudgetRecommendationRequest{value: val, isSet: true}
}

func (v NullableBudgetRecommendationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetRecommendationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
