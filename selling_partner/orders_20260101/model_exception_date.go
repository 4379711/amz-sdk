package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ExceptionDate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExceptionDate{}

// ExceptionDate Special dates when normal business hours are modified or suspended, requiring different delivery scheduling.
type ExceptionDate struct {
	// Specific calendar date when normal operating hours do not apply. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format at day granularity.
	ExceptionDate *string `json:"exceptionDate,omitempty"`
	// Operational status of the business on the specified exception date.
	ExceptionDateType *string `json:"exceptionDateType,omitempty"`
	// Alternative operating hours that apply specifically to this exception date.
	TimeWindows []TimeWindow `json:"timeWindows,omitempty"`
}

// NewExceptionDate instantiates a new ExceptionDate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExceptionDate() *ExceptionDate {
	this := ExceptionDate{}
	return &this
}

// NewExceptionDateWithDefaults instantiates a new ExceptionDate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExceptionDateWithDefaults() *ExceptionDate {
	this := ExceptionDate{}
	return &this
}

// GetExceptionDate returns the ExceptionDate field value if set, zero value otherwise.
func (o *ExceptionDate) GetExceptionDate() string {
	if o == nil || IsNil(o.ExceptionDate) {
		var ret string
		return ret
	}
	return *o.ExceptionDate
}

// GetExceptionDateOk returns a tuple with the ExceptionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionDate) GetExceptionDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExceptionDate) {
		return nil, false
	}
	return o.ExceptionDate, true
}

// HasExceptionDate returns a boolean if a field has been set.
func (o *ExceptionDate) HasExceptionDate() bool {
	if o != nil && !IsNil(o.ExceptionDate) {
		return true
	}

	return false
}

// SetExceptionDate gets a reference to the given string and assigns it to the ExceptionDate field.
func (o *ExceptionDate) SetExceptionDate(v string) {
	o.ExceptionDate = &v
}

// GetExceptionDateType returns the ExceptionDateType field value if set, zero value otherwise.
func (o *ExceptionDate) GetExceptionDateType() string {
	if o == nil || IsNil(o.ExceptionDateType) {
		var ret string
		return ret
	}
	return *o.ExceptionDateType
}

// GetExceptionDateTypeOk returns a tuple with the ExceptionDateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionDate) GetExceptionDateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExceptionDateType) {
		return nil, false
	}
	return o.ExceptionDateType, true
}

// HasExceptionDateType returns a boolean if a field has been set.
func (o *ExceptionDate) HasExceptionDateType() bool {
	if o != nil && !IsNil(o.ExceptionDateType) {
		return true
	}

	return false
}

// SetExceptionDateType gets a reference to the given string and assigns it to the ExceptionDateType field.
func (o *ExceptionDate) SetExceptionDateType(v string) {
	o.ExceptionDateType = &v
}

// GetTimeWindows returns the TimeWindows field value if set, zero value otherwise.
func (o *ExceptionDate) GetTimeWindows() []TimeWindow {
	if o == nil || IsNil(o.TimeWindows) {
		var ret []TimeWindow
		return ret
	}
	return o.TimeWindows
}

// GetTimeWindowsOk returns a tuple with the TimeWindows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionDate) GetTimeWindowsOk() ([]TimeWindow, bool) {
	if o == nil || IsNil(o.TimeWindows) {
		return nil, false
	}
	return o.TimeWindows, true
}

// HasTimeWindows returns a boolean if a field has been set.
func (o *ExceptionDate) HasTimeWindows() bool {
	if o != nil && !IsNil(o.TimeWindows) {
		return true
	}

	return false
}

// SetTimeWindows gets a reference to the given []TimeWindow and assigns it to the TimeWindows field.
func (o *ExceptionDate) SetTimeWindows(v []TimeWindow) {
	o.TimeWindows = v
}

func (o ExceptionDate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExceptionDate) {
		toSerialize["exceptionDate"] = o.ExceptionDate
	}
	if !IsNil(o.ExceptionDateType) {
		toSerialize["exceptionDateType"] = o.ExceptionDateType
	}
	if !IsNil(o.TimeWindows) {
		toSerialize["timeWindows"] = o.TimeWindows
	}
	return toSerialize, nil
}

type NullableExceptionDate struct {
	value *ExceptionDate
	isSet bool
}

func (v NullableExceptionDate) Get() *ExceptionDate {
	return v.value
}

func (v *NullableExceptionDate) Set(val *ExceptionDate) {
	v.value = val
	v.isSet = true
}

func (v NullableExceptionDate) IsSet() bool {
	return v.isSet
}

func (v *NullableExceptionDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExceptionDate(val *ExceptionDate) *NullableExceptionDate {
	return &NullableExceptionDate{value: val, isSet: true}
}

func (v NullableExceptionDate) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableExceptionDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
