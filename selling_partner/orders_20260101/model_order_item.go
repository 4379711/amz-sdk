package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderItem{}

// OrderItem Information about a single product within an order.
type OrderItem struct {
	// A unique identifier for this specific item within the order.
	OrderItemId string `json:"orderItemId"`
	// The number of units of this item that the customer ordered.
	QuantityOrdered int32        `json:"quantityOrdered"`
	Measurement     *Measurement `json:"measurement,omitempty"`
	// A list of order items associated with this item. For example, a value-add service purchased with the product.
	AssociatedOrderItems []AssociatedOrderItem `json:"associatedOrderItems,omitempty"`
	// Special programs that apply specifically to this item within the order.  **Possible values**: `TRANSPARENCY`, `SUBSCRIBE_AND_SAVE`
	Programs     []string          `json:"programs,omitempty"`
	Product      ItemProduct       `json:"product"`
	Proceeds     *ItemProceeds     `json:"proceeds,omitempty"`
	Expense      *ItemExpense      `json:"expense,omitempty"`
	Promotion    *ItemPromotion    `json:"promotion,omitempty"`
	Cancellation *ItemCancellation `json:"cancellation,omitempty"`
	Fulfillment  *ItemFulfillment  `json:"fulfillment,omitempty"`
	Tax          *ItemTax          `json:"tax,omitempty"`
}

// NewOrderItem instantiates a new OrderItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderItem(orderItemId string, quantityOrdered int32, product ItemProduct) *OrderItem {
	this := OrderItem{}
	this.OrderItemId = orderItemId
	this.QuantityOrdered = quantityOrdered
	this.Product = product
	return &this
}

// NewOrderItemWithDefaults instantiates a new OrderItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderItemWithDefaults() *OrderItem {
	this := OrderItem{}
	return &this
}

// GetOrderItemId returns the OrderItemId field value
func (o *OrderItem) GetOrderItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderItemId
}

// GetOrderItemIdOk returns a tuple with the OrderItemId field value
// and a boolean to check if the value has been set.
func (o *OrderItem) GetOrderItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderItemId, true
}

// SetOrderItemId sets field value
func (o *OrderItem) SetOrderItemId(v string) {
	o.OrderItemId = v
}

// GetQuantityOrdered returns the QuantityOrdered field value
func (o *OrderItem) GetQuantityOrdered() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QuantityOrdered
}

// GetQuantityOrderedOk returns a tuple with the QuantityOrdered field value
// and a boolean to check if the value has been set.
func (o *OrderItem) GetQuantityOrderedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuantityOrdered, true
}

// SetQuantityOrdered sets field value
func (o *OrderItem) SetQuantityOrdered(v int32) {
	o.QuantityOrdered = v
}

// GetMeasurement returns the Measurement field value if set, zero value otherwise.
func (o *OrderItem) GetMeasurement() Measurement {
	if o == nil || IsNil(o.Measurement) {
		var ret Measurement
		return ret
	}
	return *o.Measurement
}

// GetMeasurementOk returns a tuple with the Measurement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetMeasurementOk() (*Measurement, bool) {
	if o == nil || IsNil(o.Measurement) {
		return nil, false
	}
	return o.Measurement, true
}

// HasMeasurement returns a boolean if a field has been set.
func (o *OrderItem) HasMeasurement() bool {
	if o != nil && !IsNil(o.Measurement) {
		return true
	}

	return false
}

// SetMeasurement gets a reference to the given Measurement and assigns it to the Measurement field.
func (o *OrderItem) SetMeasurement(v Measurement) {
	o.Measurement = &v
}

// GetAssociatedOrderItems returns the AssociatedOrderItems field value if set, zero value otherwise.
func (o *OrderItem) GetAssociatedOrderItems() []AssociatedOrderItem {
	if o == nil || IsNil(o.AssociatedOrderItems) {
		var ret []AssociatedOrderItem
		return ret
	}
	return o.AssociatedOrderItems
}

// GetAssociatedOrderItemsOk returns a tuple with the AssociatedOrderItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetAssociatedOrderItemsOk() ([]AssociatedOrderItem, bool) {
	if o == nil || IsNil(o.AssociatedOrderItems) {
		return nil, false
	}
	return o.AssociatedOrderItems, true
}

