package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GetRMSBudgetRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRMSBudgetRuleResponse{}

// GetRMSBudgetRuleResponse struct for GetRMSBudgetRuleResponse
type GetRMSBudgetRuleResponse struct {
	BudgetRule *RMSBudgetRule `json:"budgetRule,omitempty"`
}

// NewGetRMSBudgetRuleResponse instantiates a new GetRMSBudgetRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRMSBudgetRuleResponse() *GetRMSBudgetRuleResponse {
	this := GetRMSBudgetRuleResponse{}
	return &this
}

// NewGetRMSBudgetRuleResponseWithDefaults instantiates a new GetRMSBudgetRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRMSBudgetRuleResponseWithDefaults() *GetRMSBudgetRuleResponse {
	this := GetRMSBudgetRuleResponse{}
	return &this
}

// GetBudgetRule returns the BudgetRule field value if set, zero value otherwise.
func (o *GetRMSBudgetRuleResponse) GetBudgetRule() RMSBudgetRule {
	if o == nil || IsNil(o.BudgetRule) {
		var ret RMSBudgetRule
		return ret
	}
	return *o.BudgetRule
}

// GetBudgetRuleOk returns a tuple with the BudgetRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRMSBudgetRuleResponse) GetBudgetRuleOk() (*RMSBudgetRule, bool) {
	if o == nil || IsNil(o.BudgetRule) {
		return nil, false
	}
	return o.BudgetRule, true
}

// HasBudgetRule returns a boolean if a field has been set.
func (o *GetRMSBudgetRuleResponse) HasBudgetRule() bool {
	if o != nil && !IsNil(o.BudgetRule) {
		return true
	}

	return false
}

// SetBudgetRule gets a reference to the given RMSBudgetRule and assigns it to the BudgetRule field.
func (o *GetRMSBudgetRuleResponse) SetBudgetRule(v RMSBudgetRule) {
	o.BudgetRule = &v
}

func (o GetRMSBudgetRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetRule) {
		toSerialize["budgetRule"] = o.BudgetRule
	}
	return toSerialize, nil
}

type NullableGetRMSBudgetRuleResponse struct {
	value *GetRMSBudgetRuleResponse
	isSet bool
}

func (v NullableGetRMSBudgetRuleResponse) Get() *GetRMSBudgetRuleResponse {
	return v.value
}

func (v *NullableGetRMSBudgetRuleResponse) Set(val *GetRMSBudgetRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRMSBudgetRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRMSBudgetRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRMSBudgetRuleResponse(val *GetRMSBudgetRuleResponse) *NullableGetRMSBudgetRuleResponse {
	return &NullableGetRMSBudgetRuleResponse{value: val, isSet: true}
}

func (v NullableGetRMSBudgetRuleResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetRMSBudgetRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
