package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GetSPBudgetRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSPBudgetRuleResponse{}

// GetSPBudgetRuleResponse struct for GetSPBudgetRuleResponse
type GetSPBudgetRuleResponse struct {
	BudgetRule *SPBudgetRule `json:"budgetRule,omitempty"`
}

// NewGetSPBudgetRuleResponse instantiates a new GetSPBudgetRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSPBudgetRuleResponse() *GetSPBudgetRuleResponse {
	this := GetSPBudgetRuleResponse{}
	return &this
}

// NewGetSPBudgetRuleResponseWithDefaults instantiates a new GetSPBudgetRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSPBudgetRuleResponseWithDefaults() *GetSPBudgetRuleResponse {
	this := GetSPBudgetRuleResponse{}
	return &this
}

// GetBudgetRule returns the BudgetRule field value if set, zero value otherwise.
func (o *GetSPBudgetRuleResponse) GetBudgetRule() SPBudgetRule {
	if o == nil || IsNil(o.BudgetRule) {
		var ret SPBudgetRule
		return ret
	}
	return *o.BudgetRule
}

// GetBudgetRuleOk returns a tuple with the BudgetRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSPBudgetRuleResponse) GetBudgetRuleOk() (*SPBudgetRule, bool) {
	if o == nil || IsNil(o.BudgetRule) {
		return nil, false
	}
	return o.BudgetRule, true
}

// HasBudgetRule returns a boolean if a field has been set.
func (o *GetSPBudgetRuleResponse) HasBudgetRule() bool {
	if o != nil && !IsNil(o.BudgetRule) {
		return true
	}

	return false
}

// SetBudgetRule gets a reference to the given SPBudgetRule and assigns it to the BudgetRule field.
func (o *GetSPBudgetRuleResponse) SetBudgetRule(v SPBudgetRule) {
	o.BudgetRule = &v
}

func (o GetSPBudgetRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetRule) {
		toSerialize["budgetRule"] = o.BudgetRule
	}
	return toSerialize, nil
}

type NullableGetSPBudgetRuleResponse struct {
	value *GetSPBudgetRuleResponse
	isSet bool
}

func (v NullableGetSPBudgetRuleResponse) Get() *GetSPBudgetRuleResponse {
	return v.value
}

func (v *NullableGetSPBudgetRuleResponse) Set(val *GetSPBudgetRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSPBudgetRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSPBudgetRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSPBudgetRuleResponse(val *GetSPBudgetRuleResponse) *NullableGetSPBudgetRuleResponse {
	return &NullableGetSPBudgetRuleResponse{value: val, isSet: true}
}

func (v NullableGetSPBudgetRuleResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetSPBudgetRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
