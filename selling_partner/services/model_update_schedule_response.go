package services

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateScheduleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateScheduleResponse{}

// UpdateScheduleResponse Response schema for the `updateSchedule` operation.
type UpdateScheduleResponse struct {
	// Contains the `UpdateScheduleRecords` for which the error/warning has occurred.
	Payload []UpdateScheduleRecord `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewUpdateScheduleResponse instantiates a new UpdateScheduleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateScheduleResponse() *UpdateScheduleResponse {
	this := UpdateScheduleResponse{}
	return &this
}

// NewUpdateScheduleResponseWithDefaults instantiates a new UpdateScheduleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateScheduleResponseWithDefaults() *UpdateScheduleResponse {
	this := UpdateScheduleResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *UpdateScheduleResponse) GetPayload() []UpdateScheduleRecord {
	if o == nil || IsNil(o.Payload) {
		var ret []UpdateScheduleRecord
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScheduleResponse) GetPayloadOk() ([]UpdateScheduleRecord, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *UpdateScheduleResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given []UpdateScheduleRecord and assigns it to the Payload field.
func (o *UpdateScheduleResponse) SetPayload(v []UpdateScheduleRecord) {
	o.Payload = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *UpdateScheduleResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScheduleResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *UpdateScheduleResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *UpdateScheduleResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o UpdateScheduleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableUpdateScheduleResponse struct {
	value *UpdateScheduleResponse
	isSet bool
}

func (v NullableUpdateScheduleResponse) Get() *UpdateScheduleResponse {
	return v.value
}

func (v *NullableUpdateScheduleResponse) Set(val *UpdateScheduleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateScheduleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateScheduleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateScheduleResponse(val *UpdateScheduleResponse) *NullableUpdateScheduleResponse {
	return &NullableUpdateScheduleResponse{value: val, isSet: true}
}

func (v NullableUpdateScheduleResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateScheduleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
