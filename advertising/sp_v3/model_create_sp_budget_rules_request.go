package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSPBudgetRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSPBudgetRulesRequest{}

// CreateSPBudgetRulesRequest struct for CreateSPBudgetRulesRequest
type CreateSPBudgetRulesRequest struct {
	// A list of budget rule details.
	BudgetRulesDetails []SPBudgetRuleDetails `json:"budgetRulesDetails,omitempty"`
}

// NewCreateSPBudgetRulesRequest instantiates a new CreateSPBudgetRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSPBudgetRulesRequest() *CreateSPBudgetRulesRequest {
	this := CreateSPBudgetRulesRequest{}
	return &this
}

// NewCreateSPBudgetRulesRequestWithDefaults instantiates a new CreateSPBudgetRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSPBudgetRulesRequestWithDefaults() *CreateSPBudgetRulesRequest {
	this := CreateSPBudgetRulesRequest{}
	return &this
}

// GetBudgetRulesDetails returns the BudgetRulesDetails field value if set, zero value otherwise.
func (o *CreateSPBudgetRulesRequest) GetBudgetRulesDetails() []SPBudgetRuleDetails {
	if o == nil || IsNil(o.BudgetRulesDetails) {
		var ret []SPBudgetRuleDetails
		return ret
	}
	return o.BudgetRulesDetails
}

// GetBudgetRulesDetailsOk returns a tuple with the BudgetRulesDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSPBudgetRulesRequest) GetBudgetRulesDetailsOk() ([]SPBudgetRuleDetails, bool) {
	if o == nil || IsNil(o.BudgetRulesDetails) {
		return nil, false
	}
	return o.BudgetRulesDetails, true
}

// HasBudgetRulesDetails returns a boolean if a field has been set.
func (o *CreateSPBudgetRulesRequest) HasBudgetRulesDetails() bool {
	if o != nil && !IsNil(o.BudgetRulesDetails) {
		return true
	}

	return false
}

// SetBudgetRulesDetails gets a reference to the given []SPBudgetRuleDetails and assigns it to the BudgetRulesDetails field.
func (o *CreateSPBudgetRulesRequest) SetBudgetRulesDetails(v []SPBudgetRuleDetails) {
	o.BudgetRulesDetails = v
}

func (o CreateSPBudgetRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetRulesDetails) {
		toSerialize["budgetRulesDetails"] = o.BudgetRulesDetails
	}
	return toSerialize, nil
}

type NullableCreateSPBudgetRulesRequest struct {
	value *CreateSPBudgetRulesRequest
	isSet bool
}

func (v NullableCreateSPBudgetRulesRequest) Get() *CreateSPBudgetRulesRequest {
	return v.value
}

func (v *NullableCreateSPBudgetRulesRequest) Set(val *CreateSPBudgetRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSPBudgetRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSPBudgetRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSPBudgetRulesRequest(val *CreateSPBudgetRulesRequest) *NullableCreateSPBudgetRulesRequest {
	return &NullableCreateSPBudgetRulesRequest{value: val, isSet: true}
}

func (v NullableCreateSPBudgetRulesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSPBudgetRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
