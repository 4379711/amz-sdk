package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBCampaignPerformanceForecastsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBCampaignPerformanceForecastsRequestContent{}

// SBCampaignPerformanceForecastsRequestContent struct for SBCampaignPerformanceForecastsRequestContent
type SBCampaignPerformanceForecastsRequestContent struct {
	Campaigns []SBForecastingRequestCampaignObject `json:"campaigns"`
}

type _SBCampaignPerformanceForecastsRequestContent SBCampaignPerformanceForecastsRequestContent

// NewSBCampaignPerformanceForecastsRequestContent instantiates a new SBCampaignPerformanceForecastsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBCampaignPerformanceForecastsRequestContent(campaigns []SBForecastingRequestCampaignObject) *SBCampaignPerformanceForecastsRequestContent {
	this := SBCampaignPerformanceForecastsRequestContent{}
	this.Campaigns = campaigns
	return &this
}

// NewSBCampaignPerformanceForecastsRequestContentWithDefaults instantiates a new SBCampaignPerformanceForecastsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBCampaignPerformanceForecastsRequestContentWithDefaults() *SBCampaignPerformanceForecastsRequestContent {
	this := SBCampaignPerformanceForecastsRequestContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *SBCampaignPerformanceForecastsRequestContent) GetCampaigns() []SBForecastingRequestCampaignObject {
	if o == nil {
		var ret []SBForecastingRequestCampaignObject
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *SBCampaignPerformanceForecastsRequestContent) GetCampaignsOk() ([]SBForecastingRequestCampaignObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Campaigns, true
}

// SetCampaigns sets field value
func (o *SBCampaignPerformanceForecastsRequestContent) SetCampaigns(v []SBForecastingRequestCampaignObject) {
	o.Campaigns = v
}

func (o SBCampaignPerformanceForecastsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableSBCampaignPerformanceForecastsRequestContent struct {
	value *SBCampaignPerformanceForecastsRequestContent
	isSet bool
}

func (v NullableSBCampaignPerformanceForecastsRequestContent) Get() *SBCampaignPerformanceForecastsRequestContent {
	return v.value
}

func (v *NullableSBCampaignPerformanceForecastsRequestContent) Set(val *SBCampaignPerformanceForecastsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBCampaignPerformanceForecastsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBCampaignPerformanceForecastsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBCampaignPerformanceForecastsRequestContent(val *SBCampaignPerformanceForecastsRequestContent) *NullableSBCampaignPerformanceForecastsRequestContent {
	return &NullableSBCampaignPerformanceForecastsRequestContent{value: val, isSet: true}
}

func (v NullableSBCampaignPerformanceForecastsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBCampaignPerformanceForecastsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
