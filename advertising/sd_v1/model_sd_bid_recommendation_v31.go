package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDBidRecommendationV31 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDBidRecommendationV31{}

// SDBidRecommendationV31 A recommended bid range to use for a target.
type SDBidRecommendationV31 struct {
	// The lowest recommended bid to use to win an ad placement for this target.
	RangeLower float32 `json:"rangeLower"`
	// The highest recommended bid to use to win an ad placement for this target.
	RangeUpper float32 `json:"rangeUpper"`
	// The recommended bid to use to win an ad placement for this target.
	Recommended float32 `json:"recommended"`
}

type _SDBidRecommendationV31 SDBidRecommendationV31

// NewSDBidRecommendationV31 instantiates a new SDBidRecommendationV31 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDBidRecommendationV31(rangeLower float32, rangeUpper float32, recommended float32) *SDBidRecommendationV31 {
	this := SDBidRecommendationV31{}
	this.RangeLower = rangeLower
	this.RangeUpper = rangeUpper
	this.Recommended = recommended
	return &this
}

// NewSDBidRecommendationV31WithDefaults instantiates a new SDBidRecommendationV31 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDBidRecommendationV31WithDefaults() *SDBidRecommendationV31 {
	this := SDBidRecommendationV31{}
	return &this
}

// GetRangeLower returns the RangeLower field value
func (o *SDBidRecommendationV31) GetRangeLower() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RangeLower
}

// GetRangeLowerOk returns a tuple with the RangeLower field value
// and a boolean to check if the value has been set.
func (o *SDBidRecommendationV31) GetRangeLowerOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RangeLower, true
}

// SetRangeLower sets field value
func (o *SDBidRecommendationV31) SetRangeLower(v float32) {
	o.RangeLower = v
}

// GetRangeUpper returns the RangeUpper field value
func (o *SDBidRecommendationV31) GetRangeUpper() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RangeUpper
}

// GetRangeUpperOk returns a tuple with the RangeUpper field value
// and a boolean to check if the value has been set.
func (o *SDBidRecommendationV31) GetRangeUpperOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RangeUpper, true
}

// SetRangeUpper sets field value
func (o *SDBidRecommendationV31) SetRangeUpper(v float32) {
	o.RangeUpper = v
}

// GetRecommended returns the Recommended field value
func (o *SDBidRecommendationV31) GetRecommended() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Recommended
}

// GetRecommendedOk returns a tuple with the Recommended field value
// and a boolean to check if the value has been set.
func (o *SDBidRecommendationV31) GetRecommendedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recommended, true
}

// SetRecommended sets field value
func (o *SDBidRecommendationV31) SetRecommended(v float32) {
	o.Recommended = v
}

func (o SDBidRecommendationV31) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rangeLower"] = o.RangeLower
	toSerialize["rangeUpper"] = o.RangeUpper
	toSerialize["recommended"] = o.Recommended
	return toSerialize, nil
}

type NullableSDBidRecommendationV31 struct {
	value *SDBidRecommendationV31
	isSet bool
}

func (v NullableSDBidRecommendationV31) Get() *SDBidRecommendationV31 {
	return v.value
}

func (v *NullableSDBidRecommendationV31) Set(val *SDBidRecommendationV31) {
	v.value = val
	v.isSet = true
}

func (v NullableSDBidRecommendationV31) IsSet() bool {
	return v.isSet
}

func (v *NullableSDBidRecommendationV31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDBidRecommendationV31(val *SDBidRecommendationV31) *NullableSDBidRecommendationV31 {
	return &NullableSDBidRecommendationV31{value: val, isSet: true}
}

func (v NullableSDBidRecommendationV31) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDBidRecommendationV31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
