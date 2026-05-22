package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the RankedKeywordTargetsForAdGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RankedKeywordTargetsForAdGroupRequest{}

// RankedKeywordTargetsForAdGroupRequest This request type is used to retrieve recommended keyword targets for an existing ad group. Set the recommendationType to KEYWORDS_FOR_ADGROUP to use this request type.
type RankedKeywordTargetsForAdGroupRequest struct {
	// The identifier of the campaign
	CampaignId string `json:"campaignId"`
	// The recommendationType to retrieve recommended keyword targets for an existing ad group.
	RecommendationType string `json:"recommendationType"`
	// Set this parameter to false if you do not want to retrieve bid suggestions for your keyword targets. Defaults to true.
	BidsEnabled *bool `json:"bidsEnabled,omitempty"`
	// The identifier of the ad group
	AdGroupId string `json:"adGroupId"`
	// A list of targets that need to be ranked
	Targets []KeywordTarget `json:"targets,omitempty"`
	// The max size of recommended target. Set it to 0 if you only want to rank user-defined keywords.
	MaxRecommendations *float32 `json:"maxRecommendations,omitempty"`
	// The ranking metric value. Supported values: CLICKS, CONVERSIONS, DEFAULT. DEFAULT will be applied if no value passed in.
	SortDimension *string `json:"sortDimension,omitempty"`
	// Translations are for readability and do not affect the targeting of ads. Supported marketplace to locale mappings can be found at the <a href='https://advertising.amazon.com/API/docs/en-us/localization/#/Keyword%20Localization'>POST /keywords/localize</a> endpoint. Note: Translations will be null if locale is unsupported.
	Locale *string `json:"locale,omitempty"`
}

type _RankedKeywordTargetsForAdGroupRequest RankedKeywordTargetsForAdGroupRequest

// NewRankedKeywordTargetsForAdGroupRequest instantiates a new RankedKeywordTargetsForAdGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRankedKeywordTargetsForAdGroupRequest(campaignId string, recommendationType string, adGroupId string) *RankedKeywordTargetsForAdGroupRequest {
	this := RankedKeywordTargetsForAdGroupRequest{}
	return &this
}

// NewRankedKeywordTargetsForAdGroupRequestWithDefaults instantiates a new RankedKeywordTargetsForAdGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRankedKeywordTargetsForAdGroupRequestWithDefaults() *RankedKeywordTargetsForAdGroupRequest {
	this := RankedKeywordTargetsForAdGroupRequest{}
	var bidsEnabled bool = true
	this.BidsEnabled = &bidsEnabled
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *RankedKeywordTargetsForAdGroupRequest) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *RankedKeywordTargetsForAdGroupRequest) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *RankedKeywordTargetsForAdGroupRequest) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetRecommendationType returns the RecommendationType field value
func (o *RankedKeywordTargetsForAdGroupRequest) GetRecommendationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecommendationType
}

// GetRecommendationTypeOk returns a tuple with the RecommendationType field value
// and a boolean to check if the value has been set.
func (o *RankedKeywordTargetsForAdGroupRequest) GetRecommendationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendationType, true
}

// SetRecommendationType sets field value
func (o *RankedKeywordTargetsForAdGroupRequest) SetRecommendationType(v string) {
	o.RecommendationType = v
}

// GetBidsEnabled returns the BidsEnabled field value if set, zero value otherwise.
func (o *RankedKeywordTargetsForAdGroupRequest) GetBidsEnabled() bool {
	if o == nil || IsNil(o.BidsEnabled) {
		var ret bool
		return ret
	}
	return *o.BidsEnabled
}

// GetBidsEnabledOk returns a tuple with the BidsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RankedKeywordTargetsForAdGroupRequest) GetBidsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BidsEnabled) {
		return nil, false
	}
	return o.BidsEnabled, true
}

// HasBidsEnabled returns a boolean if a field has been set.
func (o *RankedKeywordTargetsForAdGroupRequest) HasBidsEnabled() bool {
	if o != nil && !IsNil(o.BidsEnabled) {
		return true
	}

	return false
}

// SetBidsEnabled gets a reference to the given bool and assigns it to the BidsEnabled field.
func (o *RankedKeywordTargetsForAdGroupRequest) SetBidsEnabled(v bool) {
	o.BidsEnabled = &v
}

