package services

import (
	"github.com/bytedance/sonic"
)

// checks if the AssignAppointmentResourcesResponsePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignAppointmentResourcesResponsePayload{}

// AssignAppointmentResourcesResponsePayload The payload for the `assignAppointmentResource` operation.
type AssignAppointmentResourcesResponsePayload struct {
	// A list of warnings returned in the sucessful execution response of an API request.
	Warnings []Warning `json:"warnings,omitempty"`
}

// NewAssignAppointmentResourcesResponsePayload instantiates a new AssignAppointmentResourcesResponsePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignAppointmentResourcesResponsePayload() *AssignAppointmentResourcesResponsePayload {
	this := AssignAppointmentResourcesResponsePayload{}
	return &this
}

// NewAssignAppointmentResourcesResponsePayloadWithDefaults instantiates a new AssignAppointmentResourcesResponsePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignAppointmentResourcesResponsePayloadWithDefaults() *AssignAppointmentResourcesResponsePayload {
	this := AssignAppointmentResourcesResponsePayload{}
	return &this
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *AssignAppointmentResourcesResponsePayload) GetWarnings() []Warning {
	if o == nil || IsNil(o.Warnings) {
		var ret []Warning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignAppointmentResourcesResponsePayload) GetWarningsOk() ([]Warning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *AssignAppointmentResourcesResponsePayload) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Warning and assigns it to the Warnings field.
func (o *AssignAppointmentResourcesResponsePayload) SetWarnings(v []Warning) {
	o.Warnings = v
}

func (o AssignAppointmentResourcesResponsePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableAssignAppointmentResourcesResponsePayload struct {
	value *AssignAppointmentResourcesResponsePayload
	isSet bool
}

func (v NullableAssignAppointmentResourcesResponsePayload) Get() *AssignAppointmentResourcesResponsePayload {
	return v.value
}

func (v *NullableAssignAppointmentResourcesResponsePayload) Set(val *AssignAppointmentResourcesResponsePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignAppointmentResourcesResponsePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignAppointmentResourcesResponsePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignAppointmentResourcesResponsePayload(val *AssignAppointmentResourcesResponsePayload) *NullableAssignAppointmentResourcesResponsePayload {
	return &NullableAssignAppointmentResourcesResponsePayload{value: val, isSet: true}
}

func (v NullableAssignAppointmentResourcesResponsePayload) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAssignAppointmentResourcesResponsePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
