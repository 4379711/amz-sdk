package vendor_shipments

import (
	"github.com/bytedance/sonic"
)

// checks if the Volume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Volume{}

// Volume The volume of the shipment.
type Volume struct {
	// The unit of measurement.
	UnitOfMeasure string `json:"unitOfMeasure"`
	// A decimal number with no loss of precision. Useful when precision loss is unacceptable, as with currencies. Follows RFC7159 for number representation. <br>**Pattern** : `^-?(0|([1-9]\\d*))(\\.\\d+)?([eE][+-]?\\d+)?$`.
	Value string `json:"value"`
}

type _Volume Volume

// NewVolume instantiates a new Volume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolume(unitOfMeasure string, value string) *Volume {
	this := Volume{}
	this.UnitOfMeasure = unitOfMeasure
	this.Value = value
	return &this
}

// NewVolumeWithDefaults instantiates a new Volume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeWithDefaults() *Volume {
	this := Volume{}
	return &this
}

// GetUnitOfMeasure returns the UnitOfMeasure field value
func (o *Volume) GetUnitOfMeasure() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value
// and a boolean to check if the value has been set.
func (o *Volume) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitOfMeasure, true
}

// SetUnitOfMeasure sets field value
func (o *Volume) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = v
}

// GetValue returns the Value field value
func (o *Volume) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Volume) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Volume) SetValue(v string) {
	o.Value = v
}

func (o Volume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["unitOfMeasure"] = o.UnitOfMeasure
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableVolume struct {
	value *Volume
	isSet bool
}

func (v NullableVolume) Get() *Volume {
	return v.value
}

func (v *NullableVolume) Set(val *Volume) {
	v.value = val
	v.isSet = true
}

func (v NullableVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolume(val *Volume) *NullableVolume {
	return &NullableVolume{value: val, isSet: true}
}

func (v NullableVolume) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
