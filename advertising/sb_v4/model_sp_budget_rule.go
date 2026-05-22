package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SPBudgetRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SPBudgetRule{}

// SPBudgetRule struct for SPBudgetRule
type SPBudgetRule struct {
	RuleState *State `json:"ruleState,omitempty"`
	// Epoch time of budget rule update. Read-only.
	LastUpdatedDate *float32 `json:"lastUpdatedDate,omitempty"`
	// Epoch time of budget rule creation. Read-only.
	CreatedDate *float32             `json:"createdDate,omitempty"`
	RuleDetails *SPBudgetRuleDetails `json:"ruleDetails,omitempty"`
	// The budget rule identifier.
	RuleId string `json:"ruleId"`
	// The budget rule status. Read-only.
	RuleStatus *string `json:"ruleStatus,omitempty"`
}

type _SPBudgetRule SPBudgetRule

// NewSPBudgetRule instantiates a new SPBudgetRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSPBudgetRule(ruleId string) *SPBudgetRule {
	this := SPBudgetRule{}
	this.RuleId = ruleId
	return &this
}

// NewSPBudgetRuleWithDefaults instantiates a new SPBudgetRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSPBudgetRuleWithDefaults() *SPBudgetRule {
	this := SPBudgetRule{}
	return &this
}

// GetRuleState returns the RuleState field value if set, zero value otherwise.
func (o *SPBudgetRule) GetRuleState() State {
	if o == nil || IsNil(o.RuleState) {
		var ret State
		return ret
	}
	return *o.RuleState
}

// GetRuleStateOk returns a tuple with the RuleState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPBudgetRule) GetRuleStateOk() (*State, bool) {
	if o == nil || IsNil(o.RuleState) {
		return nil, false
	}
	return o.RuleState, true
}

// HasRuleState returns a boolean if a field has been set.
func (o *SPBudgetRule) HasRuleState() bool {
	if o != nil && !IsNil(o.RuleState) {
		return true
	}

	return false
}

// SetRuleState gets a reference to the given State and assigns it to the RuleState field.
func (o *SPBudgetRule) SetRuleState(v State) {
	o.RuleState = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *SPBudgetRule) GetLastUpdatedDate() float32 {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret float32
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPBudgetRule) GetLastUpdatedDateOk() (*float32, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *SPBudgetRule) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given float32 and assigns it to the LastUpdatedDate field.
func (o *SPBudgetRule) SetLastUpdatedDate(v float32) {
	o.LastUpdatedDate = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *SPBudgetRule) GetCreatedDate() float32 {
	if o == nil || IsNil(o.CreatedDate) {
		var ret float32
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPBudgetRule) GetCreatedDateOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *SPBudgetRule) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given float32 and assigns it to the CreatedDate field.
func (o *SPBudgetRule) SetCreatedDate(v float32) {
	o.CreatedDate = &v
}

// GetRuleDetails returns the RuleDetails field value if set, zero value otherwise.
func (o *SPBudgetRule) GetRuleDetails() SPBudgetRuleDetails {
	if o == nil || IsNil(o.RuleDetails) {
		var ret SPBudgetRuleDetails
		return ret
	}
	return *o.RuleDetails
}

// GetRuleDetailsOk returns a tuple with the RuleDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPBudgetRule) GetRuleDetailsOk() (*SPBudgetRuleDetails, bool) {
	if o == nil || IsNil(o.RuleDetails) {
		return nil, false
	}
	return o.RuleDetails, true
}

// HasRuleDetails returns a boolean if a field has been set.
func (o *SPBudgetRule) HasRuleDetails() bool {
	if o != nil && !IsNil(o.RuleDetails) {
		return true
	}

	return false
}

// SetRuleDetails gets a reference to the given SPBudgetRuleDetails and assigns it to the RuleDetails field.
func (o *SPBudgetRule) SetRuleDetails(v SPBudgetRuleDetails) {
	o.RuleDetails = &v
}

// GetRuleId returns the RuleId field value
func (o *SPBudgetRule) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *SPBudgetRule) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value
func (o *SPBudgetRule) SetRuleId(v string) {
	o.RuleId = v
}

// GetRuleStatus returns the RuleStatus field value if set, zero value otherwise.
func (o *SPBudgetRule) GetRuleStatus() string {
	if o == nil || IsNil(o.RuleStatus) {
		var ret string
		return ret
	}
	return *o.RuleStatus
}

// GetRuleStatusOk returns a tuple with the RuleStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPBudgetRule) GetRuleStatusOk() (*string, bool) {
	if o == nil || IsNil(o.RuleStatus) {
		return nil, false
	}
	return o.RuleStatus, true
}

// HasRuleStatus returns a boolean if a field has been set.
func (o *SPBudgetRule) HasRuleStatus() bool {
	if o != nil && !IsNil(o.RuleStatus) {
		return true
	}

	return false
}

// SetRuleStatus gets a reference to the given string and assigns it to the RuleStatus field.
func (o *SPBudgetRule) SetRuleStatus(v string) {
	o.RuleStatus = &v
}

func (o SPBudgetRule) ToMap() (map[string]interface{}, error) {
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

type NullableSPBudgetRule struct {
	value *SPBudgetRule
	isSet bool
}

func (v NullableSPBudgetRule) Get() *SPBudgetRule {
	return v.value
}

func (v *NullableSPBudgetRule) Set(val *SPBudgetRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSPBudgetRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSPBudgetRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPBudgetRule(val *SPBudgetRule) *NullableSPBudgetRule {
	return &NullableSPBudgetRule{value: val, isSet: true}
}

func (v NullableSPBudgetRule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPBudgetRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
