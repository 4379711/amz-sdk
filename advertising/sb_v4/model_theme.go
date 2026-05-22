package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the Theme type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Theme{}

// Theme Structure for theme details
type Theme struct {
	ThemeForDisplay string `json:"themeForDisplay"`
	ThemeId         string `json:"themeId"`
}

type _Theme Theme

// NewTheme instantiates a new Theme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTheme(themeForDisplay string, themeId string) *Theme {
	this := Theme{}
	this.ThemeForDisplay = themeForDisplay
	this.ThemeId = themeId
	return &this
}

// NewThemeWithDefaults instantiates a new Theme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThemeWithDefaults() *Theme {
	this := Theme{}
	return &this
}

// GetThemeForDisplay returns the ThemeForDisplay field value
func (o *Theme) GetThemeForDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThemeForDisplay
}

// GetThemeForDisplayOk returns a tuple with the ThemeForDisplay field value
// and a boolean to check if the value has been set.
func (o *Theme) GetThemeForDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThemeForDisplay, true
}

// SetThemeForDisplay sets field value
func (o *Theme) SetThemeForDisplay(v string) {
	o.ThemeForDisplay = v
}

// GetThemeId returns the ThemeId field value
func (o *Theme) GetThemeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThemeId
}

// GetThemeIdOk returns a tuple with the ThemeId field value
// and a boolean to check if the value has been set.
func (o *Theme) GetThemeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThemeId, true
}

// SetThemeId sets field value
func (o *Theme) SetThemeId(v string) {
	o.ThemeId = v
}

func (o Theme) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["themeForDisplay"] = o.ThemeForDisplay
	toSerialize["themeId"] = o.ThemeId
	return toSerialize, nil
}

type NullableTheme struct {
	value *Theme
	isSet bool
}

func (v NullableTheme) Get() *Theme {
	return v.value
}

func (v *NullableTheme) Set(val *Theme) {
	v.value = val
	v.isSet = true
}

func (v NullableTheme) IsSet() bool {
	return v.isSet
}

func (v *NullableTheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTheme(val *Theme) *NullableTheme {
	return &NullableTheme{value: val, isSet: true}
}

func (v NullableTheme) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
