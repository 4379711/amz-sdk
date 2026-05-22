package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBKeywordRecommendationThemeSuggestion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBKeywordRecommendationThemeSuggestion{}

// SBKeywordRecommendationThemeSuggestion struct for SBKeywordRecommendationThemeSuggestion
type SBKeywordRecommendationThemeSuggestion struct {
	Keywords []SBKeywordRecommendationThemeKeyword `json:"keywords,omitempty"`
	Type     *SBKeywordRecommendationThemeType     `json:"type,omitempty"`
}

// NewSBKeywordRecommendationThemeSuggestion instantiates a new SBKeywordRecommendationThemeSuggestion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBKeywordRecommendationThemeSuggestion() *SBKeywordRecommendationThemeSuggestion {
	this := SBKeywordRecommendationThemeSuggestion{}
	return &this
}

// NewSBKeywordRecommendationThemeSuggestionWithDefaults instantiates a new SBKeywordRecommendationThemeSuggestion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBKeywordRecommendationThemeSuggestionWithDefaults() *SBKeywordRecommendationThemeSuggestion {
	this := SBKeywordRecommendationThemeSuggestion{}
	return &this
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *SBKeywordRecommendationThemeSuggestion) GetKeywords() []SBKeywordRecommendationThemeKeyword {
	if o == nil || IsNil(o.Keywords) {
		var ret []SBKeywordRecommendationThemeKeyword
		return ret
	}
	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBKeywordRecommendationThemeSuggestion) GetKeywordsOk() ([]SBKeywordRecommendationThemeKeyword, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *SBKeywordRecommendationThemeSuggestion) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given []SBKeywordRecommendationThemeKeyword and assigns it to the Keywords field.
func (o *SBKeywordRecommendationThemeSuggestion) SetKeywords(v []SBKeywordRecommendationThemeKeyword) {
	o.Keywords = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SBKeywordRecommendationThemeSuggestion) GetType() SBKeywordRecommendationThemeType {
	if o == nil || IsNil(o.Type) {
		var ret SBKeywordRecommendationThemeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBKeywordRecommendationThemeSuggestion) GetTypeOk() (*SBKeywordRecommendationThemeType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SBKeywordRecommendationThemeSuggestion) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SBKeywordRecommendationThemeType and assigns it to the Type field.
func (o *SBKeywordRecommendationThemeSuggestion) SetType(v SBKeywordRecommendationThemeType) {
	o.Type = &v
}

func (o SBKeywordRecommendationThemeSuggestion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Keywords) {
		toSerialize["keywords"] = o.Keywords
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableSBKeywordRecommendationThemeSuggestion struct {
	value *SBKeywordRecommendationThemeSuggestion
	isSet bool
}

func (v NullableSBKeywordRecommendationThemeSuggestion) Get() *SBKeywordRecommendationThemeSuggestion {
	return v.value
}

func (v *NullableSBKeywordRecommendationThemeSuggestion) Set(val *SBKeywordRecommendationThemeSuggestion) {
	v.value = val
	v.isSet = true
}

func (v NullableSBKeywordRecommendationThemeSuggestion) IsSet() bool {
	return v.isSet
}

func (v *NullableSBKeywordRecommendationThemeSuggestion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBKeywordRecommendationThemeSuggestion(val *SBKeywordRecommendationThemeSuggestion) *NullableSBKeywordRecommendationThemeSuggestion {
	return &NullableSBKeywordRecommendationThemeSuggestion{value: val, isSet: true}
}

func (v NullableSBKeywordRecommendationThemeSuggestion) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBKeywordRecommendationThemeSuggestion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
