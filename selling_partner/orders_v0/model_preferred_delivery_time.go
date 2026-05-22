package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the PreferredDeliveryTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreferredDeliveryTime{}

// PreferredDeliveryTime The time window when the delivery is preferred.
type PreferredDeliveryTime struct {
	// Business hours when the business is open for deliveries.
	BusinessHours []BusinessHours `json:"BusinessHours,omitempty"`
	// Dates when the business is closed during the next 30 days.
	ExceptionDates []ExceptionDates `json:"ExceptionDates,omitempty"`
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
func (o *PreferredDeliveryTime) GetBusinessHours() []BusinessHours {
	if o == nil || IsNil(o.BusinessHours) {
		var ret []BusinessHours
		return ret
	}
	return o.BusinessHours
}

// GetBusinessHoursOk returns a tuple with the BusinessHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredDeliveryTime) GetBusinessHoursOk() ([]BusinessHours, bool) {
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

// SetBusinessHours gets a reference to the given []BusinessHours and assigns it to the BusinessHours field.
func (o *PreferredDeliveryTime) SetBusinessHours(v []BusinessHours) {
	o.BusinessHours = v
}

// GetExceptionDates returns the ExceptionDates field value if set, zero value otherwise.
func (o *PreferredDeliveryTime) GetExceptionDates() []ExceptionDates {
	if o == nil || IsNil(o.ExceptionDates) {
		var ret []ExceptionDates
		return ret
	}
	return o.ExceptionDates
}

// GetExceptionDatesOk returns a tuple with the ExceptionDates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredDeliveryTime) GetExceptionDatesOk() ([]ExceptionDates, bool) {
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

// SetExceptionDates gets a reference to the given []ExceptionDates and assigns it to the ExceptionDates field.
func (o *PreferredDeliveryTime) SetExceptionDates(v []ExceptionDates) {
	o.ExceptionDates = v
}

func (o PreferredDeliveryTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessHours) {
		toSerialize["BusinessHours"] = o.BusinessHours
	}
	if !IsNil(o.ExceptionDates) {
		toSerialize["ExceptionDates"] = o.ExceptionDates
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
