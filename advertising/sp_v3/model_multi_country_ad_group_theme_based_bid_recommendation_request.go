package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the MultiCountryAdGroupThemeBasedBidRecommendationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiCountryAdGroupThemeBasedBidRecommendationRequest{}

// MultiCountryAdGroupThemeBasedBidRecommendationRequest struct for MultiCountryAdGroupThemeBasedBidRecommendationRequest
type MultiCountryAdGroupThemeBasedBidRecommendationRequest struct {
	// The list of targeting expressions. Maximum of 100 per request per country, use pagination for more if needed.
	TargetingExpressions []MultiCountryTargetingExpression `json:"targetingExpressions"`
	// A list of country codes. Supported country codes: | Country Code |  Country            | |-------------|----------------------| | US          | United States        | | CA          | Canada               | | MX          | Mexico               | | BR          | Brazil               | | UK          | United Kingdom       | | DE          | Germany              | | FR          | France               | | ES          | Spain                | | IT          | Italy                | | IN          | India                | | AE          | United Arab Emirates | | SA          | Saudi Arabia         | | NL          | Netherlands          | | PL          | Poland               | | BE          | Belgium              | | SE          | Sweden               | | TR          | Turkey               | | EG          | Egypt                | | JP          | Japan                | | AU          | Australia            | | SG          | Singapore            |
	CountryCodes []string `json:"countryCodes,omitempty"`
	// The campaign identifier.
	CampaignId string `json:"campaignId"`
	// The bid recommendation type.
	RecommendationType string `json:"recommendationType"`
	// The ad group identifier.
	AdGroupId string `json:"adGroupId"`
}

type _MultiCountryAdGroupThemeBasedBidRecommendationRequest MultiCountryAdGroupThemeBasedBidRecommendationRequest

// NewMultiCountryAdGroupThemeBasedBidRecommendationRequest instantiates a new MultiCountryAdGroupThemeBasedBidRecommendationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiCountryAdGroupThemeBasedBidRecommendationRequest(targetingExpressions []MultiCountryTargetingExpression, campaignId string, recommendationType string, adGroupId string) *MultiCountryAdGroupThemeBasedBidRecommendationRequest {
	this := MultiCountryAdGroupThemeBasedBidRecommendationRequest{}
	this.TargetingExpressions = targetingExpressions
	this.CampaignId = campaignId
	this.RecommendationType = recommendationType
	this.AdGroupId = adGroupId
	return &this
}

// NewMultiCountryAdGroupThemeBasedBidRecommendationRequestWithDefaults instantiates a new MultiCountryAdGroupThemeBasedBidRecommendationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiCountryAdGroupThemeBasedBidRecommendationRequestWithDefaults() *MultiCountryAdGroupThemeBasedBidRecommendationRequest {
	this := MultiCountryAdGroupThemeBasedBidRecommendationRequest{}
	return &this
}

// GetTargetingExpressions returns the TargetingExpressions field value
func (o *MultiCountryAdGroupThemeBasedBidRecommendationRequest) GetTargetingExpressions() []MultiCountryTargetingExpression {
	if o == nil {
		var ret []MultiCountryTargetingExpression
		return ret
	}

	return o.TargetingExpressions
}

// GetTargetingExpressionsOk returns a tuple with the TargetingExpressions field value
// and a boolean to check if the value has been set.
func (o *MultiCountryAdGroupThemeBasedBidRecommendationRequest) GetTargetingExpressionsOk() ([]MultiCountryTargetingExpression, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetingExpressions, true
}

// SetTargetingExpressions sets field value
func (o *MultiCountryAdGroupThemeBasedBidRecommendationRequest) SetTargetingExpressions(v []MultiCountryTargetingExpression) {
	o.TargetingExpressions = v
}

// GetCountryCodes returns the CountryCodes field value if set, zero value otherwise.
func (o *MultiCountryAdGroupThemeBasedBidRecommendationRequest) GetCountryCodes() []string {
	if o == nil || IsNil(o.CountryCodes) {
		var ret []string
		return ret
	}
	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiCountryAdGroupThemeBasedBidRecommendationRequest) GetCountryCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryCodes) {
		return nil, false
	}
	return o.CountryCodes, true
}

// HasCountryCodes returns a boolean if a field has been set.
func (o *MultiCountryAdGroupThemeBasedBidRecommendationRequest) HasCountryCodes() bool {
	if o != nil && !IsNil(o.CountryCodes) {
		return true
	}

	return false
}

// SetCountryCodes gets a reference to the given []string and assigns it to the CountryCodes field.
func (o *MultiCountryAdGroupThemeBasedBidRecommendationRequest) SetCountryCodes(v []string) {
	o.CountryCodes = v
}

// GetCampaignId returns the CampaignId field value
func (o *MultiCountryAdGroupThemeBasedBidRecommendationRequest) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *MultiCountryAdGroupThemeBasedBidRecommendationRequest) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *MultiCountryAdGroupThemeBasedBidRecommendationRequest) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetRecommendationType returns the RecommendationType field value
func (o *MultiCountryAdGroupThemeBasedBidRecommendationRequest) GetRecommendationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecommendationType
}

// GetRecommendationTypeOk returns a tuple with the RecommendationType field value
// and a boolean to check if the value has been set.
func (o *MultiCountryAdGroupThemeBasedBidRecommendationRequest) GetRecommendationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendationType, true
}

// SetRecommendationType sets field value
func (o *MultiCountryAdGroupThemeBasedBidRecommendationRequest) SetRecommendationType(v string) {
	o.RecommendationType = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *MultiCountryAdGroupThemeBasedBidRecommendationRequest) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *MultiCountryAdGroupThemeBasedBidRecommendationRequest) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *MultiCountryAdGroupThemeBasedBidRecommendationRequest) SetAdGroupId(v string) {
	o.AdGroupId = v
}

func (o MultiCountryAdGroupThemeBasedBidRecommendationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingExpressions"] = o.TargetingExpressions
	if !IsNil(o.CountryCodes) {
		toSerialize["countryCodes"] = o.CountryCodes
	}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["recommendationType"] = o.RecommendationType
	toSerialize["adGroupId"] = o.AdGroupId
	return toSerialize, nil
}

type NullableMultiCountryAdGroupThemeBasedBidRecommendationRequest struct {
	value *MultiCountryAdGroupThemeBasedBidRecommendationRequest
	isSet bool
}

func (v NullableMultiCountryAdGroupThemeBasedBidRecommendationRequest) Get() *MultiCountryAdGroupThemeBasedBidRecommendationRequest {
	return v.value
}

func (v *NullableMultiCountryAdGroupThemeBasedBidRecommendationRequest) Set(val *MultiCountryAdGroupThemeBasedBidRecommendationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiCountryAdGroupThemeBasedBidRecommendationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiCountryAdGroupThemeBasedBidRecommendationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiCountryAdGroupThemeBasedBidRecommendationRequest(val *MultiCountryAdGroupThemeBasedBidRecommendationRequest) *NullableMultiCountryAdGroupThemeBasedBidRecommendationRequest {
	return &NullableMultiCountryAdGroupThemeBasedBidRecommendationRequest{value: val, isSet: true}
}

func (v NullableMultiCountryAdGroupThemeBasedBidRecommendationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMultiCountryAdGroupThemeBasedBidRecommendationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
