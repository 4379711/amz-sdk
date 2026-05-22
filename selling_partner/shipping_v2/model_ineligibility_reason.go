package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the IneligibilityReason type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IneligibilityReason{}

// IneligibilityReason The reason why a shipping service offering is ineligible.
type IneligibilityReason struct {
	Code IneligibilityReasonCode `json:"code"`
	// The ineligibility reason.
	Message string `json:"message"`
}

type _IneligibilityReason IneligibilityReason

// NewIneligibilityReason instantiates a new IneligibilityReason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIneligibilityReason(code IneligibilityReasonCode, message string) *IneligibilityReason {
	this := IneligibilityReason{}
	this.Code = code
	this.Message = message
	return &this
}

// NewIneligibilityReasonWithDefaults instantiates a new IneligibilityReason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIneligibilityReasonWithDefaults() *IneligibilityReason {
	this := IneligibilityReason{}
	return &this
}

// GetCode returns the Code field value
func (o *IneligibilityReason) GetCode() IneligibilityReasonCode {
	if o == nil {
		var ret IneligibilityReasonCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *IneligibilityReason) GetCodeOk() (*IneligibilityReasonCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *IneligibilityReason) SetCode(v IneligibilityReasonCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *IneligibilityReason) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *IneligibilityReason) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *IneligibilityReason) SetMessage(v string) {
	o.Message = v
}

func (o IneligibilityReason) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableIneligibilityReason struct {
	value *IneligibilityReason
	isSet bool
}

func (v NullableIneligibilityReason) Get() *IneligibilityReason {
	return v.value
}

func (v *NullableIneligibilityReason) Set(val *IneligibilityReason) {
	v.value = val
	v.isSet = true
}

func (v NullableIneligibilityReason) IsSet() bool {
	return v.isSet
}

func (v *NullableIneligibilityReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIneligibilityReason(val *IneligibilityReason) *NullableIneligibilityReason {
	return &NullableIneligibilityReason{value: val, isSet: true}
}

func (v NullableIneligibilityReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableIneligibilityReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
