package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateDigitalAccessKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDigitalAccessKeyResponse{}

// CreateDigitalAccessKeyResponse The response schema for the `createDigitalAccessKey` operation.
type CreateDigitalAccessKeyResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCreateDigitalAccessKeyResponse instantiates a new CreateDigitalAccessKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDigitalAccessKeyResponse() *CreateDigitalAccessKeyResponse {
	this := CreateDigitalAccessKeyResponse{}
	return &this
}

// NewCreateDigitalAccessKeyResponseWithDefaults instantiates a new CreateDigitalAccessKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDigitalAccessKeyResponseWithDefaults() *CreateDigitalAccessKeyResponse {
	this := CreateDigitalAccessKeyResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateDigitalAccessKeyResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDigitalAccessKeyResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateDigitalAccessKeyResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CreateDigitalAccessKeyResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CreateDigitalAccessKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCreateDigitalAccessKeyResponse struct {
	value *CreateDigitalAccessKeyResponse
	isSet bool
}

func (v NullableCreateDigitalAccessKeyResponse) Get() *CreateDigitalAccessKeyResponse {
	return v.value
}

func (v *NullableCreateDigitalAccessKeyResponse) Set(val *CreateDigitalAccessKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDigitalAccessKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDigitalAccessKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDigitalAccessKeyResponse(val *CreateDigitalAccessKeyResponse) *NullableCreateDigitalAccessKeyResponse {
	return &NullableCreateDigitalAccessKeyResponse{value: val, isSet: true}
}

func (v NullableCreateDigitalAccessKeyResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateDigitalAccessKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
