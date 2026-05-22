package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the RuleNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleNotification{}

// RuleNotification struct for RuleNotification
type RuleNotification struct {
	RuleState *RuleState `json:"ruleState,omitempty"`
	// The persistent rule identifier.
	CampaignOptimizationId *string `json:"campaignOptimizationId,omitempty"`
	// campaignId
	CampaignId *string `json:"campaignId,omitempty"`
	// Explains why the rule state is disabled
	NotificationString *string `json:"notificationString,omitempty"`
}

// NewRuleNotification instantiates a new RuleNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleNotification() *RuleNotification {
	this := RuleNotification{}
	return &this
}

// NewRuleNotificationWithDefaults instantiates a new RuleNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleNotificationWithDefaults() *RuleNotification {
	this := RuleNotification{}
	return &this
}

// GetRuleState returns the RuleState field value if set, zero value otherwise.
func (o *RuleNotification) GetRuleState() RuleState {
	if o == nil || IsNil(o.RuleState) {
		var ret RuleState
		return ret
	}
	return *o.RuleState
}

// GetRuleStateOk returns a tuple with the RuleState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleNotification) GetRuleStateOk() (*RuleState, bool) {
	if o == nil || IsNil(o.RuleState) {
		return nil, false
	}
	return o.RuleState, true
}

// HasRuleState returns a boolean if a field has been set.
func (o *RuleNotification) HasRuleState() bool {
	if o != nil && !IsNil(o.RuleState) {
		return true
	}

	return false
}

// SetRuleState gets a reference to the given RuleState and assigns it to the RuleState field.
func (o *RuleNotification) SetRuleState(v RuleState) {
	o.RuleState = &v
}

// GetCampaignOptimizationId returns the CampaignOptimizationId field value if set, zero value otherwise.
func (o *RuleNotification) GetCampaignOptimizationId() string {
	if o == nil || IsNil(o.CampaignOptimizationId) {
		var ret string
		return ret
	}
	return *o.CampaignOptimizationId
}

// GetCampaignOptimizationIdOk returns a tuple with the CampaignOptimizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleNotification) GetCampaignOptimizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignOptimizationId) {
		return nil, false
	}
	return o.CampaignOptimizationId, true
}

// HasCampaignOptimizationId returns a boolean if a field has been set.
func (o *RuleNotification) HasCampaignOptimizationId() bool {
	if o != nil && !IsNil(o.CampaignOptimizationId) {
		return true
	}

	return false
}

// SetCampaignOptimizationId gets a reference to the given string and assigns it to the CampaignOptimizationId field.
func (o *RuleNotification) SetCampaignOptimizationId(v string) {
	o.CampaignOptimizationId = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *RuleNotification) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleNotification) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *RuleNotification) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *RuleNotification) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetNotificationString returns the NotificationString field value if set, zero value otherwise.
func (o *RuleNotification) GetNotificationString() string {
	if o == nil || IsNil(o.NotificationString) {
		var ret string
		return ret
	}
	return *o.NotificationString
}

// GetNotificationStringOk returns a tuple with the NotificationString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleNotification) GetNotificationStringOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationString) {
		return nil, false
	}
	return o.NotificationString, true
}

// HasNotificationString returns a boolean if a field has been set.
func (o *RuleNotification) HasNotificationString() bool {
	if o != nil && !IsNil(o.NotificationString) {
		return true
	}

	return false
}

// SetNotificationString gets a reference to the given string and assigns it to the NotificationString field.
func (o *RuleNotification) SetNotificationString(v string) {
	o.NotificationString = &v
}

func (o RuleNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RuleState) {
		toSerialize["ruleState"] = o.RuleState
	}
	if !IsNil(o.CampaignOptimizationId) {
		toSerialize["campaignOptimizationId"] = o.CampaignOptimizationId
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	if !IsNil(o.NotificationString) {
		toSerialize["notificationString"] = o.NotificationString
	}
	return toSerialize, nil
}

type NullableRuleNotification struct {
	value *RuleNotification
	isSet bool
}

func (v NullableRuleNotification) Get() *RuleNotification {
	return v.value
}

func (v *NullableRuleNotification) Set(val *RuleNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleNotification(val *RuleNotification) *NullableRuleNotification {
	return &NullableRuleNotification{value: val, isSet: true}
}

func (v NullableRuleNotification) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRuleNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
