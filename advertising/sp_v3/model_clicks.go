package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the Clicks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Clicks{}

// Clicks Clicks benchmark.
type Clicks struct {
	// lower bound.
	Lower *int32 `json:"lower,omitempty"`
	// upper bound.
	Upper *int32 `json:"upper,omitempty"`
}

// NewClicks instantiates a new Clicks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClicks() *Clicks {
	this := Clicks{}
	return &this
}

// NewClicksWithDefaults instantiates a new Clicks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClicksWithDefaults() *Clicks {
	this := Clicks{}
	return &this
}

// GetLower returns the Lower field value if set, zero value otherwise.
func (o *Clicks) GetLower() int32 {
	if o == nil || IsNil(o.Lower) {
		var ret int32
		return ret
	}
	return *o.Lower
}

// GetLowerOk returns a tuple with the Lower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Clicks) GetLowerOk() (*int32, bool) {
	if o == nil || IsNil(o.Lower) {
		return nil, false
	}
	return o.Lower, true
}

// HasLower returns a boolean if a field has been set.
func (o *Clicks) HasLower() bool {
	if o != nil && !IsNil(o.Lower) {
		return true
	}

	return false
}

// SetLower gets a reference to the given int32 and assigns it to the Lower field.
func (o *Clicks) SetLower(v int32) {
	o.Lower = &v
}

// GetUpper returns the Upper field value if set, zero value otherwise.
func (o *Clicks) GetUpper() int32 {
	if o == nil || IsNil(o.Upper) {
		var ret int32
		return ret
	}
	return *o.Upper
}

// GetUpperOk returns a tuple with the Upper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Clicks) GetUpperOk() (*int32, bool) {
	if o == nil || IsNil(o.Upper) {
		return nil, false
	}
	return o.Upper, true
}

// HasUpper returns a boolean if a field has been set.
func (o *Clicks) HasUpper() bool {
	if o != nil && !IsNil(o.Upper) {
		return true
	}

	return false
}

// SetUpper gets a reference to the given int32 and assigns it to the Upper field.
func (o *Clicks) SetUpper(v int32) {
	o.Upper = &v
}

func (o Clicks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lower) {
		toSerialize["lower"] = o.Lower
	}
	if !IsNil(o.Upper) {
		toSerialize["upper"] = o.Upper
	}
	return toSerialize, nil
}

type NullableClicks struct {
	value *Clicks
	isSet bool
}

func (v NullableClicks) Get() *Clicks {
	return v.value
}

func (v *NullableClicks) Set(val *Clicks) {
	v.value = val
	v.isSet = true
}

func (v NullableClicks) IsSet() bool {
	return v.isSet
}

func (v *NullableClicks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClicks(val *Clicks) *NullableClicks {
	return &NullableClicks{value: val, isSet: true}
}

func (v NullableClicks) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableClicks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
