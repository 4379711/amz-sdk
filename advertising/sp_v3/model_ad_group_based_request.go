package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the AdGroupBasedRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdGroupBasedRequest{}

// AdGroupBasedRequest struct for AdGroupBasedRequest
type AdGroupBasedRequest struct {
	// The identifier of the campaign
	CampaignId *string `json:"campaignId,omitempty"`
	// The recommendationType to retrieve recommended keyword targets for an existing ad group.
	RecommendationType *string `json:"recommendationType,omitempty"`
	// Set this parameter to false if you do not want to retrieve bid suggestions for your keyword targets. Defaults to true.
	BidsEnabled *bool `json:"bidsEnabled,omitempty"`
	// The identifier of the ad group
	AdGroupId *string `json:"adGroupId,omitempty"`
	// The max size of recommended target. Set it to 0 if you only want to rank user-defined keywords.
	MaxRecommendations *float32 `json:"maxRecommendations,omitempty"`
	// The ranking metric value. Supported values: CLICKS, CONVERSIONS, DEFAULT. DEFAULT will be applied if no value passed in.
	SortDimension *string `json:"sortDimension,omitempty"`
	// Translations are for readability and do not affect the targeting of ads. Supported marketplace to locale mappings can be found at the <a href='https://advertising.amazon.com/API/docs/en-us/localization/#/Keyword%20Localization'>POST /keywords/localize</a> endpoint. Note: Translations will be null if locale is unsupported.
	Locale *string `json:"locale,omitempty"`
}

// NewAdGroupBasedRequest instantiates a new AdGroupBasedRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdGroupBasedRequest() *AdGroupBasedRequest {
	this := AdGroupBasedRequest{}
	return &this
}

// NewAdGroupBasedRequestWithDefaults instantiates a new AdGroupBasedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdGroupBasedRequestWithDefaults() *AdGroupBasedRequest {
	this := AdGroupBasedRequest{}
	var bidsEnabled bool = true
	this.BidsEnabled = &bidsEnabled
	return &this
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *AdGroupBasedRequest) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupBasedRequest) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *AdGroupBasedRequest) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *AdGroupBasedRequest) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetRecommendationType returns the RecommendationType field value if set, zero value otherwise.
func (o *AdGroupBasedRequest) GetRecommendationType() string {
	if o == nil || IsNil(o.RecommendationType) {
		var ret string
		return ret
	}
	return *o.RecommendationType
}

// GetRecommendationTypeOk returns a tuple with the RecommendationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupBasedRequest) GetRecommendationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecommendationType) {
		return nil, false
	}
	return o.RecommendationType, true
}

// HasRecommendationType returns a boolean if a field has been set.
func (o *AdGroupBasedRequest) HasRecommendationType() bool {
	if o != nil && !IsNil(o.RecommendationType) {
		return true
	}

	return false
}

// SetRecommendationType gets a reference to the given string and assigns it to the RecommendationType field.
func (o *AdGroupBasedRequest) SetRecommendationType(v string) {
	o.RecommendationType = &v
}

// GetBidsEnabled returns the BidsEnabled field value if set, zero value otherwise.
func (o *AdGroupBasedRequest) GetBidsEnabled() bool {
	if o == nil || IsNil(o.BidsEnabled) {
		var ret bool
		return ret
	}
	return *o.BidsEnabled
}

// GetBidsEnabledOk returns a tuple with the BidsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupBasedRequest) GetBidsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BidsEnabled) {
		return nil, false
	}
	return o.BidsEnabled, true
}

// HasBidsEnabled returns a boolean if a field has been set.
func (o *AdGroupBasedRequest) HasBidsEnabled() bool {
	if o != nil && !IsNil(o.BidsEnabled) {
		return true
	}

	return false
}

// SetBidsEnabled gets a reference to the given bool and assigns it to the BidsEnabled field.
func (o *AdGroupBasedRequest) SetBidsEnabled(v bool) {
	o.BidsEnabled = &v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *AdGroupBasedRequest) GetAdGroupId() string {
	if o == nil || IsNil(o.AdGroupId) {
		var ret string
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupBasedRequest) GetAdGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *AdGroupBasedRequest) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given string and assigns it to the AdGroupId field.
func (o *AdGroupBasedRequest) SetAdGroupId(v string) {
	o.AdGroupId = &v
}

// GetMaxRecommendations returns the MaxRecommendations field value if set, zero value otherwise.
func (o *AdGroupBasedRequest) GetMaxRecommendations() float32 {
	if o == nil || IsNil(o.MaxRecommendations) {
		var ret float32
		return ret
	}
	return *o.MaxRecommendations
}

// GetMaxRecommendationsOk returns a tuple with the MaxRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupBasedRequest) GetMaxRecommendationsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxRecommendations) {
		return nil, false
	}
	return o.MaxRecommendations, true
}

// HasMaxRecommendations returns a boolean if a field has been set.
func (o *AdGroupBasedRequest) HasMaxRecommendations() bool {
	if o != nil && !IsNil(o.MaxRecommendations) {
		return true
	}

	return false
}

// SetMaxRecommendations gets a reference to the given float32 and assigns it to the MaxRecommendations field.
func (o *AdGroupBasedRequest) SetMaxRecommendations(v float32) {
	o.MaxRecommendations = &v
}

// GetSortDimension returns the SortDimension field value if set, zero value otherwise.
func (o *AdGroupBasedRequest) GetSortDimension() string {
	if o == nil || IsNil(o.SortDimension) {
		var ret string
		return ret
	}
	return *o.SortDimension
}

// GetSortDimensionOk returns a tuple with the SortDimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupBasedRequest) GetSortDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.SortDimension) {
		return nil, false
	}
	return o.SortDimension, true
}

// HasSortDimension returns a boolean if a field has been set.
func (o *AdGroupBasedRequest) HasSortDimension() bool {
	if o != nil && !IsNil(o.SortDimension) {
		return true
	}

	return false
}

// SetSortDimension gets a reference to the given string and assigns it to the SortDimension field.
func (o *AdGroupBasedRequest) SetSortDimension(v string) {
	o.SortDimension = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *AdGroupBasedRequest) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupBasedRequest) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *AdGroupBasedRequest) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *AdGroupBasedRequest) SetLocale(v string) {
	o.Locale = &v
}

func (o AdGroupBasedRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	if !IsNil(o.RecommendationType) {
		toSerialize["recommendationType"] = o.RecommendationType
	}
	if !IsNil(o.BidsEnabled) {
		toSerialize["bidsEnabled"] = o.BidsEnabled
	}
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
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

type NullableAdGroupBasedRequest struct {
	value *AdGroupBasedRequest
	isSet bool
}

func (v NullableAdGroupBasedRequest) Get() *AdGroupBasedRequest {
	return v.value
}

func (v *NullableAdGroupBasedRequest) Set(val *AdGroupBasedRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAdGroupBasedRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAdGroupBasedRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdGroupBasedRequest(val *AdGroupBasedRequest) *NullableAdGroupBasedRequest {
	return &NullableAdGroupBasedRequest{value: val, isSet: true}
}

func (v NullableAdGroupBasedRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdGroupBasedRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
