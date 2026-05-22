package shipment_invoicing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the SubmitInvoiceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitInvoiceResponse{}

// SubmitInvoiceResponse The response schema for the submitInvoice operation.
type SubmitInvoiceResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewSubmitInvoiceResponse instantiates a new SubmitInvoiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitInvoiceResponse() *SubmitInvoiceResponse {
	this := SubmitInvoiceResponse{}
	return &this
}

// NewSubmitInvoiceResponseWithDefaults instantiates a new SubmitInvoiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitInvoiceResponseWithDefaults() *SubmitInvoiceResponse {
	this := SubmitInvoiceResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SubmitInvoiceResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitInvoiceResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SubmitInvoiceResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *SubmitInvoiceResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o SubmitInvoiceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSubmitInvoiceResponse struct {
	value *SubmitInvoiceResponse
	isSet bool
}

func (v NullableSubmitInvoiceResponse) Get() *SubmitInvoiceResponse {
	return v.value
}

func (v *NullableSubmitInvoiceResponse) Set(val *SubmitInvoiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitInvoiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitInvoiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitInvoiceResponse(val *SubmitInvoiceResponse) *NullableSubmitInvoiceResponse {
	return &NullableSubmitInvoiceResponse{value: val, isSet: true}
}

func (v NullableSubmitInvoiceResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubmitInvoiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
