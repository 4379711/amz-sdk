package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateConfirmServiceDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateConfirmServiceDetailsResponse{}

// CreateConfirmServiceDetailsResponse The response schema for the createConfirmServiceDetails operation.
type CreateConfirmServiceDetailsResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCreateConfirmServiceDetailsResponse instantiates a new CreateConfirmServiceDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConfirmServiceDetailsResponse() *CreateConfirmServiceDetailsResponse {
	this := CreateConfirmServiceDetailsResponse{}
	return &this
}

// NewCreateConfirmServiceDetailsResponseWithDefaults instantiates a new CreateConfirmServiceDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConfirmServiceDetailsResponseWithDefaults() *CreateConfirmServiceDetailsResponse {
	this := CreateConfirmServiceDetailsResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateConfirmServiceDetailsResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConfirmServiceDetailsResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateConfirmServiceDetailsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CreateConfirmServiceDetailsResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CreateConfirmServiceDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCreateConfirmServiceDetailsResponse struct {
	value *CreateConfirmServiceDetailsResponse
	isSet bool
}

func (v NullableCreateConfirmServiceDetailsResponse) Get() *CreateConfirmServiceDetailsResponse {
	return v.value
}

func (v *NullableCreateConfirmServiceDetailsResponse) Set(val *CreateConfirmServiceDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConfirmServiceDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConfirmServiceDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConfirmServiceDetailsResponse(val *CreateConfirmServiceDetailsResponse) *NullableCreateConfirmServiceDetailsResponse {
	return &NullableCreateConfirmServiceDetailsResponse{value: val, isSet: true}
}

func (v NullableCreateConfirmServiceDetailsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateConfirmServiceDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
