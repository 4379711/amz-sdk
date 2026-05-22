package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateLegalDisclosureRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLegalDisclosureRequest{}

// CreateLegalDisclosureRequest The request schema for the createLegalDisclosure operation.
type CreateLegalDisclosureRequest struct {
	// Attachments to include in the message to the buyer. If any text is included in the attachment, the text must be written in the buyer's language of preference, which can be retrieved from the GetAttributes operation.
	Attachments []Attachment `json:"attachments,omitempty"`
}

// NewCreateLegalDisclosureRequest instantiates a new CreateLegalDisclosureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLegalDisclosureRequest() *CreateLegalDisclosureRequest {
	this := CreateLegalDisclosureRequest{}
	return &this
}

// NewCreateLegalDisclosureRequestWithDefaults instantiates a new CreateLegalDisclosureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLegalDisclosureRequestWithDefaults() *CreateLegalDisclosureRequest {
	this := CreateLegalDisclosureRequest{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *CreateLegalDisclosureRequest) GetAttachments() []Attachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLegalDisclosureRequest) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *CreateLegalDisclosureRequest) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *CreateLegalDisclosureRequest) SetAttachments(v []Attachment) {
	o.Attachments = v
}

func (o CreateLegalDisclosureRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullableCreateLegalDisclosureRequest struct {
	value *CreateLegalDisclosureRequest
	isSet bool
}

func (v NullableCreateLegalDisclosureRequest) Get() *CreateLegalDisclosureRequest {
	return v.value
}

func (v *NullableCreateLegalDisclosureRequest) Set(val *CreateLegalDisclosureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLegalDisclosureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLegalDisclosureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLegalDisclosureRequest(val *CreateLegalDisclosureRequest) *NullableCreateLegalDisclosureRequest {
	return &NullableCreateLegalDisclosureRequest{value: val, isSet: true}
}

func (v NullableCreateLegalDisclosureRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateLegalDisclosureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
