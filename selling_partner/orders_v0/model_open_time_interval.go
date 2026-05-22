package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the OpenTimeInterval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenTimeInterval{}

// OpenTimeInterval The time when the business opens or closes.
type OpenTimeInterval struct {
	// The hour when the business opens or closes.
	Hour *int32 `json:"Hour,omitempty"`
	// The minute when the business opens or closes.
	Minute *int32 `json:"Minute,omitempty"`
}

// NewOpenTimeInterval instantiates a new OpenTimeInterval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenTimeInterval() *OpenTimeInterval {
	this := OpenTimeInterval{}
	return &this
}

// NewOpenTimeIntervalWithDefaults instantiates a new OpenTimeInterval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenTimeIntervalWithDefaults() *OpenTimeInterval {
	this := OpenTimeInterval{}
	return &this
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *OpenTimeInterval) GetHour() int32 {
	if o == nil || IsNil(o.Hour) {
		var ret int32
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenTimeInterval) GetHourOk() (*int32, bool) {
	if o == nil || IsNil(o.Hour) {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *OpenTimeInterval) HasHour() bool {
	if o != nil && !IsNil(o.Hour) {
		return true
	}

	return false
}

// SetHour gets a reference to the given int32 and assigns it to the Hour field.
func (o *OpenTimeInterval) SetHour(v int32) {
	o.Hour = &v
}

// GetMinute returns the Minute field value if set, zero value otherwise.
func (o *OpenTimeInterval) GetMinute() int32 {
	if o == nil || IsNil(o.Minute) {
		var ret int32
		return ret
	}
	return *o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenTimeInterval) GetMinuteOk() (*int32, bool) {
	if o == nil || IsNil(o.Minute) {
		return nil, false
	}
	return o.Minute, true
}

// HasMinute returns a boolean if a field has been set.
func (o *OpenTimeInterval) HasMinute() bool {
	if o != nil && !IsNil(o.Minute) {
		return true
	}

	return false
}

// SetMinute gets a reference to the given int32 and assigns it to the Minute field.
func (o *OpenTimeInterval) SetMinute(v int32) {
	o.Minute = &v
}

func (o OpenTimeInterval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hour) {
		toSerialize["Hour"] = o.Hour
	}
	if !IsNil(o.Minute) {
		toSerialize["Minute"] = o.Minute
	}
	return toSerialize, nil
}

type NullableOpenTimeInterval struct {
	value *OpenTimeInterval
	isSet bool
}

func (v NullableOpenTimeInterval) Get() *OpenTimeInterval {
	return v.value
}

func (v *NullableOpenTimeInterval) Set(val *OpenTimeInterval) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenTimeInterval) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenTimeInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenTimeInterval(val *OpenTimeInterval) *NullableOpenTimeInterval {
	return &NullableOpenTimeInterval{value: val, isSet: true}
}

func (v NullableOpenTimeInterval) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOpenTimeInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
