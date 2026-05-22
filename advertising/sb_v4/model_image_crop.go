package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ImageCrop type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageCrop{}

// ImageCrop struct for ImageCrop
type ImageCrop struct {
	// Policy violated region's top left Y-axis pixel value.
	TopLeftY *int64 `json:"topLeftY,omitempty"`
	// Policy violated region's top left X-axis pixel value.
	TopLeftX *int64 `json:"topLeftX,omitempty"`
	// Policy violated region's width in pixel.
	Width *int64 `json:"width,omitempty"`
	// Policy violated region's height in pixel.
	Height *int64 `json:"height,omitempty"`
}

// NewImageCrop instantiates a new ImageCrop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageCrop() *ImageCrop {
	this := ImageCrop{}
	return &this
}

// NewImageCropWithDefaults instantiates a new ImageCrop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageCropWithDefaults() *ImageCrop {
	this := ImageCrop{}
	return &this
}

// GetTopLeftY returns the TopLeftY field value if set, zero value otherwise.
func (o *ImageCrop) GetTopLeftY() int64 {
	if o == nil || IsNil(o.TopLeftY) {
		var ret int64
		return ret
	}
	return *o.TopLeftY
}

// GetTopLeftYOk returns a tuple with the TopLeftY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageCrop) GetTopLeftYOk() (*int64, bool) {
	if o == nil || IsNil(o.TopLeftY) {
		return nil, false
	}
	return o.TopLeftY, true
}

// HasTopLeftY returns a boolean if a field has been set.
func (o *ImageCrop) HasTopLeftY() bool {
	if o != nil && !IsNil(o.TopLeftY) {
		return true
	}

	return false
}

// SetTopLeftY gets a reference to the given int64 and assigns it to the TopLeftY field.
func (o *ImageCrop) SetTopLeftY(v int64) {
	o.TopLeftY = &v
}

// GetTopLeftX returns the TopLeftX field value if set, zero value otherwise.
func (o *ImageCrop) GetTopLeftX() int64 {
	if o == nil || IsNil(o.TopLeftX) {
		var ret int64
		return ret
	}
	return *o.TopLeftX
}

// GetTopLeftXOk returns a tuple with the TopLeftX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageCrop) GetTopLeftXOk() (*int64, bool) {
	if o == nil || IsNil(o.TopLeftX) {
		return nil, false
	}
	return o.TopLeftX, true
}

// HasTopLeftX returns a boolean if a field has been set.
func (o *ImageCrop) HasTopLeftX() bool {
	if o != nil && !IsNil(o.TopLeftX) {
		return true
	}

	return false
}

// SetTopLeftX gets a reference to the given int64 and assigns it to the TopLeftX field.
func (o *ImageCrop) SetTopLeftX(v int64) {
	o.TopLeftX = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *ImageCrop) GetWidth() int64 {
	if o == nil || IsNil(o.Width) {
		var ret int64
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageCrop) GetWidthOk() (*int64, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *ImageCrop) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int64 and assigns it to the Width field.
func (o *ImageCrop) SetWidth(v int64) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ImageCrop) GetHeight() int64 {
	if o == nil || IsNil(o.Height) {
		var ret int64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageCrop) GetHeightOk() (*int64, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ImageCrop) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int64 and assigns it to the Height field.
func (o *ImageCrop) SetHeight(v int64) {
	o.Height = &v
}

func (o ImageCrop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TopLeftY) {
		toSerialize["topLeftY"] = o.TopLeftY
	}
	if !IsNil(o.TopLeftX) {
		toSerialize["topLeftX"] = o.TopLeftX
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	return toSerialize, nil
}

type NullableImageCrop struct {
	value *ImageCrop
	isSet bool
}

func (v NullableImageCrop) Get() *ImageCrop {
	return v.value
}

func (v *NullableImageCrop) Set(val *ImageCrop) {
	v.value = val
	v.isSet = true
}

func (v NullableImageCrop) IsSet() bool {
	return v.isSet
}

func (v *NullableImageCrop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageCrop(val *ImageCrop) *NullableImageCrop {
	return &NullableImageCrop{value: val, isSet: true}
}

func (v NullableImageCrop) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableImageCrop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
