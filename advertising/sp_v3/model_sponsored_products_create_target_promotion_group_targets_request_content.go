package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent{}

// SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent Request object for creating target promotion group targets in a target promotion group.
type SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent struct {
	// The id of the target promotion group where the targets are being added.
	TargetPromotionGroupId string `json:"targetPromotionGroupId"`
	// List of targets to be added to the target promotion group.
	Targets []SponsoredProductsCreateTargetRequest `json:"targets,omitempty"`
}

type _SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent

// NewSponsoredProductsCreateTargetPromotionGroupTargetsRequestContent instantiates a new SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateTargetPromotionGroupTargetsRequestContent(targetPromotionGroupId string) *SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent {
	this := SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent{}
	this.TargetPromotionGroupId = targetPromotionGroupId
	return &this
}

// NewSponsoredProductsCreateTargetPromotionGroupTargetsRequestContentWithDefaults instantiates a new SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateTargetPromotionGroupTargetsRequestContentWithDefaults() *SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent {
	this := SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent{}
	return &this
}

// GetTargetPromotionGroupId returns the TargetPromotionGroupId field value
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent) GetTargetPromotionGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetPromotionGroupId
}

// GetTargetPromotionGroupIdOk returns a tuple with the TargetPromotionGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent) GetTargetPromotionGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPromotionGroupId, true
}

// SetTargetPromotionGroupId sets field value
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent) SetTargetPromotionGroupId(v string) {
	o.TargetPromotionGroupId = v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent) GetTargets() []SponsoredProductsCreateTargetRequest {
	if o == nil || IsNil(o.Targets) {
		var ret []SponsoredProductsCreateTargetRequest
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent) GetTargetsOk() ([]SponsoredProductsCreateTargetRequest, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []SponsoredProductsCreateTargetRequest and assigns it to the Targets field.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent) SetTargets(v []SponsoredProductsCreateTargetRequest) {
	o.Targets = v
}

func (o SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetPromotionGroupId"] = o.TargetPromotionGroupId
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateTargetPromotionGroupTargetsRequestContent struct {
	value *SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupTargetsRequestContent) Get() *SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupTargetsRequestContent) Set(val *SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupTargetsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupTargetsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateTargetPromotionGroupTargetsRequestContent(val *SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent) *NullableSponsoredProductsCreateTargetPromotionGroupTargetsRequestContent {
	return &NullableSponsoredProductsCreateTargetPromotionGroupTargetsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupTargetsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupTargetsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
