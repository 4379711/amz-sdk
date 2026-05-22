package definitions_product_types_20200901

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductTypeList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductTypeList{}

// ProductTypeList A list of Amazon product types with definitions available.
type ProductTypeList struct {
	ProductTypes []ProductType `json:"productTypes"`
	// Amazon product type version identifier.
	ProductTypeVersion string `json:"productTypeVersion"`
}

type _ProductTypeList ProductTypeList

// NewProductTypeList instantiates a new ProductTypeList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductTypeList(productTypes []ProductType, productTypeVersion string) *ProductTypeList {
	this := ProductTypeList{}
	this.ProductTypes = productTypes
	this.ProductTypeVersion = productTypeVersion
	return &this
}

// NewProductTypeListWithDefaults instantiates a new ProductTypeList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductTypeListWithDefaults() *ProductTypeList {
	this := ProductTypeList{}
	return &this
}

// GetProductTypes returns the ProductTypes field value
func (o *ProductTypeList) GetProductTypes() []ProductType {
	if o == nil {
		var ret []ProductType
		return ret
	}

	return o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value
// and a boolean to check if the value has been set.
func (o *ProductTypeList) GetProductTypesOk() ([]ProductType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductTypes, true
}

// SetProductTypes sets field value
func (o *ProductTypeList) SetProductTypes(v []ProductType) {
	o.ProductTypes = v
}

// GetProductTypeVersion returns the ProductTypeVersion field value
func (o *ProductTypeList) GetProductTypeVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTypeVersion
}

// GetProductTypeVersionOk returns a tuple with the ProductTypeVersion field value
// and a boolean to check if the value has been set.
func (o *ProductTypeList) GetProductTypeVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTypeVersion, true
}

// SetProductTypeVersion sets field value
func (o *ProductTypeList) SetProductTypeVersion(v string) {
	o.ProductTypeVersion = v
}

func (o ProductTypeList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productTypes"] = o.ProductTypes
	toSerialize["productTypeVersion"] = o.ProductTypeVersion
	return toSerialize, nil
}

type NullableProductTypeList struct {
	value *ProductTypeList
	isSet bool
}

func (v NullableProductTypeList) Get() *ProductTypeList {
	return v.value
}

func (v *NullableProductTypeList) Set(val *ProductTypeList) {
	v.value = val
	v.isSet = true
}

func (v NullableProductTypeList) IsSet() bool {
	return v.isSet
}

func (v *NullableProductTypeList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductTypeList(val *ProductTypeList) *NullableProductTypeList {
	return &NullableProductTypeList{value: val, isSet: true}
}

func (v NullableProductTypeList) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductTypeList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
