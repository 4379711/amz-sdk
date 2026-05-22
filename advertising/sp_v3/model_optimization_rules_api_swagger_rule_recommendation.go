package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIRuleRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIRuleRecommendation{}

// OptimizationRulesAPIRuleRecommendation struct for OptimizationRulesAPIRuleRecommendation
type OptimizationRulesAPIRuleRecommendation struct {
	RuleCategory *OptimizationRulesAPIRuleCategory `json:"ruleCategory,omitempty"`
	// Id of the campaign for which the rule is recommended
	CampaignId      *string                                            `json:"campaignId,omitempty"`
	RuleSubCategory *OptimizationRulesAPIRuleSubCategory               `json:"ruleSubCategory,omitempty"`
	Recommendation  *OptimizationRulesAPIOptimizationRuleWithoutRuleId `json:"recommendation,omitempty"`
	// Unique id for the recommendation.
	RecommendationId *string `json:"recommendationId,omitempty"`
	// Unique id for the rule if it already exists
	OptimizationRuleId *string `json:"optimizationRuleId,omitempty"`
	// Reason for recommendation.
	RecommendationReason *string `json:"recommendationReason,omitempty"`
}

// NewOptimizationRulesAPIRuleRecommendation instantiates a new OptimizationRulesAPIRuleRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIRuleRecommendation() *OptimizationRulesAPIRuleRecommendation {
	this := OptimizationRulesAPIRuleRecommendation{}
	return &this
}

// NewOptimizationRulesAPIRuleRecommendationWithDefaults instantiates a new OptimizationRulesAPIRuleRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIRuleRecommendationWithDefaults() *OptimizationRulesAPIRuleRecommendation {
	this := OptimizationRulesAPIRuleRecommendation{}
	return &this
}

// GetRuleCategory returns the RuleCategory field value if set, zero value otherwise.
func (o *OptimizationRulesAPIRuleRecommendation) GetRuleCategory() OptimizationRulesAPIRuleCategory {
	if o == nil || IsNil(o.RuleCategory) {
		var ret OptimizationRulesAPIRuleCategory
		return ret
	}
	return *o.RuleCategory
}

// GetRuleCategoryOk returns a tuple with the RuleCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleRecommendation) GetRuleCategoryOk() (*OptimizationRulesAPIRuleCategory, bool) {
	if o == nil || IsNil(o.RuleCategory) {
		return nil, false
	}
	return o.RuleCategory, true
}

// HasRuleCategory returns a boolean if a field has been set.
func (o *OptimizationRulesAPIRuleRecommendation) HasRuleCategory() bool {
	if o != nil && !IsNil(o.RuleCategory) {
		return true
	}

	return false
}

// SetRuleCategory gets a reference to the given OptimizationRulesAPIRuleCategory and assigns it to the RuleCategory field.
func (o *OptimizationRulesAPIRuleRecommendation) SetRuleCategory(v OptimizationRulesAPIRuleCategory) {
	o.RuleCategory = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *OptimizationRulesAPIRuleRecommendation) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleRecommendation) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *OptimizationRulesAPIRuleRecommendation) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *OptimizationRulesAPIRuleRecommendation) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetRuleSubCategory returns the RuleSubCategory field value if set, zero value otherwise.
func (o *OptimizationRulesAPIRuleRecommendation) GetRuleSubCategory() OptimizationRulesAPIRuleSubCategory {
	if o == nil || IsNil(o.RuleSubCategory) {
		var ret OptimizationRulesAPIRuleSubCategory
		return ret
	}
	return *o.RuleSubCategory
}

// GetRuleSubCategoryOk returns a tuple with the RuleSubCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleRecommendation) GetRuleSubCategoryOk() (*OptimizationRulesAPIRuleSubCategory, bool) {
	if o == nil || IsNil(o.RuleSubCategory) {
		return nil, false
	}
	return o.RuleSubCategory, true
}

// HasRuleSubCategory returns a boolean if a field has been set.
func (o *OptimizationRulesAPIRuleRecommendation) HasRuleSubCategory() bool {
	if o != nil && !IsNil(o.RuleSubCategory) {
		return true
	}

	return false
}

// SetRuleSubCategory gets a reference to the given OptimizationRulesAPIRuleSubCategory and assigns it to the RuleSubCategory field.
func (o *OptimizationRulesAPIRuleRecommendation) SetRuleSubCategory(v OptimizationRulesAPIRuleSubCategory) {
	o.RuleSubCategory = &v
}

// GetRecommendation returns the Recommendation field value if set, zero value otherwise.
func (o *OptimizationRulesAPIRuleRecommendation) GetRecommendation() OptimizationRulesAPIOptimizationRuleWithoutRuleId {
	if o == nil || IsNil(o.Recommendation) {
		var ret OptimizationRulesAPIOptimizationRuleWithoutRuleId
		return ret
	}
	return *o.Recommendation
}

// GetRecommendationOk returns a tuple with the Recommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleRecommendation) GetRecommendationOk() (*OptimizationRulesAPIOptimizationRuleWithoutRuleId, bool) {
	if o == nil || IsNil(o.Recommendation) {
		return nil, false
	}
	return o.Recommendation, true
}

