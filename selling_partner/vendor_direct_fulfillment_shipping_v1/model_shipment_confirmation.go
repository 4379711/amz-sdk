package vendor_direct_fulfillment_shipping_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ShipmentConfirmation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentConfirmation{}

// ShipmentConfirmation Represents the confirmation details of a shipment, including the purchase order number and other shipment details.
type ShipmentConfirmation struct {
	// Purchase order number corresponding to the shipment.
	PurchaseOrderNumber string              `json:"purchaseOrderNumber" validate:"regexp=^[a-zA-Z0-9]+$"`
	ShipmentDetails     ShipmentDetails     `json:"shipmentDetails"`
	SellingParty        PartyIdentification `json:"sellingParty"`
	ShipFromParty       PartyIdentification `json:"shipFromParty"`
	// Provide the details of the items in this shipment. If any of the item details field is common at a package or a pallet level, then provide them at the corresponding package.
	Items []Item `json:"items"`
	// Provide the details of the items in this shipment. If any of the item details field is common at a package or a pallet level, then provide them at the corresponding package.
	Containers []Container `json:"containers,omitempty"`
}

type _ShipmentConfirmation ShipmentConfirmation

// NewShipmentConfirmation instantiates a new ShipmentConfirmation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentConfirmation(purchaseOrderNumber string, shipmentDetails ShipmentDetails, sellingParty PartyIdentification, shipFromParty PartyIdentification, items []Item) *ShipmentConfirmation {
	this := ShipmentConfirmation{}
	this.PurchaseOrderNumber = purchaseOrderNumber
	this.ShipmentDetails = shipmentDetails
	this.SellingParty = sellingParty
	this.ShipFromParty = shipFromParty
	this.Items = items
	return &this
}

// NewShipmentConfirmationWithDefaults instantiates a new ShipmentConfirmation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentConfirmationWithDefaults() *ShipmentConfirmation {
	this := ShipmentConfirmation{}
	return &this
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value
func (o *ShipmentConfirmation) GetPurchaseOrderNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PurchaseOrderNumber
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value
// and a boolean to check if the value has been set.
func (o *ShipmentConfirmation) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseOrderNumber, true
}

// SetPurchaseOrderNumber sets field value
func (o *ShipmentConfirmation) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber = v
}

// GetShipmentDetails returns the ShipmentDetails field value
func (o *ShipmentConfirmation) GetShipmentDetails() ShipmentDetails {
	if o == nil {
		var ret ShipmentDetails
		return ret
	}

	return o.ShipmentDetails
}

// GetShipmentDetailsOk returns a tuple with the ShipmentDetails field value
// and a boolean to check if the value has been set.
func (o *ShipmentConfirmation) GetShipmentDetailsOk() (*ShipmentDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentDetails, true
}

// SetShipmentDetails sets field value
func (o *ShipmentConfirmation) SetShipmentDetails(v ShipmentDetails) {
	o.ShipmentDetails = v
}

// GetSellingParty returns the SellingParty field value
func (o *ShipmentConfirmation) GetSellingParty() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.SellingParty
}

// GetSellingPartyOk returns a tuple with the SellingParty field value
// and a boolean to check if the value has been set.
func (o *ShipmentConfirmation) GetSellingPartyOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellingParty, true
}

// SetSellingParty sets field value
func (o *ShipmentConfirmation) SetSellingParty(v PartyIdentification) {
	o.SellingParty = v
}

// GetShipFromParty returns the ShipFromParty field value
func (o *ShipmentConfirmation) GetShipFromParty() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.ShipFromParty
}

// GetShipFromPartyOk returns a tuple with the ShipFromParty field value
// and a boolean to check if the value has been set.
func (o *ShipmentConfirmation) GetShipFromPartyOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipFromParty, true
}

// SetShipFromParty sets field value
func (o *ShipmentConfirmation) SetShipFromParty(v PartyIdentification) {
	o.ShipFromParty = v
}

// GetItems returns the Items field value
func (o *ShipmentConfirmation) GetItems() []Item {
	if o == nil {
		var ret []Item
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ShipmentConfirmation) GetItemsOk() ([]Item, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ShipmentConfirmation) SetItems(v []Item) {
	o.Items = v
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *ShipmentConfirmation) GetContainers() []Container {
	if o == nil || IsNil(o.Containers) {
		var ret []Container
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentConfirmation) GetContainersOk() ([]Container, bool) {
	if o == nil || IsNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *ShipmentConfirmation) HasContainers() bool {
	if o != nil && !IsNil(o.Containers) {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []Container and assigns it to the Containers field.
func (o *ShipmentConfirmation) SetContainers(v []Container) {
	o.Containers = v
}

func (o ShipmentConfirmation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["purchaseOrderNumber"] = o.PurchaseOrderNumber
	toSerialize["shipmentDetails"] = o.ShipmentDetails
	toSerialize["sellingParty"] = o.SellingParty
	toSerialize["shipFromParty"] = o.ShipFromParty
	toSerialize["items"] = o.Items
	if !IsNil(o.Containers) {
		toSerialize["containers"] = o.Containers
	}
	return toSerialize, nil
}

type NullableShipmentConfirmation struct {
	value *ShipmentConfirmation
	isSet bool
}

func (v NullableShipmentConfirmation) Get() *ShipmentConfirmation {
	return v.value
}

func (v *NullableShipmentConfirmation) Set(val *ShipmentConfirmation) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentConfirmation) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentConfirmation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentConfirmation(val *ShipmentConfirmation) *NullableShipmentConfirmation {
	return &NullableShipmentConfirmation{value: val, isSet: true}
}

func (v NullableShipmentConfirmation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShipmentConfirmation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
