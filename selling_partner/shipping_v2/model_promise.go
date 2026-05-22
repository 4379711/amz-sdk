package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the Promise type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Promise{}

// Promise The time windows promised for pickup and delivery events.
type Promise struct {
	DeliveryWindow *TimeWindow `json:"deliveryWindow,omitempty"`
	PickupWindow   *TimeWindow `json:"pickupWindow,omitempty"`
}

// NewPromise instantiates a new Promise object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromise() *Promise {
	this := Promise{}
	return &this
}

// NewPromiseWithDefaults instantiates a new Promise object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromiseWithDefaults() *Promise {
	this := Promise{}
	return &this
}

// GetDeliveryWindow returns the DeliveryWindow field value if set, zero value otherwise.
func (o *Promise) GetDeliveryWindow() TimeWindow {
	if o == nil || IsNil(o.DeliveryWindow) {
		var ret TimeWindow
		return ret
	}
	return *o.DeliveryWindow
}

// GetDeliveryWindowOk returns a tuple with the DeliveryWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promise) GetDeliveryWindowOk() (*TimeWindow, bool) {
	if o == nil || IsNil(o.DeliveryWindow) {
		return nil, false
	}
	return o.DeliveryWindow, true
}

// HasDeliveryWindow returns a boolean if a field has been set.
func (o *Promise) HasDeliveryWindow() bool {
	if o != nil && !IsNil(o.DeliveryWindow) {
		return true
	}

	return false
}

// SetDeliveryWindow gets a reference to the given TimeWindow and assigns it to the DeliveryWindow field.
func (o *Promise) SetDeliveryWindow(v TimeWindow) {
	o.DeliveryWindow = &v
}

// GetPickupWindow returns the PickupWindow field value if set, zero value otherwise.
func (o *Promise) GetPickupWindow() TimeWindow {
	if o == nil || IsNil(o.PickupWindow) {
		var ret TimeWindow
		return ret
	}
	return *o.PickupWindow
}

// GetPickupWindowOk returns a tuple with the PickupWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promise) GetPickupWindowOk() (*TimeWindow, bool) {
	if o == nil || IsNil(o.PickupWindow) {
		return nil, false
	}
	return o.PickupWindow, true
}

// HasPickupWindow returns a boolean if a field has been set.
func (o *Promise) HasPickupWindow() bool {
	if o != nil && !IsNil(o.PickupWindow) {
		return true
	}

	return false
}

// SetPickupWindow gets a reference to the given TimeWindow and assigns it to the PickupWindow field.
func (o *Promise) SetPickupWindow(v TimeWindow) {
	o.PickupWindow = &v
}

func (o Promise) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeliveryWindow) {
		toSerialize["deliveryWindow"] = o.DeliveryWindow
	}
	if !IsNil(o.PickupWindow) {
		toSerialize["pickupWindow"] = o.PickupWindow
	}
	return toSerialize, nil
}

type NullablePromise struct {
	value *Promise
	isSet bool
}

func (v NullablePromise) Get() *Promise {
	return v.value
}

func (v *NullablePromise) Set(val *Promise) {
	v.value = val
	v.isSet = true
}

func (v NullablePromise) IsSet() bool {
	return v.isSet
}

func (v *NullablePromise) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromise(val *Promise) *NullablePromise {
	return &NullablePromise{value: val, isSet: true}
}

func (v NullablePromise) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePromise) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
