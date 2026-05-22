package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner{}

// CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner struct for CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner
type CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner struct {
	// Address of the image reviewed during moderation.
	ReviewedImageUrl *string                                                                                     `json:"reviewedImageUrl,omitempty"`
	ImageEvidences   []CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner `json:"imageEvidences,omitempty"`
}

// NewCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner instantiates a new CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner() *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner {
	this := CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner{}
	return &this
}

// NewCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerWithDefaults instantiates a new CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerWithDefaults() *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner {
	this := CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner{}
	return &this
}

// GetReviewedImageUrl returns the ReviewedImageUrl field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) GetReviewedImageUrl() string {
	if o == nil || IsNil(o.ReviewedImageUrl) {
		var ret string
		return ret
	}
	return *o.ReviewedImageUrl
}

// GetReviewedImageUrlOk returns a tuple with the ReviewedImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) GetReviewedImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewedImageUrl) {
		return nil, false
	}
	return o.ReviewedImageUrl, true
}

// HasReviewedImageUrl returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) HasReviewedImageUrl() bool {
	if o != nil && !IsNil(o.ReviewedImageUrl) {
		return true
	}

	return false
}

// SetReviewedImageUrl gets a reference to the given string and assigns it to the ReviewedImageUrl field.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) SetReviewedImageUrl(v string) {
	o.ReviewedImageUrl = &v
}

// GetImageEvidences returns the ImageEvidences field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) GetImageEvidences() []CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner {
	if o == nil || IsNil(o.ImageEvidences) {
		var ret []CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner
		return ret
	}
	return o.ImageEvidences
}

// GetImageEvidencesOk returns a tuple with the ImageEvidences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) GetImageEvidencesOk() ([]CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner, bool) {
	if o == nil || IsNil(o.ImageEvidences) {
		return nil, false
	}
	return o.ImageEvidences, true
}

// HasImageEvidences returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) HasImageEvidences() bool {
	if o != nil && !IsNil(o.ImageEvidences) {
		return true
	}

	return false
}

// SetImageEvidences gets a reference to the given []CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner and assigns it to the ImageEvidences field.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) SetImageEvidences(v []CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInner) {
	o.ImageEvidences = v
}

func (o CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReviewedImageUrl) {
		toSerialize["reviewedImageUrl"] = o.ReviewedImageUrl
	}
	if !IsNil(o.ImageEvidences) {
		toSerialize["imageEvidences"] = o.ImageEvidences
	}
	return toSerialize, nil
}

type NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner struct {
	value *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner
	isSet bool
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) Get() *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner {
	return v.value
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) Set(val *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner(val *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) *NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner {
	return &NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner{value: val, isSet: true}
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
