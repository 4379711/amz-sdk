package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the ExceptionOperatingHours type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExceptionOperatingHours{}

// ExceptionOperatingHours Defines exceptions to standard operating hours for certain date ranges.
type ExceptionOperatingHours struct {
	DateRange      *DateRange      `json:"dateRange,omitempty"`
	OperatingHours *OperatingHours `json:"operatingHours,omitempty"`
}

// NewExceptionOperatingHours instantiates a new ExceptionOperatingHours object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExceptionOperatingHours() *ExceptionOperatingHours {
	this := ExceptionOperatingHours{}
	return &this
}

// NewExceptionOperatingHoursWithDefaults instantiates a new ExceptionOperatingHours object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExceptionOperatingHoursWithDefaults() *ExceptionOperatingHours {
	this := ExceptionOperatingHours{}
	return &this
}

// GetDateRange returns the DateRange field value if set, zero value otherwise.
func (o *ExceptionOperatingHours) GetDateRange() DateRange {
	if o == nil || IsNil(o.DateRange) {
		var ret DateRange
		return ret
	}
	return *o.DateRange
}

// GetDateRangeOk returns a tuple with the DateRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionOperatingHours) GetDateRangeOk() (*DateRange, bool) {
	if o == nil || IsNil(o.DateRange) {
		return nil, false
	}
	return o.DateRange, true
}

// HasDateRange returns a boolean if a field has been set.
func (o *ExceptionOperatingHours) HasDateRange() bool {
	if o != nil && !IsNil(o.DateRange) {
		return true
	}

	return false
}

// SetDateRange gets a reference to the given DateRange and assigns it to the DateRange field.
func (o *ExceptionOperatingHours) SetDateRange(v DateRange) {
	o.DateRange = &v
}

// GetOperatingHours returns the OperatingHours field value if set, zero value otherwise.
func (o *ExceptionOperatingHours) GetOperatingHours() OperatingHours {
	if o == nil || IsNil(o.OperatingHours) {
		var ret OperatingHours
		return ret
	}
	return *o.OperatingHours
}

// GetOperatingHoursOk returns a tuple with the OperatingHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionOperatingHours) GetOperatingHoursOk() (*OperatingHours, bool) {
	if o == nil || IsNil(o.OperatingHours) {
		return nil, false
	}
	return o.OperatingHours, true
}

// HasOperatingHours returns a boolean if a field has been set.
func (o *ExceptionOperatingHours) HasOperatingHours() bool {
	if o != nil && !IsNil(o.OperatingHours) {
		return true
	}

	return false
}

// SetOperatingHours gets a reference to the given OperatingHours and assigns it to the OperatingHours field.
func (o *ExceptionOperatingHours) SetOperatingHours(v OperatingHours) {
	o.OperatingHours = &v
}

func (o ExceptionOperatingHours) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateRange) {
		toSerialize["dateRange"] = o.DateRange
	}
	if !IsNil(o.OperatingHours) {
		toSerialize["operatingHours"] = o.OperatingHours
	}
	return toSerialize, nil
}

type NullableExceptionOperatingHours struct {
	value *ExceptionOperatingHours
	isSet bool
}

func (v NullableExceptionOperatingHours) Get() *ExceptionOperatingHours {
	return v.value
}

func (v *NullableExceptionOperatingHours) Set(val *ExceptionOperatingHours) {
	v.value = val
	v.isSet = true
}

func (v NullableExceptionOperatingHours) IsSet() bool {
	return v.isSet
}

func (v *NullableExceptionOperatingHours) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExceptionOperatingHours(val *ExceptionOperatingHours) *NullableExceptionOperatingHours {
	return &NullableExceptionOperatingHours{value: val, isSet: true}
}

func (v NullableExceptionOperatingHours) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableExceptionOperatingHours) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