// HasAssociatedOrderItems returns a boolean if a field has been set.
func (o *OrderItem) HasAssociatedOrderItems() bool {
	if o != nil && !IsNil(o.AssociatedOrderItems) {
		return true
	}

	return false
}

// SetAssociatedOrderItems gets a reference to the given []AssociatedOrderItem and assigns it to the AssociatedOrderItems field.
func (o *OrderItem) SetAssociatedOrderItems(v []AssociatedOrderItem) {
	o.AssociatedOrderItems = v
}

// GetPrograms returns the Programs field value if set, zero value otherwise.
func (o *OrderItem) GetPrograms() []string {
	if o == nil || IsNil(o.Programs) {
		var ret []string
		return ret
	}
	return o.Programs
}

// GetProgramsOk returns a tuple with the Programs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetProgramsOk() ([]string, bool) {
	if o == nil || IsNil(o.Programs) {
		return nil, false
	}
	return o.Programs, true
}

// HasPrograms returns a boolean if a field has been set.
func (o *OrderItem) HasPrograms() bool {
	if o != nil && !IsNil(o.Programs) {
		return true
	}

	return false
}

// SetPrograms gets a reference to the given []string and assigns it to the Programs field.
func (o *OrderItem) SetPrograms(v []string) {
	o.Programs = v
}

// GetProduct returns the Product field value
func (o *OrderItem) GetProduct() ItemProduct {
	if o == nil {
		var ret ItemProduct
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *OrderItem) GetProductOk() (*ItemProduct, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *OrderItem) SetProduct(v ItemProduct) {
	o.Product = v
}

// GetProceeds returns the Proceeds field value if set, zero value otherwise.
func (o *OrderItem) GetProceeds() ItemProceeds {
	if o == nil || IsNil(o.Proceeds) {
		var ret ItemProceeds
		return ret
	}
	return *o.Proceeds
}

// GetProceedsOk returns a tuple with the Proceeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetProceedsOk() (*ItemProceeds, bool) {
	if o == nil || IsNil(o.Proceeds) {
		return nil, false
	}
	return o.Proceeds, true
}

// HasProceeds returns a boolean if a field has been set.
func (o *OrderItem) HasProceeds() bool {
	if o != nil && !IsNil(o.Proceeds) {
		return true
	}

	return false
}

// SetProceeds gets a reference to the given ItemProceeds and assigns it to the Proceeds field.
func (o *OrderItem) SetProceeds(v ItemProceeds) {
	o.Proceeds = &v
}

// GetExpense returns the Expense field value if set, zero value otherwise.
func (o *OrderItem) GetExpense() ItemExpense {
	if o == nil || IsNil(o.Expense) {
		var ret ItemExpense
		return ret
	}
	return *o.Expense
}

// GetExpenseOk returns a tuple with the Expense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetExpenseOk() (*ItemExpense, bool) {
	if o == nil || IsNil(o.Expense) {
		return nil, false
	}
	return o.Expense, true
}

// HasExpense returns a boolean if a field has been set.
func (o *OrderItem) HasExpense() bool {
	if o != nil && !IsNil(o.Expense) {
		return true
	}

	return false
}

// SetExpense gets a reference to the given ItemExpense and assigns it to the Expense field.
func (o *OrderItem) SetExpense(v ItemExpense) {
	o.Expense = &v
}

// GetPromotion returns the Promotion field value if set, zero value otherwise.
func (o *OrderItem) GetPromotion() ItemPromotion {
	if o == nil || IsNil(o.Promotion) {
		var ret ItemPromotion
		return ret
	}
	return *o.Promotion
}

// GetPromotionOk returns a tuple with the Promotion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetPromotionOk() (*ItemPromotion, bool) {
	if o == nil || IsNil(o.Promotion) {
		return nil, false
	}
	return o.Promotion, true
}

// HasPromotion returns a boolean if a field has been set.
func (o *OrderItem) HasPromotion() bool {
	if o != nil && !IsNil(o.Promotion) {
		return true
	}

	return false
}

// SetPromotion gets a reference to the given ItemPromotion and assigns it to the Promotion field.
func (o *OrderItem) SetPromotion(v ItemPromotion) {
	o.Promotion = &v
}

