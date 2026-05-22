package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ViolatingVideoContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViolatingVideoContent{}

// ViolatingVideoContent struct for ViolatingVideoContent
type ViolatingVideoContent struct {
	ViolatingVideoEvidences []ViolatingVideoEvidence `json:"violatingVideoEvidences,omitempty"`
	// Moderation component which marked the policy violation.
	ModeratedComponent *string `json:"moderatedComponent,omitempty"`
	// URL of the video which has the ad policy violation.
	ReviewedVideoUrl *string `json:"reviewedVideoUrl,omitempty"`
}

// NewViolatingVideoContent instantiates a new ViolatingVideoContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViolatingVideoContent() *ViolatingVideoContent {
	this := ViolatingVideoContent{}
	return &this
}

// NewViolatingVideoContentWithDefaults instantiates a new ViolatingVideoContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViolatingVideoContentWithDefaults() *ViolatingVideoContent {
	this := ViolatingVideoContent{}
	return &this
}

// GetViolatingVideoEvidences returns the ViolatingVideoEvidences field value if set, zero value otherwise.
func (o *ViolatingVideoContent) GetViolatingVideoEvidences() []ViolatingVideoEvidence {
	if o == nil || IsNil(o.ViolatingVideoEvidences) {
		var ret []ViolatingVideoEvidence
		return ret
	}
	return o.ViolatingVideoEvidences
}

// GetViolatingVideoEvidencesOk returns a tuple with the ViolatingVideoEvidences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolatingVideoContent) GetViolatingVideoEvidencesOk() ([]ViolatingVideoEvidence, bool) {
	if o == nil || IsNil(o.ViolatingVideoEvidences) {
		return nil, false
	}
	return o.ViolatingVideoEvidences, true
}

// HasViolatingVideoEvidences returns a boolean if a field has been set.
func (o *ViolatingVideoContent) HasViolatingVideoEvidences() bool {
	if o != nil && !IsNil(o.ViolatingVideoEvidences) {
		return true
	}

	return false
}

// SetViolatingVideoEvidences gets a reference to the given []ViolatingVideoEvidence and assigns it to the ViolatingVideoEvidences field.
func (o *ViolatingVideoContent) SetViolatingVideoEvidences(v []ViolatingVideoEvidence) {
	o.ViolatingVideoEvidences = v
}

// GetModeratedComponent returns the ModeratedComponent field value if set, zero value otherwise.
func (o *ViolatingVideoContent) GetModeratedComponent() string {
	if o == nil || IsNil(o.ModeratedComponent) {
		var ret string
		return ret
	}
	return *o.ModeratedComponent
}

// GetModeratedComponentOk returns a tuple with the ModeratedComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolatingVideoContent) GetModeratedComponentOk() (*string, bool) {
	if o == nil || IsNil(o.ModeratedComponent) {
		return nil, false
	}
	return o.ModeratedComponent, true
}

// HasModeratedComponent returns a boolean if a field has been set.
func (o *ViolatingVideoContent) HasModeratedComponent() bool {
	if o != nil && !IsNil(o.ModeratedComponent) {
		return true
	}

	return false
}

// SetModeratedComponent gets a reference to the given string and assigns it to the ModeratedComponent field.
func (o *ViolatingVideoContent) SetModeratedComponent(v string) {
	o.ModeratedComponent = &v
}

// GetReviewedVideoUrl returns the ReviewedVideoUrl field value if set, zero value otherwise.
func (o *ViolatingVideoContent) GetReviewedVideoUrl() string {
	if o == nil || IsNil(o.ReviewedVideoUrl) {
		var ret string
		return ret
	}
	return *o.ReviewedVideoUrl
}

// GetReviewedVideoUrlOk returns a tuple with the ReviewedVideoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolatingVideoContent) GetReviewedVideoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewedVideoUrl) {
		return nil, false
	}
	return o.ReviewedVideoUrl, true
}

// HasReviewedVideoUrl returns a boolean if a field has been set.
func (o *ViolatingVideoContent) HasReviewedVideoUrl() bool {
	if o != nil && !IsNil(o.ReviewedVideoUrl) {
		return true
	}

	return false
}

// SetReviewedVideoUrl gets a reference to the given string and assigns it to the ReviewedVideoUrl field.
func (o *ViolatingVideoContent) SetReviewedVideoUrl(v string) {
	o.ReviewedVideoUrl = &v
}

func (o ViolatingVideoContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ViolatingVideoEvidences) {
		toSerialize["violatingVideoEvidences"] = o.ViolatingVideoEvidences
	}
	if !IsNil(o.ModeratedComponent) {
		toSerialize["moderatedComponent"] = o.ModeratedComponent
	}
	if !IsNil(o.ReviewedVideoUrl) {
		toSerialize["reviewedVideoUrl"] = o.ReviewedVideoUrl
	}
	return toSerialize, nil
}

type NullableViolatingVideoContent struct {
	value *ViolatingVideoContent
	isSet bool
}

func (v NullableViolatingVideoContent) Get() *ViolatingVideoContent {
	return v.value
}

func (v *NullableViolatingVideoContent) Set(val *ViolatingVideoContent) {
	v.value = val
	v.isSet = true
}

func (v NullableViolatingVideoContent) IsSet() bool {
	return v.isSet
}

func (v *NullableViolatingVideoContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViolatingVideoContent(val *ViolatingVideoContent) *NullableViolatingVideoContent {
	return &NullableViolatingVideoContent{value: val, isSet: true}
}

func (v NullableViolatingVideoContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableViolatingVideoContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
