package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderTax type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderTax{}

// OrderTax Tax information about the order.
type OrderTax struct {
	// A list of tax registrations associated with the order.
	TaxRegistrations []OrderTaxRegistration `json:"taxRegistrations,omitempty"`
	TaxInvoicing     *OrderTaxInvoicing     `json:"taxInvoicing,omitempty"`
}

// NewOrderTax instantiates a new OrderTax object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTax() *OrderTax {
	this := OrderTax{}
	return &this
}

// NewOrderTaxWithDefaults instantiates a new OrderTax object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTaxWithDefaults() *OrderTax {
	this := OrderTax{}
	return &this
}

// GetTaxRegistrations returns the TaxRegistrations field value if set, zero value otherwise.
func (o *OrderTax) GetTaxRegistrations() []OrderTaxRegistration {
	if o == nil || IsNil(o.TaxRegistrations) {
		var ret []OrderTaxRegistration
		return ret
	}
	return o.TaxRegistrations
}

// GetTaxRegistrationsOk returns a tuple with the TaxRegistrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTax) GetTaxRegistrationsOk() ([]OrderTaxRegistration, bool) {
	if o == nil || IsNil(o.TaxRegistrations) {
		return nil, false
	}
	return o.TaxRegistrations, true
}

// HasTaxRegistrations returns a boolean if a field has been set.
func (o *OrderTax) HasTaxRegistrations() bool {
	if o != nil && !IsNil(o.TaxRegistrations) {
		return true
	}

	return false
}

// SetTaxRegistrations gets a reference to the given []OrderTaxRegistration and assigns it to the TaxRegistrations field.
func (o *OrderTax) SetTaxRegistrations(v []OrderTaxRegistration) {
	o.TaxRegistrations = v
}

// GetTaxInvoicing returns the TaxInvoicing field value if set, zero value otherwise.
func (o *OrderTax) GetTaxInvoicing() OrderTaxInvoicing {
	if o == nil || IsNil(o.TaxInvoicing) {
		var ret OrderTaxInvoicing
		return ret
	}
	return *o.TaxInvoicing
}

// GetTaxInvoicingOk returns a tuple with the TaxInvoicing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTax) GetTaxInvoicingOk() (*OrderTaxInvoicing, bool) {
	if o == nil || IsNil(o.TaxInvoicing) {
		return nil, false
	}
	return o.TaxInvoicing, true
}

// HasTaxInvoicing returns a boolean if a field has been set.
func (o *OrderTax) HasTaxInvoicing() bool {
	if o != nil && !IsNil(o.TaxInvoicing) {
		return true
	}

	return false
}

// SetTaxInvoicing gets a reference to the given OrderTaxInvoicing and assigns it to the TaxInvoicing field.
func (o *OrderTax) SetTaxInvoicing(v OrderTaxInvoicing) {
	o.TaxInvoicing = &v
}

func (o OrderTax) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaxRegistrations) {
		toSerialize["taxRegistrations"] = o.TaxRegistrations
	}
	if !IsNil(o.TaxInvoicing) {
		toSerialize["taxInvoicing"] = o.TaxInvoicing
	}
	return toSerialize, nil
}

type NullableOrderTax struct {
	value *OrderTax
	isSet bool
}

func (v NullableOrderTax) Get() *OrderTax {
	return v.value
}

func (v *NullableOrderTax) Set(val *OrderTax) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTax) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTax) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTax(val *OrderTax) *NullableOrderTax {
	return &NullableOrderTax{value: val, isSet: true}
}

func (v NullableOrderTax) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderTax) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
