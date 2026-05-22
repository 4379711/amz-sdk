package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SDBudgetRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDBudgetRule{}

// SDBudgetRule struct for SDBudgetRule
type SDBudgetRule struct {
	RuleState *State `json:"ruleState,omitempty"`
	// Epoch time of budget rule update. Read-only.
	LastUpdatedDate *float32 `json:"lastUpdatedDate,omitempty"`
	// Epoch time of budget rule creation. Read-only.
	CreatedDate *float32             `json:"createdDate,omitempty"`
	RuleDetails *SDBudgetRuleDetails `json:"ruleDetails,omitempty"`
	// The budget rule identifier.
	RuleId string `json:"ruleId"`
	// The budget rule status. Read-only.
	RuleStatus *string `json:"ruleStatus,omitempty"`
}

type _SDBudgetRule SDBudgetRule

// NewSDBudgetRule instantiates a new SDBudgetRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDBudgetRule(ruleId string) *SDBudgetRule {
	this := SDBudgetRule{}
	this.RuleId = ruleId
	return &this
}

// NewSDBudgetRuleWithDefaults instantiates a new SDBudgetRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDBudgetRuleWithDefaults() *SDBudgetRule {
	this := SDBudgetRule{}
	return &this
}

// GetRuleState returns the RuleState field value if set, zero value otherwise.
func (o *SDBudgetRule) GetRuleState() State {
	if o == nil || IsNil(o.RuleState) {
		var ret State
		return ret
	}
	return *o.RuleState
}

// GetRuleStateOk returns a tuple with the RuleState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDBudgetRule) GetRuleStateOk() (*State, bool) {
	if o == nil || IsNil(o.RuleState) {
		return nil, false
	}
	return o.RuleState, true
}

// HasRuleState returns a boolean if a field has been set.
func (o *SDBudgetRule) HasRuleState() bool {
	if o != nil && !IsNil(o.RuleState) {
		return true
	}

	return false
}

// SetRuleState gets a reference to the given State and assigns it to the RuleState field.
func (o *SDBudgetRule) SetRuleState(v State) {
	o.RuleState = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *SDBudgetRule) GetLastUpdatedDate() float32 {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret float32
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDBudgetRule) GetLastUpdatedDateOk() (*float32, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *SDBudgetRule) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given float32 and assigns it to the LastUpdatedDate field.
func (o *SDBudgetRule) SetLastUpdatedDate(v float32) {
	o.LastUpdatedDate = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *SDBudgetRule) GetCreatedDate() float32 {
	if o == nil || IsNil(o.CreatedDate) {
		var ret float32
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDBudgetRule) GetCreatedDateOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *SDBudgetRule) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given float32 and assigns it to the CreatedDate field.
func (o *SDBudgetRule) SetCreatedDate(v float32) {
	o.CreatedDate = &v
}

// GetRuleDetails returns the RuleDetails field value if set, zero value otherwise.
func (o *SDBudgetRule) GetRuleDetails() SDBudgetRuleDetails {
	if o == nil || IsNil(o.RuleDetails) {
		var ret SDBudgetRuleDetails
		return ret
	}
	return *o.RuleDetails
}

// GetRuleDetailsOk returns a tuple with the RuleDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDBudgetRule) GetRuleDetailsOk() (*SDBudgetRuleDetails, bool) {
	if o == nil || IsNil(o.RuleDetails) {
		return nil, false
	}
	return o.RuleDetails, true
}

// HasRuleDetails returns a boolean if a field has been set.
func (o *SDBudgetRule) HasRuleDetails() bool {
	if o != nil && !IsNil(o.RuleDetails) {
		return true
	}

	return false
}

// SetRuleDetails gets a reference to the given SDBudgetRuleDetails and assigns it to the RuleDetails field.
func (o *SDBudgetRule) SetRuleDetails(v SDBudgetRuleDetails) {
	o.RuleDetails = &v
}

// GetRuleId returns the RuleId field value
func (o *SDBudgetRule) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *SDBudgetRule) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value
func (o *SDBudgetRule) SetRuleId(v string) {
	o.RuleId = v
}

// GetRuleStatus returns the RuleStatus field value if set, zero value otherwise.
func (o *SDBudgetRule) GetRuleStatus() string {
	if o == nil || IsNil(o.RuleStatus) {
		var ret string
		return ret
	}
	return *o.RuleStatus
}

// GetRuleStatusOk returns a tuple with the RuleStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDBudgetRule) GetRuleStatusOk() (*string, bool) {
	if o == nil || IsNil(o.RuleStatus) {
		return nil, false
	}
	return o.RuleStatus, true
}

// HasRuleStatus returns a boolean if a field has been set.
func (o *SDBudgetRule) HasRuleStatus() bool {
	if o != nil && !IsNil(o.RuleStatus) {
		return true
	}

	return false
}

// SetRuleStatus gets a reference to the given string and assigns it to the RuleStatus field.
func (o *SDBudgetRule) SetRuleStatus(v string) {
	o.RuleStatus = &v
}

func (o SDBudgetRule) ToMap() (map[string]interface{}, error) {
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

type NullableSDBudgetRule struct {
	value *SDBudgetRule
	isSet bool
}

func (v NullableSDBudgetRule) Get() *SDBudgetRule {
	return v.value
}

func (v *NullableSDBudgetRule) Set(val *SDBudgetRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSDBudgetRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSDBudgetRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDBudgetRule(val *SDBudgetRule) *NullableSDBudgetRule {
	return &NullableSDBudgetRule{value: val, isSet: true}
}

func (v NullableSDBudgetRule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDBudgetRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
