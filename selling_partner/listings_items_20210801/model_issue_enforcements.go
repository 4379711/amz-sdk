package listings_items_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the IssueEnforcements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssueEnforcements{}

// IssueEnforcements This field provides information about the enforcement actions taken by Amazon that affect the publishing or status of a listing. It also includes details about any associated exemptions.
type IssueEnforcements struct {
	// List of enforcement actions taken by Amazon that affect the publishing or status of a listing.
	Actions   []IssueEnforcementAction `json:"actions"`
	Exemption IssueExemption           `json:"exemption"`
}

type _IssueEnforcements IssueEnforcements

// NewIssueEnforcements instantiates a new IssueEnforcements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueEnforcements(actions []IssueEnforcementAction, exemption IssueExemption) *IssueEnforcements {
	this := IssueEnforcements{}
	this.Actions = actions
	this.Exemption = exemption
	return &this
}

// NewIssueEnforcementsWithDefaults instantiates a new IssueEnforcements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueEnforcementsWithDefaults() *IssueEnforcements {
	this := IssueEnforcements{}
	return &this
}

// GetActions returns the Actions field value
func (o *IssueEnforcements) GetActions() []IssueEnforcementAction {
	if o == nil {
		var ret []IssueEnforcementAction
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *IssueEnforcements) GetActionsOk() ([]IssueEnforcementAction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *IssueEnforcements) SetActions(v []IssueEnforcementAction) {
	o.Actions = v
}

// GetExemption returns the Exemption field value
func (o *IssueEnforcements) GetExemption() IssueExemption {
	if o == nil {
		var ret IssueExemption
		return ret
	}

	return o.Exemption
}

// GetExemptionOk returns a tuple with the Exemption field value
// and a boolean to check if the value has been set.
func (o *IssueEnforcements) GetExemptionOk() (*IssueExemption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exemption, true
}

// SetExemption sets field value
func (o *IssueEnforcements) SetExemption(v IssueExemption) {
	o.Exemption = v
}

func (o IssueEnforcements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["actions"] = o.Actions
	toSerialize["exemption"] = o.Exemption
	return toSerialize, nil
}

type NullableIssueEnforcements struct {
	value *IssueEnforcements
	isSet bool
}

func (v NullableIssueEnforcements) Get() *IssueEnforcements {
	return v.value
}

func (v *NullableIssueEnforcements) Set(val *IssueEnforcements) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueEnforcements) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueEnforcements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueEnforcements(val *IssueEnforcements) *NullableIssueEnforcements {
	return &NullableIssueEnforcements{value: val, isSet: true}
}

func (v NullableIssueEnforcements) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableIssueEnforcements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
