package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingAdGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingAdGroup{}

// SBForecastingAdGroup The ad group settings.
type SBForecastingAdGroup struct {
	Targets          []SBForecastingProductTarget         `json:"targets,omitempty"`
	NegativeTargets  []SBForecastingNegativeProductTarget `json:"negativeTargets,omitempty"`
	LandingPages     []SBForecastingLandingPageObject     `json:"landingPages,omitempty"`
	Themes           []SBForecastingTheme                 `json:"themes,omitempty"`
	Keywords         []SBForecastingKeyword               `json:"keywords,omitempty"`
	NegativeKeywords []SBForecastingNegativeKeyword       `json:"negativeKeywords,omitempty"`
	CreativeAsins    []string                             `json:"creativeAsins,omitempty"`
}

// NewSBForecastingAdGroup instantiates a new SBForecastingAdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingAdGroup() *SBForecastingAdGroup {
	this := SBForecastingAdGroup{}
	return &this
}

// NewSBForecastingAdGroupWithDefaults instantiates a new SBForecastingAdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingAdGroupWithDefaults() *SBForecastingAdGroup {
	this := SBForecastingAdGroup{}
	return &this
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *SBForecastingAdGroup) GetTargets() []SBForecastingProductTarget {
	if o == nil || IsNil(o.Targets) {
		var ret []SBForecastingProductTarget
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingAdGroup) GetTargetsOk() ([]SBForecastingProductTarget, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *SBForecastingAdGroup) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []SBForecastingProductTarget and assigns it to the Targets field.
func (o *SBForecastingAdGroup) SetTargets(v []SBForecastingProductTarget) {
	o.Targets = v
}

// GetNegativeTargets returns the NegativeTargets field value if set, zero value otherwise.
func (o *SBForecastingAdGroup) GetNegativeTargets() []SBForecastingNegativeProductTarget {
	if o == nil || IsNil(o.NegativeTargets) {
		var ret []SBForecastingNegativeProductTarget
		return ret
	}
	return o.NegativeTargets
}

// GetNegativeTargetsOk returns a tuple with the NegativeTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingAdGroup) GetNegativeTargetsOk() ([]SBForecastingNegativeProductTarget, bool) {
	if o == nil || IsNil(o.NegativeTargets) {
		return nil, false
	}
	return o.NegativeTargets, true
}

// HasNegativeTargets returns a boolean if a field has been set.
func (o *SBForecastingAdGroup) HasNegativeTargets() bool {
	if o != nil && !IsNil(o.NegativeTargets) {
		return true
	}

	return false
}

// SetNegativeTargets gets a reference to the given []SBForecastingNegativeProductTarget and assigns it to the NegativeTargets field.
func (o *SBForecastingAdGroup) SetNegativeTargets(v []SBForecastingNegativeProductTarget) {
	o.NegativeTargets = v
}

// GetLandingPages returns the LandingPages field value if set, zero value otherwise.
func (o *SBForecastingAdGroup) GetLandingPages() []SBForecastingLandingPageObject {
	if o == nil || IsNil(o.LandingPages) {
		var ret []SBForecastingLandingPageObject
		return ret
	}
	return o.LandingPages
}

// GetLandingPagesOk returns a tuple with the LandingPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingAdGroup) GetLandingPagesOk() ([]SBForecastingLandingPageObject, bool) {
	if o == nil || IsNil(o.LandingPages) {
		return nil, false
	}
	return o.LandingPages, true
}

// HasLandingPages returns a boolean if a field has been set.
func (o *SBForecastingAdGroup) HasLandingPages() bool {
	if o != nil && !IsNil(o.LandingPages) {
		return true
	}

	return false
}

// SetLandingPages gets a reference to the given []SBForecastingLandingPageObject and assigns it to the LandingPages field.
func (o *SBForecastingAdGroup) SetLandingPages(v []SBForecastingLandingPageObject) {
	o.LandingPages = v
}

// GetThemes returns the Themes field value if set, zero value otherwise.
func (o *SBForecastingAdGroup) GetThemes() []SBForecastingTheme {
	if o == nil || IsNil(o.Themes) {
		var ret []SBForecastingTheme
		return ret
	}
	return o.Themes
}

// GetThemesOk returns a tuple with the Themes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingAdGroup) GetThemesOk() ([]SBForecastingTheme, bool) {
	if o == nil || IsNil(o.Themes) {
		return nil, false
	}
	return o.Themes, true
}

