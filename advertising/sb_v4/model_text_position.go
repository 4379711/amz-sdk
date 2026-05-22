package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the TextPosition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextPosition{}

// TextPosition struct for TextPosition
type TextPosition struct {
	// Zero-based index into the text in reviewedText where the text specified in violatingText starts.
	Start *int64 `json:"start,omitempty"`
	// Zero-based index into the text in reviewedText where the text specified in violatingText ends.
	End *int64 `json:"end,omitempty"`
}

// NewTextPosition instantiates a new TextPosition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextPosition() *TextPosition {
	this := TextPosition{}
	return &this
}

// NewTextPositionWithDefaults instantiates a new TextPosition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextPositionWithDefaults() *TextPosition {
	this := TextPosition{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *TextPosition) GetStart() int64 {
	if o == nil || IsNil(o.Start) {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextPosition) GetStartOk() (*int64, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *TextPosition) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *TextPosition) SetStart(v int64) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *TextPosition) GetEnd() int64 {
	if o == nil || IsNil(o.End) {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextPosition) GetEndOk() (*int64, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *TextPosition) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *TextPosition) SetEnd(v int64) {
	o.End = &v
}

func (o TextPosition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableTextPosition struct {
	value *TextPosition
	isSet bool
}

func (v NullableTextPosition) Get() *TextPosition {
	return v.value
}

func (v *NullableTextPosition) Set(val *TextPosition) {
	v.value = val
	v.isSet = true
}

func (v NullableTextPosition) IsSet() bool {
	return v.isSet
}

func (v *NullableTextPosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextPosition(val *TextPosition) *NullableTextPosition {
	return &NullableTextPosition{value: val, isSet: true}
}

func (v NullableTextPosition) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTextPosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
