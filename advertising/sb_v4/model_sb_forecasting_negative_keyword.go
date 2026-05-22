package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingNegativeKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingNegativeKeyword{}

// SBForecastingNegativeKeyword Negative keyword associated with the campaign.
type SBForecastingNegativeKeyword struct {
	// The keyword text. Maximum of 10 words.
	KeywordText *string `json:"keywordText,omitempty"`
	// The negative match type. Valid value: NEGATIVE_EXACT, NEGATIVE_PHRASE. For more information, see [negative keyword match types](https://advertising.amazon.com/help#GHTRFDZRJPW6764R) in the Amazon Advertising support center.
	MatchType *string `json:"matchType,omitempty"`
}

// NewSBForecastingNegativeKeyword instantiates a new SBForecastingNegativeKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingNegativeKeyword() *SBForecastingNegativeKeyword {
	this := SBForecastingNegativeKeyword{}
	return &this
}

// NewSBForecastingNegativeKeywordWithDefaults instantiates a new SBForecastingNegativeKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingNegativeKeywordWithDefaults() *SBForecastingNegativeKeyword {
	this := SBForecastingNegativeKeyword{}
	return &this
}

// GetKeywordText returns the KeywordText field value if set, zero value otherwise.
func (o *SBForecastingNegativeKeyword) GetKeywordText() string {
	if o == nil || IsNil(o.KeywordText) {
		var ret string
		return ret
	}
	return *o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingNegativeKeyword) GetKeywordTextOk() (*string, bool) {
	if o == nil || IsNil(o.KeywordText) {
		return nil, false
	}
	return o.KeywordText, true
}

// HasKeywordText returns a boolean if a field has been set.
func (o *SBForecastingNegativeKeyword) HasKeywordText() bool {
	if o != nil && !IsNil(o.KeywordText) {
		return true
	}

	return false
}

// SetKeywordText gets a reference to the given string and assigns it to the KeywordText field.
func (o *SBForecastingNegativeKeyword) SetKeywordText(v string) {
	o.KeywordText = &v
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *SBForecastingNegativeKeyword) GetMatchType() string {
	if o == nil || IsNil(o.MatchType) {
		var ret string
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingNegativeKeyword) GetMatchTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MatchType) {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *SBForecastingNegativeKeyword) HasMatchType() bool {
	if o != nil && !IsNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given string and assigns it to the MatchType field.
func (o *SBForecastingNegativeKeyword) SetMatchType(v string) {
	o.MatchType = &v
}

func (o SBForecastingNegativeKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeywordText) {
		toSerialize["keywordText"] = o.KeywordText
	}
	if !IsNil(o.MatchType) {
		toSerialize["matchType"] = o.MatchType
	}
	return toSerialize, nil
}

type NullableSBForecastingNegativeKeyword struct {
	value *SBForecastingNegativeKeyword
	isSet bool
}

func (v NullableSBForecastingNegativeKeyword) Get() *SBForecastingNegativeKeyword {
	return v.value
}

func (v *NullableSBForecastingNegativeKeyword) Set(val *SBForecastingNegativeKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingNegativeKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingNegativeKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingNegativeKeyword(val *SBForecastingNegativeKeyword) *NullableSBForecastingNegativeKeyword {
	return &NullableSBForecastingNegativeKeyword{value: val, isSet: true}
}

func (v NullableSBForecastingNegativeKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingNegativeKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
