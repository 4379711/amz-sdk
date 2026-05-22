package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the TextEvidence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextEvidence{}

// TextEvidence Structure of a text evidence
type TextEvidence struct {
	// The specific text determined to violate the specified policy in reviewedText.
	ViolatingText *string               `json:"violatingText,omitempty"`
	Position      *TextEvidencePosition `json:"position,omitempty"`
}

// NewTextEvidence instantiates a new TextEvidence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextEvidence() *TextEvidence {
	this := TextEvidence{}
	return &this
}

// NewTextEvidenceWithDefaults instantiates a new TextEvidence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextEvidenceWithDefaults() *TextEvidence {
	this := TextEvidence{}
	return &this
}

// GetViolatingText returns the ViolatingText field value if set, zero value otherwise.
func (o *TextEvidence) GetViolatingText() string {
	if o == nil || IsNil(o.ViolatingText) {
		var ret string
		return ret
	}
	return *o.ViolatingText
}

// GetViolatingTextOk returns a tuple with the ViolatingText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextEvidence) GetViolatingTextOk() (*string, bool) {
	if o == nil || IsNil(o.ViolatingText) {
		return nil, false
	}
	return o.ViolatingText, true
}

// HasViolatingText returns a boolean if a field has been set.
func (o *TextEvidence) HasViolatingText() bool {
	if o != nil && !IsNil(o.ViolatingText) {
		return true
	}

	return false
}

// SetViolatingText gets a reference to the given string and assigns it to the ViolatingText field.
func (o *TextEvidence) SetViolatingText(v string) {
	o.ViolatingText = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *TextEvidence) GetPosition() TextEvidencePosition {
	if o == nil || IsNil(o.Position) {
		var ret TextEvidencePosition
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextEvidence) GetPositionOk() (*TextEvidencePosition, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *TextEvidence) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given TextEvidencePosition and assigns it to the Position field.
func (o *TextEvidence) SetPosition(v TextEvidencePosition) {
	o.Position = &v
}

func (o TextEvidence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ViolatingText) {
		toSerialize["violatingText"] = o.ViolatingText
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	return toSerialize, nil
}

type NullableTextEvidence struct {
	value *TextEvidence
	isSet bool
}

func (v NullableTextEvidence) Get() *TextEvidence {
	return v.value
}

func (v *NullableTextEvidence) Set(val *TextEvidence) {
	v.value = val
	v.isSet = true
}

func (v NullableTextEvidence) IsSet() bool {
	return v.isSet
}

func (v *NullableTextEvidence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextEvidence(val *TextEvidence) *NullableTextEvidence {
	return &NullableTextEvidence{value: val, isSet: true}
}

func (v NullableTextEvidence) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTextEvidence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
