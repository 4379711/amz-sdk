package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemPacking type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemPacking{}

// ItemPacking Information related to the packaging process for an order item.
type ItemPacking struct {
	GiftOption              *GiftOption              `json:"giftOption,omitempty"`
	SerialNumberRequirement *SerialNumberRequirement `json:"serialNumberRequirement,omitempty"`
}

// NewItemPacking instantiates a new ItemPacking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemPacking() *ItemPacking {
	this := ItemPacking{}
	return &this
}

// NewItemPackingWithDefaults instantiates a new ItemPacking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemPackingWithDefaults() *ItemPacking {
	this := ItemPacking{}
	return &this
}

// GetGiftOption returns the GiftOption field value if set, zero value otherwise.
func (o *ItemPacking) GetGiftOption() GiftOption {
	if o == nil || IsNil(o.GiftOption) {
		var ret GiftOption
		return ret
	}
	return *o.GiftOption
}

// GetGiftOptionOk returns a tuple with the GiftOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemPacking) GetGiftOptionOk() (*GiftOption, bool) {
	if o == nil || IsNil(o.GiftOption) {
		return nil, false
	}
	return o.GiftOption, true
}

// HasGiftOption returns a boolean if a field has been set.
func (o *ItemPacking) HasGiftOption() bool {
	if o != nil && !IsNil(o.GiftOption) {
		return true
	}

	return false
}

// SetGiftOption gets a reference to the given GiftOption and assigns it to the GiftOption field.
func (o *ItemPacking) SetGiftOption(v GiftOption) {
	o.GiftOption = &v
}

// GetSerialNumberRequirement returns the SerialNumberRequirement field value if set, zero value otherwise.
func (o *ItemPacking) GetSerialNumberRequirement() SerialNumberRequirement {
	if o == nil || IsNil(o.SerialNumberRequirement) {
		var ret SerialNumberRequirement
		return ret
	}
	return *o.SerialNumberRequirement
}

// GetSerialNumberRequirementOk returns a tuple with the SerialNumberRequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemPacking) GetSerialNumberRequirementOk() (*SerialNumberRequirement, bool) {
	if o == nil || IsNil(o.SerialNumberRequirement) {
		return nil, false
	}
	return o.SerialNumberRequirement, true
}

// HasSerialNumberRequirement returns a boolean if a field has been set.
func (o *ItemPacking) HasSerialNumberRequirement() bool {
	if o != nil && !IsNil(o.SerialNumberRequirement) {
		return true
	}

	return false
}

// SetSerialNumberRequirement gets a reference to the given SerialNumberRequirement and assigns it to the SerialNumberRequirement field.
func (o *ItemPacking) SetSerialNumberRequirement(v SerialNumberRequirement) {
	o.SerialNumberRequirement = &v
}

func (o ItemPacking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GiftOption) {
		toSerialize["giftOption"] = o.GiftOption
	}
	if !IsNil(o.SerialNumberRequirement) {
		toSerialize["serialNumberRequirement"] = o.SerialNumberRequirement
	}
	return toSerialize, nil
}

type NullableItemPacking struct {
	value *ItemPacking
	isSet bool
}

func (v NullableItemPacking) Get() *ItemPacking {
	return v.value
}

func (v *NullableItemPacking) Set(val *ItemPacking) {
	v.value = val
	v.isSet = true
}

func (v NullableItemPacking) IsSet() bool {
	return v.isSet
}

func (v *NullableItemPacking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemPacking(val *ItemPacking) *NullableItemPacking {
	return &NullableItemPacking{value: val, isSet: true}
}

func (v NullableItemPacking) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemPacking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