// GetAdGroupId returns the AdGroupId field value
func (o *RankedKeywordTargetsForAdGroupRequest) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *RankedKeywordTargetsForAdGroupRequest) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *RankedKeywordTargetsForAdGroupRequest) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *RankedKeywordTargetsForAdGroupRequest) GetTargets() []KeywordTarget {
	if o == nil || IsNil(o.Targets) {
		var ret []KeywordTarget
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RankedKeywordTargetsForAdGroupRequest) GetTargetsOk() ([]KeywordTarget, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *RankedKeywordTargetsForAdGroupRequest) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []KeywordTarget and assigns it to the Targets field.
func (o *RankedKeywordTargetsForAdGroupRequest) SetTargets(v []KeywordTarget) {
	o.Targets = v
}

// GetMaxRecommendations returns the MaxRecommendations field value if set, zero value otherwise.
func (o *RankedKeywordTargetsForAdGroupRequest) GetMaxRecommendations() float32 {
	if o == nil || IsNil(o.MaxRecommendations) {
		var ret float32
		return ret
	}
	return *o.MaxRecommendations
}

// GetMaxRecommendationsOk returns a tuple with the MaxRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RankedKeywordTargetsForAdGroupRequest) GetMaxRecommendationsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxRecommendations) {
		return nil, false
	}
	return o.MaxRecommendations, true
}

// HasMaxRecommendations returns a boolean if a field has been set.
func (o *RankedKeywordTargetsForAdGroupRequest) HasMaxRecommendations() bool {
	if o != nil && !IsNil(o.MaxRecommendations) {
		return true
	}

	return false
}

// SetMaxRecommendations gets a reference to the given float32 and assigns it to the MaxRecommendations field.
func (o *RankedKeywordTargetsForAdGroupRequest) SetMaxRecommendations(v float32) {
	o.MaxRecommendations = &v
}

// GetSortDimension returns the SortDimension field value if set, zero value otherwise.
func (o *RankedKeywordTargetsForAdGroupRequest) GetSortDimension() string {
	if o == nil || IsNil(o.SortDimension) {
		var ret string
		return ret
	}
	return *o.SortDimension
}

// GetSortDimensionOk returns a tuple with the SortDimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RankedKeywordTargetsForAdGroupRequest) GetSortDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.SortDimension) {
		return nil, false
	}
	return o.SortDimension, true
}

// HasSortDimension returns a boolean if a field has been set.
func (o *RankedKeywordTargetsForAdGroupRequest) HasSortDimension() bool {
	if o != nil && !IsNil(o.SortDimension) {
		return true
	}

	return false
}

// SetSortDimension gets a reference to the given string and assigns it to the SortDimension field.
func (o *RankedKeywordTargetsForAdGroupRequest) SetSortDimension(v string) {
	o.SortDimension = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *RankedKeywordTargetsForAdGroupRequest) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RankedKeywordTargetsForAdGroupRequest) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *RankedKeywordTargetsForAdGroupRequest) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *RankedKeywordTargetsForAdGroupRequest) SetLocale(v string) {
	o.Locale = &v
}

func (o RankedKeywordTargetsForAdGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["recommendationType"] = o.RecommendationType
	if !IsNil(o.BidsEnabled) {
		toSerialize["bidsEnabled"] = o.BidsEnabled
	}
	toSerialize["adGroupId"] = o.AdGroupId
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

type NullableRankedKeywordTargetsForAdGroupRequest struct {
	value *RankedKeywordTargetsForAdGroupRequest
	isSet bool
}

func (v NullableRankedKeywordTargetsForAdGroupRequest) Get() *RankedKeywordTargetsForAdGroupRequest {
	return v.value
}

func (v *NullableRankedKeywordTargetsForAdGroupRequest) Set(val *RankedKeywordTargetsForAdGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRankedKeywordTargetsForAdGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRankedKeywordTargetsForAdGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRankedKeywordTargetsForAdGroupRequest(val *RankedKeywordTargetsForAdGroupRequest) *NullableRankedKeywordTargetsForAdGroupRequest {
	return &NullableRankedKeywordTargetsForAdGroupRequest{value: val, isSet: true}
}

func (v NullableRankedKeywordTargetsForAdGroupRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRankedKeywordTargetsForAdGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
