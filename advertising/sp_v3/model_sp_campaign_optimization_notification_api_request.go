package sp_v3

import "github.com/bytedance/sonic"

// checks if the SPCampaignOptimizationNotificationAPIRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SPCampaignOptimizationNotificationAPIRequest{}

// SPCampaignOptimizationNotificationAPIRequest struct for SPCampaignOptimizationNotificationAPIRequest
type SPCampaignOptimizationNotificationAPIRequest struct {
	// A list of campaign ids
	CampaignIds []string `json:"campaignIds"`
}

type _SPCampaignOptimizationNotificationAPIRequest SPCampaignOptimizationNotificationAPIRequest

// NewSPCampaignOptimizationNotificationAPIRequest instantiates a new SPCampaignOptimizationNotificationAPIRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSPCampaignOptimizationNotificationAPIRequest(campaignIds []string) *SPCampaignOptimizationNotificationAPIRequest {
	this := SPCampaignOptimizationNotificationAPIRequest{}
	this.CampaignIds = campaignIds
	return &this
}

// NewSPCampaignOptimizationNotificationAPIRequestWithDefaults instantiates a new SPCampaignOptimizationNotificationAPIRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSPCampaignOptimizationNotificationAPIRequestWithDefaults() *SPCampaignOptimizationNotificationAPIRequest {
	this := SPCampaignOptimizationNotificationAPIRequest{}
	return &this
}

// GetCampaignIds returns the CampaignIds field value
func (o *SPCampaignOptimizationNotificationAPIRequest) GetCampaignIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CampaignIds
}

// GetCampaignIdsOk returns a tuple with the CampaignIds field value
// and a boolean to check if the value has been set.
func (o *SPCampaignOptimizationNotificationAPIRequest) GetCampaignIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CampaignIds, true
}

// SetCampaignIds sets field value
func (o *SPCampaignOptimizationNotificationAPIRequest) SetCampaignIds(v []string) {
	o.CampaignIds = v
}

func (o SPCampaignOptimizationNotificationAPIRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignIds"] = o.CampaignIds
	return toSerialize, nil
}

type NullableSPCampaignOptimizationNotificationAPIRequest struct {
	value *SPCampaignOptimizationNotificationAPIRequest
	isSet bool
}

func (v NullableSPCampaignOptimizationNotificationAPIRequest) Get() *SPCampaignOptimizationNotificationAPIRequest {
	return v.value
}

func (v *NullableSPCampaignOptimizationNotificationAPIRequest) Set(val *SPCampaignOptimizationNotificationAPIRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSPCampaignOptimizationNotificationAPIRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSPCampaignOptimizationNotificationAPIRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPCampaignOptimizationNotificationAPIRequest(val *SPCampaignOptimizationNotificationAPIRequest) *NullableSPCampaignOptimizationNotificationAPIRequest {
	return &NullableSPCampaignOptimizationNotificationAPIRequest{value: val, isSet: true}
}

func (v NullableSPCampaignOptimizationNotificationAPIRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPCampaignOptimizationNotificationAPIRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
