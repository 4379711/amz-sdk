package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSBBudgetRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSBBudgetRulesRequest{}

// CreateSBBudgetRulesRequest struct for CreateSBBudgetRulesRequest
type CreateSBBudgetRulesRequest struct {
	// A list of budget rule details.
	BudgetRulesDetails []SBBudgetRuleDetails `json:"budgetRulesDetails,omitempty"`
}

// NewCreateSBBudgetRulesRequest instantiates a new CreateSBBudgetRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSBBudgetRulesRequest() *CreateSBBudgetRulesRequest {
	this := CreateSBBudgetRulesRequest{}
	return &this
}

// NewCreateSBBudgetRulesRequestWithDefaults instantiates a new CreateSBBudgetRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSBBudgetRulesRequestWithDefaults() *CreateSBBudgetRulesRequest {
	this := CreateSBBudgetRulesRequest{}
	return &this
}

// GetBudgetRulesDetails returns the BudgetRulesDetails field value if set, zero value otherwise.
func (o *CreateSBBudgetRulesRequest) GetBudgetRulesDetails() []SBBudgetRuleDetails {
	if o == nil || IsNil(o.BudgetRulesDetails) {
		var ret []SBBudgetRuleDetails
		return ret
	}
	return o.BudgetRulesDetails
}

// GetBudgetRulesDetailsOk returns a tuple with the BudgetRulesDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSBBudgetRulesRequest) GetBudgetRulesDetailsOk() ([]SBBudgetRuleDetails, bool) {
	if o == nil || IsNil(o.BudgetRulesDetails) {
		return nil, false
	}
	return o.BudgetRulesDetails, true
}

// HasBudgetRulesDetails returns a boolean if a field has been set.
func (o *CreateSBBudgetRulesRequest) HasBudgetRulesDetails() bool {
	if o != nil && !IsNil(o.BudgetRulesDetails) {
		return true
	}

	return false
}

// SetBudgetRulesDetails gets a reference to the given []SBBudgetRuleDetails and assigns it to the BudgetRulesDetails field.
func (o *CreateSBBudgetRulesRequest) SetBudgetRulesDetails(v []SBBudgetRuleDetails) {
	o.BudgetRulesDetails = v
}

func (o CreateSBBudgetRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetRulesDetails) {
		toSerialize["budgetRulesDetails"] = o.BudgetRulesDetails
	}
	return toSerialize, nil
}

type NullableCreateSBBudgetRulesRequest struct {
	value *CreateSBBudgetRulesRequest
	isSet bool
}

func (v NullableCreateSBBudgetRulesRequest) Get() *CreateSBBudgetRulesRequest {
	return v.value
}

func (v *NullableCreateSBBudgetRulesRequest) Set(val *CreateSBBudgetRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSBBudgetRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSBBudgetRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSBBudgetRulesRequest(val *CreateSBBudgetRulesRequest) *NullableCreateSBBudgetRulesRequest {
	return &NullableCreateSBBudgetRulesRequest{value: val, isSet: true}
}

func (v NullableCreateSBBudgetRulesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSBBudgetRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
