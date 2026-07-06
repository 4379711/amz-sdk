package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the PreferredDeliveryTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreferredDeliveryTime{}

// PreferredDeliveryTime Customer-specified time preferences for when deliveries should be attempted at the destination address.
type PreferredDeliveryTime struct {
	// Business hours when the business is open for deliveries.
	BusinessHours []BusinessHour `json:"businessHours,omitempty"`
	// Specific dates within the next 30 days when normal business hours do not apply.
	ExceptionDates []ExceptionDate `json:"exceptionDates,omitempty"`
}

// NewPreferredDeliveryTime instantiates a new PreferredDeliveryTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreferredDeliveryTime() *PreferredDeliveryTime {
	this := PreferredDeliveryTime{}
	return &this
}

// NewPreferredDeliveryTimeWithDefaults instantiates a new PreferredDeliveryTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferredDeliveryTimeWithDefaults() *PreferredDeliveryTime {
	this := PreferredDeliveryTime{}
	return &this
}

// GetBusinessHours returns the BusinessHours field value if set, zero value otherwise.
func (o *PreferredDeliveryTime) GetBusinessHours() []BusinessHour {
	if o == nil || IsNil(o.BusinessHours) {
		var ret []BusinessHour
		return ret
	}
	return o.BusinessHours
}

// GetBusinessHoursOk returns a tuple with the BusinessHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredDeliveryTime) GetBusinessHoursOk() ([]BusinessHour, bool) {
	if o == nil || IsNil(o.BusinessHours) {
		return nil, false
	}
	return o.BusinessHours, true
}

// HasBusinessHours returns a boolean if a field has been set.
func (o *PreferredDeliveryTime) HasBusinessHours() bool {
	if o != nil && !IsNil(o.BusinessHours) {
		return true
	}

	return false
}

// SetBusinessHours gets a reference to the given []BusinessHour and assigns it to the BusinessHours field.
func (o *PreferredDeliveryTime) SetBusinessHours(v []BusinessHour) {
	o.BusinessHours = v
}

// GetExceptionDates returns the ExceptionDates field value if set, zero value otherwise.
func (o *PreferredDeliveryTime) GetExceptionDates() []ExceptionDate {
	if o == nil || IsNil(o.ExceptionDates) {
		var ret []ExceptionDate
		return ret
	}
	return o.ExceptionDates
}

// GetExceptionDatesOk returns a tuple with the ExceptionDates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredDeliveryTime) GetExceptionDatesOk() ([]ExceptionDate, bool) {
	if o == nil || IsNil(o.ExceptionDates) {
		return nil, false
	}
	return o.ExceptionDates, true
}

// HasExceptionDates returns a boolean if a field has been set.
func (o *PreferredDeliveryTime) HasExceptionDates() bool {
	if o != nil && !IsNil(o.ExceptionDates) {
		return true
	}

	return false
}

// SetExceptionDates gets a reference to the given []ExceptionDate and assigns it to the ExceptionDates field.
func (o *PreferredDeliveryTime) SetExceptionDates(v []ExceptionDate) {
	o.ExceptionDates = v
}

func (o PreferredDeliveryTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessHours) {
		toSerialize["businessHours"] = o.BusinessHours
	}
	if !IsNil(o.ExceptionDates) {
		toSerialize["exceptionDates"] = o.ExceptionDates
	}
	return toSerialize, nil
}

type NullablePreferredDeliveryTime struct {
	value *PreferredDeliveryTime
	isSet bool
}

func (v NullablePreferredDeliveryTime) Get() *PreferredDeliveryTime {
	return v.value
}

func (v *NullablePreferredDeliveryTime) Set(val *PreferredDeliveryTime) {
	v.value = val
	v.isSet = true
}

func (v NullablePreferredDeliveryTime) IsSet() bool {
	return v.isSet
}

func (v *NullablePreferredDeliveryTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreferredDeliveryTime(val *PreferredDeliveryTime) *NullablePreferredDeliveryTime {
	return &NullablePreferredDeliveryTime{value: val, isSet: true}
}

func (v NullablePreferredDeliveryTime) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePreferredDeliveryTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
