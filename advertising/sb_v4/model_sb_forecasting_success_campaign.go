package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingSuccessCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingSuccessCampaign{}

// SBForecastingSuccessCampaign struct for SBForecastingSuccessCampaign
type SBForecastingSuccessCampaign struct {
	Forecasts []SBForecastingMetric `json:"forecasts,omitempty"`
	// The forecast timestamp.
	ForecastTimestamp *string `json:"forecastTimestamp,omitempty"`
}

// NewSBForecastingSuccessCampaign instantiates a new SBForecastingSuccessCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingSuccessCampaign() *SBForecastingSuccessCampaign {
	this := SBForecastingSuccessCampaign{}
	return &this
}

// NewSBForecastingSuccessCampaignWithDefaults instantiates a new SBForecastingSuccessCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingSuccessCampaignWithDefaults() *SBForecastingSuccessCampaign {
	this := SBForecastingSuccessCampaign{}
	return &this
}

// GetForecasts returns the Forecasts field value if set, zero value otherwise.
func (o *SBForecastingSuccessCampaign) GetForecasts() []SBForecastingMetric {
	if o == nil || IsNil(o.Forecasts) {
		var ret []SBForecastingMetric
		return ret
	}
	return o.Forecasts
}

// GetForecastsOk returns a tuple with the Forecasts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingSuccessCampaign) GetForecastsOk() ([]SBForecastingMetric, bool) {
	if o == nil || IsNil(o.Forecasts) {
		return nil, false
	}
	return o.Forecasts, true
}

// HasForecasts returns a boolean if a field has been set.
func (o *SBForecastingSuccessCampaign) HasForecasts() bool {
	if o != nil && !IsNil(o.Forecasts) {
		return true
	}

	return false
}

// SetForecasts gets a reference to the given []SBForecastingMetric and assigns it to the Forecasts field.
func (o *SBForecastingSuccessCampaign) SetForecasts(v []SBForecastingMetric) {
	o.Forecasts = v
}

// GetForecastTimestamp returns the ForecastTimestamp field value if set, zero value otherwise.
func (o *SBForecastingSuccessCampaign) GetForecastTimestamp() string {
	if o == nil || IsNil(o.ForecastTimestamp) {
		var ret string
		return ret
	}
	return *o.ForecastTimestamp
}

// GetForecastTimestampOk returns a tuple with the ForecastTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingSuccessCampaign) GetForecastTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.ForecastTimestamp) {
		return nil, false
	}
	return o.ForecastTimestamp, true
}

// HasForecastTimestamp returns a boolean if a field has been set.
func (o *SBForecastingSuccessCampaign) HasForecastTimestamp() bool {
	if o != nil && !IsNil(o.ForecastTimestamp) {
		return true
	}

	return false
}

// SetForecastTimestamp gets a reference to the given string and assigns it to the ForecastTimestamp field.
func (o *SBForecastingSuccessCampaign) SetForecastTimestamp(v string) {
	o.ForecastTimestamp = &v
}

func (o SBForecastingSuccessCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Forecasts) {
		toSerialize["forecasts"] = o.Forecasts
	}
	if !IsNil(o.ForecastTimestamp) {
		toSerialize["forecastTimestamp"] = o.ForecastTimestamp
	}
	return toSerialize, nil
}

type NullableSBForecastingSuccessCampaign struct {
	value *SBForecastingSuccessCampaign
	isSet bool
}

func (v NullableSBForecastingSuccessCampaign) Get() *SBForecastingSuccessCampaign {
	return v.value
}

func (v *NullableSBForecastingSuccessCampaign) Set(val *SBForecastingSuccessCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingSuccessCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingSuccessCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingSuccessCampaign(val *SBForecastingSuccessCampaign) *NullableSBForecastingSuccessCampaign {
	return &NullableSBForecastingSuccessCampaign{value: val, isSet: true}
}

func (v NullableSBForecastingSuccessCampaign) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingSuccessCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
