package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetAdditionalSellerInputsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAdditionalSellerInputsRequest{}

// GetAdditionalSellerInputsRequest Request schema.
type GetAdditionalSellerInputsRequest struct {
	// An Amazon-defined shipping service identifier.
	ShippingServiceId string  `json:"ShippingServiceId"`
	ShipFromAddress   Address `json:"ShipFromAddress"`
	// An Amazon-defined order identifier, in 3-7-7 format.
	OrderId string `json:"OrderId"`
}

type _GetAdditionalSellerInputsRequest GetAdditionalSellerInputsRequest

// NewGetAdditionalSellerInputsRequest instantiates a new GetAdditionalSellerInputsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAdditionalSellerInputsRequest(shippingServiceId string, shipFromAddress Address, orderId string) *GetAdditionalSellerInputsRequest {
	this := GetAdditionalSellerInputsRequest{}
	this.ShippingServiceId = shippingServiceId
	this.ShipFromAddress = shipFromAddress
	this.OrderId = orderId
	return &this
}

// NewGetAdditionalSellerInputsRequestWithDefaults instantiates a new GetAdditionalSellerInputsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAdditionalSellerInputsRequestWithDefaults() *GetAdditionalSellerInputsRequest {
	this := GetAdditionalSellerInputsRequest{}
	return &this
}

// GetShippingServiceId returns the ShippingServiceId field value
func (o *GetAdditionalSellerInputsRequest) GetShippingServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShippingServiceId
}

// GetShippingServiceIdOk returns a tuple with the ShippingServiceId field value
// and a boolean to check if the value has been set.
func (o *GetAdditionalSellerInputsRequest) GetShippingServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingServiceId, true
}

// SetShippingServiceId sets field value
func (o *GetAdditionalSellerInputsRequest) SetShippingServiceId(v string) {
	o.ShippingServiceId = v
}

// GetShipFromAddress returns the ShipFromAddress field value
func (o *GetAdditionalSellerInputsRequest) GetShipFromAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.ShipFromAddress
}

// GetShipFromAddressOk returns a tuple with the ShipFromAddress field value
// and a boolean to check if the value has been set.
func (o *GetAdditionalSellerInputsRequest) GetShipFromAddressOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipFromAddress, true
}

// SetShipFromAddress sets field value
func (o *GetAdditionalSellerInputsRequest) SetShipFromAddress(v Address) {
	o.ShipFromAddress = v
}

// GetOrderId returns the OrderId field value
func (o *GetAdditionalSellerInputsRequest) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *GetAdditionalSellerInputsRequest) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *GetAdditionalSellerInputsRequest) SetOrderId(v string) {
	o.OrderId = v
}

func (o GetAdditionalSellerInputsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ShippingServiceId"] = o.ShippingServiceId
	toSerialize["ShipFromAddress"] = o.ShipFromAddress
	toSerialize["OrderId"] = o.OrderId
	return toSerialize, nil
}

type NullableGetAdditionalSellerInputsRequest struct {
	value *GetAdditionalSellerInputsRequest
	isSet bool
}

func (v NullableGetAdditionalSellerInputsRequest) Get() *GetAdditionalSellerInputsRequest {
	return v.value
}

func (v *NullableGetAdditionalSellerInputsRequest) Set(val *GetAdditionalSellerInputsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAdditionalSellerInputsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAdditionalSellerInputsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAdditionalSellerInputsRequest(val *GetAdditionalSellerInputsRequest) *NullableGetAdditionalSellerInputsRequest {
	return &NullableGetAdditionalSellerInputsRequest{value: val, isSet: true}
}

func (v NullableGetAdditionalSellerInputsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetAdditionalSellerInputsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
