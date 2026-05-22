package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the InventoryQuantity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryQuantity{}

// InventoryQuantity Quantity of inventory with an associated measurement unit context.
type InventoryQuantity struct {
	// Quantity of the respective inventory.
	Quantity          float32                    `json:"quantity"`
	UnitOfMeasurement InventoryUnitOfMeasurement `json:"unitOfMeasurement"`
}

type _InventoryQuantity InventoryQuantity

// NewInventoryQuantity instantiates a new InventoryQuantity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryQuantity(quantity float32, unitOfMeasurement InventoryUnitOfMeasurement) *InventoryQuantity {
	this := InventoryQuantity{}
	this.Quantity = quantity
	this.UnitOfMeasurement = unitOfMeasurement
	return &this
}

// NewInventoryQuantityWithDefaults instantiates a new InventoryQuantity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryQuantityWithDefaults() *InventoryQuantity {
	this := InventoryQuantity{}
	return &this
}

// GetQuantity returns the Quantity field value
func (o *InventoryQuantity) GetQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *InventoryQuantity) GetQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *InventoryQuantity) SetQuantity(v float32) {
	o.Quantity = v
}

// GetUnitOfMeasurement returns the UnitOfMeasurement field value
func (o *InventoryQuantity) GetUnitOfMeasurement() InventoryUnitOfMeasurement {
	if o == nil {
		var ret InventoryUnitOfMeasurement
		return ret
	}

	return o.UnitOfMeasurement
}

// GetUnitOfMeasurementOk returns a tuple with the UnitOfMeasurement field value
// and a boolean to check if the value has been set.
func (o *InventoryQuantity) GetUnitOfMeasurementOk() (*InventoryUnitOfMeasurement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitOfMeasurement, true
}

// SetUnitOfMeasurement sets field value
func (o *InventoryQuantity) SetUnitOfMeasurement(v InventoryUnitOfMeasurement) {
	o.UnitOfMeasurement = v
}

func (o InventoryQuantity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["quantity"] = o.Quantity
	toSerialize["unitOfMeasurement"] = o.UnitOfMeasurement
	return toSerialize, nil
}

type NullableInventoryQuantity struct {
	value *InventoryQuantity
	isSet bool
}

func (v NullableInventoryQuantity) Get() *InventoryQuantity {
	return v.value
}

func (v *NullableInventoryQuantity) Set(val *InventoryQuantity) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryQuantity) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryQuantity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryQuantity(val *InventoryQuantity) *NullableInventoryQuantity {
	return &NullableInventoryQuantity{value: val, isSet: true}
}

func (v NullableInventoryQuantity) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInventoryQuantity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
