package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the AssociatedCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssociatedCampaign{}

// AssociatedCampaign struct for AssociatedCampaign
type AssociatedCampaign struct {
	// The campaign identifier.
	CampaignId string `json:"campaignId"`
	// The budget rule evaluation status for this campaign. Read-only.
	RuleStatus string `json:"ruleStatus"`
	// The campaign name.
	CampaignName string `json:"campaignName"`
}

type _AssociatedCampaign AssociatedCampaign

// NewAssociatedCampaign instantiates a new AssociatedCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociatedCampaign(campaignId string, ruleStatus string, campaignName string) *AssociatedCampaign {
	this := AssociatedCampaign{}
	this.CampaignId = campaignId
	this.RuleStatus = ruleStatus
	this.CampaignName = campaignName
	return &this
}

// NewAssociatedCampaignWithDefaults instantiates a new AssociatedCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociatedCampaignWithDefaults() *AssociatedCampaign {
	this := AssociatedCampaign{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *AssociatedCampaign) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *AssociatedCampaign) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *AssociatedCampaign) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetRuleStatus returns the RuleStatus field value
func (o *AssociatedCampaign) GetRuleStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleStatus
}

// GetRuleStatusOk returns a tuple with the RuleStatus field value
// and a boolean to check if the value has been set.
func (o *AssociatedCampaign) GetRuleStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleStatus, true
}

// SetRuleStatus sets field value
func (o *AssociatedCampaign) SetRuleStatus(v string) {
	o.RuleStatus = v
}

// GetCampaignName returns the CampaignName field value
func (o *AssociatedCampaign) GetCampaignName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignName
}

// GetCampaignNameOk returns a tuple with the CampaignName field value
// and a boolean to check if the value has been set.
func (o *AssociatedCampaign) GetCampaignNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignName, true
}

// SetCampaignName sets field value
func (o *AssociatedCampaign) SetCampaignName(v string) {
	o.CampaignName = v
}

func (o AssociatedCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["ruleStatus"] = o.RuleStatus
	toSerialize["campaignName"] = o.CampaignName
	return toSerialize, nil
}

type NullableAssociatedCampaign struct {
	value *AssociatedCampaign
	isSet bool
}

func (v NullableAssociatedCampaign) Get() *AssociatedCampaign {
	return v.value
}

func (v *NullableAssociatedCampaign) Set(val *AssociatedCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociatedCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociatedCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociatedCampaign(val *AssociatedCampaign) *NullableAssociatedCampaign {
	return &NullableAssociatedCampaign{value: val, isSet: true}
}

func (v NullableAssociatedCampaign) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAssociatedCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
