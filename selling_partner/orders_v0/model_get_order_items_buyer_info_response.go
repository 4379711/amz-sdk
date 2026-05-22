package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetOrderItemsBuyerInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrderItemsBuyerInfoResponse{}

// GetOrderItemsBuyerInfoResponse The response schema for the `getOrderItemsBuyerInfo` operation.
type GetOrderItemsBuyerInfoResponse struct {
	Payload *OrderItemsBuyerInfoList `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetOrderItemsBuyerInfoResponse instantiates a new GetOrderItemsBuyerInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderItemsBuyerInfoResponse() *GetOrderItemsBuyerInfoResponse {
	this := GetOrderItemsBuyerInfoResponse{}
	return &this
}

// NewGetOrderItemsBuyerInfoResponseWithDefaults instantiates a new GetOrderItemsBuyerInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderItemsBuyerInfoResponseWithDefaults() *GetOrderItemsBuyerInfoResponse {
	this := GetOrderItemsBuyerInfoResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetOrderItemsBuyerInfoResponse) GetPayload() OrderItemsBuyerInfoList {
	if o == nil || IsNil(o.Payload) {
		var ret OrderItemsBuyerInfoList
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderItemsBuyerInfoResponse) GetPayloadOk() (*OrderItemsBuyerInfoList, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetOrderItemsBuyerInfoResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given OrderItemsBuyerInfoList and assigns it to the Payload field.
func (o *GetOrderItemsBuyerInfoResponse) SetPayload(v OrderItemsBuyerInfoList) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetOrderItemsBuyerInfoResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderItemsBuyerInfoResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetOrderItemsBuyerInfoResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetOrderItemsBuyerInfoResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetOrderItemsBuyerInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetOrderItemsBuyerInfoResponse struct {
	value *GetOrderItemsBuyerInfoResponse
	isSet bool
}

func (v NullableGetOrderItemsBuyerInfoResponse) Get() *GetOrderItemsBuyerInfoResponse {
	return v.value
}

func (v *NullableGetOrderItemsBuyerInfoResponse) Set(val *GetOrderItemsBuyerInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderItemsBuyerInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderItemsBuyerInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderItemsBuyerInfoResponse(val *GetOrderItemsBuyerInfoResponse) *NullableGetOrderItemsBuyerInfoResponse {
	return &NullableGetOrderItemsBuyerInfoResponse{value: val, isSet: true}
}

func (v NullableGetOrderItemsBuyerInfoResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetOrderItemsBuyerInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
