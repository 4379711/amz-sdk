package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the DateRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DateRange{}

// DateRange Date Range for query the results.
type DateRange struct {
	// Start Date for query .
	StartDate *string `json:"startDate,omitempty"`
	// end date for query.
	EndDate *string `json:"endDate,omitempty"`
}

// NewDateRange instantiates a new DateRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateRange() *DateRange {
	this := DateRange{}
	return &this
}

// NewDateRangeWithDefaults instantiates a new DateRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateRangeWithDefaults() *DateRange {
	this := DateRange{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *DateRange) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateRange) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *DateRange) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *DateRange) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *DateRange) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateRange) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *DateRange) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *DateRange) SetEndDate(v string) {
	o.EndDate = &v
}

func (o DateRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableDateRange struct {
	value *DateRange
	isSet bool
}

func (v NullableDateRange) Get() *DateRange {
	return v.value
}

func (v *NullableDateRange) Set(val *DateRange) {
	v.value = val
	v.isSet = true
}

func (v NullableDateRange) IsSet() bool {
	return v.isSet
}

func (v *NullableDateRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateRange(val *DateRange) *NullableDateRange {
	return &NullableDateRange{value: val, isSet: true}
}

func (v NullableDateRange) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDateRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
