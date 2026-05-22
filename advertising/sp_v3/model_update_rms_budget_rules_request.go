package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateRMSBudgetRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRMSBudgetRulesRequest{}

// UpdateRMSBudgetRulesRequest Request object for updating budget rule for RMS campaign
type UpdateRMSBudgetRulesRequest struct {
	// A list of budget rule details.
	BudgetRulesDetails []RMSBudgetRule `json:"budgetRulesDetails,omitempty"`
}

// NewUpdateRMSBudgetRulesRequest instantiates a new UpdateRMSBudgetRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRMSBudgetRulesRequest() *UpdateRMSBudgetRulesRequest {
	this := UpdateRMSBudgetRulesRequest{}
	return &this
}

// NewUpdateRMSBudgetRulesRequestWithDefaults instantiates a new UpdateRMSBudgetRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRMSBudgetRulesRequestWithDefaults() *UpdateRMSBudgetRulesRequest {
	this := UpdateRMSBudgetRulesRequest{}
	return &this
}

// GetBudgetRulesDetails returns the BudgetRulesDetails field value if set, zero value otherwise.
func (o *UpdateRMSBudgetRulesRequest) GetBudgetRulesDetails() []RMSBudgetRule {
	if o == nil || IsNil(o.BudgetRulesDetails) {
		var ret []RMSBudgetRule
		return ret
	}
	return o.BudgetRulesDetails
}

// GetBudgetRulesDetailsOk returns a tuple with the BudgetRulesDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRMSBudgetRulesRequest) GetBudgetRulesDetailsOk() ([]RMSBudgetRule, bool) {
	if o == nil || IsNil(o.BudgetRulesDetails) {
		return nil, false
	}
	return o.BudgetRulesDetails, true
}

// HasBudgetRulesDetails returns a boolean if a field has been set.
func (o *UpdateRMSBudgetRulesRequest) HasBudgetRulesDetails() bool {
	if o != nil && !IsNil(o.BudgetRulesDetails) {
		return true
	}

	return false
}

// SetBudgetRulesDetails gets a reference to the given []RMSBudgetRule and assigns it to the BudgetRulesDetails field.
func (o *UpdateRMSBudgetRulesRequest) SetBudgetRulesDetails(v []RMSBudgetRule) {
	o.BudgetRulesDetails = v
}

func (o UpdateRMSBudgetRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetRulesDetails) {
		toSerialize["budgetRulesDetails"] = o.BudgetRulesDetails
	}
	return toSerialize, nil
}

type NullableUpdateRMSBudgetRulesRequest struct {
	value *UpdateRMSBudgetRulesRequest
	isSet bool
}

func (v NullableUpdateRMSBudgetRulesRequest) Get() *UpdateRMSBudgetRulesRequest {
	return v.value
}

func (v *NullableUpdateRMSBudgetRulesRequest) Set(val *UpdateRMSBudgetRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRMSBudgetRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRMSBudgetRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRMSBudgetRulesRequest(val *UpdateRMSBudgetRulesRequest) *NullableUpdateRMSBudgetRulesRequest {
	return &NullableUpdateRMSBudgetRulesRequest{value: val, isSet: true}
}

func (v NullableUpdateRMSBudgetRulesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateRMSBudgetRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
