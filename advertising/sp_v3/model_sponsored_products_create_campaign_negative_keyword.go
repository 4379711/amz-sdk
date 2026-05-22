package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateCampaignNegativeKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateCampaignNegativeKeyword{}

// SponsoredProductsCreateCampaignNegativeKeyword struct for SponsoredProductsCreateCampaignNegativeKeyword
type SponsoredProductsCreateCampaignNegativeKeyword struct {
	// The identifier of the campaign to which the keyword is associated.
	CampaignId string                                           `json:"campaignId"`
	MatchType  SponsoredProductsCreateOrUpdateNegativeMatchType `json:"matchType"`
	State      SponsoredProductsCreateOrUpdateEntityState       `json:"state"`
	// The keyword text.
	KeywordText string `json:"keywordText"`
}

type _SponsoredProductsCreateCampaignNegativeKeyword SponsoredProductsCreateCampaignNegativeKeyword

// NewSponsoredProductsCreateCampaignNegativeKeyword instantiates a new SponsoredProductsCreateCampaignNegativeKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateCampaignNegativeKeyword(campaignId string, matchType SponsoredProductsCreateOrUpdateNegativeMatchType, state SponsoredProductsCreateOrUpdateEntityState, keywordText string) *SponsoredProductsCreateCampaignNegativeKeyword {
	this := SponsoredProductsCreateCampaignNegativeKeyword{}
	this.CampaignId = campaignId
	this.MatchType = matchType
	this.State = state
	this.KeywordText = keywordText
	return &this
}

// NewSponsoredProductsCreateCampaignNegativeKeywordWithDefaults instantiates a new SponsoredProductsCreateCampaignNegativeKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateCampaignNegativeKeywordWithDefaults() *SponsoredProductsCreateCampaignNegativeKeyword {
	this := SponsoredProductsCreateCampaignNegativeKeyword{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateCampaignNegativeKeyword) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateCampaignNegativeKeyword) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateCampaignNegativeKeyword) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetMatchType returns the MatchType field value
func (o *SponsoredProductsCreateCampaignNegativeKeyword) GetMatchType() SponsoredProductsCreateOrUpdateNegativeMatchType {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateNegativeMatchType
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateCampaignNegativeKeyword) GetMatchTypeOk() (*SponsoredProductsCreateOrUpdateNegativeMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *SponsoredProductsCreateCampaignNegativeKeyword) SetMatchType(v SponsoredProductsCreateOrUpdateNegativeMatchType) {
	o.MatchType = v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateCampaignNegativeKeyword) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateCampaignNegativeKeyword) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateCampaignNegativeKeyword) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = v
}

// GetKeywordText returns the KeywordText field value
func (o *SponsoredProductsCreateCampaignNegativeKeyword) GetKeywordText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateCampaignNegativeKeyword) GetKeywordTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordText, true
}

// SetKeywordText sets field value
func (o *SponsoredProductsCreateCampaignNegativeKeyword) SetKeywordText(v string) {
	o.KeywordText = v
}

func (o SponsoredProductsCreateCampaignNegativeKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["matchType"] = o.MatchType
	toSerialize["state"] = o.State
	toSerialize["keywordText"] = o.KeywordText
	return toSerialize, nil
}

type NullableSponsoredProductsCreateCampaignNegativeKeyword struct {
	value *SponsoredProductsCreateCampaignNegativeKeyword
	isSet bool
}

func (v NullableSponsoredProductsCreateCampaignNegativeKeyword) Get() *SponsoredProductsCreateCampaignNegativeKeyword {
	return v.value
}

func (v *NullableSponsoredProductsCreateCampaignNegativeKeyword) Set(val *SponsoredProductsCreateCampaignNegativeKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateCampaignNegativeKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateCampaignNegativeKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateCampaignNegativeKeyword(val *SponsoredProductsCreateCampaignNegativeKeyword) *NullableSponsoredProductsCreateCampaignNegativeKeyword {
	return &NullableSponsoredProductsCreateCampaignNegativeKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateCampaignNegativeKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateCampaignNegativeKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
