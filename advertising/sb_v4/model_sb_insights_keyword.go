package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBInsightsKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBInsightsKeyword{}

// SBInsightsKeyword Keyword associated with the campaign.
type SBInsightsKeyword struct {
	MatchType SBInsightsMatchType `json:"matchType"`
	// The associated bid. Note that this value must be less than the budget associated with the Advertiser account. For more information, see [supported features](https://advertising.amazon.com/API/docs/v2/guides/supported_features).
	Bid float64 `json:"bid"`
	// The keyword text. Maximum of 10 words.
	KeywordText string `json:"keywordText"`
}

type _SBInsightsKeyword SBInsightsKeyword

// NewSBInsightsKeyword instantiates a new SBInsightsKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBInsightsKeyword(matchType SBInsightsMatchType, bid float64, keywordText string) *SBInsightsKeyword {
	this := SBInsightsKeyword{}
	this.MatchType = matchType
	this.Bid = bid
	this.KeywordText = keywordText
	return &this
}

// NewSBInsightsKeywordWithDefaults instantiates a new SBInsightsKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBInsightsKeywordWithDefaults() *SBInsightsKeyword {
	this := SBInsightsKeyword{}
	return &this
}

// GetMatchType returns the MatchType field value
func (o *SBInsightsKeyword) GetMatchType() SBInsightsMatchType {
	if o == nil {
		var ret SBInsightsMatchType
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *SBInsightsKeyword) GetMatchTypeOk() (*SBInsightsMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *SBInsightsKeyword) SetMatchType(v SBInsightsMatchType) {
	o.MatchType = v
}

// GetBid returns the Bid field value
func (o *SBInsightsKeyword) GetBid() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Bid
}

// GetBidOk returns a tuple with the Bid field value
// and a boolean to check if the value has been set.
func (o *SBInsightsKeyword) GetBidOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bid, true
}

// SetBid sets field value
func (o *SBInsightsKeyword) SetBid(v float64) {
	o.Bid = v
}

// GetKeywordText returns the KeywordText field value
func (o *SBInsightsKeyword) GetKeywordText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value
// and a boolean to check if the value has been set.
func (o *SBInsightsKeyword) GetKeywordTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordText, true
}

// SetKeywordText sets field value
func (o *SBInsightsKeyword) SetKeywordText(v string) {
	o.KeywordText = v
}

func (o SBInsightsKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["matchType"] = o.MatchType
	toSerialize["bid"] = o.Bid
	toSerialize["keywordText"] = o.KeywordText
	return toSerialize, nil
}

type NullableSBInsightsKeyword struct {
	value *SBInsightsKeyword
	isSet bool
}

func (v NullableSBInsightsKeyword) Get() *SBInsightsKeyword {
	return v.value
}

func (v *NullableSBInsightsKeyword) Set(val *SBInsightsKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSBInsightsKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSBInsightsKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBInsightsKeyword(val *SBInsightsKeyword) *NullableSBInsightsKeyword {
	return &NullableSBInsightsKeyword{value: val, isSet: true}
}

func (v NullableSBInsightsKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBInsightsKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
