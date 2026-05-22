package easy_ship_20220323

import (
	"github.com/bytedance/sonic"
)

// checks if the Weight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Weight{}

// Weight The weight of the scheduled package
type Weight struct {
	// The weight of the package.
	Value *float32      `json:"value,omitempty"`
	Unit  *UnitOfWeight `json:"unit,omitempty"`
}

// NewWeight instantiates a new Weight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeight() *Weight {
	this := Weight{}
	return &this
}

// NewWeightWithDefaults instantiates a new Weight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeightWithDefaults() *Weight {
	this := Weight{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Weight) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weight) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Weight) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *Weight) SetValue(v float32) {
	o.Value = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Weight) GetUnit() UnitOfWeight {
	if o == nil || IsNil(o.Unit) {
		var ret UnitOfWeight
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weight) GetUnitOk() (*UnitOfWeight, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Weight) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given UnitOfWeight and assigns it to the Unit field.
func (o *Weight) SetUnit(v UnitOfWeight) {
	o.Unit = &v
}

func (o Weight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	return toSerialize, nil
}

type NullableWeight struct {
	value *Weight
	isSet bool
}

func (v NullableWeight) Get() *Weight {
	return v.value
}

func (v *NullableWeight) Set(val *Weight) {
	v.value = val
	v.isSet = true
}

func (v NullableWeight) IsSet() bool {
	return v.isSet
}

func (v *NullableWeight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeight(val *Weight) *NullableWeight {
	return &NullableWeight{value: val, isSet: true}
}

func (v NullableWeight) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableWeight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
