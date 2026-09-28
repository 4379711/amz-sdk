package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the QuantityThreshold type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuantityThreshold{}

// QuantityThreshold Quantity-based purchase conditions for basket building promotions.
type QuantityThreshold struct {
	// How the quantity requirement is evaluated.
	Type string `json:"type"`
	// The number of items required to meet this purchase condition.
	Quantity int32 `json:"quantity"`
}

// NewQuantityThreshold instantiates a new QuantityThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuantityThreshold(type_ string, quantity int32) *QuantityThreshold {
	this := QuantityThreshold{}
	this.Type = type_
	this.Quantity = quantity
	return &this
}

// NewQuantityThresholdWithDefaults instantiates a new QuantityThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuantityThresholdWithDefaults() *QuantityThreshold {
	this := QuantityThreshold{}
	return &this
}

// GetType returns the Type field value
func (o *QuantityThreshold) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *QuantityThreshold) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *QuantityThreshold) SetType(v string) {
	o.Type = v
}

// GetQuantity returns the Quantity field value
func (o *QuantityThreshold) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *QuantityThreshold) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *QuantityThreshold) SetQuantity(v int32) {
	o.Quantity = v
}

func (o QuantityThreshold) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["quantity"] = o.Quantity
	return toSerialize, nil
}

type NullableQuantityThreshold struct {
	value *QuantityThreshold
	isSet bool
}

func (v NullableQuantityThreshold) Get() *QuantityThreshold {
	return v.value
}

func (v *NullableQuantityThreshold) Set(val *QuantityThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableQuantityThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableQuantityThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuantityThreshold(val *QuantityThreshold) *NullableQuantityThreshold {
	return &NullableQuantityThreshold{value: val, isSet: true}
}

func (v NullableQuantityThreshold) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableQuantityThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
