package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the SkuQuantity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuQuantity{}

// SkuQuantity Quantity details for a SKU as part of a shipment
type SkuQuantity struct {
	ExpectedQuantity InventoryQuantity  `json:"expectedQuantity"`
	ReceivedQuantity *InventoryQuantity `json:"receivedQuantity,omitempty"`
	// The merchant stock keeping unit
	Sku string `json:"sku"`
}

type _SkuQuantity SkuQuantity

// NewSkuQuantity instantiates a new SkuQuantity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuQuantity(expectedQuantity InventoryQuantity, sku string) *SkuQuantity {
	this := SkuQuantity{}
	this.ExpectedQuantity = expectedQuantity
	this.Sku = sku
	return &this
}

// NewSkuQuantityWithDefaults instantiates a new SkuQuantity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuQuantityWithDefaults() *SkuQuantity {
	this := SkuQuantity{}
	return &this
}

// GetExpectedQuantity returns the ExpectedQuantity field value
func (o *SkuQuantity) GetExpectedQuantity() InventoryQuantity {
	if o == nil {
		var ret InventoryQuantity
		return ret
	}

	return o.ExpectedQuantity
}

// GetExpectedQuantityOk returns a tuple with the ExpectedQuantity field value
// and a boolean to check if the value has been set.
func (o *SkuQuantity) GetExpectedQuantityOk() (*InventoryQuantity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpectedQuantity, true
}

// SetExpectedQuantity sets field value
func (o *SkuQuantity) SetExpectedQuantity(v InventoryQuantity) {
	o.ExpectedQuantity = v
}

// GetReceivedQuantity returns the ReceivedQuantity field value if set, zero value otherwise.
func (o *SkuQuantity) GetReceivedQuantity() InventoryQuantity {
	if o == nil || IsNil(o.ReceivedQuantity) {
		var ret InventoryQuantity
		return ret
	}
	return *o.ReceivedQuantity
}

// GetReceivedQuantityOk returns a tuple with the ReceivedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuQuantity) GetReceivedQuantityOk() (*InventoryQuantity, bool) {
	if o == nil || IsNil(o.ReceivedQuantity) {
		return nil, false
	}
	return o.ReceivedQuantity, true
}

// HasReceivedQuantity returns a boolean if a field has been set.
func (o *SkuQuantity) HasReceivedQuantity() bool {
	if o != nil && !IsNil(o.ReceivedQuantity) {
		return true
	}

	return false
}

// SetReceivedQuantity gets a reference to the given InventoryQuantity and assigns it to the ReceivedQuantity field.
func (o *SkuQuantity) SetReceivedQuantity(v InventoryQuantity) {
	o.ReceivedQuantity = &v
}

// GetSku returns the Sku field value
func (o *SkuQuantity) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *SkuQuantity) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *SkuQuantity) SetSku(v string) {
	o.Sku = v
}

func (o SkuQuantity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expectedQuantity"] = o.ExpectedQuantity
	if !IsNil(o.ReceivedQuantity) {
		toSerialize["receivedQuantity"] = o.ReceivedQuantity
	}
	toSerialize["sku"] = o.Sku
	return toSerialize, nil
}

type NullableSkuQuantity struct {
	value *SkuQuantity
	isSet bool
}

func (v NullableSkuQuantity) Get() *SkuQuantity {
	return v.value
}

func (v *NullableSkuQuantity) Set(val *SkuQuantity) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuQuantity) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuQuantity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuQuantity(val *SkuQuantity) *NullableSkuQuantity {
	return &NullableSkuQuantity{value: val, isSet: true}
}

func (v NullableSkuQuantity) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSkuQuantity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
