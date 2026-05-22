package supply_sources_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the OperationalConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationalConfiguration{}

// OperationalConfiguration The operational configuration of `supplySources`.
type OperationalConfiguration struct {
	ContactDetails      *ContactDetails      `json:"contactDetails,omitempty"`
	ThroughputConfig    *ThroughputConfig    `json:"throughputConfig,omitempty"`
	OperatingHoursByDay *OperatingHoursByDay `json:"operatingHoursByDay,omitempty"`
	HandlingTime        *Duration            `json:"handlingTime,omitempty"`
}

// NewOperationalConfiguration instantiates a new OperationalConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationalConfiguration() *OperationalConfiguration {
	this := OperationalConfiguration{}
	return &this
}

// NewOperationalConfigurationWithDefaults instantiates a new OperationalConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationalConfigurationWithDefaults() *OperationalConfiguration {
	this := OperationalConfiguration{}
	return &this
}

// GetContactDetails returns the ContactDetails field value if set, zero value otherwise.
func (o *OperationalConfiguration) GetContactDetails() ContactDetails {
	if o == nil || IsNil(o.ContactDetails) {
		var ret ContactDetails
		return ret
	}
	return *o.ContactDetails
}

// GetContactDetailsOk returns a tuple with the ContactDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationalConfiguration) GetContactDetailsOk() (*ContactDetails, bool) {
	if o == nil || IsNil(o.ContactDetails) {
		return nil, false
	}
	return o.ContactDetails, true
}

// HasContactDetails returns a boolean if a field has been set.
func (o *OperationalConfiguration) HasContactDetails() bool {
	if o != nil && !IsNil(o.ContactDetails) {
		return true
	}

	return false
}

// SetContactDetails gets a reference to the given ContactDetails and assigns it to the ContactDetails field.
func (o *OperationalConfiguration) SetContactDetails(v ContactDetails) {
	o.ContactDetails = &v
}

// GetThroughputConfig returns the ThroughputConfig field value if set, zero value otherwise.
func (o *OperationalConfiguration) GetThroughputConfig() ThroughputConfig {
	if o == nil || IsNil(o.ThroughputConfig) {
		var ret ThroughputConfig
		return ret
	}
	return *o.ThroughputConfig
}

// GetThroughputConfigOk returns a tuple with the ThroughputConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationalConfiguration) GetThroughputConfigOk() (*ThroughputConfig, bool) {
	if o == nil || IsNil(o.ThroughputConfig) {
		return nil, false
	}
	return o.ThroughputConfig, true
}

// HasThroughputConfig returns a boolean if a field has been set.
func (o *OperationalConfiguration) HasThroughputConfig() bool {
	if o != nil && !IsNil(o.ThroughputConfig) {
		return true
	}

	return false
}

// SetThroughputConfig gets a reference to the given ThroughputConfig and assigns it to the ThroughputConfig field.
func (o *OperationalConfiguration) SetThroughputConfig(v ThroughputConfig) {
	o.ThroughputConfig = &v
}

// GetOperatingHoursByDay returns the OperatingHoursByDay field value if set, zero value otherwise.
func (o *OperationalConfiguration) GetOperatingHoursByDay() OperatingHoursByDay {
	if o == nil || IsNil(o.OperatingHoursByDay) {
		var ret OperatingHoursByDay
		return ret
	}
	return *o.OperatingHoursByDay
}

// GetOperatingHoursByDayOk returns a tuple with the OperatingHoursByDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationalConfiguration) GetOperatingHoursByDayOk() (*OperatingHoursByDay, bool) {
	if o == nil || IsNil(o.OperatingHoursByDay) {
		return nil, false
	}
	return o.OperatingHoursByDay, true
}

// HasOperatingHoursByDay returns a boolean if a field has been set.
func (o *OperationalConfiguration) HasOperatingHoursByDay() bool {
	if o != nil && !IsNil(o.OperatingHoursByDay) {
		return true
	}

	return false
}

// SetOperatingHoursByDay gets a reference to the given OperatingHoursByDay and assigns it to the OperatingHoursByDay field.
func (o *OperationalConfiguration) SetOperatingHoursByDay(v OperatingHoursByDay) {
	o.OperatingHoursByDay = &v
}

// GetHandlingTime returns the HandlingTime field value if set, zero value otherwise.
func (o *OperationalConfiguration) GetHandlingTime() Duration {
	if o == nil || IsNil(o.HandlingTime) {
		var ret Duration
		return ret
	}
	return *o.HandlingTime
}

// GetHandlingTimeOk returns a tuple with the HandlingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationalConfiguration) GetHandlingTimeOk() (*Duration, bool) {
	if o == nil || IsNil(o.HandlingTime) {
		return nil, false
	}
	return o.HandlingTime, true
}

// HasHandlingTime returns a boolean if a field has been set.
func (o *OperationalConfiguration) HasHandlingTime() bool {
	if o != nil && !IsNil(o.HandlingTime) {
		return true
	}

	return false
}

// SetHandlingTime gets a reference to the given Duration and assigns it to the HandlingTime field.
func (o *OperationalConfiguration) SetHandlingTime(v Duration) {
	o.HandlingTime = &v
}

func (o OperationalConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContactDetails) {
		toSerialize["contactDetails"] = o.ContactDetails
	}
	if !IsNil(o.ThroughputConfig) {
		toSerialize["throughputConfig"] = o.ThroughputConfig
	}
	if !IsNil(o.OperatingHoursByDay) {
		toSerialize["operatingHoursByDay"] = o.OperatingHoursByDay
	}
	if !IsNil(o.HandlingTime) {
		toSerialize["handlingTime"] = o.HandlingTime
	}
	return toSerialize, nil
}

type NullableOperationalConfiguration struct {
	value *OperationalConfiguration
	isSet bool
}

func (v NullableOperationalConfiguration) Get() *OperationalConfiguration {
	return v.value
}

func (v *NullableOperationalConfiguration) Set(val *OperationalConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationalConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationalConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationalConfiguration(val *OperationalConfiguration) *NullableOperationalConfiguration {
	return &NullableOperationalConfiguration{value: val, isSet: true}
}

func (v NullableOperationalConfiguration) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOperationalConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
