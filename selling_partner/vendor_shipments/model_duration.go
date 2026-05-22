package vendor_shipments

import (
	"github.com/bytedance/sonic"
)

// checks if the Duration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Duration{}

// Duration Duration after manufacturing date during which the product is valid for consumption.
type Duration struct {
	// Unit for duration.
	DurationUnit string `json:"durationUnit"`
	// Value for the duration in terms of the durationUnit.
	DurationValue int32 `json:"durationValue"`
}

type _Duration Duration

// NewDuration instantiates a new Duration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDuration(durationUnit string, durationValue int32) *Duration {
	this := Duration{}
	this.DurationUnit = durationUnit
	this.DurationValue = durationValue
	return &this
}

// NewDurationWithDefaults instantiates a new Duration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDurationWithDefaults() *Duration {
	this := Duration{}
	return &this
}

// GetDurationUnit returns the DurationUnit field value
func (o *Duration) GetDurationUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DurationUnit
}

// GetDurationUnitOk returns a tuple with the DurationUnit field value
// and a boolean to check if the value has been set.
func (o *Duration) GetDurationUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationUnit, true
}

// SetDurationUnit sets field value
func (o *Duration) SetDurationUnit(v string) {
	o.DurationUnit = v
}

// GetDurationValue returns the DurationValue field value
func (o *Duration) GetDurationValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DurationValue
}

// GetDurationValueOk returns a tuple with the DurationValue field value
// and a boolean to check if the value has been set.
func (o *Duration) GetDurationValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationValue, true
}

// SetDurationValue sets field value
func (o *Duration) SetDurationValue(v int32) {
	o.DurationValue = v
}

func (o Duration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["durationUnit"] = o.DurationUnit
	toSerialize["durationValue"] = o.DurationValue
	return toSerialize, nil
}

type NullableDuration struct {
	value *Duration
	isSet bool
}

func (v NullableDuration) Get() *Duration {
	return v.value
}

func (v *NullableDuration) Set(val *Duration) {
	v.value = val
	v.isSet = true
}

func (v NullableDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDuration(val *Duration) *NullableDuration {
	return &NullableDuration{value: val, isSet: true}
}

func (v NullableDuration) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
