package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the Weight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Weight{}

// Weight The weight in the units indicated.
type Weight struct {
	// The unit of measurement.
	Unit string `json:"unit"`
	// The measurement value.
	Value float32 `json:"value"`
}

type _Weight Weight

// NewWeight instantiates a new Weight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeight(unit string, value float32) *Weight {
	this := Weight{}
	this.Unit = unit
	this.Value = value
	return &this
}

// NewWeightWithDefaults instantiates a new Weight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeightWithDefaults() *Weight {
	this := Weight{}
	return &this
}

// GetUnit returns the Unit field value
func (o *Weight) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *Weight) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *Weight) SetUnit(v string) {
	o.Unit = v
}

// GetValue returns the Value field value
func (o *Weight) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Weight) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Weight) SetValue(v float32) {
	o.Value = v
}

func (o Weight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["unit"] = o.Unit
	toSerialize["value"] = o.Value
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
