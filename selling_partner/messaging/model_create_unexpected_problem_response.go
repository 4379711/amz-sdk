package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateUnexpectedProblemResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUnexpectedProblemResponse{}

// CreateUnexpectedProblemResponse The response schema for the createUnexpectedProblem operation.
type CreateUnexpectedProblemResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCreateUnexpectedProblemResponse instantiates a new CreateUnexpectedProblemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUnexpectedProblemResponse() *CreateUnexpectedProblemResponse {
	this := CreateUnexpectedProblemResponse{}
	return &this
}

// NewCreateUnexpectedProblemResponseWithDefaults instantiates a new CreateUnexpectedProblemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUnexpectedProblemResponseWithDefaults() *CreateUnexpectedProblemResponse {
	this := CreateUnexpectedProblemResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateUnexpectedProblemResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUnexpectedProblemResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateUnexpectedProblemResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CreateUnexpectedProblemResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CreateUnexpectedProblemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCreateUnexpectedProblemResponse struct {
	value *CreateUnexpectedProblemResponse
	isSet bool
}

func (v NullableCreateUnexpectedProblemResponse) Get() *CreateUnexpectedProblemResponse {
	return v.value
}

func (v *NullableCreateUnexpectedProblemResponse) Set(val *CreateUnexpectedProblemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUnexpectedProblemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUnexpectedProblemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUnexpectedProblemResponse(val *CreateUnexpectedProblemResponse) *NullableCreateUnexpectedProblemResponse {
	return &NullableCreateUnexpectedProblemResponse{value: val, isSet: true}
}

func (v NullableCreateUnexpectedProblemResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateUnexpectedProblemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
