package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingKeyword{}

// SBForecastingKeyword Keyword associated with the campaign.
type SBForecastingKeyword struct {
	// The keyword text. Maximum of 10 words.
	KeywordText *string `json:"keywordText,omitempty"`
	// The match type. Valid value: EXACT, PHRASE, BROAD. For more information, see [match types](https://advertising.amazon.com/help#GHTRFDZRJPW6764R) in the Amazon Advertising support center.
	MatchType *string `json:"matchType,omitempty"`
	// The associated bid. Note that this value must be less than the budget associated with the Advertiser account.
	Bid *float32 `json:"bid,omitempty"`
}

// NewSBForecastingKeyword instantiates a new SBForecastingKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingKeyword() *SBForecastingKeyword {
	this := SBForecastingKeyword{}
	return &this
}

// NewSBForecastingKeywordWithDefaults instantiates a new SBForecastingKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingKeywordWithDefaults() *SBForecastingKeyword {
	this := SBForecastingKeyword{}
	return &this
}

// GetKeywordText returns the KeywordText field value if set, zero value otherwise.
func (o *SBForecastingKeyword) GetKeywordText() string {
	if o == nil || IsNil(o.KeywordText) {
		var ret string
		return ret
	}
	return *o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingKeyword) GetKeywordTextOk() (*string, bool) {
	if o == nil || IsNil(o.KeywordText) {
		return nil, false
	}
	return o.KeywordText, true
}

// HasKeywordText returns a boolean if a field has been set.
func (o *SBForecastingKeyword) HasKeywordText() bool {
	if o != nil && !IsNil(o.KeywordText) {
		return true
	}

	return false
}

// SetKeywordText gets a reference to the given string and assigns it to the KeywordText field.
func (o *SBForecastingKeyword) SetKeywordText(v string) {
	o.KeywordText = &v
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *SBForecastingKeyword) GetMatchType() string {
	if o == nil || IsNil(o.MatchType) {
		var ret string
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingKeyword) GetMatchTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MatchType) {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *SBForecastingKeyword) HasMatchType() bool {
	if o != nil && !IsNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given string and assigns it to the MatchType field.
func (o *SBForecastingKeyword) SetMatchType(v string) {
	o.MatchType = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *SBForecastingKeyword) GetBid() float32 {
	if o == nil || IsNil(o.Bid) {
		var ret float32
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingKeyword) GetBidOk() (*float32, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *SBForecastingKeyword) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given float32 and assigns it to the Bid field.
func (o *SBForecastingKeyword) SetBid(v float32) {
	o.Bid = &v
}

func (o SBForecastingKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeywordText) {
		toSerialize["keywordText"] = o.KeywordText
	}
	if !IsNil(o.MatchType) {
		toSerialize["matchType"] = o.MatchType
	}
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	return toSerialize, nil
}

type NullableSBForecastingKeyword struct {
	value *SBForecastingKeyword
	isSet bool
}

func (v NullableSBForecastingKeyword) Get() *SBForecastingKeyword {
	return v.value
}

func (v *NullableSBForecastingKeyword) Set(val *SBForecastingKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingKeyword(val *SBForecastingKeyword) *NullableSBForecastingKeyword {
	return &NullableSBForecastingKeyword{value: val, isSet: true}
}

func (v NullableSBForecastingKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
