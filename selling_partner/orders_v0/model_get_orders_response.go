package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetOrdersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrdersResponse{}

// GetOrdersResponse The response schema for the `getOrders` operation.
type GetOrdersResponse struct {
	Payload *OrdersList `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetOrdersResponse instantiates a new GetOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrdersResponse() *GetOrdersResponse {
	this := GetOrdersResponse{}
	return &this
}

// NewGetOrdersResponseWithDefaults instantiates a new GetOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrdersResponseWithDefaults() *GetOrdersResponse {
	this := GetOrdersResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetOrdersResponse) GetPayload() OrdersList {
	if o == nil || IsNil(o.Payload) {
		var ret OrdersList
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersResponse) GetPayloadOk() (*OrdersList, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetOrdersResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given OrdersList and assigns it to the Payload field.
func (o *GetOrdersResponse) SetPayload(v OrdersList) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetOrdersResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetOrdersResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetOrdersResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetOrdersResponse struct {
	value *GetOrdersResponse
	isSet bool
}

func (v NullableGetOrdersResponse) Get() *GetOrdersResponse {
	return v.value
}

func (v *NullableGetOrdersResponse) Set(val *GetOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrdersResponse(val *GetOrdersResponse) *NullableGetOrdersResponse {
	return &NullableGetOrdersResponse{value: val, isSet: true}
}

func (v NullableGetOrdersResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
