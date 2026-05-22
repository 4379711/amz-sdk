package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CustomImageCrop type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomImageCrop{}

// CustomImageCrop The crop to apply to the selected Custom image. A Custom image must have a 1200x628 aspect ratio, with a .01 delta for floating point precision. If a customImageAssetId is supplied but a crop is not, the crop will be defaulted to the whole image.
type CustomImageCrop struct {
	Top    *float32 `json:"top,omitempty"`
	Left   *float32 `json:"left,omitempty"`
	Width  *float32 `json:"width,omitempty"`
	Height *float32 `json:"height,omitempty"`
}

// NewCustomImageCrop instantiates a new CustomImageCrop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomImageCrop() *CustomImageCrop {
	this := CustomImageCrop{}
	return &this
}

// NewCustomImageCropWithDefaults instantiates a new CustomImageCrop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomImageCropWithDefaults() *CustomImageCrop {
	this := CustomImageCrop{}
	return &this
}

// GetTop returns the Top field value if set, zero value otherwise.
func (o *CustomImageCrop) GetTop() float32 {
	if o == nil || IsNil(o.Top) {
		var ret float32
		return ret
	}
	return *o.Top
}

// GetTopOk returns a tuple with the Top field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomImageCrop) GetTopOk() (*float32, bool) {
	if o == nil || IsNil(o.Top) {
		return nil, false
	}
	return o.Top, true
}

// HasTop returns a boolean if a field has been set.
func (o *CustomImageCrop) HasTop() bool {
	if o != nil && !IsNil(o.Top) {
		return true
	}

	return false
}

// SetTop gets a reference to the given float32 and assigns it to the Top field.
func (o *CustomImageCrop) SetTop(v float32) {
	o.Top = &v
}

// GetLeft returns the Left field value if set, zero value otherwise.
func (o *CustomImageCrop) GetLeft() float32 {
	if o == nil || IsNil(o.Left) {
		var ret float32
		return ret
	}
	return *o.Left
}

// GetLeftOk returns a tuple with the Left field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomImageCrop) GetLeftOk() (*float32, bool) {
	if o == nil || IsNil(o.Left) {
		return nil, false
	}
	return o.Left, true
}

// HasLeft returns a boolean if a field has been set.
func (o *CustomImageCrop) HasLeft() bool {
	if o != nil && !IsNil(o.Left) {
		return true
	}

	return false
}

// SetLeft gets a reference to the given float32 and assigns it to the Left field.
func (o *CustomImageCrop) SetLeft(v float32) {
	o.Left = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *CustomImageCrop) GetWidth() float32 {
	if o == nil || IsNil(o.Width) {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomImageCrop) GetWidthOk() (*float32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *CustomImageCrop) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *CustomImageCrop) SetWidth(v float32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *CustomImageCrop) GetHeight() float32 {
	if o == nil || IsNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomImageCrop) GetHeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *CustomImageCrop) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *CustomImageCrop) SetHeight(v float32) {
	o.Height = &v
}

func (o CustomImageCrop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Top) {
		toSerialize["top"] = o.Top
	}
	if !IsNil(o.Left) {
		toSerialize["left"] = o.Left
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	return toSerialize, nil
}

type NullableCustomImageCrop struct {
	value *CustomImageCrop
	isSet bool
}

func (v NullableCustomImageCrop) Get() *CustomImageCrop {
	return v.value
}

func (v *NullableCustomImageCrop) Set(val *CustomImageCrop) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomImageCrop) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomImageCrop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomImageCrop(val *CustomImageCrop) *NullableCustomImageCrop {
	return &NullableCustomImageCrop{value: val, isSet: true}
}

func (v NullableCustomImageCrop) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCustomImageCrop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
