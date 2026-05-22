package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the Measurement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Measurement{}

// Measurement Measurement information for an order item.
type Measurement struct {
	// The unit of measure.
	Unit string `json:"Unit"`
	// The measurement value.
	Value float32 `json:"Value"`
}

type _Measurement Measurement

// NewMeasurement instantiates a new Measurement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeasurement(unit string, value float32) *Measurement {
	this := Measurement{}
	this.Unit = unit
	this.Value = value
	return &this
}

// NewMeasurementWithDefaults instantiates a new Measurement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeasurementWithDefaults() *Measurement {
	this := Measurement{}
	return &this
}

// GetUnit returns the Unit field value
func (o *Measurement) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *Measurement) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *Measurement) SetUnit(v string) {
	o.Unit = v
}

// GetValue returns the Value field value
func (o *Measurement) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Measurement) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Measurement) SetValue(v float32) {
	o.Value = v
}

func (o Measurement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Unit"] = o.Unit
	toSerialize["Value"] = o.Value
	return toSerialize, nil
}

type NullableMeasurement struct {
	value *Measurement
	isSet bool
}

func (v NullableMeasurement) Get() *Measurement {
	return v.value
}

func (v *NullableMeasurement) Set(val *Measurement) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurement) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurement(val *Measurement) *NullableMeasurement {
	return &NullableMeasurement{value: val, isSet: true}
}

func (v NullableMeasurement) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMeasurement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
