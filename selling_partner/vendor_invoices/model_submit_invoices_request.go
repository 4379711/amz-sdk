package vendor_invoices

import (
	"github.com/bytedance/sonic"
)

// checks if the SubmitInvoicesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitInvoicesRequest{}

// SubmitInvoicesRequest The request schema for the submitInvoices operation.
type SubmitInvoicesRequest struct {
	// An array of Invoice objects representing the invoices or credit notes to be submitted.
	Invoices []Invoice `json:"invoices,omitempty"`
}

// NewSubmitInvoicesRequest instantiates a new SubmitInvoicesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitInvoicesRequest() *SubmitInvoicesRequest {
	this := SubmitInvoicesRequest{}
	return &this
}

// NewSubmitInvoicesRequestWithDefaults instantiates a new SubmitInvoicesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitInvoicesRequestWithDefaults() *SubmitInvoicesRequest {
	this := SubmitInvoicesRequest{}
	return &this
}

// GetInvoices returns the Invoices field value if set, zero value otherwise.
func (o *SubmitInvoicesRequest) GetInvoices() []Invoice {
	if o == nil || IsNil(o.Invoices) {
		var ret []Invoice
		return ret
	}
	return o.Invoices
}

// GetInvoicesOk returns a tuple with the Invoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitInvoicesRequest) GetInvoicesOk() ([]Invoice, bool) {
	if o == nil || IsNil(o.Invoices) {
		return nil, false
	}
	return o.Invoices, true
}

// HasInvoices returns a boolean if a field has been set.
func (o *SubmitInvoicesRequest) HasInvoices() bool {
	if o != nil && !IsNil(o.Invoices) {
		return true
	}

	return false
}

// SetInvoices gets a reference to the given []Invoice and assigns it to the Invoices field.
func (o *SubmitInvoicesRequest) SetInvoices(v []Invoice) {
	o.Invoices = v
}

func (o SubmitInvoicesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Invoices) {
		toSerialize["invoices"] = o.Invoices
	}
	return toSerialize, nil
}

type NullableSubmitInvoicesRequest struct {
	value *SubmitInvoicesRequest
	isSet bool
}

func (v NullableSubmitInvoicesRequest) Get() *SubmitInvoicesRequest {
	return v.value
}

func (v *NullableSubmitInvoicesRequest) Set(val *SubmitInvoicesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitInvoicesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitInvoicesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitInvoicesRequest(val *SubmitInvoicesRequest) *NullableSubmitInvoicesRequest {
	return &NullableSubmitInvoicesRequest{value: val, isSet: true}
}

func (v NullableSubmitInvoicesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubmitInvoicesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
