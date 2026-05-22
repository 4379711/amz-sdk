package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBKeywordRecommendationThemes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBKeywordRecommendationThemes{}

// SBKeywordRecommendationThemes struct for SBKeywordRecommendationThemes
type SBKeywordRecommendationThemes struct {
	ThemeType *SBKeywordRecommendationThemeType `json:"themeType,omitempty"`
}

// NewSBKeywordRecommendationThemes instantiates a new SBKeywordRecommendationThemes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBKeywordRecommendationThemes() *SBKeywordRecommendationThemes {
	this := SBKeywordRecommendationThemes{}
	return &this
}

// NewSBKeywordRecommendationThemesWithDefaults instantiates a new SBKeywordRecommendationThemes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBKeywordRecommendationThemesWithDefaults() *SBKeywordRecommendationThemes {
	this := SBKeywordRecommendationThemes{}
	return &this
}

// GetThemeType returns the ThemeType field value if set, zero value otherwise.
func (o *SBKeywordRecommendationThemes) GetThemeType() SBKeywordRecommendationThemeType {
	if o == nil || IsNil(o.ThemeType) {
		var ret SBKeywordRecommendationThemeType
		return ret
	}
	return *o.ThemeType
}

// GetThemeTypeOk returns a tuple with the ThemeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBKeywordRecommendationThemes) GetThemeTypeOk() (*SBKeywordRecommendationThemeType, bool) {
	if o == nil || IsNil(o.ThemeType) {
		return nil, false
	}
	return o.ThemeType, true
}

// HasThemeType returns a boolean if a field has been set.
func (o *SBKeywordRecommendationThemes) HasThemeType() bool {
	if o != nil && !IsNil(o.ThemeType) {
		return true
	}

	return false
}

// SetThemeType gets a reference to the given SBKeywordRecommendationThemeType and assigns it to the ThemeType field.
func (o *SBKeywordRecommendationThemes) SetThemeType(v SBKeywordRecommendationThemeType) {
	o.ThemeType = &v
}

func (o SBKeywordRecommendationThemes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ThemeType) {
		toSerialize["themeType"] = o.ThemeType
	}
	return toSerialize, nil
}

type NullableSBKeywordRecommendationThemes struct {
	value *SBKeywordRecommendationThemes
	isSet bool
}

func (v NullableSBKeywordRecommendationThemes) Get() *SBKeywordRecommendationThemes {
	return v.value
}

func (v *NullableSBKeywordRecommendationThemes) Set(val *SBKeywordRecommendationThemes) {
	v.value = val
	v.isSet = true
}

func (v NullableSBKeywordRecommendationThemes) IsSet() bool {
	return v.isSet
}

func (v *NullableSBKeywordRecommendationThemes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBKeywordRecommendationThemes(val *SBKeywordRecommendationThemes) *NullableSBKeywordRecommendationThemes {
	return &NullableSBKeywordRecommendationThemes{value: val, isSet: true}
}

func (v NullableSBKeywordRecommendationThemes) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBKeywordRecommendationThemes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
