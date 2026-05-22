package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateGlobalNegativeKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateGlobalNegativeKeyword{}

// SponsoredProductsCreateGlobalNegativeKeyword struct for SponsoredProductsCreateGlobalNegativeKeyword
type SponsoredProductsCreateGlobalNegativeKeyword struct {
	// The identifier of the campaign to which the keyword is associated.
	CampaignId string                                           `json:"campaignId"`
	MatchType  SponsoredProductsCreateOrUpdateNegativeMatchType `json:"matchType"`
	// Name for the keyword
	Name  *string                                          `json:"name,omitempty"`
	State SponsoredProductsCreateOrUpdateGlobalEntityState `json:"state"`
	// The identifier of the ad group to which this keyword is associated.
	AdGroupId   string                                     `json:"adGroupId"`
	KeywordText SponsoredProductsGlobalNegativeKeywordText `json:"keywordText"`
}

type _SponsoredProductsCreateGlobalNegativeKeyword SponsoredProductsCreateGlobalNegativeKeyword

// NewSponsoredProductsCreateGlobalNegativeKeyword instantiates a new SponsoredProductsCreateGlobalNegativeKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateGlobalNegativeKeyword(campaignId string, matchType SponsoredProductsCreateOrUpdateNegativeMatchType, state SponsoredProductsCreateOrUpdateGlobalEntityState, adGroupId string, keywordText SponsoredProductsGlobalNegativeKeywordText) *SponsoredProductsCreateGlobalNegativeKeyword {
	this := SponsoredProductsCreateGlobalNegativeKeyword{}
	this.CampaignId = campaignId
	this.MatchType = matchType
	this.State = state
	this.AdGroupId = adGroupId
	this.KeywordText = keywordText
	return &this
}

// NewSponsoredProductsCreateGlobalNegativeKeywordWithDefaults instantiates a new SponsoredProductsCreateGlobalNegativeKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateGlobalNegativeKeywordWithDefaults() *SponsoredProductsCreateGlobalNegativeKeyword {
	this := SponsoredProductsCreateGlobalNegativeKeyword{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateGlobalNegativeKeyword) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalNegativeKeyword) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateGlobalNegativeKeyword) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetMatchType returns the MatchType field value
func (o *SponsoredProductsCreateGlobalNegativeKeyword) GetMatchType() SponsoredProductsCreateOrUpdateNegativeMatchType {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateNegativeMatchType
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalNegativeKeyword) GetMatchTypeOk() (*SponsoredProductsCreateOrUpdateNegativeMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *SponsoredProductsCreateGlobalNegativeKeyword) SetMatchType(v SponsoredProductsCreateOrUpdateNegativeMatchType) {
	o.MatchType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsCreateGlobalNegativeKeyword) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalNegativeKeyword) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsCreateGlobalNegativeKeyword) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsCreateGlobalNegativeKeyword) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateGlobalNegativeKeyword) GetState() SponsoredProductsCreateOrUpdateGlobalEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateGlobalEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalNegativeKeyword) GetStateOk() (*SponsoredProductsCreateOrUpdateGlobalEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateGlobalNegativeKeyword) SetState(v SponsoredProductsCreateOrUpdateGlobalEntityState) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsCreateGlobalNegativeKeyword) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalNegativeKeyword) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsCreateGlobalNegativeKeyword) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetKeywordText returns the KeywordText field value
func (o *SponsoredProductsCreateGlobalNegativeKeyword) GetKeywordText() SponsoredProductsGlobalNegativeKeywordText {
	if o == nil {
		var ret SponsoredProductsGlobalNegativeKeywordText
		return ret
	}

	return o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalNegativeKeyword) GetKeywordTextOk() (*SponsoredProductsGlobalNegativeKeywordText, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordText, true
}

// SetKeywordText sets field value
func (o *SponsoredProductsCreateGlobalNegativeKeyword) SetKeywordText(v SponsoredProductsGlobalNegativeKeywordText) {
	o.KeywordText = v
}

func (o SponsoredProductsCreateGlobalNegativeKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["matchType"] = o.MatchType
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["state"] = o.State
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["keywordText"] = o.KeywordText
	return toSerialize, nil
}

type NullableSponsoredProductsCreateGlobalNegativeKeyword struct {
	value *SponsoredProductsCreateGlobalNegativeKeyword
	isSet bool
}

func (v NullableSponsoredProductsCreateGlobalNegativeKeyword) Get() *SponsoredProductsCreateGlobalNegativeKeyword {
	return v.value
}

func (v *NullableSponsoredProductsCreateGlobalNegativeKeyword) Set(val *SponsoredProductsCreateGlobalNegativeKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateGlobalNegativeKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateGlobalNegativeKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateGlobalNegativeKeyword(val *SponsoredProductsCreateGlobalNegativeKeyword) *NullableSponsoredProductsCreateGlobalNegativeKeyword {
	return &NullableSponsoredProductsCreateGlobalNegativeKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateGlobalNegativeKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateGlobalNegativeKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
