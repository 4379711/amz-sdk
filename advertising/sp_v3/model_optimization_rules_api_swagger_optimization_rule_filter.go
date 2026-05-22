package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIOptimizationRuleFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIOptimizationRuleFilter{}

// OptimizationRulesAPIOptimizationRuleFilter Filter on optimization rules.
type OptimizationRulesAPIOptimizationRuleFilter struct {
	RuleCategory       *OptimizationRulesAPIEntityFieldFilter `json:"ruleCategory,omitempty"`
	RuleSubCategory    *OptimizationRulesAPIEntityFieldFilter `json:"ruleSubCategory,omitempty"`
	OptimizationRuleId *OptimizationRulesAPIEntityFieldFilter `json:"optimizationRuleId,omitempty"`
}

// NewOptimizationRulesAPIOptimizationRuleFilter instantiates a new OptimizationRulesAPIOptimizationRuleFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIOptimizationRuleFilter() *OptimizationRulesAPIOptimizationRuleFilter {
	this := OptimizationRulesAPIOptimizationRuleFilter{}
	return &this
}

// NewOptimizationRulesAPIOptimizationRuleFilterWithDefaults instantiates a new OptimizationRulesAPIOptimizationRuleFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIOptimizationRuleFilterWithDefaults() *OptimizationRulesAPIOptimizationRuleFilter {
	this := OptimizationRulesAPIOptimizationRuleFilter{}
	return &this
}

// GetRuleCategory returns the RuleCategory field value if set, zero value otherwise.
func (o *OptimizationRulesAPIOptimizationRuleFilter) GetRuleCategory() OptimizationRulesAPIEntityFieldFilter {
	if o == nil || IsNil(o.RuleCategory) {
		var ret OptimizationRulesAPIEntityFieldFilter
		return ret
	}
	return *o.RuleCategory
}

// GetRuleCategoryOk returns a tuple with the RuleCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIOptimizationRuleFilter) GetRuleCategoryOk() (*OptimizationRulesAPIEntityFieldFilter, bool) {
	if o == nil || IsNil(o.RuleCategory) {
		return nil, false
	}
	return o.RuleCategory, true
}

// HasRuleCategory returns a boolean if a field has been set.
func (o *OptimizationRulesAPIOptimizationRuleFilter) HasRuleCategory() bool {
	if o != nil && !IsNil(o.RuleCategory) {
		return true
	}

	return false
}

// SetRuleCategory gets a reference to the given OptimizationRulesAPIEntityFieldFilter and assigns it to the RuleCategory field.
func (o *OptimizationRulesAPIOptimizationRuleFilter) SetRuleCategory(v OptimizationRulesAPIEntityFieldFilter) {
	o.RuleCategory = &v
}

// GetRuleSubCategory returns the RuleSubCategory field value if set, zero value otherwise.
func (o *OptimizationRulesAPIOptimizationRuleFilter) GetRuleSubCategory() OptimizationRulesAPIEntityFieldFilter {
	if o == nil || IsNil(o.RuleSubCategory) {
		var ret OptimizationRulesAPIEntityFieldFilter
		return ret
	}
	return *o.RuleSubCategory
}

// GetRuleSubCategoryOk returns a tuple with the RuleSubCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIOptimizationRuleFilter) GetRuleSubCategoryOk() (*OptimizationRulesAPIEntityFieldFilter, bool) {
	if o == nil || IsNil(o.RuleSubCategory) {
		return nil, false
	}
	return o.RuleSubCategory, true
}

// HasRuleSubCategory returns a boolean if a field has been set.
func (o *OptimizationRulesAPIOptimizationRuleFilter) HasRuleSubCategory() bool {
	if o != nil && !IsNil(o.RuleSubCategory) {
		return true
	}

	return false
}

// SetRuleSubCategory gets a reference to the given OptimizationRulesAPIEntityFieldFilter and assigns it to the RuleSubCategory field.
func (o *OptimizationRulesAPIOptimizationRuleFilter) SetRuleSubCategory(v OptimizationRulesAPIEntityFieldFilter) {
	o.RuleSubCategory = &v
}

// GetOptimizationRuleId returns the OptimizationRuleId field value if set, zero value otherwise.
func (o *OptimizationRulesAPIOptimizationRuleFilter) GetOptimizationRuleId() OptimizationRulesAPIEntityFieldFilter {
	if o == nil || IsNil(o.OptimizationRuleId) {
		var ret OptimizationRulesAPIEntityFieldFilter
		return ret
	}
	return *o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIOptimizationRuleFilter) GetOptimizationRuleIdOk() (*OptimizationRulesAPIEntityFieldFilter, bool) {
	if o == nil || IsNil(o.OptimizationRuleId) {
		return nil, false
	}
	return o.OptimizationRuleId, true
}

// HasOptimizationRuleId returns a boolean if a field has been set.
func (o *OptimizationRulesAPIOptimizationRuleFilter) HasOptimizationRuleId() bool {
	if o != nil && !IsNil(o.OptimizationRuleId) {
		return true
	}

	return false
}

// SetOptimizationRuleId gets a reference to the given OptimizationRulesAPIEntityFieldFilter and assigns it to the OptimizationRuleId field.
func (o *OptimizationRulesAPIOptimizationRuleFilter) SetOptimizationRuleId(v OptimizationRulesAPIEntityFieldFilter) {
	o.OptimizationRuleId = &v
}

func (o OptimizationRulesAPIOptimizationRuleFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RuleCategory) {
		toSerialize["ruleCategory"] = o.RuleCategory
	}
	if !IsNil(o.RuleSubCategory) {
		toSerialize["ruleSubCategory"] = o.RuleSubCategory
	}
	if !IsNil(o.OptimizationRuleId) {
		toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIOptimizationRuleFilter struct {
	value *OptimizationRulesAPIOptimizationRuleFilter
	isSet bool
}

func (v NullableOptimizationRulesAPIOptimizationRuleFilter) Get() *OptimizationRulesAPIOptimizationRuleFilter {
	return v.value
}

func (v *NullableOptimizationRulesAPIOptimizationRuleFilter) Set(val *OptimizationRulesAPIOptimizationRuleFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIOptimizationRuleFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIOptimizationRuleFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIOptimizationRuleFilter(val *OptimizationRulesAPIOptimizationRuleFilter) *NullableOptimizationRulesAPIOptimizationRuleFilter {
	return &NullableOptimizationRulesAPIOptimizationRuleFilter{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIOptimizationRuleFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIOptimizationRuleFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
