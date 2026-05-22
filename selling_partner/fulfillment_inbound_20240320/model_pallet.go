package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the Pallet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pallet{}

// Pallet Contains information about a pallet that is used in the inbound plan. The pallet is a container that holds multiple items or boxes.
type Pallet struct {
	Dimensions *Dimensions `json:"dimensions,omitempty"`
	// Primary key to uniquely identify a Package (Box or Pallet).
	PackageId string `json:"packageId" validate:"regexp=^[a-zA-Z0-9-]*$"`
	// The number of containers where all other properties like weight or dimensions are identical.
	Quantity     *int32        `json:"quantity,omitempty"`
	Stackability *Stackability `json:"stackability,omitempty"`
	Weight       *Weight       `json:"weight,omitempty"`
}

type _Pallet Pallet

// NewPallet instantiates a new Pallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPallet(packageId string) *Pallet {
	this := Pallet{}
	this.PackageId = packageId
	return &this
}

// NewPalletWithDefaults instantiates a new Pallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalletWithDefaults() *Pallet {
	this := Pallet{}
	return &this
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *Pallet) GetDimensions() Dimensions {
	if o == nil || IsNil(o.Dimensions) {
		var ret Dimensions
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pallet) GetDimensionsOk() (*Dimensions, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *Pallet) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given Dimensions and assigns it to the Dimensions field.
func (o *Pallet) SetDimensions(v Dimensions) {
	o.Dimensions = &v
}

// GetPackageId returns the PackageId field value
func (o *Pallet) GetPackageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageId
}

// GetPackageIdOk returns a tuple with the PackageId field value
// and a boolean to check if the value has been set.
func (o *Pallet) GetPackageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageId, true
}

// SetPackageId sets field value
func (o *Pallet) SetPackageId(v string) {
	o.PackageId = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *Pallet) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pallet) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *Pallet) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *Pallet) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetStackability returns the Stackability field value if set, zero value otherwise.
func (o *Pallet) GetStackability() Stackability {
	if o == nil || IsNil(o.Stackability) {
		var ret Stackability
		return ret
	}
	return *o.Stackability
}

// GetStackabilityOk returns a tuple with the Stackability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pallet) GetStackabilityOk() (*Stackability, bool) {
	if o == nil || IsNil(o.Stackability) {
		return nil, false
	}
	return o.Stackability, true
}

// HasStackability returns a boolean if a field has been set.
func (o *Pallet) HasStackability() bool {
	if o != nil && !IsNil(o.Stackability) {
		return true
	}

	return false
}

// SetStackability gets a reference to the given Stackability and assigns it to the Stackability field.
func (o *Pallet) SetStackability(v Stackability) {
	o.Stackability = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *Pallet) GetWeight() Weight {
	if o == nil || IsNil(o.Weight) {
		var ret Weight
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pallet) GetWeightOk() (*Weight, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *Pallet) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given Weight and assigns it to the Weight field.
func (o *Pallet) SetWeight(v Weight) {
	o.Weight = &v
}

func (o Pallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	toSerialize["packageId"] = o.PackageId
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Stackability) {
		toSerialize["stackability"] = o.Stackability
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullablePallet struct {
	value *Pallet
	isSet bool
}

func (v NullablePallet) Get() *Pallet {
	return v.value
}

func (v *NullablePallet) Set(val *Pallet) {
	v.value = val
	v.isSet = true
}

func (v NullablePallet) IsSet() bool {
	return v.isSet
}

func (v *NullablePallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePallet(val *Pallet) *NullablePallet {
	return &NullablePallet{value: val, isSet: true}
}

func (v NullablePallet) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
