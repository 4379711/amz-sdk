package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSPCampaignOptimizationRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSPCampaignOptimizationRulesRequest{}

// UpdateSPCampaignOptimizationRulesRequest Request object for updating campaign optimization rule
type UpdateSPCampaignOptimizationRulesRequest struct {
	Recurrence RecurrenceType `json:"recurrence"`
	RuleAction RuleAction     `json:"ruleAction"`
	// The persistent rule identifier.
	CampaignOptimizationId string          `json:"campaignOptimizationId"`
	RuleCondition          []RuleCondition `json:"ruleCondition"`
	RuleType               RuleType        `json:"ruleType"`
	// The campaign optimization rule name.
	RuleName *string `json:"ruleName,omitempty"`
	// A list of campaign ids
	CampaignIds []string `json:"campaignIds"`
}

type _UpdateSPCampaignOptimizationRulesRequest UpdateSPCampaignOptimizationRulesRequest

// NewUpdateSPCampaignOptimizationRulesRequest instantiates a new UpdateSPCampaignOptimizationRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSPCampaignOptimizationRulesRequest(recurrence RecurrenceType, ruleAction RuleAction, campaignOptimizationId string, ruleCondition []RuleCondition, ruleType RuleType, campaignIds []string) *UpdateSPCampaignOptimizationRulesRequest {
	this := UpdateSPCampaignOptimizationRulesRequest{}
	this.Recurrence = recurrence
	this.RuleAction = ruleAction
	this.CampaignOptimizationId = campaignOptimizationId
	this.RuleCondition = ruleCondition
	this.RuleType = ruleType
	this.CampaignIds = campaignIds
	return &this
}

// NewUpdateSPCampaignOptimizationRulesRequestWithDefaults instantiates a new UpdateSPCampaignOptimizationRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSPCampaignOptimizationRulesRequestWithDefaults() *UpdateSPCampaignOptimizationRulesRequest {
	this := UpdateSPCampaignOptimizationRulesRequest{}
	return &this
}

// GetRecurrence returns the Recurrence field value
func (o *UpdateSPCampaignOptimizationRulesRequest) GetRecurrence() RecurrenceType {
	if o == nil {
		var ret RecurrenceType
		return ret
	}

	return o.Recurrence
}

// GetRecurrenceOk returns a tuple with the Recurrence field value
// and a boolean to check if the value has been set.
func (o *UpdateSPCampaignOptimizationRulesRequest) GetRecurrenceOk() (*RecurrenceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recurrence, true
}

// SetRecurrence sets field value
func (o *UpdateSPCampaignOptimizationRulesRequest) SetRecurrence(v RecurrenceType) {
	o.Recurrence = v
}

// GetRuleAction returns the RuleAction field value
func (o *UpdateSPCampaignOptimizationRulesRequest) GetRuleAction() RuleAction {
	if o == nil {
		var ret RuleAction
		return ret
	}

	return o.RuleAction
}

// GetRuleActionOk returns a tuple with the RuleAction field value
// and a boolean to check if the value has been set.
func (o *UpdateSPCampaignOptimizationRulesRequest) GetRuleActionOk() (*RuleAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleAction, true
}

// SetRuleAction sets field value
func (o *UpdateSPCampaignOptimizationRulesRequest) SetRuleAction(v RuleAction) {
	o.RuleAction = v
}

// GetCampaignOptimizationId returns the CampaignOptimizationId field value
func (o *UpdateSPCampaignOptimizationRulesRequest) GetCampaignOptimizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignOptimizationId
}

// GetCampaignOptimizationIdOk returns a tuple with the CampaignOptimizationId field value
// and a boolean to check if the value has been set.
func (o *UpdateSPCampaignOptimizationRulesRequest) GetCampaignOptimizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignOptimizationId, true
}

// SetCampaignOptimizationId sets field value
func (o *UpdateSPCampaignOptimizationRulesRequest) SetCampaignOptimizationId(v string) {
	o.CampaignOptimizationId = v
}

// GetRuleCondition returns the RuleCondition field value
func (o *UpdateSPCampaignOptimizationRulesRequest) GetRuleCondition() []RuleCondition {
	if o == nil {
		var ret []RuleCondition
		return ret
	}

	return o.RuleCondition
}

// GetRuleConditionOk returns a tuple with the RuleCondition field value
// and a boolean to check if the value has been set.
func (o *UpdateSPCampaignOptimizationRulesRequest) GetRuleConditionOk() ([]RuleCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleCondition, true
}

// SetRuleCondition sets field value
func (o *UpdateSPCampaignOptimizationRulesRequest) SetRuleCondition(v []RuleCondition) {
	o.RuleCondition = v
}

// GetRuleType returns the RuleType field value
func (o *UpdateSPCampaignOptimizationRulesRequest) GetRuleType() RuleType {
	if o == nil {
		var ret RuleType
		return ret
	}

	return o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value
// and a boolean to check if the value has been set.
func (o *UpdateSPCampaignOptimizationRulesRequest) GetRuleTypeOk() (*RuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleType, true
}

// SetRuleType sets field value
func (o *UpdateSPCampaignOptimizationRulesRequest) SetRuleType(v RuleType) {
	o.RuleType = v
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *UpdateSPCampaignOptimizationRulesRequest) GetRuleName() string {
	if o == nil || IsNil(o.RuleName) {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSPCampaignOptimizationRulesRequest) GetRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RuleName) {
		return nil, false
	}
	return o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *UpdateSPCampaignOptimizationRulesRequest) HasRuleName() bool {
	if o != nil && !IsNil(o.RuleName) {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *UpdateSPCampaignOptimizationRulesRequest) SetRuleName(v string) {
	o.RuleName = &v
}

// GetCampaignIds returns the CampaignIds field value
func (o *UpdateSPCampaignOptimizationRulesRequest) GetCampaignIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CampaignIds
}

// GetCampaignIdsOk returns a tuple with the CampaignIds field value
// and a boolean to check if the value has been set.
func (o *UpdateSPCampaignOptimizationRulesRequest) GetCampaignIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CampaignIds, true
}

// SetCampaignIds sets field value
func (o *UpdateSPCampaignOptimizationRulesRequest) SetCampaignIds(v []string) {
	o.CampaignIds = v
}

func (o UpdateSPCampaignOptimizationRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["recurrence"] = o.Recurrence
	toSerialize["ruleAction"] = o.RuleAction
	toSerialize["campaignOptimizationId"] = o.CampaignOptimizationId
	toSerialize["ruleCondition"] = o.RuleCondition
	toSerialize["ruleType"] = o.RuleType
	if !IsNil(o.RuleName) {
		toSerialize["ruleName"] = o.RuleName
	}
	toSerialize["campaignIds"] = o.CampaignIds
	return toSerialize, nil
}

type NullableUpdateSPCampaignOptimizationRulesRequest struct {
	value *UpdateSPCampaignOptimizationRulesRequest
	isSet bool
}

func (v NullableUpdateSPCampaignOptimizationRulesRequest) Get() *UpdateSPCampaignOptimizationRulesRequest {
	return v.value
}

func (v *NullableUpdateSPCampaignOptimizationRulesRequest) Set(val *UpdateSPCampaignOptimizationRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSPCampaignOptimizationRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSPCampaignOptimizationRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSPCampaignOptimizationRulesRequest(val *UpdateSPCampaignOptimizationRulesRequest) *NullableUpdateSPCampaignOptimizationRulesRequest {
	return &NullableUpdateSPCampaignOptimizationRulesRequest{value: val, isSet: true}
}

func (v NullableUpdateSPCampaignOptimizationRulesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSPCampaignOptimizationRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
