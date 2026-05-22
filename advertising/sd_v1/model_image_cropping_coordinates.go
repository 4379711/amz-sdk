package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ImageCroppingCoordinates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageCroppingCoordinates{}

// ImageCroppingCoordinates Optional cropping coordinates to apply to the image.
type ImageCroppingCoordinates struct {
	// Pixel distance from the top edge of the cropping zone to the top edge of the original image.
	Top int32 `json:"top"`
	// Pixel distance from the left edge of the cropping zone to the left edge of the original image.
	Left int32 `json:"left"`
	// Pixel width of the cropping zone.
	Width int32 `json:"width"`
	// Pixel height of the cropping zone.
	Height int32 `json:"height"`
}

type _ImageCroppingCoordinates ImageCroppingCoordinates

// NewImageCroppingCoordinates instantiates a new ImageCroppingCoordinates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageCroppingCoordinates(top int32, left int32, width int32, height int32) *ImageCroppingCoordinates {
	this := ImageCroppingCoordinates{}
	this.Top = top
	this.Left = left
	this.Width = width
	this.Height = height
	return &this
}

// NewImageCroppingCoordinatesWithDefaults instantiates a new ImageCroppingCoordinates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageCroppingCoordinatesWithDefaults() *ImageCroppingCoordinates {
	this := ImageCroppingCoordinates{}
	return &this
}

// GetTop returns the Top field value
func (o *ImageCroppingCoordinates) GetTop() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Top
}

// GetTopOk returns a tuple with the Top field value
// and a boolean to check if the value has been set.
func (o *ImageCroppingCoordinates) GetTopOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Top, true
}

// SetTop sets field value
func (o *ImageCroppingCoordinates) SetTop(v int32) {
	o.Top = v
}

// GetLeft returns the Left field value
func (o *ImageCroppingCoordinates) GetLeft() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Left
}

// GetLeftOk returns a tuple with the Left field value
// and a boolean to check if the value has been set.
func (o *ImageCroppingCoordinates) GetLeftOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Left, true
}

// SetLeft sets field value
func (o *ImageCroppingCoordinates) SetLeft(v int32) {
	o.Left = v
}

// GetWidth returns the Width field value
func (o *ImageCroppingCoordinates) GetWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *ImageCroppingCoordinates) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *ImageCroppingCoordinates) SetWidth(v int32) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *ImageCroppingCoordinates) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *ImageCroppingCoordinates) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *ImageCroppingCoordinates) SetHeight(v int32) {
	o.Height = v
}

func (o ImageCroppingCoordinates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["top"] = o.Top
	toSerialize["left"] = o.Left
	toSerialize["width"] = o.Width
	toSerialize["height"] = o.Height
	return toSerialize, nil
}

type NullableImageCroppingCoordinates struct {
	value *ImageCroppingCoordinates
	isSet bool
}

func (v NullableImageCroppingCoordinates) Get() *ImageCroppingCoordinates {
	return v.value
}

func (v *NullableImageCroppingCoordinates) Set(val *ImageCroppingCoordinates) {
	v.value = val
	v.isSet = true
}

func (v NullableImageCroppingCoordinates) IsSet() bool {
	return v.isSet
}

func (v *NullableImageCroppingCoordinates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageCroppingCoordinates(val *ImageCroppingCoordinates) *NullableImageCroppingCoordinates {
	return &NullableImageCroppingCoordinates{value: val, isSet: true}
}

func (v NullableImageCroppingCoordinates) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableImageCroppingCoordinates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
