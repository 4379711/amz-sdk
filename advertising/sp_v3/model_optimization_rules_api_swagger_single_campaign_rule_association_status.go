package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPISingleCampaignRuleAssociationStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPISingleCampaignRuleAssociationStatus{}

// OptimizationRulesAPISingleCampaignRuleAssociationStatus Object that describes the association status of an optimization rule and a single campaign.
type OptimizationRulesAPISingleCampaignRuleAssociationStatus struct {
	RuleState *string `json:"ruleState,omitempty"`
	// Sp campaign identifier.
	CampaignId *string `json:"campaignId,omitempty"`
	// The rule identifier.
	OptimizationRuleId *string `json:"optimizationRuleId,omitempty"`
	NotificationString *string `json:"notificationString,omitempty"`
}

// NewOptimizationRulesAPISingleCampaignRuleAssociationStatus instantiates a new OptimizationRulesAPISingleCampaignRuleAssociationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPISingleCampaignRuleAssociationStatus() *OptimizationRulesAPISingleCampaignRuleAssociationStatus {
	this := OptimizationRulesAPISingleCampaignRuleAssociationStatus{}
	return &this
}

// NewOptimizationRulesAPISingleCampaignRuleAssociationStatusWithDefaults instantiates a new OptimizationRulesAPISingleCampaignRuleAssociationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPISingleCampaignRuleAssociationStatusWithDefaults() *OptimizationRulesAPISingleCampaignRuleAssociationStatus {
	this := OptimizationRulesAPISingleCampaignRuleAssociationStatus{}
	return &this
}

// GetRuleState returns the RuleState field value if set, zero value otherwise.
func (o *OptimizationRulesAPISingleCampaignRuleAssociationStatus) GetRuleState() string {
	if o == nil || IsNil(o.RuleState) {
		var ret string
		return ret
	}
	return *o.RuleState
}

// GetRuleStateOk returns a tuple with the RuleState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISingleCampaignRuleAssociationStatus) GetRuleStateOk() (*string, bool) {
	if o == nil || IsNil(o.RuleState) {
		return nil, false
	}
	return o.RuleState, true
}

// HasRuleState returns a boolean if a field has been set.
func (o *OptimizationRulesAPISingleCampaignRuleAssociationStatus) HasRuleState() bool {
	if o != nil && !IsNil(o.RuleState) {
		return true
	}

	return false
}

// SetRuleState gets a reference to the given string and assigns it to the RuleState field.
func (o *OptimizationRulesAPISingleCampaignRuleAssociationStatus) SetRuleState(v string) {
	o.RuleState = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *OptimizationRulesAPISingleCampaignRuleAssociationStatus) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISingleCampaignRuleAssociationStatus) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *OptimizationRulesAPISingleCampaignRuleAssociationStatus) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *OptimizationRulesAPISingleCampaignRuleAssociationStatus) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetOptimizationRuleId returns the OptimizationRuleId field value if set, zero value otherwise.
func (o *OptimizationRulesAPISingleCampaignRuleAssociationStatus) GetOptimizationRuleId() string {
	if o == nil || IsNil(o.OptimizationRuleId) {
		var ret string
		return ret
	}
	return *o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISingleCampaignRuleAssociationStatus) GetOptimizationRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.OptimizationRuleId) {
		return nil, false
	}
	return o.OptimizationRuleId, true
}

// HasOptimizationRuleId returns a boolean if a field has been set.
func (o *OptimizationRulesAPISingleCampaignRuleAssociationStatus) HasOptimizationRuleId() bool {
	if o != nil && !IsNil(o.OptimizationRuleId) {
		return true
	}

	return false
}

// SetOptimizationRuleId gets a reference to the given string and assigns it to the OptimizationRuleId field.
func (o *OptimizationRulesAPISingleCampaignRuleAssociationStatus) SetOptimizationRuleId(v string) {
	o.OptimizationRuleId = &v
}

// GetNotificationString returns the NotificationString field value if set, zero value otherwise.
func (o *OptimizationRulesAPISingleCampaignRuleAssociationStatus) GetNotificationString() string {
	if o == nil || IsNil(o.NotificationString) {
		var ret string
		return ret
	}
	return *o.NotificationString
}

// GetNotificationStringOk returns a tuple with the NotificationString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISingleCampaignRuleAssociationStatus) GetNotificationStringOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationString) {
		return nil, false
	}
	return o.NotificationString, true
}

// HasNotificationString returns a boolean if a field has been set.
func (o *OptimizationRulesAPISingleCampaignRuleAssociationStatus) HasNotificationString() bool {
	if o != nil && !IsNil(o.NotificationString) {
		return true
	}

	return false
}

// SetNotificationString gets a reference to the given string and assigns it to the NotificationString field.
func (o *OptimizationRulesAPISingleCampaignRuleAssociationStatus) SetNotificationString(v string) {
	o.NotificationString = &v
}

func (o OptimizationRulesAPISingleCampaignRuleAssociationStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RuleState) {
		toSerialize["ruleState"] = o.RuleState
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	if !IsNil(o.OptimizationRuleId) {
		toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	}
	if !IsNil(o.NotificationString) {
		toSerialize["notificationString"] = o.NotificationString
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPISingleCampaignRuleAssociationStatus struct {
	value *OptimizationRulesAPISingleCampaignRuleAssociationStatus
	isSet bool
}

func (v NullableOptimizationRulesAPISingleCampaignRuleAssociationStatus) Get() *OptimizationRulesAPISingleCampaignRuleAssociationStatus {
	return v.value
}

func (v *NullableOptimizationRulesAPISingleCampaignRuleAssociationStatus) Set(val *OptimizationRulesAPISingleCampaignRuleAssociationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPISingleCampaignRuleAssociationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPISingleCampaignRuleAssociationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPISingleCampaignRuleAssociationStatus(val *OptimizationRulesAPISingleCampaignRuleAssociationStatus) *NullableOptimizationRulesAPISingleCampaignRuleAssociationStatus {
	return &NullableOptimizationRulesAPISingleCampaignRuleAssociationStatus{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPISingleCampaignRuleAssociationStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPISingleCampaignRuleAssociationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