// HasThemes returns a boolean if a field has been set.
func (o *SBForecastingAdGroup) HasThemes() bool {
	if o != nil && !IsNil(o.Themes) {
		return true
	}

	return false
}

// SetThemes gets a reference to the given []SBForecastingTheme and assigns it to the Themes field.
func (o *SBForecastingAdGroup) SetThemes(v []SBForecastingTheme) {
	o.Themes = v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *SBForecastingAdGroup) GetKeywords() []SBForecastingKeyword {
	if o == nil || IsNil(o.Keywords) {
		var ret []SBForecastingKeyword
		return ret
	}
	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingAdGroup) GetKeywordsOk() ([]SBForecastingKeyword, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *SBForecastingAdGroup) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given []SBForecastingKeyword and assigns it to the Keywords field.
func (o *SBForecastingAdGroup) SetKeywords(v []SBForecastingKeyword) {
	o.Keywords = v
}

// GetNegativeKeywords returns the NegativeKeywords field value if set, zero value otherwise.
func (o *SBForecastingAdGroup) GetNegativeKeywords() []SBForecastingNegativeKeyword {
	if o == nil || IsNil(o.NegativeKeywords) {
		var ret []SBForecastingNegativeKeyword
		return ret
	}
	return o.NegativeKeywords
}

// GetNegativeKeywordsOk returns a tuple with the NegativeKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingAdGroup) GetNegativeKeywordsOk() ([]SBForecastingNegativeKeyword, bool) {
	if o == nil || IsNil(o.NegativeKeywords) {
		return nil, false
	}
	return o.NegativeKeywords, true
}

// HasNegativeKeywords returns a boolean if a field has been set.
func (o *SBForecastingAdGroup) HasNegativeKeywords() bool {
	if o != nil && !IsNil(o.NegativeKeywords) {
		return true
	}

	return false
}

// SetNegativeKeywords gets a reference to the given []SBForecastingNegativeKeyword and assigns it to the NegativeKeywords field.
func (o *SBForecastingAdGroup) SetNegativeKeywords(v []SBForecastingNegativeKeyword) {
	o.NegativeKeywords = v
}

// GetCreativeAsins returns the CreativeAsins field value if set, zero value otherwise.
func (o *SBForecastingAdGroup) GetCreativeAsins() []string {
	if o == nil || IsNil(o.CreativeAsins) {
		var ret []string
		return ret
	}
	return o.CreativeAsins
}

// GetCreativeAsinsOk returns a tuple with the CreativeAsins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingAdGroup) GetCreativeAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.CreativeAsins) {
		return nil, false
	}
	return o.CreativeAsins, true
}

// HasCreativeAsins returns a boolean if a field has been set.
func (o *SBForecastingAdGroup) HasCreativeAsins() bool {
	if o != nil && !IsNil(o.CreativeAsins) {
		return true
	}

	return false
}

// SetCreativeAsins gets a reference to the given []string and assigns it to the CreativeAsins field.
func (o *SBForecastingAdGroup) SetCreativeAsins(v []string) {
	o.CreativeAsins = v
}

func (o SBForecastingAdGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	if !IsNil(o.NegativeTargets) {
		toSerialize["negativeTargets"] = o.NegativeTargets
	}
	if !IsNil(o.LandingPages) {
		toSerialize["landingPages"] = o.LandingPages
	}
	if !IsNil(o.Themes) {
		toSerialize["themes"] = o.Themes
	}
	if !IsNil(o.Keywords) {
		toSerialize["keywords"] = o.Keywords
	}
	if !IsNil(o.NegativeKeywords) {
		toSerialize["negativeKeywords"] = o.NegativeKeywords
	}
	if !IsNil(o.CreativeAsins) {
		toSerialize["creativeAsins"] = o.CreativeAsins
	}
	return toSerialize, nil
}

type NullableSBForecastingAdGroup struct {
	value *SBForecastingAdGroup
	isSet bool
}

func (v NullableSBForecastingAdGroup) Get() *SBForecastingAdGroup {
	return v.value
}

func (v *NullableSBForecastingAdGroup) Set(val *SBForecastingAdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingAdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingAdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingAdGroup(val *SBForecastingAdGroup) *NullableSBForecastingAdGroup {
	return &NullableSBForecastingAdGroup{value: val, isSet: true}
}

func (v NullableSBForecastingAdGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingAdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
