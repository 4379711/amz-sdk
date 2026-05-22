package fulfillment_inbound_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetShipmentItemsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetShipmentItemsResponse{}

// GetShipmentItemsResponse The response schema for the getShipmentItems operation.
type GetShipmentItemsResponse struct {
	Payload *GetShipmentItemsResult `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetShipmentItemsResponse instantiates a new GetShipmentItemsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetShipmentItemsResponse() *GetShipmentItemsResponse {
	this := GetShipmentItemsResponse{}
	return &this
}

// NewGetShipmentItemsResponseWithDefaults instantiates a new GetShipmentItemsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetShipmentItemsResponseWithDefaults() *GetShipmentItemsResponse {
	this := GetShipmentItemsResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetShipmentItemsResponse) GetPayload() GetShipmentItemsResult {
	if o == nil || IsNil(o.Payload) {
		var ret GetShipmentItemsResult
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentItemsResponse) GetPayloadOk() (*GetShipmentItemsResult, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetShipmentItemsResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given GetShipmentItemsResult and assigns it to the Payload field.
func (o *GetShipmentItemsResponse) SetPayload(v GetShipmentItemsResult) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetShipmentItemsResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentItemsResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetShipmentItemsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetShipmentItemsResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetShipmentItemsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetShipmentItemsResponse struct {
	value *GetShipmentItemsResponse
	isSet bool
}

func (v NullableGetShipmentItemsResponse) Get() *GetShipmentItemsResponse {
	return v.value
}

func (v *NullableGetShipmentItemsResponse) Set(val *GetShipmentItemsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetShipmentItemsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetShipmentItemsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetShipmentItemsResponse(val *GetShipmentItemsResponse) *NullableGetShipmentItemsResponse {
	return &NullableGetShipmentItemsResponse{value: val, isSet: true}
}

func (v NullableGetShipmentItemsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetShipmentItemsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
