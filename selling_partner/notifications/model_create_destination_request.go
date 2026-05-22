package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateDestinationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDestinationRequest{}

// CreateDestinationRequest The request schema for the `createDestination` operation.
type CreateDestinationRequest struct {
	ResourceSpecification DestinationResourceSpecification `json:"resourceSpecification"`
	// A developer-defined name to help identify this destination.
	Name string `json:"name"`
}

type _CreateDestinationRequest CreateDestinationRequest

// NewCreateDestinationRequest instantiates a new CreateDestinationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDestinationRequest(resourceSpecification DestinationResourceSpecification, name string) *CreateDestinationRequest {
	this := CreateDestinationRequest{}
	this.ResourceSpecification = resourceSpecification
	this.Name = name
	return &this
}

// NewCreateDestinationRequestWithDefaults instantiates a new CreateDestinationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDestinationRequestWithDefaults() *CreateDestinationRequest {
	this := CreateDestinationRequest{}
	return &this
}

// GetResourceSpecification returns the ResourceSpecification field value
func (o *CreateDestinationRequest) GetResourceSpecification() DestinationResourceSpecification {
	if o == nil {
		var ret DestinationResourceSpecification
		return ret
	}

	return o.ResourceSpecification
}

// GetResourceSpecificationOk returns a tuple with the ResourceSpecification field value
// and a boolean to check if the value has been set.
func (o *CreateDestinationRequest) GetResourceSpecificationOk() (*DestinationResourceSpecification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceSpecification, true
}

// SetResourceSpecification sets field value
func (o *CreateDestinationRequest) SetResourceSpecification(v DestinationResourceSpecification) {
	o.ResourceSpecification = v
}

// GetName returns the Name field value
func (o *CreateDestinationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateDestinationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateDestinationRequest) SetName(v string) {
	o.Name = v
}

func (o CreateDestinationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceSpecification"] = o.ResourceSpecification
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableCreateDestinationRequest struct {
	value *CreateDestinationRequest
	isSet bool
}

func (v NullableCreateDestinationRequest) Get() *CreateDestinationRequest {
	return v.value
}

func (v *NullableCreateDestinationRequest) Set(val *CreateDestinationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDestinationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDestinationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDestinationRequest(val *CreateDestinationRequest) *NullableCreateDestinationRequest {
	return &NullableCreateDestinationRequest{value: val, isSet: true}
}

func (v NullableCreateDestinationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateDestinationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
