package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the SkuIneligibilityReason type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuIneligibilityReason{}

// SkuIneligibilityReason Represents the ineligibility reason for one SKU.
type SkuIneligibilityReason struct {
	// Code for the SKU ineligibility.
	Code string `json:"code"`
	// Detailed description of the SKU ineligibility.
	Description string `json:"description"`
}

type _SkuIneligibilityReason SkuIneligibilityReason

// NewSkuIneligibilityReason instantiates a new SkuIneligibilityReason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuIneligibilityReason(code string, description string) *SkuIneligibilityReason {
	this := SkuIneligibilityReason{}
	this.Code = code
	this.Description = description
	return &this
}

// NewSkuIneligibilityReasonWithDefaults instantiates a new SkuIneligibilityReason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuIneligibilityReasonWithDefaults() *SkuIneligibilityReason {
	this := SkuIneligibilityReason{}
	return &this
}

// GetCode returns the Code field value
func (o *SkuIneligibilityReason) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SkuIneligibilityReason) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SkuIneligibilityReason) SetCode(v string) {
	o.Code = v
}

// GetDescription returns the Description field value
func (o *SkuIneligibilityReason) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SkuIneligibilityReason) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SkuIneligibilityReason) SetDescription(v string) {
	o.Description = v
}

func (o SkuIneligibilityReason) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

type NullableSkuIneligibilityReason struct {
	value *SkuIneligibilityReason
	isSet bool
}

func (v NullableSkuIneligibilityReason) Get() *SkuIneligibilityReason {
	return v.value
}

func (v *NullableSkuIneligibilityReason) Set(val *SkuIneligibilityReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuIneligibilityReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuIneligibilityReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuIneligibilityReason(val *SkuIneligibilityReason) *NullableSkuIneligibilityReason {
	return &NullableSkuIneligibilityReason{value: val, isSet: true}
}

func (v NullableSkuIneligibilityReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSkuIneligibilityReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
