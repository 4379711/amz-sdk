package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBBudgetRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBBudgetRule{}

// SBBudgetRule struct for SBBudgetRule
type SBBudgetRule struct {
	RuleState *State `json:"ruleState,omitempty"`
	// Epoch time of budget rule update. Read-only.
	LastUpdatedDate *float32 `json:"lastUpdatedDate,omitempty"`
	// Epoch time of budget rule creation. Read-only.
	CreatedDate *float32             `json:"createdDate,omitempty"`
	RuleDetails *SBBudgetRuleDetails `json:"ruleDetails,omitempty"`
	// The budget rule identifier.
	RuleId string `json:"ruleId"`
	// The budget rule status. Read-only.
	RuleStatus *string `json:"ruleStatus,omitempty"`
}

type _SBBudgetRule SBBudgetRule

// NewSBBudgetRule instantiates a new SBBudgetRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBBudgetRule(ruleId string) *SBBudgetRule {
	this := SBBudgetRule{}
	this.RuleId = ruleId
	return &this
}

// NewSBBudgetRuleWithDefaults instantiates a new SBBudgetRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBBudgetRuleWithDefaults() *SBBudgetRule {
	this := SBBudgetRule{}
	return &this
}

// GetRuleState returns the RuleState field value if set, zero value otherwise.
func (o *SBBudgetRule) GetRuleState() State {
	if o == nil || IsNil(o.RuleState) {
		var ret State
		return ret
	}
	return *o.RuleState
}

// GetRuleStateOk returns a tuple with the RuleState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBBudgetRule) GetRuleStateOk() (*State, bool) {
	if o == nil || IsNil(o.RuleState) {
		return nil, false
	}
	return o.RuleState, true
}

// HasRuleState returns a boolean if a field has been set.
func (o *SBBudgetRule) HasRuleState() bool {
	if o != nil && !IsNil(o.RuleState) {
		return true
	}

	return false
}

// SetRuleState gets a reference to the given State and assigns it to the RuleState field.
func (o *SBBudgetRule) SetRuleState(v State) {
	o.RuleState = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *SBBudgetRule) GetLastUpdatedDate() float32 {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret float32
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBBudgetRule) GetLastUpdatedDateOk() (*float32, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *SBBudgetRule) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given float32 and assigns it to the LastUpdatedDate field.
func (o *SBBudgetRule) SetLastUpdatedDate(v float32) {
	o.LastUpdatedDate = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *SBBudgetRule) GetCreatedDate() float32 {
	if o == nil || IsNil(o.CreatedDate) {
		var ret float32
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBBudgetRule) GetCreatedDateOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *SBBudgetRule) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given float32 and assigns it to the CreatedDate field.
func (o *SBBudgetRule) SetCreatedDate(v float32) {
	o.CreatedDate = &v
}

// GetRuleDetails returns the RuleDetails field value if set, zero value otherwise.
func (o *SBBudgetRule) GetRuleDetails() SBBudgetRuleDetails {
	if o == nil || IsNil(o.RuleDetails) {
		var ret SBBudgetRuleDetails
		return ret
	}
	return *o.RuleDetails
}

// GetRuleDetailsOk returns a tuple with the RuleDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBBudgetRule) GetRuleDetailsOk() (*SBBudgetRuleDetails, bool) {
	if o == nil || IsNil(o.RuleDetails) {
		return nil, false
	}
	return o.RuleDetails, true
}

// HasRuleDetails returns a boolean if a field has been set.
func (o *SBBudgetRule) HasRuleDetails() bool {
	if o != nil && !IsNil(o.RuleDetails) {
		return true
	}

	return false
}

// SetRuleDetails gets a reference to the given SBBudgetRuleDetails and assigns it to the RuleDetails field.
func (o *SBBudgetRule) SetRuleDetails(v SBBudgetRuleDetails) {
	o.RuleDetails = &v
}

// GetRuleId returns the RuleId field value
func (o *SBBudgetRule) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *SBBudgetRule) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value
func (o *SBBudgetRule) SetRuleId(v string) {
	o.RuleId = v
}

// GetRuleStatus returns the RuleStatus field value if set, zero value otherwise.
func (o *SBBudgetRule) GetRuleStatus() string {
	if o == nil || IsNil(o.RuleStatus) {
		var ret string
		return ret
	}
	return *o.RuleStatus
}

// GetRuleStatusOk returns a tuple with the RuleStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBBudgetRule) GetRuleStatusOk() (*string, bool) {
	if o == nil || IsNil(o.RuleStatus) {
		return nil, false
	}
	return o.RuleStatus, true
}

// HasRuleStatus returns a boolean if a field has been set.
func (o *SBBudgetRule) HasRuleStatus() bool {
	if o != nil && !IsNil(o.RuleStatus) {
		return true
	}

	return false
}

// SetRuleStatus gets a reference to the given string and assigns it to the RuleStatus field.
func (o *SBBudgetRule) SetRuleStatus(v string) {
	o.RuleStatus = &v
}

func (o SBBudgetRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RuleState) {
		toSerialize["ruleState"] = o.RuleState
	}
	if !IsNil(o.LastUpdatedDate) {
		toSerialize["lastUpdatedDate"] = o.LastUpdatedDate
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.RuleDetails) {
		toSerialize["ruleDetails"] = o.RuleDetails
	}
	toSerialize["ruleId"] = o.RuleId
	if !IsNil(o.RuleStatus) {
		toSerialize["ruleStatus"] = o.RuleStatus
	}
	return toSerialize, nil
}

type NullableSBBudgetRule struct {
	value *SBBudgetRule
	isSet bool
}

func (v NullableSBBudgetRule) Get() *SBBudgetRule {
	return v.value
}

func (v *NullableSBBudgetRule) Set(val *SBBudgetRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSBBudgetRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSBBudgetRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBBudgetRule(val *SBBudgetRule) *NullableSBBudgetRule {
	return &NullableSBBudgetRule{value: val, isSet: true}
}

func (v NullableSBBudgetRule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBBudgetRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
