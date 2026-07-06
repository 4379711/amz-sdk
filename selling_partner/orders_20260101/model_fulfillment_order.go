package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the FulfillmentOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FulfillmentOrder{}

// FulfillmentOrder Information about a fulfillment order associated with a customer order. A fulfillment order represents a unit of fulfillment created by Amazon for the order. **Note:** Only available for EasyShip orders at present.
type FulfillmentOrder struct {
	// The Fulfillment Order ID assigned by Amazon after fulfillment planning. This identifier is identical to the Shipment ID required by External Fulfillment APIs.
	FulfillmentOrderId string `json:"fulfillmentOrderId"`
}

// NewFulfillmentOrder instantiates a new FulfillmentOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillmentOrder(fulfillmentOrderId string) *FulfillmentOrder {
	this := FulfillmentOrder{}
	this.FulfillmentOrderId = fulfillmentOrderId
	return &this
}

// NewFulfillmentOrderWithDefaults instantiates a new FulfillmentOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillmentOrderWithDefaults() *FulfillmentOrder {
	this := FulfillmentOrder{}
	return &this
}

// GetFulfillmentOrderId returns the FulfillmentOrderId field value
func (o *FulfillmentOrder) GetFulfillmentOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FulfillmentOrderId
}

// GetFulfillmentOrderIdOk returns a tuple with the FulfillmentOrderId field value
// and a boolean to check if the value has been set.
func (o *FulfillmentOrder) GetFulfillmentOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfillmentOrderId, true
}

// SetFulfillmentOrderId sets field value
func (o *FulfillmentOrder) SetFulfillmentOrderId(v string) {
	o.FulfillmentOrderId = v
}

func (o FulfillmentOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fulfillmentOrderId"] = o.FulfillmentOrderId
	return toSerialize, nil
}

type NullableFulfillmentOrder struct {
	value *FulfillmentOrder
	isSet bool
}

func (v NullableFulfillmentOrder) Get() *FulfillmentOrder {
	return v.value
}

func (v *NullableFulfillmentOrder) Set(val *FulfillmentOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentOrder(val *FulfillmentOrder) *NullableFulfillmentOrder {
	return &NullableFulfillmentOrder{value: val, isSet: true}
}

func (v NullableFulfillmentOrder) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFulfillmentOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
