package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalProductEligibilityRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalProductEligibilityRequest{}

// GlobalProductEligibilityRequest A multi-country product advertising global eligibility request object.
type GlobalProductEligibilityRequest struct {
	// Set to 'sp' to check product eligibility for Sponsored Products advertisements. Set to 'sb' to check product eligibility for Sponsored Brands advertisements. Set to 'sd' to check product eligibility for Sponsored Display advertisements.
	AdType *string `json:"adType,omitempty"`
	// Set to the locale string in the table below to specify the language in which the response is returned. |Locale|Language (ISO 639)|Country (ISO 3166)| |------|------------------|------------------| |ar-AE|Arabic (ar)|United Arab Emirates (AE)| |zh-CN|Chinese (zh)|China (CN)| |nl-NL|Dutch (nl)|Netherlands (NL)| |en-AU|English (en)|Australia (AU)| |en-CA|English (en)|Canada (CA)| |en-IN|English (en)|India (IN)| |en-GB|English (en)|United Kingdom (GB)| |en-US|English (en)|United States (US)| |fr-CA|French (fr)|Canada (CA)| |fr-FR|French (fr)|France (FR)| |de-DE|German (de)|Germany (DE)| |it-IT|Italian (it)|Italy (IT)| |ja-JP|Japanese (ja)|Japan (JP)| |ko-KR|Korean (ko)|South Korea (KR)| |pt-BR|Portuguese (pt)|Brazil (BR)| |es-ES|Spanish (es)|Spain (ES)| |es-US|Spanish (es)|United States (US)| |es-MX|Spanish (es)|Mexico (MX)| |tr-TR|Turkish (tr)|Turkey (TR)|
	Locale *string `json:"locale,omitempty"`
	// A country code, in the ISO 3166-1 for example US for USA alpha-2 format, and the county's corresponding product identifiers list.
	CountryProductDetailsMap map[string][]ProductDetails `json:"countryProductDetailsMap"`
}

type _GlobalProductEligibilityRequest GlobalProductEligibilityRequest

// NewGlobalProductEligibilityRequest instantiates a new GlobalProductEligibilityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalProductEligibilityRequest(countryProductDetailsMap map[string][]ProductDetails) *GlobalProductEligibilityRequest {
	this := GlobalProductEligibilityRequest{}
	var adType string = "sp"
	this.AdType = &adType
	var locale string = "en-US"
	this.Locale = &locale
	this.CountryProductDetailsMap = countryProductDetailsMap
	return &this
}

// NewGlobalProductEligibilityRequestWithDefaults instantiates a new GlobalProductEligibilityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalProductEligibilityRequestWithDefaults() *GlobalProductEligibilityRequest {
	this := GlobalProductEligibilityRequest{}
	var adType string = "sp"
	this.AdType = &adType
	var locale string = "en-US"
	this.Locale = &locale
	return &this
}

// GetAdType returns the AdType field value if set, zero value otherwise.
func (o *GlobalProductEligibilityRequest) GetAdType() string {
	if o == nil || IsNil(o.AdType) {
		var ret string
		return ret
	}
	return *o.AdType
}

// GetAdTypeOk returns a tuple with the AdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalProductEligibilityRequest) GetAdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AdType) {
		return nil, false
	}
	return o.AdType, true
}

// HasAdType returns a boolean if a field has been set.
func (o *GlobalProductEligibilityRequest) HasAdType() bool {
	if o != nil && !IsNil(o.AdType) {
		return true
	}

	return false
}

// SetAdType gets a reference to the given string and assigns it to the AdType field.
func (o *GlobalProductEligibilityRequest) SetAdType(v string) {
	o.AdType = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *GlobalProductEligibilityRequest) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalProductEligibilityRequest) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *GlobalProductEligibilityRequest) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *GlobalProductEligibilityRequest) SetLocale(v string) {
	o.Locale = &v
}

// GetCountryProductDetailsMap returns the CountryProductDetailsMap field value
func (o *GlobalProductEligibilityRequest) GetCountryProductDetailsMap() map[string][]ProductDetails {
	if o == nil {
		var ret map[string][]ProductDetails
		return ret
	}

	return o.CountryProductDetailsMap
}

// GetCountryProductDetailsMapOk returns a tuple with the CountryProductDetailsMap field value
// and a boolean to check if the value has been set.
func (o *GlobalProductEligibilityRequest) GetCountryProductDetailsMapOk() (*map[string][]ProductDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryProductDetailsMap, true
}

// SetCountryProductDetailsMap sets field value
func (o *GlobalProductEligibilityRequest) SetCountryProductDetailsMap(v map[string][]ProductDetails) {
	o.CountryProductDetailsMap = v
}

func (o GlobalProductEligibilityRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdType) {
		toSerialize["adType"] = o.AdType
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	toSerialize["countryProductDetailsMap"] = o.CountryProductDetailsMap
	return toSerialize, nil
}

type NullableGlobalProductEligibilityRequest struct {
	value *GlobalProductEligibilityRequest
	isSet bool
}

func (v NullableGlobalProductEligibilityRequest) Get() *GlobalProductEligibilityRequest {
	return v.value
}

func (v *NullableGlobalProductEligibilityRequest) Set(val *GlobalProductEligibilityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalProductEligibilityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalProductEligibilityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalProductEligibilityRequest(val *GlobalProductEligibilityRequest) *NullableGlobalProductEligibilityRequest {
	return &NullableGlobalProductEligibilityRequest{value: val, isSet: true}
}

func (v NullableGlobalProductEligibilityRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalProductEligibilityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
