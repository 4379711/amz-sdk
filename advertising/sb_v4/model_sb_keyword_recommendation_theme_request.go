package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBKeywordRecommendationThemeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBKeywordRecommendationThemeRequest{}

// SBKeywordRecommendationThemeRequest struct for SBKeywordRecommendationThemeRequest
type SBKeywordRecommendationThemeRequest struct {
	Themes []SBKeywordRecommendationThemes `json:"themes,omitempty"`
	// Maximum number of suggestions to return for each theme. Max value is 1000. If not provided, default to 100.
	MaxNumSuggestions *int32                               `json:"maxNumSuggestions,omitempty"`
	LandingPages      []SBKeywordRecommendationLandingPage `json:"landingPages,omitempty"`
}

// NewSBKeywordRecommendationThemeRequest instantiates a new SBKeywordRecommendationThemeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBKeywordRecommendationThemeRequest() *SBKeywordRecommendationThemeRequest {
	this := SBKeywordRecommendationThemeRequest{}
	return &this
}

// NewSBKeywordRecommendationThemeRequestWithDefaults instantiates a new SBKeywordRecommendationThemeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBKeywordRecommendationThemeRequestWithDefaults() *SBKeywordRecommendationThemeRequest {
	this := SBKeywordRecommendationThemeRequest{}
	return &this
}

// GetThemes returns the Themes field value if set, zero value otherwise.
func (o *SBKeywordRecommendationThemeRequest) GetThemes() []SBKeywordRecommendationThemes {
	if o == nil || IsNil(o.Themes) {
		var ret []SBKeywordRecommendationThemes
		return ret
	}
	return o.Themes
}

// GetThemesOk returns a tuple with the Themes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBKeywordRecommendationThemeRequest) GetThemesOk() ([]SBKeywordRecommendationThemes, bool) {
	if o == nil || IsNil(o.Themes) {
		return nil, false
	}
	return o.Themes, true
}

// HasThemes returns a boolean if a field has been set.
func (o *SBKeywordRecommendationThemeRequest) HasThemes() bool {
	if o != nil && !IsNil(o.Themes) {
		return true
	}

	return false
}

// SetThemes gets a reference to the given []SBKeywordRecommendationThemes and assigns it to the Themes field.
func (o *SBKeywordRecommendationThemeRequest) SetThemes(v []SBKeywordRecommendationThemes) {
	o.Themes = v
}

// GetMaxNumSuggestions returns the MaxNumSuggestions field value if set, zero value otherwise.
func (o *SBKeywordRecommendationThemeRequest) GetMaxNumSuggestions() int32 {
	if o == nil || IsNil(o.MaxNumSuggestions) {
		var ret int32
		return ret
	}
	return *o.MaxNumSuggestions
}

// GetMaxNumSuggestionsOk returns a tuple with the MaxNumSuggestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBKeywordRecommendationThemeRequest) GetMaxNumSuggestionsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumSuggestions) {
		return nil, false
	}
	return o.MaxNumSuggestions, true
}

// HasMaxNumSuggestions returns a boolean if a field has been set.
func (o *SBKeywordRecommendationThemeRequest) HasMaxNumSuggestions() bool {
	if o != nil && !IsNil(o.MaxNumSuggestions) {
		return true
	}

	return false
}

// SetMaxNumSuggestions gets a reference to the given int32 and assigns it to the MaxNumSuggestions field.
func (o *SBKeywordRecommendationThemeRequest) SetMaxNumSuggestions(v int32) {
	o.MaxNumSuggestions = &v
}

// GetLandingPages returns the LandingPages field value if set, zero value otherwise.
func (o *SBKeywordRecommendationThemeRequest) GetLandingPages() []SBKeywordRecommendationLandingPage {
	if o == nil || IsNil(o.LandingPages) {
		var ret []SBKeywordRecommendationLandingPage
		return ret
	}
	return o.LandingPages
}

// GetLandingPagesOk returns a tuple with the LandingPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBKeywordRecommendationThemeRequest) GetLandingPagesOk() ([]SBKeywordRecommendationLandingPage, bool) {
	if o == nil || IsNil(o.LandingPages) {
		return nil, false
	}
	return o.LandingPages, true
}

// HasLandingPages returns a boolean if a field has been set.
func (o *SBKeywordRecommendationThemeRequest) HasLandingPages() bool {
	if o != nil && !IsNil(o.LandingPages) {
		return true
	}

	return false
}

// SetLandingPages gets a reference to the given []SBKeywordRecommendationLandingPage and assigns it to the LandingPages field.
func (o *SBKeywordRecommendationThemeRequest) SetLandingPages(v []SBKeywordRecommendationLandingPage) {
	o.LandingPages = v
}

func (o SBKeywordRecommendationThemeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Themes) {
		toSerialize["themes"] = o.Themes
	}
	if !IsNil(o.MaxNumSuggestions) {
		toSerialize["maxNumSuggestions"] = o.MaxNumSuggestions
	}
	if !IsNil(o.LandingPages) {
		toSerialize["landingPages"] = o.LandingPages
	}
	return toSerialize, nil
}

type NullableSBKeywordRecommendationThemeRequest struct {
	value *SBKeywordRecommendationThemeRequest
	isSet bool
}

func (v NullableSBKeywordRecommendationThemeRequest) Get() *SBKeywordRecommendationThemeRequest {
	return v.value
}

func (v *NullableSBKeywordRecommendationThemeRequest) Set(val *SBKeywordRecommendationThemeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSBKeywordRecommendationThemeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSBKeywordRecommendationThemeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBKeywordRecommendationThemeRequest(val *SBKeywordRecommendationThemeRequest) *NullableSBKeywordRecommendationThemeRequest {
	return &NullableSBKeywordRecommendationThemeRequest{value: val, isSet: true}
}

func (v NullableSBKeywordRecommendationThemeRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBKeywordRecommendationThemeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
