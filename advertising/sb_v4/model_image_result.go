package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ImageResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageResult{}

// ImageResult struct for ImageResult
type ImageResult struct {
	// Alt text for this image
	ImageAltText *string `json:"imageAltText,omitempty"`
	ImageUrl     *string `json:"imageUrl,omitempty"`
}

// NewImageResult instantiates a new ImageResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageResult() *ImageResult {
	this := ImageResult{}
	return &this
}

// NewImageResultWithDefaults instantiates a new ImageResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageResultWithDefaults() *ImageResult {
	this := ImageResult{}
	return &this
}

// GetImageAltText returns the ImageAltText field value if set, zero value otherwise.
func (o *ImageResult) GetImageAltText() string {
	if o == nil || IsNil(o.ImageAltText) {
		var ret string
		return ret
	}
	return *o.ImageAltText
}

// GetImageAltTextOk returns a tuple with the ImageAltText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageResult) GetImageAltTextOk() (*string, bool) {
	if o == nil || IsNil(o.ImageAltText) {
		return nil, false
	}
	return o.ImageAltText, true
}

// HasImageAltText returns a boolean if a field has been set.
func (o *ImageResult) HasImageAltText() bool {
	if o != nil && !IsNil(o.ImageAltText) {
		return true
	}

	return false
}

// SetImageAltText gets a reference to the given string and assigns it to the ImageAltText field.
func (o *ImageResult) SetImageAltText(v string) {
	o.ImageAltText = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *ImageResult) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageResult) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *ImageResult) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *ImageResult) SetImageUrl(v string) {
	o.ImageUrl = &v
}

func (o ImageResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageAltText) {
		toSerialize["imageAltText"] = o.ImageAltText
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	return toSerialize, nil
}

type NullableImageResult struct {
	value *ImageResult
	isSet bool
}

func (v NullableImageResult) Get() *ImageResult {
	return v.value
}

func (v *NullableImageResult) Set(val *ImageResult) {
	v.value = val
	v.isSet = true
}

func (v NullableImageResult) IsSet() bool {
	return v.isSet
}

func (v *NullableImageResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageResult(val *ImageResult) *NullableImageResult {
	return &NullableImageResult{value: val, isSet: true}
}

func (v NullableImageResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableImageResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
