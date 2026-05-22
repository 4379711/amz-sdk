package services

import (
	"github.com/bytedance/sonic"
)

// checks if the GetServiceJobsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetServiceJobsResponse{}

// GetServiceJobsResponse Response schema for the `getServiceJobs` operation.
type GetServiceJobsResponse struct {
	Payload *JobListing `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetServiceJobsResponse instantiates a new GetServiceJobsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServiceJobsResponse() *GetServiceJobsResponse {
	this := GetServiceJobsResponse{}
	return &this
}

// NewGetServiceJobsResponseWithDefaults instantiates a new GetServiceJobsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServiceJobsResponseWithDefaults() *GetServiceJobsResponse {
	this := GetServiceJobsResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetServiceJobsResponse) GetPayload() JobListing {
	if o == nil || IsNil(o.Payload) {
		var ret JobListing
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServiceJobsResponse) GetPayloadOk() (*JobListing, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetServiceJobsResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given JobListing and assigns it to the Payload field.
func (o *GetServiceJobsResponse) SetPayload(v JobListing) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetServiceJobsResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServiceJobsResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetServiceJobsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetServiceJobsResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetServiceJobsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetServiceJobsResponse struct {
	value *GetServiceJobsResponse
	isSet bool
}

func (v NullableGetServiceJobsResponse) Get() *GetServiceJobsResponse {
	return v.value
}

func (v *NullableGetServiceJobsResponse) Set(val *GetServiceJobsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServiceJobsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServiceJobsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServiceJobsResponse(val *GetServiceJobsResponse) *NullableGetServiceJobsResponse {
	return &NullableGetServiceJobsResponse{value: val, isSet: true}
}

func (v NullableGetServiceJobsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetServiceJobsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
