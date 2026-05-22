package vendor_direct_fulfillment_payments_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemQuantity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemQuantity{}

// ItemQuantity Details of item quantity.
type ItemQuantity struct {
	// Quantity of units available for a specific item.
	Amount int32 `json:"amount"`
	// Unit of measure for the available quantity.
	UnitOfMeasure string `json:"unitOfMeasure"`
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

func (o ItemQuantity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["unitOfMeasure"] = o.UnitOfMeasure
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
