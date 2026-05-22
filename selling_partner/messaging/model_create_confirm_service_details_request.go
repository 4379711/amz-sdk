package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateConfirmServiceDetailsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateConfirmServiceDetailsRequest{}

// CreateConfirmServiceDetailsRequest The request schema for the createConfirmServiceDetails operation.
type CreateConfirmServiceDetailsRequest struct {
	// The text to be sent to the buyer. Only links related to Home Service calls are allowed. Do not include HTML or email addresses. The text must be written in the buyer's language of preference, which can be retrieved from the GetAttributes operation.
	Text *string `json:"text,omitempty"`
}

// NewCreateConfirmServiceDetailsRequest instantiates a new CreateConfirmServiceDetailsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConfirmServiceDetailsRequest() *CreateConfirmServiceDetailsRequest {
	this := CreateConfirmServiceDetailsRequest{}
	return &this
}

// NewCreateConfirmServiceDetailsRequestWithDefaults instantiates a new CreateConfirmServiceDetailsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConfirmServiceDetailsRequestWithDefaults() *CreateConfirmServiceDetailsRequest {
	this := CreateConfirmServiceDetailsRequest{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *CreateConfirmServiceDetailsRequest) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConfirmServiceDetailsRequest) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *CreateConfirmServiceDetailsRequest) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *CreateConfirmServiceDetailsRequest) SetText(v string) {
	o.Text = &v
}

func (o CreateConfirmServiceDetailsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableCreateConfirmServiceDetailsRequest struct {
	value *CreateConfirmServiceDetailsRequest
	isSet bool
}

func (v NullableCreateConfirmServiceDetailsRequest) Get() *CreateConfirmServiceDetailsRequest {
	return v.value
}

func (v *NullableCreateConfirmServiceDetailsRequest) Set(val *CreateConfirmServiceDetailsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConfirmServiceDetailsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConfirmServiceDetailsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConfirmServiceDetailsRequest(val *CreateConfirmServiceDetailsRequest) *NullableCreateConfirmServiceDetailsRequest {
	return &NullableCreateConfirmServiceDetailsRequest{value: val, isSet: true}
}

func (v NullableCreateConfirmServiceDetailsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateConfirmServiceDetailsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
