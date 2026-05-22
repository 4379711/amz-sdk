package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the TransportationSelection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportationSelection{}

// TransportationSelection The transportation option selected to confirm.
type TransportationSelection struct {
	ContactInformation *ContactInformation `json:"contactInformation,omitempty"`
	// Shipment ID that the transportation Option is for.
	ShipmentId string `json:"shipmentId" validate:"regexp=^[a-zA-Z0-9-]*$"`
	// Transportation option being selected for the provided shipment.
	TransportationOptionId string `json:"transportationOptionId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _TransportationSelection TransportationSelection

// NewTransportationSelection instantiates a new TransportationSelection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportationSelection(shipmentId string, transportationOptionId string) *TransportationSelection {
	this := TransportationSelection{}
	this.ShipmentId = shipmentId
	this.TransportationOptionId = transportationOptionId
	return &this
}

// NewTransportationSelectionWithDefaults instantiates a new TransportationSelection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportationSelectionWithDefaults() *TransportationSelection {
	this := TransportationSelection{}
	return &this
}

// GetContactInformation returns the ContactInformation field value if set, zero value otherwise.
func (o *TransportationSelection) GetContactInformation() ContactInformation {
	if o == nil || IsNil(o.ContactInformation) {
		var ret ContactInformation
		return ret
	}
	return *o.ContactInformation
}

// GetContactInformationOk returns a tuple with the ContactInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportationSelection) GetContactInformationOk() (*ContactInformation, bool) {
	if o == nil || IsNil(o.ContactInformation) {
		return nil, false
	}
	return o.ContactInformation, true
}

// HasContactInformation returns a boolean if a field has been set.
func (o *TransportationSelection) HasContactInformation() bool {
	if o != nil && !IsNil(o.ContactInformation) {
		return true
	}

	return false
}

// SetContactInformation gets a reference to the given ContactInformation and assigns it to the ContactInformation field.
func (o *TransportationSelection) SetContactInformation(v ContactInformation) {
	o.ContactInformation = &v
}

// GetShipmentId returns the ShipmentId field value
func (o *TransportationSelection) GetShipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value
// and a boolean to check if the value has been set.
func (o *TransportationSelection) GetShipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentId, true
}

// SetShipmentId sets field value
func (o *TransportationSelection) SetShipmentId(v string) {
	o.ShipmentId = v
}

// GetTransportationOptionId returns the TransportationOptionId field value
func (o *TransportationSelection) GetTransportationOptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransportationOptionId
}

// GetTransportationOptionIdOk returns a tuple with the TransportationOptionId field value
// and a boolean to check if the value has been set.
func (o *TransportationSelection) GetTransportationOptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransportationOptionId, true
}

// SetTransportationOptionId sets field value
func (o *TransportationSelection) SetTransportationOptionId(v string) {
	o.TransportationOptionId = v
}

func (o TransportationSelection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContactInformation) {
		toSerialize["contactInformation"] = o.ContactInformation
	}
	toSerialize["shipmentId"] = o.ShipmentId
	toSerialize["transportationOptionId"] = o.TransportationOptionId
	return toSerialize, nil
}

type NullableTransportationSelection struct {
	value *TransportationSelection
	isSet bool
}

func (v NullableTransportationSelection) Get() *TransportationSelection {
	return v.value
}

func (v *NullableTransportationSelection) Set(val *TransportationSelection) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportationSelection) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportationSelection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportationSelection(val *TransportationSelection) *NullableTransportationSelection {
	return &NullableTransportationSelection{value: val, isSet: true}
}

func (v NullableTransportationSelection) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTransportationSelection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
