package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDCategoryRecommendationTargetableAsinCountRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDCategoryRecommendationTargetableAsinCountRange{}

// SDCategoryRecommendationTargetableAsinCountRange The range of ASINs available within the category catalogue. If no targetable ASIN counts are available then the targetableAsinCountRange value will be null without any properties.
type SDCategoryRecommendationTargetableAsinCountRange struct {
	RangeLower *int32 `json:"rangeLower,omitempty"`
	RangeUpper *int32 `json:"rangeUpper,omitempty"`
}

// NewSDCategoryRecommendationTargetableAsinCountRange instantiates a new SDCategoryRecommendationTargetableAsinCountRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDCategoryRecommendationTargetableAsinCountRange() *SDCategoryRecommendationTargetableAsinCountRange {
	this := SDCategoryRecommendationTargetableAsinCountRange{}
	return &this
}

// NewSDCategoryRecommendationTargetableAsinCountRangeWithDefaults instantiates a new SDCategoryRecommendationTargetableAsinCountRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDCategoryRecommendationTargetableAsinCountRangeWithDefaults() *SDCategoryRecommendationTargetableAsinCountRange {
	this := SDCategoryRecommendationTargetableAsinCountRange{}
	return &this
}

// GetRangeLower returns the RangeLower field value if set, zero value otherwise.
func (o *SDCategoryRecommendationTargetableAsinCountRange) GetRangeLower() int32 {
	if o == nil || IsNil(o.RangeLower) {
		var ret int32
		return ret
	}
	return *o.RangeLower
}

// GetRangeLowerOk returns a tuple with the RangeLower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDCategoryRecommendationTargetableAsinCountRange) GetRangeLowerOk() (*int32, bool) {
	if o == nil || IsNil(o.RangeLower) {
		return nil, false
	}
	return o.RangeLower, true
}

// HasRangeLower returns a boolean if a field has been set.
func (o *SDCategoryRecommendationTargetableAsinCountRange) HasRangeLower() bool {
	if o != nil && !IsNil(o.RangeLower) {
		return true
	}

	return false
}

// SetRangeLower gets a reference to the given int32 and assigns it to the RangeLower field.
func (o *SDCategoryRecommendationTargetableAsinCountRange) SetRangeLower(v int32) {
	o.RangeLower = &v
}

// GetRangeUpper returns the RangeUpper field value if set, zero value otherwise.
func (o *SDCategoryRecommendationTargetableAsinCountRange) GetRangeUpper() int32 {
	if o == nil || IsNil(o.RangeUpper) {
		var ret int32
		return ret
	}
	return *o.RangeUpper
}

// GetRangeUpperOk returns a tuple with the RangeUpper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDCategoryRecommendationTargetableAsinCountRange) GetRangeUpperOk() (*int32, bool) {
	if o == nil || IsNil(o.RangeUpper) {
		return nil, false
	}
	return o.RangeUpper, true
}

// HasRangeUpper returns a boolean if a field has been set.
func (o *SDCategoryRecommendationTargetableAsinCountRange) HasRangeUpper() bool {
	if o != nil && !IsNil(o.RangeUpper) {
		return true
	}

	return false
}

// SetRangeUpper gets a reference to the given int32 and assigns it to the RangeUpper field.
func (o *SDCategoryRecommendationTargetableAsinCountRange) SetRangeUpper(v int32) {
	o.RangeUpper = &v
}

func (o SDCategoryRecommendationTargetableAsinCountRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RangeLower) {
		toSerialize["rangeLower"] = o.RangeLower
	}
	if !IsNil(o.RangeUpper) {
		toSerialize["rangeUpper"] = o.RangeUpper
	}
	return toSerialize, nil
}

type NullableSDCategoryRecommendationTargetableAsinCountRange struct {
	value *SDCategoryRecommendationTargetableAsinCountRange
	isSet bool
}

func (v NullableSDCategoryRecommendationTargetableAsinCountRange) Get() *SDCategoryRecommendationTargetableAsinCountRange {
	return v.value
}

func (v *NullableSDCategoryRecommendationTargetableAsinCountRange) Set(val *SDCategoryRecommendationTargetableAsinCountRange) {
	v.value = val
	v.isSet = true
}

func (v NullableSDCategoryRecommendationTargetableAsinCountRange) IsSet() bool {
	return v.isSet
}

func (v *NullableSDCategoryRecommendationTargetableAsinCountRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDCategoryRecommendationTargetableAsinCountRange(val *SDCategoryRecommendationTargetableAsinCountRange) *NullableSDCategoryRecommendationTargetableAsinCountRange {
	return &NullableSDCategoryRecommendationTargetableAsinCountRange{value: val, isSet: true}
}

func (v NullableSDCategoryRecommendationTargetableAsinCountRange) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDCategoryRecommendationTargetableAsinCountRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
