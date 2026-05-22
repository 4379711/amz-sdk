package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateVerificationStatusErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVerificationStatusErrorResponse{}

// UpdateVerificationStatusErrorResponse The error response schema for the `UpdateVerificationStatus` operation.
type UpdateVerificationStatusErrorResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewUpdateVerificationStatusErrorResponse instantiates a new UpdateVerificationStatusErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVerificationStatusErrorResponse() *UpdateVerificationStatusErrorResponse {
	this := UpdateVerificationStatusErrorResponse{}
	return &this
}

// NewUpdateVerificationStatusErrorResponseWithDefaults instantiates a new UpdateVerificationStatusErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVerificationStatusErrorResponseWithDefaults() *UpdateVerificationStatusErrorResponse {
	this := UpdateVerificationStatusErrorResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *UpdateVerificationStatusErrorResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVerificationStatusErrorResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *UpdateVerificationStatusErrorResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *UpdateVerificationStatusErrorResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o UpdateVerificationStatusErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableUpdateVerificationStatusErrorResponse struct {
	value *UpdateVerificationStatusErrorResponse
	isSet bool
}

func (v NullableUpdateVerificationStatusErrorResponse) Get() *UpdateVerificationStatusErrorResponse {
	return v.value
}

func (v *NullableUpdateVerificationStatusErrorResponse) Set(val *UpdateVerificationStatusErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVerificationStatusErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVerificationStatusErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVerificationStatusErrorResponse(val *UpdateVerificationStatusErrorResponse) *NullableUpdateVerificationStatusErrorResponse {
	return &NullableUpdateVerificationStatusErrorResponse{value: val, isSet: true}
}

func (v NullableUpdateVerificationStatusErrorResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateVerificationStatusErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
