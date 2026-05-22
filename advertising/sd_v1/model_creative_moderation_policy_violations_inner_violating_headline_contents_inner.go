package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner{}

// CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner struct for CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner
type CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner struct {
	// The specific text reviewed during moderation.
	ReviewedText *string                                                                                  `json:"reviewedText,omitempty"`
	TextEvidence []CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner `json:"textEvidence,omitempty"`
}

// NewCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner instantiates a new CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner() *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner {
	this := CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner{}
	return &this
}

// NewCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerWithDefaults instantiates a new CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerWithDefaults() *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner {
	this := CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner{}
	return &this
}

// GetReviewedText returns the ReviewedText field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner) GetReviewedText() string {
	if o == nil || IsNil(o.ReviewedText) {
		var ret string
		return ret
	}
	return *o.ReviewedText
}

// GetReviewedTextOk returns a tuple with the ReviewedText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner) GetReviewedTextOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewedText) {
		return nil, false
	}
	return o.ReviewedText, true
}

// HasReviewedText returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner) HasReviewedText() bool {
	if o != nil && !IsNil(o.ReviewedText) {
		return true
	}

	return false
}

// SetReviewedText gets a reference to the given string and assigns it to the ReviewedText field.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner) SetReviewedText(v string) {
	o.ReviewedText = &v
}

// GetTextEvidence returns the TextEvidence field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner) GetTextEvidence() []CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner {
	if o == nil || IsNil(o.TextEvidence) {
		var ret []CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner
		return ret
	}
	return o.TextEvidence
}

// GetTextEvidenceOk returns a tuple with the TextEvidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner) GetTextEvidenceOk() ([]CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner, bool) {
	if o == nil || IsNil(o.TextEvidence) {
		return nil, false
	}
	return o.TextEvidence, true
}

// HasTextEvidence returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner) HasTextEvidence() bool {
	if o != nil && !IsNil(o.TextEvidence) {
		return true
	}

	return false
}

// SetTextEvidence gets a reference to the given []CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner and assigns it to the TextEvidence field.
func (o *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner) SetTextEvidence(v []CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInnerTextEvidenceInner) {
	o.TextEvidence = v
}

func (o CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReviewedText) {
		toSerialize["reviewedText"] = o.ReviewedText
	}
	if !IsNil(o.TextEvidence) {
		toSerialize["textEvidence"] = o.TextEvidence
	}
	return toSerialize, nil
}

type NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner struct {
	value *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner
	isSet bool
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner) Get() *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner {
	return v.value
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner) Set(val *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner(val *CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner) *NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner {
	return &NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner{value: val, isSet: true}
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
