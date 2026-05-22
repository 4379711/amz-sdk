package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateBudgetRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateBudgetRulesResponse{}

// UpdateBudgetRulesResponse struct for UpdateBudgetRulesResponse
type UpdateBudgetRulesResponse struct {
	Responses []BudgetRuleResponse `json:"responses,omitempty"`
}

// NewUpdateBudgetRulesResponse instantiates a new UpdateBudgetRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateBudgetRulesResponse() *UpdateBudgetRulesResponse {
	this := UpdateBudgetRulesResponse{}
	return &this
}

// NewUpdateBudgetRulesResponseWithDefaults instantiates a new UpdateBudgetRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateBudgetRulesResponseWithDefaults() *UpdateBudgetRulesResponse {
	this := UpdateBudgetRulesResponse{}
	return &this
}

// GetResponses returns the Responses field value if set, zero value otherwise.
func (o *UpdateBudgetRulesResponse) GetResponses() []BudgetRuleResponse {
	if o == nil || IsNil(o.Responses) {
		var ret []BudgetRuleResponse
		return ret
	}
	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBudgetRulesResponse) GetResponsesOk() ([]BudgetRuleResponse, bool) {
	if o == nil || IsNil(o.Responses) {
		return nil, false
	}
	return o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *UpdateBudgetRulesResponse) HasResponses() bool {
	if o != nil && !IsNil(o.Responses) {
		return true
	}

	return false
}

// SetResponses gets a reference to the given []BudgetRuleResponse and assigns it to the Responses field.
func (o *UpdateBudgetRulesResponse) SetResponses(v []BudgetRuleResponse) {
	o.Responses = v
}

func (o UpdateBudgetRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Responses) {
		toSerialize["responses"] = o.Responses
	}
	return toSerialize, nil
}

type NullableUpdateBudgetRulesResponse struct {
	value *UpdateBudgetRulesResponse
	isSet bool
}

func (v NullableUpdateBudgetRulesResponse) Get() *UpdateBudgetRulesResponse {
	return v.value
}

func (v *NullableUpdateBudgetRulesResponse) Set(val *UpdateBudgetRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateBudgetRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateBudgetRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateBudgetRulesResponse(val *UpdateBudgetRulesResponse) *NullableUpdateBudgetRulesResponse {
	return &NullableUpdateBudgetRulesResponse{value: val, isSet: true}
}

func (v NullableUpdateBudgetRulesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateBudgetRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
