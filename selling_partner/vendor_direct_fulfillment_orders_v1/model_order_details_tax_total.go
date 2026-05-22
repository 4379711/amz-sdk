package vendor_direct_fulfillment_orders_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderDetailsTaxTotal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailsTaxTotal{}

// OrderDetailsTaxTotal The total Tax object within shipment that relates to the order.
type OrderDetailsTaxTotal struct {
	// A list of tax line items.
	TaxLineItem []TaxDetails `json:"taxLineItem,omitempty"`
}

// NewOrderDetailsTaxTotal instantiates a new OrderDetailsTaxTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailsTaxTotal() *OrderDetailsTaxTotal {
	this := OrderDetailsTaxTotal{}
	return &this
}

// NewOrderDetailsTaxTotalWithDefaults instantiates a new OrderDetailsTaxTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailsTaxTotalWithDefaults() *OrderDetailsTaxTotal {
	this := OrderDetailsTaxTotal{}
	return &this
}

// GetTaxLineItem returns the TaxLineItem field value if set, zero value otherwise.
func (o *OrderDetailsTaxTotal) GetTaxLineItem() []TaxDetails {
	if o == nil || IsNil(o.TaxLineItem) {
		var ret []TaxDetails
		return ret
	}
	return o.TaxLineItem
}

// GetTaxLineItemOk returns a tuple with the TaxLineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailsTaxTotal) GetTaxLineItemOk() ([]TaxDetails, bool) {
	if o == nil || IsNil(o.TaxLineItem) {
		return nil, false
	}
	return o.TaxLineItem, true
}

// HasTaxLineItem returns a boolean if a field has been set.
func (o *OrderDetailsTaxTotal) HasTaxLineItem() bool {
	if o != nil && !IsNil(o.TaxLineItem) {
		return true
	}

	return false
}

// SetTaxLineItem gets a reference to the given []TaxDetails and assigns it to the TaxLineItem field.
func (o *OrderDetailsTaxTotal) SetTaxLineItem(v []TaxDetails) {
	o.TaxLineItem = v
}

func (o OrderDetailsTaxTotal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaxLineItem) {
		toSerialize["taxLineItem"] = o.TaxLineItem
	}
	return toSerialize, nil
}

type NullableOrderDetailsTaxTotal struct {
	value *OrderDetailsTaxTotal
	isSet bool
}

func (v NullableOrderDetailsTaxTotal) Get() *OrderDetailsTaxTotal {
	return v.value
}

func (v *NullableOrderDetailsTaxTotal) Set(val *OrderDetailsTaxTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailsTaxTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailsTaxTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailsTaxTotal(val *OrderDetailsTaxTotal) *NullableOrderDetailsTaxTotal {
	return &NullableOrderDetailsTaxTotal{value: val, isSet: true}
}

func (v NullableOrderDetailsTaxTotal) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderDetailsTaxTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
