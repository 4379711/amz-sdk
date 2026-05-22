package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GetSDBudgetRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSDBudgetRuleResponse{}

// GetSDBudgetRuleResponse struct for GetSDBudgetRuleResponse
type GetSDBudgetRuleResponse struct {
	BudgetRule *SDBudgetRule `json:"budgetRule,omitempty"`
}

// NewGetSDBudgetRuleResponse instantiates a new GetSDBudgetRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSDBudgetRuleResponse() *GetSDBudgetRuleResponse {
	this := GetSDBudgetRuleResponse{}
	return &this
}

// NewGetSDBudgetRuleResponseWithDefaults instantiates a new GetSDBudgetRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSDBudgetRuleResponseWithDefaults() *GetSDBudgetRuleResponse {
	this := GetSDBudgetRuleResponse{}
	return &this
}

// GetBudgetRule returns the BudgetRule field value if set, zero value otherwise.
func (o *GetSDBudgetRuleResponse) GetBudgetRule() SDBudgetRule {
	if o == nil || IsNil(o.BudgetRule) {
		var ret SDBudgetRule
		return ret
	}
	return *o.BudgetRule
}

// GetBudgetRuleOk returns a tuple with the BudgetRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSDBudgetRuleResponse) GetBudgetRuleOk() (*SDBudgetRule, bool) {
	if o == nil || IsNil(o.BudgetRule) {
		return nil, false
	}
	return o.BudgetRule, true
}

// HasBudgetRule returns a boolean if a field has been set.
func (o *GetSDBudgetRuleResponse) HasBudgetRule() bool {
	if o != nil && !IsNil(o.BudgetRule) {
		return true
	}

	return false
}

// SetBudgetRule gets a reference to the given SDBudgetRule and assigns it to the BudgetRule field.
func (o *GetSDBudgetRuleResponse) SetBudgetRule(v SDBudgetRule) {
	o.BudgetRule = &v
}

func (o GetSDBudgetRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetRule) {
		toSerialize["budgetRule"] = o.BudgetRule
	}
	return toSerialize, nil
}

type NullableGetSDBudgetRuleResponse struct {
	value *GetSDBudgetRuleResponse
	isSet bool
}

func (v NullableGetSDBudgetRuleResponse) Get() *GetSDBudgetRuleResponse {
	return v.value
}

func (v *NullableGetSDBudgetRuleResponse) Set(val *GetSDBudgetRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSDBudgetRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSDBudgetRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSDBudgetRuleResponse(val *GetSDBudgetRuleResponse) *NullableGetSDBudgetRuleResponse {
	return &NullableGetSDBudgetRuleResponse{value: val, isSet: true}
}

func (v NullableGetSDBudgetRuleResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetSDBudgetRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
