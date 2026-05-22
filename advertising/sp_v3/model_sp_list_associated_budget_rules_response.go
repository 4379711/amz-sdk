package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SPListAssociatedBudgetRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SPListAssociatedBudgetRulesResponse{}

// SPListAssociatedBudgetRulesResponse struct for SPListAssociatedBudgetRulesResponse
type SPListAssociatedBudgetRulesResponse struct {
	// A list of associated budget rules.
	AssociatedRules []SPCampaignBudgetRule `json:"associatedRules,omitempty"`
}

// NewSPListAssociatedBudgetRulesResponse instantiates a new SPListAssociatedBudgetRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSPListAssociatedBudgetRulesResponse() *SPListAssociatedBudgetRulesResponse {
	this := SPListAssociatedBudgetRulesResponse{}
	return &this
}

// NewSPListAssociatedBudgetRulesResponseWithDefaults instantiates a new SPListAssociatedBudgetRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSPListAssociatedBudgetRulesResponseWithDefaults() *SPListAssociatedBudgetRulesResponse {
	this := SPListAssociatedBudgetRulesResponse{}
	return &this
}

// GetAssociatedRules returns the AssociatedRules field value if set, zero value otherwise.
func (o *SPListAssociatedBudgetRulesResponse) GetAssociatedRules() []SPCampaignBudgetRule {
	if o == nil || IsNil(o.AssociatedRules) {
		var ret []SPCampaignBudgetRule
		return ret
	}
	return o.AssociatedRules
}

// GetAssociatedRulesOk returns a tuple with the AssociatedRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPListAssociatedBudgetRulesResponse) GetAssociatedRulesOk() ([]SPCampaignBudgetRule, bool) {
	if o == nil || IsNil(o.AssociatedRules) {
		return nil, false
	}
	return o.AssociatedRules, true
}

// HasAssociatedRules returns a boolean if a field has been set.
func (o *SPListAssociatedBudgetRulesResponse) HasAssociatedRules() bool {
	if o != nil && !IsNil(o.AssociatedRules) {
		return true
	}

	return false
}

// SetAssociatedRules gets a reference to the given []SPCampaignBudgetRule and assigns it to the AssociatedRules field.
func (o *SPListAssociatedBudgetRulesResponse) SetAssociatedRules(v []SPCampaignBudgetRule) {
	o.AssociatedRules = v
}

func (o SPListAssociatedBudgetRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociatedRules) {
		toSerialize["associatedRules"] = o.AssociatedRules
	}
	return toSerialize, nil
}

type NullableSPListAssociatedBudgetRulesResponse struct {
	value *SPListAssociatedBudgetRulesResponse
	isSet bool
}

func (v NullableSPListAssociatedBudgetRulesResponse) Get() *SPListAssociatedBudgetRulesResponse {
	return v.value
}

func (v *NullableSPListAssociatedBudgetRulesResponse) Set(val *SPListAssociatedBudgetRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSPListAssociatedBudgetRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSPListAssociatedBudgetRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPListAssociatedBudgetRulesResponse(val *SPListAssociatedBudgetRulesResponse) *NullableSPListAssociatedBudgetRulesResponse {
	return &NullableSPListAssociatedBudgetRulesResponse{value: val, isSet: true}
}

func (v NullableSPListAssociatedBudgetRulesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPListAssociatedBudgetRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
