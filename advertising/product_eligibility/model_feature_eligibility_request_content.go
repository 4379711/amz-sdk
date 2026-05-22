package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the FeatureEligibilityRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureEligibilityRequestContent{}

// FeatureEligibilityRequestContent A request to evaluate feature eligibility
type FeatureEligibilityRequestContent struct {
	// The List of features and marketplaces of which you wish the feature to be evalulated in
	Features []Feature `json:"features"`
}

type _FeatureEligibilityRequestContent FeatureEligibilityRequestContent

// NewFeatureEligibilityRequestContent instantiates a new FeatureEligibilityRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureEligibilityRequestContent(features []Feature) *FeatureEligibilityRequestContent {
	this := FeatureEligibilityRequestContent{}
	this.Features = features
	return &this
}

// NewFeatureEligibilityRequestContentWithDefaults instantiates a new FeatureEligibilityRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureEligibilityRequestContentWithDefaults() *FeatureEligibilityRequestContent {
	this := FeatureEligibilityRequestContent{}
	return &this
}

// GetFeatures returns the Features field value
func (o *FeatureEligibilityRequestContent) GetFeatures() []Feature {
	if o == nil {
		var ret []Feature
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *FeatureEligibilityRequestContent) GetFeaturesOk() ([]Feature, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *FeatureEligibilityRequestContent) SetFeatures(v []Feature) {
	o.Features = v
}

func (o FeatureEligibilityRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["features"] = o.Features
	return toSerialize, nil
}

type NullableFeatureEligibilityRequestContent struct {
	value *FeatureEligibilityRequestContent
	isSet bool
}

func (v NullableFeatureEligibilityRequestContent) Get() *FeatureEligibilityRequestContent {
	return v.value
}

func (v *NullableFeatureEligibilityRequestContent) Set(val *FeatureEligibilityRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureEligibilityRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureEligibilityRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureEligibilityRequestContent(val *FeatureEligibilityRequestContent) *NullableFeatureEligibilityRequestContent {
	return &NullableFeatureEligibilityRequestContent{value: val, isSet: true}
}

func (v NullableFeatureEligibilityRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeatureEligibilityRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
