package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the VideoPosition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoPosition{}

// VideoPosition struct for VideoPosition
type VideoPosition struct {
	// Start time of the video having the policy violation.
	Start *int64 `json:"start,omitempty"`
	// End time of the video having the policy violation.
	End *int64 `json:"end,omitempty"`
}

// NewVideoPosition instantiates a new VideoPosition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoPosition() *VideoPosition {
	this := VideoPosition{}
	return &this
}

// NewVideoPositionWithDefaults instantiates a new VideoPosition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoPositionWithDefaults() *VideoPosition {
	this := VideoPosition{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *VideoPosition) GetStart() int64 {
	if o == nil || IsNil(o.Start) {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPosition) GetStartOk() (*int64, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *VideoPosition) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *VideoPosition) SetStart(v int64) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *VideoPosition) GetEnd() int64 {
	if o == nil || IsNil(o.End) {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPosition) GetEndOk() (*int64, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *VideoPosition) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *VideoPosition) SetEnd(v int64) {
	o.End = &v
}

func (o VideoPosition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableVideoPosition struct {
	value *VideoPosition
	isSet bool
}

func (v NullableVideoPosition) Get() *VideoPosition {
	return v.value
}

func (v *NullableVideoPosition) Set(val *VideoPosition) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoPosition) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoPosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoPosition(val *VideoPosition) *NullableVideoPosition {
	return &NullableVideoPosition{value: val, isSet: true}
}

func (v NullableVideoPosition) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableVideoPosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
