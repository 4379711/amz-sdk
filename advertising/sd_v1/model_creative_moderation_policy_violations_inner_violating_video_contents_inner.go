package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner{}

// CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner struct for CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner
type CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner struct {
	// Address of the video reviewed during moderation.
	ReviewedVideoUrl *string                                                                                 `json:"reviewedVideoUrl,omitempty"`
	VideoEvidences   []CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner `json:"videoEvidences,omitempty"`
}

// NewCreativeModerationPolicyViolationsInnerViolatingVideoContentsInner instantiates a new CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeModerationPolicyViolationsInnerViolatingVideoContentsInner() *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner {
	this := CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner{}
	return &this
}

// NewCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerWithDefaults instantiates a new CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerWithDefaults() *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner {
	this := CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner{}
	return &this
}

// GetReviewedVideoUrl returns the ReviewedVideoUrl field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner) GetReviewedVideoUrl() string {
	if o == nil || IsNil(o.ReviewedVideoUrl) {
		var ret string
		return ret
	}
	return *o.ReviewedVideoUrl
}

// GetReviewedVideoUrlOk returns a tuple with the ReviewedVideoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner) GetReviewedVideoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewedVideoUrl) {
		return nil, false
	}
	return o.ReviewedVideoUrl, true
}

// HasReviewedVideoUrl returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner) HasReviewedVideoUrl() bool {
	if o != nil && !IsNil(o.ReviewedVideoUrl) {
		return true
	}

	return false
}

// SetReviewedVideoUrl gets a reference to the given string and assigns it to the ReviewedVideoUrl field.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner) SetReviewedVideoUrl(v string) {
	o.ReviewedVideoUrl = &v
}

// GetVideoEvidences returns the VideoEvidences field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner) GetVideoEvidences() []CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner {
	if o == nil || IsNil(o.VideoEvidences) {
		var ret []CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner
		return ret
	}
	return o.VideoEvidences
}

// GetVideoEvidencesOk returns a tuple with the VideoEvidences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner) GetVideoEvidencesOk() ([]CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner, bool) {
	if o == nil || IsNil(o.VideoEvidences) {
		return nil, false
	}
	return o.VideoEvidences, true
}

// HasVideoEvidences returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner) HasVideoEvidences() bool {
	if o != nil && !IsNil(o.VideoEvidences) {
		return true
	}

	return false
}

// SetVideoEvidences gets a reference to the given []CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner and assigns it to the VideoEvidences field.
func (o *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner) SetVideoEvidences(v []CreativeModerationPolicyViolationsInnerViolatingVideoContentsInnerVideoEvidencesInner) {
	o.VideoEvidences = v
}

func (o CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReviewedVideoUrl) {
		toSerialize["reviewedVideoUrl"] = o.ReviewedVideoUrl
	}
	if !IsNil(o.VideoEvidences) {
		toSerialize["videoEvidences"] = o.VideoEvidences
	}
	return toSerialize, nil
}

type NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInner struct {
	value *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner
	isSet bool
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInner) Get() *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner {
	return v.value
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInner) Set(val *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInner(val *CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner) *NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInner {
	return &NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInner{value: val, isSet: true}
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingVideoContentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
