package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the GetOrderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrderResponse{}

// GetOrderResponse Order details.
type GetOrderResponse struct {
	Order Order `json:"order"`
}

// NewGetOrderResponse instantiates a new GetOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderResponse(order Order) *GetOrderResponse {
	this := GetOrderResponse{}
	this.Order = order
	return &this
}

// NewGetOrderResponseWithDefaults instantiates a new GetOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderResponseWithDefaults() *GetOrderResponse {
	this := GetOrderResponse{}
	return &this
}

// GetOrder returns the Order field value
func (o *GetOrderResponse) GetOrder() Order {
	if o == nil {
		var ret Order
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetOrderOk() (*Order, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *GetOrderResponse) SetOrder(v Order) {
	o.Order = v
}

func (o GetOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["order"] = o.Order
	return toSerialize, nil
}

type NullableGetOrderResponse struct {
	value *GetOrderResponse
	isSet bool
}

func (v NullableGetOrderResponse) Get() *GetOrderResponse {
	return v.value
}

func (v *NullableGetOrderResponse) Set(val *GetOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderResponse(val *GetOrderResponse) *NullableGetOrderResponse {
	return &NullableGetOrderResponse{value: val, isSet: true}
}

func (v NullableGetOrderResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
