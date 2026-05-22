package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the RMSListAssociatedBudgetRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RMSListAssociatedBudgetRulesResponse{}

// RMSListAssociatedBudgetRulesResponse struct for RMSListAssociatedBudgetRulesResponse
type RMSListAssociatedBudgetRulesResponse struct {
	// A list of associated budget rules.
	AssociatedRules []RMSBudgetRule `json:"associatedRules,omitempty"`
}

// NewRMSListAssociatedBudgetRulesResponse instantiates a new RMSListAssociatedBudgetRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRMSListAssociatedBudgetRulesResponse() *RMSListAssociatedBudgetRulesResponse {
	this := RMSListAssociatedBudgetRulesResponse{}
	return &this
}

// NewRMSListAssociatedBudgetRulesResponseWithDefaults instantiates a new RMSListAssociatedBudgetRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRMSListAssociatedBudgetRulesResponseWithDefaults() *RMSListAssociatedBudgetRulesResponse {
	this := RMSListAssociatedBudgetRulesResponse{}
	return &this
}

// GetAssociatedRules returns the AssociatedRules field value if set, zero value otherwise.
func (o *RMSListAssociatedBudgetRulesResponse) GetAssociatedRules() []RMSBudgetRule {
	if o == nil || IsNil(o.AssociatedRules) {
		var ret []RMSBudgetRule
		return ret
	}
	return o.AssociatedRules
}

// GetAssociatedRulesOk returns a tuple with the AssociatedRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RMSListAssociatedBudgetRulesResponse) GetAssociatedRulesOk() ([]RMSBudgetRule, bool) {
	if o == nil || IsNil(o.AssociatedRules) {
		return nil, false
	}
	return o.AssociatedRules, true
}

// HasAssociatedRules returns a boolean if a field has been set.
func (o *RMSListAssociatedBudgetRulesResponse) HasAssociatedRules() bool {
	if o != nil && !IsNil(o.AssociatedRules) {
		return true
	}

	return false
}

// SetAssociatedRules gets a reference to the given []RMSBudgetRule and assigns it to the AssociatedRules field.
func (o *RMSListAssociatedBudgetRulesResponse) SetAssociatedRules(v []RMSBudgetRule) {
	o.AssociatedRules = v
}

func (o RMSListAssociatedBudgetRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociatedRules) {
		toSerialize["associatedRules"] = o.AssociatedRules
	}
	return toSerialize, nil
}

type NullableRMSListAssociatedBudgetRulesResponse struct {
	value *RMSListAssociatedBudgetRulesResponse
	isSet bool
}

func (v NullableRMSListAssociatedBudgetRulesResponse) Get() *RMSListAssociatedBudgetRulesResponse {
	return v.value
}

func (v *NullableRMSListAssociatedBudgetRulesResponse) Set(val *RMSListAssociatedBudgetRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRMSListAssociatedBudgetRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRMSListAssociatedBudgetRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRMSListAssociatedBudgetRulesResponse(val *RMSListAssociatedBudgetRulesResponse) *NullableRMSListAssociatedBudgetRulesResponse {
	return &NullableRMSListAssociatedBudgetRulesResponse{value: val, isSet: true}
}

func (v NullableRMSListAssociatedBudgetRulesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRMSListAssociatedBudgetRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
