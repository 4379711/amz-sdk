package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetRuleRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetRuleRecommendation{}

// BudgetRuleRecommendation struct for BudgetRuleRecommendation
type BudgetRuleRecommendation struct {
	// suggested increase percent
	SuggestedBudgetIncreasePercent *float32 `json:"suggestedBudgetIncreasePercent,omitempty"`
	// rule name for the recomemendation
	RuleName *string `json:"ruleName,omitempty"`
	// rule id for the recomemendation
	RuleId *string `json:"ruleId,omitempty"`
}

// NewBudgetRuleRecommendation instantiates a new BudgetRuleRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetRuleRecommendation() *BudgetRuleRecommendation {
	this := BudgetRuleRecommendation{}
	return &this
}

// NewBudgetRuleRecommendationWithDefaults instantiates a new BudgetRuleRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetRuleRecommendationWithDefaults() *BudgetRuleRecommendation {
	this := BudgetRuleRecommendation{}
	return &this
}

// GetSuggestedBudgetIncreasePercent returns the SuggestedBudgetIncreasePercent field value if set, zero value otherwise.
func (o *BudgetRuleRecommendation) GetSuggestedBudgetIncreasePercent() float32 {
	if o == nil || IsNil(o.SuggestedBudgetIncreasePercent) {
		var ret float32
		return ret
	}
	return *o.SuggestedBudgetIncreasePercent
}

// GetSuggestedBudgetIncreasePercentOk returns a tuple with the SuggestedBudgetIncreasePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetRuleRecommendation) GetSuggestedBudgetIncreasePercentOk() (*float32, bool) {
	if o == nil || IsNil(o.SuggestedBudgetIncreasePercent) {
		return nil, false
	}
	return o.SuggestedBudgetIncreasePercent, true
}

// HasSuggestedBudgetIncreasePercent returns a boolean if a field has been set.
func (o *BudgetRuleRecommendation) HasSuggestedBudgetIncreasePercent() bool {
	if o != nil && !IsNil(o.SuggestedBudgetIncreasePercent) {
		return true
	}

	return false
}

// SetSuggestedBudgetIncreasePercent gets a reference to the given float32 and assigns it to the SuggestedBudgetIncreasePercent field.
func (o *BudgetRuleRecommendation) SetSuggestedBudgetIncreasePercent(v float32) {
	o.SuggestedBudgetIncreasePercent = &v
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *BudgetRuleRecommendation) GetRuleName() string {
	if o == nil || IsNil(o.RuleName) {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetRuleRecommendation) GetRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RuleName) {
		return nil, false
	}
	return o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *BudgetRuleRecommendation) HasRuleName() bool {
	if o != nil && !IsNil(o.RuleName) {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *BudgetRuleRecommendation) SetRuleName(v string) {
	o.RuleName = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *BudgetRuleRecommendation) GetRuleId() string {
	if o == nil || IsNil(o.RuleId) {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetRuleRecommendation) GetRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *BudgetRuleRecommendation) HasRuleId() bool {
	if o != nil && !IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *BudgetRuleRecommendation) SetRuleId(v string) {
	o.RuleId = &v
}

func (o BudgetRuleRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SuggestedBudgetIncreasePercent) {
		toSerialize["suggestedBudgetIncreasePercent"] = o.SuggestedBudgetIncreasePercent
	}
	if !IsNil(o.RuleName) {
		toSerialize["ruleName"] = o.RuleName
	}
	if !IsNil(o.RuleId) {
		toSerialize["ruleId"] = o.RuleId
	}
	return toSerialize, nil
}

type NullableBudgetRuleRecommendation struct {
	value *BudgetRuleRecommendation
	isSet bool
}

func (v NullableBudgetRuleRecommendation) Get() *BudgetRuleRecommendation {
	return v.value
}

func (v *NullableBudgetRuleRecommendation) Set(val *BudgetRuleRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetRuleRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetRuleRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetRuleRecommendation(val *BudgetRuleRecommendation) *NullableBudgetRuleRecommendation {
	return &NullableBudgetRuleRecommendation{value: val, isSet: true}
}

func (v NullableBudgetRuleRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetRuleRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
