package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ImagePolicyViolation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImagePolicyViolation{}

// ImagePolicyViolation Structure of policy violation for a image component
type ImagePolicyViolation struct {
	// A human-readable description of the policy.
	PolicyDescription *string `json:"policyDescription,omitempty"`
	// List of evidences for the policy violations detected on the image component.
	ImageEvidences []ImageEvidence `json:"imageEvidences,omitempty"`
	// A policy violation code.
	Name *string `json:"name,omitempty"`
	// Type of policy violation.
	Type *string `json:"type,omitempty"`
	// Address of the policy documentation. Follow the link to learn more about the specified policy.
	PolicyLinkUrl *string `json:"policyLinkUrl,omitempty"`
	// Policy violation on an image can be detected on the ocr detected text on the image as well. This list of text evidences will have the policy violations detected on the text on top of the image.
	TextEvidences []TextEvidence `json:"textEvidences,omitempty"`
}

// NewImagePolicyViolation instantiates a new ImagePolicyViolation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImagePolicyViolation() *ImagePolicyViolation {
	this := ImagePolicyViolation{}
	return &this
}

// NewImagePolicyViolationWithDefaults instantiates a new ImagePolicyViolation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImagePolicyViolationWithDefaults() *ImagePolicyViolation {
	this := ImagePolicyViolation{}
	return &this
}

// GetPolicyDescription returns the PolicyDescription field value if set, zero value otherwise.
func (o *ImagePolicyViolation) GetPolicyDescription() string {
	if o == nil || IsNil(o.PolicyDescription) {
		var ret string
		return ret
	}
	return *o.PolicyDescription
}

// GetPolicyDescriptionOk returns a tuple with the PolicyDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImagePolicyViolation) GetPolicyDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyDescription) {
		return nil, false
	}
	return o.PolicyDescription, true
}

// HasPolicyDescription returns a boolean if a field has been set.
func (o *ImagePolicyViolation) HasPolicyDescription() bool {
	if o != nil && !IsNil(o.PolicyDescription) {
		return true
	}

	return false
}

// SetPolicyDescription gets a reference to the given string and assigns it to the PolicyDescription field.
func (o *ImagePolicyViolation) SetPolicyDescription(v string) {
	o.PolicyDescription = &v
}

// GetImageEvidences returns the ImageEvidences field value if set, zero value otherwise.
func (o *ImagePolicyViolation) GetImageEvidences() []ImageEvidence {
	if o == nil || IsNil(o.ImageEvidences) {
		var ret []ImageEvidence
		return ret
	}
	return o.ImageEvidences
}

// GetImageEvidencesOk returns a tuple with the ImageEvidences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImagePolicyViolation) GetImageEvidencesOk() ([]ImageEvidence, bool) {
	if o == nil || IsNil(o.ImageEvidences) {
		return nil, false
	}
	return o.ImageEvidences, true
}

// HasImageEvidences returns a boolean if a field has been set.
func (o *ImagePolicyViolation) HasImageEvidences() bool {
	if o != nil && !IsNil(o.ImageEvidences) {
		return true
	}

	return false
}

// SetImageEvidences gets a reference to the given []ImageEvidence and assigns it to the ImageEvidences field.
func (o *ImagePolicyViolation) SetImageEvidences(v []ImageEvidence) {
	o.ImageEvidences = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ImagePolicyViolation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImagePolicyViolation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ImagePolicyViolation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ImagePolicyViolation) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ImagePolicyViolation) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImagePolicyViolation) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ImagePolicyViolation) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ImagePolicyViolation) SetType(v string) {
	o.Type = &v
}

// GetPolicyLinkUrl returns the PolicyLinkUrl field value if set, zero value otherwise.
func (o *ImagePolicyViolation) GetPolicyLinkUrl() string {
	if o == nil || IsNil(o.PolicyLinkUrl) {
		var ret string
		return ret
	}
	return *o.PolicyLinkUrl
}

// GetPolicyLinkUrlOk returns a tuple with the PolicyLinkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImagePolicyViolation) GetPolicyLinkUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyLinkUrl) {
		return nil, false
	}
	return o.PolicyLinkUrl, true
}

// HasPolicyLinkUrl returns a boolean if a field has been set.
func (o *ImagePolicyViolation) HasPolicyLinkUrl() bool {
	if o != nil && !IsNil(o.PolicyLinkUrl) {
		return true
	}

	return false
}

// SetPolicyLinkUrl gets a reference to the given string and assigns it to the PolicyLinkUrl field.
func (o *ImagePolicyViolation) SetPolicyLinkUrl(v string) {
	o.PolicyLinkUrl = &v
}

// GetTextEvidences returns the TextEvidences field value if set, zero value otherwise.
func (o *ImagePolicyViolation) GetTextEvidences() []TextEvidence {
	if o == nil || IsNil(o.TextEvidences) {
		var ret []TextEvidence
		return ret
	}
	return o.TextEvidences
}

// GetTextEvidencesOk returns a tuple with the TextEvidences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImagePolicyViolation) GetTextEvidencesOk() ([]TextEvidence, bool) {
	if o == nil || IsNil(o.TextEvidences) {
		return nil, false
	}
	return o.TextEvidences, true
}

// HasTextEvidences returns a boolean if a field has been set.
func (o *ImagePolicyViolation) HasTextEvidences() bool {
	if o != nil && !IsNil(o.TextEvidences) {
		return true
	}

	return false
}

// SetTextEvidences gets a reference to the given []TextEvidence and assigns it to the TextEvidences field.
func (o *ImagePolicyViolation) SetTextEvidences(v []TextEvidence) {
	o.TextEvidences = v
}

func (o ImagePolicyViolation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PolicyDescription) {
		toSerialize["policyDescription"] = o.PolicyDescription
	}
	if !IsNil(o.ImageEvidences) {
		toSerialize["imageEvidences"] = o.ImageEvidences
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.PolicyLinkUrl) {
		toSerialize["policyLinkUrl"] = o.PolicyLinkUrl
	}
	if !IsNil(o.TextEvidences) {
		toSerialize["textEvidences"] = o.TextEvidences
	}
	return toSerialize, nil
}

type NullableImagePolicyViolation struct {
	value *ImagePolicyViolation
	isSet bool
}

func (v NullableImagePolicyViolation) Get() *ImagePolicyViolation {
	return v.value
}

func (v *NullableImagePolicyViolation) Set(val *ImagePolicyViolation) {
	v.value = val
	v.isSet = true
}

func (v NullableImagePolicyViolation) IsSet() bool {
	return v.isSet
}

func (v *NullableImagePolicyViolation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImagePolicyViolation(val *ImagePolicyViolation) *NullableImagePolicyViolation {
	return &NullableImagePolicyViolation{value: val, isSet: true}
}

func (v NullableImagePolicyViolation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableImagePolicyViolation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
