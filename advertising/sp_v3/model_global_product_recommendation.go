package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalProductRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalProductRecommendation{}

// GlobalProductRecommendation Recommended asin and related information.
type GlobalProductRecommendation struct {
	// List of themes associated with this recommended ASIN.
	Themes []string `json:"themes,omitempty"`
	// Recommended ASIN
	RecommendedAsin *string `json:"recommendedAsin,omitempty"`
}

// NewGlobalProductRecommendation instantiates a new GlobalProductRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalProductRecommendation() *GlobalProductRecommendation {
	this := GlobalProductRecommendation{}
	return &this
}

// NewGlobalProductRecommendationWithDefaults instantiates a new GlobalProductRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalProductRecommendationWithDefaults() *GlobalProductRecommendation {
	this := GlobalProductRecommendation{}
	return &this
}

// GetThemes returns the Themes field value if set, zero value otherwise.
func (o *GlobalProductRecommendation) GetThemes() []string {
	if o == nil || IsNil(o.Themes) {
		var ret []string
		return ret
	}
	return o.Themes
}

// GetThemesOk returns a tuple with the Themes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalProductRecommendation) GetThemesOk() ([]string, bool) {
	if o == nil || IsNil(o.Themes) {
		return nil, false
	}
	return o.Themes, true
}

// HasThemes returns a boolean if a field has been set.
func (o *GlobalProductRecommendation) HasThemes() bool {
	if o != nil && !IsNil(o.Themes) {
		return true
	}

	return false
}

// SetThemes gets a reference to the given []string and assigns it to the Themes field.
func (o *GlobalProductRecommendation) SetThemes(v []string) {
	o.Themes = v
}

// GetRecommendedAsin returns the RecommendedAsin field value if set, zero value otherwise.
func (o *GlobalProductRecommendation) GetRecommendedAsin() string {
	if o == nil || IsNil(o.RecommendedAsin) {
		var ret string
		return ret
	}
	return *o.RecommendedAsin
}

// GetRecommendedAsinOk returns a tuple with the RecommendedAsin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalProductRecommendation) GetRecommendedAsinOk() (*string, bool) {
	if o == nil || IsNil(o.RecommendedAsin) {
		return nil, false
	}
	return o.RecommendedAsin, true
}

// HasRecommendedAsin returns a boolean if a field has been set.
func (o *GlobalProductRecommendation) HasRecommendedAsin() bool {
	if o != nil && !IsNil(o.RecommendedAsin) {
		return true
	}

	return false
}

// SetRecommendedAsin gets a reference to the given string and assigns it to the RecommendedAsin field.
func (o *GlobalProductRecommendation) SetRecommendedAsin(v string) {
	o.RecommendedAsin = &v
}

func (o GlobalProductRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Themes) {
		toSerialize["themes"] = o.Themes
	}
	if !IsNil(o.RecommendedAsin) {
		toSerialize["recommendedAsin"] = o.RecommendedAsin
	}
	return toSerialize, nil
}

type NullableGlobalProductRecommendation struct {
	value *GlobalProductRecommendation
	isSet bool
}

func (v NullableGlobalProductRecommendation) Get() *GlobalProductRecommendation {
	return v.value
}

func (v *NullableGlobalProductRecommendation) Set(val *GlobalProductRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalProductRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalProductRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalProductRecommendation(val *GlobalProductRecommendation) *NullableGlobalProductRecommendation {
	return &NullableGlobalProductRecommendation{value: val, isSet: true}
}

func (v NullableGlobalProductRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalProductRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
