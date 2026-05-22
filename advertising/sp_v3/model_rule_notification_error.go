package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the RuleNotificationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleNotificationError{}

// RuleNotificationError struct for RuleNotificationError
type RuleNotificationError struct {
	// campaignId
	CampaignId *string                        `json:"campaignId,omitempty"`
	Error      *CampaignOptimizationRuleError `json:"Error,omitempty"`
}

// NewRuleNotificationError instantiates a new RuleNotificationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleNotificationError() *RuleNotificationError {
	this := RuleNotificationError{}
	return &this
}

// NewRuleNotificationErrorWithDefaults instantiates a new RuleNotificationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleNotificationErrorWithDefaults() *RuleNotificationError {
	this := RuleNotificationError{}
	return &this
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *RuleNotificationError) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleNotificationError) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *RuleNotificationError) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *RuleNotificationError) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RuleNotificationError) GetError() CampaignOptimizationRuleError {
	if o == nil || IsNil(o.Error) {
		var ret CampaignOptimizationRuleError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleNotificationError) GetErrorOk() (*CampaignOptimizationRuleError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RuleNotificationError) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given CampaignOptimizationRuleError and assigns it to the Error field.
func (o *RuleNotificationError) SetError(v CampaignOptimizationRuleError) {
	o.Error = &v
}

func (o RuleNotificationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	if !IsNil(o.Error) {
		toSerialize["Error"] = o.Error
	}
	return toSerialize, nil
}

type NullableRuleNotificationError struct {
	value *RuleNotificationError
	isSet bool
}

func (v NullableRuleNotificationError) Get() *RuleNotificationError {
	return v.value
}

func (v *NullableRuleNotificationError) Set(val *RuleNotificationError) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleNotificationError) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleNotificationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleNotificationError(val *RuleNotificationError) *NullableRuleNotificationError {
	return &NullableRuleNotificationError{value: val, isSet: true}
}

func (v NullableRuleNotificationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRuleNotificationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
