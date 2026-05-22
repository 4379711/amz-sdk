package fulfillment_inbound_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetShipmentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetShipmentsResponse{}

// GetShipmentsResponse The response schema for the getShipments operation.
type GetShipmentsResponse struct {
	Payload *GetShipmentsResult `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetShipmentsResponse instantiates a new GetShipmentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetShipmentsResponse() *GetShipmentsResponse {
	this := GetShipmentsResponse{}
	return &this
}

// NewGetShipmentsResponseWithDefaults instantiates a new GetShipmentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetShipmentsResponseWithDefaults() *GetShipmentsResponse {
	this := GetShipmentsResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetShipmentsResponse) GetPayload() GetShipmentsResult {
	if o == nil || IsNil(o.Payload) {
		var ret GetShipmentsResult
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentsResponse) GetPayloadOk() (*GetShipmentsResult, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetShipmentsResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given GetShipmentsResult and assigns it to the Payload field.
func (o *GetShipmentsResponse) SetPayload(v GetShipmentsResult) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetShipmentsResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentsResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetShipmentsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetShipmentsResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetShipmentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetShipmentsResponse struct {
	value *GetShipmentsResponse
	isSet bool
}

func (v NullableGetShipmentsResponse) Get() *GetShipmentsResponse {
	return v.value
}

func (v *NullableGetShipmentsResponse) Set(val *GetShipmentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetShipmentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetShipmentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetShipmentsResponse(val *GetShipmentsResponse) *NullableGetShipmentsResponse {
	return &NullableGetShipmentsResponse{value: val, isSet: true}
}

func (v NullableGetShipmentsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetShipmentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
