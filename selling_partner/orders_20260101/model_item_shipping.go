package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemShipping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemShipping{}

// ItemShipping Information related to the shipping and delivery process for an order item.
type ItemShipping struct {
	ScheduledDeliveryWindow *DateTimeRange             `json:"scheduledDeliveryWindow,omitempty"`
	ShippingConstraints     *ItemShippingConstraints   `json:"shippingConstraints,omitempty"`
	InternationalShipping   *ItemInternationalShipping `json:"internationalShipping,omitempty"`
}

// NewItemShipping instantiates a new ItemShipping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemShipping() *ItemShipping {
	this := ItemShipping{}
	return &this
}

// NewItemShippingWithDefaults instantiates a new ItemShipping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemShippingWithDefaults() *ItemShipping {
	this := ItemShipping{}
	return &this
}

// GetScheduledDeliveryWindow returns the ScheduledDeliveryWindow field value if set, zero value otherwise.
func (o *ItemShipping) GetScheduledDeliveryWindow() DateTimeRange {
	if o == nil || IsNil(o.ScheduledDeliveryWindow) {
		var ret DateTimeRange
		return ret
	}
	return *o.ScheduledDeliveryWindow
}

// GetScheduledDeliveryWindowOk returns a tuple with the ScheduledDeliveryWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemShipping) GetScheduledDeliveryWindowOk() (*DateTimeRange, bool) {
	if o == nil || IsNil(o.ScheduledDeliveryWindow) {
		return nil, false
	}
	return o.ScheduledDeliveryWindow, true
}

// HasScheduledDeliveryWindow returns a boolean if a field has been set.
func (o *ItemShipping) HasScheduledDeliveryWindow() bool {
	if o != nil && !IsNil(o.ScheduledDeliveryWindow) {
		return true
	}

	return false
}

// SetScheduledDeliveryWindow gets a reference to the given DateTimeRange and assigns it to the ScheduledDeliveryWindow field.
func (o *ItemShipping) SetScheduledDeliveryWindow(v DateTimeRange) {
	o.ScheduledDeliveryWindow = &v
}

// GetShippingConstraints returns the ShippingConstraints field value if set, zero value otherwise.
func (o *ItemShipping) GetShippingConstraints() ItemShippingConstraints {
	if o == nil || IsNil(o.ShippingConstraints) {
		var ret ItemShippingConstraints
		return ret
	}
	return *o.ShippingConstraints
}

// GetShippingConstraintsOk returns a tuple with the ShippingConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemShipping) GetShippingConstraintsOk() (*ItemShippingConstraints, bool) {
	if o == nil || IsNil(o.ShippingConstraints) {
		return nil, false
	}
	return o.ShippingConstraints, true
}

// HasShippingConstraints returns a boolean if a field has been set.
func (o *ItemShipping) HasShippingConstraints() bool {
	if o != nil && !IsNil(o.ShippingConstraints) {
		return true
	}

	return false
}

// SetShippingConstraints gets a reference to the given ItemShippingConstraints and assigns it to the ShippingConstraints field.
func (o *ItemShipping) SetShippingConstraints(v ItemShippingConstraints) {
	o.ShippingConstraints = &v
}

// GetInternationalShipping returns the InternationalShipping field value if set, zero value otherwise.
func (o *ItemShipping) GetInternationalShipping() ItemInternationalShipping {
	if o == nil || IsNil(o.InternationalShipping) {
		var ret ItemInternationalShipping
		return ret
	}
	return *o.InternationalShipping
}

// GetInternationalShippingOk returns a tuple with the InternationalShipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemShipping) GetInternationalShippingOk() (*ItemInternationalShipping, bool) {
	if o == nil || IsNil(o.InternationalShipping) {
		return nil, false
	}
	return o.InternationalShipping, true
}

// HasInternationalShipping returns a boolean if a field has been set.
func (o *ItemShipping) HasInternationalShipping() bool {
	if o != nil && !IsNil(o.InternationalShipping) {
		return true
	}

	return false
}

// SetInternationalShipping gets a reference to the given ItemInternationalShipping and assigns it to the InternationalShipping field.
func (o *ItemShipping) SetInternationalShipping(v ItemInternationalShipping) {
	o.InternationalShipping = &v
}

func (o ItemShipping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScheduledDeliveryWindow) {
		toSerialize["scheduledDeliveryWindow"] = o.ScheduledDeliveryWindow
	}
	if !IsNil(o.ShippingConstraints) {
		toSerialize["shippingConstraints"] = o.ShippingConstraints
	}
	if !IsNil(o.InternationalShipping) {
		toSerialize["internationalShipping"] = o.InternationalShipping
	}
	return toSerialize, nil
}

type NullableItemShipping struct {
	value *ItemShipping
	isSet bool
}

func (v NullableItemShipping) Get() *ItemShipping {
	return v.value
}

func (v *NullableItemShipping) Set(val *ItemShipping) {
	v.value = val
	v.isSet = true
}

func (v NullableItemShipping) IsSet() bool {
	return v.isSet
}

func (v *NullableItemShipping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemShipping(val *ItemShipping) *NullableItemShipping {
	return &NullableItemShipping{value: val, isSet: true}
}

func (v NullableItemShipping) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemShipping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
