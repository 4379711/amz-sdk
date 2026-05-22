package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSPCampaignOptimizationRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSPCampaignOptimizationRulesRequest{}

// CreateSPCampaignOptimizationRulesRequest struct for CreateSPCampaignOptimizationRulesRequest
type CreateSPCampaignOptimizationRulesRequest struct {
	Recurrence    RecurrenceType  `json:"recurrence"`
	RuleAction    RuleAction      `json:"ruleAction"`
	RuleCondition []RuleCondition `json:"ruleCondition"`
	RuleType      RuleType        `json:"ruleType"`
	// The campaign optimization rule name.
	RuleName *string `json:"ruleName,omitempty"`
	// A list of campaign ids
	CampaignIds []string `json:"campaignIds"`
}

type _CreateSPCampaignOptimizationRulesRequest CreateSPCampaignOptimizationRulesRequest

// NewCreateSPCampaignOptimizationRulesRequest instantiates a new CreateSPCampaignOptimizationRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSPCampaignOptimizationRulesRequest(recurrence RecurrenceType, ruleAction RuleAction, ruleCondition []RuleCondition, ruleType RuleType, campaignIds []string) *CreateSPCampaignOptimizationRulesRequest {
	this := CreateSPCampaignOptimizationRulesRequest{}
	this.Recurrence = recurrence
	this.RuleAction = ruleAction
	this.RuleCondition = ruleCondition
	this.RuleType = ruleType
	this.CampaignIds = campaignIds
	return &this
}

// NewCreateSPCampaignOptimizationRulesRequestWithDefaults instantiates a new CreateSPCampaignOptimizationRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSPCampaignOptimizationRulesRequestWithDefaults() *CreateSPCampaignOptimizationRulesRequest {
	this := CreateSPCampaignOptimizationRulesRequest{}
	return &this
}

// GetRecurrence returns the Recurrence field value
func (o *CreateSPCampaignOptimizationRulesRequest) GetRecurrence() RecurrenceType {
	if o == nil {
		var ret RecurrenceType
		return ret
	}

	return o.Recurrence
}

// GetRecurrenceOk returns a tuple with the Recurrence field value
// and a boolean to check if the value has been set.
func (o *CreateSPCampaignOptimizationRulesRequest) GetRecurrenceOk() (*RecurrenceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recurrence, true
}

// SetRecurrence sets field value
func (o *CreateSPCampaignOptimizationRulesRequest) SetRecurrence(v RecurrenceType) {
	o.Recurrence = v
}

// GetRuleAction returns the RuleAction field value
func (o *CreateSPCampaignOptimizationRulesRequest) GetRuleAction() RuleAction {
	if o == nil {
		var ret RuleAction
		return ret
	}

	return o.RuleAction
}

// GetRuleActionOk returns a tuple with the RuleAction field value
// and a boolean to check if the value has been set.
func (o *CreateSPCampaignOptimizationRulesRequest) GetRuleActionOk() (*RuleAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleAction, true
}

// SetRuleAction sets field value
func (o *CreateSPCampaignOptimizationRulesRequest) SetRuleAction(v RuleAction) {
	o.RuleAction = v
}

// GetRuleCondition returns the RuleCondition field value
func (o *CreateSPCampaignOptimizationRulesRequest) GetRuleCondition() []RuleCondition {
	if o == nil {
		var ret []RuleCondition
		return ret
	}

	return o.RuleCondition
}

// GetRuleConditionOk returns a tuple with the RuleCondition field value
// and a boolean to check if the value has been set.
func (o *CreateSPCampaignOptimizationRulesRequest) GetRuleConditionOk() ([]RuleCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleCondition, true
}

// SetRuleCondition sets field value
func (o *CreateSPCampaignOptimizationRulesRequest) SetRuleCondition(v []RuleCondition) {
	o.RuleCondition = v
}

// GetRuleType returns the RuleType field value
func (o *CreateSPCampaignOptimizationRulesRequest) GetRuleType() RuleType {
	if o == nil {
		var ret RuleType
		return ret
	}

	return o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value
// and a boolean to check if the value has been set.
func (o *CreateSPCampaignOptimizationRulesRequest) GetRuleTypeOk() (*RuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleType, true
}

// SetRuleType sets field value
func (o *CreateSPCampaignOptimizationRulesRequest) SetRuleType(v RuleType) {
	o.RuleType = v
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *CreateSPCampaignOptimizationRulesRequest) GetRuleName() string {
	if o == nil || IsNil(o.RuleName) {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSPCampaignOptimizationRulesRequest) GetRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RuleName) {
		return nil, false
	}
	return o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *CreateSPCampaignOptimizationRulesRequest) HasRuleName() bool {
	if o != nil && !IsNil(o.RuleName) {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *CreateSPCampaignOptimizationRulesRequest) SetRuleName(v string) {
	o.RuleName = &v
}

// GetCampaignIds returns the CampaignIds field value
func (o *CreateSPCampaignOptimizationRulesRequest) GetCampaignIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CampaignIds
}

// GetCampaignIdsOk returns a tuple with the CampaignIds field value
// and a boolean to check if the value has been set.
func (o *CreateSPCampaignOptimizationRulesRequest) GetCampaignIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CampaignIds, true
}

// SetCampaignIds sets field value
func (o *CreateSPCampaignOptimizationRulesRequest) SetCampaignIds(v []string) {
	o.CampaignIds = v
}

func (o CreateSPCampaignOptimizationRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["recurrence"] = o.Recurrence
	toSerialize["ruleAction"] = o.RuleAction
	toSerialize["ruleCondition"] = o.RuleCondition
	toSerialize["ruleType"] = o.RuleType
	if !IsNil(o.RuleName) {
		toSerialize["ruleName"] = o.RuleName
	}
	toSerialize["campaignIds"] = o.CampaignIds
	return toSerialize, nil
}

type NullableCreateSPCampaignOptimizationRulesRequest struct {
	value *CreateSPCampaignOptimizationRulesRequest
	isSet bool
}

func (v NullableCreateSPCampaignOptimizationRulesRequest) Get() *CreateSPCampaignOptimizationRulesRequest {
	return v.value
}

func (v *NullableCreateSPCampaignOptimizationRulesRequest) Set(val *CreateSPCampaignOptimizationRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSPCampaignOptimizationRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSPCampaignOptimizationRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSPCampaignOptimizationRulesRequest(val *CreateSPCampaignOptimizationRulesRequest) *NullableCreateSPCampaignOptimizationRulesRequest {
	return &NullableCreateSPCampaignOptimizationRulesRequest{value: val, isSet: true}
}

func (v NullableCreateSPCampaignOptimizationRulesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSPCampaignOptimizationRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
