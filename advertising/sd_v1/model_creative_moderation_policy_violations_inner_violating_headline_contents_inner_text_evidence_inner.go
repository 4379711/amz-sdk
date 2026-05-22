package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner{}

// CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner struct for CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner
type CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner struct {
	// The specific text determined to violate the specified policy in reviewedText
	ViolatingText         *string                                                                                                      `json:"violatingText,omitempty"`
	ViolatingTextPosition *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition `json:"violatingTextPosition,omitempty"`
}

// NewCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner instantiates a new CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner() *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner {
	this := CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner{}
	return &this
}

// NewCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerWithDefaults instantiates a new CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerWithDefaults() *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner {
	this := CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner{}
	return &this
}

// GetViolatingText returns the ViolatingText field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner) GetViolatingText() string {
	if o == nil || IsNil(o.ViolatingText) {
		var ret string
		return ret
	}
	return *o.ViolatingText
}

// GetViolatingTextOk returns a tuple with the ViolatingText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner) GetViolatingTextOk() (*string, bool) {
	if o == nil || IsNil(o.ViolatingText) {
		return nil, false
	}
	return o.ViolatingText, true
}

// HasViolatingText returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner) HasViolatingText() bool {
	if o != nil && !IsNil(o.ViolatingText) {
		return true
	}

	return false
}

// SetViolatingText gets a reference to the given string and assigns it to the ViolatingText field.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner) SetViolatingText(v string) {
	o.ViolatingText = &v
}

// GetViolatingTextPosition returns the ViolatingTextPosition field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner) GetViolatingTextPosition() CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition {
	if o == nil || IsNil(o.ViolatingTextPosition) {
		var ret CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition
		return ret
	}
	return *o.ViolatingTextPosition
}

// GetViolatingTextPositionOk returns a tuple with the ViolatingTextPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner) GetViolatingTextPositionOk() (*CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition, bool) {
	if o == nil || IsNil(o.ViolatingTextPosition) {
		return nil, false
	}
	return o.ViolatingTextPosition, true
}

// HasViolatingTextPosition returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner) HasViolatingTextPosition() bool {
	if o != nil && !IsNil(o.ViolatingTextPosition) {
		return true
	}

	return false
}

// SetViolatingTextPosition gets a reference to the given CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition and assigns it to the ViolatingTextPosition field.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner) SetViolatingTextPosition(v CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition) {
	o.ViolatingTextPosition = &v
}

func (o CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ViolatingText) {
		toSerialize["violatingText"] = o.ViolatingText
	}
	if !IsNil(o.ViolatingTextPosition) {
		toSerialize["violatingTextPosition"] = o.ViolatingTextPosition
	}
	return toSerialize, nil
}

type NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner struct {
	value *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner
	isSet bool
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner) Get() *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner {
	return v.value
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner) Set(val *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner(val *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner) *NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner {
	return &NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner{value: val, isSet: true}
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
