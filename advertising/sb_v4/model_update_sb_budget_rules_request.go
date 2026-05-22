package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSBBudgetRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSBBudgetRulesRequest{}

// UpdateSBBudgetRulesRequest struct for UpdateSBBudgetRulesRequest
type UpdateSBBudgetRulesRequest struct {
	// A list of budget rule details.
	BudgetRulesDetails []SBBudgetRule `json:"budgetRulesDetails,omitempty"`
}

// NewUpdateSBBudgetRulesRequest instantiates a new UpdateSBBudgetRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSBBudgetRulesRequest() *UpdateSBBudgetRulesRequest {
	this := UpdateSBBudgetRulesRequest{}
	return &this
}

// NewUpdateSBBudgetRulesRequestWithDefaults instantiates a new UpdateSBBudgetRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSBBudgetRulesRequestWithDefaults() *UpdateSBBudgetRulesRequest {
	this := UpdateSBBudgetRulesRequest{}
	return &this
}

// GetBudgetRulesDetails returns the BudgetRulesDetails field value if set, zero value otherwise.
func (o *UpdateSBBudgetRulesRequest) GetBudgetRulesDetails() []SBBudgetRule {
	if o == nil || IsNil(o.BudgetRulesDetails) {
		var ret []SBBudgetRule
		return ret
	}
	return o.BudgetRulesDetails
}

// GetBudgetRulesDetailsOk returns a tuple with the BudgetRulesDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSBBudgetRulesRequest) GetBudgetRulesDetailsOk() ([]SBBudgetRule, bool) {
	if o == nil || IsNil(o.BudgetRulesDetails) {
		return nil, false
	}
	return o.BudgetRulesDetails, true
}

// HasBudgetRulesDetails returns a boolean if a field has been set.
func (o *UpdateSBBudgetRulesRequest) HasBudgetRulesDetails() bool {
	if o != nil && !IsNil(o.BudgetRulesDetails) {
		return true
	}

	return false
}

// SetBudgetRulesDetails gets a reference to the given []SBBudgetRule and assigns it to the BudgetRulesDetails field.
func (o *UpdateSBBudgetRulesRequest) SetBudgetRulesDetails(v []SBBudgetRule) {
	o.BudgetRulesDetails = v
}

func (o UpdateSBBudgetRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetRulesDetails) {
		toSerialize["budgetRulesDetails"] = o.BudgetRulesDetails
	}
	return toSerialize, nil
}

type NullableUpdateSBBudgetRulesRequest struct {
	value *UpdateSBBudgetRulesRequest
	isSet bool
}

func (v NullableUpdateSBBudgetRulesRequest) Get() *UpdateSBBudgetRulesRequest {
	return v.value
}

func (v *NullableUpdateSBBudgetRulesRequest) Set(val *UpdateSBBudgetRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSBBudgetRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSBBudgetRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSBBudgetRulesRequest(val *UpdateSBBudgetRulesRequest) *NullableUpdateSBBudgetRulesRequest {
	return &NullableUpdateSBBudgetRulesRequest{value: val, isSet: true}
}

func (v NullableUpdateSBBudgetRulesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSBBudgetRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
