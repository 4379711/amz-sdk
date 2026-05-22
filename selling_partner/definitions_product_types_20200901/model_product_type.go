package definitions_product_types_20200901

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductType{}

// ProductType An Amazon product type with a definition available.
type ProductType struct {
	// The name of the Amazon product type.
	Name string `json:"name"`
	// The human-readable and localized description of the Amazon product type.
	DisplayName string `json:"displayName"`
	// The Amazon marketplace identifiers for which the product type definition is available.
	MarketplaceIds []string `json:"marketplaceIds"`
}

type _ProductType ProductType

// NewProductType instantiates a new ProductType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductType(name string, displayName string, marketplaceIds []string) *ProductType {
	this := ProductType{}
	this.Name = name
	this.DisplayName = displayName
	this.MarketplaceIds = marketplaceIds
	return &this
}

// NewProductTypeWithDefaults instantiates a new ProductType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductTypeWithDefaults() *ProductType {
	this := ProductType{}
	return &this
}

// GetName returns the Name field value
func (o *ProductType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProductType) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *ProductType) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ProductType) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ProductType) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetMarketplaceIds returns the MarketplaceIds field value
func (o *ProductType) GetMarketplaceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MarketplaceIds
}

// GetMarketplaceIdsOk returns a tuple with the MarketplaceIds field value
// and a boolean to check if the value has been set.
func (o *ProductType) GetMarketplaceIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MarketplaceIds, true
}

// SetMarketplaceIds sets field value
func (o *ProductType) SetMarketplaceIds(v []string) {
	o.MarketplaceIds = v
}

func (o ProductType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["displayName"] = o.DisplayName
	toSerialize["marketplaceIds"] = o.MarketplaceIds
	return toSerialize, nil
}

type NullableProductType struct {
	value *ProductType
	isSet bool
}

func (v NullableProductType) Get() *ProductType {
	return v.value
}

func (v *NullableProductType) Set(val *ProductType) {
	v.value = val
	v.isSet = true
}

func (v NullableProductType) IsSet() bool {
	return v.isSet
}

func (v *NullableProductType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductType(val *ProductType) *NullableProductType {
	return &NullableProductType{value: val, isSet: true}
}

func (v NullableProductType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
