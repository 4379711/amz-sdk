package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the ImageDimensions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageDimensions{}

// ImageDimensions The dimensions that extend from the top left corner of the image (this applies to cropped and uncropped images). `ImageDimensions` units must be in pixels.
type ImageDimensions struct {
	Width  IntegerWithUnits `json:"width"`
	Height IntegerWithUnits `json:"height"`
}

type _ImageDimensions ImageDimensions

// NewImageDimensions instantiates a new ImageDimensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageDimensions(width IntegerWithUnits, height IntegerWithUnits) *ImageDimensions {
	this := ImageDimensions{}
	this.Width = width
	this.Height = height
	return &this
}

// NewImageDimensionsWithDefaults instantiates a new ImageDimensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageDimensionsWithDefaults() *ImageDimensions {
	this := ImageDimensions{}
	return &this
}

// GetWidth returns the Width field value
func (o *ImageDimensions) GetWidth() IntegerWithUnits {
	if o == nil {
		var ret IntegerWithUnits
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *ImageDimensions) GetWidthOk() (*IntegerWithUnits, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *ImageDimensions) SetWidth(v IntegerWithUnits) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *ImageDimensions) GetHeight() IntegerWithUnits {
	if o == nil {
		var ret IntegerWithUnits
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *ImageDimensions) GetHeightOk() (*IntegerWithUnits, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *ImageDimensions) SetHeight(v IntegerWithUnits) {
	o.Height = v
}

func (o ImageDimensions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["width"] = o.Width
	toSerialize["height"] = o.Height
	return toSerialize, nil
}

type NullableImageDimensions struct {
	value *ImageDimensions
	isSet bool
}

func (v NullableImageDimensions) Get() *ImageDimensions {
	return v.value
}

func (v *NullableImageDimensions) Set(val *ImageDimensions) {
	v.value = val
	v.isSet = true
}

func (v NullableImageDimensions) IsSet() bool {
	return v.isSet
}

func (v *NullableImageDimensions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageDimensions(val *ImageDimensions) *NullableImageDimensions {
	return &NullableImageDimensions{value: val, isSet: true}
}

func (v NullableImageDimensions) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableImageDimensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
