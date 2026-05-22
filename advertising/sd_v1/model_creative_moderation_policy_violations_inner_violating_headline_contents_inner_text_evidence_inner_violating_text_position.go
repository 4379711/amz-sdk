package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition{}

// CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition struct for CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition
type CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition struct {
	// Zero-based index into the text in reviewedText where the text specified in violatingText starts
	Start *int32 `json:"start,omitempty"`
	// Zero-based index into the text in reviewedText where the text specified in violatingText ends
	End *int32 `json:"end,omitempty"`
}

// NewCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition instantiates a new CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition() *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition {
	this := CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition{}
	return &this
}

// NewCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPositionWithDefaults instantiates a new CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPositionWithDefaults() *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition {
	this := CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition) GetStart() int32 {
	if o == nil || IsNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition) GetStartOk() (*int32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition) SetStart(v int32) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition) GetEnd() int32 {
	if o == nil || IsNil(o.End) {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition) GetEndOk() (*int32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition) SetEnd(v int32) {
	o.End = &v
}

func (o CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition struct {
	value *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition
	isSet bool
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition) Get() *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition {
	return v.value
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition) Set(val *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition(val *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition) *NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition {
	return &NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition{value: val, isSet: true}
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInnerViolatingTextPosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
