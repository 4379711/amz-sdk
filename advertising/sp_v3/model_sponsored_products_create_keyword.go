package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateKeyword{}

// SponsoredProductsCreateKeyword struct for SponsoredProductsCreateKeyword
type SponsoredProductsCreateKeyword struct {
	// The unlocalized keyword text in the preferred locale of the advertiser.
	NativeLanguageKeyword *string `json:"nativeLanguageKeyword,omitempty"`
	// The locale preference of the advertiser. For example, if the advertiser’s preferred language is Simplified Chinese, set the locale to zh_CN. Supported locales include: Simplified Chinese (locale: zh_CN) for US, UK and CA. English (locale: en_GB) for DE, FR, IT and ES.
	NativeLanguageLocale *string `json:"nativeLanguageLocale,omitempty"`
	// The identifer of the campaign to which the keyword is associated.
	CampaignId string                                     `json:"campaignId"`
	MatchType  SponsoredProductsCreateOrUpdateMatchType   `json:"matchType"`
	State      SponsoredProductsCreateOrUpdateEntityState `json:"state"`
	// Bid associated with this keyword. Applicable to biddable match types only. For more information about bid constraints by marketplace, see [bid limits](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace).
	Bid NullableFloat64 `json:"bid,omitempty"`
	// The identifier of the ad group to which this keyword is associated.
	AdGroupId string `json:"adGroupId"`
	// The keyword text.
	KeywordText string `json:"keywordText"`
}

type _SponsoredProductsCreateKeyword SponsoredProductsCreateKeyword

// NewSponsoredProductsCreateKeyword instantiates a new SponsoredProductsCreateKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateKeyword(campaignId string, matchType SponsoredProductsCreateOrUpdateMatchType, state SponsoredProductsCreateOrUpdateEntityState, adGroupId string, keywordText string) *SponsoredProductsCreateKeyword {
	this := SponsoredProductsCreateKeyword{}
	this.CampaignId = campaignId
	this.MatchType = matchType
	this.State = state
	this.AdGroupId = adGroupId
	this.KeywordText = keywordText
	return &this
}

// NewSponsoredProductsCreateKeywordWithDefaults instantiates a new SponsoredProductsCreateKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateKeywordWithDefaults() *SponsoredProductsCreateKeyword {
	this := SponsoredProductsCreateKeyword{}
	return &this
}

// GetNativeLanguageKeyword returns the NativeLanguageKeyword field value if set, zero value otherwise.
func (o *SponsoredProductsCreateKeyword) GetNativeLanguageKeyword() string {
	if o == nil || IsNil(o.NativeLanguageKeyword) {
		var ret string
		return ret
	}
	return *o.NativeLanguageKeyword
}

// GetNativeLanguageKeywordOk returns a tuple with the NativeLanguageKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateKeyword) GetNativeLanguageKeywordOk() (*string, bool) {
	if o == nil || IsNil(o.NativeLanguageKeyword) {
		return nil, false
	}
	return o.NativeLanguageKeyword, true
}

// HasNativeLanguageKeyword returns a boolean if a field has been set.
func (o *SponsoredProductsCreateKeyword) HasNativeLanguageKeyword() bool {
	if o != nil && !IsNil(o.NativeLanguageKeyword) {
		return true
	}

	return false
}

// SetNativeLanguageKeyword gets a reference to the given string and assigns it to the NativeLanguageKeyword field.
func (o *SponsoredProductsCreateKeyword) SetNativeLanguageKeyword(v string) {
	o.NativeLanguageKeyword = &v
}

// GetNativeLanguageLocale returns the NativeLanguageLocale field value if set, zero value otherwise.
func (o *SponsoredProductsCreateKeyword) GetNativeLanguageLocale() string {
	if o == nil || IsNil(o.NativeLanguageLocale) {
		var ret string
		return ret
	}
	return *o.NativeLanguageLocale
}

// GetNativeLanguageLocaleOk returns a tuple with the NativeLanguageLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateKeyword) GetNativeLanguageLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.NativeLanguageLocale) {
		return nil, false
	}
	return o.NativeLanguageLocale, true
}

// HasNativeLanguageLocale returns a boolean if a field has been set.
func (o *SponsoredProductsCreateKeyword) HasNativeLanguageLocale() bool {
	if o != nil && !IsNil(o.NativeLanguageLocale) {
		return true
	}

	return false
}

// SetNativeLanguageLocale gets a reference to the given string and assigns it to the NativeLanguageLocale field.
func (o *SponsoredProductsCreateKeyword) SetNativeLanguageLocale(v string) {
	o.NativeLanguageLocale = &v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateKeyword) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateKeyword) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateKeyword) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetMatchType returns the MatchType field value
func (o *SponsoredProductsCreateKeyword) GetMatchType() SponsoredProductsCreateOrUpdateMatchType {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateMatchType
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateKeyword) GetMatchTypeOk() (*SponsoredProductsCreateOrUpdateMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *SponsoredProductsCreateKeyword) SetMatchType(v SponsoredProductsCreateOrUpdateMatchType) {
	o.MatchType = v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateKeyword) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateKeyword) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateKeyword) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = v
}

// GetBid returns the Bid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SponsoredProductsCreateKeyword) GetBid() float64 {
	if o == nil || IsNil(o.Bid.Get()) {
		var ret float64
		return ret
	}
	return *o.Bid.Get()
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SponsoredProductsCreateKeyword) GetBidOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bid.Get(), o.Bid.IsSet()
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsCreateKeyword) HasBid() bool {
	if o != nil && o.Bid.IsSet() {
		return true
	}

	return false
}

// SetBid gets a reference to the given NullableFloat64 and assigns it to the Bid field.
func (o *SponsoredProductsCreateKeyword) SetBid(v float64) {
	o.Bid.Set(&v)
}

// SetBidNil sets the value for Bid to be an explicit nil
func (o *SponsoredProductsCreateKeyword) SetBidNil() {
	o.Bid.Set(nil)
}

// UnsetBid ensures that no value is present for Bid, not even an explicit nil
func (o *SponsoredProductsCreateKeyword) UnsetBid() {
	o.Bid.Unset()
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsCreateKeyword) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateKeyword) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsCreateKeyword) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetKeywordText returns the KeywordText field value
func (o *SponsoredProductsCreateKeyword) GetKeywordText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateKeyword) GetKeywordTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordText, true
}

// SetKeywordText sets field value
func (o *SponsoredProductsCreateKeyword) SetKeywordText(v string) {
	o.KeywordText = v
}

func (o SponsoredProductsCreateKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NativeLanguageKeyword) {
		toSerialize["nativeLanguageKeyword"] = o.NativeLanguageKeyword
	}
	if !IsNil(o.NativeLanguageLocale) {
		toSerialize["nativeLanguageLocale"] = o.NativeLanguageLocale
	}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["matchType"] = o.MatchType
	toSerialize["state"] = o.State
	if o.Bid.IsSet() {
		toSerialize["bid"] = o.Bid.Get()
	}
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["keywordText"] = o.KeywordText
	return toSerialize, nil
}

type NullableSponsoredProductsCreateKeyword struct {
	value *SponsoredProductsCreateKeyword
	isSet bool
}

func (v NullableSponsoredProductsCreateKeyword) Get() *SponsoredProductsCreateKeyword {
	return v.value
}

func (v *NullableSponsoredProductsCreateKeyword) Set(val *SponsoredProductsCreateKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateKeyword(val *SponsoredProductsCreateKeyword) *NullableSponsoredProductsCreateKeyword {
	return &NullableSponsoredProductsCreateKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
