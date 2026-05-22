package services

import (
	"github.com/bytedance/sonic"
)

// checks if the GetAppointmentSlotsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAppointmentSlotsResponse{}

// GetAppointmentSlotsResponse The response of fetching appointment slots based on service context.
type GetAppointmentSlotsResponse struct {
	Payload *AppointmentSlotReport `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetAppointmentSlotsResponse instantiates a new GetAppointmentSlotsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAppointmentSlotsResponse() *GetAppointmentSlotsResponse {
	this := GetAppointmentSlotsResponse{}
	return &this
}

// NewGetAppointmentSlotsResponseWithDefaults instantiates a new GetAppointmentSlotsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAppointmentSlotsResponseWithDefaults() *GetAppointmentSlotsResponse {
	this := GetAppointmentSlotsResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetAppointmentSlotsResponse) GetPayload() AppointmentSlotReport {
	if o == nil || IsNil(o.Payload) {
		var ret AppointmentSlotReport
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppointmentSlotsResponse) GetPayloadOk() (*AppointmentSlotReport, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetAppointmentSlotsResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given AppointmentSlotReport and assigns it to the Payload field.
func (o *GetAppointmentSlotsResponse) SetPayload(v AppointmentSlotReport) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetAppointmentSlotsResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppointmentSlotsResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetAppointmentSlotsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetAppointmentSlotsResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetAppointmentSlotsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetAppointmentSlotsResponse struct {
	value *GetAppointmentSlotsResponse
	isSet bool
}

func (v NullableGetAppointmentSlotsResponse) Get() *GetAppointmentSlotsResponse {
	return v.value
}

func (v *NullableGetAppointmentSlotsResponse) Set(val *GetAppointmentSlotsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAppointmentSlotsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAppointmentSlotsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAppointmentSlotsResponse(val *GetAppointmentSlotsResponse) *NullableGetAppointmentSlotsResponse {
	return &NullableGetAppointmentSlotsResponse{value: val, isSet: true}
}

func (v NullableGetAppointmentSlotsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetAppointmentSlotsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
