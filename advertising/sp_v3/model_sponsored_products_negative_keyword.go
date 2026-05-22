package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNegativeKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNegativeKeyword{}

// SponsoredProductsNegativeKeyword struct for SponsoredProductsNegativeKeyword
type SponsoredProductsNegativeKeyword struct {
	// The identifier of the keyword.
	KeywordId string `json:"keywordId"`
	// The unlocalized keyword text in the preferred locale of the advertiser
	NativeLanguageKeyword *string `json:"nativeLanguageKeyword,omitempty"`
	// The locale preference of the advertiser.
	NativeLanguageLocale *string `json:"nativeLanguageLocale,omitempty"`
	// The identifier of the campaign to which the keyword is associated.
	CampaignId string                             `json:"campaignId"`
	MatchType  SponsoredProductsNegativeMatchType `json:"matchType"`
	State      SponsoredProductsEntityState       `json:"state"`
	// The identifier of the ad group to which this keyword is associated.
	AdGroupId string `json:"adGroupId"`
	// The keyword text.
	KeywordText  string                                        `json:"keywordText"`
	ExtendedData *SponsoredProductsNegativeKeywordExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsNegativeKeyword SponsoredProductsNegativeKeyword

// NewSponsoredProductsNegativeKeyword instantiates a new SponsoredProductsNegativeKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNegativeKeyword(keywordId string, campaignId string, matchType SponsoredProductsNegativeMatchType, state SponsoredProductsEntityState, adGroupId string, keywordText string) *SponsoredProductsNegativeKeyword {
	this := SponsoredProductsNegativeKeyword{}
	this.KeywordId = keywordId
	this.CampaignId = campaignId
	this.MatchType = matchType
	this.State = state
	this.AdGroupId = adGroupId
	this.KeywordText = keywordText
	return &this
}

// NewSponsoredProductsNegativeKeywordWithDefaults instantiates a new SponsoredProductsNegativeKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNegativeKeywordWithDefaults() *SponsoredProductsNegativeKeyword {
	this := SponsoredProductsNegativeKeyword{}
	return &this
}

// GetKeywordId returns the KeywordId field value
func (o *SponsoredProductsNegativeKeyword) GetKeywordId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordId
}

// GetKeywordIdOk returns a tuple with the KeywordId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeyword) GetKeywordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordId, true
}

// SetKeywordId sets field value
func (o *SponsoredProductsNegativeKeyword) SetKeywordId(v string) {
	o.KeywordId = v
}

// GetNativeLanguageKeyword returns the NativeLanguageKeyword field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeyword) GetNativeLanguageKeyword() string {
	if o == nil || IsNil(o.NativeLanguageKeyword) {
		var ret string
		return ret
	}
	return *o.NativeLanguageKeyword
}

// GetNativeLanguageKeywordOk returns a tuple with the NativeLanguageKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeyword) GetNativeLanguageKeywordOk() (*string, bool) {
	if o == nil || IsNil(o.NativeLanguageKeyword) {
		return nil, false
	}
	return o.NativeLanguageKeyword, true
}

// HasNativeLanguageKeyword returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeyword) HasNativeLanguageKeyword() bool {
	if o != nil && !IsNil(o.NativeLanguageKeyword) {
		return true
	}

	return false
}

// SetNativeLanguageKeyword gets a reference to the given string and assigns it to the NativeLanguageKeyword field.
func (o *SponsoredProductsNegativeKeyword) SetNativeLanguageKeyword(v string) {
	o.NativeLanguageKeyword = &v
}

// GetNativeLanguageLocale returns the NativeLanguageLocale field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeyword) GetNativeLanguageLocale() string {
	if o == nil || IsNil(o.NativeLanguageLocale) {
		var ret string
		return ret
	}
	return *o.NativeLanguageLocale
}

// GetNativeLanguageLocaleOk returns a tuple with the NativeLanguageLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeyword) GetNativeLanguageLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.NativeLanguageLocale) {
		return nil, false
	}
	return o.NativeLanguageLocale, true
}

// HasNativeLanguageLocale returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeyword) HasNativeLanguageLocale() bool {
	if o != nil && !IsNil(o.NativeLanguageLocale) {
		return true
	}

	return false
}

// SetNativeLanguageLocale gets a reference to the given string and assigns it to the NativeLanguageLocale field.
func (o *SponsoredProductsNegativeKeyword) SetNativeLanguageLocale(v string) {
	o.NativeLanguageLocale = &v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsNegativeKeyword) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeyword) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsNegativeKeyword) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetMatchType returns the MatchType field value
func (o *SponsoredProductsNegativeKeyword) GetMatchType() SponsoredProductsNegativeMatchType {
	if o == nil {
		var ret SponsoredProductsNegativeMatchType
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeyword) GetMatchTypeOk() (*SponsoredProductsNegativeMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *SponsoredProductsNegativeKeyword) SetMatchType(v SponsoredProductsNegativeMatchType) {
	o.MatchType = v
}

// GetState returns the State field value
func (o *SponsoredProductsNegativeKeyword) GetState() SponsoredProductsEntityState {
	if o == nil {
		var ret SponsoredProductsEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeyword) GetStateOk() (*SponsoredProductsEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsNegativeKeyword) SetState(v SponsoredProductsEntityState) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsNegativeKeyword) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeyword) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsNegativeKeyword) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetKeywordText returns the KeywordText field value
func (o *SponsoredProductsNegativeKeyword) GetKeywordText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeyword) GetKeywordTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordText, true
}

// SetKeywordText sets field value
func (o *SponsoredProductsNegativeKeyword) SetKeywordText(v string) {
	o.KeywordText = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeyword) GetExtendedData() SponsoredProductsNegativeKeywordExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsNegativeKeywordExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeyword) GetExtendedDataOk() (*SponsoredProductsNegativeKeywordExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeyword) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsNegativeKeywordExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsNegativeKeyword) SetExtendedData(v SponsoredProductsNegativeKeywordExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsNegativeKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywordId"] = o.KeywordId
	if !IsNil(o.NativeLanguageKeyword) {
		toSerialize["nativeLanguageKeyword"] = o.NativeLanguageKeyword
	}
	if !IsNil(o.NativeLanguageLocale) {
		toSerialize["nativeLanguageLocale"] = o.NativeLanguageLocale
	}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["matchType"] = o.MatchType
	toSerialize["state"] = o.State
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["keywordText"] = o.KeywordText
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsNegativeKeyword struct {
	value *SponsoredProductsNegativeKeyword
	isSet bool
}

func (v NullableSponsoredProductsNegativeKeyword) Get() *SponsoredProductsNegativeKeyword {
	return v.value
}

func (v *NullableSponsoredProductsNegativeKeyword) Set(val *SponsoredProductsNegativeKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNegativeKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNegativeKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNegativeKeyword(val *SponsoredProductsNegativeKeyword) *NullableSponsoredProductsNegativeKeyword {
	return &NullableSponsoredProductsNegativeKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsNegativeKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNegativeKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
