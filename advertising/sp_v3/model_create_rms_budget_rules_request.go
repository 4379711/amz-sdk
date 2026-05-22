package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateRMSBudgetRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRMSBudgetRulesRequest{}

// CreateRMSBudgetRulesRequest struct for CreateRMSBudgetRulesRequest
type CreateRMSBudgetRulesRequest struct {
	// A list of budget rule details.
	BudgetRulesDetails []RMSBudgetRuleDetails `json:"budgetRulesDetails,omitempty"`
}

// NewCreateRMSBudgetRulesRequest instantiates a new CreateRMSBudgetRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRMSBudgetRulesRequest() *CreateRMSBudgetRulesRequest {
	this := CreateRMSBudgetRulesRequest{}
	return &this
}

// NewCreateRMSBudgetRulesRequestWithDefaults instantiates a new CreateRMSBudgetRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRMSBudgetRulesRequestWithDefaults() *CreateRMSBudgetRulesRequest {
	this := CreateRMSBudgetRulesRequest{}
	return &this
}

// GetBudgetRulesDetails returns the BudgetRulesDetails field value if set, zero value otherwise.
func (o *CreateRMSBudgetRulesRequest) GetBudgetRulesDetails() []RMSBudgetRuleDetails {
	if o == nil || IsNil(o.BudgetRulesDetails) {
		var ret []RMSBudgetRuleDetails
		return ret
	}
	return o.BudgetRulesDetails
}

// GetBudgetRulesDetailsOk returns a tuple with the BudgetRulesDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRMSBudgetRulesRequest) GetBudgetRulesDetailsOk() ([]RMSBudgetRuleDetails, bool) {
	if o == nil || IsNil(o.BudgetRulesDetails) {
		return nil, false
	}
	return o.BudgetRulesDetails, true
}

// HasBudgetRulesDetails returns a boolean if a field has been set.
func (o *CreateRMSBudgetRulesRequest) HasBudgetRulesDetails() bool {
	if o != nil && !IsNil(o.BudgetRulesDetails) {
		return true
	}

	return false
}

// SetBudgetRulesDetails gets a reference to the given []RMSBudgetRuleDetails and assigns it to the BudgetRulesDetails field.
func (o *CreateRMSBudgetRulesRequest) SetBudgetRulesDetails(v []RMSBudgetRuleDetails) {
	o.BudgetRulesDetails = v
}

func (o CreateRMSBudgetRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetRulesDetails) {
		toSerialize["budgetRulesDetails"] = o.BudgetRulesDetails
	}
	return toSerialize, nil
}

type NullableCreateRMSBudgetRulesRequest struct {
	value *CreateRMSBudgetRulesRequest
	isSet bool
}

func (v NullableCreateRMSBudgetRulesRequest) Get() *CreateRMSBudgetRulesRequest {
	return v.value
}

func (v *NullableCreateRMSBudgetRulesRequest) Set(val *CreateRMSBudgetRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRMSBudgetRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRMSBudgetRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRMSBudgetRulesRequest(val *CreateRMSBudgetRulesRequest) *NullableCreateRMSBudgetRulesRequest {
	return &NullableCreateRMSBudgetRulesRequest{value: val, isSet: true}
}

func (v NullableCreateRMSBudgetRulesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateRMSBudgetRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
