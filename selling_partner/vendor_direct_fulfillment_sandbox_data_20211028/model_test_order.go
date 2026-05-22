package vendor_direct_fulfillment_sandbox_data_20211028

import (
	"github.com/bytedance/sonic"
)

// checks if the TestOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestOrder{}

// TestOrder Error response returned when the request is unsuccessful.
type TestOrder struct {
	// An error code that identifies the type of error that occurred.
	OrderId string `json:"orderId"`
}

type _TestOrder TestOrder

// NewTestOrder instantiates a new TestOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestOrder(orderId string) *TestOrder {
	this := TestOrder{}
	this.OrderId = orderId
	return &this
}

// NewTestOrderWithDefaults instantiates a new TestOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestOrderWithDefaults() *TestOrder {
	this := TestOrder{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *TestOrder) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *TestOrder) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *TestOrder) SetOrderId(v string) {
	o.OrderId = v
}

func (o TestOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orderId"] = o.OrderId
	return toSerialize, nil
}

type NullableTestOrder struct {
	value *TestOrder
	isSet bool
}

func (v NullableTestOrder) Get() *TestOrder {
	return v.value
}

func (v *NullableTestOrder) Set(val *TestOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableTestOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableTestOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestOrder(val *TestOrder) *NullableTestOrder {
	return &NullableTestOrder{value: val, isSet: true}
}

func (v NullableTestOrder) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTestOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
