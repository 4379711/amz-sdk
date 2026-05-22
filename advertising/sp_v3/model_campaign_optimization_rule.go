package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CampaignOptimizationRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignOptimizationRule{}

// CampaignOptimizationRule struct for CampaignOptimizationRule
type CampaignOptimizationRule struct {
	Recurrence *RecurrenceType `json:"recurrence,omitempty"`
	RuleAction *RuleAction     `json:"ruleAction,omitempty"`
	// The persistent rule identifier.
	CampaignOptimizationId string `json:"campaignOptimizationId"`
	// Time of campaign optimization rule creation in ISO 8061. Read-only.
	CreatedDate   *string         `json:"createdDate,omitempty"`
	RuleCondition []RuleCondition `json:"ruleCondition,omitempty"`
	RuleType      *RuleType       `json:"ruleType,omitempty"`
	// The campaign optimization rule name.
	RuleName    *string     `json:"ruleName,omitempty"`
	CampaignIds []string    `json:"campaignIds,omitempty"`
	RuleStatus  *RuleStatus `json:"ruleStatus,omitempty"`
}

type _CampaignOptimizationRule CampaignOptimizationRule

// NewCampaignOptimizationRule instantiates a new CampaignOptimizationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignOptimizationRule(campaignOptimizationId string) *CampaignOptimizationRule {
	this := CampaignOptimizationRule{}
	this.CampaignOptimizationId = campaignOptimizationId
	return &this
}

// NewCampaignOptimizationRuleWithDefaults instantiates a new CampaignOptimizationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignOptimizationRuleWithDefaults() *CampaignOptimizationRule {
	this := CampaignOptimizationRule{}
	return &this
}

// GetRecurrence returns the Recurrence field value if set, zero value otherwise.
func (o *CampaignOptimizationRule) GetRecurrence() RecurrenceType {
	if o == nil || IsNil(o.Recurrence) {
		var ret RecurrenceType
		return ret
	}
	return *o.Recurrence
}

// GetRecurrenceOk returns a tuple with the Recurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignOptimizationRule) GetRecurrenceOk() (*RecurrenceType, bool) {
	if o == nil || IsNil(o.Recurrence) {
		return nil, false
	}
	return o.Recurrence, true
}

// HasRecurrence returns a boolean if a field has been set.
func (o *CampaignOptimizationRule) HasRecurrence() bool {
	if o != nil && !IsNil(o.Recurrence) {
		return true
	}

	return false
}

// SetRecurrence gets a reference to the given RecurrenceType and assigns it to the Recurrence field.
func (o *CampaignOptimizationRule) SetRecurrence(v RecurrenceType) {
	o.Recurrence = &v
}

// GetRuleAction returns the RuleAction field value if set, zero value otherwise.
func (o *CampaignOptimizationRule) GetRuleAction() RuleAction {
	if o == nil || IsNil(o.RuleAction) {
		var ret RuleAction
		return ret
	}
	return *o.RuleAction
}

// GetRuleActionOk returns a tuple with the RuleAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignOptimizationRule) GetRuleActionOk() (*RuleAction, bool) {
	if o == nil || IsNil(o.RuleAction) {
		return nil, false
	}
	return o.RuleAction, true
}

// HasRuleAction returns a boolean if a field has been set.
func (o *CampaignOptimizationRule) HasRuleAction() bool {
	if o != nil && !IsNil(o.RuleAction) {
		return true
	}

	return false
}

// SetRuleAction gets a reference to the given RuleAction and assigns it to the RuleAction field.
func (o *CampaignOptimizationRule) SetRuleAction(v RuleAction) {
	o.RuleAction = &v
}

// GetCampaignOptimizationId returns the CampaignOptimizationId field value
func (o *CampaignOptimizationRule) GetCampaignOptimizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignOptimizationId
}

// GetCampaignOptimizationIdOk returns a tuple with the CampaignOptimizationId field value
// and a boolean to check if the value has been set.
func (o *CampaignOptimizationRule) GetCampaignOptimizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignOptimizationId, true
}

// SetCampaignOptimizationId sets field value
func (o *CampaignOptimizationRule) SetCampaignOptimizationId(v string) {
	o.CampaignOptimizationId = v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *CampaignOptimizationRule) GetCreatedDate() string {
	if o == nil || IsNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignOptimizationRule) GetCreatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *CampaignOptimizationRule) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *CampaignOptimizationRule) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetRuleCondition returns the RuleCondition field value if set, zero value otherwise.
func (o *CampaignOptimizationRule) GetRuleCondition() []RuleCondition {
	if o == nil || IsNil(o.RuleCondition) {
		var ret []RuleCondition
		return ret
	}
	return o.RuleCondition
}

// GetRuleConditionOk returns a tuple with the RuleCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignOptimizationRule) GetRuleConditionOk() ([]RuleCondition, bool) {
	if o == nil || IsNil(o.RuleCondition) {
		return nil, false
	}
	return o.RuleCondition, true
}

