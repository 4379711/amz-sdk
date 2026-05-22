package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ExceptionDates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExceptionDates{}

// ExceptionDates Dates when the business is closed or open with a different time window.
type ExceptionDates struct {
	// Date when the business is closed, in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date format.
	ExceptionDate *string `json:"ExceptionDate,omitempty"`
	// Boolean indicating if the business is closed or open on that date.
	IsOpen *bool `json:"IsOpen,omitempty"`
	// Time window during the day when the business is open.
	OpenIntervals []OpenInterval `json:"OpenIntervals,omitempty"`
}

// NewExceptionDates instantiates a new ExceptionDates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExceptionDates() *ExceptionDates {
	this := ExceptionDates{}
	return &this
}

// NewExceptionDatesWithDefaults instantiates a new ExceptionDates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExceptionDatesWithDefaults() *ExceptionDates {
	this := ExceptionDates{}
	return &this
}

// GetExceptionDate returns the ExceptionDate field value if set, zero value otherwise.
func (o *ExceptionDates) GetExceptionDate() string {
	if o == nil || IsNil(o.ExceptionDate) {
		var ret string
		return ret
	}
	return *o.ExceptionDate
}

// GetExceptionDateOk returns a tuple with the ExceptionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionDates) GetExceptionDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExceptionDate) {
		return nil, false
	}
	return o.ExceptionDate, true
}

// HasExceptionDate returns a boolean if a field has been set.
func (o *ExceptionDates) HasExceptionDate() bool {
	if o != nil && !IsNil(o.ExceptionDate) {
		return true
	}

	return false
}

// SetExceptionDate gets a reference to the given string and assigns it to the ExceptionDate field.
func (o *ExceptionDates) SetExceptionDate(v string) {
	o.ExceptionDate = &v
}

// GetIsOpen returns the IsOpen field value if set, zero value otherwise.
func (o *ExceptionDates) GetIsOpen() bool {
	if o == nil || IsNil(o.IsOpen) {
		var ret bool
		return ret
	}
	return *o.IsOpen
}

// GetIsOpenOk returns a tuple with the IsOpen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionDates) GetIsOpenOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOpen) {
		return nil, false
	}
	return o.IsOpen, true
}

// HasIsOpen returns a boolean if a field has been set.
func (o *ExceptionDates) HasIsOpen() bool {
	if o != nil && !IsNil(o.IsOpen) {
		return true
	}

	return false
}

// SetIsOpen gets a reference to the given bool and assigns it to the IsOpen field.
func (o *ExceptionDates) SetIsOpen(v bool) {
	o.IsOpen = &v
}

// GetOpenIntervals returns the OpenIntervals field value if set, zero value otherwise.
func (o *ExceptionDates) GetOpenIntervals() []OpenInterval {
	if o == nil || IsNil(o.OpenIntervals) {
		var ret []OpenInterval
		return ret
	}
	return o.OpenIntervals
}

// GetOpenIntervalsOk returns a tuple with the OpenIntervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionDates) GetOpenIntervalsOk() ([]OpenInterval, bool) {
	if o == nil || IsNil(o.OpenIntervals) {
		return nil, false
	}
	return o.OpenIntervals, true
}

// HasOpenIntervals returns a boolean if a field has been set.
func (o *ExceptionDates) HasOpenIntervals() bool {
	if o != nil && !IsNil(o.OpenIntervals) {
		return true
	}

	return false
}

// SetOpenIntervals gets a reference to the given []OpenInterval and assigns it to the OpenIntervals field.
func (o *ExceptionDates) SetOpenIntervals(v []OpenInterval) {
	o.OpenIntervals = v
}

func (o ExceptionDates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExceptionDate) {
		toSerialize["ExceptionDate"] = o.ExceptionDate
	}
	if !IsNil(o.IsOpen) {
		toSerialize["IsOpen"] = o.IsOpen
	}
	if !IsNil(o.OpenIntervals) {
		toSerialize["OpenIntervals"] = o.OpenIntervals
	}
	return toSerialize, nil
}

type NullableExceptionDates struct {
	value *ExceptionDates
	isSet bool
}

func (v NullableExceptionDates) Get() *ExceptionDates {
	return v.value
}

func (v *NullableExceptionDates) Set(val *ExceptionDates) {
	v.value = val
	v.isSet = true
}

func (v NullableExceptionDates) IsSet() bool {
	return v.isSet
}

func (v *NullableExceptionDates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExceptionDates(val *ExceptionDates) *NullableExceptionDates {
	return &NullableExceptionDates{value: val, isSet: true}
}

func (v NullableExceptionDates) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableExceptionDates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
