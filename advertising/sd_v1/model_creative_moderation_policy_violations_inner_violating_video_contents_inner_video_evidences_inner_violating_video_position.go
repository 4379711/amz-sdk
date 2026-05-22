package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition{}

// CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition struct for CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition
type CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition struct {
	// Time at which policy violation within video asset starts.
	Start *int32 `json:"start,omitempty"`
	// Time at which policy violation within the video asset ends.
	End *int32 `json:"end,omitempty"`
}

// NewCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition instantiates a new CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition() *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition {
	this := CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition{}
	return &this
}

// NewCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPositionWithDefaults instantiates a new CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPositionWithDefaults() *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition {
	this := CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition) GetStart() int32 {
	if o == nil || IsNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition) GetStartOk() (*int32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition) SetStart(v int32) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition) GetEnd() int32 {
	if o == nil || IsNil(o.End) {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition) GetEndOk() (*int32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition) SetEnd(v int32) {
	o.End = &v
}

func (o CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition struct {
	value *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition
	isSet bool
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition) Get() *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition {
	return v.value
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition) Set(val *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition(val *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition) *NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition {
	return &NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition{value: val, isSet: true}
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
