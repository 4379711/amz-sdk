package vendor_direct_fulfillment_shipping_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CustomerInvoiceList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerInvoiceList{}

// CustomerInvoiceList Represents a list of customer invoices, potentially paginated.
type CustomerInvoiceList struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	// Represents a customer invoice within the CustomerInvoiceList.
	CustomerInvoices []CustomerInvoice `json:"customerInvoices,omitempty"`
}

// NewCustomerInvoiceList instantiates a new CustomerInvoiceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerInvoiceList() *CustomerInvoiceList {
	this := CustomerInvoiceList{}
	return &this
}

// NewCustomerInvoiceListWithDefaults instantiates a new CustomerInvoiceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerInvoiceListWithDefaults() *CustomerInvoiceList {
	this := CustomerInvoiceList{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *CustomerInvoiceList) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerInvoiceList) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *CustomerInvoiceList) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *CustomerInvoiceList) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetCustomerInvoices returns the CustomerInvoices field value if set, zero value otherwise.
func (o *CustomerInvoiceList) GetCustomerInvoices() []CustomerInvoice {
	if o == nil || IsNil(o.CustomerInvoices) {
		var ret []CustomerInvoice
		return ret
	}
	return o.CustomerInvoices
}

// GetCustomerInvoicesOk returns a tuple with the CustomerInvoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerInvoiceList) GetCustomerInvoicesOk() ([]CustomerInvoice, bool) {
	if o == nil || IsNil(o.CustomerInvoices) {
		return nil, false
	}
	return o.CustomerInvoices, true
}

// HasCustomerInvoices returns a boolean if a field has been set.
func (o *CustomerInvoiceList) HasCustomerInvoices() bool {
	if o != nil && !IsNil(o.CustomerInvoices) {
		return true
	}

	return false
}

// SetCustomerInvoices gets a reference to the given []CustomerInvoice and assigns it to the CustomerInvoices field.
func (o *CustomerInvoiceList) SetCustomerInvoices(v []CustomerInvoice) {
	o.CustomerInvoices = v
}

func (o CustomerInvoiceList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.CustomerInvoices) {
		toSerialize["customerInvoices"] = o.CustomerInvoices
	}
	return toSerialize, nil
}

type NullableCustomerInvoiceList struct {
	value *CustomerInvoiceList
	isSet bool
}

func (v NullableCustomerInvoiceList) Get() *CustomerInvoiceList {
	return v.value
}

func (v *NullableCustomerInvoiceList) Set(val *CustomerInvoiceList) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerInvoiceList) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerInvoiceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerInvoiceList(val *CustomerInvoiceList) *NullableCustomerInvoiceList {
	return &NullableCustomerInvoiceList{value: val, isSet: true}
}

func (v NullableCustomerInvoiceList) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCustomerInvoiceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
