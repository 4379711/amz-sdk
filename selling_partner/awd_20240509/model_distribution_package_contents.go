package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the DistributionPackageContents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DistributionPackageContents{}

// DistributionPackageContents Represents the contents inside a package, which can be products or a nested package.
type DistributionPackageContents struct {
	// This is required only when `DistributionPackageType=PALLET`.
	Packages []DistributionPackageQuantity `json:"packages,omitempty"`
	// This is required only when `DistributionPackageType=CASE`.
	Products []ProductQuantity `json:"products,omitempty"`
}

// NewDistributionPackageContents instantiates a new DistributionPackageContents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistributionPackageContents() *DistributionPackageContents {
	this := DistributionPackageContents{}
	return &this
}

// NewDistributionPackageContentsWithDefaults instantiates a new DistributionPackageContents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributionPackageContentsWithDefaults() *DistributionPackageContents {
	this := DistributionPackageContents{}
	return &this
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *DistributionPackageContents) GetPackages() []DistributionPackageQuantity {
	if o == nil || IsNil(o.Packages) {
		var ret []DistributionPackageQuantity
		return ret
	}
	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionPackageContents) GetPackagesOk() ([]DistributionPackageQuantity, bool) {
	if o == nil || IsNil(o.Packages) {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *DistributionPackageContents) HasPackages() bool {
	if o != nil && !IsNil(o.Packages) {
		return true
	}

	return false
}

// SetPackages gets a reference to the given []DistributionPackageQuantity and assigns it to the Packages field.
func (o *DistributionPackageContents) SetPackages(v []DistributionPackageQuantity) {
	o.Packages = v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *DistributionPackageContents) GetProducts() []ProductQuantity {
	if o == nil || IsNil(o.Products) {
		var ret []ProductQuantity
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionPackageContents) GetProductsOk() ([]ProductQuantity, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *DistributionPackageContents) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []ProductQuantity and assigns it to the Products field.
func (o *DistributionPackageContents) SetProducts(v []ProductQuantity) {
	o.Products = v
}

func (o DistributionPackageContents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Packages) {
		toSerialize["packages"] = o.Packages
	}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	return toSerialize, nil
}

type NullableDistributionPackageContents struct {
	value *DistributionPackageContents
	isSet bool
}

func (v NullableDistributionPackageContents) Get() *DistributionPackageContents {
	return v.value
}

func (v *NullableDistributionPackageContents) Set(val *DistributionPackageContents) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionPackageContents) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionPackageContents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionPackageContents(val *DistributionPackageContents) *NullableDistributionPackageContents {
	return &NullableDistributionPackageContents{value: val, isSet: true}
}

func (v NullableDistributionPackageContents) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDistributionPackageContents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
