package orders_20260101

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the Order type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Order{}

// Order Comprehensive information about a customer order.
type Order struct {
	// An Amazon-defined order identifier.
	OrderId string `json:"orderId"`
	// Alternative identifiers that can be used to reference this order, such as seller-defined order numbers.
	OrderAliases []Alias `json:"orderAliases,omitempty"`
	// The time when the customer placed the order. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.
	CreatedTime time.Time `json:"createdTime"`
	// The most recent time when any aspect of this order was modified by Amazon or the seller. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.
	LastUpdatedTime time.Time `json:"lastUpdatedTime"`
	// Special programs associated with this order that may affect fulfillment or customer experience.   **Possible values**: `AMAZON_BAZAAR`, `AMAZON_BUSINESS`, `AMAZON_EASY_SHIP`, `AMAZON_HAUL`, `DELIVERY_BY_AMAZON`, `FBM_SHIP_PLUS`, `INVOICE_BY_AMAZON`, `IN_STORE_PICK_UP`, `PREMIUM`, `PREORDER`, `PRIME`
	Programs []string `json:"programs,omitempty"`
	// Other orders that have a direct relationship to this order, such as replacement or exchange orders.
	AssociatedOrders []AssociatedOrder `json:"associatedOrders,omitempty"`
	SalesChannel     SalesChannel      `json:"salesChannel"`
	Buyer            *Buyer            `json:"buyer,omitempty"`
	Recipient        *Recipient        `json:"recipient,omitempty"`
	Proceeds         *OrderProceeds    `json:"proceeds,omitempty"`
	Payment          *OrderPayment     `json:"payment,omitempty"`
	Tax              *OrderTax         `json:"tax,omitempty"`
	Fulfillment      *OrderFulfillment `json:"fulfillment,omitempty"`
	// The list of all order items included in this order.
	OrderItems []OrderItem `json:"orderItems"`
	// Shipping packages created for this order, including tracking information. **Note:** Only available for merchant-fulfilled (FBM) orders.
	Packages []OrderPackage `json:"packages,omitempty"`
	// The list of fulfillment orders associated with this customer order. Each entry corresponds to one fulfillment unit created by Amazon for this order. **Note:** Only available for EasyShip orders at present.
	FulfillmentOrders []FulfillmentOrder `json:"fulfillmentOrders,omitempty"`
}

// NewOrder instantiates a new Order object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrder(orderId string, createdTime time.Time, lastUpdatedTime time.Time, salesChannel SalesChannel, orderItems []OrderItem) *Order {
	this := Order{}
	this.OrderId = orderId
	this.CreatedTime = createdTime
	this.LastUpdatedTime = lastUpdatedTime
	this.SalesChannel = salesChannel
	this.OrderItems = orderItems
	return &this
}

// NewOrderWithDefaults instantiates a new Order object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderWithDefaults() *Order {
	this := Order{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *Order) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *Order) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *Order) SetOrderId(v string) {
	o.OrderId = v
}

// GetOrderAliases returns the OrderAliases field value if set, zero value otherwise.
func (o *Order) GetOrderAliases() []Alias {
	if o == nil || IsNil(o.OrderAliases) {
		var ret []Alias
		return ret
	}
	return o.OrderAliases
}

// GetOrderAliasesOk returns a tuple with the OrderAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetOrderAliasesOk() ([]Alias, bool) {
	if o == nil || IsNil(o.OrderAliases) {
		return nil, false
	}
	return o.OrderAliases, true
}

// HasOrderAliases returns a boolean if a field has been set.
func (o *Order) HasOrderAliases() bool {
	if o != nil && !IsNil(o.OrderAliases) {
		return true
	}

	return false
}

// SetOrderAliases gets a reference to the given []Alias and assigns it to the OrderAliases field.
func (o *Order) SetOrderAliases(v []Alias) {
	o.OrderAliases = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *Order) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *Order) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *Order) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value
func (o *Order) GetLastUpdatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value
// and a boolean to check if the value has been set.
func (o *Order) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedTime, true
}

// SetLastUpdatedTime sets field value
func (o *Order) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = v
}

// GetPrograms returns the Programs field value if set, zero value otherwise.
func (o *Order) GetPrograms() []string {
	if o == nil || IsNil(o.Programs) {
		var ret []string
		return ret
	}
	return o.Programs
}

// GetProgramsOk returns a tuple with the Programs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetProgramsOk() ([]string, bool) {
	if o == nil || IsNil(o.Programs) {
		return nil, false
	}
	return o.Programs, true
}

// HasPrograms returns a boolean if a field has been set.
func (o *Order) HasPrograms() bool {
	if o != nil && !IsNil(o.Programs) {
		return true
	}

	return false
}

// SetPrograms gets a reference to the given []string and assigns it to the Programs field.
func (o *Order) SetPrograms(v []string) {
	o.Programs = v
}

