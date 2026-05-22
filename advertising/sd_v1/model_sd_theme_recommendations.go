package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDThemeRecommendations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDThemeRecommendations{}

// SDThemeRecommendations struct for SDThemeRecommendations
type SDThemeRecommendations struct {
	Themes *SDThemeRecommendationsThemes `json:"themes,omitempty"`
}

// NewSDThemeRecommendations instantiates a new SDThemeRecommendations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDThemeRecommendations() *SDThemeRecommendations {
	this := SDThemeRecommendations{}
	return &this
}

// NewSDThemeRecommendationsWithDefaults instantiates a new SDThemeRecommendations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDThemeRecommendationsWithDefaults() *SDThemeRecommendations {
	this := SDThemeRecommendations{}
	return &this
}

// GetThemes returns the Themes field value if set, zero value otherwise.
func (o *SDThemeRecommendations) GetThemes() SDThemeRecommendationsThemes {
	if o == nil || IsNil(o.Themes) {
		var ret SDThemeRecommendationsThemes
		return ret
	}
	return *o.Themes
}

// GetThemesOk returns a tuple with the Themes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDThemeRecommendations) GetThemesOk() (*SDThemeRecommendationsThemes, bool) {
	if o == nil || IsNil(o.Themes) {
		return nil, false
	}
	return o.Themes, true
}

// HasThemes returns a boolean if a field has been set.
func (o *SDThemeRecommendations) HasThemes() bool {
	if o != nil && !IsNil(o.Themes) {
		return true
	}

	return false
}

// SetThemes gets a reference to the given SDThemeRecommendationsThemes and assigns it to the Themes field.
func (o *SDThemeRecommendations) SetThemes(v SDThemeRecommendationsThemes) {
	o.Themes = &v
}

func (o SDThemeRecommendations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Themes) {
		toSerialize["themes"] = o.Themes
	}
	return toSerialize, nil
}

type NullableSDThemeRecommendations struct {
	value *SDThemeRecommendations
	isSet bool
}

func (v NullableSDThemeRecommendations) Get() *SDThemeRecommendations {
	return v.value
}

func (v *NullableSDThemeRecommendations) Set(val *SDThemeRecommendations) {
	v.value = val
	v.isSet = true
}

func (v NullableSDThemeRecommendations) IsSet() bool {
	return v.isSet
}

func (v *NullableSDThemeRecommendations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDThemeRecommendations(val *SDThemeRecommendations) *NullableSDThemeRecommendations {
	return &NullableSDThemeRecommendations{value: val, isSet: true}
}

func (v NullableSDThemeRecommendations) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDThemeRecommendations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
