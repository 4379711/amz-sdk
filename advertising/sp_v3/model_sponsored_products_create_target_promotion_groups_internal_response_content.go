package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent{}

// SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent Response object for creating a target promotion group.
type SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent struct {
	TargetPromotionGroup *SponsoredProductsTargetPromotionGroup `json:"targetPromotionGroup,omitempty"`
}

// NewSponsoredProductsCreateTargetPromotionGroupsInternalResponseContent instantiates a new SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateTargetPromotionGroupsInternalResponseContent() *SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent {
	this := SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent{}
	return &this
}

// NewSponsoredProductsCreateTargetPromotionGroupsInternalResponseContentWithDefaults instantiates a new SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateTargetPromotionGroupsInternalResponseContentWithDefaults() *SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent {
	this := SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent{}
	return &this
}

// GetTargetPromotionGroup returns the TargetPromotionGroup field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent) GetTargetPromotionGroup() SponsoredProductsTargetPromotionGroup {
	if o == nil || IsNil(o.TargetPromotionGroup) {
		var ret SponsoredProductsTargetPromotionGroup
		return ret
	}
	return *o.TargetPromotionGroup
}

// GetTargetPromotionGroupOk returns a tuple with the TargetPromotionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent) GetTargetPromotionGroupOk() (*SponsoredProductsTargetPromotionGroup, bool) {
	if o == nil || IsNil(o.TargetPromotionGroup) {
		return nil, false
	}
	return o.TargetPromotionGroup, true
}

// HasTargetPromotionGroup returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent) HasTargetPromotionGroup() bool {
	if o != nil && !IsNil(o.TargetPromotionGroup) {
		return true
	}

	return false
}

// SetTargetPromotionGroup gets a reference to the given SponsoredProductsTargetPromotionGroup and assigns it to the TargetPromotionGroup field.
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent) SetTargetPromotionGroup(v SponsoredProductsTargetPromotionGroup) {
	o.TargetPromotionGroup = &v
}

func (o SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetPromotionGroup) {
		toSerialize["targetPromotionGroup"] = o.TargetPromotionGroup
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateTargetPromotionGroupsInternalResponseContent struct {
	value *SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupsInternalResponseContent) Get() *SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupsInternalResponseContent) Set(val *SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupsInternalResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupsInternalResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateTargetPromotionGroupsInternalResponseContent(val *SponsoredProductsCreateTargetPromotionGroupsInternalResponseContent) *NullableSponsoredProductsCreateTargetPromotionGroupsInternalResponseContent {
	return &NullableSponsoredProductsCreateTargetPromotionGroupsInternalResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupsInternalResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupsInternalResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
