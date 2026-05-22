package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner{}

// CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner struct for CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner
type CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner struct {
	ViolatingImageCrop *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop `json:"violatingImageCrop,omitempty"`
}

// NewCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner instantiates a new CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner() *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner {
	this := CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner{}
	return &this
}

// NewCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerWithDefaults instantiates a new CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerWithDefaults() *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner {
	this := CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner{}
	return &this
}

// GetViolatingImageCrop returns the ViolatingImageCrop field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner) GetViolatingImageCrop() CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop {
	if o == nil || IsNil(o.ViolatingImageCrop) {
		var ret CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop
		return ret
	}
	return *o.ViolatingImageCrop
}

// GetViolatingImageCropOk returns a tuple with the ViolatingImageCrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner) GetViolatingImageCropOk() (*CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop, bool) {
	if o == nil || IsNil(o.ViolatingImageCrop) {
		return nil, false
	}
	return o.ViolatingImageCrop, true
}

// HasViolatingImageCrop returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner) HasViolatingImageCrop() bool {
	if o != nil && !IsNil(o.ViolatingImageCrop) {
		return true
	}

	return false
}

// SetViolatingImageCrop gets a reference to the given CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop and assigns it to the ViolatingImageCrop field.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner) SetViolatingImageCrop(v CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) {
	o.ViolatingImageCrop = &v
}

func (o CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ViolatingImageCrop) {
		toSerialize["violatingImageCrop"] = o.ViolatingImageCrop
	}
	return toSerialize, nil
}

type NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner struct {
	value *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner
	isSet bool
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner) Get() *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner {
	return v.value
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner) Set(val *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner(val *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner) *NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner {
	return &NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner{value: val, isSet: true}
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
