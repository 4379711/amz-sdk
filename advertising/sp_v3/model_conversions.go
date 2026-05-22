package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the Conversions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Conversions{}

// Conversions Conversions benchmark.
type Conversions struct {
	// lower bound.
	Lower *int32 `json:"lower,omitempty"`
	// upper bound.
	Upper *int32 `json:"upper,omitempty"`
}

// NewConversions instantiates a new Conversions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConversions() *Conversions {
	this := Conversions{}
	return &this
}

// NewConversionsWithDefaults instantiates a new Conversions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConversionsWithDefaults() *Conversions {
	this := Conversions{}
	return &this
}

// GetLower returns the Lower field value if set, zero value otherwise.
func (o *Conversions) GetLower() int32 {
	if o == nil || IsNil(o.Lower) {
		var ret int32
		return ret
	}
	return *o.Lower
}

// GetLowerOk returns a tuple with the Lower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversions) GetLowerOk() (*int32, bool) {
	if o == nil || IsNil(o.Lower) {
		return nil, false
	}
	return o.Lower, true
}

// HasLower returns a boolean if a field has been set.
func (o *Conversions) HasLower() bool {
	if o != nil && !IsNil(o.Lower) {
		return true
	}

	return false
}

// SetLower gets a reference to the given int32 and assigns it to the Lower field.
func (o *Conversions) SetLower(v int32) {
	o.Lower = &v
}

// GetUpper returns the Upper field value if set, zero value otherwise.
func (o *Conversions) GetUpper() int32 {
	if o == nil || IsNil(o.Upper) {
		var ret int32
		return ret
	}
	return *o.Upper
}

// GetUpperOk returns a tuple with the Upper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversions) GetUpperOk() (*int32, bool) {
	if o == nil || IsNil(o.Upper) {
		return nil, false
	}
	return o.Upper, true
}

// HasUpper returns a boolean if a field has been set.
func (o *Conversions) HasUpper() bool {
	if o != nil && !IsNil(o.Upper) {
		return true
	}

	return false
}

// SetUpper gets a reference to the given int32 and assigns it to the Upper field.
func (o *Conversions) SetUpper(v int32) {
	o.Upper = &v
}

func (o Conversions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lower) {
		toSerialize["lower"] = o.Lower
	}
	if !IsNil(o.Upper) {
		toSerialize["upper"] = o.Upper
	}
	return toSerialize, nil
}

type NullableConversions struct {
	value *Conversions
	isSet bool
}

func (v NullableConversions) Get() *Conversions {
	return v.value
}

func (v *NullableConversions) Set(val *Conversions) {
	v.value = val
	v.isSet = true
}

func (v NullableConversions) IsSet() bool {
	return v.isSet
}

func (v *NullableConversions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConversions(val *Conversions) *NullableConversions {
	return &NullableConversions{value: val, isSet: true}
}

func (v NullableConversions) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableConversions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
