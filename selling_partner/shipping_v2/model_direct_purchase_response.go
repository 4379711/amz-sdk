package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the DirectPurchaseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectPurchaseResponse{}

// DirectPurchaseResponse The response schema for the directPurchaseShipment operation.
type DirectPurchaseResponse struct {
	Payload *DirectPurchaseResult `json:"payload,omitempty"`
}

// NewDirectPurchaseResponse instantiates a new DirectPurchaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectPurchaseResponse() *DirectPurchaseResponse {
	this := DirectPurchaseResponse{}
	return &this
}

// NewDirectPurchaseResponseWithDefaults instantiates a new DirectPurchaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectPurchaseResponseWithDefaults() *DirectPurchaseResponse {
	this := DirectPurchaseResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *DirectPurchaseResponse) GetPayload() DirectPurchaseResult {
	if o == nil || IsNil(o.Payload) {
		var ret DirectPurchaseResult
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectPurchaseResponse) GetPayloadOk() (*DirectPurchaseResult, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *DirectPurchaseResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given DirectPurchaseResult and assigns it to the Payload field.
func (o *DirectPurchaseResponse) SetPayload(v DirectPurchaseResult) {
	o.Payload = &v
}

func (o DirectPurchaseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	return toSerialize, nil
}

type NullableDirectPurchaseResponse struct {
	value *DirectPurchaseResponse
	isSet bool
}

func (v NullableDirectPurchaseResponse) Get() *DirectPurchaseResponse {
	return v.value
}

func (v *NullableDirectPurchaseResponse) Set(val *DirectPurchaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectPurchaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectPurchaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectPurchaseResponse(val *DirectPurchaseResponse) *NullableDirectPurchaseResponse {
	return &NullableDirectPurchaseResponse{value: val, isSet: true}
}

func (v NullableDirectPurchaseResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDirectPurchaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
