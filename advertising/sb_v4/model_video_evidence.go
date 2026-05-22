package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the VideoEvidence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoEvidence{}

// VideoEvidence Structure of a video evidence
type VideoEvidence struct {
	// The start position (in seconds) of the content that violates the specified policy within the video.
	Start *int32 `json:"start,omitempty"`
	// The end position (in seconds) of the content that violates the specified policy within the video.
	End *int32 `json:"end,omitempty"`
}

// NewVideoEvidence instantiates a new VideoEvidence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoEvidence() *VideoEvidence {
	this := VideoEvidence{}
	return &this
}

// NewVideoEvidenceWithDefaults instantiates a new VideoEvidence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoEvidenceWithDefaults() *VideoEvidence {
	this := VideoEvidence{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *VideoEvidence) GetStart() int32 {
	if o == nil || IsNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoEvidence) GetStartOk() (*int32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *VideoEvidence) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *VideoEvidence) SetStart(v int32) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *VideoEvidence) GetEnd() int32 {
	if o == nil || IsNil(o.End) {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoEvidence) GetEndOk() (*int32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *VideoEvidence) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *VideoEvidence) SetEnd(v int32) {
	o.End = &v
}

func (o VideoEvidence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableVideoEvidence struct {
	value *VideoEvidence
	isSet bool
}

func (v NullableVideoEvidence) Get() *VideoEvidence {
	return v.value
}

func (v *NullableVideoEvidence) Set(val *VideoEvidence) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoEvidence) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoEvidence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoEvidence(val *VideoEvidence) *NullableVideoEvidence {
	return &NullableVideoEvidence{value: val, isSet: true}
}

func (v NullableVideoEvidence) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableVideoEvidence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
