package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the AmazonOrderDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmazonOrderDetails{}

// AmazonOrderDetails Amazon order information. This is required if the shipment source channel is Amazon.
type AmazonOrderDetails struct {
	// The Amazon order ID associated with the Amazon order fulfilled by this shipment.
	OrderId string `json:"orderId"`
}

type _AmazonOrderDetails AmazonOrderDetails

// NewAmazonOrderDetails instantiates a new AmazonOrderDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonOrderDetails(orderId string) *AmazonOrderDetails {
	this := AmazonOrderDetails{}
	this.OrderId = orderId
	return &this
}

// NewAmazonOrderDetailsWithDefaults instantiates a new AmazonOrderDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonOrderDetailsWithDefaults() *AmazonOrderDetails {
	this := AmazonOrderDetails{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *AmazonOrderDetails) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *AmazonOrderDetails) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *AmazonOrderDetails) SetOrderId(v string) {
	o.OrderId = v
}

func (o AmazonOrderDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orderId"] = o.OrderId
	return toSerialize, nil
}

type NullableAmazonOrderDetails struct {
	value *AmazonOrderDetails
	isSet bool
}

func (v NullableAmazonOrderDetails) Get() *AmazonOrderDetails {
	return v.value
}

func (v *NullableAmazonOrderDetails) Set(val *AmazonOrderDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonOrderDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonOrderDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonOrderDetails(val *AmazonOrderDetails) *NullableAmazonOrderDetails {
	return &NullableAmazonOrderDetails{value: val, isSet: true}
}

func (v NullableAmazonOrderDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAmazonOrderDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
