package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the OpenInterval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenInterval{}

// OpenInterval The time interval for which the business is open.
type OpenInterval struct {
	StartTime *OpenTimeInterval `json:"StartTime,omitempty"`
	EndTime   *OpenTimeInterval `json:"EndTime,omitempty"`
}

// NewOpenInterval instantiates a new OpenInterval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenInterval() *OpenInterval {
	this := OpenInterval{}
	return &this
}

// NewOpenIntervalWithDefaults instantiates a new OpenInterval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenIntervalWithDefaults() *OpenInterval {
	this := OpenInterval{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *OpenInterval) GetStartTime() OpenTimeInterval {
	if o == nil || IsNil(o.StartTime) {
		var ret OpenTimeInterval
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterval) GetStartTimeOk() (*OpenTimeInterval, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *OpenInterval) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given OpenTimeInterval and assigns it to the StartTime field.
func (o *OpenInterval) SetStartTime(v OpenTimeInterval) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *OpenInterval) GetEndTime() OpenTimeInterval {
	if o == nil || IsNil(o.EndTime) {
		var ret OpenTimeInterval
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterval) GetEndTimeOk() (*OpenTimeInterval, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *OpenInterval) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given OpenTimeInterval and assigns it to the EndTime field.
func (o *OpenInterval) SetEndTime(v OpenTimeInterval) {
	o.EndTime = &v
}

func (o OpenInterval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTime) {
		toSerialize["StartTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["EndTime"] = o.EndTime
	}
	return toSerialize, nil
}

type NullableOpenInterval struct {
	value *OpenInterval
	isSet bool
}

func (v NullableOpenInterval) Get() *OpenInterval {
	return v.value
}

func (v *NullableOpenInterval) Set(val *OpenInterval) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenInterval) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenInterval(val *OpenInterval) *NullableOpenInterval {
	return &NullableOpenInterval{value: val, isSet: true}
}

func (v NullableOpenInterval) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOpenInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
