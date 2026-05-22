package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the Length type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Length{}

// Length The length.
type Length struct {
	// The value in units.
	Value *float32      `json:"value,omitempty"`
	Unit  *UnitOfLength `json:"unit,omitempty"`
}

// NewLength instantiates a new Length object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLength() *Length {
	this := Length{}
	return &this
}

// NewLengthWithDefaults instantiates a new Length object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLengthWithDefaults() *Length {
	this := Length{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Length) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Length) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Length) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *Length) SetValue(v float32) {
	o.Value = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Length) GetUnit() UnitOfLength {
	if o == nil || IsNil(o.Unit) {
		var ret UnitOfLength
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Length) GetUnitOk() (*UnitOfLength, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Length) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given UnitOfLength and assigns it to the Unit field.
func (o *Length) SetUnit(v UnitOfLength) {
	o.Unit = &v
}

func (o Length) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	return toSerialize, nil
}

type NullableLength struct {
	value *Length
	isSet bool
}

func (v NullableLength) Get() *Length {
	return v.value
}

func (v *NullableLength) Set(val *Length) {
	v.value = val
	v.isSet = true
}

func (v NullableLength) IsSet() bool {
	return v.isSet
}

func (v *NullableLength) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLength(val *Length) *NullableLength {
	return &NullableLength{value: val, isSet: true}
}

func (v NullableLength) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLength) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
