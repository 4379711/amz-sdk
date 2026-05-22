package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the GetAdditionalInputsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAdditionalInputsResponse{}

// GetAdditionalInputsResponse The response schema for the getAdditionalInputs operation.
type GetAdditionalInputsResponse struct {
	// The JSON schema to use to provide additional inputs when required to purchase a shipping offering.
	Payload map[string]map[string]interface{} `json:"payload,omitempty"`
}

// NewGetAdditionalInputsResponse instantiates a new GetAdditionalInputsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAdditionalInputsResponse() *GetAdditionalInputsResponse {
	this := GetAdditionalInputsResponse{}
	return &this
}

// NewGetAdditionalInputsResponseWithDefaults instantiates a new GetAdditionalInputsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAdditionalInputsResponseWithDefaults() *GetAdditionalInputsResponse {
	this := GetAdditionalInputsResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetAdditionalInputsResponse) GetPayload() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Payload) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdditionalInputsResponse) GetPayloadOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Payload) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetAdditionalInputsResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given map[string]map[string]interface{} and assigns it to the Payload field.
func (o *GetAdditionalInputsResponse) SetPayload(v map[string]map[string]interface{}) {
	o.Payload = v
}

func (o GetAdditionalInputsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	return toSerialize, nil
}

type NullableGetAdditionalInputsResponse struct {
	value *GetAdditionalInputsResponse
	isSet bool
}

func (v NullableGetAdditionalInputsResponse) Get() *GetAdditionalInputsResponse {
	return v.value
}

func (v *NullableGetAdditionalInputsResponse) Set(val *GetAdditionalInputsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAdditionalInputsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAdditionalInputsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAdditionalInputsResponse(val *GetAdditionalInputsResponse) *NullableGetAdditionalInputsResponse {
	return &NullableGetAdditionalInputsResponse{value: val, isSet: true}
}

func (v NullableGetAdditionalInputsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetAdditionalInputsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
