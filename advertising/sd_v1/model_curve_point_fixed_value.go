package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CurvePointFixedValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurvePointFixedValue{}

// CurvePointFixedValue struct for CurvePointFixedValue
type CurvePointFixedValue struct {
	Value interface{} `json:"value,omitempty"`
}

// NewCurvePointFixedValue instantiates a new CurvePointFixedValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurvePointFixedValue() *CurvePointFixedValue {
	this := CurvePointFixedValue{}
	return &this
}

// NewCurvePointFixedValueWithDefaults instantiates a new CurvePointFixedValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurvePointFixedValueWithDefaults() *CurvePointFixedValue {
	this := CurvePointFixedValue{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurvePointFixedValue) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurvePointFixedValue) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CurvePointFixedValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *CurvePointFixedValue) SetValue(v interface{}) {
	o.Value = v
}

func (o CurvePointFixedValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCurvePointFixedValue struct {
	value *CurvePointFixedValue
	isSet bool
}

func (v NullableCurvePointFixedValue) Get() *CurvePointFixedValue {
	return v.value
}

func (v *NullableCurvePointFixedValue) Set(val *CurvePointFixedValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCurvePointFixedValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCurvePointFixedValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurvePointFixedValue(val *CurvePointFixedValue) *NullableCurvePointFixedValue {
	return &NullableCurvePointFixedValue{value: val, isSet: true}
}

func (v NullableCurvePointFixedValue) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCurvePointFixedValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
