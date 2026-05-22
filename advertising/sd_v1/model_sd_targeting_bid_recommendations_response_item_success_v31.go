package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingBidRecommendationsResponseItemSuccessV31 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingBidRecommendationsResponseItemSuccessV31{}

// SDTargetingBidRecommendationsResponseItemSuccessV31 A recommended bid range to use for a target.
type SDTargetingBidRecommendationsResponseItemSuccessV31 struct {
	// The lowest recommended bid to use to win an ad placement for this target.
	RangeLower float32 `json:"rangeLower"`
	// The highest recommended bid to use to win an ad placement for this target.
	RangeUpper float32 `json:"rangeUpper"`
	// The recommended bid to use to win an ad placement for this target.
	Recommended float32 `json:"recommended"`
	// The HTTP status code of this item.
	Code string `json:"code"`
}

type _SDTargetingBidRecommendationsResponseItemSuccessV31 SDTargetingBidRecommendationsResponseItemSuccessV31

// NewSDTargetingBidRecommendationsResponseItemSuccessV31 instantiates a new SDTargetingBidRecommendationsResponseItemSuccessV31 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingBidRecommendationsResponseItemSuccessV31(rangeLower float32, rangeUpper float32, recommended float32, code string) *SDTargetingBidRecommendationsResponseItemSuccessV31 {
	this := SDTargetingBidRecommendationsResponseItemSuccessV31{}
	this.RangeLower = rangeLower
	this.RangeUpper = rangeUpper
	this.Recommended = recommended
	this.Code = code
	return &this
}

// NewSDTargetingBidRecommendationsResponseItemSuccessV31WithDefaults instantiates a new SDTargetingBidRecommendationsResponseItemSuccessV31 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingBidRecommendationsResponseItemSuccessV31WithDefaults() *SDTargetingBidRecommendationsResponseItemSuccessV31 {
	this := SDTargetingBidRecommendationsResponseItemSuccessV31{}
	return &this
}

// GetRangeLower returns the RangeLower field value
func (o *SDTargetingBidRecommendationsResponseItemSuccessV31) GetRangeLower() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RangeLower
}

// GetRangeLowerOk returns a tuple with the RangeLower field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsResponseItemSuccessV31) GetRangeLowerOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RangeLower, true
}

// SetRangeLower sets field value
func (o *SDTargetingBidRecommendationsResponseItemSuccessV31) SetRangeLower(v float32) {
	o.RangeLower = v
}

// GetRangeUpper returns the RangeUpper field value
func (o *SDTargetingBidRecommendationsResponseItemSuccessV31) GetRangeUpper() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RangeUpper
}

// GetRangeUpperOk returns a tuple with the RangeUpper field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsResponseItemSuccessV31) GetRangeUpperOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RangeUpper, true
}

// SetRangeUpper sets field value
func (o *SDTargetingBidRecommendationsResponseItemSuccessV31) SetRangeUpper(v float32) {
	o.RangeUpper = v
}

// GetRecommended returns the Recommended field value
func (o *SDTargetingBidRecommendationsResponseItemSuccessV31) GetRecommended() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Recommended
}

// GetRecommendedOk returns a tuple with the Recommended field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsResponseItemSuccessV31) GetRecommendedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recommended, true
}

// SetRecommended sets field value
func (o *SDTargetingBidRecommendationsResponseItemSuccessV31) SetRecommended(v float32) {
	o.Recommended = v
}

// GetCode returns the Code field value
func (o *SDTargetingBidRecommendationsResponseItemSuccessV31) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsResponseItemSuccessV31) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SDTargetingBidRecommendationsResponseItemSuccessV31) SetCode(v string) {
	o.Code = v
}

func (o SDTargetingBidRecommendationsResponseItemSuccessV31) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rangeLower"] = o.RangeLower
	toSerialize["rangeUpper"] = o.RangeUpper
	toSerialize["recommended"] = o.Recommended
	toSerialize["code"] = o.Code
	return toSerialize, nil
}

type NullableSDTargetingBidRecommendationsResponseItemSuccessV31 struct {
	value *SDTargetingBidRecommendationsResponseItemSuccessV31
	isSet bool
}

func (v NullableSDTargetingBidRecommendationsResponseItemSuccessV31) Get() *SDTargetingBidRecommendationsResponseItemSuccessV31 {
	return v.value
}

func (v *NullableSDTargetingBidRecommendationsResponseItemSuccessV31) Set(val *SDTargetingBidRecommendationsResponseItemSuccessV31) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingBidRecommendationsResponseItemSuccessV31) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingBidRecommendationsResponseItemSuccessV31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingBidRecommendationsResponseItemSuccessV31(val *SDTargetingBidRecommendationsResponseItemSuccessV31) *NullableSDTargetingBidRecommendationsResponseItemSuccessV31 {
	return &NullableSDTargetingBidRecommendationsResponseItemSuccessV31{value: val, isSet: true}
}

func (v NullableSDTargetingBidRecommendationsResponseItemSuccessV31) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingBidRecommendationsResponseItemSuccessV31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
