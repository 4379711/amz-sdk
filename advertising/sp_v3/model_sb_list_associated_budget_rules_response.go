package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SBListAssociatedBudgetRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBListAssociatedBudgetRulesResponse{}

// SBListAssociatedBudgetRulesResponse struct for SBListAssociatedBudgetRulesResponse
type SBListAssociatedBudgetRulesResponse struct {
	// A list of associated budget rules.
	AssociatedRules []SBCampaignBudgetRule `json:"associatedRules,omitempty"`
}

// NewSBListAssociatedBudgetRulesResponse instantiates a new SBListAssociatedBudgetRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBListAssociatedBudgetRulesResponse() *SBListAssociatedBudgetRulesResponse {
	this := SBListAssociatedBudgetRulesResponse{}
	return &this
}

// NewSBListAssociatedBudgetRulesResponseWithDefaults instantiates a new SBListAssociatedBudgetRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBListAssociatedBudgetRulesResponseWithDefaults() *SBListAssociatedBudgetRulesResponse {
	this := SBListAssociatedBudgetRulesResponse{}
	return &this
}

// GetAssociatedRules returns the AssociatedRules field value if set, zero value otherwise.
func (o *SBListAssociatedBudgetRulesResponse) GetAssociatedRules() []SBCampaignBudgetRule {
	if o == nil || IsNil(o.AssociatedRules) {
		var ret []SBCampaignBudgetRule
		return ret
	}
	return o.AssociatedRules
}

// GetAssociatedRulesOk returns a tuple with the AssociatedRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBListAssociatedBudgetRulesResponse) GetAssociatedRulesOk() ([]SBCampaignBudgetRule, bool) {
	if o == nil || IsNil(o.AssociatedRules) {
		return nil, false
	}
	return o.AssociatedRules, true
}

// HasAssociatedRules returns a boolean if a field has been set.
func (o *SBListAssociatedBudgetRulesResponse) HasAssociatedRules() bool {
	if o != nil && !IsNil(o.AssociatedRules) {
		return true
	}

	return false
}

// SetAssociatedRules gets a reference to the given []SBCampaignBudgetRule and assigns it to the AssociatedRules field.
func (o *SBListAssociatedBudgetRulesResponse) SetAssociatedRules(v []SBCampaignBudgetRule) {
	o.AssociatedRules = v
}

func (o SBListAssociatedBudgetRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociatedRules) {
		toSerialize["associatedRules"] = o.AssociatedRules
	}
	return toSerialize, nil
}

type NullableSBListAssociatedBudgetRulesResponse struct {
	value *SBListAssociatedBudgetRulesResponse
	isSet bool
}

func (v NullableSBListAssociatedBudgetRulesResponse) Get() *SBListAssociatedBudgetRulesResponse {
	return v.value
}

func (v *NullableSBListAssociatedBudgetRulesResponse) Set(val *SBListAssociatedBudgetRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSBListAssociatedBudgetRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSBListAssociatedBudgetRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBListAssociatedBudgetRulesResponse(val *SBListAssociatedBudgetRulesResponse) *NullableSBListAssociatedBudgetRulesResponse {
	return &NullableSBListAssociatedBudgetRulesResponse{value: val, isSet: true}
}

func (v NullableSBListAssociatedBudgetRulesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBListAssociatedBudgetRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
