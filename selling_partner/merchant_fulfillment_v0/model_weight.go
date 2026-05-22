package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the Weight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Weight{}

// Weight The weight.
type Weight struct {
	// The weight value.
	Value float64      `json:"Value"`
	Unit  UnitOfWeight `json:"Unit"`
}

type _Weight Weight

// NewWeight instantiates a new Weight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeight(value float64, unit UnitOfWeight) *Weight {
	this := Weight{}
	this.Value = value
	this.Unit = unit
	return &this
}

// NewWeightWithDefaults instantiates a new Weight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeightWithDefaults() *Weight {
	this := Weight{}
	return &this
}

// GetValue returns the Value field value
func (o *Weight) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Weight) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Weight) SetValue(v float64) {
	o.Value = v
}

// GetUnit returns the Unit field value
func (o *Weight) GetUnit() UnitOfWeight {
	if o == nil {
		var ret UnitOfWeight
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *Weight) GetUnitOk() (*UnitOfWeight, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *Weight) SetUnit(v UnitOfWeight) {
	o.Unit = v
}

func (o Weight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Value"] = o.Value
	toSerialize["Unit"] = o.Unit
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
