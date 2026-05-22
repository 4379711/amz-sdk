package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignNegativeKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignNegativeKeyword{}

// SponsoredProductsCampaignNegativeKeyword struct for SponsoredProductsCampaignNegativeKeyword
type SponsoredProductsCampaignNegativeKeyword struct {
	// The identifier of the keyword.
	KeywordId string `json:"keywordId"`
	// The identifier of the campaign to which the keyword is associated.
	CampaignId string                             `json:"campaignId"`
	MatchType  SponsoredProductsNegativeMatchType `json:"matchType"`
	State      SponsoredProductsEntityState       `json:"state"`
	// The keyword text.
	KeywordText  string                                                `json:"keywordText"`
	ExtendedData *SponsoredProductsCampaignNegativeKeywordExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsCampaignNegativeKeyword SponsoredProductsCampaignNegativeKeyword

// NewSponsoredProductsCampaignNegativeKeyword instantiates a new SponsoredProductsCampaignNegativeKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignNegativeKeyword(keywordId string, campaignId string, matchType SponsoredProductsNegativeMatchType, state SponsoredProductsEntityState, keywordText string) *SponsoredProductsCampaignNegativeKeyword {
	this := SponsoredProductsCampaignNegativeKeyword{}
	this.KeywordId = keywordId
	this.CampaignId = campaignId
	this.MatchType = matchType
	this.State = state
	this.KeywordText = keywordText
	return &this
}

// NewSponsoredProductsCampaignNegativeKeywordWithDefaults instantiates a new SponsoredProductsCampaignNegativeKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignNegativeKeywordWithDefaults() *SponsoredProductsCampaignNegativeKeyword {
	this := SponsoredProductsCampaignNegativeKeyword{}
	return &this
}

// GetKeywordId returns the KeywordId field value
func (o *SponsoredProductsCampaignNegativeKeyword) GetKeywordId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordId
}

// GetKeywordIdOk returns a tuple with the KeywordId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeyword) GetKeywordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordId, true
}

// SetKeywordId sets field value
func (o *SponsoredProductsCampaignNegativeKeyword) SetKeywordId(v string) {
	o.KeywordId = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCampaignNegativeKeyword) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeyword) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCampaignNegativeKeyword) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetMatchType returns the MatchType field value
func (o *SponsoredProductsCampaignNegativeKeyword) GetMatchType() SponsoredProductsNegativeMatchType {
	if o == nil {
		var ret SponsoredProductsNegativeMatchType
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeyword) GetMatchTypeOk() (*SponsoredProductsNegativeMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *SponsoredProductsCampaignNegativeKeyword) SetMatchType(v SponsoredProductsNegativeMatchType) {
	o.MatchType = v
}

// GetState returns the State field value
func (o *SponsoredProductsCampaignNegativeKeyword) GetState() SponsoredProductsEntityState {
	if o == nil {
		var ret SponsoredProductsEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeyword) GetStateOk() (*SponsoredProductsEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCampaignNegativeKeyword) SetState(v SponsoredProductsEntityState) {
	o.State = v
}

// GetKeywordText returns the KeywordText field value
func (o *SponsoredProductsCampaignNegativeKeyword) GetKeywordText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeyword) GetKeywordTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordText, true
}

// SetKeywordText sets field value
func (o *SponsoredProductsCampaignNegativeKeyword) SetKeywordText(v string) {
	o.KeywordText = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignNegativeKeyword) GetExtendedData() SponsoredProductsCampaignNegativeKeywordExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsCampaignNegativeKeywordExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeyword) GetExtendedDataOk() (*SponsoredProductsCampaignNegativeKeywordExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignNegativeKeyword) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsCampaignNegativeKeywordExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsCampaignNegativeKeyword) SetExtendedData(v SponsoredProductsCampaignNegativeKeywordExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsCampaignNegativeKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywordId"] = o.KeywordId
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["matchType"] = o.MatchType
	toSerialize["state"] = o.State
	toSerialize["keywordText"] = o.KeywordText
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignNegativeKeyword struct {
	value *SponsoredProductsCampaignNegativeKeyword
	isSet bool
}

func (v NullableSponsoredProductsCampaignNegativeKeyword) Get() *SponsoredProductsCampaignNegativeKeyword {
	return v.value
}

func (v *NullableSponsoredProductsCampaignNegativeKeyword) Set(val *SponsoredProductsCampaignNegativeKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignNegativeKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignNegativeKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignNegativeKeyword(val *SponsoredProductsCampaignNegativeKeyword) *NullableSponsoredProductsCampaignNegativeKeyword {
	return &NullableSponsoredProductsCampaignNegativeKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignNegativeKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignNegativeKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
