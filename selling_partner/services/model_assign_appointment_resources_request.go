package services

import (
	"github.com/bytedance/sonic"
)

// checks if the AssignAppointmentResourcesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignAppointmentResourcesRequest{}

// AssignAppointmentResourcesRequest Request schema for the `assignAppointmentResources` operation.
type AssignAppointmentResourcesRequest struct {
	// List of resources that performs or performed job appointment fulfillment.
	Resources []AppointmentResource `json:"resources"`
}

type _AssignAppointmentResourcesRequest AssignAppointmentResourcesRequest

// NewAssignAppointmentResourcesRequest instantiates a new AssignAppointmentResourcesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignAppointmentResourcesRequest(resources []AppointmentResource) *AssignAppointmentResourcesRequest {
	this := AssignAppointmentResourcesRequest{}
	this.Resources = resources
	return &this
}

// NewAssignAppointmentResourcesRequestWithDefaults instantiates a new AssignAppointmentResourcesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignAppointmentResourcesRequestWithDefaults() *AssignAppointmentResourcesRequest {
	this := AssignAppointmentResourcesRequest{}
	return &this
}

// GetResources returns the Resources field value
func (o *AssignAppointmentResourcesRequest) GetResources() []AppointmentResource {
	if o == nil {
		var ret []AppointmentResource
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *AssignAppointmentResourcesRequest) GetResourcesOk() ([]AppointmentResource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *AssignAppointmentResourcesRequest) SetResources(v []AppointmentResource) {
	o.Resources = v
}

func (o AssignAppointmentResourcesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resources"] = o.Resources
	return toSerialize, nil
}

type NullableAssignAppointmentResourcesRequest struct {
	value *AssignAppointmentResourcesRequest
	isSet bool
}

func (v NullableAssignAppointmentResourcesRequest) Get() *AssignAppointmentResourcesRequest {
	return v.value
}

func (v *NullableAssignAppointmentResourcesRequest) Set(val *AssignAppointmentResourcesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignAppointmentResourcesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignAppointmentResourcesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignAppointmentResourcesRequest(val *AssignAppointmentResourcesRequest) *NullableAssignAppointmentResourcesRequest {
	return &NullableAssignAppointmentResourcesRequest{value: val, isSet: true}
}

func (v NullableAssignAppointmentResourcesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAssignAppointmentResourcesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
