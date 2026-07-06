package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderProceedsBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderProceedsBreakdown{}

// OrderProceedsBreakdown An entry detailing proceeds information.
type OrderProceedsBreakdown struct {
	// The proceeds category.   **Possible values**: `ITEM`, `SHIPPING`, `GIFT_WRAP`, `COD_FEE`, `TAX`, `DISCOUNT`, `DELIVERY_TIP`, `OTHER`. **Note:** `DELIVERY_TIP` is charged separately and not attributed to a specific item. The remaining categories are aggregated across all order items.
	Type string `json:"type"`
	// The processing status of the charge. Only present for categories processed separately after checkout, such as `DELIVERY_TIP`.  **Possible values**: `PENDING`, `FINALIZED`.
	Status   *string `json:"status,omitempty"`
	Subtotal Money   `json:"subtotal"`
}

// NewOrderProceedsBreakdown instantiates a new OrderProceedsBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderProceedsBreakdown(type_ string, subtotal Money) *OrderProceedsBreakdown {
	this := OrderProceedsBreakdown{}
	this.Type = type_
	this.Subtotal = subtotal
	return &this
}

// NewOrderProceedsBreakdownWithDefaults instantiates a new OrderProceedsBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderProceedsBreakdownWithDefaults() *OrderProceedsBreakdown {
	this := OrderProceedsBreakdown{}
	return &this
}

// GetType returns the Type field value
func (o *OrderProceedsBreakdown) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrderProceedsBreakdown) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderProceedsBreakdown) SetType(v string) {
	o.Type = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderProceedsBreakdown) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderProceedsBreakdown) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderProceedsBreakdown) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrderProceedsBreakdown) SetStatus(v string) {
	o.Status = &v
}

// GetSubtotal returns the Subtotal field value
func (o *OrderProceedsBreakdown) GetSubtotal() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.Subtotal
}

// GetSubtotalOk returns a tuple with the Subtotal field value
// and a boolean to check if the value has been set.
func (o *OrderProceedsBreakdown) GetSubtotalOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtotal, true
}

// SetSubtotal sets field value
func (o *OrderProceedsBreakdown) SetSubtotal(v Money) {
	o.Subtotal = v
}

func (o OrderProceedsBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["subtotal"] = o.Subtotal
	return toSerialize, nil
}

type NullableOrderProceedsBreakdown struct {
	value *OrderProceedsBreakdown
	isSet bool
}

func (v NullableOrderProceedsBreakdown) Get() *OrderProceedsBreakdown {
	return v.value
}

func (v *NullableOrderProceedsBreakdown) Set(val *OrderProceedsBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderProceedsBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderProceedsBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderProceedsBreakdown(val *OrderProceedsBreakdown) *NullableOrderProceedsBreakdown {
	return &NullableOrderProceedsBreakdown{value: val, isSet: true}
}

func (v NullableOrderProceedsBreakdown) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderProceedsBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
