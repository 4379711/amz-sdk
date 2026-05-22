package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the PalletInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalletInput{}

// PalletInput Contains input information about a pallet to be used in the inbound plan.
type PalletInput struct {
	Dimensions *Dimensions `json:"dimensions,omitempty"`
	// The number of containers where all other properties like weight or dimensions are identical.
	Quantity     int32         `json:"quantity"`
	Stackability *Stackability `json:"stackability,omitempty"`
	Weight       *Weight       `json:"weight,omitempty"`
}

type _PalletInput PalletInput

// NewPalletInput instantiates a new PalletInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalletInput(quantity int32) *PalletInput {
	this := PalletInput{}
	this.Quantity = quantity
	return &this
}

// NewPalletInputWithDefaults instantiates a new PalletInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalletInputWithDefaults() *PalletInput {
	this := PalletInput{}
	return &this
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *PalletInput) GetDimensions() Dimensions {
	if o == nil || IsNil(o.Dimensions) {
		var ret Dimensions
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PalletInput) GetDimensionsOk() (*Dimensions, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *PalletInput) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given Dimensions and assigns it to the Dimensions field.
func (o *PalletInput) SetDimensions(v Dimensions) {
	o.Dimensions = &v
}

// GetQuantity returns the Quantity field value
func (o *PalletInput) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *PalletInput) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *PalletInput) SetQuantity(v int32) {
	o.Quantity = v
}

// GetStackability returns the Stackability field value if set, zero value otherwise.
func (o *PalletInput) GetStackability() Stackability {
	if o == nil || IsNil(o.Stackability) {
		var ret Stackability
		return ret
	}
	return *o.Stackability
}

// GetStackabilityOk returns a tuple with the Stackability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PalletInput) GetStackabilityOk() (*Stackability, bool) {
	if o == nil || IsNil(o.Stackability) {
		return nil, false
	}
	return o.Stackability, true
}

// HasStackability returns a boolean if a field has been set.
func (o *PalletInput) HasStackability() bool {
	if o != nil && !IsNil(o.Stackability) {
		return true
	}

	return false
}

// SetStackability gets a reference to the given Stackability and assigns it to the Stackability field.
func (o *PalletInput) SetStackability(v Stackability) {
	o.Stackability = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *PalletInput) GetWeight() Weight {
	if o == nil || IsNil(o.Weight) {
		var ret Weight
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PalletInput) GetWeightOk() (*Weight, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *PalletInput) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given Weight and assigns it to the Weight field.
func (o *PalletInput) SetWeight(v Weight) {
	o.Weight = &v
}

func (o PalletInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	toSerialize["quantity"] = o.Quantity
	if !IsNil(o.Stackability) {
		toSerialize["stackability"] = o.Stackability
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullablePalletInput struct {
	value *PalletInput
	isSet bool
}

func (v NullablePalletInput) Get() *PalletInput {
	return v.value
}

func (v *NullablePalletInput) Set(val *PalletInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePalletInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePalletInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalletInput(val *PalletInput) *NullablePalletInput {
	return &NullablePalletInput{value: val, isSet: true}
}

func (v NullablePalletInput) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePalletInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
