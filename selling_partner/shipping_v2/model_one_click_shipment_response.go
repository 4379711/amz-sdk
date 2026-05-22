package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the OneClickShipmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OneClickShipmentResponse{}

// OneClickShipmentResponse The response schema for the OneClickShipment operation.
type OneClickShipmentResponse struct {
	Payload *OneClickShipmentResult `json:"payload,omitempty"`
}

// NewOneClickShipmentResponse instantiates a new OneClickShipmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOneClickShipmentResponse() *OneClickShipmentResponse {
	this := OneClickShipmentResponse{}
	return &this
}

// NewOneClickShipmentResponseWithDefaults instantiates a new OneClickShipmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOneClickShipmentResponseWithDefaults() *OneClickShipmentResponse {
	this := OneClickShipmentResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *OneClickShipmentResponse) GetPayload() OneClickShipmentResult {
	if o == nil || IsNil(o.Payload) {
		var ret OneClickShipmentResult
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneClickShipmentResponse) GetPayloadOk() (*OneClickShipmentResult, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *OneClickShipmentResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given OneClickShipmentResult and assigns it to the Payload field.
func (o *OneClickShipmentResponse) SetPayload(v OneClickShipmentResult) {
	o.Payload = &v
}

func (o OneClickShipmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	return toSerialize, nil
}

type NullableOneClickShipmentResponse struct {
	value *OneClickShipmentResponse
	isSet bool
}

func (v NullableOneClickShipmentResponse) Get() *OneClickShipmentResponse {
	return v.value
}

func (v *NullableOneClickShipmentResponse) Set(val *OneClickShipmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOneClickShipmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOneClickShipmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneClickShipmentResponse(val *OneClickShipmentResponse) *NullableOneClickShipmentResponse {
	return &NullableOneClickShipmentResponse{value: val, isSet: true}
}

func (v NullableOneClickShipmentResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOneClickShipmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
