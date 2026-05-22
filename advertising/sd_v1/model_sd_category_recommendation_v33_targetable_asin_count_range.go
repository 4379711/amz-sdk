package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDCategoryRecommendationV33TargetableAsinCountRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDCategoryRecommendationV33TargetableAsinCountRange{}

// SDCategoryRecommendationV33TargetableAsinCountRange The range of ASINs available within the category catalogue.
type SDCategoryRecommendationV33TargetableAsinCountRange struct {
	RangeLower *int32 `json:"rangeLower,omitempty"`
	RangeUpper *int32 `json:"rangeUpper,omitempty"`
}

// NewSDCategoryRecommendationV33TargetableAsinCountRange instantiates a new SDCategoryRecommendationV33TargetableAsinCountRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDCategoryRecommendationV33TargetableAsinCountRange() *SDCategoryRecommendationV33TargetableAsinCountRange {
	this := SDCategoryRecommendationV33TargetableAsinCountRange{}
	return &this
}

// NewSDCategoryRecommendationV33TargetableAsinCountRangeWithDefaults instantiates a new SDCategoryRecommendationV33TargetableAsinCountRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDCategoryRecommendationV33TargetableAsinCountRangeWithDefaults() *SDCategoryRecommendationV33TargetableAsinCountRange {
	this := SDCategoryRecommendationV33TargetableAsinCountRange{}
	return &this
}

// GetRangeLower returns the RangeLower field value if set, zero value otherwise.
func (o *SDCategoryRecommendationV33TargetableAsinCountRange) GetRangeLower() int32 {
	if o == nil || IsNil(o.RangeLower) {
		var ret int32
		return ret
	}
	return *o.RangeLower
}

// GetRangeLowerOk returns a tuple with the RangeLower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDCategoryRecommendationV33TargetableAsinCountRange) GetRangeLowerOk() (*int32, bool) {
	if o == nil || IsNil(o.RangeLower) {
		return nil, false
	}
	return o.RangeLower, true
}

// HasRangeLower returns a boolean if a field has been set.
func (o *SDCategoryRecommendationV33TargetableAsinCountRange) HasRangeLower() bool {
	if o != nil && !IsNil(o.RangeLower) {
		return true
	}

	return false
}

// SetRangeLower gets a reference to the given int32 and assigns it to the RangeLower field.
func (o *SDCategoryRecommendationV33TargetableAsinCountRange) SetRangeLower(v int32) {
	o.RangeLower = &v
}

// GetRangeUpper returns the RangeUpper field value if set, zero value otherwise.
func (o *SDCategoryRecommendationV33TargetableAsinCountRange) GetRangeUpper() int32 {
	if o == nil || IsNil(o.RangeUpper) {
		var ret int32
		return ret
	}
	return *o.RangeUpper
}

// GetRangeUpperOk returns a tuple with the RangeUpper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDCategoryRecommendationV33TargetableAsinCountRange) GetRangeUpperOk() (*int32, bool) {
	if o == nil || IsNil(o.RangeUpper) {
		return nil, false
	}
	return o.RangeUpper, true
}

// HasRangeUpper returns a boolean if a field has been set.
func (o *SDCategoryRecommendationV33TargetableAsinCountRange) HasRangeUpper() bool {
	if o != nil && !IsNil(o.RangeUpper) {
		return true
	}

	return false
}

// SetRangeUpper gets a reference to the given int32 and assigns it to the RangeUpper field.
func (o *SDCategoryRecommendationV33TargetableAsinCountRange) SetRangeUpper(v int32) {
	o.RangeUpper = &v
}

func (o SDCategoryRecommendationV33TargetableAsinCountRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RangeLower) {
		toSerialize["rangeLower"] = o.RangeLower
	}
	if !IsNil(o.RangeUpper) {
		toSerialize["rangeUpper"] = o.RangeUpper
	}
	return toSerialize, nil
}

type NullableSDCategoryRecommendationV33TargetableAsinCountRange struct {
	value *SDCategoryRecommendationV33TargetableAsinCountRange
	isSet bool
}

func (v NullableSDCategoryRecommendationV33TargetableAsinCountRange) Get() *SDCategoryRecommendationV33TargetableAsinCountRange {
	return v.value
}

func (v *NullableSDCategoryRecommendationV33TargetableAsinCountRange) Set(val *SDCategoryRecommendationV33TargetableAsinCountRange) {
	v.value = val
	v.isSet = true
}

func (v NullableSDCategoryRecommendationV33TargetableAsinCountRange) IsSet() bool {
	return v.isSet
}

func (v *NullableSDCategoryRecommendationV33TargetableAsinCountRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDCategoryRecommendationV33TargetableAsinCountRange(val *SDCategoryRecommendationV33TargetableAsinCountRange) *NullableSDCategoryRecommendationV33TargetableAsinCountRange {
	return &NullableSDCategoryRecommendationV33TargetableAsinCountRange{value: val, isSet: true}
}

func (v NullableSDCategoryRecommendationV33TargetableAsinCountRange) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDCategoryRecommendationV33TargetableAsinCountRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
