package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ShipmentTransportationConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentTransportationConfiguration{}

// ShipmentTransportationConfiguration Details needed to generate the transportation options.
type ShipmentTransportationConfiguration struct {
	ContactInformation *ContactInformation `json:"contactInformation,omitempty"`
	FreightInformation *FreightInformation `json:"freightInformation,omitempty"`
	// List of pallet configuration inputs.
	Pallets           []PalletInput `json:"pallets,omitempty"`
	ReadyToShipWindow WindowInput   `json:"readyToShipWindow"`
	// Identifier of a shipment. A shipment contains the boxes and units being inbounded.
	ShipmentId string `json:"shipmentId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _ShipmentTransportationConfiguration ShipmentTransportationConfiguration

// NewShipmentTransportationConfiguration instantiates a new ShipmentTransportationConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentTransportationConfiguration(readyToShipWindow WindowInput, shipmentId string) *ShipmentTransportationConfiguration {
	this := ShipmentTransportationConfiguration{}
	this.ReadyToShipWindow = readyToShipWindow
	this.ShipmentId = shipmentId
	return &this
}

// NewShipmentTransportationConfigurationWithDefaults instantiates a new ShipmentTransportationConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentTransportationConfigurationWithDefaults() *ShipmentTransportationConfiguration {
	this := ShipmentTransportationConfiguration{}
	return &this
}

// GetContactInformation returns the ContactInformation field value if set, zero value otherwise.
func (o *ShipmentTransportationConfiguration) GetContactInformation() ContactInformation {
	if o == nil || IsNil(o.ContactInformation) {
		var ret ContactInformation
		return ret
	}
	return *o.ContactInformation
}

// GetContactInformationOk returns a tuple with the ContactInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentTransportationConfiguration) GetContactInformationOk() (*ContactInformation, bool) {
	if o == nil || IsNil(o.ContactInformation) {
		return nil, false
	}
	return o.ContactInformation, true
}

// HasContactInformation returns a boolean if a field has been set.
func (o *ShipmentTransportationConfiguration) HasContactInformation() bool {
	if o != nil && !IsNil(o.ContactInformation) {
		return true
	}

	return false
}

// SetContactInformation gets a reference to the given ContactInformation and assigns it to the ContactInformation field.
func (o *ShipmentTransportationConfiguration) SetContactInformation(v ContactInformation) {
	o.ContactInformation = &v
}

// GetFreightInformation returns the FreightInformation field value if set, zero value otherwise.
func (o *ShipmentTransportationConfiguration) GetFreightInformation() FreightInformation {
	if o == nil || IsNil(o.FreightInformation) {
		var ret FreightInformation
		return ret
	}
	return *o.FreightInformation
}

// GetFreightInformationOk returns a tuple with the FreightInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentTransportationConfiguration) GetFreightInformationOk() (*FreightInformation, bool) {
	if o == nil || IsNil(o.FreightInformation) {
		return nil, false
	}
	return o.FreightInformation, true
}

// HasFreightInformation returns a boolean if a field has been set.
func (o *ShipmentTransportationConfiguration) HasFreightInformation() bool {
	if o != nil && !IsNil(o.FreightInformation) {
		return true
	}

	return false
}

// SetFreightInformation gets a reference to the given FreightInformation and assigns it to the FreightInformation field.
func (o *ShipmentTransportationConfiguration) SetFreightInformation(v FreightInformation) {
	o.FreightInformation = &v
}

// GetPallets returns the Pallets field value if set, zero value otherwise.
func (o *ShipmentTransportationConfiguration) GetPallets() []PalletInput {
	if o == nil || IsNil(o.Pallets) {
		var ret []PalletInput
		return ret
	}
	return o.Pallets
}

// GetPalletsOk returns a tuple with the Pallets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentTransportationConfiguration) GetPalletsOk() ([]PalletInput, bool) {
	if o == nil || IsNil(o.Pallets) {
		return nil, false
	}
	return o.Pallets, true
}

// HasPallets returns a boolean if a field has been set.
func (o *ShipmentTransportationConfiguration) HasPallets() bool {
	if o != nil && !IsNil(o.Pallets) {
		return true
	}

	return false
}

// SetPallets gets a reference to the given []PalletInput and assigns it to the Pallets field.
func (o *ShipmentTransportationConfiguration) SetPallets(v []PalletInput) {
	o.Pallets = v
}

// GetReadyToShipWindow returns the ReadyToShipWindow field value
func (o *ShipmentTransportationConfiguration) GetReadyToShipWindow() WindowInput {
	if o == nil {
		var ret WindowInput
		return ret
	}

	return o.ReadyToShipWindow
}

// GetReadyToShipWindowOk returns a tuple with the ReadyToShipWindow field value
// and a boolean to check if the value has been set.
func (o *ShipmentTransportationConfiguration) GetReadyToShipWindowOk() (*WindowInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadyToShipWindow, true
}

// SetReadyToShipWindow sets field value
func (o *ShipmentTransportationConfiguration) SetReadyToShipWindow(v WindowInput) {
	o.ReadyToShipWindow = v
}

// GetShipmentId returns the ShipmentId field value
func (o *ShipmentTransportationConfiguration) GetShipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value
// and a boolean to check if the value has been set.
func (o *ShipmentTransportationConfiguration) GetShipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentId, true
}

// SetShipmentId sets field value
func (o *ShipmentTransportationConfiguration) SetShipmentId(v string) {
	o.ShipmentId = v
}

func (o ShipmentTransportationConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContactInformation) {
		toSerialize["contactInformation"] = o.ContactInformation
	}
	if !IsNil(o.FreightInformation) {
		toSerialize["freightInformation"] = o.FreightInformation
	}
	if !IsNil(o.Pallets) {
		toSerialize["pallets"] = o.Pallets
	}
	toSerialize["readyToShipWindow"] = o.ReadyToShipWindow
	toSerialize["shipmentId"] = o.ShipmentId
	return toSerialize, nil
}

type NullableShipmentTransportationConfiguration struct {
	value *ShipmentTransportationConfiguration
	isSet bool
}

func (v NullableShipmentTransportationConfiguration) Get() *ShipmentTransportationConfiguration {
	return v.value
}

func (v *NullableShipmentTransportationConfiguration) Set(val *ShipmentTransportationConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentTransportationConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentTransportationConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentTransportationConfiguration(val *ShipmentTransportationConfiguration) *NullableShipmentTransportationConfiguration {
	return &NullableShipmentTransportationConfiguration{value: val, isSet: true}
}

func (v NullableShipmentTransportationConfiguration) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShipmentTransportationConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
