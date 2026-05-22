package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the ImpactMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImpactMetrics{}

// ImpactMetrics For the CONVERSION_OPPORTUNITIES theme, the impact metrics are weekly clicks and orders received for similar products. For other event-based themes, the impact metrics are clicks and orders received for similar products during the event days. <br> Note: This object is nullable
type ImpactMetrics struct {
	Clicks NullableImpactMetric `json:"clicks,omitempty"`
	Orders NullableImpactMetric `json:"orders,omitempty"`
}

// NewImpactMetrics instantiates a new ImpactMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImpactMetrics() *ImpactMetrics {
	this := ImpactMetrics{}
	return &this
}

// NewImpactMetricsWithDefaults instantiates a new ImpactMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImpactMetricsWithDefaults() *ImpactMetrics {
	this := ImpactMetrics{}
	return &this
}

// GetClicks returns the Clicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImpactMetrics) GetClicks() ImpactMetric {
	if o == nil || IsNil(o.Clicks.Get()) {
		var ret ImpactMetric
		return ret
	}
	return *o.Clicks.Get()
}

// GetClicksOk returns a tuple with the Clicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImpactMetrics) GetClicksOk() (*ImpactMetric, bool) {
	if o == nil {
		return nil, false
	}
	return o.Clicks.Get(), o.Clicks.IsSet()
}

// HasClicks returns a boolean if a field has been set.
func (o *ImpactMetrics) HasClicks() bool {
	if o != nil && o.Clicks.IsSet() {
		return true
	}

	return false
}

// SetClicks gets a reference to the given NullableImpactMetric and assigns it to the Clicks field.
func (o *ImpactMetrics) SetClicks(v ImpactMetric) {
	o.Clicks.Set(&v)
}

// SetClicksNil sets the value for Clicks to be an explicit nil
func (o *ImpactMetrics) SetClicksNil() {
	o.Clicks.Set(nil)
}

// UnsetClicks ensures that no value is present for Clicks, not even an explicit nil
func (o *ImpactMetrics) UnsetClicks() {
	o.Clicks.Unset()
}

// GetOrders returns the Orders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImpactMetrics) GetOrders() ImpactMetric {
	if o == nil || IsNil(o.Orders.Get()) {
		var ret ImpactMetric
		return ret
	}
	return *o.Orders.Get()
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImpactMetrics) GetOrdersOk() (*ImpactMetric, bool) {
	if o == nil {
		return nil, false
	}
	return o.Orders.Get(), o.Orders.IsSet()
}

// HasOrders returns a boolean if a field has been set.
func (o *ImpactMetrics) HasOrders() bool {
	if o != nil && o.Orders.IsSet() {
		return true
	}

	return false
}

// SetOrders gets a reference to the given NullableImpactMetric and assigns it to the Orders field.
func (o *ImpactMetrics) SetOrders(v ImpactMetric) {
	o.Orders.Set(&v)
}

// SetOrdersNil sets the value for Orders to be an explicit nil
func (o *ImpactMetrics) SetOrdersNil() {
	o.Orders.Set(nil)
}

// UnsetOrders ensures that no value is present for Orders, not even an explicit nil
func (o *ImpactMetrics) UnsetOrders() {
	o.Orders.Unset()
}

func (o ImpactMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Clicks.IsSet() {
		toSerialize["clicks"] = o.Clicks.Get()
	}
	if o.Orders.IsSet() {
		toSerialize["orders"] = o.Orders.Get()
	}
	return toSerialize, nil
}

type NullableImpactMetrics struct {
	value *ImpactMetrics
	isSet bool
}

func (v NullableImpactMetrics) Get() *ImpactMetrics {
	return v.value
}

func (v *NullableImpactMetrics) Set(val *ImpactMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableImpactMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableImpactMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImpactMetrics(val *ImpactMetrics) *NullableImpactMetrics {
	return &NullableImpactMetrics{value: val, isSet: true}
}

func (v NullableImpactMetrics) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableImpactMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
