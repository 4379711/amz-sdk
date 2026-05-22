package vendor_direct_fulfillment_orders_v1

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the ShipmentDates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentDates{}

// ShipmentDates Shipment dates.
type ShipmentDates struct {
	// Time by which the vendor is required to ship the order.
	RequiredShipDate time.Time `json:"requiredShipDate"`
	// Delivery date promised to the Amazon customer.
	PromisedDeliveryDate *time.Time `json:"promisedDeliveryDate,omitempty"`
}

type _ShipmentDates ShipmentDates

// NewShipmentDates instantiates a new ShipmentDates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentDates(requiredShipDate time.Time) *ShipmentDates {
	this := ShipmentDates{}
	this.RequiredShipDate = requiredShipDate
	return &this
}

// NewShipmentDatesWithDefaults instantiates a new ShipmentDates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentDatesWithDefaults() *ShipmentDates {
	this := ShipmentDates{}
	return &this
}

// GetRequiredShipDate returns the RequiredShipDate field value
func (o *ShipmentDates) GetRequiredShipDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.RequiredShipDate
}

// GetRequiredShipDateOk returns a tuple with the RequiredShipDate field value
// and a boolean to check if the value has been set.
func (o *ShipmentDates) GetRequiredShipDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequiredShipDate, true
}

// SetRequiredShipDate sets field value
func (o *ShipmentDates) SetRequiredShipDate(v time.Time) {
	o.RequiredShipDate = v
}

// GetPromisedDeliveryDate returns the PromisedDeliveryDate field value if set, zero value otherwise.
func (o *ShipmentDates) GetPromisedDeliveryDate() time.Time {
	if o == nil || IsNil(o.PromisedDeliveryDate) {
		var ret time.Time
		return ret
	}
	return *o.PromisedDeliveryDate
}

// GetPromisedDeliveryDateOk returns a tuple with the PromisedDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDates) GetPromisedDeliveryDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PromisedDeliveryDate) {
		return nil, false
	}
	return o.PromisedDeliveryDate, true
}

// HasPromisedDeliveryDate returns a boolean if a field has been set.
func (o *ShipmentDates) HasPromisedDeliveryDate() bool {
	if o != nil && !IsNil(o.PromisedDeliveryDate) {
		return true
	}

	return false
}

// SetPromisedDeliveryDate gets a reference to the given time.Time and assigns it to the PromisedDeliveryDate field.
func (o *ShipmentDates) SetPromisedDeliveryDate(v time.Time) {
	o.PromisedDeliveryDate = &v
}

func (o ShipmentDates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requiredShipDate"] = o.RequiredShipDate
	if !IsNil(o.PromisedDeliveryDate) {
		toSerialize["promisedDeliveryDate"] = o.PromisedDeliveryDate
	}
	return toSerialize, nil
}

type NullableShipmentDates struct {
	value *ShipmentDates
	isSet bool
}

func (v NullableShipmentDates) Get() *ShipmentDates {
	return v.value
}

func (v *NullableShipmentDates) Set(val *ShipmentDates) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentDates) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentDates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentDates(val *ShipmentDates) *NullableShipmentDates {
	return &NullableShipmentDates{value: val, isSet: true}
}

func (v NullableShipmentDates) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShipmentDates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
