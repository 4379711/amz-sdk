package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalRankedKeywordTargetsForAdGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalRankedKeywordTargetsForAdGroupRequest{}

// GlobalRankedKeywordTargetsForAdGroupRequest This request type is used to retrieve recommended keyword targets for an existing ad group. Set the recommendationType to KEYWORDS_FOR_ADGROUP to use this request type.
type GlobalRankedKeywordTargetsForAdGroupRequest struct {
	// A list of countries with targets that need to be ranked. The 2-letter country code is the key, and an object with targets is the value.
	CountryCodes *map[string]CountryWithTargets `json:"countryCodes,omitempty"`
	// The identifier of the campaign
	CampaignId string `json:"campaignId"`
	// The recommendationType to retrieve recommended keyword targets for an existing ad group.
	RecommendationType string `json:"recommendationType"`
	// Set this parameter to false if you do not want to retrieve bid suggestions for your keyword targets. Defaults to true.
	BidsEnabled *bool `json:"bidsEnabled,omitempty"`
	// The identifier of the ad group
	AdGroupId string `json:"adGroupId"`
	// The max size of recommended target. Set it to 0 if you only want to rank user-defined keywords.
	MaxRecommendations *float32 `json:"maxRecommendations,omitempty"`
	// The ranking metric value. Supported values: CLICKS, CONVERSIONS, DEFAULT. DEFAULT will be applied if no value passed in.
	SortDimension *string `json:"sortDimension,omitempty"`
	// Translations are for readability and do not affect the targeting of ads. Supported marketplace to locale mappings can be found at the <a href='https://advertising.amazon.com/API/docs/en-us/localization/#/Keyword%20Localization'>POST /keywords/localize</a> endpoint. Note: Translations will be null if locale is unsupported.
	Locale *string `json:"locale,omitempty"`
}

type _GlobalRankedKeywordTargetsForAdGroupRequest GlobalRankedKeywordTargetsForAdGroupRequest

// NewGlobalRankedKeywordTargetsForAdGroupRequest instantiates a new GlobalRankedKeywordTargetsForAdGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalRankedKeywordTargetsForAdGroupRequest(campaignId string, recommendationType string, adGroupId string) *GlobalRankedKeywordTargetsForAdGroupRequest {
	this := GlobalRankedKeywordTargetsForAdGroupRequest{}
	this.CampaignId = campaignId
	this.RecommendationType = recommendationType
	var bidsEnabled bool = true
	this.BidsEnabled = &bidsEnabled
	this.AdGroupId = adGroupId
	return &this
}

// NewGlobalRankedKeywordTargetsForAdGroupRequestWithDefaults instantiates a new GlobalRankedKeywordTargetsForAdGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalRankedKeywordTargetsForAdGroupRequestWithDefaults() *GlobalRankedKeywordTargetsForAdGroupRequest {
	this := GlobalRankedKeywordTargetsForAdGroupRequest{}
	var bidsEnabled bool = true
	this.BidsEnabled = &bidsEnabled
	return &this
}

// GetCountryCodes returns the CountryCodes field value if set, zero value otherwise.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) GetCountryCodes() map[string]CountryWithTargets {
	if o == nil || IsNil(o.CountryCodes) {
		var ret map[string]CountryWithTargets
		return ret
	}
	return *o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) GetCountryCodesOk() (*map[string]CountryWithTargets, bool) {
	if o == nil || IsNil(o.CountryCodes) {
		return nil, false
	}
	return o.CountryCodes, true
}

// HasCountryCodes returns a boolean if a field has been set.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) HasCountryCodes() bool {
	if o != nil && !IsNil(o.CountryCodes) {
		return true
	}

	return false
}

// SetCountryCodes gets a reference to the given map[string]CountryWithTargets and assigns it to the CountryCodes field.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) SetCountryCodes(v map[string]CountryWithTargets) {
	o.CountryCodes = &v
}

// GetCampaignId returns the CampaignId field value
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetRecommendationType returns the RecommendationType field value
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) GetRecommendationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecommendationType
}

