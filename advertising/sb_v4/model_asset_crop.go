package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the AssetCrop type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetCrop{}

// AssetCrop Asset cropping attributes
type AssetCrop struct {
	// The highest pixel from which to begin cropping
	Top *float32 `json:"top,omitempty"`
	// The leftmost pixel from which to begin cropping
	Left *float32 `json:"left,omitempty"`
	// The number of pixels to crop rightwards from the value specified as left
	Width *float32 `json:"width,omitempty"`
	// The number of pixels to crop down from the value specified as top
	Height *float32 `json:"height,omitempty"`
}

// NewAssetCrop instantiates a new AssetCrop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetCrop() *AssetCrop {
	this := AssetCrop{}
	return &this
}

// NewAssetCropWithDefaults instantiates a new AssetCrop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetCropWithDefaults() *AssetCrop {
	this := AssetCrop{}
	return &this
}

// GetTop returns the Top field value if set, zero value otherwise.
func (o *AssetCrop) GetTop() float32 {
	if o == nil || IsNil(o.Top) {
		var ret float32
		return ret
	}
	return *o.Top
}

// GetTopOk returns a tuple with the Top field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCrop) GetTopOk() (*float32, bool) {
	if o == nil || IsNil(o.Top) {
		return nil, false
	}
	return o.Top, true
}

// HasTop returns a boolean if a field has been set.
func (o *AssetCrop) HasTop() bool {
	if o != nil && !IsNil(o.Top) {
		return true
	}

	return false
}

// SetTop gets a reference to the given float32 and assigns it to the Top field.
func (o *AssetCrop) SetTop(v float32) {
	o.Top = &v
}

// GetLeft returns the Left field value if set, zero value otherwise.
func (o *AssetCrop) GetLeft() float32 {
	if o == nil || IsNil(o.Left) {
		var ret float32
		return ret
	}
	return *o.Left
}

// GetLeftOk returns a tuple with the Left field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCrop) GetLeftOk() (*float32, bool) {
	if o == nil || IsNil(o.Left) {
		return nil, false
	}
	return o.Left, true
}

// HasLeft returns a boolean if a field has been set.
func (o *AssetCrop) HasLeft() bool {
	if o != nil && !IsNil(o.Left) {
		return true
	}

	return false
}

// SetLeft gets a reference to the given float32 and assigns it to the Left field.
func (o *AssetCrop) SetLeft(v float32) {
	o.Left = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *AssetCrop) GetWidth() float32 {
	if o == nil || IsNil(o.Width) {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCrop) GetWidthOk() (*float32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *AssetCrop) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *AssetCrop) SetWidth(v float32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *AssetCrop) GetHeight() float32 {
	if o == nil || IsNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCrop) GetHeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *AssetCrop) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *AssetCrop) SetHeight(v float32) {
	o.Height = &v
}

func (o AssetCrop) ToMap() (map[string]interface{}, error) {
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

type NullableAssetCrop struct {
	value *AssetCrop
	isSet bool
}

func (v NullableAssetCrop) Get() *AssetCrop {
	return v.value
}

func (v *NullableAssetCrop) Set(val *AssetCrop) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetCrop) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetCrop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetCrop(val *AssetCrop) *NullableAssetCrop {
	return &NullableAssetCrop{value: val, isSet: true}
}

func (v NullableAssetCrop) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAssetCrop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
