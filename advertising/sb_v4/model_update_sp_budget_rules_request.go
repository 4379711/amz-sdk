package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSPBudgetRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSPBudgetRulesRequest{}

// UpdateSPBudgetRulesRequest Request object for updating budget rule for SP campaign
type UpdateSPBudgetRulesRequest struct {
	// A list of budget rule details.
	BudgetRulesDetails []SPBudgetRule `json:"budgetRulesDetails,omitempty"`
}

// NewUpdateSPBudgetRulesRequest instantiates a new UpdateSPBudgetRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSPBudgetRulesRequest() *UpdateSPBudgetRulesRequest {
	this := UpdateSPBudgetRulesRequest{}
	return &this
}

// NewUpdateSPBudgetRulesRequestWithDefaults instantiates a new UpdateSPBudgetRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSPBudgetRulesRequestWithDefaults() *UpdateSPBudgetRulesRequest {
	this := UpdateSPBudgetRulesRequest{}
	return &this
}

// GetBudgetRulesDetails returns the BudgetRulesDetails field value if set, zero value otherwise.
func (o *UpdateSPBudgetRulesRequest) GetBudgetRulesDetails() []SPBudgetRule {
	if o == nil || IsNil(o.BudgetRulesDetails) {
		var ret []SPBudgetRule
		return ret
	}
	return o.BudgetRulesDetails
}

// GetBudgetRulesDetailsOk returns a tuple with the BudgetRulesDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSPBudgetRulesRequest) GetBudgetRulesDetailsOk() ([]SPBudgetRule, bool) {
	if o == nil || IsNil(o.BudgetRulesDetails) {
		return nil, false
	}
	return o.BudgetRulesDetails, true
}

// HasBudgetRulesDetails returns a boolean if a field has been set.
func (o *UpdateSPBudgetRulesRequest) HasBudgetRulesDetails() bool {
	if o != nil && !IsNil(o.BudgetRulesDetails) {
		return true
	}

	return false
}

// SetBudgetRulesDetails gets a reference to the given []SPBudgetRule and assigns it to the BudgetRulesDetails field.
func (o *UpdateSPBudgetRulesRequest) SetBudgetRulesDetails(v []SPBudgetRule) {
	o.BudgetRulesDetails = v
}

func (o UpdateSPBudgetRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetRulesDetails) {
		toSerialize["budgetRulesDetails"] = o.BudgetRulesDetails
	}
	return toSerialize, nil
}

type NullableUpdateSPBudgetRulesRequest struct {
	value *UpdateSPBudgetRulesRequest
	isSet bool
}

func (v NullableUpdateSPBudgetRulesRequest) Get() *UpdateSPBudgetRulesRequest {
	return v.value
}

func (v *NullableUpdateSPBudgetRulesRequest) Set(val *UpdateSPBudgetRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSPBudgetRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSPBudgetRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSPBudgetRulesRequest(val *UpdateSPBudgetRulesRequest) *NullableUpdateSPBudgetRulesRequest {
	return &NullableUpdateSPBudgetRulesRequest{value: val, isSet: true}
}

func (v NullableUpdateSPBudgetRulesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSPBudgetRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
