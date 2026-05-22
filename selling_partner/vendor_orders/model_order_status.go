package vendor_orders

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the OrderStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderStatus{}

// OrderStatus Current status of a purchase order.
type OrderStatus struct {
	// The buyer's purchase order number for this order. Formatting Notes: 8-character alpha-numeric code.
	PurchaseOrderNumber string `json:"purchaseOrderNumber"`
	// The status of the buyer's purchase order for this order.
	PurchaseOrderStatus string `json:"purchaseOrderStatus"`
	// The date the purchase order was placed. Must be in ISO-8601 date/time format.
	PurchaseOrderDate time.Time `json:"purchaseOrderDate"`
	// The date when the purchase order was last updated. Must be in ISO-8601 date/time format.
	LastUpdatedDate *time.Time          `json:"lastUpdatedDate,omitempty"`
	SellingParty    PartyIdentification `json:"sellingParty"`
	ShipToParty     PartyIdentification `json:"shipToParty"`
	// Detailed description of items order status.
	ItemStatus []OrderItemStatus `json:"itemStatus"`
}

type _OrderStatus OrderStatus

// NewOrderStatus instantiates a new OrderStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderStatus(purchaseOrderNumber string, purchaseOrderStatus string, purchaseOrderDate time.Time, sellingParty PartyIdentification, shipToParty PartyIdentification, itemStatus []OrderItemStatus) *OrderStatus {
	this := OrderStatus{}
	this.PurchaseOrderNumber = purchaseOrderNumber
	this.PurchaseOrderStatus = purchaseOrderStatus
	this.PurchaseOrderDate = purchaseOrderDate
	this.SellingParty = sellingParty
	this.ShipToParty = shipToParty
	this.ItemStatus = itemStatus
	return &this
}

// NewOrderStatusWithDefaults instantiates a new OrderStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderStatusWithDefaults() *OrderStatus {
	this := OrderStatus{}
	return &this
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value
func (o *OrderStatus) GetPurchaseOrderNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PurchaseOrderNumber
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value
// and a boolean to check if the value has been set.
func (o *OrderStatus) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseOrderNumber, true
}

// SetPurchaseOrderNumber sets field value
func (o *OrderStatus) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber = v
}

// GetPurchaseOrderStatus returns the PurchaseOrderStatus field value
func (o *OrderStatus) GetPurchaseOrderStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PurchaseOrderStatus
}

// GetPurchaseOrderStatusOk returns a tuple with the PurchaseOrderStatus field value
// and a boolean to check if the value has been set.
func (o *OrderStatus) GetPurchaseOrderStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseOrderStatus, true
}

// SetPurchaseOrderStatus sets field value
func (o *OrderStatus) SetPurchaseOrderStatus(v string) {
	o.PurchaseOrderStatus = v
}

// GetPurchaseOrderDate returns the PurchaseOrderDate field value
func (o *OrderStatus) GetPurchaseOrderDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PurchaseOrderDate
}

// GetPurchaseOrderDateOk returns a tuple with the PurchaseOrderDate field value
// and a boolean to check if the value has been set.
func (o *OrderStatus) GetPurchaseOrderDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseOrderDate, true
}

// SetPurchaseOrderDate sets field value
func (o *OrderStatus) SetPurchaseOrderDate(v time.Time) {
	o.PurchaseOrderDate = v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *OrderStatus) GetLastUpdatedDate() time.Time {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatus) GetLastUpdatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *OrderStatus) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given time.Time and assigns it to the LastUpdatedDate field.
func (o *OrderStatus) SetLastUpdatedDate(v time.Time) {
	o.LastUpdatedDate = &v
}

// GetSellingParty returns the SellingParty field value
func (o *OrderStatus) GetSellingParty() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.SellingParty
}

// GetSellingPartyOk returns a tuple with the SellingParty field value
// and a boolean to check if the value has been set.
func (o *OrderStatus) GetSellingPartyOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellingParty, true
}

// SetSellingParty sets field value
func (o *OrderStatus) SetSellingParty(v PartyIdentification) {
	o.SellingParty = v
}

// GetShipToParty returns the ShipToParty field value
func (o *OrderStatus) GetShipToParty() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.ShipToParty
}

// GetShipToPartyOk returns a tuple with the ShipToParty field value
// and a boolean to check if the value has been set.
func (o *OrderStatus) GetShipToPartyOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipToParty, true
}

// SetShipToParty sets field value
func (o *OrderStatus) SetShipToParty(v PartyIdentification) {
	o.ShipToParty = v
}

// GetItemStatus returns the ItemStatus field value
func (o *OrderStatus) GetItemStatus() []OrderItemStatus {
	if o == nil {
		var ret []OrderItemStatus
		return ret
	}

	return o.ItemStatus
}

// GetItemStatusOk returns a tuple with the ItemStatus field value
// and a boolean to check if the value has been set.
func (o *OrderStatus) GetItemStatusOk() ([]OrderItemStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemStatus, true
}

// SetItemStatus sets field value
func (o *OrderStatus) SetItemStatus(v []OrderItemStatus) {
	o.ItemStatus = v
}

func (o OrderStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["purchaseOrderNumber"] = o.PurchaseOrderNumber
	toSerialize["purchaseOrderStatus"] = o.PurchaseOrderStatus
	toSerialize["purchaseOrderDate"] = o.PurchaseOrderDate
	if !IsNil(o.LastUpdatedDate) {
		toSerialize["lastUpdatedDate"] = o.LastUpdatedDate
	}
	toSerialize["sellingParty"] = o.SellingParty
	toSerialize["shipToParty"] = o.ShipToParty
	toSerialize["itemStatus"] = o.ItemStatus
	return toSerialize, nil
}

type NullableOrderStatus struct {
	value *OrderStatus
	isSet bool
}

func (v NullableOrderStatus) Get() *OrderStatus {
	return v.value
}

func (v *NullableOrderStatus) Set(val *OrderStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderStatus(val *OrderStatus) *NullableOrderStatus {
	return &NullableOrderStatus{value: val, isSet: true}
}

func (v NullableOrderStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
