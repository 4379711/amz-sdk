package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the ImageCropSpecification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageCropSpecification{}

// ImageCropSpecification The instructions for optionally cropping an image. If you don't want to crop the image, set the dimensions to the original image size. If the image is cropped and you don't include offset values, the coordinates of the top left corner of the cropped image are set to (0,0) by default.
type ImageCropSpecification struct {
	Size   ImageDimensions `json:"size"`
	Offset *ImageOffsets   `json:"offset,omitempty"`
}

type _ImageCropSpecification ImageCropSpecification

// NewImageCropSpecification instantiates a new ImageCropSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageCropSpecification(size ImageDimensions) *ImageCropSpecification {
	this := ImageCropSpecification{}
	this.Size = size
	return &this
}

// NewImageCropSpecificationWithDefaults instantiates a new ImageCropSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageCropSpecificationWithDefaults() *ImageCropSpecification {
	this := ImageCropSpecification{}
	return &this
}

// GetSize returns the Size field value
func (o *ImageCropSpecification) GetSize() ImageDimensions {
	if o == nil {
		var ret ImageDimensions
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ImageCropSpecification) GetSizeOk() (*ImageDimensions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ImageCropSpecification) SetSize(v ImageDimensions) {
	o.Size = v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ImageCropSpecification) GetOffset() ImageOffsets {
	if o == nil || IsNil(o.Offset) {
		var ret ImageOffsets
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageCropSpecification) GetOffsetOk() (*ImageOffsets, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ImageCropSpecification) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given ImageOffsets and assigns it to the Offset field.
func (o *ImageCropSpecification) SetOffset(v ImageOffsets) {
	o.Offset = &v
}

func (o ImageCropSpecification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["size"] = o.Size
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	return toSerialize, nil
}

type NullableImageCropSpecification struct {
	value *ImageCropSpecification
	isSet bool
}

func (v NullableImageCropSpecification) Get() *ImageCropSpecification {
	return v.value
}

func (v *NullableImageCropSpecification) Set(val *ImageCropSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableImageCropSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableImageCropSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageCropSpecification(val *ImageCropSpecification) *NullableImageCropSpecification {
	return &NullableImageCropSpecification{value: val, isSet: true}
}

func (v NullableImageCropSpecification) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableImageCropSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
