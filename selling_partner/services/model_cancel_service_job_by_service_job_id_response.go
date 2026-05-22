package services

import (
	"github.com/bytedance/sonic"
)

// checks if the CancelServiceJobByServiceJobIdResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelServiceJobByServiceJobIdResponse{}

// CancelServiceJobByServiceJobIdResponse Response schema for the `cancelServiceJobByServiceJobId` operation.
type CancelServiceJobByServiceJobIdResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCancelServiceJobByServiceJobIdResponse instantiates a new CancelServiceJobByServiceJobIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelServiceJobByServiceJobIdResponse() *CancelServiceJobByServiceJobIdResponse {
	this := CancelServiceJobByServiceJobIdResponse{}
	return &this
}

// NewCancelServiceJobByServiceJobIdResponseWithDefaults instantiates a new CancelServiceJobByServiceJobIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelServiceJobByServiceJobIdResponseWithDefaults() *CancelServiceJobByServiceJobIdResponse {
	this := CancelServiceJobByServiceJobIdResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CancelServiceJobByServiceJobIdResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelServiceJobByServiceJobIdResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CancelServiceJobByServiceJobIdResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CancelServiceJobByServiceJobIdResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CancelServiceJobByServiceJobIdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCancelServiceJobByServiceJobIdResponse struct {
	value *CancelServiceJobByServiceJobIdResponse
	isSet bool
}

func (v NullableCancelServiceJobByServiceJobIdResponse) Get() *CancelServiceJobByServiceJobIdResponse {
	return v.value
}

func (v *NullableCancelServiceJobByServiceJobIdResponse) Set(val *CancelServiceJobByServiceJobIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelServiceJobByServiceJobIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelServiceJobByServiceJobIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelServiceJobByServiceJobIdResponse(val *CancelServiceJobByServiceJobIdResponse) *NullableCancelServiceJobByServiceJobIdResponse {
	return &NullableCancelServiceJobByServiceJobIdResponse{value: val, isSet: true}
}

func (v NullableCancelServiceJobByServiceJobIdResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCancelServiceJobByServiceJobIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
