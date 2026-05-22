package vendor_invoices

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemQuantity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemQuantity{}

// ItemQuantity Details of quantity.
type ItemQuantity struct {
	// Quantity of an item. This value should not be zero.
	Amount int32 `json:"amount"`
	// Unit of measure for the quantity.
	UnitOfMeasure string `json:"unitOfMeasure"`
	// The case size, if the unit of measure value is Cases.
	UnitSize    *int32       `json:"unitSize,omitempty"`
	TotalWeight *TotalWeight `json:"totalWeight,omitempty"`
}

type _ItemQuantity ItemQuantity

// NewItemQuantity instantiates a new ItemQuantity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemQuantity(amount int32, unitOfMeasure string) *ItemQuantity {
	this := ItemQuantity{}
	this.Amount = amount
	this.UnitOfMeasure = unitOfMeasure
	return &this
}

// NewItemQuantityWithDefaults instantiates a new ItemQuantity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemQuantityWithDefaults() *ItemQuantity {
	this := ItemQuantity{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ItemQuantity) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ItemQuantity) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ItemQuantity) SetAmount(v int32) {
	o.Amount = v
}

// GetUnitOfMeasure returns the UnitOfMeasure field value
func (o *ItemQuantity) GetUnitOfMeasure() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value
// and a boolean to check if the value has been set.
func (o *ItemQuantity) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitOfMeasure, true
}

// SetUnitOfMeasure sets field value
func (o *ItemQuantity) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = v
}

// GetUnitSize returns the UnitSize field value if set, zero value otherwise.
func (o *ItemQuantity) GetUnitSize() int32 {
	if o == nil || IsNil(o.UnitSize) {
		var ret int32
		return ret
	}
	return *o.UnitSize
}

// GetUnitSizeOk returns a tuple with the UnitSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemQuantity) GetUnitSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.UnitSize) {
		return nil, false
	}
	return o.UnitSize, true
}

// HasUnitSize returns a boolean if a field has been set.
func (o *ItemQuantity) HasUnitSize() bool {
	if o != nil && !IsNil(o.UnitSize) {
		return true
	}

	return false
}

// SetUnitSize gets a reference to the given int32 and assigns it to the UnitSize field.
func (o *ItemQuantity) SetUnitSize(v int32) {
	o.UnitSize = &v
}

// GetTotalWeight returns the TotalWeight field value if set, zero value otherwise.
func (o *ItemQuantity) GetTotalWeight() TotalWeight {
	if o == nil || IsNil(o.TotalWeight) {
		var ret TotalWeight
		return ret
	}
	return *o.TotalWeight
}

// GetTotalWeightOk returns a tuple with the TotalWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemQuantity) GetTotalWeightOk() (*TotalWeight, bool) {
	if o == nil || IsNil(o.TotalWeight) {
		return nil, false
	}
	return o.TotalWeight, true
}

// HasTotalWeight returns a boolean if a field has been set.
func (o *ItemQuantity) HasTotalWeight() bool {
	if o != nil && !IsNil(o.TotalWeight) {
		return true
	}

	return false
}

// SetTotalWeight gets a reference to the given TotalWeight and assigns it to the TotalWeight field.
func (o *ItemQuantity) SetTotalWeight(v TotalWeight) {
	o.TotalWeight = &v
}

func (o ItemQuantity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["unitOfMeasure"] = o.UnitOfMeasure
	if !IsNil(o.UnitSize) {
		toSerialize["unitSize"] = o.UnitSize
	}
	if !IsNil(o.TotalWeight) {
		toSerialize["totalWeight"] = o.TotalWeight
	}
	return toSerialize, nil
}

type NullableItemQuantity struct {
	value *ItemQuantity
	isSet bool
}

func (v NullableItemQuantity) Get() *ItemQuantity {
	return v.value
}

func (v *NullableItemQuantity) Set(val *ItemQuantity) {
	v.value = val
	v.isSet = true
}

func (v NullableItemQuantity) IsSet() bool {
	return v.isSet
}

func (v *NullableItemQuantity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemQuantity(val *ItemQuantity) *NullableItemQuantity {
	return &NullableItemQuantity{value: val, isSet: true}
}

func (v NullableItemQuantity) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemQuantity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
