package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetAdditionalSellerInputsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAdditionalSellerInputsResponse{}

// GetAdditionalSellerInputsResponse Response schema.
type GetAdditionalSellerInputsResponse struct {
	Payload *GetAdditionalSellerInputsResult `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetAdditionalSellerInputsResponse instantiates a new GetAdditionalSellerInputsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAdditionalSellerInputsResponse() *GetAdditionalSellerInputsResponse {
	this := GetAdditionalSellerInputsResponse{}
	return &this
}

// NewGetAdditionalSellerInputsResponseWithDefaults instantiates a new GetAdditionalSellerInputsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAdditionalSellerInputsResponseWithDefaults() *GetAdditionalSellerInputsResponse {
	this := GetAdditionalSellerInputsResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetAdditionalSellerInputsResponse) GetPayload() GetAdditionalSellerInputsResult {
	if o == nil || IsNil(o.Payload) {
		var ret GetAdditionalSellerInputsResult
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdditionalSellerInputsResponse) GetPayloadOk() (*GetAdditionalSellerInputsResult, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetAdditionalSellerInputsResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given GetAdditionalSellerInputsResult and assigns it to the Payload field.
func (o *GetAdditionalSellerInputsResponse) SetPayload(v GetAdditionalSellerInputsResult) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetAdditionalSellerInputsResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdditionalSellerInputsResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetAdditionalSellerInputsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetAdditionalSellerInputsResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetAdditionalSellerInputsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetAdditionalSellerInputsResponse struct {
	value *GetAdditionalSellerInputsResponse
	isSet bool
}

func (v NullableGetAdditionalSellerInputsResponse) Get() *GetAdditionalSellerInputsResponse {
	return v.value
}

func (v *NullableGetAdditionalSellerInputsResponse) Set(val *GetAdditionalSellerInputsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAdditionalSellerInputsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAdditionalSellerInputsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAdditionalSellerInputsResponse(val *GetAdditionalSellerInputsResponse) *NullableGetAdditionalSellerInputsResponse {
	return &NullableGetAdditionalSellerInputsResponse{value: val, isSet: true}
}

func (v NullableGetAdditionalSellerInputsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetAdditionalSellerInputsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
