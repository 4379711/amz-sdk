package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SDListAssociatedBudgetRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDListAssociatedBudgetRulesResponse{}

// SDListAssociatedBudgetRulesResponse struct for SDListAssociatedBudgetRulesResponse
type SDListAssociatedBudgetRulesResponse struct {
	// A list of associated budget rules.
	AssociatedRules []SDBudgetRule `json:"associatedRules,omitempty"`
}

// NewSDListAssociatedBudgetRulesResponse instantiates a new SDListAssociatedBudgetRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDListAssociatedBudgetRulesResponse() *SDListAssociatedBudgetRulesResponse {
	this := SDListAssociatedBudgetRulesResponse{}
	return &this
}

// NewSDListAssociatedBudgetRulesResponseWithDefaults instantiates a new SDListAssociatedBudgetRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDListAssociatedBudgetRulesResponseWithDefaults() *SDListAssociatedBudgetRulesResponse {
	this := SDListAssociatedBudgetRulesResponse{}
	return &this
}

// GetAssociatedRules returns the AssociatedRules field value if set, zero value otherwise.
func (o *SDListAssociatedBudgetRulesResponse) GetAssociatedRules() []SDBudgetRule {
	if o == nil || IsNil(o.AssociatedRules) {
		var ret []SDBudgetRule
		return ret
	}
	return o.AssociatedRules
}

// GetAssociatedRulesOk returns a tuple with the AssociatedRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDListAssociatedBudgetRulesResponse) GetAssociatedRulesOk() ([]SDBudgetRule, bool) {
	if o == nil || IsNil(o.AssociatedRules) {
		return nil, false
	}
	return o.AssociatedRules, true
}

// HasAssociatedRules returns a boolean if a field has been set.
func (o *SDListAssociatedBudgetRulesResponse) HasAssociatedRules() bool {
	if o != nil && !IsNil(o.AssociatedRules) {
		return true
	}

	return false
}

// SetAssociatedRules gets a reference to the given []SDBudgetRule and assigns it to the AssociatedRules field.
func (o *SDListAssociatedBudgetRulesResponse) SetAssociatedRules(v []SDBudgetRule) {
	o.AssociatedRules = v
}

func (o SDListAssociatedBudgetRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociatedRules) {
		toSerialize["associatedRules"] = o.AssociatedRules
	}
	return toSerialize, nil
}

type NullableSDListAssociatedBudgetRulesResponse struct {
	value *SDListAssociatedBudgetRulesResponse
	isSet bool
}

func (v NullableSDListAssociatedBudgetRulesResponse) Get() *SDListAssociatedBudgetRulesResponse {
	return v.value
}

func (v *NullableSDListAssociatedBudgetRulesResponse) Set(val *SDListAssociatedBudgetRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSDListAssociatedBudgetRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSDListAssociatedBudgetRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDListAssociatedBudgetRulesResponse(val *SDListAssociatedBudgetRulesResponse) *NullableSDListAssociatedBudgetRulesResponse {
	return &NullableSDListAssociatedBudgetRulesResponse{value: val, isSet: true}
}

func (v NullableSDListAssociatedBudgetRulesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDListAssociatedBudgetRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
