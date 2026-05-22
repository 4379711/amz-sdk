package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ViolatingTextContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViolatingTextContent{}

// ViolatingTextContent Information about the specific text that violates the specified policy in the campaign.
type ViolatingTextContent struct {
	// The actual text on which the moderation was done.
	ReviewedText           *string                 `json:"reviewedText,omitempty"`
	ViolatingTextEvidences []ViolatingTextEvidence `json:"violatingTextEvidences,omitempty"`
	// Moderation component which marked the policy violation.
	ModeratedComponent *string `json:"moderatedComponent,omitempty"`
}

// NewViolatingTextContent instantiates a new ViolatingTextContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViolatingTextContent() *ViolatingTextContent {
	this := ViolatingTextContent{}
	return &this
}

// NewViolatingTextContentWithDefaults instantiates a new ViolatingTextContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViolatingTextContentWithDefaults() *ViolatingTextContent {
	this := ViolatingTextContent{}
	return &this
}

// GetReviewedText returns the ReviewedText field value if set, zero value otherwise.
func (o *ViolatingTextContent) GetReviewedText() string {
	if o == nil || IsNil(o.ReviewedText) {
		var ret string
		return ret
	}
	return *o.ReviewedText
}

// GetReviewedTextOk returns a tuple with the ReviewedText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolatingTextContent) GetReviewedTextOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewedText) {
		return nil, false
	}
	return o.ReviewedText, true
}

// HasReviewedText returns a boolean if a field has been set.
func (o *ViolatingTextContent) HasReviewedText() bool {
	if o != nil && !IsNil(o.ReviewedText) {
		return true
	}

	return false
}

// SetReviewedText gets a reference to the given string and assigns it to the ReviewedText field.
func (o *ViolatingTextContent) SetReviewedText(v string) {
	o.ReviewedText = &v
}

// GetViolatingTextEvidences returns the ViolatingTextEvidences field value if set, zero value otherwise.
func (o *ViolatingTextContent) GetViolatingTextEvidences() []ViolatingTextEvidence {
	if o == nil || IsNil(o.ViolatingTextEvidences) {
		var ret []ViolatingTextEvidence
		return ret
	}
	return o.ViolatingTextEvidences
}

// GetViolatingTextEvidencesOk returns a tuple with the ViolatingTextEvidences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolatingTextContent) GetViolatingTextEvidencesOk() ([]ViolatingTextEvidence, bool) {
	if o == nil || IsNil(o.ViolatingTextEvidences) {
		return nil, false
	}
	return o.ViolatingTextEvidences, true
}

// HasViolatingTextEvidences returns a boolean if a field has been set.
func (o *ViolatingTextContent) HasViolatingTextEvidences() bool {
	if o != nil && !IsNil(o.ViolatingTextEvidences) {
		return true
	}

	return false
}

// SetViolatingTextEvidences gets a reference to the given []ViolatingTextEvidence and assigns it to the ViolatingTextEvidences field.
func (o *ViolatingTextContent) SetViolatingTextEvidences(v []ViolatingTextEvidence) {
	o.ViolatingTextEvidences = v
}

// GetModeratedComponent returns the ModeratedComponent field value if set, zero value otherwise.
func (o *ViolatingTextContent) GetModeratedComponent() string {
	if o == nil || IsNil(o.ModeratedComponent) {
		var ret string
		return ret
	}
	return *o.ModeratedComponent
}

// GetModeratedComponentOk returns a tuple with the ModeratedComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolatingTextContent) GetModeratedComponentOk() (*string, bool) {
	if o == nil || IsNil(o.ModeratedComponent) {
		return nil, false
	}
	return o.ModeratedComponent, true
}

// HasModeratedComponent returns a boolean if a field has been set.
func (o *ViolatingTextContent) HasModeratedComponent() bool {
	if o != nil && !IsNil(o.ModeratedComponent) {
		return true
	}

	return false
}

// SetModeratedComponent gets a reference to the given string and assigns it to the ModeratedComponent field.
func (o *ViolatingTextContent) SetModeratedComponent(v string) {
	o.ModeratedComponent = &v
}

func (o ViolatingTextContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReviewedText) {
		toSerialize["reviewedText"] = o.ReviewedText
	}
	if !IsNil(o.ViolatingTextEvidences) {
		toSerialize["violatingTextEvidences"] = o.ViolatingTextEvidences
	}
	if !IsNil(o.ModeratedComponent) {
		toSerialize["moderatedComponent"] = o.ModeratedComponent
	}
	return toSerialize, nil
}

type NullableViolatingTextContent struct {
	value *ViolatingTextContent
	isSet bool
}

func (v NullableViolatingTextContent) Get() *ViolatingTextContent {
	return v.value
}

func (v *NullableViolatingTextContent) Set(val *ViolatingTextContent) {
	v.value = val
	v.isSet = true
}

func (v NullableViolatingTextContent) IsSet() bool {
	return v.isSet
}

func (v *NullableViolatingTextContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViolatingTextContent(val *ViolatingTextContent) *NullableViolatingTextContent {
	return &NullableViolatingTextContent{value: val, isSet: true}
}

func (v NullableViolatingTextContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableViolatingTextContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
