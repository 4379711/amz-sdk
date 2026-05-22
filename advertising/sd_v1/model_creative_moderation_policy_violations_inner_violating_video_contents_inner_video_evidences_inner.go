package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner{}

// CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner struct for CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner
type CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner struct {
	ViolatingVideoPosition *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition `json:"violatingVideoPosition,omitempty"`
}

// NewCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner instantiates a new CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner() *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner {
	this := CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner{}
	return &this
}

// NewCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerWithDefaults instantiates a new CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerWithDefaults() *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner {
	this := CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner{}
	return &this
}

// GetViolatingVideoPosition returns the ViolatingVideoPosition field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner) GetViolatingVideoPosition() CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition {
	if o == nil || IsNil(o.ViolatingVideoPosition) {
		var ret CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition
		return ret
	}
	return *o.ViolatingVideoPosition
}

// GetViolatingVideoPositionOk returns a tuple with the ViolatingVideoPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner) GetViolatingVideoPositionOk() (*CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition, bool) {
	if o == nil || IsNil(o.ViolatingVideoPosition) {
		return nil, false
	}
	return o.ViolatingVideoPosition, true
}

// HasViolatingVideoPosition returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner) HasViolatingVideoPosition() bool {
	if o != nil && !IsNil(o.ViolatingVideoPosition) {
		return true
	}

	return false
}

// SetViolatingVideoPosition gets a reference to the given CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition and assigns it to the ViolatingVideoPosition field.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner) SetViolatingVideoPosition(v CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInnerViolatingVideoPosition) {
	o.ViolatingVideoPosition = &v
}

func (o CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ViolatingVideoPosition) {
		toSerialize["violatingVideoPosition"] = o.ViolatingVideoPosition
	}
	return toSerialize, nil
}

type NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner struct {
	value *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner
	isSet bool
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner) Get() *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner {
	return v.value
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner) Set(val *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner(val *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner) *NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner {
	return &NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner{value: val, isSet: true}
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
