package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the LiquidVolume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiquidVolume{}

// LiquidVolume Liquid volume.
type LiquidVolume struct {
	// The unit of measurement.
	Unit string `json:"Unit"`
	// The measurement value.
	Value float32 `json:"Value"`
}

type _LiquidVolume LiquidVolume

// NewLiquidVolume instantiates a new LiquidVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiquidVolume(unit string, value float32) *LiquidVolume {
	this := LiquidVolume{}
	this.Unit = unit
	this.Value = value
	return &this
}

// NewLiquidVolumeWithDefaults instantiates a new LiquidVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiquidVolumeWithDefaults() *LiquidVolume {
	this := LiquidVolume{}
	return &this
}

// GetUnit returns the Unit field value
func (o *LiquidVolume) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *LiquidVolume) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *LiquidVolume) SetUnit(v string) {
	o.Unit = v
}

// GetValue returns the Value field value
func (o *LiquidVolume) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *LiquidVolume) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *LiquidVolume) SetValue(v float32) {
	o.Value = v
}

func (o LiquidVolume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Unit"] = o.Unit
	toSerialize["Value"] = o.Value
	return toSerialize, nil
}

type NullableLiquidVolume struct {
	value *LiquidVolume
	isSet bool
}

func (v NullableLiquidVolume) Get() *LiquidVolume {
	return v.value
}

func (v *NullableLiquidVolume) Set(val *LiquidVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableLiquidVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableLiquidVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiquidVolume(val *LiquidVolume) *NullableLiquidVolume {
	return &NullableLiquidVolume{value: val, isSet: true}
}

func (v NullableLiquidVolume) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLiquidVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
