package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the ShippingPromiseSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingPromiseSet{}

// ShippingPromiseSet The promised delivery time and pickup time.
type ShippingPromiseSet struct {
	DeliveryWindow *TimeRange `json:"deliveryWindow,omitempty"`
	ReceiveWindow  *TimeRange `json:"receiveWindow,omitempty"`
}

// NewShippingPromiseSet instantiates a new ShippingPromiseSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingPromiseSet() *ShippingPromiseSet {
	this := ShippingPromiseSet{}
	return &this
}

// NewShippingPromiseSetWithDefaults instantiates a new ShippingPromiseSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingPromiseSetWithDefaults() *ShippingPromiseSet {
	this := ShippingPromiseSet{}
	return &this
}

// GetDeliveryWindow returns the DeliveryWindow field value if set, zero value otherwise.
func (o *ShippingPromiseSet) GetDeliveryWindow() TimeRange {
	if o == nil || IsNil(o.DeliveryWindow) {
		var ret TimeRange
		return ret
	}
	return *o.DeliveryWindow
}

// GetDeliveryWindowOk returns a tuple with the DeliveryWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingPromiseSet) GetDeliveryWindowOk() (*TimeRange, bool) {
	if o == nil || IsNil(o.DeliveryWindow) {
		return nil, false
	}
	return o.DeliveryWindow, true
}

// HasDeliveryWindow returns a boolean if a field has been set.
func (o *ShippingPromiseSet) HasDeliveryWindow() bool {
	if o != nil && !IsNil(o.DeliveryWindow) {
		return true
	}

	return false
}

// SetDeliveryWindow gets a reference to the given TimeRange and assigns it to the DeliveryWindow field.
func (o *ShippingPromiseSet) SetDeliveryWindow(v TimeRange) {
	o.DeliveryWindow = &v
}

// GetReceiveWindow returns the ReceiveWindow field value if set, zero value otherwise.
func (o *ShippingPromiseSet) GetReceiveWindow() TimeRange {
	if o == nil || IsNil(o.ReceiveWindow) {
		var ret TimeRange
		return ret
	}
	return *o.ReceiveWindow
}

// GetReceiveWindowOk returns a tuple with the ReceiveWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingPromiseSet) GetReceiveWindowOk() (*TimeRange, bool) {
	if o == nil || IsNil(o.ReceiveWindow) {
		return nil, false
	}
	return o.ReceiveWindow, true
}

// HasReceiveWindow returns a boolean if a field has been set.
func (o *ShippingPromiseSet) HasReceiveWindow() bool {
	if o != nil && !IsNil(o.ReceiveWindow) {
		return true
	}

	return false
}

// SetReceiveWindow gets a reference to the given TimeRange and assigns it to the ReceiveWindow field.
func (o *ShippingPromiseSet) SetReceiveWindow(v TimeRange) {
	o.ReceiveWindow = &v
}

func (o ShippingPromiseSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeliveryWindow) {
		toSerialize["deliveryWindow"] = o.DeliveryWindow
	}
	if !IsNil(o.ReceiveWindow) {
		toSerialize["receiveWindow"] = o.ReceiveWindow
	}
	return toSerialize, nil
}

type NullableShippingPromiseSet struct {
	value *ShippingPromiseSet
	isSet bool
}

func (v NullableShippingPromiseSet) Get() *ShippingPromiseSet {
	return v.value
}

func (v *NullableShippingPromiseSet) Set(val *ShippingPromiseSet) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingPromiseSet) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingPromiseSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingPromiseSet(val *ShippingPromiseSet) *NullableShippingPromiseSet {
	return &NullableShippingPromiseSet{value: val, isSet: true}
}

func (v NullableShippingPromiseSet) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShippingPromiseSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
