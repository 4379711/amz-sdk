package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the KeywordTargetRankRecommendationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeywordTargetRankRecommendationRequest{}

// KeywordTargetRankRecommendationRequest struct for KeywordTargetRankRecommendationRequest
type KeywordTargetRankRecommendationRequest struct {
	// A list of targets that need to be ranked
	Targets []KeywordTarget `json:"targets,omitempty"`
	// The max size of recommended target. Set it to 0 if you only want to rank user-defined keywords.
	MaxRecommendations *float32 `json:"maxRecommendations,omitempty"`
	// The ranking metric value. Supported values: CLICKS, CONVERSIONS, DEFAULT. DEFAULT will be applied if no value passed in.
	SortDimension *string `json:"sortDimension,omitempty"`
	// Translations are for readability and do not affect the targeting of ads. Supported marketplace to locale mappings can be found at the <a href='https://advertising.amazon.com/API/docs/en-us/localization/#/Keyword%20Localization'>POST /keywords/localize</a> endpoint. Note: Translations will be null if locale is unsupported.
	Locale *string `json:"locale,omitempty"`
}

// NewKeywordTargetRankRecommendationRequest instantiates a new KeywordTargetRankRecommendationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeywordTargetRankRecommendationRequest() *KeywordTargetRankRecommendationRequest {
	this := KeywordTargetRankRecommendationRequest{}
	return &this
}

// NewKeywordTargetRankRecommendationRequestWithDefaults instantiates a new KeywordTargetRankRecommendationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeywordTargetRankRecommendationRequestWithDefaults() *KeywordTargetRankRecommendationRequest {
	this := KeywordTargetRankRecommendationRequest{}
	return &this
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *KeywordTargetRankRecommendationRequest) GetTargets() []KeywordTarget {
	if o == nil || IsNil(o.Targets) {
		var ret []KeywordTarget
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTargetRankRecommendationRequest) GetTargetsOk() ([]KeywordTarget, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *KeywordTargetRankRecommendationRequest) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []KeywordTarget and assigns it to the Targets field.
func (o *KeywordTargetRankRecommendationRequest) SetTargets(v []KeywordTarget) {
	o.Targets = v
}

// GetMaxRecommendations returns the MaxRecommendations field value if set, zero value otherwise.
func (o *KeywordTargetRankRecommendationRequest) GetMaxRecommendations() float32 {
	if o == nil || IsNil(o.MaxRecommendations) {
		var ret float32
		return ret
	}
	return *o.MaxRecommendations
}

// GetMaxRecommendationsOk returns a tuple with the MaxRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTargetRankRecommendationRequest) GetMaxRecommendationsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxRecommendations) {
		return nil, false
	}
	return o.MaxRecommendations, true
}

// HasMaxRecommendations returns a boolean if a field has been set.
func (o *KeywordTargetRankRecommendationRequest) HasMaxRecommendations() bool {
	if o != nil && !IsNil(o.MaxRecommendations) {
		return true
	}

	return false
}

// SetMaxRecommendations gets a reference to the given float32 and assigns it to the MaxRecommendations field.
func (o *KeywordTargetRankRecommendationRequest) SetMaxRecommendations(v float32) {
	o.MaxRecommendations = &v
}

// GetSortDimension returns the SortDimension field value if set, zero value otherwise.
func (o *KeywordTargetRankRecommendationRequest) GetSortDimension() string {
	if o == nil || IsNil(o.SortDimension) {
		var ret string
		return ret
	}
	return *o.SortDimension
}

// GetSortDimensionOk returns a tuple with the SortDimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTargetRankRecommendationRequest) GetSortDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.SortDimension) {
		return nil, false
	}
	return o.SortDimension, true
}

// HasSortDimension returns a boolean if a field has been set.
func (o *KeywordTargetRankRecommendationRequest) HasSortDimension() bool {
	if o != nil && !IsNil(o.SortDimension) {
		return true
	}

	return false
}

// SetSortDimension gets a reference to the given string and assigns it to the SortDimension field.
func (o *KeywordTargetRankRecommendationRequest) SetSortDimension(v string) {
	o.SortDimension = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *KeywordTargetRankRecommendationRequest) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTargetRankRecommendationRequest) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *KeywordTargetRankRecommendationRequest) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *KeywordTargetRankRecommendationRequest) SetLocale(v string) {
	o.Locale = &v
}

func (o KeywordTargetRankRecommendationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
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

type NullableKeywordTargetRankRecommendationRequest struct {
	value *KeywordTargetRankRecommendationRequest
	isSet bool
}

func (v NullableKeywordTargetRankRecommendationRequest) Get() *KeywordTargetRankRecommendationRequest {
	return v.value
}

func (v *NullableKeywordTargetRankRecommendationRequest) Set(val *KeywordTargetRankRecommendationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeywordTargetRankRecommendationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeywordTargetRankRecommendationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeywordTargetRankRecommendationRequest(val *KeywordTargetRankRecommendationRequest) *NullableKeywordTargetRankRecommendationRequest {
	return &NullableKeywordTargetRankRecommendationRequest{value: val, isSet: true}
}

func (v NullableKeywordTargetRankRecommendationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableKeywordTargetRankRecommendationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
