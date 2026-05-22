package vendor_direct_fulfillment_orders_v1

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the OrderAcknowledgementItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderAcknowledgementItem{}

// OrderAcknowledgementItem Details of an individual order being acknowledged.
type OrderAcknowledgementItem struct {
	// The purchase order number for this order. Formatting Notes: alpha-numeric code.
	PurchaseOrderNumber string `json:"purchaseOrderNumber"`
	// The vendor's order number for this order.
	VendorOrderNumber string `json:"vendorOrderNumber"`
	// The date and time when the order is acknowledged, in ISO-8601 date/time format. For example: 2018-07-16T23:00:00Z / 2018-07-16T23:00:00-05:00 / 2018-07-16T23:00:00-08:00.
	AcknowledgementDate   time.Time             `json:"acknowledgementDate"`
	AcknowledgementStatus AcknowledgementStatus `json:"acknowledgementStatus"`
	SellingParty          PartyIdentification   `json:"sellingParty"`
	ShipFromParty         PartyIdentification   `json:"shipFromParty"`
	// Item details including acknowledged quantity.
	ItemAcknowledgements []OrderItemAcknowledgement `json:"itemAcknowledgements"`
}

type _OrderAcknowledgementItem OrderAcknowledgementItem

// NewOrderAcknowledgementItem instantiates a new OrderAcknowledgementItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAcknowledgementItem(purchaseOrderNumber string, vendorOrderNumber string, acknowledgementDate time.Time, acknowledgementStatus AcknowledgementStatus, sellingParty PartyIdentification, shipFromParty PartyIdentification, itemAcknowledgements []OrderItemAcknowledgement) *OrderAcknowledgementItem {
	this := OrderAcknowledgementItem{}
	this.PurchaseOrderNumber = purchaseOrderNumber
	this.VendorOrderNumber = vendorOrderNumber
	this.AcknowledgementDate = acknowledgementDate
	this.AcknowledgementStatus = acknowledgementStatus
	this.SellingParty = sellingParty
	this.ShipFromParty = shipFromParty
	this.ItemAcknowledgements = itemAcknowledgements
	return &this
}

// NewOrderAcknowledgementItemWithDefaults instantiates a new OrderAcknowledgementItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAcknowledgementItemWithDefaults() *OrderAcknowledgementItem {
	this := OrderAcknowledgementItem{}
	return &this
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value
func (o *OrderAcknowledgementItem) GetPurchaseOrderNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PurchaseOrderNumber
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgementItem) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseOrderNumber, true
}

// SetPurchaseOrderNumber sets field value
func (o *OrderAcknowledgementItem) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber = v
}

// GetVendorOrderNumber returns the VendorOrderNumber field value
func (o *OrderAcknowledgementItem) GetVendorOrderNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VendorOrderNumber
}

// GetVendorOrderNumberOk returns a tuple with the VendorOrderNumber field value
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgementItem) GetVendorOrderNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VendorOrderNumber, true
}

// SetVendorOrderNumber sets field value
func (o *OrderAcknowledgementItem) SetVendorOrderNumber(v string) {
	o.VendorOrderNumber = v
}

// GetAcknowledgementDate returns the AcknowledgementDate field value
func (o *OrderAcknowledgementItem) GetAcknowledgementDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.AcknowledgementDate
}

// GetAcknowledgementDateOk returns a tuple with the AcknowledgementDate field value
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgementItem) GetAcknowledgementDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcknowledgementDate, true
}

// SetAcknowledgementDate sets field value
func (o *OrderAcknowledgementItem) SetAcknowledgementDate(v time.Time) {
	o.AcknowledgementDate = v
}

// GetAcknowledgementStatus returns the AcknowledgementStatus field value
func (o *OrderAcknowledgementItem) GetAcknowledgementStatus() AcknowledgementStatus {
	if o == nil {
		var ret AcknowledgementStatus
		return ret
	}

	return o.AcknowledgementStatus
}

// GetAcknowledgementStatusOk returns a tuple with the AcknowledgementStatus field value
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgementItem) GetAcknowledgementStatusOk() (*AcknowledgementStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcknowledgementStatus, true
}

// SetAcknowledgementStatus sets field value
func (o *OrderAcknowledgementItem) SetAcknowledgementStatus(v AcknowledgementStatus) {
	o.AcknowledgementStatus = v
}

// GetSellingParty returns the SellingParty field value
func (o *OrderAcknowledgementItem) GetSellingParty() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.SellingParty
}

// GetSellingPartyOk returns a tuple with the SellingParty field value
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgementItem) GetSellingPartyOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellingParty, true
}

// SetSellingParty sets field value
func (o *OrderAcknowledgementItem) SetSellingParty(v PartyIdentification) {
	o.SellingParty = v
}

// GetShipFromParty returns the ShipFromParty field value
func (o *OrderAcknowledgementItem) GetShipFromParty() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.ShipFromParty
}

// GetShipFromPartyOk returns a tuple with the ShipFromParty field value
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgementItem) GetShipFromPartyOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipFromParty, true
}

// SetShipFromParty sets field value
func (o *OrderAcknowledgementItem) SetShipFromParty(v PartyIdentification) {
	o.ShipFromParty = v
}

// GetItemAcknowledgements returns the ItemAcknowledgements field value
func (o *OrderAcknowledgementItem) GetItemAcknowledgements() []OrderItemAcknowledgement {
	if o == nil {
		var ret []OrderItemAcknowledgement
		return ret
	}

	return o.ItemAcknowledgements
}

// GetItemAcknowledgementsOk returns a tuple with the ItemAcknowledgements field value
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgementItem) GetItemAcknowledgementsOk() ([]OrderItemAcknowledgement, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemAcknowledgements, true
}

// SetItemAcknowledgements sets field value
func (o *OrderAcknowledgementItem) SetItemAcknowledgements(v []OrderItemAcknowledgement) {
	o.ItemAcknowledgements = v
}

func (o OrderAcknowledgementItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["purchaseOrderNumber"] = o.PurchaseOrderNumber
	toSerialize["vendorOrderNumber"] = o.VendorOrderNumber
	toSerialize["acknowledgementDate"] = o.AcknowledgementDate
	toSerialize["acknowledgementStatus"] = o.AcknowledgementStatus
	toSerialize["sellingParty"] = o.SellingParty
	toSerialize["shipFromParty"] = o.ShipFromParty
	toSerialize["itemAcknowledgements"] = o.ItemAcknowledgements
	return toSerialize, nil
}

type NullableOrderAcknowledgementItem struct {
	value *OrderAcknowledgementItem
	isSet bool
}

func (v NullableOrderAcknowledgementItem) Get() *OrderAcknowledgementItem {
	return v.value
}

func (v *NullableOrderAcknowledgementItem) Set(val *OrderAcknowledgementItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAcknowledgementItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAcknowledgementItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAcknowledgementItem(val *OrderAcknowledgementItem) *NullableOrderAcknowledgementItem {
	return &NullableOrderAcknowledgementItem{value: val, isSet: true}
}

func (v NullableOrderAcknowledgementItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderAcknowledgementItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
