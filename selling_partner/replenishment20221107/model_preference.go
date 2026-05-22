package replenishment20221107

import (
	"github.com/bytedance/sonic"
)

// checks if the Preference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Preference{}

// Preference Offer preferences that you can include in the result filter criteria.
type Preference struct {
	// Filters the results to only include offers with the auto-enrollment preference specified.
	AutoEnrollment []AutoEnrollmentPreference `json:"autoEnrollment,omitempty"`
}

// NewPreference instantiates a new Preference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreference() *Preference {
	this := Preference{}
	return &this
}

// NewPreferenceWithDefaults instantiates a new Preference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferenceWithDefaults() *Preference {
	this := Preference{}
	return &this
}

// GetAutoEnrollment returns the AutoEnrollment field value if set, zero value otherwise.
func (o *Preference) GetAutoEnrollment() []AutoEnrollmentPreference {
	if o == nil || IsNil(o.AutoEnrollment) {
		var ret []AutoEnrollmentPreference
		return ret
	}
	return o.AutoEnrollment
}

// GetAutoEnrollmentOk returns a tuple with the AutoEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preference) GetAutoEnrollmentOk() ([]AutoEnrollmentPreference, bool) {
	if o == nil || IsNil(o.AutoEnrollment) {
		return nil, false
	}
	return o.AutoEnrollment, true
}

// HasAutoEnrollment returns a boolean if a field has been set.
func (o *Preference) HasAutoEnrollment() bool {
	if o != nil && !IsNil(o.AutoEnrollment) {
		return true
	}

	return false
}

// SetAutoEnrollment gets a reference to the given []AutoEnrollmentPreference and assigns it to the AutoEnrollment field.
func (o *Preference) SetAutoEnrollment(v []AutoEnrollmentPreference) {
	o.AutoEnrollment = v
}

func (o Preference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoEnrollment) {
		toSerialize["autoEnrollment"] = o.AutoEnrollment
	}
	return toSerialize, nil
}

type NullablePreference struct {
	value *Preference
	isSet bool
}

func (v NullablePreference) Get() *Preference {
	return v.value
}

func (v *NullablePreference) Set(val *Preference) {
	v.value = val
	v.isSet = true
}

func (v NullablePreference) IsSet() bool {
	return v.isSet
}

func (v *NullablePreference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreference(val *Preference) *NullablePreference {
	return &NullablePreference{value: val, isSet: true}
}

func (v NullablePreference) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePreference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
