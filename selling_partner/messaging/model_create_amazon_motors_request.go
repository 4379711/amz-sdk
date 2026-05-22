package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateAmazonMotorsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAmazonMotorsRequest{}

// CreateAmazonMotorsRequest The request schema for the createAmazonMotors operation.
type CreateAmazonMotorsRequest struct {
	// Attachments to include in the message to the buyer. If any text is included in the attachment, the text must be written in the buyer's language of preference, which can be retrieved from the GetAttributes operation.
	Attachments []Attachment `json:"attachments,omitempty"`
}

// NewCreateAmazonMotorsRequest instantiates a new CreateAmazonMotorsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAmazonMotorsRequest() *CreateAmazonMotorsRequest {
	this := CreateAmazonMotorsRequest{}
	return &this
}

// NewCreateAmazonMotorsRequestWithDefaults instantiates a new CreateAmazonMotorsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAmazonMotorsRequestWithDefaults() *CreateAmazonMotorsRequest {
	this := CreateAmazonMotorsRequest{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *CreateAmazonMotorsRequest) GetAttachments() []Attachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAmazonMotorsRequest) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *CreateAmazonMotorsRequest) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *CreateAmazonMotorsRequest) SetAttachments(v []Attachment) {
	o.Attachments = v
}

func (o CreateAmazonMotorsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullableCreateAmazonMotorsRequest struct {
	value *CreateAmazonMotorsRequest
	isSet bool
}

func (v NullableCreateAmazonMotorsRequest) Get() *CreateAmazonMotorsRequest {
	return v.value
}

func (v *NullableCreateAmazonMotorsRequest) Set(val *CreateAmazonMotorsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAmazonMotorsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAmazonMotorsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAmazonMotorsRequest(val *CreateAmazonMotorsRequest) *NullableCreateAmazonMotorsRequest {
	return &NullableCreateAmazonMotorsRequest{value: val, isSet: true}
}

func (v NullableCreateAmazonMotorsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateAmazonMotorsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
