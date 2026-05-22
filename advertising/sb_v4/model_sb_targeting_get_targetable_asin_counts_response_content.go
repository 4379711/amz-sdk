package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingGetTargetableASINCountsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingGetTargetableASINCountsResponseContent{}

// SBTargetingGetTargetableASINCountsResponseContent Response object for /sb/targets/products/count to get number of targetable asins for refinements provided by the user
type SBTargetingGetTargetableASINCountsResponseContent struct {
	AsinCounts *SBTargetingIntegerRange `json:"asinCounts,omitempty"`
}

// NewSBTargetingGetTargetableASINCountsResponseContent instantiates a new SBTargetingGetTargetableASINCountsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingGetTargetableASINCountsResponseContent() *SBTargetingGetTargetableASINCountsResponseContent {
	this := SBTargetingGetTargetableASINCountsResponseContent{}
	return &this
}

// NewSBTargetingGetTargetableASINCountsResponseContentWithDefaults instantiates a new SBTargetingGetTargetableASINCountsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingGetTargetableASINCountsResponseContentWithDefaults() *SBTargetingGetTargetableASINCountsResponseContent {
	this := SBTargetingGetTargetableASINCountsResponseContent{}
	return &this
}

// GetAsinCounts returns the AsinCounts field value if set, zero value otherwise.
func (o *SBTargetingGetTargetableASINCountsResponseContent) GetAsinCounts() SBTargetingIntegerRange {
	if o == nil || IsNil(o.AsinCounts) {
		var ret SBTargetingIntegerRange
		return ret
	}
	return *o.AsinCounts
}

// GetAsinCountsOk returns a tuple with the AsinCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingGetTargetableASINCountsResponseContent) GetAsinCountsOk() (*SBTargetingIntegerRange, bool) {
	if o == nil || IsNil(o.AsinCounts) {
		return nil, false
	}
	return o.AsinCounts, true
}

// HasAsinCounts returns a boolean if a field has been set.
func (o *SBTargetingGetTargetableASINCountsResponseContent) HasAsinCounts() bool {
	if o != nil && !IsNil(o.AsinCounts) {
		return true
	}

	return false
}

// SetAsinCounts gets a reference to the given SBTargetingIntegerRange and assigns it to the AsinCounts field.
func (o *SBTargetingGetTargetableASINCountsResponseContent) SetAsinCounts(v SBTargetingIntegerRange) {
	o.AsinCounts = &v
}

func (o SBTargetingGetTargetableASINCountsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AsinCounts) {
		toSerialize["asinCounts"] = o.AsinCounts
	}
	return toSerialize, nil
}

type NullableSBTargetingGetTargetableASINCountsResponseContent struct {
	value *SBTargetingGetTargetableASINCountsResponseContent
	isSet bool
}

func (v NullableSBTargetingGetTargetableASINCountsResponseContent) Get() *SBTargetingGetTargetableASINCountsResponseContent {
	return v.value
}

func (v *NullableSBTargetingGetTargetableASINCountsResponseContent) Set(val *SBTargetingGetTargetableASINCountsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingGetTargetableASINCountsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingGetTargetableASINCountsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingGetTargetableASINCountsResponseContent(val *SBTargetingGetTargetableASINCountsResponseContent) *NullableSBTargetingGetTargetableASINCountsResponseContent {
	return &NullableSBTargetingGetTargetableASINCountsResponseContent{value: val, isSet: true}
}

func (v NullableSBTargetingGetTargetableASINCountsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingGetTargetableASINCountsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
