package vendor_direct_fulfillment_shipping_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ShipmentStatusUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentStatusUpdate{}

// ShipmentStatusUpdate Represents an update to the status of a shipment.
type ShipmentStatusUpdate struct {
	// Purchase order number of the shipment for which to update the shipment status.
	PurchaseOrderNumber string              `json:"purchaseOrderNumber" validate:"regexp=^[a-zA-Z0-9]+$"`
	SellingParty        PartyIdentification `json:"sellingParty"`
	ShipFromParty       PartyIdentification `json:"shipFromParty"`
	StatusUpdateDetails StatusUpdateDetails `json:"statusUpdateDetails"`
}

type _ShipmentStatusUpdate ShipmentStatusUpdate

// NewShipmentStatusUpdate instantiates a new ShipmentStatusUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentStatusUpdate(purchaseOrderNumber string, sellingParty PartyIdentification, shipFromParty PartyIdentification, statusUpdateDetails StatusUpdateDetails) *ShipmentStatusUpdate {
	this := ShipmentStatusUpdate{}
	this.PurchaseOrderNumber = purchaseOrderNumber
	this.SellingParty = sellingParty
	this.ShipFromParty = shipFromParty
	this.StatusUpdateDetails = statusUpdateDetails
	return &this
}

// NewShipmentStatusUpdateWithDefaults instantiates a new ShipmentStatusUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentStatusUpdateWithDefaults() *ShipmentStatusUpdate {
	this := ShipmentStatusUpdate{}
	return &this
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value
func (o *ShipmentStatusUpdate) GetPurchaseOrderNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PurchaseOrderNumber
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value
// and a boolean to check if the value has been set.
func (o *ShipmentStatusUpdate) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseOrderNumber, true
}

// SetPurchaseOrderNumber sets field value
func (o *ShipmentStatusUpdate) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber = v
}

// GetSellingParty returns the SellingParty field value
func (o *ShipmentStatusUpdate) GetSellingParty() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.SellingParty
}

// GetSellingPartyOk returns a tuple with the SellingParty field value
// and a boolean to check if the value has been set.
func (o *ShipmentStatusUpdate) GetSellingPartyOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellingParty, true
}

// SetSellingParty sets field value
func (o *ShipmentStatusUpdate) SetSellingParty(v PartyIdentification) {
	o.SellingParty = v
}

// GetShipFromParty returns the ShipFromParty field value
func (o *ShipmentStatusUpdate) GetShipFromParty() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.ShipFromParty
}

// GetShipFromPartyOk returns a tuple with the ShipFromParty field value
// and a boolean to check if the value has been set.
func (o *ShipmentStatusUpdate) GetShipFromPartyOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipFromParty, true
}

// SetShipFromParty sets field value
func (o *ShipmentStatusUpdate) SetShipFromParty(v PartyIdentification) {
	o.ShipFromParty = v
}

// GetStatusUpdateDetails returns the StatusUpdateDetails field value
func (o *ShipmentStatusUpdate) GetStatusUpdateDetails() StatusUpdateDetails {
	if o == nil {
		var ret StatusUpdateDetails
		return ret
	}

	return o.StatusUpdateDetails
}

// GetStatusUpdateDetailsOk returns a tuple with the StatusUpdateDetails field value
// and a boolean to check if the value has been set.
func (o *ShipmentStatusUpdate) GetStatusUpdateDetailsOk() (*StatusUpdateDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusUpdateDetails, true
}

// SetStatusUpdateDetails sets field value
func (o *ShipmentStatusUpdate) SetStatusUpdateDetails(v StatusUpdateDetails) {
	o.StatusUpdateDetails = v
}

func (o ShipmentStatusUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["purchaseOrderNumber"] = o.PurchaseOrderNumber
	toSerialize["sellingParty"] = o.SellingParty
	toSerialize["shipFromParty"] = o.ShipFromParty
	toSerialize["statusUpdateDetails"] = o.StatusUpdateDetails
	return toSerialize, nil
}

type NullableShipmentStatusUpdate struct {
	value *ShipmentStatusUpdate
	isSet bool
}

func (v NullableShipmentStatusUpdate) Get() *ShipmentStatusUpdate {
	return v.value
}

func (v *NullableShipmentStatusUpdate) Set(val *ShipmentStatusUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentStatusUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentStatusUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentStatusUpdate(val *ShipmentStatusUpdate) *NullableShipmentStatusUpdate {
	return &NullableShipmentStatusUpdate{value: val, isSet: true}
}

func (v NullableShipmentStatusUpdate) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShipmentStatusUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
