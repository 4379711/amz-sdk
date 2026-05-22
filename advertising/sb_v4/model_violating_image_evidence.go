package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ViolatingImageEvidence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViolatingImageEvidence{}

// ViolatingImageEvidence struct for ViolatingImageEvidence
type ViolatingImageEvidence struct {
	ViolatingImageCrop *ImageCrop `json:"violatingImageCrop,omitempty"`
}

// NewViolatingImageEvidence instantiates a new ViolatingImageEvidence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViolatingImageEvidence() *ViolatingImageEvidence {
	this := ViolatingImageEvidence{}
	return &this
}

// NewViolatingImageEvidenceWithDefaults instantiates a new ViolatingImageEvidence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViolatingImageEvidenceWithDefaults() *ViolatingImageEvidence {
	this := ViolatingImageEvidence{}
	return &this
}

// GetViolatingImageCrop returns the ViolatingImageCrop field value if set, zero value otherwise.
func (o *ViolatingImageEvidence) GetViolatingImageCrop() ImageCrop {
	if o == nil || IsNil(o.ViolatingImageCrop) {
		var ret ImageCrop
		return ret
	}
	return *o.ViolatingImageCrop
}

// GetViolatingImageCropOk returns a tuple with the ViolatingImageCrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolatingImageEvidence) GetViolatingImageCropOk() (*ImageCrop, bool) {
	if o == nil || IsNil(o.ViolatingImageCrop) {
		return nil, false
	}
	return o.ViolatingImageCrop, true
}

// HasViolatingImageCrop returns a boolean if a field has been set.
func (o *ViolatingImageEvidence) HasViolatingImageCrop() bool {
	if o != nil && !IsNil(o.ViolatingImageCrop) {
		return true
	}

	return false
}

// SetViolatingImageCrop gets a reference to the given ImageCrop and assigns it to the ViolatingImageCrop field.
func (o *ViolatingImageEvidence) SetViolatingImageCrop(v ImageCrop) {
	o.ViolatingImageCrop = &v
}

func (o ViolatingImageEvidence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ViolatingImageCrop) {
		toSerialize["violatingImageCrop"] = o.ViolatingImageCrop
	}
	return toSerialize, nil
}

type NullableViolatingImageEvidence struct {
	value *ViolatingImageEvidence
	isSet bool
}

func (v NullableViolatingImageEvidence) Get() *ViolatingImageEvidence {
	return v.value
}

func (v *NullableViolatingImageEvidence) Set(val *ViolatingImageEvidence) {
	v.value = val
	v.isSet = true
}

func (v NullableViolatingImageEvidence) IsSet() bool {
	return v.isSet
}

func (v *NullableViolatingImageEvidence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViolatingImageEvidence(val *ViolatingImageEvidence) *NullableViolatingImageEvidence {
	return &NullableViolatingImageEvidence{value: val, isSet: true}
}

func (v NullableViolatingImageEvidence) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableViolatingImageEvidence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
