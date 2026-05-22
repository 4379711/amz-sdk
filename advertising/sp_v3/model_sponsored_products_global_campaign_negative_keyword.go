package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalCampaignNegativeKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalCampaignNegativeKeyword{}

// SponsoredProductsGlobalCampaignNegativeKeyword struct for SponsoredProductsGlobalCampaignNegativeKeyword
type SponsoredProductsGlobalCampaignNegativeKeyword struct {
	// The identifier of the keyword.
	KeywordId string `json:"keywordId"`
	// The identifier of the campaign to which the keyword is associated.
	CampaignId string                             `json:"campaignId"`
	MatchType  SponsoredProductsNegativeMatchType `json:"matchType"`
	// Name for the keyword
	Name         *string                                                     `json:"name,omitempty"`
	State        SponsoredProductsGlobalEntityState                          `json:"state"`
	KeywordText  SponsoredProductsGlobalNegativeKeywordText                  `json:"keywordText"`
	ExtendedData *SponsoredProductsGlobalCampaignNegativeKeywordExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsGlobalCampaignNegativeKeyword SponsoredProductsGlobalCampaignNegativeKeyword

// NewSponsoredProductsGlobalCampaignNegativeKeyword instantiates a new SponsoredProductsGlobalCampaignNegativeKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalCampaignNegativeKeyword(keywordId string, campaignId string, matchType SponsoredProductsNegativeMatchType, state SponsoredProductsGlobalEntityState, keywordText SponsoredProductsGlobalNegativeKeywordText) *SponsoredProductsGlobalCampaignNegativeKeyword {
	this := SponsoredProductsGlobalCampaignNegativeKeyword{}
	this.KeywordId = keywordId
	this.CampaignId = campaignId
	this.MatchType = matchType
	this.State = state
	this.KeywordText = keywordText
	return &this
}

// NewSponsoredProductsGlobalCampaignNegativeKeywordWithDefaults instantiates a new SponsoredProductsGlobalCampaignNegativeKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalCampaignNegativeKeywordWithDefaults() *SponsoredProductsGlobalCampaignNegativeKeyword {
	this := SponsoredProductsGlobalCampaignNegativeKeyword{}
	return &this
}

// GetKeywordId returns the KeywordId field value
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) GetKeywordId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordId
}

// GetKeywordIdOk returns a tuple with the KeywordId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) GetKeywordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordId, true
}

// SetKeywordId sets field value
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) SetKeywordId(v string) {
	o.KeywordId = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetMatchType returns the MatchType field value
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) GetMatchType() SponsoredProductsNegativeMatchType {
	if o == nil {
		var ret SponsoredProductsNegativeMatchType
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) GetMatchTypeOk() (*SponsoredProductsNegativeMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) SetMatchType(v SponsoredProductsNegativeMatchType) {
	o.MatchType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) GetState() SponsoredProductsGlobalEntityState {
	if o == nil {
		var ret SponsoredProductsGlobalEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) GetStateOk() (*SponsoredProductsGlobalEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) SetState(v SponsoredProductsGlobalEntityState) {
	o.State = v
}

// GetKeywordText returns the KeywordText field value
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) GetKeywordText() SponsoredProductsGlobalNegativeKeywordText {
	if o == nil {
		var ret SponsoredProductsGlobalNegativeKeywordText
		return ret
	}

	return o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) GetKeywordTextOk() (*SponsoredProductsGlobalNegativeKeywordText, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordText, true
}

// SetKeywordText sets field value
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) SetKeywordText(v SponsoredProductsGlobalNegativeKeywordText) {
	o.KeywordText = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) GetExtendedData() SponsoredProductsGlobalCampaignNegativeKeywordExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsGlobalCampaignNegativeKeywordExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) GetExtendedDataOk() (*SponsoredProductsGlobalCampaignNegativeKeywordExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsGlobalCampaignNegativeKeywordExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsGlobalCampaignNegativeKeyword) SetExtendedData(v SponsoredProductsGlobalCampaignNegativeKeywordExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsGlobalCampaignNegativeKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywordId"] = o.KeywordId
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["matchType"] = o.MatchType
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["state"] = o.State
	toSerialize["keywordText"] = o.KeywordText
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalCampaignNegativeKeyword struct {
	value *SponsoredProductsGlobalCampaignNegativeKeyword
	isSet bool
}

func (v NullableSponsoredProductsGlobalCampaignNegativeKeyword) Get() *SponsoredProductsGlobalCampaignNegativeKeyword {
	return v.value
}

func (v *NullableSponsoredProductsGlobalCampaignNegativeKeyword) Set(val *SponsoredProductsGlobalCampaignNegativeKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalCampaignNegativeKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalCampaignNegativeKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalCampaignNegativeKeyword(val *SponsoredProductsGlobalCampaignNegativeKeyword) *NullableSponsoredProductsGlobalCampaignNegativeKeyword {
	return &NullableSponsoredProductsGlobalCampaignNegativeKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalCampaignNegativeKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalCampaignNegativeKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
