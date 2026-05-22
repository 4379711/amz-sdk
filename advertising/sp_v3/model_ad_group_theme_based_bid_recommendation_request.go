package sp_v3

import "github.com/bytedance/sonic"

// checks if the AdGroupThemeBasedBidRecommendationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdGroupThemeBasedBidRecommendationRequest{}

// AdGroupThemeBasedBidRecommendationRequest struct for AdGroupThemeBasedBidRecommendationRequest
type AdGroupThemeBasedBidRecommendationRequest struct {
	// The list of targeting expressions. Maximum of 100 per request, use pagination for more if needed.
	TargetingExpressions []TargetingExpression `json:"targetingExpressions"`
	// The campaign identifier.
	CampaignId string `json:"campaignId"`
	// The bid recommendation type.
	RecommendationType string `json:"recommendationType"`
	// The ad group identifier.
	AdGroupId string `json:"adGroupId"`
}

type _AdGroupThemeBasedBidRecommendationRequest AdGroupThemeBasedBidRecommendationRequest

// NewAdGroupThemeBasedBidRecommendationRequest instantiates a new AdGroupThemeBasedBidRecommendationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdGroupThemeBasedBidRecommendationRequest(targetingExpressions []TargetingExpression, campaignId string, recommendationType string, adGroupId string) *AdGroupThemeBasedBidRecommendationRequest {
	this := AdGroupThemeBasedBidRecommendationRequest{}
	this.TargetingExpressions = targetingExpressions
	this.CampaignId = campaignId
	this.RecommendationType = recommendationType
	this.AdGroupId = adGroupId
	return &this
}

// NewAdGroupThemeBasedBidRecommendationRequestWithDefaults instantiates a new AdGroupThemeBasedBidRecommendationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdGroupThemeBasedBidRecommendationRequestWithDefaults() *AdGroupThemeBasedBidRecommendationRequest {
	this := AdGroupThemeBasedBidRecommendationRequest{}
	return &this
}

// GetTargetingExpressions returns the TargetingExpressions field value
func (o *AdGroupThemeBasedBidRecommendationRequest) GetTargetingExpressions() []TargetingExpression {
	if o == nil {
		var ret []TargetingExpression
		return ret
	}

	return o.TargetingExpressions
}

// GetTargetingExpressionsOk returns a tuple with the TargetingExpressions field value
// and a boolean to check if the value has been set.
func (o *AdGroupThemeBasedBidRecommendationRequest) GetTargetingExpressionsOk() ([]TargetingExpression, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetingExpressions, true
}

// SetTargetingExpressions sets field value
func (o *AdGroupThemeBasedBidRecommendationRequest) SetTargetingExpressions(v []TargetingExpression) {
	o.TargetingExpressions = v
}

// GetCampaignId returns the CampaignId field value
func (o *AdGroupThemeBasedBidRecommendationRequest) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *AdGroupThemeBasedBidRecommendationRequest) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *AdGroupThemeBasedBidRecommendationRequest) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetRecommendationType returns the RecommendationType field value
func (o *AdGroupThemeBasedBidRecommendationRequest) GetRecommendationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecommendationType
}

// GetRecommendationTypeOk returns a tuple with the RecommendationType field value
// and a boolean to check if the value has been set.
func (o *AdGroupThemeBasedBidRecommendationRequest) GetRecommendationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendationType, true
}

// SetRecommendationType sets field value
func (o *AdGroupThemeBasedBidRecommendationRequest) SetRecommendationType(v string) {
	o.RecommendationType = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *AdGroupThemeBasedBidRecommendationRequest) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *AdGroupThemeBasedBidRecommendationRequest) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *AdGroupThemeBasedBidRecommendationRequest) SetAdGroupId(v string) {
	o.AdGroupId = v
}

func (o AdGroupThemeBasedBidRecommendationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingExpressions"] = o.TargetingExpressions
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["recommendationType"] = o.RecommendationType
	toSerialize["adGroupId"] = o.AdGroupId
	return toSerialize, nil
}

type NullableAdGroupThemeBasedBidRecommendationRequest struct {
	value *AdGroupThemeBasedBidRecommendationRequest
	isSet bool
}

func (v NullableAdGroupThemeBasedBidRecommendationRequest) Get() *AdGroupThemeBasedBidRecommendationRequest {
	return v.value
}

func (v *NullableAdGroupThemeBasedBidRecommendationRequest) Set(val *AdGroupThemeBasedBidRecommendationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAdGroupThemeBasedBidRecommendationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAdGroupThemeBasedBidRecommendationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdGroupThemeBasedBidRecommendationRequest(val *AdGroupThemeBasedBidRecommendationRequest) *NullableAdGroupThemeBasedBidRecommendationRequest {
	return &NullableAdGroupThemeBasedBidRecommendationRequest{value: val, isSet: true}
}

func (v NullableAdGroupThemeBasedBidRecommendationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdGroupThemeBasedBidRecommendationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
