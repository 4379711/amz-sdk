package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SPCampaignOptimizationRecommendationsAPIRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SPCampaignOptimizationRecommendationsAPIRequest{}

// SPCampaignOptimizationRecommendationsAPIRequest struct for SPCampaignOptimizationRecommendationsAPIRequest
type SPCampaignOptimizationRecommendationsAPIRequest struct {
	// A list of campaign ids
	CampaignIds []string `json:"campaignIds"`
}

type _SPCampaignOptimizationRecommendationsAPIRequest SPCampaignOptimizationRecommendationsAPIRequest

// NewSPCampaignOptimizationRecommendationsAPIRequest instantiates a new SPCampaignOptimizationRecommendationsAPIRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSPCampaignOptimizationRecommendationsAPIRequest(campaignIds []string) *SPCampaignOptimizationRecommendationsAPIRequest {
	this := SPCampaignOptimizationRecommendationsAPIRequest{}
	this.CampaignIds = campaignIds
	return &this
}

// NewSPCampaignOptimizationRecommendationsAPIRequestWithDefaults instantiates a new SPCampaignOptimizationRecommendationsAPIRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSPCampaignOptimizationRecommendationsAPIRequestWithDefaults() *SPCampaignOptimizationRecommendationsAPIRequest {
	this := SPCampaignOptimizationRecommendationsAPIRequest{}
	return &this
}

// GetCampaignIds returns the CampaignIds field value
func (o *SPCampaignOptimizationRecommendationsAPIRequest) GetCampaignIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CampaignIds
}

// GetCampaignIdsOk returns a tuple with the CampaignIds field value
// and a boolean to check if the value has been set.
func (o *SPCampaignOptimizationRecommendationsAPIRequest) GetCampaignIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CampaignIds, true
}

// SetCampaignIds sets field value
func (o *SPCampaignOptimizationRecommendationsAPIRequest) SetCampaignIds(v []string) {
	o.CampaignIds = v
}

func (o SPCampaignOptimizationRecommendationsAPIRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignIds"] = o.CampaignIds
	return toSerialize, nil
}

type NullableSPCampaignOptimizationRecommendationsAPIRequest struct {
	value *SPCampaignOptimizationRecommendationsAPIRequest
	isSet bool
}

func (v NullableSPCampaignOptimizationRecommendationsAPIRequest) Get() *SPCampaignOptimizationRecommendationsAPIRequest {
	return v.value
}

func (v *NullableSPCampaignOptimizationRecommendationsAPIRequest) Set(val *SPCampaignOptimizationRecommendationsAPIRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSPCampaignOptimizationRecommendationsAPIRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSPCampaignOptimizationRecommendationsAPIRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPCampaignOptimizationRecommendationsAPIRequest(val *SPCampaignOptimizationRecommendationsAPIRequest) *NullableSPCampaignOptimizationRecommendationsAPIRequest {
	return &NullableSPCampaignOptimizationRecommendationsAPIRequest{value: val, isSet: true}
}

func (v NullableSPCampaignOptimizationRecommendationsAPIRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPCampaignOptimizationRecommendationsAPIRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
