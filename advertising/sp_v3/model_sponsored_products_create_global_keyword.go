package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateGlobalKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateGlobalKeyword{}

// SponsoredProductsCreateGlobalKeyword struct for SponsoredProductsCreateGlobalKeyword
type SponsoredProductsCreateGlobalKeyword struct {
	// The identifer of the campaign to which the keyword is associated.
	CampaignId string                                   `json:"campaignId"`
	MatchType  SponsoredProductsCreateOrUpdateMatchType `json:"matchType"`
	// Name for the Keyword
	Name  *string                                          `json:"name,omitempty"`
	State SponsoredProductsCreateOrUpdateGlobalEntityState `json:"state"`
	Bid   *SponsoredProductsGlobalBid                      `json:"bid,omitempty"`
	// The identifier of the ad group to which this keyword is associated.
	AdGroupId   string                             `json:"adGroupId"`
	KeywordText SponsoredProductsGlobalKeywordText `json:"keywordText"`
}

type _SponsoredProductsCreateGlobalKeyword SponsoredProductsCreateGlobalKeyword

// NewSponsoredProductsCreateGlobalKeyword instantiates a new SponsoredProductsCreateGlobalKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateGlobalKeyword(campaignId string, matchType SponsoredProductsCreateOrUpdateMatchType, state SponsoredProductsCreateOrUpdateGlobalEntityState, adGroupId string, keywordText SponsoredProductsGlobalKeywordText) *SponsoredProductsCreateGlobalKeyword {
	this := SponsoredProductsCreateGlobalKeyword{}
	this.CampaignId = campaignId
	this.MatchType = matchType
	this.State = state
	this.AdGroupId = adGroupId
	this.KeywordText = keywordText
	return &this
}

// NewSponsoredProductsCreateGlobalKeywordWithDefaults instantiates a new SponsoredProductsCreateGlobalKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateGlobalKeywordWithDefaults() *SponsoredProductsCreateGlobalKeyword {
	this := SponsoredProductsCreateGlobalKeyword{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateGlobalKeyword) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalKeyword) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateGlobalKeyword) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetMatchType returns the MatchType field value
func (o *SponsoredProductsCreateGlobalKeyword) GetMatchType() SponsoredProductsCreateOrUpdateMatchType {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateMatchType
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalKeyword) GetMatchTypeOk() (*SponsoredProductsCreateOrUpdateMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *SponsoredProductsCreateGlobalKeyword) SetMatchType(v SponsoredProductsCreateOrUpdateMatchType) {
	o.MatchType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsCreateGlobalKeyword) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalKeyword) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsCreateGlobalKeyword) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsCreateGlobalKeyword) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateGlobalKeyword) GetState() SponsoredProductsCreateOrUpdateGlobalEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateGlobalEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalKeyword) GetStateOk() (*SponsoredProductsCreateOrUpdateGlobalEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateGlobalKeyword) SetState(v SponsoredProductsCreateOrUpdateGlobalEntityState) {
	o.State = v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *SponsoredProductsCreateGlobalKeyword) GetBid() SponsoredProductsGlobalBid {
	if o == nil || IsNil(o.Bid) {
		var ret SponsoredProductsGlobalBid
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalKeyword) GetBidOk() (*SponsoredProductsGlobalBid, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsCreateGlobalKeyword) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given SponsoredProductsGlobalBid and assigns it to the Bid field.
func (o *SponsoredProductsCreateGlobalKeyword) SetBid(v SponsoredProductsGlobalBid) {
	o.Bid = &v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsCreateGlobalKeyword) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalKeyword) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsCreateGlobalKeyword) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetKeywordText returns the KeywordText field value
func (o *SponsoredProductsCreateGlobalKeyword) GetKeywordText() SponsoredProductsGlobalKeywordText {
	if o == nil {
		var ret SponsoredProductsGlobalKeywordText
		return ret
	}

	return o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalKeyword) GetKeywordTextOk() (*SponsoredProductsGlobalKeywordText, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordText, true
}

// SetKeywordText sets field value
func (o *SponsoredProductsCreateGlobalKeyword) SetKeywordText(v SponsoredProductsGlobalKeywordText) {
	o.KeywordText = v
}

func (o SponsoredProductsCreateGlobalKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["matchType"] = o.MatchType
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["state"] = o.State
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["keywordText"] = o.KeywordText
	return toSerialize, nil
}

type NullableSponsoredProductsCreateGlobalKeyword struct {
	value *SponsoredProductsCreateGlobalKeyword
	isSet bool
}

func (v NullableSponsoredProductsCreateGlobalKeyword) Get() *SponsoredProductsCreateGlobalKeyword {
	return v.value
}

func (v *NullableSponsoredProductsCreateGlobalKeyword) Set(val *SponsoredProductsCreateGlobalKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateGlobalKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateGlobalKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateGlobalKeyword(val *SponsoredProductsCreateGlobalKeyword) *NullableSponsoredProductsCreateGlobalKeyword {
	return &NullableSponsoredProductsCreateGlobalKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateGlobalKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateGlobalKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
