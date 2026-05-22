package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the RecommendationOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecommendationOptions{}

// RecommendationOptions struct for RecommendationOptions
type RecommendationOptions struct {
	// The max size of recommended target. Set it to 0 if you only want to rank user-defined keywords.
	MaxRecommendations *float32 `json:"maxRecommendations,omitempty"`
	// The ranking metric value. Supported values: CLICKS, CONVERSIONS, DEFAULT. DEFAULT will be applied if no value passed in.
	SortDimension *string `json:"sortDimension,omitempty"`
	// Translations are for readability and do not affect the targeting of ads. Supported marketplace to locale mappings can be found at the <a href='https://advertising.amazon.com/API/docs/en-us/localization/#/Keyword%20Localization'>POST /keywords/localize</a> endpoint. Note: Translations will be null if locale is unsupported.
	Locale *string `json:"locale,omitempty"`
}

// NewRecommendationOptions instantiates a new RecommendationOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationOptions() *RecommendationOptions {
	this := RecommendationOptions{}
	return &this
}

// NewRecommendationOptionsWithDefaults instantiates a new RecommendationOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationOptionsWithDefaults() *RecommendationOptions {
	this := RecommendationOptions{}
	return &this
}

// GetMaxRecommendations returns the MaxRecommendations field value if set, zero value otherwise.
func (o *RecommendationOptions) GetMaxRecommendations() float32 {
	if o == nil || IsNil(o.MaxRecommendations) {
		var ret float32
		return ret
	}
	return *o.MaxRecommendations
}

// GetMaxRecommendationsOk returns a tuple with the MaxRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationOptions) GetMaxRecommendationsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxRecommendations) {
		return nil, false
	}
	return o.MaxRecommendations, true
}

// HasMaxRecommendations returns a boolean if a field has been set.
func (o *RecommendationOptions) HasMaxRecommendations() bool {
	if o != nil && !IsNil(o.MaxRecommendations) {
		return true
	}

	return false
}

// SetMaxRecommendations gets a reference to the given float32 and assigns it to the MaxRecommendations field.
func (o *RecommendationOptions) SetMaxRecommendations(v float32) {
	o.MaxRecommendations = &v
}

// GetSortDimension returns the SortDimension field value if set, zero value otherwise.
func (o *RecommendationOptions) GetSortDimension() string {
	if o == nil || IsNil(o.SortDimension) {
		var ret string
		return ret
	}
	return *o.SortDimension
}

// GetSortDimensionOk returns a tuple with the SortDimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationOptions) GetSortDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.SortDimension) {
		return nil, false
	}
	return o.SortDimension, true
}

// HasSortDimension returns a boolean if a field has been set.
func (o *RecommendationOptions) HasSortDimension() bool {
	if o != nil && !IsNil(o.SortDimension) {
		return true
	}

	return false
}

// SetSortDimension gets a reference to the given string and assigns it to the SortDimension field.
func (o *RecommendationOptions) SetSortDimension(v string) {
	o.SortDimension = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *RecommendationOptions) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationOptions) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *RecommendationOptions) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *RecommendationOptions) SetLocale(v string) {
	o.Locale = &v
}

func (o RecommendationOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxRecommendations) {
		toSerialize["maxRecommendations"] = o.MaxRecommendations
	}
	if !IsNil(o.SortDimension) {
		toSerialize["sortDimension"] = o.SortDimension
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	return toSerialize, nil
}

type NullableRecommendationOptions struct {
	value *RecommendationOptions
	isSet bool
}

func (v NullableRecommendationOptions) Get() *RecommendationOptions {
	return v.value
}

func (v *NullableRecommendationOptions) Set(val *RecommendationOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationOptions(val *RecommendationOptions) *NullableRecommendationOptions {
	return &NullableRecommendationOptions{value: val, isSet: true}
}

func (v NullableRecommendationOptions) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRecommendationOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
