package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ViolatingTextEvidence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViolatingTextEvidence{}

// ViolatingTextEvidence struct for ViolatingTextEvidence
type ViolatingTextEvidence struct {
	ViolatingTextPosition *TextPosition `json:"violatingTextPosition,omitempty"`
	// The specific text determined to violate the specified policy in reviewedText.
	ViolatingText *string `json:"violatingText,omitempty"`
}

// NewViolatingTextEvidence instantiates a new ViolatingTextEvidence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViolatingTextEvidence() *ViolatingTextEvidence {
	this := ViolatingTextEvidence{}
	return &this
}

// NewViolatingTextEvidenceWithDefaults instantiates a new ViolatingTextEvidence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViolatingTextEvidenceWithDefaults() *ViolatingTextEvidence {
	this := ViolatingTextEvidence{}
	return &this
}

// GetViolatingTextPosition returns the ViolatingTextPosition field value if set, zero value otherwise.
func (o *ViolatingTextEvidence) GetViolatingTextPosition() TextPosition {
	if o == nil || IsNil(o.ViolatingTextPosition) {
		var ret TextPosition
		return ret
	}
	return *o.ViolatingTextPosition
}

// GetViolatingTextPositionOk returns a tuple with the ViolatingTextPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolatingTextEvidence) GetViolatingTextPositionOk() (*TextPosition, bool) {
	if o == nil || IsNil(o.ViolatingTextPosition) {
		return nil, false
	}
	return o.ViolatingTextPosition, true
}

// HasViolatingTextPosition returns a boolean if a field has been set.
func (o *ViolatingTextEvidence) HasViolatingTextPosition() bool {
	if o != nil && !IsNil(o.ViolatingTextPosition) {
		return true
	}

	return false
}

// SetViolatingTextPosition gets a reference to the given TextPosition and assigns it to the ViolatingTextPosition field.
func (o *ViolatingTextEvidence) SetViolatingTextPosition(v TextPosition) {
	o.ViolatingTextPosition = &v
}

// GetViolatingText returns the ViolatingText field value if set, zero value otherwise.
func (o *ViolatingTextEvidence) GetViolatingText() string {
	if o == nil || IsNil(o.ViolatingText) {
		var ret string
		return ret
	}
	return *o.ViolatingText
}

// GetViolatingTextOk returns a tuple with the ViolatingText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolatingTextEvidence) GetViolatingTextOk() (*string, bool) {
	if o == nil || IsNil(o.ViolatingText) {
		return nil, false
	}
	return o.ViolatingText, true
}

// HasViolatingText returns a boolean if a field has been set.
func (o *ViolatingTextEvidence) HasViolatingText() bool {
	if o != nil && !IsNil(o.ViolatingText) {
		return true
	}

	return false
}

// SetViolatingText gets a reference to the given string and assigns it to the ViolatingText field.
func (o *ViolatingTextEvidence) SetViolatingText(v string) {
	o.ViolatingText = &v
}

func (o ViolatingTextEvidence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ViolatingTextPosition) {
		toSerialize["violatingTextPosition"] = o.ViolatingTextPosition
	}
	if !IsNil(o.ViolatingText) {
		toSerialize["violatingText"] = o.ViolatingText
	}
	return toSerialize, nil
}

type NullableViolatingTextEvidence struct {
	value *ViolatingTextEvidence
	isSet bool
}

func (v NullableViolatingTextEvidence) Get() *ViolatingTextEvidence {
	return v.value
}

func (v *NullableViolatingTextEvidence) Set(val *ViolatingTextEvidence) {
	v.value = val
	v.isSet = true
}

func (v NullableViolatingTextEvidence) IsSet() bool {
	return v.isSet
}

func (v *NullableViolatingTextEvidence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViolatingTextEvidence(val *ViolatingTextEvidence) *NullableViolatingTextEvidence {
	return &NullableViolatingTextEvidence{value: val, isSet: true}
}

func (v NullableViolatingTextEvidence) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableViolatingTextEvidence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
