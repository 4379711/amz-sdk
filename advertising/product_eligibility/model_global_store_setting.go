package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalStoreSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalStoreSetting{}

// GlobalStoreSetting Fields required to check eligibility for [GlobalStore Program](https://sellercentral.amazon.com/help/hub/reference/external/202139180) Ads.
type GlobalStoreSetting struct {
	// Country code of source marketplace where seller has listed the product. Possible source country codes include US, UK, DE, JP, and AE.
	CatalogSourceCountryCode *string `json:"catalogSourceCountryCode,omitempty"`
}

// NewGlobalStoreSetting instantiates a new GlobalStoreSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalStoreSetting() *GlobalStoreSetting {
	this := GlobalStoreSetting{}
	return &this
}

// NewGlobalStoreSettingWithDefaults instantiates a new GlobalStoreSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalStoreSettingWithDefaults() *GlobalStoreSetting {
	this := GlobalStoreSetting{}
	return &this
}

// GetCatalogSourceCountryCode returns the CatalogSourceCountryCode field value if set, zero value otherwise.
func (o *GlobalStoreSetting) GetCatalogSourceCountryCode() string {
	if o == nil || IsNil(o.CatalogSourceCountryCode) {
		var ret string
		return ret
	}
	return *o.CatalogSourceCountryCode
}

// GetCatalogSourceCountryCodeOk returns a tuple with the CatalogSourceCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalStoreSetting) GetCatalogSourceCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CatalogSourceCountryCode) {
		return nil, false
	}
	return o.CatalogSourceCountryCode, true
}

// HasCatalogSourceCountryCode returns a boolean if a field has been set.
func (o *GlobalStoreSetting) HasCatalogSourceCountryCode() bool {
	if o != nil && !IsNil(o.CatalogSourceCountryCode) {
		return true
	}

	return false
}

// SetCatalogSourceCountryCode gets a reference to the given string and assigns it to the CatalogSourceCountryCode field.
func (o *GlobalStoreSetting) SetCatalogSourceCountryCode(v string) {
	o.CatalogSourceCountryCode = &v
}

func (o GlobalStoreSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CatalogSourceCountryCode) {
		toSerialize["catalogSourceCountryCode"] = o.CatalogSourceCountryCode
	}
	return toSerialize, nil
}

type NullableGlobalStoreSetting struct {
	value *GlobalStoreSetting
	isSet bool
}

func (v NullableGlobalStoreSetting) Get() *GlobalStoreSetting {
	return v.value
}

func (v *NullableGlobalStoreSetting) Set(val *GlobalStoreSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalStoreSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalStoreSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalStoreSetting(val *GlobalStoreSetting) *NullableGlobalStoreSetting {
	return &NullableGlobalStoreSetting{value: val, isSet: true}
}

func (v NullableGlobalStoreSetting) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalStoreSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
