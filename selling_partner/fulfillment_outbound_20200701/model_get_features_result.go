package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the GetFeaturesResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFeaturesResult{}

// GetFeaturesResult The payload for the `getFeatures` operation.
type GetFeaturesResult struct {
	// An array of features.
	Features []Feature `json:"features"`
}

type _GetFeaturesResult GetFeaturesResult

// NewGetFeaturesResult instantiates a new GetFeaturesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFeaturesResult(features []Feature) *GetFeaturesResult {
	this := GetFeaturesResult{}
	this.Features = features
	return &this
}

// NewGetFeaturesResultWithDefaults instantiates a new GetFeaturesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFeaturesResultWithDefaults() *GetFeaturesResult {
	this := GetFeaturesResult{}
	return &this
}

// GetFeatures returns the Features field value
func (o *GetFeaturesResult) GetFeatures() []Feature {
	if o == nil {
		var ret []Feature
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *GetFeaturesResult) GetFeaturesOk() ([]Feature, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *GetFeaturesResult) SetFeatures(v []Feature) {
	o.Features = v
}

func (o GetFeaturesResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["features"] = o.Features
	return toSerialize, nil
}

type NullableGetFeaturesResult struct {
	value *GetFeaturesResult
	isSet bool
}

func (v NullableGetFeaturesResult) Get() *GetFeaturesResult {
	return v.value
}

func (v *NullableGetFeaturesResult) Set(val *GetFeaturesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFeaturesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFeaturesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFeaturesResult(val *GetFeaturesResult) *NullableGetFeaturesResult {
	return &NullableGetFeaturesResult{value: val, isSet: true}
}

func (v NullableGetFeaturesResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetFeaturesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
