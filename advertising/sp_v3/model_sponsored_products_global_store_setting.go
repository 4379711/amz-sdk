package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalStoreSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalStoreSetting{}

// SponsoredProductsGlobalStoreSetting struct for SponsoredProductsGlobalStoreSetting
type SponsoredProductsGlobalStoreSetting struct {
	// Country code of source marketplace where seller has listed the product. Possible source country codes include US, UK, DE, JP, and AE.
	CatalogSourceCountryCode *string `json:"catalogSourceCountryCode,omitempty"`
}

// NewSponsoredProductsGlobalStoreSetting instantiates a new SponsoredProductsGlobalStoreSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalStoreSetting() *SponsoredProductsGlobalStoreSetting {
	this := SponsoredProductsGlobalStoreSetting{}
	return &this
}

// NewSponsoredProductsGlobalStoreSettingWithDefaults instantiates a new SponsoredProductsGlobalStoreSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalStoreSettingWithDefaults() *SponsoredProductsGlobalStoreSetting {
	this := SponsoredProductsGlobalStoreSetting{}
	return &this
}

// GetCatalogSourceCountryCode returns the CatalogSourceCountryCode field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalStoreSetting) GetCatalogSourceCountryCode() string {
	if o == nil || IsNil(o.CatalogSourceCountryCode) {
		var ret string
		return ret
	}
	return *o.CatalogSourceCountryCode
}

// GetCatalogSourceCountryCodeOk returns a tuple with the CatalogSourceCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalStoreSetting) GetCatalogSourceCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CatalogSourceCountryCode) {
		return nil, false
	}
	return o.CatalogSourceCountryCode, true
}

// HasCatalogSourceCountryCode returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalStoreSetting) HasCatalogSourceCountryCode() bool {
	if o != nil && !IsNil(o.CatalogSourceCountryCode) {
		return true
	}

	return false
}

// SetCatalogSourceCountryCode gets a reference to the given string and assigns it to the CatalogSourceCountryCode field.
func (o *SponsoredProductsGlobalStoreSetting) SetCatalogSourceCountryCode(v string) {
	o.CatalogSourceCountryCode = &v
}

func (o SponsoredProductsGlobalStoreSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CatalogSourceCountryCode) {
		toSerialize["catalogSourceCountryCode"] = o.CatalogSourceCountryCode
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalStoreSetting struct {
	value *SponsoredProductsGlobalStoreSetting
	isSet bool
}

func (v NullableSponsoredProductsGlobalStoreSetting) Get() *SponsoredProductsGlobalStoreSetting {
	return v.value
}

func (v *NullableSponsoredProductsGlobalStoreSetting) Set(val *SponsoredProductsGlobalStoreSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalStoreSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalStoreSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalStoreSetting(val *SponsoredProductsGlobalStoreSetting) *NullableSponsoredProductsGlobalStoreSetting {
	return &NullableSponsoredProductsGlobalStoreSetting{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalStoreSetting) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalStoreSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
