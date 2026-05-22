package vendor_shipments

import (
	"github.com/bytedance/sonic"
)

// checks if the Weight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Weight{}

// Weight The weight of the shipment.
type Weight struct {
	// The unit of measurement.
	UnitOfMeasure string `json:"unitOfMeasure"`
	// A decimal number with no loss of precision. Useful when precision loss is unacceptable, as with currencies. Follows RFC7159 for number representation. <br>**Pattern** : `^-?(0|([1-9]\\d*))(\\.\\d+)?([eE][+-]?\\d+)?$`.
	Value string `json:"value"`
}

type _Weight Weight

// NewWeight instantiates a new Weight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeight(unitOfMeasure string, value string) *Weight {
	this := Weight{}
	this.UnitOfMeasure = unitOfMeasure
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

// GetUnitOfMeasure returns the UnitOfMeasure field value
func (o *Weight) GetUnitOfMeasure() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value
// and a boolean to check if the value has been set.
func (o *Weight) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitOfMeasure, true
}

// SetUnitOfMeasure sets field value
func (o *Weight) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = v
}

// GetValue returns the Value field value
func (o *Weight) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Weight) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Weight) SetValue(v string) {
	o.Value = v
}

func (o Weight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["unitOfMeasure"] = o.UnitOfMeasure
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
