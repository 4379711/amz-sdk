package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the InvoiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceRequest{}

// InvoiceRequest The request schema for the `sendInvoice` operation.
type InvoiceRequest struct {
	// Attachments to include in the message to the buyer.
	Attachments []Attachment `json:"attachments,omitempty"`
}

// NewInvoiceRequest instantiates a new InvoiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceRequest() *InvoiceRequest {
	this := InvoiceRequest{}
	return &this
}

// NewInvoiceRequestWithDefaults instantiates a new InvoiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceRequestWithDefaults() *InvoiceRequest {
	this := InvoiceRequest{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *InvoiceRequest) GetAttachments() []Attachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *InvoiceRequest) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *InvoiceRequest) SetAttachments(v []Attachment) {
	o.Attachments = v
}

func (o InvoiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullableInvoiceRequest struct {
	value *InvoiceRequest
	isSet bool
}

func (v NullableInvoiceRequest) Get() *InvoiceRequest {
	return v.value
}

func (v *NullableInvoiceRequest) Set(val *InvoiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceRequest(val *InvoiceRequest) *NullableInvoiceRequest {
	return &NullableInvoiceRequest{value: val, isSet: true}
}

func (v NullableInvoiceRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInvoiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
