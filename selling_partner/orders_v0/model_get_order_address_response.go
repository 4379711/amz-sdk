package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetOrderAddressResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrderAddressResponse{}

// GetOrderAddressResponse The response schema for the `getOrderAddress` operation.
type GetOrderAddressResponse struct {
	Payload *OrderAddress `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetOrderAddressResponse instantiates a new GetOrderAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderAddressResponse() *GetOrderAddressResponse {
	this := GetOrderAddressResponse{}
	return &this
}

// NewGetOrderAddressResponseWithDefaults instantiates a new GetOrderAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderAddressResponseWithDefaults() *GetOrderAddressResponse {
	this := GetOrderAddressResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetOrderAddressResponse) GetPayload() OrderAddress {
	if o == nil || IsNil(o.Payload) {
		var ret OrderAddress
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderAddressResponse) GetPayloadOk() (*OrderAddress, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetOrderAddressResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given OrderAddress and assigns it to the Payload field.
func (o *GetOrderAddressResponse) SetPayload(v OrderAddress) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetOrderAddressResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderAddressResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetOrderAddressResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetOrderAddressResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetOrderAddressResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetOrderAddressResponse struct {
	value *GetOrderAddressResponse
	isSet bool
}

func (v NullableGetOrderAddressResponse) Get() *GetOrderAddressResponse {
	return v.value
}

func (v *NullableGetOrderAddressResponse) Set(val *GetOrderAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderAddressResponse(val *GetOrderAddressResponse) *NullableGetOrderAddressResponse {
	return &NullableGetOrderAddressResponse{value: val, isSet: true}
}

func (v NullableGetOrderAddressResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetOrderAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
