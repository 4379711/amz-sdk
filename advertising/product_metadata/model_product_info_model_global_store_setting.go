package product_metadata

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductInfoModelGlobalStoreSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductInfoModelGlobalStoreSetting{}

// ProductInfoModelGlobalStoreSetting This denotes the fields related to [GlobalStore Program](https://sellercentral.amazon.com/help/hub/reference/external/202139180).
type ProductInfoModelGlobalStoreSetting struct {
	// Country code of source marketplace where seller has listed the product. Possible source country codes include US, UK, DE, JP, and AE.
	CatalogSourceCountryCode *string `json:"catalogSourceCountryCode,omitempty"`
}

// NewProductInfoModelGlobalStoreSetting instantiates a new ProductInfoModelGlobalStoreSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductInfoModelGlobalStoreSetting() *ProductInfoModelGlobalStoreSetting {
	this := ProductInfoModelGlobalStoreSetting{}
	return &this
}

// NewProductInfoModelGlobalStoreSettingWithDefaults instantiates a new ProductInfoModelGlobalStoreSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductInfoModelGlobalStoreSettingWithDefaults() *ProductInfoModelGlobalStoreSetting {
	this := ProductInfoModelGlobalStoreSetting{}
	return &this
}

// GetCatalogSourceCountryCode returns the CatalogSourceCountryCode field value if set, zero value otherwise.
func (o *ProductInfoModelGlobalStoreSetting) GetCatalogSourceCountryCode() string {
	if o == nil || IsNil(o.CatalogSourceCountryCode) {
		var ret string
		return ret
	}
	return *o.CatalogSourceCountryCode
}

// GetCatalogSourceCountryCodeOk returns a tuple with the CatalogSourceCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModelGlobalStoreSetting) GetCatalogSourceCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CatalogSourceCountryCode) {
		return nil, false
	}
	return o.CatalogSourceCountryCode, true
}

// HasCatalogSourceCountryCode returns a boolean if a field has been set.
func (o *ProductInfoModelGlobalStoreSetting) HasCatalogSourceCountryCode() bool {
	if o != nil && !IsNil(o.CatalogSourceCountryCode) {
		return true
	}

	return false
}

// SetCatalogSourceCountryCode gets a reference to the given string and assigns it to the CatalogSourceCountryCode field.
func (o *ProductInfoModelGlobalStoreSetting) SetCatalogSourceCountryCode(v string) {
	o.CatalogSourceCountryCode = &v
}

func (o ProductInfoModelGlobalStoreSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CatalogSourceCountryCode) {
		toSerialize["catalogSourceCountryCode"] = o.CatalogSourceCountryCode
	}
	return toSerialize, nil
}

type NullableProductInfoModelGlobalStoreSetting struct {
	value *ProductInfoModelGlobalStoreSetting
	isSet bool
}

func (v NullableProductInfoModelGlobalStoreSetting) Get() *ProductInfoModelGlobalStoreSetting {
	return v.value
}

func (v *NullableProductInfoModelGlobalStoreSetting) Set(val *ProductInfoModelGlobalStoreSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableProductInfoModelGlobalStoreSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableProductInfoModelGlobalStoreSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductInfoModelGlobalStoreSetting(val *ProductInfoModelGlobalStoreSetting) *NullableProductInfoModelGlobalStoreSetting {
	return &NullableProductInfoModelGlobalStoreSetting{value: val, isSet: true}
}

func (v NullableProductInfoModelGlobalStoreSetting) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductInfoModelGlobalStoreSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
