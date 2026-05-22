package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductDetailsGlobalStoreSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductDetailsGlobalStoreSetting{}

// ProductDetailsGlobalStoreSetting This denotes the fields related to [GlobalStore Program](https://sellercentral.amazon.com/help/hub/reference/external/202139180).
type ProductDetailsGlobalStoreSetting struct {
	// Country code of source marketplace where seller has listed the product. Possible source country codes include US, UK, DE, JP, and AE.
	CatalogSourceCountryCode *string `json:"catalogSourceCountryCode,omitempty"`
}

// NewProductDetailsGlobalStoreSetting instantiates a new ProductDetailsGlobalStoreSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductDetailsGlobalStoreSetting() *ProductDetailsGlobalStoreSetting {
	this := ProductDetailsGlobalStoreSetting{}
	return &this
}

// NewProductDetailsGlobalStoreSettingWithDefaults instantiates a new ProductDetailsGlobalStoreSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductDetailsGlobalStoreSettingWithDefaults() *ProductDetailsGlobalStoreSetting {
	this := ProductDetailsGlobalStoreSetting{}
	return &this
}

// GetCatalogSourceCountryCode returns the CatalogSourceCountryCode field value if set, zero value otherwise.
func (o *ProductDetailsGlobalStoreSetting) GetCatalogSourceCountryCode() string {
	if o == nil || IsNil(o.CatalogSourceCountryCode) {
		var ret string
		return ret
	}
	return *o.CatalogSourceCountryCode
}

// GetCatalogSourceCountryCodeOk returns a tuple with the CatalogSourceCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailsGlobalStoreSetting) GetCatalogSourceCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CatalogSourceCountryCode) {
		return nil, false
	}
	return o.CatalogSourceCountryCode, true
}

// HasCatalogSourceCountryCode returns a boolean if a field has been set.
func (o *ProductDetailsGlobalStoreSetting) HasCatalogSourceCountryCode() bool {
	if o != nil && !IsNil(o.CatalogSourceCountryCode) {
		return true
	}

	return false
}

// SetCatalogSourceCountryCode gets a reference to the given string and assigns it to the CatalogSourceCountryCode field.
func (o *ProductDetailsGlobalStoreSetting) SetCatalogSourceCountryCode(v string) {
	o.CatalogSourceCountryCode = &v
}

func (o ProductDetailsGlobalStoreSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CatalogSourceCountryCode) {
		toSerialize["catalogSourceCountryCode"] = o.CatalogSourceCountryCode
	}
	return toSerialize, nil
}

type NullableProductDetailsGlobalStoreSetting struct {
	value *ProductDetailsGlobalStoreSetting
	isSet bool
}

func (v NullableProductDetailsGlobalStoreSetting) Get() *ProductDetailsGlobalStoreSetting {
	return v.value
}

func (v *NullableProductDetailsGlobalStoreSetting) Set(val *ProductDetailsGlobalStoreSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableProductDetailsGlobalStoreSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableProductDetailsGlobalStoreSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductDetailsGlobalStoreSetting(val *ProductDetailsGlobalStoreSetting) *NullableProductDetailsGlobalStoreSetting {
	return &NullableProductDetailsGlobalStoreSetting{value: val, isSet: true}
}

func (v NullableProductDetailsGlobalStoreSetting) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductDetailsGlobalStoreSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
