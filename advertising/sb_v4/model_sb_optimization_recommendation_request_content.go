package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBOptimizationRecommendationRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBOptimizationRecommendationRequestContent{}

// SBOptimizationRecommendationRequestContent struct for SBOptimizationRecommendationRequestContent
type SBOptimizationRecommendationRequestContent struct {
	CostControlMetric CostControlMetric `json:"costControlMetric"`
	LandingPages      []LandingPage     `json:"landingPages"`
}

type _SBOptimizationRecommendationRequestContent SBOptimizationRecommendationRequestContent

// NewSBOptimizationRecommendationRequestContent instantiates a new SBOptimizationRecommendationRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBOptimizationRecommendationRequestContent(costControlMetric CostControlMetric, landingPages []LandingPage) *SBOptimizationRecommendationRequestContent {
	this := SBOptimizationRecommendationRequestContent{}
	this.CostControlMetric = costControlMetric
	this.LandingPages = landingPages
	return &this
}

// NewSBOptimizationRecommendationRequestContentWithDefaults instantiates a new SBOptimizationRecommendationRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBOptimizationRecommendationRequestContentWithDefaults() *SBOptimizationRecommendationRequestContent {
	this := SBOptimizationRecommendationRequestContent{}
	return &this
}

// GetCostControlMetric returns the CostControlMetric field value
func (o *SBOptimizationRecommendationRequestContent) GetCostControlMetric() CostControlMetric {
	if o == nil {
		var ret CostControlMetric
		return ret
	}

	return o.CostControlMetric
}

// GetCostControlMetricOk returns a tuple with the CostControlMetric field value
// and a boolean to check if the value has been set.
func (o *SBOptimizationRecommendationRequestContent) GetCostControlMetricOk() (*CostControlMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CostControlMetric, true
}

// SetCostControlMetric sets field value
func (o *SBOptimizationRecommendationRequestContent) SetCostControlMetric(v CostControlMetric) {
	o.CostControlMetric = v
}

// GetLandingPages returns the LandingPages field value
func (o *SBOptimizationRecommendationRequestContent) GetLandingPages() []LandingPage {
	if o == nil {
		var ret []LandingPage
		return ret
	}

	return o.LandingPages
}

// GetLandingPagesOk returns a tuple with the LandingPages field value
// and a boolean to check if the value has been set.
func (o *SBOptimizationRecommendationRequestContent) GetLandingPagesOk() ([]LandingPage, bool) {
	if o == nil {
		return nil, false
	}
	return o.LandingPages, true
}

// SetLandingPages sets field value
func (o *SBOptimizationRecommendationRequestContent) SetLandingPages(v []LandingPage) {
	o.LandingPages = v
}

func (o SBOptimizationRecommendationRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["costControlMetric"] = o.CostControlMetric
	toSerialize["landingPages"] = o.LandingPages
	return toSerialize, nil
}

type NullableSBOptimizationRecommendationRequestContent struct {
	value *SBOptimizationRecommendationRequestContent
	isSet bool
}

func (v NullableSBOptimizationRecommendationRequestContent) Get() *SBOptimizationRecommendationRequestContent {
	return v.value
}

func (v *NullableSBOptimizationRecommendationRequestContent) Set(val *SBOptimizationRecommendationRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBOptimizationRecommendationRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBOptimizationRecommendationRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBOptimizationRecommendationRequestContent(val *SBOptimizationRecommendationRequestContent) *NullableSBOptimizationRecommendationRequestContent {
	return &NullableSBOptimizationRecommendationRequestContent{value: val, isSet: true}
}

func (v NullableSBOptimizationRecommendationRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBOptimizationRecommendationRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
