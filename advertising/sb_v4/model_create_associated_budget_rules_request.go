package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateAssociatedBudgetRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAssociatedBudgetRulesRequest{}

// CreateAssociatedBudgetRulesRequest struct for CreateAssociatedBudgetRulesRequest
type CreateAssociatedBudgetRulesRequest struct {
	// A list of budget rule identifiers.
	BudgetRuleIds []string `json:"budgetRuleIds,omitempty"`
}

// NewCreateAssociatedBudgetRulesRequest instantiates a new CreateAssociatedBudgetRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAssociatedBudgetRulesRequest() *CreateAssociatedBudgetRulesRequest {
	this := CreateAssociatedBudgetRulesRequest{}
	return &this
}

// NewCreateAssociatedBudgetRulesRequestWithDefaults instantiates a new CreateAssociatedBudgetRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAssociatedBudgetRulesRequestWithDefaults() *CreateAssociatedBudgetRulesRequest {
	this := CreateAssociatedBudgetRulesRequest{}
	return &this
}

// GetBudgetRuleIds returns the BudgetRuleIds field value if set, zero value otherwise.
func (o *CreateAssociatedBudgetRulesRequest) GetBudgetRuleIds() []string {
	if o == nil || IsNil(o.BudgetRuleIds) {
		var ret []string
		return ret
	}
	return o.BudgetRuleIds
}

// GetBudgetRuleIdsOk returns a tuple with the BudgetRuleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssociatedBudgetRulesRequest) GetBudgetRuleIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.BudgetRuleIds) {
		return nil, false
	}
	return o.BudgetRuleIds, true
}

// HasBudgetRuleIds returns a boolean if a field has been set.
func (o *CreateAssociatedBudgetRulesRequest) HasBudgetRuleIds() bool {
	if o != nil && !IsNil(o.BudgetRuleIds) {
		return true
	}

	return false
}

// SetBudgetRuleIds gets a reference to the given []string and assigns it to the BudgetRuleIds field.
func (o *CreateAssociatedBudgetRulesRequest) SetBudgetRuleIds(v []string) {
	o.BudgetRuleIds = v
}

func (o CreateAssociatedBudgetRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetRuleIds) {
		toSerialize["budgetRuleIds"] = o.BudgetRuleIds
	}
	return toSerialize, nil
}

type NullableCreateAssociatedBudgetRulesRequest struct {
	value *CreateAssociatedBudgetRulesRequest
	isSet bool
}

func (v NullableCreateAssociatedBudgetRulesRequest) Get() *CreateAssociatedBudgetRulesRequest {
	return v.value
}

func (v *NullableCreateAssociatedBudgetRulesRequest) Set(val *CreateAssociatedBudgetRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAssociatedBudgetRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAssociatedBudgetRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAssociatedBudgetRulesRequest(val *CreateAssociatedBudgetRulesRequest) *NullableCreateAssociatedBudgetRulesRequest {
	return &NullableCreateAssociatedBudgetRulesRequest{value: val, isSet: true}
}

func (v NullableCreateAssociatedBudgetRulesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateAssociatedBudgetRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