// HasRuleCondition returns a boolean if a field has been set.
func (o *CampaignOptimizationRule) HasRuleCondition() bool {
	if o != nil && !IsNil(o.RuleCondition) {
		return true
	}

	return false
}

// SetRuleCondition gets a reference to the given []RuleCondition and assigns it to the RuleCondition field.
func (o *CampaignOptimizationRule) SetRuleCondition(v []RuleCondition) {
	o.RuleCondition = v
}

// GetRuleType returns the RuleType field value if set, zero value otherwise.
func (o *CampaignOptimizationRule) GetRuleType() RuleType {
	if o == nil || IsNil(o.RuleType) {
		var ret RuleType
		return ret
	}
	return *o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignOptimizationRule) GetRuleTypeOk() (*RuleType, bool) {
	if o == nil || IsNil(o.RuleType) {
		return nil, false
	}
	return o.RuleType, true
}

// HasRuleType returns a boolean if a field has been set.
func (o *CampaignOptimizationRule) HasRuleType() bool {
	if o != nil && !IsNil(o.RuleType) {
		return true
	}

	return false
}

// SetRuleType gets a reference to the given RuleType and assigns it to the RuleType field.
func (o *CampaignOptimizationRule) SetRuleType(v RuleType) {
	o.RuleType = &v
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *CampaignOptimizationRule) GetRuleName() string {
	if o == nil || IsNil(o.RuleName) {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignOptimizationRule) GetRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RuleName) {
		return nil, false
	}
	return o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *CampaignOptimizationRule) HasRuleName() bool {
	if o != nil && !IsNil(o.RuleName) {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *CampaignOptimizationRule) SetRuleName(v string) {
	o.RuleName = &v
}

// GetCampaignIds returns the CampaignIds field value if set, zero value otherwise.
func (o *CampaignOptimizationRule) GetCampaignIds() []string {
	if o == nil || IsNil(o.CampaignIds) {
		var ret []string
		return ret
	}
	return o.CampaignIds
}

// GetCampaignIdsOk returns a tuple with the CampaignIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignOptimizationRule) GetCampaignIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CampaignIds) {
		return nil, false
	}
	return o.CampaignIds, true
}

// HasCampaignIds returns a boolean if a field has been set.
func (o *CampaignOptimizationRule) HasCampaignIds() bool {
	if o != nil && !IsNil(o.CampaignIds) {
		return true
	}

	return false
}

// SetCampaignIds gets a reference to the given []string and assigns it to the CampaignIds field.
func (o *CampaignOptimizationRule) SetCampaignIds(v []string) {
	o.CampaignIds = v
}

// GetRuleStatus returns the RuleStatus field value if set, zero value otherwise.
func (o *CampaignOptimizationRule) GetRuleStatus() RuleStatus {
	if o == nil || IsNil(o.RuleStatus) {
		var ret RuleStatus
		return ret
	}
	return *o.RuleStatus
}

// GetRuleStatusOk returns a tuple with the RuleStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignOptimizationRule) GetRuleStatusOk() (*RuleStatus, bool) {
	if o == nil || IsNil(o.RuleStatus) {
		return nil, false
	}
	return o.RuleStatus, true
}

// HasRuleStatus returns a boolean if a field has been set.
func (o *CampaignOptimizationRule) HasRuleStatus() bool {
	if o != nil && !IsNil(o.RuleStatus) {
		return true
	}

	return false
}

// SetRuleStatus gets a reference to the given RuleStatus and assigns it to the RuleStatus field.
func (o *CampaignOptimizationRule) SetRuleStatus(v RuleStatus) {
	o.RuleStatus = &v
}

func (o CampaignOptimizationRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Recurrence) {
		toSerialize["recurrence"] = o.Recurrence
	}
	if !IsNil(o.RuleAction) {
		toSerialize["ruleAction"] = o.RuleAction
	}
	toSerialize["campaignOptimizationId"] = o.CampaignOptimizationId
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.RuleCondition) {
		toSerialize["ruleCondition"] = o.RuleCondition
	}
	if !IsNil(o.RuleType) {
		toSerialize["ruleType"] = o.RuleType
	}
	if !IsNil(o.RuleName) {
		toSerialize["ruleName"] = o.RuleName
	}
	if !IsNil(o.CampaignIds) {
		toSerialize["campaignIds"] = o.CampaignIds
	}
	if !IsNil(o.RuleStatus) {
		toSerialize["ruleStatus"] = o.RuleStatus
	}
	return toSerialize, nil
}

type NullableCampaignOptimizationRule struct {
	value *CampaignOptimizationRule
	isSet bool
}

func (v NullableCampaignOptimizationRule) Get() *CampaignOptimizationRule {
	return v.value
}

func (v *NullableCampaignOptimizationRule) Set(val *CampaignOptimizationRule) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignOptimizationRule) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignOptimizationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignOptimizationRule(val *CampaignOptimizationRule) *NullableCampaignOptimizationRule {
	return &NullableCampaignOptimizationRule{value: val, isSet: true}
}

func (v NullableCampaignOptimizationRule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCampaignOptimizationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
