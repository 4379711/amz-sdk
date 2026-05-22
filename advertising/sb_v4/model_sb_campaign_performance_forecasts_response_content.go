package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBCampaignPerformanceForecastsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBCampaignPerformanceForecastsResponseContent{}

// SBCampaignPerformanceForecastsResponseContent Response object for /sb/forecasts containing a list of performance forecast for the campaign.
type SBCampaignPerformanceForecastsResponseContent struct {
	Campaigns *SBForecastingResponseCampaignObject `json:"campaigns,omitempty"`
}

// NewSBCampaignPerformanceForecastsResponseContent instantiates a new SBCampaignPerformanceForecastsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBCampaignPerformanceForecastsResponseContent() *SBCampaignPerformanceForecastsResponseContent {
	this := SBCampaignPerformanceForecastsResponseContent{}
	return &this
}

// NewSBCampaignPerformanceForecastsResponseContentWithDefaults instantiates a new SBCampaignPerformanceForecastsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBCampaignPerformanceForecastsResponseContentWithDefaults() *SBCampaignPerformanceForecastsResponseContent {
	this := SBCampaignPerformanceForecastsResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise.
func (o *SBCampaignPerformanceForecastsResponseContent) GetCampaigns() SBForecastingResponseCampaignObject {
	if o == nil || IsNil(o.Campaigns) {
		var ret SBForecastingResponseCampaignObject
		return ret
	}
	return *o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBCampaignPerformanceForecastsResponseContent) GetCampaignsOk() (*SBForecastingResponseCampaignObject, bool) {
	if o == nil || IsNil(o.Campaigns) {
		return nil, false
	}
	return o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *SBCampaignPerformanceForecastsResponseContent) HasCampaigns() bool {
	if o != nil && !IsNil(o.Campaigns) {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given SBForecastingResponseCampaignObject and assigns it to the Campaigns field.
func (o *SBCampaignPerformanceForecastsResponseContent) SetCampaigns(v SBForecastingResponseCampaignObject) {
	o.Campaigns = &v
}

func (o SBCampaignPerformanceForecastsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Campaigns) {
		toSerialize["campaigns"] = o.Campaigns
	}
	return toSerialize, nil
}

type NullableSBCampaignPerformanceForecastsResponseContent struct {
	value *SBCampaignPerformanceForecastsResponseContent
	isSet bool
}

func (v NullableSBCampaignPerformanceForecastsResponseContent) Get() *SBCampaignPerformanceForecastsResponseContent {
	return v.value
}

func (v *NullableSBCampaignPerformanceForecastsResponseContent) Set(val *SBCampaignPerformanceForecastsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBCampaignPerformanceForecastsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBCampaignPerformanceForecastsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBCampaignPerformanceForecastsResponseContent(val *SBCampaignPerformanceForecastsResponseContent) *NullableSBCampaignPerformanceForecastsResponseContent {
	return &NullableSBCampaignPerformanceForecastsResponseContent{value: val, isSet: true}
}

func (v NullableSBCampaignPerformanceForecastsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBCampaignPerformanceForecastsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