// HasRecommendation returns a boolean if a field has been set.
func (o *OptimizationRulesAPIRuleRecommendation) HasRecommendation() bool {
	if o != nil && !IsNil(o.Recommendation) {
		return true
	}

	return false
}

// SetRecommendation gets a reference to the given OptimizationRulesAPIOptimizationRuleWithoutRuleId and assigns it to the Recommendation field.
func (o *OptimizationRulesAPIRuleRecommendation) SetRecommendation(v OptimizationRulesAPIOptimizationRuleWithoutRuleId) {
	o.Recommendation = &v
}

// GetRecommendationId returns the RecommendationId field value if set, zero value otherwise.
func (o *OptimizationRulesAPIRuleRecommendation) GetRecommendationId() string {
	if o == nil || IsNil(o.RecommendationId) {
		var ret string
		return ret
	}
	return *o.RecommendationId
}

// GetRecommendationIdOk returns a tuple with the RecommendationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleRecommendation) GetRecommendationIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecommendationId) {
		return nil, false
	}
	return o.RecommendationId, true
}

// HasRecommendationId returns a boolean if a field has been set.
func (o *OptimizationRulesAPIRuleRecommendation) HasRecommendationId() bool {
	if o != nil && !IsNil(o.RecommendationId) {
		return true
	}

	return false
}

// SetRecommendationId gets a reference to the given string and assigns it to the RecommendationId field.
func (o *OptimizationRulesAPIRuleRecommendation) SetRecommendationId(v string) {
	o.RecommendationId = &v
}

// GetOptimizationRuleId returns the OptimizationRuleId field value if set, zero value otherwise.
func (o *OptimizationRulesAPIRuleRecommendation) GetOptimizationRuleId() string {
	if o == nil || IsNil(o.OptimizationRuleId) {
		var ret string
		return ret
	}
	return *o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleRecommendation) GetOptimizationRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.OptimizationRuleId) {
		return nil, false
	}
	return o.OptimizationRuleId, true
}

// HasOptimizationRuleId returns a boolean if a field has been set.
func (o *OptimizationRulesAPIRuleRecommendation) HasOptimizationRuleId() bool {
	if o != nil && !IsNil(o.OptimizationRuleId) {
		return true
	}

	return false
}

// SetOptimizationRuleId gets a reference to the given string and assigns it to the OptimizationRuleId field.
func (o *OptimizationRulesAPIRuleRecommendation) SetOptimizationRuleId(v string) {
	o.OptimizationRuleId = &v
}

// GetRecommendationReason returns the RecommendationReason field value if set, zero value otherwise.
func (o *OptimizationRulesAPIRuleRecommendation) GetRecommendationReason() string {
	if o == nil || IsNil(o.RecommendationReason) {
		var ret string
		return ret
	}
	return *o.RecommendationReason
}

// GetRecommendationReasonOk returns a tuple with the RecommendationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleRecommendation) GetRecommendationReasonOk() (*string, bool) {
	if o == nil || IsNil(o.RecommendationReason) {
		return nil, false
	}
	return o.RecommendationReason, true
}

// HasRecommendationReason returns a boolean if a field has been set.
func (o *OptimizationRulesAPIRuleRecommendation) HasRecommendationReason() bool {
	if o != nil && !IsNil(o.RecommendationReason) {
		return true
	}

	return false
}

// SetRecommendationReason gets a reference to the given string and assigns it to the RecommendationReason field.
func (o *OptimizationRulesAPIRuleRecommendation) SetRecommendationReason(v string) {
	o.RecommendationReason = &v
}

func (o OptimizationRulesAPIRuleRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RuleCategory) {
		toSerialize["ruleCategory"] = o.RuleCategory
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	if !IsNil(o.RuleSubCategory) {
		toSerialize["ruleSubCategory"] = o.RuleSubCategory
	}
	if !IsNil(o.Recommendation) {
		toSerialize["recommendation"] = o.Recommendation
	}
	if !IsNil(o.RecommendationId) {
		toSerialize["recommendationId"] = o.RecommendationId
	}
	if !IsNil(o.OptimizationRuleId) {
		toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	}
	if !IsNil(o.RecommendationReason) {
		toSerialize["recommendationReason"] = o.RecommendationReason
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIRuleRecommendation struct {
	value *OptimizationRulesAPIRuleRecommendation
	isSet bool
}

func (v NullableOptimizationRulesAPIRuleRecommendation) Get() *OptimizationRulesAPIRuleRecommendation {
	return v.value
}

func (v *NullableOptimizationRulesAPIRuleRecommendation) Set(val *OptimizationRulesAPIRuleRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIRuleRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIRuleRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIRuleRecommendation(val *OptimizationRulesAPIRuleRecommendation) *NullableOptimizationRulesAPIRuleRecommendation {
	return &NullableOptimizationRulesAPIRuleRecommendation{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIRuleRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIRuleRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
