package vendor_orders

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the OrderAcknowledgement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderAcknowledgement{}

// OrderAcknowledgement Represents an acknowledgement for an order, including the purchase order number, selling party details, acknowledgement date, and a list of acknowledged items.
type OrderAcknowledgement struct {
	// The purchase order number. Formatting Notes: 8-character alpha-numeric code.
	PurchaseOrderNumber string              `json:"purchaseOrderNumber"`
	SellingParty        PartyIdentification `json:"sellingParty"`
	// The date and time when the purchase order is acknowledged, in ISO-8601 date/time format.
	AcknowledgementDate time.Time `json:"acknowledgementDate"`
	// A list of the items being acknowledged with associated details.
	Items []OrderAcknowledgementItem `json:"items"`
}

type _OrderAcknowledgement OrderAcknowledgement

// NewOrderAcknowledgement instantiates a new OrderAcknowledgement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAcknowledgement(purchaseOrderNumber string, sellingParty PartyIdentification, acknowledgementDate time.Time, items []OrderAcknowledgementItem) *OrderAcknowledgement {
	this := OrderAcknowledgement{}
	this.PurchaseOrderNumber = purchaseOrderNumber
	this.SellingParty = sellingParty
	this.AcknowledgementDate = acknowledgementDate
	this.Items = items
	return &this
}

// NewOrderAcknowledgementWithDefaults instantiates a new OrderAcknowledgement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAcknowledgementWithDefaults() *OrderAcknowledgement {
	this := OrderAcknowledgement{}
	return &this
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value
func (o *OrderAcknowledgement) GetPurchaseOrderNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PurchaseOrderNumber
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgement) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseOrderNumber, true
}

// SetPurchaseOrderNumber sets field value
func (o *OrderAcknowledgement) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber = v
}

// GetSellingParty returns the SellingParty field value
func (o *OrderAcknowledgement) GetSellingParty() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.SellingParty
}

// GetSellingPartyOk returns a tuple with the SellingParty field value
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgement) GetSellingPartyOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellingParty, true
}

// SetSellingParty sets field value
func (o *OrderAcknowledgement) SetSellingParty(v PartyIdentification) {
	o.SellingParty = v
}

// GetAcknowledgementDate returns the AcknowledgementDate field value
func (o *OrderAcknowledgement) GetAcknowledgementDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.AcknowledgementDate
}

// GetAcknowledgementDateOk returns a tuple with the AcknowledgementDate field value
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgement) GetAcknowledgementDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcknowledgementDate, true
}

// SetAcknowledgementDate sets field value
func (o *OrderAcknowledgement) SetAcknowledgementDate(v time.Time) {
	o.AcknowledgementDate = v
}

// GetItems returns the Items field value
func (o *OrderAcknowledgement) GetItems() []OrderAcknowledgementItem {
	if o == nil {
		var ret []OrderAcknowledgementItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgement) GetItemsOk() ([]OrderAcknowledgementItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *OrderAcknowledgement) SetItems(v []OrderAcknowledgementItem) {
	o.Items = v
}

func (o OrderAcknowledgement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["purchaseOrderNumber"] = o.PurchaseOrderNumber
	toSerialize["sellingParty"] = o.SellingParty
	toSerialize["acknowledgementDate"] = o.AcknowledgementDate
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableOrderAcknowledgement struct {
	value *OrderAcknowledgement
	isSet bool
}

func (v NullableOrderAcknowledgement) Get() *OrderAcknowledgement {
	return v.value
}

func (v *NullableOrderAcknowledgement) Set(val *OrderAcknowledgement) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAcknowledgement) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAcknowledgement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAcknowledgement(val *OrderAcknowledgement) *NullableOrderAcknowledgement {
	return &NullableOrderAcknowledgement{value: val, isSet: true}
}

func (v NullableOrderAcknowledgement) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderAcknowledgement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
