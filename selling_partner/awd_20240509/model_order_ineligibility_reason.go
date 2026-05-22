package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderIneligibilityReason type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderIneligibilityReason{}

// OrderIneligibilityReason Represents one ineligibility reason for the order (there can be multiple reasons).
type OrderIneligibilityReason struct {
	// Code for the order ineligibility.
	Code string `json:"code"`
	// Description detailing the ineligibility reason of the order.
	Description string `json:"description"`
}

type _OrderIneligibilityReason OrderIneligibilityReason

// NewOrderIneligibilityReason instantiates a new OrderIneligibilityReason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderIneligibilityReason(code string, description string) *OrderIneligibilityReason {
	this := OrderIneligibilityReason{}
	this.Code = code
	this.Description = description
	return &this
}

// NewOrderIneligibilityReasonWithDefaults instantiates a new OrderIneligibilityReason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderIneligibilityReasonWithDefaults() *OrderIneligibilityReason {
	this := OrderIneligibilityReason{}
	return &this
}

// GetCode returns the Code field value
func (o *OrderIneligibilityReason) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OrderIneligibilityReason) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OrderIneligibilityReason) SetCode(v string) {
	o.Code = v
}

// GetDescription returns the Description field value
func (o *OrderIneligibilityReason) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *OrderIneligibilityReason) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *OrderIneligibilityReason) SetDescription(v string) {
	o.Description = v
}

func (o OrderIneligibilityReason) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

type NullableOrderIneligibilityReason struct {
	value *OrderIneligibilityReason
	isSet bool
}

func (v NullableOrderIneligibilityReason) Get() *OrderIneligibilityReason {
	return v.value
}

func (v *NullableOrderIneligibilityReason) Set(val *OrderIneligibilityReason) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderIneligibilityReason) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderIneligibilityReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderIneligibilityReason(val *OrderIneligibilityReason) *NullableOrderIneligibilityReason {
	return &NullableOrderIneligibilityReason{value: val, isSet: true}
}

func (v NullableOrderIneligibilityReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderIneligibilityReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
