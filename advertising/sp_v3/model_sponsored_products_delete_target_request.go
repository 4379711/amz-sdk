package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteTargetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteTargetRequest{}

// SponsoredProductsDeleteTargetRequest struct for SponsoredProductsDeleteTargetRequest
type SponsoredProductsDeleteTargetRequest struct {
	// Entity object identifier
	TargetPromotionGroupId string `json:"targetPromotionGroupId"`
	// Entity object identifier
	TargetId string `json:"targetId"`
	// Entity object identifier
	ManualTargetingAdGroupId string `json:"manualTargetingAdGroupId"`
	// Entity object identifier
	AutoTargetingCampaignAdGroupId string `json:"autoTargetingCampaignAdGroupId"`
}

type _SponsoredProductsDeleteTargetRequest SponsoredProductsDeleteTargetRequest

// NewSponsoredProductsDeleteTargetRequest instantiates a new SponsoredProductsDeleteTargetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteTargetRequest(targetPromotionGroupId string, targetId string, manualTargetingAdGroupId string, autoTargetingCampaignAdGroupId string) *SponsoredProductsDeleteTargetRequest {
	this := SponsoredProductsDeleteTargetRequest{}
	this.TargetPromotionGroupId = targetPromotionGroupId
	this.TargetId = targetId
	this.ManualTargetingAdGroupId = manualTargetingAdGroupId
	this.AutoTargetingCampaignAdGroupId = autoTargetingCampaignAdGroupId
	return &this
}

// NewSponsoredProductsDeleteTargetRequestWithDefaults instantiates a new SponsoredProductsDeleteTargetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteTargetRequestWithDefaults() *SponsoredProductsDeleteTargetRequest {
	this := SponsoredProductsDeleteTargetRequest{}
	return &this
}

// GetTargetPromotionGroupId returns the TargetPromotionGroupId field value
func (o *SponsoredProductsDeleteTargetRequest) GetTargetPromotionGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetPromotionGroupId
}

// GetTargetPromotionGroupIdOk returns a tuple with the TargetPromotionGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteTargetRequest) GetTargetPromotionGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPromotionGroupId, true
}

// SetTargetPromotionGroupId sets field value
func (o *SponsoredProductsDeleteTargetRequest) SetTargetPromotionGroupId(v string) {
	o.TargetPromotionGroupId = v
}

// GetTargetId returns the TargetId field value
func (o *SponsoredProductsDeleteTargetRequest) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteTargetRequest) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *SponsoredProductsDeleteTargetRequest) SetTargetId(v string) {
	o.TargetId = v
}

// GetManualTargetingAdGroupId returns the ManualTargetingAdGroupId field value
func (o *SponsoredProductsDeleteTargetRequest) GetManualTargetingAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ManualTargetingAdGroupId
}

// GetManualTargetingAdGroupIdOk returns a tuple with the ManualTargetingAdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteTargetRequest) GetManualTargetingAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManualTargetingAdGroupId, true
}

// SetManualTargetingAdGroupId sets field value
func (o *SponsoredProductsDeleteTargetRequest) SetManualTargetingAdGroupId(v string) {
	o.ManualTargetingAdGroupId = v
}

// GetAutoTargetingCampaignAdGroupId returns the AutoTargetingCampaignAdGroupId field value
func (o *SponsoredProductsDeleteTargetRequest) GetAutoTargetingCampaignAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AutoTargetingCampaignAdGroupId
}

// GetAutoTargetingCampaignAdGroupIdOk returns a tuple with the AutoTargetingCampaignAdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteTargetRequest) GetAutoTargetingCampaignAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoTargetingCampaignAdGroupId, true
}

// SetAutoTargetingCampaignAdGroupId sets field value
func (o *SponsoredProductsDeleteTargetRequest) SetAutoTargetingCampaignAdGroupId(v string) {
	o.AutoTargetingCampaignAdGroupId = v
}

func (o SponsoredProductsDeleteTargetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetPromotionGroupId"] = o.TargetPromotionGroupId
	toSerialize["targetId"] = o.TargetId
	toSerialize["manualTargetingAdGroupId"] = o.ManualTargetingAdGroupId
	toSerialize["autoTargetingCampaignAdGroupId"] = o.AutoTargetingCampaignAdGroupId
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteTargetRequest struct {
	value *SponsoredProductsDeleteTargetRequest
	isSet bool
}

func (v NullableSponsoredProductsDeleteTargetRequest) Get() *SponsoredProductsDeleteTargetRequest {
	return v.value
}

func (v *NullableSponsoredProductsDeleteTargetRequest) Set(val *SponsoredProductsDeleteTargetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteTargetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteTargetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteTargetRequest(val *SponsoredProductsDeleteTargetRequest) *NullableSponsoredProductsDeleteTargetRequest {
	return &NullableSponsoredProductsDeleteTargetRequest{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteTargetRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteTargetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
