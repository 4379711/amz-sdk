package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingTheme type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingTheme{}

// SBForecastingTheme The theme.
type SBForecastingTheme struct {
	// The theme target type. Valid value: KEYWORDS_RELATED_TO_YOUR_BRAND, KEYWORDS_RELATED_TO_YOUR_LANDING_PAGES.   KEYWORDS_RELATED_TO_YOUR_BRAND - keywords related to brands.   KEYWORDS_RELATED_TO_YOUR_LANDING_PAGES - keywords related to your landing pages.
	ThemeType *string `json:"themeType,omitempty"`
	// The associated bid. Note that this value must be less than the budget associated with the Advertiser account.
	Bid *float32 `json:"bid,omitempty"`
}

// NewSBForecastingTheme instantiates a new SBForecastingTheme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingTheme() *SBForecastingTheme {
	this := SBForecastingTheme{}
	return &this
}

// NewSBForecastingThemeWithDefaults instantiates a new SBForecastingTheme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingThemeWithDefaults() *SBForecastingTheme {
	this := SBForecastingTheme{}
	return &this
}

// GetThemeType returns the ThemeType field value if set, zero value otherwise.
func (o *SBForecastingTheme) GetThemeType() string {
	if o == nil || IsNil(o.ThemeType) {
		var ret string
		return ret
	}
	return *o.ThemeType
}

// GetThemeTypeOk returns a tuple with the ThemeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingTheme) GetThemeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ThemeType) {
		return nil, false
	}
	return o.ThemeType, true
}

// HasThemeType returns a boolean if a field has been set.
func (o *SBForecastingTheme) HasThemeType() bool {
	if o != nil && !IsNil(o.ThemeType) {
		return true
	}

	return false
}

// SetThemeType gets a reference to the given string and assigns it to the ThemeType field.
func (o *SBForecastingTheme) SetThemeType(v string) {
	o.ThemeType = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *SBForecastingTheme) GetBid() float32 {
	if o == nil || IsNil(o.Bid) {
		var ret float32
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingTheme) GetBidOk() (*float32, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *SBForecastingTheme) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given float32 and assigns it to the Bid field.
func (o *SBForecastingTheme) SetBid(v float32) {
	o.Bid = &v
}

func (o SBForecastingTheme) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ThemeType) {
		toSerialize["themeType"] = o.ThemeType
	}
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	return toSerialize, nil
}

type NullableSBForecastingTheme struct {
	value *SBForecastingTheme
	isSet bool
}

func (v NullableSBForecastingTheme) Get() *SBForecastingTheme {
	return v.value
}

func (v *NullableSBForecastingTheme) Set(val *SBForecastingTheme) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingTheme) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingTheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingTheme(val *SBForecastingTheme) *NullableSBForecastingTheme {
	return &NullableSBForecastingTheme{value: val, isSet: true}
}

func (v NullableSBForecastingTheme) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingTheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
