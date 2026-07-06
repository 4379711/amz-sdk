package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderProceeds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderProceeds{}

// OrderProceeds The money that the seller receives from the sale of the order.
type OrderProceeds struct {
	GrandTotal *Money `json:"grandTotal,omitempty"`
	// Categorized proceeds for the order. Proceed categories are either aggregated across all order items (such as `ITEM`, `SHIPPING`, and `TAX`) or applied at the order level (such as `DELIVERY_TIP`).
	Breakdowns []OrderProceedsBreakdown `json:"breakdowns,omitempty"`
}

// NewOrderProceeds instantiates a new OrderProceeds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderProceeds() *OrderProceeds {
	this := OrderProceeds{}
	return &this
}

// NewOrderProceedsWithDefaults instantiates a new OrderProceeds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderProceedsWithDefaults() *OrderProceeds {
	this := OrderProceeds{}
	return &this
}

// GetGrandTotal returns the GrandTotal field value if set, zero value otherwise.
func (o *OrderProceeds) GetGrandTotal() Money {
	if o == nil || IsNil(o.GrandTotal) {
		var ret Money
		return ret
	}
	return *o.GrandTotal
}

// GetGrandTotalOk returns a tuple with the GrandTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderProceeds) GetGrandTotalOk() (*Money, bool) {
	if o == nil || IsNil(o.GrandTotal) {
		return nil, false
	}
	return o.GrandTotal, true
}

// HasGrandTotal returns a boolean if a field has been set.
func (o *OrderProceeds) HasGrandTotal() bool {
	if o != nil && !IsNil(o.GrandTotal) {
		return true
	}

	return false
}

// SetGrandTotal gets a reference to the given Money and assigns it to the GrandTotal field.
func (o *OrderProceeds) SetGrandTotal(v Money) {
	o.GrandTotal = &v
}

// GetBreakdowns returns the Breakdowns field value if set, zero value otherwise.
func (o *OrderProceeds) GetBreakdowns() []OrderProceedsBreakdown {
	if o == nil || IsNil(o.Breakdowns) {
		var ret []OrderProceedsBreakdown
		return ret
	}
	return o.Breakdowns
}

// GetBreakdownsOk returns a tuple with the Breakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderProceeds) GetBreakdownsOk() ([]OrderProceedsBreakdown, bool) {
	if o == nil || IsNil(o.Breakdowns) {
		return nil, false
	}
	return o.Breakdowns, true
}

// HasBreakdowns returns a boolean if a field has been set.
func (o *OrderProceeds) HasBreakdowns() bool {
	if o != nil && !IsNil(o.Breakdowns) {
		return true
	}

	return false
}

// SetBreakdowns gets a reference to the given []OrderProceedsBreakdown and assigns it to the Breakdowns field.
func (o *OrderProceeds) SetBreakdowns(v []OrderProceedsBreakdown) {
	o.Breakdowns = v
}

func (o OrderProceeds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GrandTotal) {
		toSerialize["grandTotal"] = o.GrandTotal
	}
	if !IsNil(o.Breakdowns) {
		toSerialize["breakdowns"] = o.Breakdowns
	}
	return toSerialize, nil
}

type NullableOrderProceeds struct {
	value *OrderProceeds
	isSet bool
}

func (v NullableOrderProceeds) Get() *OrderProceeds {
	return v.value
}

func (v *NullableOrderProceeds) Set(val *OrderProceeds) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderProceeds) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderProceeds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderProceeds(val *OrderProceeds) *NullableOrderProceeds {
	return &NullableOrderProceeds{value: val, isSet: true}
}

func (v NullableOrderProceeds) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderProceeds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