// GetCancellation returns the Cancellation field value if set, zero value otherwise.
func (o *OrderItem) GetCancellation() ItemCancellation {
	if o == nil || IsNil(o.Cancellation) {
		var ret ItemCancellation
		return ret
	}
	return *o.Cancellation
}

// GetCancellationOk returns a tuple with the Cancellation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetCancellationOk() (*ItemCancellation, bool) {
	if o == nil || IsNil(o.Cancellation) {
		return nil, false
	}
	return o.Cancellation, true
}

// HasCancellation returns a boolean if a field has been set.
func (o *OrderItem) HasCancellation() bool {
	if o != nil && !IsNil(o.Cancellation) {
		return true
	}

	return false
}

// SetCancellation gets a reference to the given ItemCancellation and assigns it to the Cancellation field.
func (o *OrderItem) SetCancellation(v ItemCancellation) {
	o.Cancellation = &v
}

// GetFulfillment returns the Fulfillment field value if set, zero value otherwise.
func (o *OrderItem) GetFulfillment() ItemFulfillment {
	if o == nil || IsNil(o.Fulfillment) {
		var ret ItemFulfillment
		return ret
	}
	return *o.Fulfillment
}

// GetFulfillmentOk returns a tuple with the Fulfillment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetFulfillmentOk() (*ItemFulfillment, bool) {
	if o == nil || IsNil(o.Fulfillment) {
		return nil, false
	}
	return o.Fulfillment, true
}

// HasFulfillment returns a boolean if a field has been set.
func (o *OrderItem) HasFulfillment() bool {
	if o != nil && !IsNil(o.Fulfillment) {
		return true
	}

	return false
}

// SetFulfillment gets a reference to the given ItemFulfillment and assigns it to the Fulfillment field.
func (o *OrderItem) SetFulfillment(v ItemFulfillment) {
	o.Fulfillment = &v
}

// GetTax returns the Tax field value if set, zero value otherwise.
func (o *OrderItem) GetTax() ItemTax {
	if o == nil || IsNil(o.Tax) {
		var ret ItemTax
		return ret
	}
	return *o.Tax
}

// GetTaxOk returns a tuple with the Tax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetTaxOk() (*ItemTax, bool) {
	if o == nil || IsNil(o.Tax) {
		return nil, false
	}
	return o.Tax, true
}

// HasTax returns a boolean if a field has been set.
func (o *OrderItem) HasTax() bool {
	if o != nil && !IsNil(o.Tax) {
		return true
	}

	return false
}

// SetTax gets a reference to the given ItemTax and assigns it to the Tax field.
func (o *OrderItem) SetTax(v ItemTax) {
	o.Tax = &v
}

func (o OrderItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orderItemId"] = o.OrderItemId
	toSerialize["quantityOrdered"] = o.QuantityOrdered
	if !IsNil(o.Measurement) {
		toSerialize["measurement"] = o.Measurement
	}
	if !IsNil(o.AssociatedOrderItems) {
		toSerialize["associatedOrderItems"] = o.AssociatedOrderItems
	}
	if !IsNil(o.Programs) {
		toSerialize["programs"] = o.Programs
	}
	toSerialize["product"] = o.Product
	if !IsNil(o.Proceeds) {
		toSerialize["proceeds"] = o.Proceeds
	}
	if !IsNil(o.Expense) {
		toSerialize["expense"] = o.Expense
	}
	if !IsNil(o.Promotion) {
		toSerialize["promotion"] = o.Promotion
	}
	if !IsNil(o.Cancellation) {
		toSerialize["cancellation"] = o.Cancellation
	}
	if !IsNil(o.Fulfillment) {
		toSerialize["fulfillment"] = o.Fulfillment
	}
	if !IsNil(o.Tax) {
		toSerialize["tax"] = o.Tax
	}
	return toSerialize, nil
}

type NullableOrderItem struct {
	value *OrderItem
	isSet bool
}

func (v NullableOrderItem) Get() *OrderItem {
	return v.value
}

func (v *NullableOrderItem) Set(val *OrderItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderItem(val *OrderItem) *NullableOrderItem {
	return &NullableOrderItem{value: val, isSet: true}
}

func (v NullableOrderItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
