package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ViolatingAsinContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViolatingAsinContent{}

// ViolatingAsinContent struct for ViolatingAsinContent
type ViolatingAsinContent struct {
	ViolatingAsinEvidences []ViolatingAsinEvidence `json:"violatingAsinEvidences,omitempty"`
	// Moderation component which marked the policy violation.
	ModeratedComponent *string `json:"moderatedComponent,omitempty"`
}

// NewViolatingAsinContent instantiates a new ViolatingAsinContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViolatingAsinContent() *ViolatingAsinContent {
	this := ViolatingAsinContent{}
	return &this
}

// NewViolatingAsinContentWithDefaults instantiates a new ViolatingAsinContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViolatingAsinContentWithDefaults() *ViolatingAsinContent {
	this := ViolatingAsinContent{}
	return &this
}

// GetViolatingAsinEvidences returns the ViolatingAsinEvidences field value if set, zero value otherwise.
func (o *ViolatingAsinContent) GetViolatingAsinEvidences() []ViolatingAsinEvidence {
	if o == nil || IsNil(o.ViolatingAsinEvidences) {
		var ret []ViolatingAsinEvidence
		return ret
	}
	return o.ViolatingAsinEvidences
}

// GetViolatingAsinEvidencesOk returns a tuple with the ViolatingAsinEvidences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolatingAsinContent) GetViolatingAsinEvidencesOk() ([]ViolatingAsinEvidence, bool) {
	if o == nil || IsNil(o.ViolatingAsinEvidences) {
		return nil, false
	}
	return o.ViolatingAsinEvidences, true
}

// HasViolatingAsinEvidences returns a boolean if a field has been set.
func (o *ViolatingAsinContent) HasViolatingAsinEvidences() bool {
	if o != nil && !IsNil(o.ViolatingAsinEvidences) {
		return true
	}

	return false
}

// SetViolatingAsinEvidences gets a reference to the given []ViolatingAsinEvidence and assigns it to the ViolatingAsinEvidences field.
func (o *ViolatingAsinContent) SetViolatingAsinEvidences(v []ViolatingAsinEvidence) {
	o.ViolatingAsinEvidences = v
}

// GetModeratedComponent returns the ModeratedComponent field value if set, zero value otherwise.
func (o *ViolatingAsinContent) GetModeratedComponent() string {
	if o == nil || IsNil(o.ModeratedComponent) {
		var ret string
		return ret
	}
	return *o.ModeratedComponent
}

// GetModeratedComponentOk returns a tuple with the ModeratedComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolatingAsinContent) GetModeratedComponentOk() (*string, bool) {
	if o == nil || IsNil(o.ModeratedComponent) {
		return nil, false
	}
	return o.ModeratedComponent, true
}

// HasModeratedComponent returns a boolean if a field has been set.
func (o *ViolatingAsinContent) HasModeratedComponent() bool {
	if o != nil && !IsNil(o.ModeratedComponent) {
		return true
	}

	return false
}

// SetModeratedComponent gets a reference to the given string and assigns it to the ModeratedComponent field.
func (o *ViolatingAsinContent) SetModeratedComponent(v string) {
	o.ModeratedComponent = &v
}

func (o ViolatingAsinContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ViolatingAsinEvidences) {
		toSerialize["violatingAsinEvidences"] = o.ViolatingAsinEvidences
	}
	if !IsNil(o.ModeratedComponent) {
		toSerialize["moderatedComponent"] = o.ModeratedComponent
	}
	return toSerialize, nil
}

type NullableViolatingAsinContent struct {
	value *ViolatingAsinContent
	isSet bool
}

func (v NullableViolatingAsinContent) Get() *ViolatingAsinContent {
	return v.value
}

func (v *NullableViolatingAsinContent) Set(val *ViolatingAsinContent) {
	v.value = val
	v.isSet = true
}

func (v NullableViolatingAsinContent) IsSet() bool {
	return v.isSet
}

func (v *NullableViolatingAsinContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViolatingAsinContent(val *ViolatingAsinContent) *NullableViolatingAsinContent {
	return &NullableViolatingAsinContent{value: val, isSet: true}
}

func (v NullableViolatingAsinContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableViolatingAsinContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