// GetAssociatedOrders returns the AssociatedOrders field value if set, zero value otherwise.
func (o *Order) GetAssociatedOrders() []AssociatedOrder {
	if o == nil || IsNil(o.AssociatedOrders) {
		var ret []AssociatedOrder
		return ret
	}
	return o.AssociatedOrders
}

// GetAssociatedOrdersOk returns a tuple with the AssociatedOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetAssociatedOrdersOk() ([]AssociatedOrder, bool) {
	if o == nil || IsNil(o.AssociatedOrders) {
		return nil, false
	}
	return o.AssociatedOrders, true
}

// HasAssociatedOrders returns a boolean if a field has been set.
func (o *Order) HasAssociatedOrders() bool {
	if o != nil && !IsNil(o.AssociatedOrders) {
		return true
	}

	return false
}

// SetAssociatedOrders gets a reference to the given []AssociatedOrder and assigns it to the AssociatedOrders field.
func (o *Order) SetAssociatedOrders(v []AssociatedOrder) {
	o.AssociatedOrders = v
}

// GetSalesChannel returns the SalesChannel field value
func (o *Order) GetSalesChannel() SalesChannel {
	if o == nil {
		var ret SalesChannel
		return ret
	}

	return o.SalesChannel
}

// GetSalesChannelOk returns a tuple with the SalesChannel field value
// and a boolean to check if the value has been set.
func (o *Order) GetSalesChannelOk() (*SalesChannel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SalesChannel, true
}

// SetSalesChannel sets field value
func (o *Order) SetSalesChannel(v SalesChannel) {
	o.SalesChannel = v
}

