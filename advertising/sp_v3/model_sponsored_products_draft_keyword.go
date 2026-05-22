package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftKeyword{}

// SponsoredProductsDraftKeyword struct for SponsoredProductsDraftKeyword
type SponsoredProductsDraftKeyword struct {
	// The identifier of the draft keyword.
	KeywordId string `json:"keywordId"`
	// The locale preference of the advertiser. For example, if the advertiser’s preferred language is Simplified Chinese, set the locale to zh_CN. Supported locales include: Simplified Chinese (locale: zh_CN) for US, UK and CA. English (locale: en_GB) for DE, FR, IT and ES.
	NativeLanguageLocale *string `json:"nativeLanguageLocale,omitempty"`
	// The identifier of the campaign to which the draft keyword is associated.
	CampaignId string                     `json:"campaignId"`
	MatchType  SponsoredProductsMatchType `json:"matchType"`
	// The unlocalized draft keyword text in the preferred locale of the advertiser.
	NativeLanguageDraftKeyword *string `json:"nativeLanguageDraftKeyword,omitempty"`
	// Bid associated with this draft keyword. Applicable to biddable match types only. For more information about bid constraints by marketplace, see [bid limits](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace).
	Bid NullableFloat64 `json:"bid,omitempty"`
	// The identifier of the ad group to which this draft keyword is associated.
	AdGroupId string `json:"adGroupId"`
	// The draft keyword text.
	KeywordText  string                                     `json:"keywordText"`
	ExtendedData *SponsoredProductsDraftKeywordExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsDraftKeyword SponsoredProductsDraftKeyword

// NewSponsoredProductsDraftKeyword instantiates a new SponsoredProductsDraftKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftKeyword(keywordId string, campaignId string, matchType SponsoredProductsMatchType, adGroupId string, keywordText string) *SponsoredProductsDraftKeyword {
	this := SponsoredProductsDraftKeyword{}
	this.KeywordId = keywordId
	this.CampaignId = campaignId
	this.MatchType = matchType
	this.AdGroupId = adGroupId
	this.KeywordText = keywordText
	return &this
}

// NewSponsoredProductsDraftKeywordWithDefaults instantiates a new SponsoredProductsDraftKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftKeywordWithDefaults() *SponsoredProductsDraftKeyword {
	this := SponsoredProductsDraftKeyword{}
	return &this
}

// GetKeywordId returns the KeywordId field value
func (o *SponsoredProductsDraftKeyword) GetKeywordId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordId
}

// GetKeywordIdOk returns a tuple with the KeywordId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeyword) GetKeywordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordId, true
}

// SetKeywordId sets field value
func (o *SponsoredProductsDraftKeyword) SetKeywordId(v string) {
	o.KeywordId = v
}

// GetNativeLanguageLocale returns the NativeLanguageLocale field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeyword) GetNativeLanguageLocale() string {
	if o == nil || IsNil(o.NativeLanguageLocale) {
		var ret string
		return ret
	}
	return *o.NativeLanguageLocale
}

// GetNativeLanguageLocaleOk returns a tuple with the NativeLanguageLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeyword) GetNativeLanguageLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.NativeLanguageLocale) {
		return nil, false
	}
	return o.NativeLanguageLocale, true
}

// HasNativeLanguageLocale returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeyword) HasNativeLanguageLocale() bool {
	if o != nil && !IsNil(o.NativeLanguageLocale) {
		return true
	}

	return false
}

// SetNativeLanguageLocale gets a reference to the given string and assigns it to the NativeLanguageLocale field.
func (o *SponsoredProductsDraftKeyword) SetNativeLanguageLocale(v string) {
	o.NativeLanguageLocale = &v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsDraftKeyword) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeyword) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsDraftKeyword) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetMatchType returns the MatchType field value
