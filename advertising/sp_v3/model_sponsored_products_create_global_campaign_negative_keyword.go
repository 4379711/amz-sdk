package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateGlobalCampaignNegativeKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateGlobalCampaignNegativeKeyword{}

// SponsoredProductsCreateGlobalCampaignNegativeKeyword struct for SponsoredProductsCreateGlobalCampaignNegativeKeyword
type SponsoredProductsCreateGlobalCampaignNegativeKeyword struct {
	// The identifier of the campaign to which the keyword is associated.
	CampaignId string                                           `json:"campaignId"`
	MatchType  SponsoredProductsCreateOrUpdateNegativeMatchType `json:"matchType"`
	// Name for the keyword
	Name        *string                                          `json:"name,omitempty"`
	State       SponsoredProductsCreateOrUpdateGlobalEntityState `json:"state"`
	KeywordText SponsoredProductsGlobalNegativeKeywordText       `json:"keywordText"`
}

type _SponsoredProductsCreateGlobalCampaignNegativeKeyword SponsoredProductsCreateGlobalCampaignNegativeKeyword

// NewSponsoredProductsCreateGlobalCampaignNegativeKeyword instantiates a new SponsoredProductsCreateGlobalCampaignNegativeKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateGlobalCampaignNegativeKeyword(campaignId string, matchType SponsoredProductsCreateOrUpdateNegativeMatchType, state SponsoredProductsCreateOrUpdateGlobalEntityState, keywordText SponsoredProductsGlobalNegativeKeywordText) *SponsoredProductsCreateGlobalCampaignNegativeKeyword {
	this := SponsoredProductsCreateGlobalCampaignNegativeKeyword{}
	this.CampaignId = campaignId
	this.MatchType = matchType
	this.State = state
	this.KeywordText = keywordText
	return &this
}

// NewSponsoredProductsCreateGlobalCampaignNegativeKeywordWithDefaults instantiates a new SponsoredProductsCreateGlobalCampaignNegativeKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateGlobalCampaignNegativeKeywordWithDefaults() *SponsoredProductsCreateGlobalCampaignNegativeKeyword {
	this := SponsoredProductsCreateGlobalCampaignNegativeKeyword{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateGlobalCampaignNegativeKeyword) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalCampaignNegativeKeyword) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateGlobalCampaignNegativeKeyword) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetMatchType returns the MatchType field value
func (o *SponsoredProductsCreateGlobalCampaignNegativeKeyword) GetMatchType() SponsoredProductsCreateOrUpdateNegativeMatchType {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateNegativeMatchType
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalCampaignNegativeKeyword) GetMatchTypeOk() (*SponsoredProductsCreateOrUpdateNegativeMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *SponsoredProductsCreateGlobalCampaignNegativeKeyword) SetMatchType(v SponsoredProductsCreateOrUpdateNegativeMatchType) {
	o.MatchType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsCreateGlobalCampaignNegativeKeyword) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalCampaignNegativeKeyword) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsCreateGlobalCampaignNegativeKeyword) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsCreateGlobalCampaignNegativeKeyword) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateGlobalCampaignNegativeKeyword) GetState() SponsoredProductsCreateOrUpdateGlobalEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateGlobalEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalCampaignNegativeKeyword) GetStateOk() (*SponsoredProductsCreateOrUpdateGlobalEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateGlobalCampaignNegativeKeyword) SetState(v SponsoredProductsCreateOrUpdateGlobalEntityState) {
	o.State = v
}

// GetKeywordText returns the KeywordText field value
func (o *SponsoredProductsCreateGlobalCampaignNegativeKeyword) GetKeywordText() SponsoredProductsGlobalNegativeKeywordText {
	if o == nil {
		var ret SponsoredProductsGlobalNegativeKeywordText
		return ret
	}

	return o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalCampaignNegativeKeyword) GetKeywordTextOk() (*SponsoredProductsGlobalNegativeKeywordText, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordText, true
}

// SetKeywordText sets field value
func (o *SponsoredProductsCreateGlobalCampaignNegativeKeyword) SetKeywordText(v SponsoredProductsGlobalNegativeKeywordText) {
	o.KeywordText = v
}

func (o SponsoredProductsCreateGlobalCampaignNegativeKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["matchType"] = o.MatchType
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["state"] = o.State
	toSerialize["keywordText"] = o.KeywordText
	return toSerialize, nil
}

type NullableSponsoredProductsCreateGlobalCampaignNegativeKeyword struct {
	value *SponsoredProductsCreateGlobalCampaignNegativeKeyword
	isSet bool
}

func (v NullableSponsoredProductsCreateGlobalCampaignNegativeKeyword) Get() *SponsoredProductsCreateGlobalCampaignNegativeKeyword {
	return v.value
}

func (v *NullableSponsoredProductsCreateGlobalCampaignNegativeKeyword) Set(val *SponsoredProductsCreateGlobalCampaignNegativeKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateGlobalCampaignNegativeKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateGlobalCampaignNegativeKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateGlobalCampaignNegativeKeyword(val *SponsoredProductsCreateGlobalCampaignNegativeKeyword) *NullableSponsoredProductsCreateGlobalCampaignNegativeKeyword {
	return &NullableSponsoredProductsCreateGlobalCampaignNegativeKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateGlobalCampaignNegativeKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateGlobalCampaignNegativeKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
