package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateConfirmDeliveryDetailsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateConfirmDeliveryDetailsRequest{}

// CreateConfirmDeliveryDetailsRequest The request schema for the createConfirmDeliveryDetails operation.
type CreateConfirmDeliveryDetailsRequest struct {
	// The text to be sent to the buyer. Only links related to order delivery are allowed. Do not include HTML or email addresses. The text must be written in the buyer's language of preference, which can be retrieved from the GetAttributes operation.
	Text *string `json:"text,omitempty"`
}

// NewCreateConfirmDeliveryDetailsRequest instantiates a new CreateConfirmDeliveryDetailsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConfirmDeliveryDetailsRequest() *CreateConfirmDeliveryDetailsRequest {
	this := CreateConfirmDeliveryDetailsRequest{}
	return &this
}

// NewCreateConfirmDeliveryDetailsRequestWithDefaults instantiates a new CreateConfirmDeliveryDetailsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConfirmDeliveryDetailsRequestWithDefaults() *CreateConfirmDeliveryDetailsRequest {
	this := CreateConfirmDeliveryDetailsRequest{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *CreateConfirmDeliveryDetailsRequest) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConfirmDeliveryDetailsRequest) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *CreateConfirmDeliveryDetailsRequest) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *CreateConfirmDeliveryDetailsRequest) SetText(v string) {
	o.Text = &v
}

func (o CreateConfirmDeliveryDetailsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableCreateConfirmDeliveryDetailsRequest struct {
	value *CreateConfirmDeliveryDetailsRequest
	isSet bool
}

func (v NullableCreateConfirmDeliveryDetailsRequest) Get() *CreateConfirmDeliveryDetailsRequest {
	return v.value
}

func (v *NullableCreateConfirmDeliveryDetailsRequest) Set(val *CreateConfirmDeliveryDetailsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConfirmDeliveryDetailsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConfirmDeliveryDetailsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConfirmDeliveryDetailsRequest(val *CreateConfirmDeliveryDetailsRequest) *NullableCreateConfirmDeliveryDetailsRequest {
	return &NullableCreateConfirmDeliveryDetailsRequest{value: val, isSet: true}
}

func (v NullableCreateConfirmDeliveryDetailsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateConfirmDeliveryDetailsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
