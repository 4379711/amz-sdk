package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the Shipment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Shipment{}

// Shipment The shipment related data.
type Shipment struct {
	// The unique shipment identifier.
	ShipmentId string `json:"shipmentId"`
	// Client reference id.
	ClientReferenceId string        `json:"clientReferenceId"`
	ShipFrom          Address       `json:"shipFrom"`
	ShipTo            Address       `json:"shipTo"`
	AcceptedRate      *AcceptedRate `json:"acceptedRate,omitempty"`
	Shipper           *Party        `json:"shipper,omitempty"`
	// A list of container.
	Containers []Container `json:"containers"`
}

type _Shipment Shipment

// NewShipment instantiates a new Shipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipment(shipmentId string, clientReferenceId string, shipFrom Address, shipTo Address, containers []Container) *Shipment {
	this := Shipment{}
	this.ShipmentId = shipmentId
	this.ClientReferenceId = clientReferenceId
	this.ShipFrom = shipFrom
	this.ShipTo = shipTo
	this.Containers = containers
	return &this
}

// NewShipmentWithDefaults instantiates a new Shipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentWithDefaults() *Shipment {
	this := Shipment{}
	return &this
}

// GetShipmentId returns the ShipmentId field value
func (o *Shipment) GetShipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value
// and a boolean to check if the value has been set.
func (o *Shipment) GetShipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentId, true
}

// SetShipmentId sets field value
func (o *Shipment) SetShipmentId(v string) {
	o.ShipmentId = v
}

// GetClientReferenceId returns the ClientReferenceId field value
func (o *Shipment) GetClientReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientReferenceId
}

// GetClientReferenceIdOk returns a tuple with the ClientReferenceId field value
// and a boolean to check if the value has been set.
func (o *Shipment) GetClientReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientReferenceId, true
}

// SetClientReferenceId sets field value
func (o *Shipment) SetClientReferenceId(v string) {
	o.ClientReferenceId = v
}

// GetShipFrom returns the ShipFrom field value
func (o *Shipment) GetShipFrom() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.ShipFrom
}

// GetShipFromOk returns a tuple with the ShipFrom field value
// and a boolean to check if the value has been set.
func (o *Shipment) GetShipFromOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipFrom, true
}

// SetShipFrom sets field value
func (o *Shipment) SetShipFrom(v Address) {
	o.ShipFrom = v
}

// GetShipTo returns the ShipTo field value
func (o *Shipment) GetShipTo() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.ShipTo
}

// GetShipToOk returns a tuple with the ShipTo field value
// and a boolean to check if the value has been set.
func (o *Shipment) GetShipToOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipTo, true
}

// SetShipTo sets field value
func (o *Shipment) SetShipTo(v Address) {
	o.ShipTo = v
}

// GetAcceptedRate returns the AcceptedRate field value if set, zero value otherwise.
func (o *Shipment) GetAcceptedRate() AcceptedRate {
	if o == nil || IsNil(o.AcceptedRate) {
		var ret AcceptedRate
		return ret
	}
	return *o.AcceptedRate
}

// GetAcceptedRateOk returns a tuple with the AcceptedRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetAcceptedRateOk() (*AcceptedRate, bool) {
	if o == nil || IsNil(o.AcceptedRate) {
		return nil, false
	}
	return o.AcceptedRate, true
}

// HasAcceptedRate returns a boolean if a field has been set.
func (o *Shipment) HasAcceptedRate() bool {
	if o != nil && !IsNil(o.AcceptedRate) {
		return true
	}

	return false
}

// SetAcceptedRate gets a reference to the given AcceptedRate and assigns it to the AcceptedRate field.
func (o *Shipment) SetAcceptedRate(v AcceptedRate) {
	o.AcceptedRate = &v
}

// GetShipper returns the Shipper field value if set, zero value otherwise.
func (o *Shipment) GetShipper() Party {
	if o == nil || IsNil(o.Shipper) {
		var ret Party
		return ret
	}
	return *o.Shipper
}

// GetShipperOk returns a tuple with the Shipper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetShipperOk() (*Party, bool) {
	if o == nil || IsNil(o.Shipper) {
		return nil, false
	}
	return o.Shipper, true
}

// HasShipper returns a boolean if a field has been set.
func (o *Shipment) HasShipper() bool {
	if o != nil && !IsNil(o.Shipper) {
		return true
	}

	return false
}

// SetShipper gets a reference to the given Party and assigns it to the Shipper field.
func (o *Shipment) SetShipper(v Party) {
	o.Shipper = &v
}

// GetContainers returns the Containers field value
func (o *Shipment) GetContainers() []Container {
	if o == nil {
		var ret []Container
		return ret
	}

	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value
// and a boolean to check if the value has been set.
func (o *Shipment) GetContainersOk() ([]Container, bool) {
	if o == nil {
		return nil, false
	}
	return o.Containers, true
}

// SetContainers sets field value
func (o *Shipment) SetContainers(v []Container) {
	o.Containers = v
}

func (o Shipment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipmentId"] = o.ShipmentId
	toSerialize["clientReferenceId"] = o.ClientReferenceId
	toSerialize["shipFrom"] = o.ShipFrom
	toSerialize["shipTo"] = o.ShipTo
	if !IsNil(o.AcceptedRate) {
		toSerialize["acceptedRate"] = o.AcceptedRate
	}
	if !IsNil(o.Shipper) {
		toSerialize["shipper"] = o.Shipper
	}
	toSerialize["containers"] = o.Containers
	return toSerialize, nil
}

type NullableShipment struct {
	value *Shipment
	isSet bool
}

func (v NullableShipment) Get() *Shipment {
	return v.value
}

func (v *NullableShipment) Set(val *Shipment) {
	v.value = val
	v.isSet = true
}

func (v NullableShipment) IsSet() bool {
	return v.isSet
}

func (v *NullableShipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipment(val *Shipment) *NullableShipment {
	return &NullableShipment{value: val, isSet: true}
}

func (v NullableShipment) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
