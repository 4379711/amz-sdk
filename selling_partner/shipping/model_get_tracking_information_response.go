package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the GetTrackingInformationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTrackingInformationResponse{}

// GetTrackingInformationResponse The response schema for the getTrackingInformation operation.
type GetTrackingInformationResponse struct {
	Payload *TrackingInformation `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetTrackingInformationResponse instantiates a new GetTrackingInformationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTrackingInformationResponse() *GetTrackingInformationResponse {
	this := GetTrackingInformationResponse{}
	return &this
}

// NewGetTrackingInformationResponseWithDefaults instantiates a new GetTrackingInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTrackingInformationResponseWithDefaults() *GetTrackingInformationResponse {
	this := GetTrackingInformationResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetTrackingInformationResponse) GetPayload() TrackingInformation {
	if o == nil || IsNil(o.Payload) {
		var ret TrackingInformation
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrackingInformationResponse) GetPayloadOk() (*TrackingInformation, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetTrackingInformationResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given TrackingInformation and assigns it to the Payload field.
func (o *GetTrackingInformationResponse) SetPayload(v TrackingInformation) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetTrackingInformationResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrackingInformationResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetTrackingInformationResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetTrackingInformationResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetTrackingInformationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetTrackingInformationResponse struct {
	value *GetTrackingInformationResponse
	isSet bool
}

func (v NullableGetTrackingInformationResponse) Get() *GetTrackingInformationResponse {
	return v.value
}

func (v *NullableGetTrackingInformationResponse) Set(val *GetTrackingInformationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTrackingInformationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTrackingInformationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTrackingInformationResponse(val *GetTrackingInformationResponse) *NullableGetTrackingInformationResponse {
	return &NullableGetTrackingInformationResponse{value: val, isSet: true}
}

func (v NullableGetTrackingInformationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetTrackingInformationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
