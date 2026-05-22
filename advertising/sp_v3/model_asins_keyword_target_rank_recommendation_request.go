package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the AsinsKeywordTargetRankRecommendationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsinsKeywordTargetRankRecommendationRequest{}

// AsinsKeywordTargetRankRecommendationRequest This request type is used to retrieve recommended keyword targets for ASINs. Set the recommendationType to KEYWORDS_FOR_ASINS to use this request type.
type AsinsKeywordTargetRankRecommendationRequest struct {
	// An array list of Asins
	Asins []string `json:"asins"`
	// The recommendationType to retrieve recommended keyword targets for a list of ASINs.
	RecommendationType string `json:"recommendationType"`
	// A list of targets that need to be ranked
	Targets []KeywordTarget `json:"targets,omitempty"`
	// The max size of recommended target. Set it to 0 if you only want to rank user-defined keywords.
	MaxRecommendations *float32 `json:"maxRecommendations,omitempty"`
	// The ranking metric value. Supported values: CLICKS, CONVERSIONS, DEFAULT. DEFAULT will be applied if no value passed in.
	SortDimension *string `json:"sortDimension,omitempty"`
	// Translations are for readability and do not affect the targeting of ads. Supported marketplace to locale mappings can be found at the <a href='https://advertising.amazon.com/API/docs/en-us/localization/#/Keyword%20Localization'>POST /keywords/localize</a> endpoint. Note: Translations will be null if locale is unsupported.
	Locale *string `json:"locale,omitempty"`
}

type _AsinsKeywordTargetRankRecommendationRequest AsinsKeywordTargetRankRecommendationRequest

// NewAsinsKeywordTargetRankRecommendationRequest instantiates a new AsinsKeywordTargetRankRecommendationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsinsKeywordTargetRankRecommendationRequest(asins []string, recommendationType string) *AsinsKeywordTargetRankRecommendationRequest {
	this := AsinsKeywordTargetRankRecommendationRequest{}
	return &this
}

// NewAsinsKeywordTargetRankRecommendationRequestWithDefaults instantiates a new AsinsKeywordTargetRankRecommendationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsinsKeywordTargetRankRecommendationRequestWithDefaults() *AsinsKeywordTargetRankRecommendationRequest {
	this := AsinsKeywordTargetRankRecommendationRequest{}
	return &this
}

// GetAsins returns the Asins field value
func (o *AsinsKeywordTargetRankRecommendationRequest) GetAsins() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value
// and a boolean to check if the value has been set.
func (o *AsinsKeywordTargetRankRecommendationRequest) GetAsinsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Asins, true
}

// SetAsins sets field value
func (o *AsinsKeywordTargetRankRecommendationRequest) SetAsins(v []string) {
	o.Asins = v
}

// GetRecommendationType returns the RecommendationType field value
func (o *AsinsKeywordTargetRankRecommendationRequest) GetRecommendationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecommendationType
}

// GetRecommendationTypeOk returns a tuple with the RecommendationType field value
// and a boolean to check if the value has been set.
func (o *AsinsKeywordTargetRankRecommendationRequest) GetRecommendationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendationType, true
}

// SetRecommendationType sets field value
func (o *AsinsKeywordTargetRankRecommendationRequest) SetRecommendationType(v string) {
	o.RecommendationType = v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *AsinsKeywordTargetRankRecommendationRequest) GetTargets() []KeywordTarget {
	if o == nil || IsNil(o.Targets) {
		var ret []KeywordTarget
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinsKeywordTargetRankRecommendationRequest) GetTargetsOk() ([]KeywordTarget, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *AsinsKeywordTargetRankRecommendationRequest) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []KeywordTarget and assigns it to the Targets field.
func (o *AsinsKeywordTargetRankRecommendationRequest) SetTargets(v []KeywordTarget) {
	o.Targets = v
}

// GetMaxRecommendations returns the MaxRecommendations field value if set, zero value otherwise.
func (o *AsinsKeywordTargetRankRecommendationRequest) GetMaxRecommendations() float32 {
	if o == nil || IsNil(o.MaxRecommendations) {
		var ret float32
		return ret
	}
	return *o.MaxRecommendations
}

// GetMaxRecommendationsOk returns a tuple with the MaxRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinsKeywordTargetRankRecommendationRequest) GetMaxRecommendationsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxRecommendations) {
		return nil, false
	}
	return o.MaxRecommendations, true
}

// HasMaxRecommendations returns a boolean if a field has been set.
func (o *AsinsKeywordTargetRankRecommendationRequest) HasMaxRecommendations() bool {
	if o != nil && !IsNil(o.MaxRecommendations) {
		return true
	}

	return false
}

// SetMaxRecommendations gets a reference to the given float32 and assigns it to the MaxRecommendations field.
func (o *AsinsKeywordTargetRankRecommendationRequest) SetMaxRecommendations(v float32) {
	o.MaxRecommendations = &v
}

// GetSortDimension returns the SortDimension field value if set, zero value otherwise.
func (o *AsinsKeywordTargetRankRecommendationRequest) GetSortDimension() string {
	if o == nil || IsNil(o.SortDimension) {
		var ret string
		return ret
	}
	return *o.SortDimension
}

// GetSortDimensionOk returns a tuple with the SortDimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinsKeywordTargetRankRecommendationRequest) GetSortDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.SortDimension) {
		return nil, false
	}
	return o.SortDimension, true
}

// HasSortDimension returns a boolean if a field has been set.
func (o *AsinsKeywordTargetRankRecommendationRequest) HasSortDimension() bool {
	if o != nil && !IsNil(o.SortDimension) {
		return true
	}

	return false
}

// SetSortDimension gets a reference to the given string and assigns it to the SortDimension field.
func (o *AsinsKeywordTargetRankRecommendationRequest) SetSortDimension(v string) {
	o.SortDimension = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *AsinsKeywordTargetRankRecommendationRequest) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinsKeywordTargetRankRecommendationRequest) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *AsinsKeywordTargetRankRecommendationRequest) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *AsinsKeywordTargetRankRecommendationRequest) SetLocale(v string) {
	o.Locale = &v
}

func (o AsinsKeywordTargetRankRecommendationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asins"] = o.Asins
	toSerialize["recommendationType"] = o.RecommendationType
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

type NullableAsinsKeywordTargetRankRecommendationRequest struct {
	value *AsinsKeywordTargetRankRecommendationRequest
	isSet bool
}

func (v NullableAsinsKeywordTargetRankRecommendationRequest) Get() *AsinsKeywordTargetRankRecommendationRequest {
	return v.value
}

func (v *NullableAsinsKeywordTargetRankRecommendationRequest) Set(val *AsinsKeywordTargetRankRecommendationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAsinsKeywordTargetRankRecommendationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAsinsKeywordTargetRankRecommendationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsinsKeywordTargetRankRecommendationRequest(val *AsinsKeywordTargetRankRecommendationRequest) *NullableAsinsKeywordTargetRankRecommendationRequest {
	return &NullableAsinsKeywordTargetRankRecommendationRequest{value: val, isSet: true}
}

func (v NullableAsinsKeywordTargetRankRecommendationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAsinsKeywordTargetRankRecommendationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
