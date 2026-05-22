package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeImageRecommendationRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeImageRecommendationRequestContent{}

// CreativeImageRecommendationRequestContent struct for CreativeImageRecommendationRequestContent
type CreativeImageRecommendationRequestContent struct {
	// ----------------------------------------------- List types ----------------------------------------------- A list of ASINs
	Asins        []string      `json:"asins"`
	AssetSubType *AssetSubType `json:"assetSubType,omitempty"`
	// Maximum number of recommendations that API should return. Response will [0, recommendations] recommendations (recommendations are not guaranteed).
	MaxNumRecommendations *float32 `json:"maxNumRecommendations,omitempty"`
	// Filter assets by program types. For example, if only [A_PLUS] assets are requested then only assets that were used as A+ content will be recommended. If no program type is provided, recommend assets from all programs
	AssetPrograms []ProgramType `json:"assetPrograms,omitempty"`
	// (Optional) locale of creative headline and ASIN titles. If locale is not provided, default locale of marketplace is used. Currently, only en_US and en_CA are supported.
	Locale *string `json:"locale,omitempty"`
	// The headline text. Maximum length of the string is 50 characters for all marketplaces other than Japan, which has a maximum length of 35 characters. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#headlines) for headline requirements.
	Headline *string `json:"headline,omitempty"`
}

type _CreativeImageRecommendationRequestContent CreativeImageRecommendationRequestContent

// NewCreativeImageRecommendationRequestContent instantiates a new CreativeImageRecommendationRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeImageRecommendationRequestContent(asins []string) *CreativeImageRecommendationRequestContent {
	this := CreativeImageRecommendationRequestContent{}
	this.Asins = asins
	return &this
}

// NewCreativeImageRecommendationRequestContentWithDefaults instantiates a new CreativeImageRecommendationRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeImageRecommendationRequestContentWithDefaults() *CreativeImageRecommendationRequestContent {
	this := CreativeImageRecommendationRequestContent{}
	return &this
}

// GetAsins returns the Asins field value
func (o *CreativeImageRecommendationRequestContent) GetAsins() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value
// and a boolean to check if the value has been set.
func (o *CreativeImageRecommendationRequestContent) GetAsinsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Asins, true
}

// SetAsins sets field value
func (o *CreativeImageRecommendationRequestContent) SetAsins(v []string) {
	o.Asins = v
}

// GetAssetSubType returns the AssetSubType field value if set, zero value otherwise.
func (o *CreativeImageRecommendationRequestContent) GetAssetSubType() AssetSubType {
	if o == nil || IsNil(o.AssetSubType) {
		var ret AssetSubType
		return ret
	}
	return *o.AssetSubType
}

// GetAssetSubTypeOk returns a tuple with the AssetSubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeImageRecommendationRequestContent) GetAssetSubTypeOk() (*AssetSubType, bool) {
	if o == nil || IsNil(o.AssetSubType) {
		return nil, false
	}
	return o.AssetSubType, true
}

// HasAssetSubType returns a boolean if a field has been set.
func (o *CreativeImageRecommendationRequestContent) HasAssetSubType() bool {
	if o != nil && !IsNil(o.AssetSubType) {
		return true
	}

	return false
}

// SetAssetSubType gets a reference to the given AssetSubType and assigns it to the AssetSubType field.
func (o *CreativeImageRecommendationRequestContent) SetAssetSubType(v AssetSubType) {
	o.AssetSubType = &v
}

// GetMaxNumRecommendations returns the MaxNumRecommendations field value if set, zero value otherwise.
func (o *CreativeImageRecommendationRequestContent) GetMaxNumRecommendations() float32 {
	if o == nil || IsNil(o.MaxNumRecommendations) {
		var ret float32
		return ret
	}
	return *o.MaxNumRecommendations
}

// GetMaxNumRecommendationsOk returns a tuple with the MaxNumRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeImageRecommendationRequestContent) GetMaxNumRecommendationsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxNumRecommendations) {
		return nil, false
	}
	return o.MaxNumRecommendations, true
}

// HasMaxNumRecommendations returns a boolean if a field has been set.
func (o *CreativeImageRecommendationRequestContent) HasMaxNumRecommendations() bool {
	if o != nil && !IsNil(o.MaxNumRecommendations) {
		return true
	}

	return false
}

// SetMaxNumRecommendations gets a reference to the given float32 and assigns it to the MaxNumRecommendations field.
func (o *CreativeImageRecommendationRequestContent) SetMaxNumRecommendations(v float32) {
	o.MaxNumRecommendations = &v
}

// GetAssetPrograms returns the AssetPrograms field value if set, zero value otherwise.
func (o *CreativeImageRecommendationRequestContent) GetAssetPrograms() []ProgramType {
	if o == nil || IsNil(o.AssetPrograms) {
		var ret []ProgramType
		return ret
	}
	return o.AssetPrograms
}

// GetAssetProgramsOk returns a tuple with the AssetPrograms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeImageRecommendationRequestContent) GetAssetProgramsOk() ([]ProgramType, bool) {
	if o == nil || IsNil(o.AssetPrograms) {
		return nil, false
	}
	return o.AssetPrograms, true
}

// HasAssetPrograms returns a boolean if a field has been set.
func (o *CreativeImageRecommendationRequestContent) HasAssetPrograms() bool {
	if o != nil && !IsNil(o.AssetPrograms) {
		return true
	}

	return false
}

// SetAssetPrograms gets a reference to the given []ProgramType and assigns it to the AssetPrograms field.
func (o *CreativeImageRecommendationRequestContent) SetAssetPrograms(v []ProgramType) {
	o.AssetPrograms = v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *CreativeImageRecommendationRequestContent) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeImageRecommendationRequestContent) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *CreativeImageRecommendationRequestContent) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *CreativeImageRecommendationRequestContent) SetLocale(v string) {
	o.Locale = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *CreativeImageRecommendationRequestContent) GetHeadline() string {
	if o == nil || IsNil(o.Headline) {
		var ret string
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeImageRecommendationRequestContent) GetHeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *CreativeImageRecommendationRequestContent) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given string and assigns it to the Headline field.
func (o *CreativeImageRecommendationRequestContent) SetHeadline(v string) {
	o.Headline = &v
}

func (o CreativeImageRecommendationRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asins"] = o.Asins
	if !IsNil(o.AssetSubType) {
		toSerialize["assetSubType"] = o.AssetSubType
	}
	if !IsNil(o.MaxNumRecommendations) {
		toSerialize["maxNumRecommendations"] = o.MaxNumRecommendations
	}
	if !IsNil(o.AssetPrograms) {
		toSerialize["assetPrograms"] = o.AssetPrograms
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	return toSerialize, nil
}

type NullableCreativeImageRecommendationRequestContent struct {
	value *CreativeImageRecommendationRequestContent
	isSet bool
}

func (v NullableCreativeImageRecommendationRequestContent) Get() *CreativeImageRecommendationRequestContent {
	return v.value
}

func (v *NullableCreativeImageRecommendationRequestContent) Set(val *CreativeImageRecommendationRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeImageRecommendationRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeImageRecommendationRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeImageRecommendationRequestContent(val *CreativeImageRecommendationRequestContent) *NullableCreativeImageRecommendationRequestContent {
	return &NullableCreativeImageRecommendationRequestContent{value: val, isSet: true}
}

func (v NullableCreativeImageRecommendationRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeImageRecommendationRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
