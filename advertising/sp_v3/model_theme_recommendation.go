package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the ThemeRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThemeRecommendation{}

// ThemeRecommendation Recommended asins grouped by theme attribute.
type ThemeRecommendation struct {
	// A theme name representing the context around the recommended list of ASINs.
	Description *string `json:"description,omitempty"`
	// List of recommended ASINs under current theme.
	RecommendedAsins []string `json:"recommendedAsins,omitempty"`
	// A theme name representing the context around the recommended list of ASINs.
	Theme *string `json:"theme,omitempty"`
}

// NewThemeRecommendation instantiates a new ThemeRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThemeRecommendation() *ThemeRecommendation {
	this := ThemeRecommendation{}
	return &this
}

// NewThemeRecommendationWithDefaults instantiates a new ThemeRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThemeRecommendationWithDefaults() *ThemeRecommendation {
	this := ThemeRecommendation{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThemeRecommendation) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemeRecommendation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThemeRecommendation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThemeRecommendation) SetDescription(v string) {
	o.Description = &v
}

// GetRecommendedAsins returns the RecommendedAsins field value if set, zero value otherwise.
func (o *ThemeRecommendation) GetRecommendedAsins() []string {
	if o == nil || IsNil(o.RecommendedAsins) {
		var ret []string
		return ret
	}
	return o.RecommendedAsins
}

// GetRecommendedAsinsOk returns a tuple with the RecommendedAsins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemeRecommendation) GetRecommendedAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.RecommendedAsins) {
		return nil, false
	}
	return o.RecommendedAsins, true
}

// HasRecommendedAsins returns a boolean if a field has been set.
func (o *ThemeRecommendation) HasRecommendedAsins() bool {
	if o != nil && !IsNil(o.RecommendedAsins) {
		return true
	}

	return false
}

// SetRecommendedAsins gets a reference to the given []string and assigns it to the RecommendedAsins field.
func (o *ThemeRecommendation) SetRecommendedAsins(v []string) {
	o.RecommendedAsins = v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *ThemeRecommendation) GetTheme() string {
	if o == nil || IsNil(o.Theme) {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemeRecommendation) GetThemeOk() (*string, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *ThemeRecommendation) HasTheme() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *ThemeRecommendation) SetTheme(v string) {
	o.Theme = &v
}

func (o ThemeRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.RecommendedAsins) {
		toSerialize["recommendedAsins"] = o.RecommendedAsins
	}
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	return toSerialize, nil
}

type NullableThemeRecommendation struct {
	value *ThemeRecommendation
	isSet bool
}

func (v NullableThemeRecommendation) Get() *ThemeRecommendation {
	return v.value
}

func (v *NullableThemeRecommendation) Set(val *ThemeRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableThemeRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableThemeRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThemeRecommendation(val *ThemeRecommendation) *NullableThemeRecommendation {
	return &NullableThemeRecommendation{value: val, isSet: true}
}

func (v NullableThemeRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableThemeRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
