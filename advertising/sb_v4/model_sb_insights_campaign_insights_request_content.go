package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBInsightsCampaignInsightsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBInsightsCampaignInsightsRequestContent{}

// SBInsightsCampaignInsightsRequestContent struct for SBInsightsCampaignInsightsRequestContent
type SBInsightsCampaignInsightsRequestContent struct {
	AdGroups []SBInsightsAdGroup `json:"adGroups"`
}

type _SBInsightsCampaignInsightsRequestContent SBInsightsCampaignInsightsRequestContent

// NewSBInsightsCampaignInsightsRequestContent instantiates a new SBInsightsCampaignInsightsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBInsightsCampaignInsightsRequestContent(adGroups []SBInsightsAdGroup) *SBInsightsCampaignInsightsRequestContent {
	this := SBInsightsCampaignInsightsRequestContent{}
	this.AdGroups = adGroups
	return &this
}

// NewSBInsightsCampaignInsightsRequestContentWithDefaults instantiates a new SBInsightsCampaignInsightsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBInsightsCampaignInsightsRequestContentWithDefaults() *SBInsightsCampaignInsightsRequestContent {
	this := SBInsightsCampaignInsightsRequestContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value
func (o *SBInsightsCampaignInsightsRequestContent) GetAdGroups() []SBInsightsAdGroup {
	if o == nil {
		var ret []SBInsightsAdGroup
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *SBInsightsCampaignInsightsRequestContent) GetAdGroupsOk() ([]SBInsightsAdGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdGroups, true
}

// SetAdGroups sets field value
func (o *SBInsightsCampaignInsightsRequestContent) SetAdGroups(v []SBInsightsAdGroup) {
	o.AdGroups = v
}

func (o SBInsightsCampaignInsightsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroups"] = o.AdGroups
	return toSerialize, nil
}

type NullableSBInsightsCampaignInsightsRequestContent struct {
	value *SBInsightsCampaignInsightsRequestContent
	isSet bool
}

func (v NullableSBInsightsCampaignInsightsRequestContent) Get() *SBInsightsCampaignInsightsRequestContent {
	return v.value
}

func (v *NullableSBInsightsCampaignInsightsRequestContent) Set(val *SBInsightsCampaignInsightsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBInsightsCampaignInsightsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBInsightsCampaignInsightsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBInsightsCampaignInsightsRequestContent(val *SBInsightsCampaignInsightsRequestContent) *NullableSBInsightsCampaignInsightsRequestContent {
	return &NullableSBInsightsCampaignInsightsRequestContent{value: val, isSet: true}
}

func (v NullableSBInsightsCampaignInsightsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBInsightsCampaignInsightsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
