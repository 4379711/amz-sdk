package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalKeyword{}

// SponsoredProductsGlobalKeyword struct for SponsoredProductsGlobalKeyword
type SponsoredProductsGlobalKeyword struct {
	// The identifier of the keyword.
	KeywordId string `json:"keywordId"`
	// The identifier of the campaign to which the keyword is associated.
	CampaignId string                     `json:"campaignId"`
	MatchType  SponsoredProductsMatchType `json:"matchType"`
	// Name for the Keyword
	Name  *string                            `json:"name,omitempty"`
	State SponsoredProductsGlobalEntityState `json:"state"`
	Bid   SponsoredProductsGlobalBid         `json:"bid"`
	// The identifier of the ad group to which this keyword is associated.
	AdGroupId    string                                      `json:"adGroupId"`
	KeywordText  SponsoredProductsGlobalKeywordText          `json:"keywordText"`
	ExtendedData *SponsoredProductsGlobalKeywordExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsGlobalKeyword SponsoredProductsGlobalKeyword

// NewSponsoredProductsGlobalKeyword instantiates a new SponsoredProductsGlobalKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalKeyword(keywordId string, campaignId string, matchType SponsoredProductsMatchType, state SponsoredProductsGlobalEntityState, bid SponsoredProductsGlobalBid, adGroupId string, keywordText SponsoredProductsGlobalKeywordText) *SponsoredProductsGlobalKeyword {
	this := SponsoredProductsGlobalKeyword{}
	this.KeywordId = keywordId
	this.CampaignId = campaignId
	this.MatchType = matchType
	this.State = state
	this.Bid = bid
	this.AdGroupId = adGroupId
	this.KeywordText = keywordText
	return &this
}

// NewSponsoredProductsGlobalKeywordWithDefaults instantiates a new SponsoredProductsGlobalKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalKeywordWithDefaults() *SponsoredProductsGlobalKeyword {
	this := SponsoredProductsGlobalKeyword{}
	return &this
}

// GetKeywordId returns the KeywordId field value
func (o *SponsoredProductsGlobalKeyword) GetKeywordId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordId
}

// GetKeywordIdOk returns a tuple with the KeywordId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeyword) GetKeywordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordId, true
}

// SetKeywordId sets field value
func (o *SponsoredProductsGlobalKeyword) SetKeywordId(v string) {
	o.KeywordId = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsGlobalKeyword) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeyword) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsGlobalKeyword) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetMatchType returns the MatchType field value
func (o *SponsoredProductsGlobalKeyword) GetMatchType() SponsoredProductsMatchType {
	if o == nil {
		var ret SponsoredProductsMatchType
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeyword) GetMatchTypeOk() (*SponsoredProductsMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *SponsoredProductsGlobalKeyword) SetMatchType(v SponsoredProductsMatchType) {
	o.MatchType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalKeyword) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeyword) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalKeyword) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsGlobalKeyword) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value
func (o *SponsoredProductsGlobalKeyword) GetState() SponsoredProductsGlobalEntityState {
	if o == nil {
		var ret SponsoredProductsGlobalEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeyword) GetStateOk() (*SponsoredProductsGlobalEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsGlobalKeyword) SetState(v SponsoredProductsGlobalEntityState) {
	o.State = v
}

// GetBid returns the Bid field value
func (o *SponsoredProductsGlobalKeyword) GetBid() SponsoredProductsGlobalBid {
	if o == nil {
		var ret SponsoredProductsGlobalBid
		return ret
	}

	return o.Bid
}

// GetBidOk returns a tuple with the Bid field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeyword) GetBidOk() (*SponsoredProductsGlobalBid, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bid, true
}

// SetBid sets field value
func (o *SponsoredProductsGlobalKeyword) SetBid(v SponsoredProductsGlobalBid) {
	o.Bid = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsGlobalKeyword) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeyword) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsGlobalKeyword) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetKeywordText returns the KeywordText field value
func (o *SponsoredProductsGlobalKeyword) GetKeywordText() SponsoredProductsGlobalKeywordText {
	if o == nil {
		var ret SponsoredProductsGlobalKeywordText
		return ret
	}

	return o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeyword) GetKeywordTextOk() (*SponsoredProductsGlobalKeywordText, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordText, true
}

// SetKeywordText sets field value
func (o *SponsoredProductsGlobalKeyword) SetKeywordText(v SponsoredProductsGlobalKeywordText) {
	o.KeywordText = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalKeyword) GetExtendedData() SponsoredProductsGlobalKeywordExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsGlobalKeywordExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeyword) GetExtendedDataOk() (*SponsoredProductsGlobalKeywordExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalKeyword) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsGlobalKeywordExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsGlobalKeyword) SetExtendedData(v SponsoredProductsGlobalKeywordExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsGlobalKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywordId"] = o.KeywordId
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["matchType"] = o.MatchType
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["state"] = o.State
	toSerialize["bid"] = o.Bid
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["keywordText"] = o.KeywordText
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalKeyword struct {
	value *SponsoredProductsGlobalKeyword
	isSet bool
}

func (v NullableSponsoredProductsGlobalKeyword) Get() *SponsoredProductsGlobalKeyword {
	return v.value
}

func (v *NullableSponsoredProductsGlobalKeyword) Set(val *SponsoredProductsGlobalKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalKeyword(val *SponsoredProductsGlobalKeyword) *NullableSponsoredProductsGlobalKeyword {
	return &NullableSponsoredProductsGlobalKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
