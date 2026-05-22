package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the InventoryDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryDetails{}

// InventoryDetails Additional inventory details. This object is only displayed if the details parameter in the request is set to `SHOW`.
type InventoryDetails struct {
	// Quantity that is available for downstream channel replenishment.
	AvailableDistributableQuantity *int64 `json:"availableDistributableQuantity,omitempty"`
	// Quantity that is in transit from AWD and has not yet been received at FBA.
	ReplenishmentQuantity *int64 `json:"replenishmentQuantity,omitempty"`
	// Quantity that is reserved for a downstream channel replenishment order that is being prepared for shipment.
	ReservedDistributableQuantity *int64 `json:"reservedDistributableQuantity,omitempty"`
}

// NewInventoryDetails instantiates a new InventoryDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryDetails() *InventoryDetails {
	this := InventoryDetails{}
	return &this
}

// NewInventoryDetailsWithDefaults instantiates a new InventoryDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryDetailsWithDefaults() *InventoryDetails {
	this := InventoryDetails{}
	return &this
}

// GetAvailableDistributableQuantity returns the AvailableDistributableQuantity field value if set, zero value otherwise.
func (o *InventoryDetails) GetAvailableDistributableQuantity() int64 {
	if o == nil || IsNil(o.AvailableDistributableQuantity) {
		var ret int64
		return ret
	}
	return *o.AvailableDistributableQuantity
}

// GetAvailableDistributableQuantityOk returns a tuple with the AvailableDistributableQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDetails) GetAvailableDistributableQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.AvailableDistributableQuantity) {
		return nil, false
	}
	return o.AvailableDistributableQuantity, true
}

// HasAvailableDistributableQuantity returns a boolean if a field has been set.
func (o *InventoryDetails) HasAvailableDistributableQuantity() bool {
	if o != nil && !IsNil(o.AvailableDistributableQuantity) {
		return true
	}

	return false
}

// SetAvailableDistributableQuantity gets a reference to the given int64 and assigns it to the AvailableDistributableQuantity field.
func (o *InventoryDetails) SetAvailableDistributableQuantity(v int64) {
	o.AvailableDistributableQuantity = &v
}

// GetReplenishmentQuantity returns the ReplenishmentQuantity field value if set, zero value otherwise.
func (o *InventoryDetails) GetReplenishmentQuantity() int64 {
	if o == nil || IsNil(o.ReplenishmentQuantity) {
		var ret int64
		return ret
	}
	return *o.ReplenishmentQuantity
}

// GetReplenishmentQuantityOk returns a tuple with the ReplenishmentQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDetails) GetReplenishmentQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.ReplenishmentQuantity) {
		return nil, false
	}
	return o.ReplenishmentQuantity, true
}

// HasReplenishmentQuantity returns a boolean if a field has been set.
func (o *InventoryDetails) HasReplenishmentQuantity() bool {
	if o != nil && !IsNil(o.ReplenishmentQuantity) {
		return true
	}

	return false
}

// SetReplenishmentQuantity gets a reference to the given int64 and assigns it to the ReplenishmentQuantity field.
func (o *InventoryDetails) SetReplenishmentQuantity(v int64) {
	o.ReplenishmentQuantity = &v
}

// GetReservedDistributableQuantity returns the ReservedDistributableQuantity field value if set, zero value otherwise.
func (o *InventoryDetails) GetReservedDistributableQuantity() int64 {
	if o == nil || IsNil(o.ReservedDistributableQuantity) {
		var ret int64
		return ret
	}
	return *o.ReservedDistributableQuantity
}

// GetReservedDistributableQuantityOk returns a tuple with the ReservedDistributableQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDetails) GetReservedDistributableQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.ReservedDistributableQuantity) {
		return nil, false
	}
	return o.ReservedDistributableQuantity, true
}

// HasReservedDistributableQuantity returns a boolean if a field has been set.
func (o *InventoryDetails) HasReservedDistributableQuantity() bool {
	if o != nil && !IsNil(o.ReservedDistributableQuantity) {
		return true
	}

	return false
}

// SetReservedDistributableQuantity gets a reference to the given int64 and assigns it to the ReservedDistributableQuantity field.
func (o *InventoryDetails) SetReservedDistributableQuantity(v int64) {
	o.ReservedDistributableQuantity = &v
}

func (o InventoryDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvailableDistributableQuantity) {
		toSerialize["availableDistributableQuantity"] = o.AvailableDistributableQuantity
	}
	if !IsNil(o.ReplenishmentQuantity) {
		toSerialize["replenishmentQuantity"] = o.ReplenishmentQuantity
	}
	if !IsNil(o.ReservedDistributableQuantity) {
		toSerialize["reservedDistributableQuantity"] = o.ReservedDistributableQuantity
	}
	return toSerialize, nil
}

type NullableInventoryDetails struct {
	value *InventoryDetails
	isSet bool
}

func (v NullableInventoryDetails) Get() *InventoryDetails {
	return v.value
}

func (v *NullableInventoryDetails) Set(val *InventoryDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryDetails(val *InventoryDetails) *NullableInventoryDetails {
	return &NullableInventoryDetails{value: val, isSet: true}
}

func (v NullableInventoryDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInventoryDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
