package services

import (
	"github.com/bytedance/sonic"
)

// checks if the RangeCapacity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RangeCapacity{}

// RangeCapacity Range capacity entity where each entry has a capacity type and corresponding slots.
type RangeCapacity struct {
	CapacityType *CapacityType `json:"capacityType,omitempty"`
	// Array of capacity slots in range slot format.
	Slots []RangeSlot `json:"slots,omitempty"`
}

// NewRangeCapacity instantiates a new RangeCapacity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRangeCapacity() *RangeCapacity {
	this := RangeCapacity{}
	return &this
}

// NewRangeCapacityWithDefaults instantiates a new RangeCapacity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeCapacityWithDefaults() *RangeCapacity {
	this := RangeCapacity{}
	return &this
}

// GetCapacityType returns the CapacityType field value if set, zero value otherwise.
func (o *RangeCapacity) GetCapacityType() CapacityType {
	if o == nil || IsNil(o.CapacityType) {
		var ret CapacityType
		return ret
	}
	return *o.CapacityType
}

// GetCapacityTypeOk returns a tuple with the CapacityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeCapacity) GetCapacityTypeOk() (*CapacityType, bool) {
	if o == nil || IsNil(o.CapacityType) {
		return nil, false
	}
	return o.CapacityType, true
}

// HasCapacityType returns a boolean if a field has been set.
func (o *RangeCapacity) HasCapacityType() bool {
	if o != nil && !IsNil(o.CapacityType) {
		return true
	}

	return false
}

// SetCapacityType gets a reference to the given CapacityType and assigns it to the CapacityType field.
func (o *RangeCapacity) SetCapacityType(v CapacityType) {
	o.CapacityType = &v
}

// GetSlots returns the Slots field value if set, zero value otherwise.
func (o *RangeCapacity) GetSlots() []RangeSlot {
	if o == nil || IsNil(o.Slots) {
		var ret []RangeSlot
		return ret
	}
	return o.Slots
}

// GetSlotsOk returns a tuple with the Slots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeCapacity) GetSlotsOk() ([]RangeSlot, bool) {
	if o == nil || IsNil(o.Slots) {
		return nil, false
	}
	return o.Slots, true
}

// HasSlots returns a boolean if a field has been set.
func (o *RangeCapacity) HasSlots() bool {
	if o != nil && !IsNil(o.Slots) {
		return true
	}

	return false
}

// SetSlots gets a reference to the given []RangeSlot and assigns it to the Slots field.
func (o *RangeCapacity) SetSlots(v []RangeSlot) {
	o.Slots = v
}

func (o RangeCapacity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CapacityType) {
		toSerialize["capacityType"] = o.CapacityType
	}
	if !IsNil(o.Slots) {
		toSerialize["slots"] = o.Slots
	}
	return toSerialize, nil
}

type NullableRangeCapacity struct {
	value *RangeCapacity
	isSet bool
}

func (v NullableRangeCapacity) Get() *RangeCapacity {
	return v.value
}

func (v *NullableRangeCapacity) Set(val *RangeCapacity) {
	v.value = val
	v.isSet = true
}

func (v NullableRangeCapacity) IsSet() bool {
	return v.isSet
}

func (v *NullableRangeCapacity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRangeCapacity(val *RangeCapacity) *NullableRangeCapacity {
	return &NullableRangeCapacity{value: val, isSet: true}
}

func (v NullableRangeCapacity) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRangeCapacity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
