package services

import (
	"github.com/bytedance/sonic"
)

// checks if the AppointmentResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppointmentResource{}

// AppointmentResource The resource that performs or performed appointment fulfillment.
type AppointmentResource struct {
	// The resource identifier.
	ResourceId *string `json:"resourceId,omitempty"`
}

// NewAppointmentResource instantiates a new AppointmentResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppointmentResource() *AppointmentResource {
	this := AppointmentResource{}
	return &this
}

// NewAppointmentResourceWithDefaults instantiates a new AppointmentResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppointmentResourceWithDefaults() *AppointmentResource {
	this := AppointmentResource{}
	return &this
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *AppointmentResource) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResource) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *AppointmentResource) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *AppointmentResource) SetResourceId(v string) {
	o.ResourceId = &v
}

func (o AppointmentResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceId) {
		toSerialize["resourceId"] = o.ResourceId
	}
	return toSerialize, nil
}

type NullableAppointmentResource struct {
	value *AppointmentResource
	isSet bool
}

func (v NullableAppointmentResource) Get() *AppointmentResource {
	return v.value
}

func (v *NullableAppointmentResource) Set(val *AppointmentResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAppointmentResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAppointmentResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppointmentResource(val *AppointmentResource) *NullableAppointmentResource {
	return &NullableAppointmentResource{value: val, isSet: true}
}

func (v NullableAppointmentResource) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAppointmentResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
