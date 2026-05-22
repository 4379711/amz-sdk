package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the DateRangeTypeRuleDuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DateRangeTypeRuleDuration{}

// DateRangeTypeRuleDuration Object representing date range type rule duration.
type DateRangeTypeRuleDuration struct {
	// The end date of the budget rule in YYYYMMDD format. The end date is inclusive. Required to be equal or greater than `startDate`.
	EndDate *string `json:"endDate,omitempty"`
	// The start date of the budget rule in YYYYMMDD format. The start date is inclusive. Required to be greater than or equal to current date.
	StartDate string `json:"startDate"`
}

type _DateRangeTypeRuleDuration DateRangeTypeRuleDuration

// NewDateRangeTypeRuleDuration instantiates a new DateRangeTypeRuleDuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateRangeTypeRuleDuration(startDate string) *DateRangeTypeRuleDuration {
	this := DateRangeTypeRuleDuration{}
	this.StartDate = startDate
	return &this
}

// NewDateRangeTypeRuleDurationWithDefaults instantiates a new DateRangeTypeRuleDuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateRangeTypeRuleDurationWithDefaults() *DateRangeTypeRuleDuration {
	this := DateRangeTypeRuleDuration{}
	return &this
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *DateRangeTypeRuleDuration) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateRangeTypeRuleDuration) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *DateRangeTypeRuleDuration) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *DateRangeTypeRuleDuration) SetEndDate(v string) {
	o.EndDate = &v
}

// GetStartDate returns the StartDate field value
func (o *DateRangeTypeRuleDuration) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *DateRangeTypeRuleDuration) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *DateRangeTypeRuleDuration) SetStartDate(v string) {
	o.StartDate = v
}

func (o DateRangeTypeRuleDuration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	toSerialize["startDate"] = o.StartDate
	return toSerialize, nil
}

type NullableDateRangeTypeRuleDuration struct {
	value *DateRangeTypeRuleDuration
	isSet bool
}

func (v NullableDateRangeTypeRuleDuration) Get() *DateRangeTypeRuleDuration {
	return v.value
}

func (v *NullableDateRangeTypeRuleDuration) Set(val *DateRangeTypeRuleDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableDateRangeTypeRuleDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableDateRangeTypeRuleDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateRangeTypeRuleDuration(val *DateRangeTypeRuleDuration) *NullableDateRangeTypeRuleDuration {
	return &NullableDateRangeTypeRuleDuration{value: val, isSet: true}
}

func (v NullableDateRangeTypeRuleDuration) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDateRangeTypeRuleDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
