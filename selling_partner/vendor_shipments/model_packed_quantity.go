package vendor_shipments

import (
	"github.com/bytedance/sonic"
)

// checks if the PackedQuantity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackedQuantity{}

// PackedQuantity Details of item quantity.
type PackedQuantity struct {
	// Amount of units shipped for a specific item at a shipment level. If the item is present only in certain cartons or pallets within the shipment, please provide this at the appropriate carton or pallet level.
	Amount int32 `json:"amount"`
	// Unit of measure for the shipped quantity.
	UnitOfMeasure string `json:"unitOfMeasure"`
	// The case size, in the event that we ordered using cases. Otherwise, 1.
	UnitSize *int32 `json:"unitSize,omitempty"`
}

type _PackedQuantity PackedQuantity

// NewPackedQuantity instantiates a new PackedQuantity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackedQuantity(amount int32, unitOfMeasure string) *PackedQuantity {
	this := PackedQuantity{}
	this.Amount = amount
	this.UnitOfMeasure = unitOfMeasure
	return &this
}

// NewPackedQuantityWithDefaults instantiates a new PackedQuantity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackedQuantityWithDefaults() *PackedQuantity {
	this := PackedQuantity{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PackedQuantity) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PackedQuantity) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PackedQuantity) SetAmount(v int32) {
	o.Amount = v
}

// GetUnitOfMeasure returns the UnitOfMeasure field value
func (o *PackedQuantity) GetUnitOfMeasure() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value
// and a boolean to check if the value has been set.
func (o *PackedQuantity) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitOfMeasure, true
}

// SetUnitOfMeasure sets field value
func (o *PackedQuantity) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = v
}

// GetUnitSize returns the UnitSize field value if set, zero value otherwise.
func (o *PackedQuantity) GetUnitSize() int32 {
	if o == nil || IsNil(o.UnitSize) {
		var ret int32
		return ret
	}
	return *o.UnitSize
}

// GetUnitSizeOk returns a tuple with the UnitSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackedQuantity) GetUnitSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.UnitSize) {
		return nil, false
	}
	return o.UnitSize, true
}

// HasUnitSize returns a boolean if a field has been set.
func (o *PackedQuantity) HasUnitSize() bool {
	if o != nil && !IsNil(o.UnitSize) {
		return true
	}

	return false
}

// SetUnitSize gets a reference to the given int32 and assigns it to the UnitSize field.
func (o *PackedQuantity) SetUnitSize(v int32) {
	o.UnitSize = &v
}

func (o PackedQuantity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["unitOfMeasure"] = o.UnitOfMeasure
	if !IsNil(o.UnitSize) {
		toSerialize["unitSize"] = o.UnitSize
	}
	return toSerialize, nil
}

type NullablePackedQuantity struct {
	value *PackedQuantity
	isSet bool
}

func (v NullablePackedQuantity) Get() *PackedQuantity {
	return v.value
}

func (v *NullablePackedQuantity) Set(val *PackedQuantity) {
	v.value = val
	v.isSet = true
}

func (v NullablePackedQuantity) IsSet() bool {
	return v.isSet
}

func (v *NullablePackedQuantity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackedQuantity(val *PackedQuantity) *NullablePackedQuantity {
	return &NullablePackedQuantity{value: val, isSet: true}
}

func (v NullablePackedQuantity) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackedQuantity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
