package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateBudgetRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBudgetRulesResponse{}

// CreateBudgetRulesResponse struct for CreateBudgetRulesResponse
type CreateBudgetRulesResponse struct {
	Responses []BudgetRuleResponse `json:"responses,omitempty"`
}

// NewCreateBudgetRulesResponse instantiates a new CreateBudgetRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBudgetRulesResponse() *CreateBudgetRulesResponse {
	this := CreateBudgetRulesResponse{}
	return &this
}

// NewCreateBudgetRulesResponseWithDefaults instantiates a new CreateBudgetRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBudgetRulesResponseWithDefaults() *CreateBudgetRulesResponse {
	this := CreateBudgetRulesResponse{}
	return &this
}

// GetResponses returns the Responses field value if set, zero value otherwise.
func (o *CreateBudgetRulesResponse) GetResponses() []BudgetRuleResponse {
	if o == nil || IsNil(o.Responses) {
		var ret []BudgetRuleResponse
		return ret
	}
	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBudgetRulesResponse) GetResponsesOk() ([]BudgetRuleResponse, bool) {
	if o == nil || IsNil(o.Responses) {
		return nil, false
	}
	return o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *CreateBudgetRulesResponse) HasResponses() bool {
	if o != nil && !IsNil(o.Responses) {
		return true
	}

	return false
}

// SetResponses gets a reference to the given []BudgetRuleResponse and assigns it to the Responses field.
func (o *CreateBudgetRulesResponse) SetResponses(v []BudgetRuleResponse) {
	o.Responses = v
}

func (o CreateBudgetRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Responses) {
		toSerialize["responses"] = o.Responses
	}
	return toSerialize, nil
}

type NullableCreateBudgetRulesResponse struct {
	value *CreateBudgetRulesResponse
	isSet bool
}

func (v NullableCreateBudgetRulesResponse) Get() *CreateBudgetRulesResponse {
	return v.value
}

func (v *NullableCreateBudgetRulesResponse) Set(val *CreateBudgetRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBudgetRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBudgetRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBudgetRulesResponse(val *CreateBudgetRulesResponse) *NullableCreateBudgetRulesResponse {
	return &NullableCreateBudgetRulesResponse{value: val, isSet: true}
}

func (v NullableCreateBudgetRulesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateBudgetRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
