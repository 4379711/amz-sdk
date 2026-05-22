package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDThemeRecommendationsV34 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDThemeRecommendationsV34{}

// SDThemeRecommendationsV34 struct for SDThemeRecommendationsV34
type SDThemeRecommendationsV34 struct {
	Themes *SDThemeRecommendationsV34Themes `json:"themes,omitempty"`
}

// NewSDThemeRecommendationsV34 instantiates a new SDThemeRecommendationsV34 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDThemeRecommendationsV34() *SDThemeRecommendationsV34 {
	this := SDThemeRecommendationsV34{}
	return &this
}

// NewSDThemeRecommendationsV34WithDefaults instantiates a new SDThemeRecommendationsV34 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDThemeRecommendationsV34WithDefaults() *SDThemeRecommendationsV34 {
	this := SDThemeRecommendationsV34{}
	return &this
}

// GetThemes returns the Themes field value if set, zero value otherwise.
func (o *SDThemeRecommendationsV34) GetThemes() SDThemeRecommendationsV34Themes {
	if o == nil || IsNil(o.Themes) {
		var ret SDThemeRecommendationsV34Themes
		return ret
	}
	return *o.Themes
}

// GetThemesOk returns a tuple with the Themes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDThemeRecommendationsV34) GetThemesOk() (*SDThemeRecommendationsV34Themes, bool) {
	if o == nil || IsNil(o.Themes) {
		return nil, false
	}
	return o.Themes, true
}

// HasThemes returns a boolean if a field has been set.
func (o *SDThemeRecommendationsV34) HasThemes() bool {
	if o != nil && !IsNil(o.Themes) {
		return true
	}

	return false
}

// SetThemes gets a reference to the given SDThemeRecommendationsV34Themes and assigns it to the Themes field.
func (o *SDThemeRecommendationsV34) SetThemes(v SDThemeRecommendationsV34Themes) {
	o.Themes = &v
}

func (o SDThemeRecommendationsV34) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Themes) {
		toSerialize["themes"] = o.Themes
	}
	return toSerialize, nil
}

type NullableSDThemeRecommendationsV34 struct {
	value *SDThemeRecommendationsV34
	isSet bool
}

func (v NullableSDThemeRecommendationsV34) Get() *SDThemeRecommendationsV34 {
	return v.value
}

func (v *NullableSDThemeRecommendationsV34) Set(val *SDThemeRecommendationsV34) {
	v.value = val
	v.isSet = true
}

func (v NullableSDThemeRecommendationsV34) IsSet() bool {
	return v.isSet
}

func (v *NullableSDThemeRecommendationsV34) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDThemeRecommendationsV34(val *SDThemeRecommendationsV34) *NullableSDThemeRecommendationsV34 {
	return &NullableSDThemeRecommendationsV34{value: val, isSet: true}
}

func (v NullableSDThemeRecommendationsV34) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDThemeRecommendationsV34) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
