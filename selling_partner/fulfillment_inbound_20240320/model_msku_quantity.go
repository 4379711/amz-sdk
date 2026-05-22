package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the MskuQuantity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MskuQuantity{}

// MskuQuantity Represents an MSKU and the related quantity.
type MskuQuantity struct {
	// The merchant SKU, a merchant-supplied identifier for a specific SKU.
	Msku string `json:"msku"`
	// A positive integer.
	Quantity int32 `json:"quantity"`
}

type _MskuQuantity MskuQuantity

// NewMskuQuantity instantiates a new MskuQuantity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMskuQuantity(msku string, quantity int32) *MskuQuantity {
	this := MskuQuantity{}
	this.Msku = msku
	this.Quantity = quantity
	return &this
}

// NewMskuQuantityWithDefaults instantiates a new MskuQuantity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMskuQuantityWithDefaults() *MskuQuantity {
	this := MskuQuantity{}
	return &this
}

// GetMsku returns the Msku field value
func (o *MskuQuantity) GetMsku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msku
}

// GetMskuOk returns a tuple with the Msku field value
// and a boolean to check if the value has been set.
func (o *MskuQuantity) GetMskuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msku, true
}

// SetMsku sets field value
func (o *MskuQuantity) SetMsku(v string) {
	o.Msku = v
}

// GetQuantity returns the Quantity field value
func (o *MskuQuantity) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *MskuQuantity) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *MskuQuantity) SetQuantity(v int32) {
	o.Quantity = v
}

func (o MskuQuantity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["msku"] = o.Msku
	toSerialize["quantity"] = o.Quantity
	return toSerialize, nil
}

type NullableMskuQuantity struct {
	value *MskuQuantity
	isSet bool
}

func (v NullableMskuQuantity) Get() *MskuQuantity {
	return v.value
}

func (v *NullableMskuQuantity) Set(val *MskuQuantity) {
	v.value = val
	v.isSet = true
}

func (v NullableMskuQuantity) IsSet() bool {
	return v.isSet
}

func (v *NullableMskuQuantity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMskuQuantity(val *MskuQuantity) *NullableMskuQuantity {
	return &NullableMskuQuantity{value: val, isSet: true}
}

func (v NullableMskuQuantity) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMskuQuantity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
