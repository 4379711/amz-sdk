package services

import (
	"github.com/bytedance/sonic"
)

// checks if the CompleteServiceJobByServiceJobIdResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompleteServiceJobByServiceJobIdResponse{}

// CompleteServiceJobByServiceJobIdResponse Response schema for the `completeServiceJobByServiceJobId` operation.
type CompleteServiceJobByServiceJobIdResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCompleteServiceJobByServiceJobIdResponse instantiates a new CompleteServiceJobByServiceJobIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompleteServiceJobByServiceJobIdResponse() *CompleteServiceJobByServiceJobIdResponse {
	this := CompleteServiceJobByServiceJobIdResponse{}
	return &this
}

// NewCompleteServiceJobByServiceJobIdResponseWithDefaults instantiates a new CompleteServiceJobByServiceJobIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompleteServiceJobByServiceJobIdResponseWithDefaults() *CompleteServiceJobByServiceJobIdResponse {
	this := CompleteServiceJobByServiceJobIdResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CompleteServiceJobByServiceJobIdResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteServiceJobByServiceJobIdResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CompleteServiceJobByServiceJobIdResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CompleteServiceJobByServiceJobIdResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CompleteServiceJobByServiceJobIdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCompleteServiceJobByServiceJobIdResponse struct {
	value *CompleteServiceJobByServiceJobIdResponse
	isSet bool
}

func (v NullableCompleteServiceJobByServiceJobIdResponse) Get() *CompleteServiceJobByServiceJobIdResponse {
	return v.value
}

func (v *NullableCompleteServiceJobByServiceJobIdResponse) Set(val *CompleteServiceJobByServiceJobIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCompleteServiceJobByServiceJobIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCompleteServiceJobByServiceJobIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompleteServiceJobByServiceJobIdResponse(val *CompleteServiceJobByServiceJobIdResponse) *NullableCompleteServiceJobByServiceJobIdResponse {
	return &NullableCompleteServiceJobByServiceJobIdResponse{value: val, isSet: true}
}

func (v NullableCompleteServiceJobByServiceJobIdResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCompleteServiceJobByServiceJobIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
