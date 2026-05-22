package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateNegativeKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateNegativeKeyword{}

// SponsoredProductsCreateNegativeKeyword struct for SponsoredProductsCreateNegativeKeyword
type SponsoredProductsCreateNegativeKeyword struct {
	// The unlocalized keyword text in the preferred locale of the advertiser
	NativeLanguageKeyword *string `json:"nativeLanguageKeyword,omitempty"`
	// The locale preference of the advertiser.
	NativeLanguageLocale *string `json:"nativeLanguageLocale,omitempty"`
	// The identifer of the campaign to which the keyword is associated.
	CampaignId string                                           `json:"campaignId"`
	MatchType  SponsoredProductsCreateOrUpdateNegativeMatchType `json:"matchType"`
	State      SponsoredProductsCreateOrUpdateEntityState       `json:"state"`
	// The identifier of the ad group to which this keyword is associated.
	AdGroupId string `json:"adGroupId"`
	// The keyword text.
	KeywordText string `json:"keywordText"`
}

type _SponsoredProductsCreateNegativeKeyword SponsoredProductsCreateNegativeKeyword

// NewSponsoredProductsCreateNegativeKeyword instantiates a new SponsoredProductsCreateNegativeKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateNegativeKeyword(campaignId string, matchType SponsoredProductsCreateOrUpdateNegativeMatchType, state SponsoredProductsCreateOrUpdateEntityState, adGroupId string, keywordText string) *SponsoredProductsCreateNegativeKeyword {
	this := SponsoredProductsCreateNegativeKeyword{}
	this.CampaignId = campaignId
	this.MatchType = matchType
	this.State = state
	this.AdGroupId = adGroupId
	this.KeywordText = keywordText
	return &this
}

// NewSponsoredProductsCreateNegativeKeywordWithDefaults instantiates a new SponsoredProductsCreateNegativeKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateNegativeKeywordWithDefaults() *SponsoredProductsCreateNegativeKeyword {
	this := SponsoredProductsCreateNegativeKeyword{}
	return &this
}

// GetNativeLanguageKeyword returns the NativeLanguageKeyword field value if set, zero value otherwise.
func (o *SponsoredProductsCreateNegativeKeyword) GetNativeLanguageKeyword() string {
	if o == nil || IsNil(o.NativeLanguageKeyword) {
		var ret string
		return ret
	}
	return *o.NativeLanguageKeyword
}

// GetNativeLanguageKeywordOk returns a tuple with the NativeLanguageKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateNegativeKeyword) GetNativeLanguageKeywordOk() (*string, bool) {
	if o == nil || IsNil(o.NativeLanguageKeyword) {
		return nil, false
	}
	return o.NativeLanguageKeyword, true
}

// HasNativeLanguageKeyword returns a boolean if a field has been set.
func (o *SponsoredProductsCreateNegativeKeyword) HasNativeLanguageKeyword() bool {
	if o != nil && !IsNil(o.NativeLanguageKeyword) {
		return true
	}

	return false
}

// SetNativeLanguageKeyword gets a reference to the given string and assigns it to the NativeLanguageKeyword field.
func (o *SponsoredProductsCreateNegativeKeyword) SetNativeLanguageKeyword(v string) {
	o.NativeLanguageKeyword = &v
}

// GetNativeLanguageLocale returns the NativeLanguageLocale field value if set, zero value otherwise.
func (o *SponsoredProductsCreateNegativeKeyword) GetNativeLanguageLocale() string {
	if o == nil || IsNil(o.NativeLanguageLocale) {
		var ret string
		return ret
	}
	return *o.NativeLanguageLocale
}

// GetNativeLanguageLocaleOk returns a tuple with the NativeLanguageLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateNegativeKeyword) GetNativeLanguageLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.NativeLanguageLocale) {
		return nil, false
	}
	return o.NativeLanguageLocale, true
}

// HasNativeLanguageLocale returns a boolean if a field has been set.
func (o *SponsoredProductsCreateNegativeKeyword) HasNativeLanguageLocale() bool {
	if o != nil && !IsNil(o.NativeLanguageLocale) {
		return true
	}

	return false
}

// SetNativeLanguageLocale gets a reference to the given string and assigns it to the NativeLanguageLocale field.
func (o *SponsoredProductsCreateNegativeKeyword) SetNativeLanguageLocale(v string) {
	o.NativeLanguageLocale = &v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateNegativeKeyword) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateNegativeKeyword) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateNegativeKeyword) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetMatchType returns the MatchType field value
func (o *SponsoredProductsCreateNegativeKeyword) GetMatchType() SponsoredProductsCreateOrUpdateNegativeMatchType {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateNegativeMatchType
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateNegativeKeyword) GetMatchTypeOk() (*SponsoredProductsCreateOrUpdateNegativeMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *SponsoredProductsCreateNegativeKeyword) SetMatchType(v SponsoredProductsCreateOrUpdateNegativeMatchType) {
	o.MatchType = v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateNegativeKeyword) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateNegativeKeyword) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateNegativeKeyword) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsCreateNegativeKeyword) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateNegativeKeyword) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsCreateNegativeKeyword) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetKeywordText returns the KeywordText field value
func (o *SponsoredProductsCreateNegativeKeyword) GetKeywordText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateNegativeKeyword) GetKeywordTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordText, true
}

// SetKeywordText sets field value
func (o *SponsoredProductsCreateNegativeKeyword) SetKeywordText(v string) {
	o.KeywordText = v
}

func (o SponsoredProductsCreateNegativeKeyword) ToMap() (map[string]interface{}, error) {
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
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["keywordText"] = o.KeywordText
	return toSerialize, nil
}

type NullableSponsoredProductsCreateNegativeKeyword struct {
	value *SponsoredProductsCreateNegativeKeyword
	isSet bool
}

func (v NullableSponsoredProductsCreateNegativeKeyword) Get() *SponsoredProductsCreateNegativeKeyword {
	return v.value
}

func (v *NullableSponsoredProductsCreateNegativeKeyword) Set(val *SponsoredProductsCreateNegativeKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateNegativeKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateNegativeKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateNegativeKeyword(val *SponsoredProductsCreateNegativeKeyword) *NullableSponsoredProductsCreateNegativeKeyword {
	return &NullableSponsoredProductsCreateNegativeKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateNegativeKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateNegativeKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
