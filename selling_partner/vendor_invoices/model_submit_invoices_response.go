package vendor_invoices

import (
	"github.com/bytedance/sonic"
)

// checks if the SubmitInvoicesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitInvoicesResponse{}

// SubmitInvoicesResponse The response schema for the submitInvoices operation.
type SubmitInvoicesResponse struct {
	Payload *TransactionId `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewSubmitInvoicesResponse instantiates a new SubmitInvoicesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitInvoicesResponse() *SubmitInvoicesResponse {
	this := SubmitInvoicesResponse{}
	return &this
}

// NewSubmitInvoicesResponseWithDefaults instantiates a new SubmitInvoicesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitInvoicesResponseWithDefaults() *SubmitInvoicesResponse {
	this := SubmitInvoicesResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *SubmitInvoicesResponse) GetPayload() TransactionId {
	if o == nil || IsNil(o.Payload) {
		var ret TransactionId
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitInvoicesResponse) GetPayloadOk() (*TransactionId, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *SubmitInvoicesResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given TransactionId and assigns it to the Payload field.
func (o *SubmitInvoicesResponse) SetPayload(v TransactionId) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SubmitInvoicesResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitInvoicesResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SubmitInvoicesResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *SubmitInvoicesResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o SubmitInvoicesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSubmitInvoicesResponse struct {
	value *SubmitInvoicesResponse
	isSet bool
}

func (v NullableSubmitInvoicesResponse) Get() *SubmitInvoicesResponse {
	return v.value
}

func (v *NullableSubmitInvoicesResponse) Set(val *SubmitInvoicesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitInvoicesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitInvoicesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitInvoicesResponse(val *SubmitInvoicesResponse) *NullableSubmitInvoicesResponse {
	return &NullableSubmitInvoicesResponse{value: val, isSet: true}
}

func (v NullableSubmitInvoicesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubmitInvoicesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
