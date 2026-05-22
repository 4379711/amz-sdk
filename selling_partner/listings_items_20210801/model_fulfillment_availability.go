package listings_items_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the FulfillmentAvailability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FulfillmentAvailability{}

// FulfillmentAvailability The fulfillment availability details for the listings item.
type FulfillmentAvailability struct {
	// The code of the fulfillment network that will be used.
	FulfillmentChannelCode string `json:"fulfillmentChannelCode"`
	// The quantity of the item you are making available for sale.
	Quantity *int32 `json:"quantity,omitempty"`
}

type _FulfillmentAvailability FulfillmentAvailability

// NewFulfillmentAvailability instantiates a new FulfillmentAvailability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillmentAvailability(fulfillmentChannelCode string) *FulfillmentAvailability {
	this := FulfillmentAvailability{}
	this.FulfillmentChannelCode = fulfillmentChannelCode
	return &this
}

// NewFulfillmentAvailabilityWithDefaults instantiates a new FulfillmentAvailability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillmentAvailabilityWithDefaults() *FulfillmentAvailability {
	this := FulfillmentAvailability{}
	return &this
}

// GetFulfillmentChannelCode returns the FulfillmentChannelCode field value
func (o *FulfillmentAvailability) GetFulfillmentChannelCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FulfillmentChannelCode
}

// GetFulfillmentChannelCodeOk returns a tuple with the FulfillmentChannelCode field value
// and a boolean to check if the value has been set.
func (o *FulfillmentAvailability) GetFulfillmentChannelCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfillmentChannelCode, true
}

// SetFulfillmentChannelCode sets field value
func (o *FulfillmentAvailability) SetFulfillmentChannelCode(v string) {
	o.FulfillmentChannelCode = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *FulfillmentAvailability) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentAvailability) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *FulfillmentAvailability) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *FulfillmentAvailability) SetQuantity(v int32) {
	o.Quantity = &v
}

func (o FulfillmentAvailability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fulfillmentChannelCode"] = o.FulfillmentChannelCode
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableFulfillmentAvailability struct {
	value *FulfillmentAvailability
	isSet bool
}

func (v NullableFulfillmentAvailability) Get() *FulfillmentAvailability {
	return v.value
}

func (v *NullableFulfillmentAvailability) Set(val *FulfillmentAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentAvailability(val *FulfillmentAvailability) *NullableFulfillmentAvailability {
	return &NullableFulfillmentAvailability{value: val, isSet: true}
}

func (v NullableFulfillmentAvailability) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFulfillmentAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
