package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the Destination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Destination{}

// Destination Information about the destination created when you call the `createDestination` operation.
type Destination struct {
	// The developer-defined name for this destination.
	Name string `json:"name"`
	// The destination identifier generated when you created the destination.
	DestinationId string              `json:"destinationId"`
	Resource      DestinationResource `json:"resource"`
}

type _Destination Destination

// NewDestination instantiates a new Destination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestination(name string, destinationId string, resource DestinationResource) *Destination {
	this := Destination{}
	this.Name = name
	this.DestinationId = destinationId
	this.Resource = resource
	return &this
}

// NewDestinationWithDefaults instantiates a new Destination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationWithDefaults() *Destination {
	this := Destination{}
	return &this
}

// GetName returns the Name field value
func (o *Destination) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Destination) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Destination) SetName(v string) {
	o.Name = v
}

// GetDestinationId returns the DestinationId field value
func (o *Destination) GetDestinationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationId
}

// GetDestinationIdOk returns a tuple with the DestinationId field value
// and a boolean to check if the value has been set.
func (o *Destination) GetDestinationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationId, true
}

// SetDestinationId sets field value
func (o *Destination) SetDestinationId(v string) {
	o.DestinationId = v
}

// GetResource returns the Resource field value
func (o *Destination) GetResource() DestinationResource {
	if o == nil {
		var ret DestinationResource
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *Destination) GetResourceOk() (*DestinationResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *Destination) SetResource(v DestinationResource) {
	o.Resource = v
}

func (o Destination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["destinationId"] = o.DestinationId
	toSerialize["resource"] = o.Resource
	return toSerialize, nil
}

type NullableDestination struct {
	value *Destination
	isSet bool
}

func (v NullableDestination) Get() *Destination {
	return v.value
}

func (v *NullableDestination) Set(val *Destination) {
	v.value = val
	v.isSet = true
}

func (v NullableDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestination(val *Destination) *NullableDestination {
	return &NullableDestination{value: val, isSet: true}
}

func (v NullableDestination) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
