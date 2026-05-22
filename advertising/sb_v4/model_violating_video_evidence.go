package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ViolatingVideoEvidence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViolatingVideoEvidence{}

// ViolatingVideoEvidence struct for ViolatingVideoEvidence
type ViolatingVideoEvidence struct {
	ViolatingVideoPosition *VideoPosition `json:"violatingVideoPosition,omitempty"`
}

// NewViolatingVideoEvidence instantiates a new ViolatingVideoEvidence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViolatingVideoEvidence() *ViolatingVideoEvidence {
	this := ViolatingVideoEvidence{}
	return &this
}

// NewViolatingVideoEvidenceWithDefaults instantiates a new ViolatingVideoEvidence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViolatingVideoEvidenceWithDefaults() *ViolatingVideoEvidence {
	this := ViolatingVideoEvidence{}
	return &this
}

// GetViolatingVideoPosition returns the ViolatingVideoPosition field value if set, zero value otherwise.
func (o *ViolatingVideoEvidence) GetViolatingVideoPosition() VideoPosition {
	if o == nil || IsNil(o.ViolatingVideoPosition) {
		var ret VideoPosition
		return ret
	}
	return *o.ViolatingVideoPosition
}

// GetViolatingVideoPositionOk returns a tuple with the ViolatingVideoPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolatingVideoEvidence) GetViolatingVideoPositionOk() (*VideoPosition, bool) {
	if o == nil || IsNil(o.ViolatingVideoPosition) {
		return nil, false
	}
	return o.ViolatingVideoPosition, true
}

// HasViolatingVideoPosition returns a boolean if a field has been set.
func (o *ViolatingVideoEvidence) HasViolatingVideoPosition() bool {
	if o != nil && !IsNil(o.ViolatingVideoPosition) {
		return true
	}

	return false
}

// SetViolatingVideoPosition gets a reference to the given VideoPosition and assigns it to the ViolatingVideoPosition field.
func (o *ViolatingVideoEvidence) SetViolatingVideoPosition(v VideoPosition) {
	o.ViolatingVideoPosition = &v
}

func (o ViolatingVideoEvidence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ViolatingVideoPosition) {
		toSerialize["violatingVideoPosition"] = o.ViolatingVideoPosition
	}
	return toSerialize, nil
}

type NullableViolatingVideoEvidence struct {
	value *ViolatingVideoEvidence
	isSet bool
}

func (v NullableViolatingVideoEvidence) Get() *ViolatingVideoEvidence {
	return v.value
}

func (v *NullableViolatingVideoEvidence) Set(val *ViolatingVideoEvidence) {
	v.value = val
	v.isSet = true
}

func (v NullableViolatingVideoEvidence) IsSet() bool {
	return v.isSet
}

func (v *NullableViolatingVideoEvidence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViolatingVideoEvidence(val *ViolatingVideoEvidence) *NullableViolatingVideoEvidence {
	return &NullableViolatingVideoEvidence{value: val, isSet: true}
}

func (v NullableViolatingVideoEvidence) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableViolatingVideoEvidence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
