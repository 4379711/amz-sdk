package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ViolatingImageContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViolatingImageContent{}

// ViolatingImageContent struct for ViolatingImageContent
type ViolatingImageContent struct {
	ViolatingImageEvidences []ViolatingImageEvidence `json:"violatingImageEvidences,omitempty"`
	// Moderation component which marked the policy violation.
	ModeratedComponent *string `json:"moderatedComponent,omitempty"`
	// URL of the image which has the ad policy violation.
	ReviewedImageUrl *string `json:"reviewedImageUrl,omitempty"`
}

// NewViolatingImageContent instantiates a new ViolatingImageContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViolatingImageContent() *ViolatingImageContent {
	this := ViolatingImageContent{}
	return &this
}

// NewViolatingImageContentWithDefaults instantiates a new ViolatingImageContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViolatingImageContentWithDefaults() *ViolatingImageContent {
	this := ViolatingImageContent{}
	return &this
}

// GetViolatingImageEvidences returns the ViolatingImageEvidences field value if set, zero value otherwise.
func (o *ViolatingImageContent) GetViolatingImageEvidences() []ViolatingImageEvidence {
	if o == nil || IsNil(o.ViolatingImageEvidences) {
		var ret []ViolatingImageEvidence
		return ret
	}
	return o.ViolatingImageEvidences
}

// GetViolatingImageEvidencesOk returns a tuple with the ViolatingImageEvidences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolatingImageContent) GetViolatingImageEvidencesOk() ([]ViolatingImageEvidence, bool) {
	if o == nil || IsNil(o.ViolatingImageEvidences) {
		return nil, false
	}
	return o.ViolatingImageEvidences, true
}

// HasViolatingImageEvidences returns a boolean if a field has been set.
func (o *ViolatingImageContent) HasViolatingImageEvidences() bool {
	if o != nil && !IsNil(o.ViolatingImageEvidences) {
		return true
	}

	return false
}

// SetViolatingImageEvidences gets a reference to the given []ViolatingImageEvidence and assigns it to the ViolatingImageEvidences field.
func (o *ViolatingImageContent) SetViolatingImageEvidences(v []ViolatingImageEvidence) {
	o.ViolatingImageEvidences = v
}

// GetModeratedComponent returns the ModeratedComponent field value if set, zero value otherwise.
func (o *ViolatingImageContent) GetModeratedComponent() string {
	if o == nil || IsNil(o.ModeratedComponent) {
		var ret string
		return ret
	}
	return *o.ModeratedComponent
}

// GetModeratedComponentOk returns a tuple with the ModeratedComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolatingImageContent) GetModeratedComponentOk() (*string, bool) {
	if o == nil || IsNil(o.ModeratedComponent) {
		return nil, false
	}
	return o.ModeratedComponent, true
}

// HasModeratedComponent returns a boolean if a field has been set.
func (o *ViolatingImageContent) HasModeratedComponent() bool {
	if o != nil && !IsNil(o.ModeratedComponent) {
		return true
	}

	return false
}

// SetModeratedComponent gets a reference to the given string and assigns it to the ModeratedComponent field.
func (o *ViolatingImageContent) SetModeratedComponent(v string) {
	o.ModeratedComponent = &v
}

// GetReviewedImageUrl returns the ReviewedImageUrl field value if set, zero value otherwise.
func (o *ViolatingImageContent) GetReviewedImageUrl() string {
	if o == nil || IsNil(o.ReviewedImageUrl) {
		var ret string
		return ret
	}
	return *o.ReviewedImageUrl
}

// GetReviewedImageUrlOk returns a tuple with the ReviewedImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolatingImageContent) GetReviewedImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewedImageUrl) {
		return nil, false
	}
	return o.ReviewedImageUrl, true
}

// HasReviewedImageUrl returns a boolean if a field has been set.
func (o *ViolatingImageContent) HasReviewedImageUrl() bool {
	if o != nil && !IsNil(o.ReviewedImageUrl) {
		return true
	}

	return false
}

// SetReviewedImageUrl gets a reference to the given string and assigns it to the ReviewedImageUrl field.
func (o *ViolatingImageContent) SetReviewedImageUrl(v string) {
	o.ReviewedImageUrl = &v
}

func (o ViolatingImageContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ViolatingImageEvidences) {
		toSerialize["violatingImageEvidences"] = o.ViolatingImageEvidences
	}
	if !IsNil(o.ModeratedComponent) {
		toSerialize["moderatedComponent"] = o.ModeratedComponent
	}
	if !IsNil(o.ReviewedImageUrl) {
		toSerialize["reviewedImageUrl"] = o.ReviewedImageUrl
	}
	return toSerialize, nil
}

type NullableViolatingImageContent struct {
	value *ViolatingImageContent
	isSet bool
}

func (v NullableViolatingImageContent) Get() *ViolatingImageContent {
	return v.value
}

func (v *NullableViolatingImageContent) Set(val *ViolatingImageContent) {
	v.value = val
	v.isSet = true
}

func (v NullableViolatingImageContent) IsSet() bool {
	return v.isSet
}

func (v *NullableViolatingImageContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViolatingImageContent(val *ViolatingImageContent) *NullableViolatingImageContent {
	return &NullableViolatingImageContent{value: val, isSet: true}
}

func (v NullableViolatingImageContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableViolatingImageContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