// GetRecommendationTypeOk returns a tuple with the RecommendationType field value
// and a boolean to check if the value has been set.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) GetRecommendationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendationType, true
}

// SetRecommendationType sets field value
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) SetRecommendationType(v string) {
	o.RecommendationType = v
}

// GetBidsEnabled returns the BidsEnabled field value if set, zero value otherwise.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) GetBidsEnabled() bool {
	if o == nil || IsNil(o.BidsEnabled) {
		var ret bool
		return ret
	}
	return *o.BidsEnabled
}

// GetBidsEnabledOk returns a tuple with the BidsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) GetBidsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BidsEnabled) {
		return nil, false
	}
	return o.BidsEnabled, true
}

// HasBidsEnabled returns a boolean if a field has been set.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) HasBidsEnabled() bool {
	if o != nil && !IsNil(o.BidsEnabled) {
		return true
	}

	return false
}

// SetBidsEnabled gets a reference to the given bool and assigns it to the BidsEnabled field.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) SetBidsEnabled(v bool) {
	o.BidsEnabled = &v
}

// GetAdGroupId returns the AdGroupId field value
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetMaxRecommendations returns the MaxRecommendations field value if set, zero value otherwise.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) GetMaxRecommendations() float32 {
	if o == nil || IsNil(o.MaxRecommendations) {
		var ret float32
		return ret
	}
	return *o.MaxRecommendations
}

// GetMaxRecommendationsOk returns a tuple with the MaxRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) GetMaxRecommendationsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxRecommendations) {
		return nil, false
	}
	return o.MaxRecommendations, true
}

// HasMaxRecommendations returns a boolean if a field has been set.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) HasMaxRecommendations() bool {
	if o != nil && !IsNil(o.MaxRecommendations) {
		return true
	}

	return false
}

// SetMaxRecommendations gets a reference to the given float32 and assigns it to the MaxRecommendations field.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) SetMaxRecommendations(v float32) {
	o.MaxRecommendations = &v
}

// GetSortDimension returns the SortDimension field value if set, zero value otherwise.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) GetSortDimension() string {
	if o == nil || IsNil(o.SortDimension) {
		var ret string
		return ret
	}
	return *o.SortDimension
}

// GetSortDimensionOk returns a tuple with the SortDimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) GetSortDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.SortDimension) {
		return nil, false
	}
	return o.SortDimension, true
}

// HasSortDimension returns a boolean if a field has been set.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) HasSortDimension() bool {
	if o != nil && !IsNil(o.SortDimension) {
		return true
	}

	return false
}

// SetSortDimension gets a reference to the given string and assigns it to the SortDimension field.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) SetSortDimension(v string) {
	o.SortDimension = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *GlobalRankedKeywordTargetsForAdGroupRequest) SetLocale(v string) {
	o.Locale = &v
}

func (o GlobalRankedKeywordTargetsForAdGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryCodes) {
		toSerialize["countryCodes"] = o.CountryCodes
	}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["recommendationType"] = o.RecommendationType
	if !IsNil(o.BidsEnabled) {
		toSerialize["bidsEnabled"] = o.BidsEnabled
	}
	toSerialize["adGroupId"] = o.AdGroupId
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

type NullableGlobalRankedKeywordTargetsForAdGroupRequest struct {
	value *GlobalRankedKeywordTargetsForAdGroupRequest
	isSet bool
}

func (v NullableGlobalRankedKeywordTargetsForAdGroupRequest) Get() *GlobalRankedKeywordTargetsForAdGroupRequest {
	return v.value
}

func (v *NullableGlobalRankedKeywordTargetsForAdGroupRequest) Set(val *GlobalRankedKeywordTargetsForAdGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalRankedKeywordTargetsForAdGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalRankedKeywordTargetsForAdGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalRankedKeywordTargetsForAdGroupRequest(val *GlobalRankedKeywordTargetsForAdGroupRequest) *NullableGlobalRankedKeywordTargetsForAdGroupRequest {
	return &NullableGlobalRankedKeywordTargetsForAdGroupRequest{value: val, isSet: true}
}

func (v NullableGlobalRankedKeywordTargetsForAdGroupRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalRankedKeywordTargetsForAdGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
