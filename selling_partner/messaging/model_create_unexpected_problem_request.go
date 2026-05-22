package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateUnexpectedProblemRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUnexpectedProblemRequest{}

// CreateUnexpectedProblemRequest The request schema for the createUnexpectedProblem operation.
type CreateUnexpectedProblemRequest struct {
	// The text to be sent to the buyer. Only links related to unexpected problem calls are allowed. Do not include HTML or email addresses. The text must be written in the buyer's language of preference, which can be retrieved from the GetAttributes operation.
	Text *string `json:"text,omitempty"`
}

// NewCreateUnexpectedProblemRequest instantiates a new CreateUnexpectedProblemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUnexpectedProblemRequest() *CreateUnexpectedProblemRequest {
	this := CreateUnexpectedProblemRequest{}
	return &this
}

// NewCreateUnexpectedProblemRequestWithDefaults instantiates a new CreateUnexpectedProblemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUnexpectedProblemRequestWithDefaults() *CreateUnexpectedProblemRequest {
	this := CreateUnexpectedProblemRequest{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *CreateUnexpectedProblemRequest) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUnexpectedProblemRequest) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *CreateUnexpectedProblemRequest) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *CreateUnexpectedProblemRequest) SetText(v string) {
	o.Text = &v
}

func (o CreateUnexpectedProblemRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableCreateUnexpectedProblemRequest struct {
	value *CreateUnexpectedProblemRequest
	isSet bool
}

func (v NullableCreateUnexpectedProblemRequest) Get() *CreateUnexpectedProblemRequest {
	return v.value
}

func (v *NullableCreateUnexpectedProblemRequest) Set(val *CreateUnexpectedProblemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUnexpectedProblemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUnexpectedProblemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUnexpectedProblemRequest(val *CreateUnexpectedProblemRequest) *NullableCreateUnexpectedProblemRequest {
	return &NullableCreateUnexpectedProblemRequest{value: val, isSet: true}
}

func (v NullableCreateUnexpectedProblemRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateUnexpectedProblemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
