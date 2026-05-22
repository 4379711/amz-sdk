package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSDBudgetRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSDBudgetRulesRequest{}

// CreateSDBudgetRulesRequest struct for CreateSDBudgetRulesRequest
type CreateSDBudgetRulesRequest struct {
	// A list of budget rule details.
	BudgetRulesDetails []SDBudgetRuleDetails `json:"budgetRulesDetails,omitempty"`
}

// NewCreateSDBudgetRulesRequest instantiates a new CreateSDBudgetRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSDBudgetRulesRequest() *CreateSDBudgetRulesRequest {
	this := CreateSDBudgetRulesRequest{}
	return &this
}

// NewCreateSDBudgetRulesRequestWithDefaults instantiates a new CreateSDBudgetRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSDBudgetRulesRequestWithDefaults() *CreateSDBudgetRulesRequest {
	this := CreateSDBudgetRulesRequest{}
	return &this
}

// GetBudgetRulesDetails returns the BudgetRulesDetails field value if set, zero value otherwise.
func (o *CreateSDBudgetRulesRequest) GetBudgetRulesDetails() []SDBudgetRuleDetails {
	if o == nil || IsNil(o.BudgetRulesDetails) {
		var ret []SDBudgetRuleDetails
		return ret
	}
	return o.BudgetRulesDetails
}

// GetBudgetRulesDetailsOk returns a tuple with the BudgetRulesDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSDBudgetRulesRequest) GetBudgetRulesDetailsOk() ([]SDBudgetRuleDetails, bool) {
	if o == nil || IsNil(o.BudgetRulesDetails) {
		return nil, false
	}
	return o.BudgetRulesDetails, true
}

// HasBudgetRulesDetails returns a boolean if a field has been set.
func (o *CreateSDBudgetRulesRequest) HasBudgetRulesDetails() bool {
	if o != nil && !IsNil(o.BudgetRulesDetails) {
		return true
	}

	return false
}

// SetBudgetRulesDetails gets a reference to the given []SDBudgetRuleDetails and assigns it to the BudgetRulesDetails field.
func (o *CreateSDBudgetRulesRequest) SetBudgetRulesDetails(v []SDBudgetRuleDetails) {
	o.BudgetRulesDetails = v
}

func (o CreateSDBudgetRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetRulesDetails) {
		toSerialize["budgetRulesDetails"] = o.BudgetRulesDetails
	}
	return toSerialize, nil
}

type NullableCreateSDBudgetRulesRequest struct {
	value *CreateSDBudgetRulesRequest
	isSet bool
}

func (v NullableCreateSDBudgetRulesRequest) Get() *CreateSDBudgetRulesRequest {
	return v.value
}

func (v *NullableCreateSDBudgetRulesRequest) Set(val *CreateSDBudgetRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSDBudgetRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSDBudgetRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSDBudgetRulesRequest(val *CreateSDBudgetRulesRequest) *NullableCreateSDBudgetRulesRequest {
	return &NullableCreateSDBudgetRulesRequest{value: val, isSet: true}
}

func (v NullableCreateSDBudgetRulesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSDBudgetRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
