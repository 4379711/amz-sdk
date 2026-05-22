package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateTargetPromotionGroupsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateTargetPromotionGroupsResponseContent{}

// SponsoredProductsCreateTargetPromotionGroupsResponseContent Response object for creating a target promotion group.
type SponsoredProductsCreateTargetPromotionGroupsResponseContent struct {
	TargetPromotionGroup *SponsoredProductsTargetPromotionGroup `json:"targetPromotionGroup,omitempty"`
}

// NewSponsoredProductsCreateTargetPromotionGroupsResponseContent instantiates a new SponsoredProductsCreateTargetPromotionGroupsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateTargetPromotionGroupsResponseContent() *SponsoredProductsCreateTargetPromotionGroupsResponseContent {
	this := SponsoredProductsCreateTargetPromotionGroupsResponseContent{}
	return &this
}

// NewSponsoredProductsCreateTargetPromotionGroupsResponseContentWithDefaults instantiates a new SponsoredProductsCreateTargetPromotionGroupsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateTargetPromotionGroupsResponseContentWithDefaults() *SponsoredProductsCreateTargetPromotionGroupsResponseContent {
	this := SponsoredProductsCreateTargetPromotionGroupsResponseContent{}
	return &this
}

// GetTargetPromotionGroup returns the TargetPromotionGroup field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupsResponseContent) GetTargetPromotionGroup() SponsoredProductsTargetPromotionGroup {
	if o == nil || IsNil(o.TargetPromotionGroup) {
		var ret SponsoredProductsTargetPromotionGroup
		return ret
	}
	return *o.TargetPromotionGroup
}

// GetTargetPromotionGroupOk returns a tuple with the TargetPromotionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsResponseContent) GetTargetPromotionGroupOk() (*SponsoredProductsTargetPromotionGroup, bool) {
	if o == nil || IsNil(o.TargetPromotionGroup) {
		return nil, false
	}
	return o.TargetPromotionGroup, true
}

// HasTargetPromotionGroup returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsResponseContent) HasTargetPromotionGroup() bool {
	if o != nil && !IsNil(o.TargetPromotionGroup) {
		return true
	}

	return false
}

// SetTargetPromotionGroup gets a reference to the given SponsoredProductsTargetPromotionGroup and assigns it to the TargetPromotionGroup field.
func (o *SponsoredProductsCreateTargetPromotionGroupsResponseContent) SetTargetPromotionGroup(v SponsoredProductsTargetPromotionGroup) {
	o.TargetPromotionGroup = &v
}

func (o SponsoredProductsCreateTargetPromotionGroupsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetPromotionGroup) {
		toSerialize["targetPromotionGroup"] = o.TargetPromotionGroup
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateTargetPromotionGroupsResponseContent struct {
	value *SponsoredProductsCreateTargetPromotionGroupsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupsResponseContent) Get() *SponsoredProductsCreateTargetPromotionGroupsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupsResponseContent) Set(val *SponsoredProductsCreateTargetPromotionGroupsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateTargetPromotionGroupsResponseContent(val *SponsoredProductsCreateTargetPromotionGroupsResponseContent) *NullableSponsoredProductsCreateTargetPromotionGroupsResponseContent {
	return &NullableSponsoredProductsCreateTargetPromotionGroupsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
