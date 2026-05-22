package replenishment20221107

import (
	"github.com/bytedance/sonic"
)

// checks if the OfferProgramConfigurationPreferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferProgramConfigurationPreferences{}

// OfferProgramConfigurationPreferences An object which contains the preferences applied to the offer.
type OfferProgramConfigurationPreferences struct {
	AutoEnrollment *AutoEnrollmentPreference `json:"autoEnrollment,omitempty"`
}

// NewOfferProgramConfigurationPreferences instantiates a new OfferProgramConfigurationPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferProgramConfigurationPreferences() *OfferProgramConfigurationPreferences {
	this := OfferProgramConfigurationPreferences{}
	return &this
}

// NewOfferProgramConfigurationPreferencesWithDefaults instantiates a new OfferProgramConfigurationPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferProgramConfigurationPreferencesWithDefaults() *OfferProgramConfigurationPreferences {
	this := OfferProgramConfigurationPreferences{}
	return &this
}

// GetAutoEnrollment returns the AutoEnrollment field value if set, zero value otherwise.
func (o *OfferProgramConfigurationPreferences) GetAutoEnrollment() AutoEnrollmentPreference {
	if o == nil || IsNil(o.AutoEnrollment) {
		var ret AutoEnrollmentPreference
		return ret
	}
	return *o.AutoEnrollment
}

// GetAutoEnrollmentOk returns a tuple with the AutoEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferProgramConfigurationPreferences) GetAutoEnrollmentOk() (*AutoEnrollmentPreference, bool) {
	if o == nil || IsNil(o.AutoEnrollment) {
		return nil, false
	}
	return o.AutoEnrollment, true
}

// HasAutoEnrollment returns a boolean if a field has been set.
func (o *OfferProgramConfigurationPreferences) HasAutoEnrollment() bool {
	if o != nil && !IsNil(o.AutoEnrollment) {
		return true
	}

	return false
}

// SetAutoEnrollment gets a reference to the given AutoEnrollmentPreference and assigns it to the AutoEnrollment field.
func (o *OfferProgramConfigurationPreferences) SetAutoEnrollment(v AutoEnrollmentPreference) {
	o.AutoEnrollment = &v
}

func (o OfferProgramConfigurationPreferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoEnrollment) {
		toSerialize["autoEnrollment"] = o.AutoEnrollment
	}
	return toSerialize, nil
}

type NullableOfferProgramConfigurationPreferences struct {
	value *OfferProgramConfigurationPreferences
	isSet bool
}

func (v NullableOfferProgramConfigurationPreferences) Get() *OfferProgramConfigurationPreferences {
	return v.value
}

func (v *NullableOfferProgramConfigurationPreferences) Set(val *OfferProgramConfigurationPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferProgramConfigurationPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferProgramConfigurationPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferProgramConfigurationPreferences(val *OfferProgramConfigurationPreferences) *NullableOfferProgramConfigurationPreferences {
	return &NullableOfferProgramConfigurationPreferences{value: val, isSet: true}
}

func (v NullableOfferProgramConfigurationPreferences) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOfferProgramConfigurationPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
