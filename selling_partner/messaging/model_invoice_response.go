package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the InvoiceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceResponse{}

// InvoiceResponse The response schema for the sendInvoice operation.
type InvoiceResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewInvoiceResponse instantiates a new InvoiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceResponse() *InvoiceResponse {
	this := InvoiceResponse{}
	return &this
}

// NewInvoiceResponseWithDefaults instantiates a new InvoiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceResponseWithDefaults() *InvoiceResponse {
	this := InvoiceResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *InvoiceResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *InvoiceResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *InvoiceResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o InvoiceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableInvoiceResponse struct {
	value *InvoiceResponse
	isSet bool
}

func (v NullableInvoiceResponse) Get() *InvoiceResponse {
	return v.value
}

func (v *NullableInvoiceResponse) Set(val *InvoiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceResponse(val *InvoiceResponse) *NullableInvoiceResponse {
	return &NullableInvoiceResponse{value: val, isSet: true}
}

func (v NullableInvoiceResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInvoiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
