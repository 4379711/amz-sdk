package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSDBudgetRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSDBudgetRulesRequest{}

// UpdateSDBudgetRulesRequest Request object for updating budget rule for SD campaign
type UpdateSDBudgetRulesRequest struct {
	// A list of budget rule details.
	BudgetRulesDetails []SDBudgetRule `json:"budgetRulesDetails,omitempty"`
}

// NewUpdateSDBudgetRulesRequest instantiates a new UpdateSDBudgetRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSDBudgetRulesRequest() *UpdateSDBudgetRulesRequest {
	this := UpdateSDBudgetRulesRequest{}
	return &this
}

// NewUpdateSDBudgetRulesRequestWithDefaults instantiates a new UpdateSDBudgetRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSDBudgetRulesRequestWithDefaults() *UpdateSDBudgetRulesRequest {
	this := UpdateSDBudgetRulesRequest{}
	return &this
}

// GetBudgetRulesDetails returns the BudgetRulesDetails field value if set, zero value otherwise.
func (o *UpdateSDBudgetRulesRequest) GetBudgetRulesDetails() []SDBudgetRule {
	if o == nil || IsNil(o.BudgetRulesDetails) {
		var ret []SDBudgetRule
		return ret
	}
	return o.BudgetRulesDetails
}

// GetBudgetRulesDetailsOk returns a tuple with the BudgetRulesDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSDBudgetRulesRequest) GetBudgetRulesDetailsOk() ([]SDBudgetRule, bool) {
	if o == nil || IsNil(o.BudgetRulesDetails) {
		return nil, false
	}
	return o.BudgetRulesDetails, true
}

// HasBudgetRulesDetails returns a boolean if a field has been set.
func (o *UpdateSDBudgetRulesRequest) HasBudgetRulesDetails() bool {
	if o != nil && !IsNil(o.BudgetRulesDetails) {
		return true
	}

	return false
}

// SetBudgetRulesDetails gets a reference to the given []SDBudgetRule and assigns it to the BudgetRulesDetails field.
func (o *UpdateSDBudgetRulesRequest) SetBudgetRulesDetails(v []SDBudgetRule) {
	o.BudgetRulesDetails = v
}

func (o UpdateSDBudgetRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetRulesDetails) {
		toSerialize["budgetRulesDetails"] = o.BudgetRulesDetails
	}
	return toSerialize, nil
}

type NullableUpdateSDBudgetRulesRequest struct {
	value *UpdateSDBudgetRulesRequest
	isSet bool
}

func (v NullableUpdateSDBudgetRulesRequest) Get() *UpdateSDBudgetRulesRequest {
	return v.value
}

func (v *NullableUpdateSDBudgetRulesRequest) Set(val *UpdateSDBudgetRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSDBudgetRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSDBudgetRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSDBudgetRulesRequest(val *UpdateSDBudgetRulesRequest) *NullableUpdateSDBudgetRulesRequest {
	return &NullableUpdateSDBudgetRulesRequest{value: val, isSet: true}
}

func (v NullableUpdateSDBudgetRulesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSDBudgetRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
