package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateNegativeFeedbackRemovalResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNegativeFeedbackRemovalResponse{}

// CreateNegativeFeedbackRemovalResponse The response schema for the createNegativeFeedbackRemoval operation.
type CreateNegativeFeedbackRemovalResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCreateNegativeFeedbackRemovalResponse instantiates a new CreateNegativeFeedbackRemovalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNegativeFeedbackRemovalResponse() *CreateNegativeFeedbackRemovalResponse {
	this := CreateNegativeFeedbackRemovalResponse{}
	return &this
}

// NewCreateNegativeFeedbackRemovalResponseWithDefaults instantiates a new CreateNegativeFeedbackRemovalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNegativeFeedbackRemovalResponseWithDefaults() *CreateNegativeFeedbackRemovalResponse {
	this := CreateNegativeFeedbackRemovalResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateNegativeFeedbackRemovalResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNegativeFeedbackRemovalResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateNegativeFeedbackRemovalResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CreateNegativeFeedbackRemovalResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CreateNegativeFeedbackRemovalResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCreateNegativeFeedbackRemovalResponse struct {
	value *CreateNegativeFeedbackRemovalResponse
	isSet bool
}

func (v NullableCreateNegativeFeedbackRemovalResponse) Get() *CreateNegativeFeedbackRemovalResponse {
	return v.value
}

func (v *NullableCreateNegativeFeedbackRemovalResponse) Set(val *CreateNegativeFeedbackRemovalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNegativeFeedbackRemovalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNegativeFeedbackRemovalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNegativeFeedbackRemovalResponse(val *CreateNegativeFeedbackRemovalResponse) *NullableCreateNegativeFeedbackRemovalResponse {
	return &NullableCreateNegativeFeedbackRemovalResponse{value: val, isSet: true}
}

func (v NullableCreateNegativeFeedbackRemovalResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateNegativeFeedbackRemovalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
