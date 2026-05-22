package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the GetTrackingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTrackingResponse{}

// GetTrackingResponse The response schema for the getTracking operation.
type GetTrackingResponse struct {
	Payload *GetTrackingResult `json:"payload,omitempty"`
}

// NewGetTrackingResponse instantiates a new GetTrackingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTrackingResponse() *GetTrackingResponse {
	this := GetTrackingResponse{}
	return &this
}

// NewGetTrackingResponseWithDefaults instantiates a new GetTrackingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTrackingResponseWithDefaults() *GetTrackingResponse {
	this := GetTrackingResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetTrackingResponse) GetPayload() GetTrackingResult {
	if o == nil || IsNil(o.Payload) {
		var ret GetTrackingResult
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrackingResponse) GetPayloadOk() (*GetTrackingResult, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetTrackingResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given GetTrackingResult and assigns it to the Payload field.
func (o *GetTrackingResponse) SetPayload(v GetTrackingResult) {
	o.Payload = &v
}

func (o GetTrackingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	return toSerialize, nil
}

type NullableGetTrackingResponse struct {
	value *GetTrackingResponse
	isSet bool
}

func (v NullableGetTrackingResponse) Get() *GetTrackingResponse {
	return v.value
}

func (v *NullableGetTrackingResponse) Set(val *GetTrackingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTrackingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTrackingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTrackingResponse(val *GetTrackingResponse) *NullableGetTrackingResponse {
	return &NullableGetTrackingResponse{value: val, isSet: true}
}

func (v NullableGetTrackingResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetTrackingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
