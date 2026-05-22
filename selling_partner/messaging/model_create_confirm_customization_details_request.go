package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateConfirmCustomizationDetailsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateConfirmCustomizationDetailsRequest{}

// CreateConfirmCustomizationDetailsRequest The request schema for the confirmCustomizationDetails operation.
type CreateConfirmCustomizationDetailsRequest struct {
	// The text to be sent to the buyer. Only links related to customization details are allowed. Do not include HTML or email addresses. The text must be written in the buyer's language of preference, which can be retrieved from the GetAttributes operation.
	Text *string `json:"text,omitempty"`
	// Attachments to include in the message to the buyer.
	Attachments []Attachment `json:"attachments,omitempty"`
}

// NewCreateConfirmCustomizationDetailsRequest instantiates a new CreateConfirmCustomizationDetailsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConfirmCustomizationDetailsRequest() *CreateConfirmCustomizationDetailsRequest {
	this := CreateConfirmCustomizationDetailsRequest{}
	return &this
}

// NewCreateConfirmCustomizationDetailsRequestWithDefaults instantiates a new CreateConfirmCustomizationDetailsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConfirmCustomizationDetailsRequestWithDefaults() *CreateConfirmCustomizationDetailsRequest {
	this := CreateConfirmCustomizationDetailsRequest{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *CreateConfirmCustomizationDetailsRequest) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConfirmCustomizationDetailsRequest) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *CreateConfirmCustomizationDetailsRequest) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *CreateConfirmCustomizationDetailsRequest) SetText(v string) {
	o.Text = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *CreateConfirmCustomizationDetailsRequest) GetAttachments() []Attachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConfirmCustomizationDetailsRequest) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *CreateConfirmCustomizationDetailsRequest) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *CreateConfirmCustomizationDetailsRequest) SetAttachments(v []Attachment) {
	o.Attachments = v
}

func (o CreateConfirmCustomizationDetailsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullableCreateConfirmCustomizationDetailsRequest struct {
	value *CreateConfirmCustomizationDetailsRequest
	isSet bool
}

func (v NullableCreateConfirmCustomizationDetailsRequest) Get() *CreateConfirmCustomizationDetailsRequest {
	return v.value
}

func (v *NullableCreateConfirmCustomizationDetailsRequest) Set(val *CreateConfirmCustomizationDetailsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConfirmCustomizationDetailsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConfirmCustomizationDetailsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConfirmCustomizationDetailsRequest(val *CreateConfirmCustomizationDetailsRequest) *NullableCreateConfirmCustomizationDetailsRequest {
	return &NullableCreateConfirmCustomizationDetailsRequest{value: val, isSet: true}
}

func (v NullableCreateConfirmCustomizationDetailsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateConfirmCustomizationDetailsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
