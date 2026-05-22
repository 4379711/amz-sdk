package supply_sources_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the OperatingHour type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperatingHour{}

// OperatingHour The operating hour schema
type OperatingHour struct {
	// The opening time, ISO 8601 formatted timestamp without date, HH:mm.
	StartTime *string `json:"startTime,omitempty"`
	// The closing time, ISO 8601 formatted timestamp without date, HH:mm.
	EndTime *string `json:"endTime,omitempty"`
}

// NewOperatingHour instantiates a new OperatingHour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatingHour() *OperatingHour {
	this := OperatingHour{}
	return &this
}

// NewOperatingHourWithDefaults instantiates a new OperatingHour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatingHourWithDefaults() *OperatingHour {
	this := OperatingHour{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *OperatingHour) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingHour) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *OperatingHour) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *OperatingHour) SetStartTime(v string) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *OperatingHour) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingHour) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *OperatingHour) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *OperatingHour) SetEndTime(v string) {
	o.EndTime = &v
}

func (o OperatingHour) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	return toSerialize, nil
}

type NullableOperatingHour struct {
	value *OperatingHour
	isSet bool
}

func (v NullableOperatingHour) Get() *OperatingHour {
	return v.value
}

func (v *NullableOperatingHour) Set(val *OperatingHour) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatingHour) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatingHour) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatingHour(val *OperatingHour) *NullableOperatingHour {
	return &NullableOperatingHour{value: val, isSet: true}
}

func (v NullableOperatingHour) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOperatingHour) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
