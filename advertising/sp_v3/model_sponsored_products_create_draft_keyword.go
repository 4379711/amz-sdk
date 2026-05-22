package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateDraftKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateDraftKeyword{}

// SponsoredProductsCreateDraftKeyword struct for SponsoredProductsCreateDraftKeyword
type SponsoredProductsCreateDraftKeyword struct {
	// The unlocalized draft keyword text in the preferred locale of the advertiser.
	NativeLanguageKeyword *string `json:"nativeLanguageKeyword,omitempty"`
	// The locale preference of the advertiser. For example, if the advertiser’s preferred language is Simplified Chinese, set the locale to zh_CN. Supported locales include: Simplified Chinese (locale: zh_CN) for US, UK and CA. English (locale: en_GB) for DE, FR, IT and ES.
	NativeLanguageLocale *string `json:"nativeLanguageLocale,omitempty"`
	// The identifer of the campaign to which the draft keyword is associated.
	CampaignId string                                    `json:"campaignId"`
	MatchType  *SponsoredProductsCreateOrUpdateMatchType `json:"matchType,omitempty"`
	// Bid associated with this draft keyword. Applicable to biddable match types only. For more information about bid constraints by marketplace, see [bid limits](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace).
	Bid *float64 `json:"bid,omitempty"`
	// The identifier of the ad group to which this draft keyword is associated.
	AdGroupId string `json:"adGroupId"`
	// The draft keyword text.
	KeywordText string `json:"keywordText"`
}

type _SponsoredProductsCreateDraftKeyword SponsoredProductsCreateDraftKeyword

// NewSponsoredProductsCreateDraftKeyword instantiates a new SponsoredProductsCreateDraftKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateDraftKeyword(campaignId string, adGroupId string, keywordText string) *SponsoredProductsCreateDraftKeyword {
	this := SponsoredProductsCreateDraftKeyword{}
	this.CampaignId = campaignId
	this.AdGroupId = adGroupId
	this.KeywordText = keywordText
	return &this
}

// NewSponsoredProductsCreateDraftKeywordWithDefaults instantiates a new SponsoredProductsCreateDraftKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateDraftKeywordWithDefaults() *SponsoredProductsCreateDraftKeyword {
	this := SponsoredProductsCreateDraftKeyword{}
	return &this
}

// GetNativeLanguageKeyword returns the NativeLanguageKeyword field value if set, zero value otherwise.
func (o *SponsoredProductsCreateDraftKeyword) GetNativeLanguageKeyword() string {
	if o == nil || IsNil(o.NativeLanguageKeyword) {
		var ret string
		return ret
	}
	return *o.NativeLanguageKeyword
}

// GetNativeLanguageKeywordOk returns a tuple with the NativeLanguageKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftKeyword) GetNativeLanguageKeywordOk() (*string, bool) {
	if o == nil || IsNil(o.NativeLanguageKeyword) {
		return nil, false
	}
	return o.NativeLanguageKeyword, true
}

// HasNativeLanguageKeyword returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftKeyword) HasNativeLanguageKeyword() bool {
	if o != nil && !IsNil(o.NativeLanguageKeyword) {
		return true
	}

	return false
}

// SetNativeLanguageKeyword gets a reference to the given string and assigns it to the NativeLanguageKeyword field.
func (o *SponsoredProductsCreateDraftKeyword) SetNativeLanguageKeyword(v string) {
	o.NativeLanguageKeyword = &v
}

// GetNativeLanguageLocale returns the NativeLanguageLocale field value if set, zero value otherwise.
func (o *SponsoredProductsCreateDraftKeyword) GetNativeLanguageLocale() string {
	if o == nil || IsNil(o.NativeLanguageLocale) {
		var ret string
		return ret
	}
	return *o.NativeLanguageLocale
}

// GetNativeLanguageLocaleOk returns a tuple with the NativeLanguageLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftKeyword) GetNativeLanguageLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.NativeLanguageLocale) {
		return nil, false
	}
	return o.NativeLanguageLocale, true
}

// HasNativeLanguageLocale returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftKeyword) HasNativeLanguageLocale() bool {
	if o != nil && !IsNil(o.NativeLanguageLocale) {
		return true
	}

	return false
}

// SetNativeLanguageLocale gets a reference to the given string and assigns it to the NativeLanguageLocale field.
func (o *SponsoredProductsCreateDraftKeyword) SetNativeLanguageLocale(v string) {
	o.NativeLanguageLocale = &v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateDraftKeyword) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftKeyword) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateDraftKeyword) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *SponsoredProductsCreateDraftKeyword) GetMatchType() SponsoredProductsCreateOrUpdateMatchType {
	if o == nil || IsNil(o.MatchType) {
		var ret SponsoredProductsCreateOrUpdateMatchType
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftKeyword) GetMatchTypeOk() (*SponsoredProductsCreateOrUpdateMatchType, bool) {
	if o == nil || IsNil(o.MatchType) {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftKeyword) HasMatchType() bool {
	if o != nil && !IsNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given SponsoredProductsCreateOrUpdateMatchType and assigns it to the MatchType field.
func (o *SponsoredProductsCreateDraftKeyword) SetMatchType(v SponsoredProductsCreateOrUpdateMatchType) {
	o.MatchType = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *SponsoredProductsCreateDraftKeyword) GetBid() float64 {
	if o == nil || IsNil(o.Bid) {
		var ret float64
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftKeyword) GetBidOk() (*float64, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftKeyword) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given float64 and assigns it to the Bid field.
func (o *SponsoredProductsCreateDraftKeyword) SetBid(v float64) {
	o.Bid = &v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsCreateDraftKeyword) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftKeyword) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsCreateDraftKeyword) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetKeywordText returns the KeywordText field value
func (o *SponsoredProductsCreateDraftKeyword) GetKeywordText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftKeyword) GetKeywordTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordText, true
}

// SetKeywordText sets field value
func (o *SponsoredProductsCreateDraftKeyword) SetKeywordText(v string) {
	o.KeywordText = v
}

func (o SponsoredProductsCreateDraftKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NativeLanguageKeyword) {
		toSerialize["nativeLanguageKeyword"] = o.NativeLanguageKeyword
	}
	if !IsNil(o.NativeLanguageLocale) {
		toSerialize["nativeLanguageLocale"] = o.NativeLanguageLocale
	}
	toSerialize["campaignId"] = o.CampaignId
	if !IsNil(o.MatchType) {
		toSerialize["matchType"] = o.MatchType
	}
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["keywordText"] = o.KeywordText
	return toSerialize, nil
}

type NullableSponsoredProductsCreateDraftKeyword struct {
	value *SponsoredProductsCreateDraftKeyword
	isSet bool
}

func (v NullableSponsoredProductsCreateDraftKeyword) Get() *SponsoredProductsCreateDraftKeyword {
	return v.value
}

func (v *NullableSponsoredProductsCreateDraftKeyword) Set(val *SponsoredProductsCreateDraftKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateDraftKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateDraftKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateDraftKeyword(val *SponsoredProductsCreateDraftKeyword) *NullableSponsoredProductsCreateDraftKeyword {
	return &NullableSponsoredProductsCreateDraftKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateDraftKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateDraftKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
