package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemExpense type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemExpense{}

// ItemExpense The expense information related to this specific item.
type ItemExpense struct {
	PointsCost *ItemPointsCost `json:"pointsCost,omitempty"`
}

// NewItemExpense instantiates a new ItemExpense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemExpense() *ItemExpense {
	this := ItemExpense{}
	return &this
}

// NewItemExpenseWithDefaults instantiates a new ItemExpense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemExpenseWithDefaults() *ItemExpense {
	this := ItemExpense{}
	return &this
}

// GetPointsCost returns the PointsCost field value if set, zero value otherwise.
func (o *ItemExpense) GetPointsCost() ItemPointsCost {
	if o == nil || IsNil(o.PointsCost) {
		var ret ItemPointsCost
		return ret
	}
	return *o.PointsCost
}

// GetPointsCostOk returns a tuple with the PointsCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemExpense) GetPointsCostOk() (*ItemPointsCost, bool) {
	if o == nil || IsNil(o.PointsCost) {
		return nil, false
	}
	return o.PointsCost, true
}

// HasPointsCost returns a boolean if a field has been set.
func (o *ItemExpense) HasPointsCost() bool {
	if o != nil && !IsNil(o.PointsCost) {
		return true
	}

	return false
}

// SetPointsCost gets a reference to the given ItemPointsCost and assigns it to the PointsCost field.
func (o *ItemExpense) SetPointsCost(v ItemPointsCost) {
	o.PointsCost = &v
}

func (o ItemExpense) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PointsCost) {
		toSerialize["pointsCost"] = o.PointsCost
	}
	return toSerialize, nil
}

type NullableItemExpense struct {
	value *ItemExpense
	isSet bool
}

func (v NullableItemExpense) Get() *ItemExpense {
	return v.value
}

func (v *NullableItemExpense) Set(val *ItemExpense) {
	v.value = val
	v.isSet = true
}

func (v NullableItemExpense) IsSet() bool {
	return v.isSet
}

func (v *NullableItemExpense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemExpense(val *ItemExpense) *NullableItemExpense {
	return &NullableItemExpense{value: val, isSet: true}
}

func (v NullableItemExpense) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemExpense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
