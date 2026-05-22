package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBKeywordRecommendationThemeKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBKeywordRecommendationThemeKeyword{}

// SBKeywordRecommendationThemeKeyword struct for SBKeywordRecommendationThemeKeyword
type SBKeywordRecommendationThemeKeyword struct {
	// Unique ID for each recommendation.
	RecommendationId *string `json:"recommendationId,omitempty"`
	// Recommended keyword value.
	Value *string `json:"value,omitempty"`
}

// NewSBKeywordRecommendationThemeKeyword instantiates a new SBKeywordRecommendationThemeKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBKeywordRecommendationThemeKeyword() *SBKeywordRecommendationThemeKeyword {
	this := SBKeywordRecommendationThemeKeyword{}
	return &this
}

// NewSBKeywordRecommendationThemeKeywordWithDefaults instantiates a new SBKeywordRecommendationThemeKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBKeywordRecommendationThemeKeywordWithDefaults() *SBKeywordRecommendationThemeKeyword {
	this := SBKeywordRecommendationThemeKeyword{}
	return &this
}

// GetRecommendationId returns the RecommendationId field value if set, zero value otherwise.
func (o *SBKeywordRecommendationThemeKeyword) GetRecommendationId() string {
	if o == nil || IsNil(o.RecommendationId) {
		var ret string
		return ret
	}
	return *o.RecommendationId
}

// GetRecommendationIdOk returns a tuple with the RecommendationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBKeywordRecommendationThemeKeyword) GetRecommendationIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecommendationId) {
		return nil, false
	}
	return o.RecommendationId, true
}

// HasRecommendationId returns a boolean if a field has been set.
func (o *SBKeywordRecommendationThemeKeyword) HasRecommendationId() bool {
	if o != nil && !IsNil(o.RecommendationId) {
		return true
	}

	return false
}

// SetRecommendationId gets a reference to the given string and assigns it to the RecommendationId field.
func (o *SBKeywordRecommendationThemeKeyword) SetRecommendationId(v string) {
	o.RecommendationId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SBKeywordRecommendationThemeKeyword) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBKeywordRecommendationThemeKeyword) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SBKeywordRecommendationThemeKeyword) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SBKeywordRecommendationThemeKeyword) SetValue(v string) {
	o.Value = &v
}

func (o SBKeywordRecommendationThemeKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecommendationId) {
		toSerialize["recommendationId"] = o.RecommendationId
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSBKeywordRecommendationThemeKeyword struct {
	value *SBKeywordRecommendationThemeKeyword
	isSet bool
}

func (v NullableSBKeywordRecommendationThemeKeyword) Get() *SBKeywordRecommendationThemeKeyword {
	return v.value
}

func (v *NullableSBKeywordRecommendationThemeKeyword) Set(val *SBKeywordRecommendationThemeKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSBKeywordRecommendationThemeKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSBKeywordRecommendationThemeKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBKeywordRecommendationThemeKeyword(val *SBKeywordRecommendationThemeKeyword) *NullableSBKeywordRecommendationThemeKeyword {
	return &NullableSBKeywordRecommendationThemeKeyword{value: val, isSet: true}
}

func (v NullableSBKeywordRecommendationThemeKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBKeywordRecommendationThemeKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
