package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsSPKeywordTargetDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsSPKeywordTargetDetails{}

// SponsoredProductsSPKeywordTargetDetails struct for SponsoredProductsSPKeywordTargetDetails
type SponsoredProductsSPKeywordTargetDetails struct {
	// The unlocalized keyword text in the preferred locale of the advertiser.
	NativeLanguageKeyword *string `json:"nativeLanguageKeyword,omitempty"`
	// The locale preference of the advertiser. For example, if the advertiser’s preferred language is Simplified Chinese, set the locale to zh_CN. Supported locales include: Simplified Chinese (locale: zh_CN) for US, UK and CA. English (locale: en_GB) for DE, FR, IT and ES.
	NativeLanguageLocale *string                           `json:"nativeLanguageLocale,omitempty"`
	MatchType            SponsoredProductsKeywordMatchType `json:"matchType"`
	// The keyword text to target.
	Keyword string `json:"keyword"`
}

type _SponsoredProductsSPKeywordTargetDetails SponsoredProductsSPKeywordTargetDetails

// NewSponsoredProductsSPKeywordTargetDetails instantiates a new SponsoredProductsSPKeywordTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsSPKeywordTargetDetails(matchType SponsoredProductsKeywordMatchType, keyword string) *SponsoredProductsSPKeywordTargetDetails {
	this := SponsoredProductsSPKeywordTargetDetails{}
	this.MatchType = matchType
	this.Keyword = keyword
	return &this
}

// NewSponsoredProductsSPKeywordTargetDetailsWithDefaults instantiates a new SponsoredProductsSPKeywordTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsSPKeywordTargetDetailsWithDefaults() *SponsoredProductsSPKeywordTargetDetails {
	this := SponsoredProductsSPKeywordTargetDetails{}
	return &this
}

// GetNativeLanguageKeyword returns the NativeLanguageKeyword field value if set, zero value otherwise.
func (o *SponsoredProductsSPKeywordTargetDetails) GetNativeLanguageKeyword() string {
	if o == nil || IsNil(o.NativeLanguageKeyword) {
		var ret string
		return ret
	}
	return *o.NativeLanguageKeyword
}

// GetNativeLanguageKeywordOk returns a tuple with the NativeLanguageKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPKeywordTargetDetails) GetNativeLanguageKeywordOk() (*string, bool) {
	if o == nil || IsNil(o.NativeLanguageKeyword) {
		return nil, false
	}
	return o.NativeLanguageKeyword, true
}

// HasNativeLanguageKeyword returns a boolean if a field has been set.
func (o *SponsoredProductsSPKeywordTargetDetails) HasNativeLanguageKeyword() bool {
	if o != nil && !IsNil(o.NativeLanguageKeyword) {
		return true
	}

	return false
}

// SetNativeLanguageKeyword gets a reference to the given string and assigns it to the NativeLanguageKeyword field.
func (o *SponsoredProductsSPKeywordTargetDetails) SetNativeLanguageKeyword(v string) {
	o.NativeLanguageKeyword = &v
}

// GetNativeLanguageLocale returns the NativeLanguageLocale field value if set, zero value otherwise.
func (o *SponsoredProductsSPKeywordTargetDetails) GetNativeLanguageLocale() string {
	if o == nil || IsNil(o.NativeLanguageLocale) {
		var ret string
		return ret
	}
	return *o.NativeLanguageLocale
}

// GetNativeLanguageLocaleOk returns a tuple with the NativeLanguageLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPKeywordTargetDetails) GetNativeLanguageLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.NativeLanguageLocale) {
		return nil, false
	}
	return o.NativeLanguageLocale, true
}

// HasNativeLanguageLocale returns a boolean if a field has been set.
func (o *SponsoredProductsSPKeywordTargetDetails) HasNativeLanguageLocale() bool {
	if o != nil && !IsNil(o.NativeLanguageLocale) {
		return true
	}

	return false
}

// SetNativeLanguageLocale gets a reference to the given string and assigns it to the NativeLanguageLocale field.
func (o *SponsoredProductsSPKeywordTargetDetails) SetNativeLanguageLocale(v string) {
	o.NativeLanguageLocale = &v
}

// GetMatchType returns the MatchType field value
func (o *SponsoredProductsSPKeywordTargetDetails) GetMatchType() SponsoredProductsKeywordMatchType {
	if o == nil {
		var ret SponsoredProductsKeywordMatchType
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPKeywordTargetDetails) GetMatchTypeOk() (*SponsoredProductsKeywordMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *SponsoredProductsSPKeywordTargetDetails) SetMatchType(v SponsoredProductsKeywordMatchType) {
	o.MatchType = v
}

// GetKeyword returns the Keyword field value
func (o *SponsoredProductsSPKeywordTargetDetails) GetKeyword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPKeywordTargetDetails) GetKeywordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keyword, true
}

// SetKeyword sets field value
func (o *SponsoredProductsSPKeywordTargetDetails) SetKeyword(v string) {
	o.Keyword = v
}

func (o SponsoredProductsSPKeywordTargetDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NativeLanguageKeyword) {
		toSerialize["nativeLanguageKeyword"] = o.NativeLanguageKeyword
	}
	if !IsNil(o.NativeLanguageLocale) {
		toSerialize["nativeLanguageLocale"] = o.NativeLanguageLocale
	}
	toSerialize["matchType"] = o.MatchType
	toSerialize["keyword"] = o.Keyword
	return toSerialize, nil
}

type NullableSponsoredProductsSPKeywordTargetDetails struct {
	value *SponsoredProductsSPKeywordTargetDetails
	isSet bool
}

func (v NullableSponsoredProductsSPKeywordTargetDetails) Get() *SponsoredProductsSPKeywordTargetDetails {
	return v.value
}

func (v *NullableSponsoredProductsSPKeywordTargetDetails) Set(val *SponsoredProductsSPKeywordTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsSPKeywordTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsSPKeywordTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsSPKeywordTargetDetails(val *SponsoredProductsSPKeywordTargetDetails) *NullableSponsoredProductsSPKeywordTargetDetails {
	return &NullableSponsoredProductsSPKeywordTargetDetails{value: val, isSet: true}
}

func (v NullableSponsoredProductsSPKeywordTargetDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsSPKeywordTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
