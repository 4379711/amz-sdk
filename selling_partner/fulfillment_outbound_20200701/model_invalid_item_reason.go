package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the InvalidItemReason type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvalidItemReason{}

// InvalidItemReason The reason that the item is invalid for return.
type InvalidItemReason struct {
	InvalidItemReasonCode InvalidItemReasonCode `json:"invalidItemReasonCode"`
	// A human readable description of the invalid item reason code.
	Description string `json:"description"`
}

type _InvalidItemReason InvalidItemReason

// NewInvalidItemReason instantiates a new InvalidItemReason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidItemReason(invalidItemReasonCode InvalidItemReasonCode, description string) *InvalidItemReason {
	this := InvalidItemReason{}
	this.InvalidItemReasonCode = invalidItemReasonCode
	this.Description = description
	return &this
}

// NewInvalidItemReasonWithDefaults instantiates a new InvalidItemReason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidItemReasonWithDefaults() *InvalidItemReason {
	this := InvalidItemReason{}
	return &this
}

// GetInvalidItemReasonCode returns the InvalidItemReasonCode field value
func (o *InvalidItemReason) GetInvalidItemReasonCode() InvalidItemReasonCode {
	if o == nil {
		var ret InvalidItemReasonCode
		return ret
	}

	return o.InvalidItemReasonCode
}

// GetInvalidItemReasonCodeOk returns a tuple with the InvalidItemReasonCode field value
// and a boolean to check if the value has been set.
func (o *InvalidItemReason) GetInvalidItemReasonCodeOk() (*InvalidItemReasonCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvalidItemReasonCode, true
}

// SetInvalidItemReasonCode sets field value
func (o *InvalidItemReason) SetInvalidItemReasonCode(v InvalidItemReasonCode) {
	o.InvalidItemReasonCode = v
}

// GetDescription returns the Description field value
func (o *InvalidItemReason) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InvalidItemReason) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *InvalidItemReason) SetDescription(v string) {
	o.Description = v
}

func (o InvalidItemReason) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invalidItemReasonCode"] = o.InvalidItemReasonCode
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

type NullableInvalidItemReason struct {
	value *InvalidItemReason
	isSet bool
}

func (v NullableInvalidItemReason) Get() *InvalidItemReason {
	return v.value
}

func (v *NullableInvalidItemReason) Set(val *InvalidItemReason) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidItemReason) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidItemReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidItemReason(val *InvalidItemReason) *NullableInvalidItemReason {
	return &NullableInvalidItemReason{value: val, isSet: true}
}

func (v NullableInvalidItemReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInvalidItemReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
