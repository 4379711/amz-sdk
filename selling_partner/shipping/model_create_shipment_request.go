package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateShipmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateShipmentRequest{}

// CreateShipmentRequest The request schema for the createShipment operation.
type CreateShipmentRequest struct {
	// Client reference id.
	ClientReferenceId string  `json:"clientReferenceId"`
	ShipTo            Address `json:"shipTo"`
	ShipFrom          Address `json:"shipFrom"`
	// A list of container.
	Containers []Container `json:"containers"`
}

type _CreateShipmentRequest CreateShipmentRequest

// NewCreateShipmentRequest instantiates a new CreateShipmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShipmentRequest(clientReferenceId string, shipTo Address, shipFrom Address, containers []Container) *CreateShipmentRequest {
	this := CreateShipmentRequest{}
	this.ClientReferenceId = clientReferenceId
	this.ShipTo = shipTo
	this.ShipFrom = shipFrom
	this.Containers = containers
	return &this
}

// NewCreateShipmentRequestWithDefaults instantiates a new CreateShipmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShipmentRequestWithDefaults() *CreateShipmentRequest {
	this := CreateShipmentRequest{}
	return &this
}

// GetClientReferenceId returns the ClientReferenceId field value
func (o *CreateShipmentRequest) GetClientReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientReferenceId
}

// GetClientReferenceIdOk returns a tuple with the ClientReferenceId field value
// and a boolean to check if the value has been set.
func (o *CreateShipmentRequest) GetClientReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientReferenceId, true
}

// SetClientReferenceId sets field value
func (o *CreateShipmentRequest) SetClientReferenceId(v string) {
	o.ClientReferenceId = v
}

// GetShipTo returns the ShipTo field value
func (o *CreateShipmentRequest) GetShipTo() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.ShipTo
}

// GetShipToOk returns a tuple with the ShipTo field value
// and a boolean to check if the value has been set.
func (o *CreateShipmentRequest) GetShipToOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipTo, true
}

// SetShipTo sets field value
func (o *CreateShipmentRequest) SetShipTo(v Address) {
	o.ShipTo = v
}

// GetShipFrom returns the ShipFrom field value
func (o *CreateShipmentRequest) GetShipFrom() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.ShipFrom
}

// GetShipFromOk returns a tuple with the ShipFrom field value
// and a boolean to check if the value has been set.
func (o *CreateShipmentRequest) GetShipFromOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipFrom, true
}

// SetShipFrom sets field value
func (o *CreateShipmentRequest) SetShipFrom(v Address) {
	o.ShipFrom = v
}

// GetContainers returns the Containers field value
func (o *CreateShipmentRequest) GetContainers() []Container {
	if o == nil {
		var ret []Container
		return ret
	}

	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value
// and a boolean to check if the value has been set.
func (o *CreateShipmentRequest) GetContainersOk() ([]Container, bool) {
	if o == nil {
		return nil, false
	}
	return o.Containers, true
}

// SetContainers sets field value
func (o *CreateShipmentRequest) SetContainers(v []Container) {
	o.Containers = v
}

func (o CreateShipmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientReferenceId"] = o.ClientReferenceId
	toSerialize["shipTo"] = o.ShipTo
	toSerialize["shipFrom"] = o.ShipFrom
	toSerialize["containers"] = o.Containers
	return toSerialize, nil
}

type NullableCreateShipmentRequest struct {
	value *CreateShipmentRequest
	isSet bool
}

func (v NullableCreateShipmentRequest) Get() *CreateShipmentRequest {
	return v.value
}

func (v *NullableCreateShipmentRequest) Set(val *CreateShipmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShipmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShipmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShipmentRequest(val *CreateShipmentRequest) *NullableCreateShipmentRequest {
	return &NullableCreateShipmentRequest{value: val, isSet: true}
}

func (v NullableCreateShipmentRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateShipmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
