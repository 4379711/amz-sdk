package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateAssociatedBudgetRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAssociatedBudgetRulesResponse{}

// CreateAssociatedBudgetRulesResponse struct for CreateAssociatedBudgetRulesResponse
type CreateAssociatedBudgetRulesResponse struct {
	Responses []AssociatedBudgetRuleResponse `json:"responses,omitempty"`
}

// NewCreateAssociatedBudgetRulesResponse instantiates a new CreateAssociatedBudgetRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAssociatedBudgetRulesResponse() *CreateAssociatedBudgetRulesResponse {
	this := CreateAssociatedBudgetRulesResponse{}
	return &this
}

// NewCreateAssociatedBudgetRulesResponseWithDefaults instantiates a new CreateAssociatedBudgetRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAssociatedBudgetRulesResponseWithDefaults() *CreateAssociatedBudgetRulesResponse {
	this := CreateAssociatedBudgetRulesResponse{}
	return &this
}

// GetResponses returns the Responses field value if set, zero value otherwise.
func (o *CreateAssociatedBudgetRulesResponse) GetResponses() []AssociatedBudgetRuleResponse {
	if o == nil || IsNil(o.Responses) {
		var ret []AssociatedBudgetRuleResponse
		return ret
	}
	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssociatedBudgetRulesResponse) GetResponsesOk() ([]AssociatedBudgetRuleResponse, bool) {
	if o == nil || IsNil(o.Responses) {
		return nil, false
	}
	return o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *CreateAssociatedBudgetRulesResponse) HasResponses() bool {
	if o != nil && !IsNil(o.Responses) {
		return true
	}

	return false
}

// SetResponses gets a reference to the given []AssociatedBudgetRuleResponse and assigns it to the Responses field.
func (o *CreateAssociatedBudgetRulesResponse) SetResponses(v []AssociatedBudgetRuleResponse) {
	o.Responses = v
}

func (o CreateAssociatedBudgetRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Responses) {
		toSerialize["responses"] = o.Responses
	}
	return toSerialize, nil
}

type NullableCreateAssociatedBudgetRulesResponse struct {
	value *CreateAssociatedBudgetRulesResponse
	isSet bool
}

func (v NullableCreateAssociatedBudgetRulesResponse) Get() *CreateAssociatedBudgetRulesResponse {
	return v.value
}

func (v *NullableCreateAssociatedBudgetRulesResponse) Set(val *CreateAssociatedBudgetRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAssociatedBudgetRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAssociatedBudgetRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAssociatedBudgetRulesResponse(val *CreateAssociatedBudgetRulesResponse) *NullableCreateAssociatedBudgetRulesResponse {
	return &NullableCreateAssociatedBudgetRulesResponse{value: val, isSet: true}
}

func (v NullableCreateAssociatedBudgetRulesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateAssociatedBudgetRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
