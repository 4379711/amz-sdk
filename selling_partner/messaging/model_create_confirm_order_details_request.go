package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateConfirmOrderDetailsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateConfirmOrderDetailsRequest{}

// CreateConfirmOrderDetailsRequest The request schema for the createConfirmOrderDetails operation.
type CreateConfirmOrderDetailsRequest struct {
	// The text to be sent to the buyer. Only links related to order completion are allowed. Do not include HTML or email addresses. The text must be written in the buyer's language of preference, which can be retrieved from the GetAttributes operation.
	Text *string `json:"text,omitempty"`
}

// NewCreateConfirmOrderDetailsRequest instantiates a new CreateConfirmOrderDetailsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConfirmOrderDetailsRequest() *CreateConfirmOrderDetailsRequest {
	this := CreateConfirmOrderDetailsRequest{}
	return &this
}

// NewCreateConfirmOrderDetailsRequestWithDefaults instantiates a new CreateConfirmOrderDetailsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConfirmOrderDetailsRequestWithDefaults() *CreateConfirmOrderDetailsRequest {
	this := CreateConfirmOrderDetailsRequest{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *CreateConfirmOrderDetailsRequest) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConfirmOrderDetailsRequest) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *CreateConfirmOrderDetailsRequest) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *CreateConfirmOrderDetailsRequest) SetText(v string) {
	o.Text = &v
}

func (o CreateConfirmOrderDetailsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableCreateConfirmOrderDetailsRequest struct {
	value *CreateConfirmOrderDetailsRequest
	isSet bool
}

func (v NullableCreateConfirmOrderDetailsRequest) Get() *CreateConfirmOrderDetailsRequest {
	return v.value
}

func (v *NullableCreateConfirmOrderDetailsRequest) Set(val *CreateConfirmOrderDetailsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConfirmOrderDetailsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConfirmOrderDetailsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConfirmOrderDetailsRequest(val *CreateConfirmOrderDetailsRequest) *NullableCreateConfirmOrderDetailsRequest {
	return &NullableCreateConfirmOrderDetailsRequest{value: val, isSet: true}
}

func (v NullableCreateConfirmOrderDetailsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateConfirmOrderDetailsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
