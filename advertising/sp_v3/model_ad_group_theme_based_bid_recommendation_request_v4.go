package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the AdGroupThemeBasedBidRecommendationRequestV4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdGroupThemeBasedBidRecommendationRequestV4{}

// AdGroupThemeBasedBidRecommendationRequestV4 struct for AdGroupThemeBasedBidRecommendationRequestV4
type AdGroupThemeBasedBidRecommendationRequestV4 struct {
	// The list of targeting expressions. Maximum of 100 per request, use pagination for more if needed.
	TargetingExpressions []TargetingExpressionV4 `json:"targetingExpressions"`
	// The campaign identifier.
	CampaignId string `json:"campaignId"`
	// The bid recommendation type.
	RecommendationType string `json:"recommendationType"`
	// The ad group identifier.
	AdGroupId string `json:"adGroupId"`
}

type _AdGroupThemeBasedBidRecommendationRequestV4 AdGroupThemeBasedBidRecommendationRequestV4

// NewAdGroupThemeBasedBidRecommendationRequestV4 instantiates a new AdGroupThemeBasedBidRecommendationRequestV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdGroupThemeBasedBidRecommendationRequestV4(targetingExpressions []TargetingExpressionV4, campaignId string, recommendationType string, adGroupId string) *AdGroupThemeBasedBidRecommendationRequestV4 {
	this := AdGroupThemeBasedBidRecommendationRequestV4{}
	this.TargetingExpressions = targetingExpressions
	this.CampaignId = campaignId
	this.RecommendationType = recommendationType
	this.AdGroupId = adGroupId
	return &this
}

// NewAdGroupThemeBasedBidRecommendationRequestV4WithDefaults instantiates a new AdGroupThemeBasedBidRecommendationRequestV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdGroupThemeBasedBidRecommendationRequestV4WithDefaults() *AdGroupThemeBasedBidRecommendationRequestV4 {
	this := AdGroupThemeBasedBidRecommendationRequestV4{}
	return &this
}

// GetTargetingExpressions returns the TargetingExpressions field value
func (o *AdGroupThemeBasedBidRecommendationRequestV4) GetTargetingExpressions() []TargetingExpressionV4 {
	if o == nil {
		var ret []TargetingExpressionV4
		return ret
	}

	return o.TargetingExpressions
}

// GetTargetingExpressionsOk returns a tuple with the TargetingExpressions field value
// and a boolean to check if the value has been set.
func (o *AdGroupThemeBasedBidRecommendationRequestV4) GetTargetingExpressionsOk() ([]TargetingExpressionV4, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetingExpressions, true
}

// SetTargetingExpressions sets field value
func (o *AdGroupThemeBasedBidRecommendationRequestV4) SetTargetingExpressions(v []TargetingExpressionV4) {
	o.TargetingExpressions = v
}

// GetCampaignId returns the CampaignId field value
func (o *AdGroupThemeBasedBidRecommendationRequestV4) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *AdGroupThemeBasedBidRecommendationRequestV4) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *AdGroupThemeBasedBidRecommendationRequestV4) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetRecommendationType returns the RecommendationType field value
func (o *AdGroupThemeBasedBidRecommendationRequestV4) GetRecommendationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecommendationType
}

// GetRecommendationTypeOk returns a tuple with the RecommendationType field value
// and a boolean to check if the value has been set.
func (o *AdGroupThemeBasedBidRecommendationRequestV4) GetRecommendationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendationType, true
}

// SetRecommendationType sets field value
func (o *AdGroupThemeBasedBidRecommendationRequestV4) SetRecommendationType(v string) {
	o.RecommendationType = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *AdGroupThemeBasedBidRecommendationRequestV4) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *AdGroupThemeBasedBidRecommendationRequestV4) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *AdGroupThemeBasedBidRecommendationRequestV4) SetAdGroupId(v string) {
	o.AdGroupId = v
}

func (o AdGroupThemeBasedBidRecommendationRequestV4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingExpressions"] = o.TargetingExpressions
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["recommendationType"] = o.RecommendationType
	toSerialize["adGroupId"] = o.AdGroupId
	return toSerialize, nil
}

type NullableAdGroupThemeBasedBidRecommendationRequestV4 struct {
	value *AdGroupThemeBasedBidRecommendationRequestV4
	isSet bool
}

func (v NullableAdGroupThemeBasedBidRecommendationRequestV4) Get() *AdGroupThemeBasedBidRecommendationRequestV4 {
	return v.value
}

func (v *NullableAdGroupThemeBasedBidRecommendationRequestV4) Set(val *AdGroupThemeBasedBidRecommendationRequestV4) {
	v.value = val
	v.isSet = true
}

func (v NullableAdGroupThemeBasedBidRecommendationRequestV4) IsSet() bool {
	return v.isSet
}

func (v *NullableAdGroupThemeBasedBidRecommendationRequestV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdGroupThemeBasedBidRecommendationRequestV4(val *AdGroupThemeBasedBidRecommendationRequestV4) *NullableAdGroupThemeBasedBidRecommendationRequestV4 {
	return &NullableAdGroupThemeBasedBidRecommendationRequestV4{value: val, isSet: true}
}

func (v NullableAdGroupThemeBasedBidRecommendationRequestV4) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdGroupThemeBasedBidRecommendationRequestV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
