package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GetSBBudgetRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSBBudgetRuleResponse{}

// GetSBBudgetRuleResponse struct for GetSBBudgetRuleResponse
type GetSBBudgetRuleResponse struct {
	BudgetRule *SBBudgetRule `json:"budgetRule,omitempty"`
}

// NewGetSBBudgetRuleResponse instantiates a new GetSBBudgetRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSBBudgetRuleResponse() *GetSBBudgetRuleResponse {
	this := GetSBBudgetRuleResponse{}
	return &this
}

// NewGetSBBudgetRuleResponseWithDefaults instantiates a new GetSBBudgetRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSBBudgetRuleResponseWithDefaults() *GetSBBudgetRuleResponse {
	this := GetSBBudgetRuleResponse{}
	return &this
}

// GetBudgetRule returns the BudgetRule field value if set, zero value otherwise.
func (o *GetSBBudgetRuleResponse) GetBudgetRule() SBBudgetRule {
	if o == nil || IsNil(o.BudgetRule) {
		var ret SBBudgetRule
		return ret
	}
	return *o.BudgetRule
}

// GetBudgetRuleOk returns a tuple with the BudgetRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSBBudgetRuleResponse) GetBudgetRuleOk() (*SBBudgetRule, bool) {
	if o == nil || IsNil(o.BudgetRule) {
		return nil, false
	}
	return o.BudgetRule, true
}

// HasBudgetRule returns a boolean if a field has been set.
func (o *GetSBBudgetRuleResponse) HasBudgetRule() bool {
	if o != nil && !IsNil(o.BudgetRule) {
		return true
	}

	return false
}

// SetBudgetRule gets a reference to the given SBBudgetRule and assigns it to the BudgetRule field.
func (o *GetSBBudgetRuleResponse) SetBudgetRule(v SBBudgetRule) {
	o.BudgetRule = &v
}

func (o GetSBBudgetRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetRule) {
		toSerialize["budgetRule"] = o.BudgetRule
	}
	return toSerialize, nil
}

type NullableGetSBBudgetRuleResponse struct {
	value *GetSBBudgetRuleResponse
	isSet bool
}

func (v NullableGetSBBudgetRuleResponse) Get() *GetSBBudgetRuleResponse {
	return v.value
}

func (v *NullableGetSBBudgetRuleResponse) Set(val *GetSBBudgetRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSBBudgetRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSBBudgetRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSBBudgetRuleResponse(val *GetSBBudgetRuleResponse) *NullableGetSBBudgetRuleResponse {
	return &NullableGetSBBudgetRuleResponse{value: val, isSet: true}
}

func (v NullableGetSBBudgetRuleResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetSBBudgetRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
