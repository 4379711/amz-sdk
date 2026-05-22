package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignNegativeKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignNegativeKeyword{}

// SponsoredProductsDraftCampaignNegativeKeyword struct for SponsoredProductsDraftCampaignNegativeKeyword
type SponsoredProductsDraftCampaignNegativeKeyword struct {
	// The identifier of the keyword.
	KeywordId string `json:"keywordId"`
	// The identifier of the draft to which the keyword is associated.
	CampaignId string                             `json:"campaignId"`
	MatchType  SponsoredProductsNegativeMatchType `json:"matchType"`
	// The keyword text.
	KeywordText  string                                                     `json:"keywordText"`
	ExtendedData *SponsoredProductsDraftCampaignNegativeKeywordExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsDraftCampaignNegativeKeyword SponsoredProductsDraftCampaignNegativeKeyword

// NewSponsoredProductsDraftCampaignNegativeKeyword instantiates a new SponsoredProductsDraftCampaignNegativeKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignNegativeKeyword(keywordId string, campaignId string, matchType SponsoredProductsNegativeMatchType, keywordText string) *SponsoredProductsDraftCampaignNegativeKeyword {
	this := SponsoredProductsDraftCampaignNegativeKeyword{}
	this.KeywordId = keywordId
	this.CampaignId = campaignId
	this.MatchType = matchType
	this.KeywordText = keywordText
	return &this
}

// NewSponsoredProductsDraftCampaignNegativeKeywordWithDefaults instantiates a new SponsoredProductsDraftCampaignNegativeKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignNegativeKeywordWithDefaults() *SponsoredProductsDraftCampaignNegativeKeyword {
	this := SponsoredProductsDraftCampaignNegativeKeyword{}
	return &this
}

// GetKeywordId returns the KeywordId field value
func (o *SponsoredProductsDraftCampaignNegativeKeyword) GetKeywordId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordId
}

// GetKeywordIdOk returns a tuple with the KeywordId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeyword) GetKeywordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordId, true
}

// SetKeywordId sets field value
func (o *SponsoredProductsDraftCampaignNegativeKeyword) SetKeywordId(v string) {
	o.KeywordId = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsDraftCampaignNegativeKeyword) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeyword) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsDraftCampaignNegativeKeyword) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetMatchType returns the MatchType field value
func (o *SponsoredProductsDraftCampaignNegativeKeyword) GetMatchType() SponsoredProductsNegativeMatchType {
	if o == nil {
		var ret SponsoredProductsNegativeMatchType
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeyword) GetMatchTypeOk() (*SponsoredProductsNegativeMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *SponsoredProductsDraftCampaignNegativeKeyword) SetMatchType(v SponsoredProductsNegativeMatchType) {
	o.MatchType = v
}

// GetKeywordText returns the KeywordText field value
func (o *SponsoredProductsDraftCampaignNegativeKeyword) GetKeywordText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeyword) GetKeywordTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordText, true
}

// SetKeywordText sets field value
func (o *SponsoredProductsDraftCampaignNegativeKeyword) SetKeywordText(v string) {
	o.KeywordText = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignNegativeKeyword) GetExtendedData() SponsoredProductsDraftCampaignNegativeKeywordExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsDraftCampaignNegativeKeywordExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeyword) GetExtendedDataOk() (*SponsoredProductsDraftCampaignNegativeKeywordExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeyword) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsDraftCampaignNegativeKeywordExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsDraftCampaignNegativeKeyword) SetExtendedData(v SponsoredProductsDraftCampaignNegativeKeywordExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsDraftCampaignNegativeKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywordId"] = o.KeywordId
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["matchType"] = o.MatchType
	toSerialize["keywordText"] = o.KeywordText
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftCampaignNegativeKeyword struct {
	value *SponsoredProductsDraftCampaignNegativeKeyword
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeyword) Get() *SponsoredProductsDraftCampaignNegativeKeyword {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeyword) Set(val *SponsoredProductsDraftCampaignNegativeKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignNegativeKeyword(val *SponsoredProductsDraftCampaignNegativeKeyword) *NullableSponsoredProductsDraftCampaignNegativeKeyword {
	return &NullableSponsoredProductsDraftCampaignNegativeKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
