package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderTaxInvoicing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderTaxInvoicing{}

// OrderTaxInvoicing Tax invoicing information for the order.
type OrderTaxInvoicing struct {
	// The buyer's invoicing preference, which indicates whether the seller should issue an individual or a business invoice to the buyer.    **Note**: This attribute is only available in the Turkey marketplace.   **Possible values**: - `INDIVIDUAL` (Issues an individual invoice to the buyer) - `BUSINESS` (Issues a business invoice to the buyer)
	BuyerInvoicePreference *string `json:"buyerInvoicePreference,omitempty"`
	// The status of the invoice. Only available for Easy Ship orders and orders in the Brazil marketplace.  **Possible values**: - `NOT_REQUIRED` (The order does not require an electronic invoice to be uploaded) - `NOT_FOUND` (The order requires an electronic invoice but it is not uploaded) - `PROCESSING` (The required electronic invoice was uploaded and is processing) - `ERRORED` (The uploaded electronic invoice was not accepted) - `ACCEPTED` (The uploaded electronic invoice was accepted)
	InvoiceStatus *string `json:"invoiceStatus,omitempty"`
}

// NewOrderTaxInvoicing instantiates a new OrderTaxInvoicing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTaxInvoicing() *OrderTaxInvoicing {
	this := OrderTaxInvoicing{}
	return &this
}

// NewOrderTaxInvoicingWithDefaults instantiates a new OrderTaxInvoicing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTaxInvoicingWithDefaults() *OrderTaxInvoicing {
	this := OrderTaxInvoicing{}
	return &this
}

// GetBuyerInvoicePreference returns the BuyerInvoicePreference field value if set, zero value otherwise.
func (o *OrderTaxInvoicing) GetBuyerInvoicePreference() string {
	if o == nil || IsNil(o.BuyerInvoicePreference) {
		var ret string
		return ret
	}
	return *o.BuyerInvoicePreference
}

// GetBuyerInvoicePreferenceOk returns a tuple with the BuyerInvoicePreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTaxInvoicing) GetBuyerInvoicePreferenceOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerInvoicePreference) {
		return nil, false
	}
	return o.BuyerInvoicePreference, true
}

// HasBuyerInvoicePreference returns a boolean if a field has been set.
func (o *OrderTaxInvoicing) HasBuyerInvoicePreference() bool {
	if o != nil && !IsNil(o.BuyerInvoicePreference) {
		return true
	}

	return false
}

// SetBuyerInvoicePreference gets a reference to the given string and assigns it to the BuyerInvoicePreference field.
func (o *OrderTaxInvoicing) SetBuyerInvoicePreference(v string) {
	o.BuyerInvoicePreference = &v
}

// GetInvoiceStatus returns the InvoiceStatus field value if set, zero value otherwise.
func (o *OrderTaxInvoicing) GetInvoiceStatus() string {
	if o == nil || IsNil(o.InvoiceStatus) {
		var ret string
		return ret
	}
	return *o.InvoiceStatus
}

// GetInvoiceStatusOk returns a tuple with the InvoiceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTaxInvoicing) GetInvoiceStatusOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceStatus) {
		return nil, false
	}
	return o.InvoiceStatus, true
}

// HasInvoiceStatus returns a boolean if a field has been set.
func (o *OrderTaxInvoicing) HasInvoiceStatus() bool {
	if o != nil && !IsNil(o.InvoiceStatus) {
		return true
	}

	return false
}

// SetInvoiceStatus gets a reference to the given string and assigns it to the InvoiceStatus field.
func (o *OrderTaxInvoicing) SetInvoiceStatus(v string) {
	o.InvoiceStatus = &v
}

func (o OrderTaxInvoicing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuyerInvoicePreference) {
		toSerialize["buyerInvoicePreference"] = o.BuyerInvoicePreference
	}
	if !IsNil(o.InvoiceStatus) {
		toSerialize["invoiceStatus"] = o.InvoiceStatus
	}
	return toSerialize, nil
}

type NullableOrderTaxInvoicing struct {
	value *OrderTaxInvoicing
	isSet bool
}

func (v NullableOrderTaxInvoicing) Get() *OrderTaxInvoicing {
	return v.value
}

func (v *NullableOrderTaxInvoicing) Set(val *OrderTaxInvoicing) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTaxInvoicing) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTaxInvoicing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTaxInvoicing(val *OrderTaxInvoicing) *NullableOrderTaxInvoicing {
	return &NullableOrderTaxInvoicing{value: val, isSet: true}
}

func (v NullableOrderTaxInvoicing) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderTaxInvoicing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