// GetBuyer returns the Buyer field value if set, zero value otherwise.
func (o *Order) GetBuyer() Buyer {
	if o == nil || IsNil(o.Buyer) {
		var ret Buyer
		return ret
	}
	return *o.Buyer
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetBuyerOk() (*Buyer, bool) {
	if o == nil || IsNil(o.Buyer) {
		return nil, false
	}
	return o.Buyer, true
}

// HasBuyer returns a boolean if a field has been set.
func (o *Order) HasBuyer() bool {
	if o != nil && !IsNil(o.Buyer) {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given Buyer and assigns it to the Buyer field.
func (o *Order) SetBuyer(v Buyer) {
	o.Buyer = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *Order) GetRecipient() Recipient {
	if o == nil || IsNil(o.Recipient) {
		var ret Recipient
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetRecipientOk() (*Recipient, bool) {
	if o == nil || IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *Order) HasRecipient() bool {
	if o != nil && !IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given Recipient and assigns it to the Recipient field.
func (o *Order) SetRecipient(v Recipient) {
	o.Recipient = &v
}

// GetProceeds returns the Proceeds field value if set, zero value otherwise.
func (o *Order) GetProceeds() OrderProceeds {
	if o == nil || IsNil(o.Proceeds) {
		var ret OrderProceeds
		return ret
	}
	return *o.Proceeds
}

// GetProceedsOk returns a tuple with the Proceeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetProceedsOk() (*OrderProceeds, bool) {
	if o == nil || IsNil(o.Proceeds) {
		return nil, false
	}
	return o.Proceeds, true
}

// HasProceeds returns a boolean if a field has been set.
func (o *Order) HasProceeds() bool {
	if o != nil && !IsNil(o.Proceeds) {
		return true
	}

	return false
}

// SetProceeds gets a reference to the given OrderProceeds and assigns it to the Proceeds field.
func (o *Order) SetProceeds(v OrderProceeds) {
	o.Proceeds = &v
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *Order) GetPayment() OrderPayment {
	if o == nil || IsNil(o.Payment) {
		var ret OrderPayment
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetPaymentOk() (*OrderPayment, bool) {
	if o == nil || IsNil(o.Payment) {
		return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *Order) HasPayment() bool {
	if o != nil && !IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given OrderPayment and assigns it to the Payment field.
func (o *Order) SetPayment(v OrderPayment) {
	o.Payment = &v
}

// GetTax returns the Tax field value if set, zero value otherwise.
func (o *Order) GetTax() OrderTax {
	if o == nil || IsNil(o.Tax) {
		var ret OrderTax
		return ret
	}
	return *o.Tax
}

// GetTaxOk returns a tuple with the Tax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetTaxOk() (*OrderTax, bool) {
	if o == nil || IsNil(o.Tax) {
		return nil, false
	}
	return o.Tax, true
}

// HasTax returns a boolean if a field has been set.
func (o *Order) HasTax() bool {
	if o != nil && !IsNil(o.Tax) {
		return true
	}

	return false
}

// SetTax gets a reference to the given OrderTax and assigns it to the Tax field.
func (o *Order) SetTax(v OrderTax) {
	o.Tax = &v
}

// GetFulfillment returns the Fulfillment field value if set, zero value otherwise.
func (o *Order) GetFulfillment() OrderFulfillment {
	if o == nil || IsNil(o.Fulfillment) {
		var ret OrderFulfillment
		return ret
	}
	return *o.Fulfillment
}

// GetFulfillmentOk returns a tuple with the Fulfillment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetFulfillmentOk() (*OrderFulfillment, bool) {
	if o == nil || IsNil(o.Fulfillment) {
		return nil, false
	}
	return o.Fulfillment, true
}

// HasFulfillment returns a boolean if a field has been set.
func (o *Order) HasFulfillment() bool {
	if o != nil && !IsNil(o.Fulfillment) {
		return true
	}

	return false
}

// SetFulfillment gets a reference to the given OrderFulfillment and assigns it to the Fulfillment field.
func (o *Order) SetFulfillment(v OrderFulfillment) {
	o.Fulfillment = &v
}

// GetOrderItems returns the OrderItems field value
func (o *Order) GetOrderItems() []OrderItem {
	if o == nil {
		var ret []OrderItem
		return ret
	}

	return o.OrderItems
}

// GetOrderItemsOk returns a tuple with the OrderItems field value
// and a boolean to check if the value has been set.
func (o *Order) GetOrderItemsOk() ([]OrderItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderItems, true
}

// SetOrderItems sets field value
func (o *Order) SetOrderItems(v []OrderItem) {
	o.OrderItems = v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *Order) GetPackages() []OrderPackage {
	if o == nil || IsNil(o.Packages) {
		var ret []OrderPackage
		return ret
	}
	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetPackagesOk() ([]OrderPackage, bool) {
	if o == nil || IsNil(o.Packages) {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *Order) HasPackages() bool {
	if o != nil && !IsNil(o.Packages) {
		return true
	}

	return false
}

// SetPackages gets a reference to the given []OrderPackage and assigns it to the Packages field.
func (o *Order) SetPackages(v []OrderPackage) {
	o.Packages = v
}

// GetFulfillmentOrders returns the FulfillmentOrders field value if set, zero value otherwise.
func (o *Order) GetFulfillmentOrders() []FulfillmentOrder {
	if o == nil || IsNil(o.FulfillmentOrders) {
		var ret []FulfillmentOrder
		return ret
	}
	return o.FulfillmentOrders
}

// GetFulfillmentOrdersOk returns a tuple with the FulfillmentOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetFulfillmentOrdersOk() ([]FulfillmentOrder, bool) {
	if o == nil || IsNil(o.FulfillmentOrders) {
		return nil, false
	}
	return o.FulfillmentOrders, true
}

// HasFulfillmentOrders returns a boolean if a field has been set.
func (o *Order) HasFulfillmentOrders() bool {
	if o != nil && !IsNil(o.FulfillmentOrders) {
		return true
	}

	return false
}

// SetFulfillmentOrders gets a reference to the given []FulfillmentOrder and assigns it to the FulfillmentOrders field.
func (o *Order) SetFulfillmentOrders(v []FulfillmentOrder) {
	o.FulfillmentOrders = v
}

func (o Order) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orderId"] = o.OrderId
	if !IsNil(o.OrderAliases) {
		toSerialize["orderAliases"] = o.OrderAliases
	}
	toSerialize["createdTime"] = o.CreatedTime
	toSerialize["lastUpdatedTime"] = o.LastUpdatedTime
	if !IsNil(o.Programs) {
		toSerialize["programs"] = o.Programs
	}
	if !IsNil(o.AssociatedOrders) {
		toSerialize["associatedOrders"] = o.AssociatedOrders
	}
	toSerialize["salesChannel"] = o.SalesChannel
	if !IsNil(o.Buyer) {
		toSerialize["buyer"] = o.Buyer
	}
	if !IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	if !IsNil(o.Proceeds) {
		toSerialize["proceeds"] = o.Proceeds
	}
	if !IsNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}
	if !IsNil(o.Tax) {
		toSerialize["tax"] = o.Tax
	}
	if !IsNil(o.Fulfillment) {
		toSerialize["fulfillment"] = o.Fulfillment
	}
	toSerialize["orderItems"] = o.OrderItems
	if !IsNil(o.Packages) {
		toSerialize["packages"] = o.Packages
	}
	if !IsNil(o.FulfillmentOrders) {
		toSerialize["fulfillmentOrders"] = o.FulfillmentOrders
	}
	return toSerialize, nil
}

type NullableOrder struct {
	value *Order
	isSet bool
}

func (v NullableOrder) Get() *Order {
	return v.value
}

func (v *NullableOrder) Set(val *Order) {
	v.value = val
	v.isSet = true
}

func (v NullableOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrder(val *Order) *NullableOrder {
	return &NullableOrder{value: val, isSet: true}
}

func (v NullableOrder) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