func (o *SponsoredProductsDraftKeyword) GetMatchType() SponsoredProductsMatchType {
	if o == nil {
		var ret SponsoredProductsMatchType
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeyword) GetMatchTypeOk() (*SponsoredProductsMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *SponsoredProductsDraftKeyword) SetMatchType(v SponsoredProductsMatchType) {
	o.MatchType = v
}

// GetNativeLanguageDraftKeyword returns the NativeLanguageDraftKeyword field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeyword) GetNativeLanguageDraftKeyword() string {
	if o == nil || IsNil(o.NativeLanguageDraftKeyword) {
		var ret string
		return ret
	}
	return *o.NativeLanguageDraftKeyword
}

// GetNativeLanguageDraftKeywordOk returns a tuple with the NativeLanguageDraftKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeyword) GetNativeLanguageDraftKeywordOk() (*string, bool) {
	if o == nil || IsNil(o.NativeLanguageDraftKeyword) {
		return nil, false
	}
	return o.NativeLanguageDraftKeyword, true
}

// HasNativeLanguageDraftKeyword returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeyword) HasNativeLanguageDraftKeyword() bool {
	if o != nil && !IsNil(o.NativeLanguageDraftKeyword) {
		return true
	}

	return false
}

// SetNativeLanguageDraftKeyword gets a reference to the given string and assigns it to the NativeLanguageDraftKeyword field.
func (o *SponsoredProductsDraftKeyword) SetNativeLanguageDraftKeyword(v string) {
	o.NativeLanguageDraftKeyword = &v
}

// GetBid returns the Bid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SponsoredProductsDraftKeyword) GetBid() float64 {
	if o == nil || IsNil(o.Bid.Get()) {
		var ret float64
		return ret
	}
	return *o.Bid.Get()
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SponsoredProductsDraftKeyword) GetBidOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bid.Get(), o.Bid.IsSet()
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeyword) HasBid() bool {
	if o != nil && o.Bid.IsSet() {
		return true
	}

	return false
}

// SetBid gets a reference to the given NullableFloat64 and assigns it to the Bid field.
func (o *SponsoredProductsDraftKeyword) SetBid(v float64) {
	o.Bid.Set(&v)
}

// SetBidNil sets the value for Bid to be an explicit nil
func (o *SponsoredProductsDraftKeyword) SetBidNil() {
	o.Bid.Set(nil)
}

// UnsetBid ensures that no value is present for Bid, not even an explicit nil
func (o *SponsoredProductsDraftKeyword) UnsetBid() {
	o.Bid.Unset()
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsDraftKeyword) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeyword) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsDraftKeyword) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetKeywordText returns the KeywordText field value
func (o *SponsoredProductsDraftKeyword) GetKeywordText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeyword) GetKeywordTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordText, true
}

// SetKeywordText sets field value
func (o *SponsoredProductsDraftKeyword) SetKeywordText(v string) {
	o.KeywordText = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeyword) GetExtendedData() SponsoredProductsDraftKeywordExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsDraftKeywordExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeyword) GetExtendedDataOk() (*SponsoredProductsDraftKeywordExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeyword) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsDraftKeywordExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsDraftKeyword) SetExtendedData(v SponsoredProductsDraftKeywordExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsDraftKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywordId"] = o.KeywordId
	if !IsNil(o.NativeLanguageLocale) {
		toSerialize["nativeLanguageLocale"] = o.NativeLanguageLocale
	}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["matchType"] = o.MatchType
	if !IsNil(o.NativeLanguageDraftKeyword) {
		toSerialize["nativeLanguageDraftKeyword"] = o.NativeLanguageDraftKeyword
	}
	if o.Bid.IsSet() {
		toSerialize["bid"] = o.Bid.Get()
	}
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["keywordText"] = o.KeywordText
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftKeyword struct {
	value *SponsoredProductsDraftKeyword
	isSet bool
}

func (v NullableSponsoredProductsDraftKeyword) Get() *SponsoredProductsDraftKeyword {
	return v.value
}

func (v *NullableSponsoredProductsDraftKeyword) Set(val *SponsoredProductsDraftKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftKeyword(val *SponsoredProductsDraftKeyword) *NullableSponsoredProductsDraftKeyword {
	return &NullableSponsoredProductsDraftKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
