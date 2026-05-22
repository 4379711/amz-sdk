package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateDraftNegativeKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateDraftNegativeKeyword{}

// SponsoredProductsCreateDraftNegativeKeyword struct for SponsoredProductsCreateDraftNegativeKeyword
type SponsoredProductsCreateDraftNegativeKeyword struct {
	// The unlocalized keyword text in the preferred locale of the advertiser
	NativeLanguageKeyword *string `json:"nativeLanguageKeyword,omitempty"`
	// The locale preference of the advertiser.
	NativeLanguageLocale *string `json:"nativeLanguageLocale,omitempty"`
	// The identifer of the campaign to which the keyword is associated.
	CampaignId string                                           `json:"campaignId"`
	MatchType  SponsoredProductsCreateOrUpdateNegativeMatchType `json:"matchType"`
	// The identifier of the ad group to which this keyword is associated.
	AdGroupId string `json:"adGroupId"`
	// The keyword text.
	KeywordText string `json:"keywordText"`
}

type _SponsoredProductsCreateDraftNegativeKeyword SponsoredProductsCreateDraftNegativeKeyword

// NewSponsoredProductsCreateDraftNegativeKeyword instantiates a new SponsoredProductsCreateDraftNegativeKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateDraftNegativeKeyword(campaignId string, matchType SponsoredProductsCreateOrUpdateNegativeMatchType, adGroupId string, keywordText string) *SponsoredProductsCreateDraftNegativeKeyword {
	this := SponsoredProductsCreateDraftNegativeKeyword{}
	this.CampaignId = campaignId
	this.MatchType = matchType
	this.AdGroupId = adGroupId
	this.KeywordText = keywordText
	return &this
}

// NewSponsoredProductsCreateDraftNegativeKeywordWithDefaults instantiates a new SponsoredProductsCreateDraftNegativeKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateDraftNegativeKeywordWithDefaults() *SponsoredProductsCreateDraftNegativeKeyword {
	this := SponsoredProductsCreateDraftNegativeKeyword{}
	return &this
}

// GetNativeLanguageKeyword returns the NativeLanguageKeyword field value if set, zero value otherwise.
func (o *SponsoredProductsCreateDraftNegativeKeyword) GetNativeLanguageKeyword() string {
	if o == nil || IsNil(o.NativeLanguageKeyword) {
		var ret string
		return ret
	}
	return *o.NativeLanguageKeyword
}

// GetNativeLanguageKeywordOk returns a tuple with the NativeLanguageKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftNegativeKeyword) GetNativeLanguageKeywordOk() (*string, bool) {
	if o == nil || IsNil(o.NativeLanguageKeyword) {
		return nil, false
	}
	return o.NativeLanguageKeyword, true
}

// HasNativeLanguageKeyword returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftNegativeKeyword) HasNativeLanguageKeyword() bool {
	if o != nil && !IsNil(o.NativeLanguageKeyword) {
		return true
	}

	return false
}

// SetNativeLanguageKeyword gets a reference to the given string and assigns it to the NativeLanguageKeyword field.
func (o *SponsoredProductsCreateDraftNegativeKeyword) SetNativeLanguageKeyword(v string) {
	o.NativeLanguageKeyword = &v
}

// GetNativeLanguageLocale returns the NativeLanguageLocale field value if set, zero value otherwise.
func (o *SponsoredProductsCreateDraftNegativeKeyword) GetNativeLanguageLocale() string {
	if o == nil || IsNil(o.NativeLanguageLocale) {
		var ret string
		return ret
	}
	return *o.NativeLanguageLocale
}

// GetNativeLanguageLocaleOk returns a tuple with the NativeLanguageLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftNegativeKeyword) GetNativeLanguageLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.NativeLanguageLocale) {
		return nil, false
	}
	return o.NativeLanguageLocale, true
}

// HasNativeLanguageLocale returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftNegativeKeyword) HasNativeLanguageLocale() bool {
	if o != nil && !IsNil(o.NativeLanguageLocale) {
		return true
	}

	return false
}

// SetNativeLanguageLocale gets a reference to the given string and assigns it to the NativeLanguageLocale field.
func (o *SponsoredProductsCreateDraftNegativeKeyword) SetNativeLanguageLocale(v string) {
	o.NativeLanguageLocale = &v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateDraftNegativeKeyword) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftNegativeKeyword) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateDraftNegativeKeyword) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetMatchType returns the MatchType field value
func (o *SponsoredProductsCreateDraftNegativeKeyword) GetMatchType() SponsoredProductsCreateOrUpdateNegativeMatchType {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateNegativeMatchType
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftNegativeKeyword) GetMatchTypeOk() (*SponsoredProductsCreateOrUpdateNegativeMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *SponsoredProductsCreateDraftNegativeKeyword) SetMatchType(v SponsoredProductsCreateOrUpdateNegativeMatchType) {
	o.MatchType = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsCreateDraftNegativeKeyword) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftNegativeKeyword) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsCreateDraftNegativeKeyword) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetKeywordText returns the KeywordText field value
func (o *SponsoredProductsCreateDraftNegativeKeyword) GetKeywordText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftNegativeKeyword) GetKeywordTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordText, true
}

// SetKeywordText sets field value
func (o *SponsoredProductsCreateDraftNegativeKeyword) SetKeywordText(v string) {
	o.KeywordText = v
}

func (o SponsoredProductsCreateDraftNegativeKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NativeLanguageKeyword) {
		toSerialize["nativeLanguageKeyword"] = o.NativeLanguageKeyword
	}
	if !IsNil(o.NativeLanguageLocale) {
		toSerialize["nativeLanguageLocale"] = o.NativeLanguageLocale
	}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["matchType"] = o.MatchType
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["keywordText"] = o.KeywordText
	return toSerialize, nil
}

type NullableSponsoredProductsCreateDraftNegativeKeyword struct {
	value *SponsoredProductsCreateDraftNegativeKeyword
	isSet bool
}

func (v NullableSponsoredProductsCreateDraftNegativeKeyword) Get() *SponsoredProductsCreateDraftNegativeKeyword {
	return v.value
}

func (v *NullableSponsoredProductsCreateDraftNegativeKeyword) Set(val *SponsoredProductsCreateDraftNegativeKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateDraftNegativeKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateDraftNegativeKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateDraftNegativeKeyword(val *SponsoredProductsCreateDraftNegativeKeyword) *NullableSponsoredProductsCreateDraftNegativeKeyword {
	return &NullableSponsoredProductsCreateDraftNegativeKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateDraftNegativeKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateDraftNegativeKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
