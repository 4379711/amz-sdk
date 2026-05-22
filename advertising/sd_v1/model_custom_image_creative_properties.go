package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CustomImageCreativeProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomImageCreativeProperties{}

// CustomImageCreativeProperties User-customizable properties of a custom image creative.
type CustomImageCreativeProperties struct {
	RectCustomImage   *Image `json:"rectCustomImage,omitempty"`
	SquareCustomImage *Image `json:"squareCustomImage,omitempty"`
	// An optional collection of 1:1 square images which are displayed on the ad.
	SquareImages []Image `json:"squareImages,omitempty"`
	// An optional collection of 1.91:1 horizontal images which are displayed on the ad.
	HorizontalImages []Image `json:"horizontalImages,omitempty"`
	// An optional collection of 9:16 vertical images which are displayed on the ad.
	VerticalImages []Image `json:"verticalImages,omitempty"`
}

// NewCustomImageCreativeProperties instantiates a new CustomImageCreativeProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomImageCreativeProperties() *CustomImageCreativeProperties {
	this := CustomImageCreativeProperties{}
	return &this
}

// NewCustomImageCreativePropertiesWithDefaults instantiates a new CustomImageCreativeProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomImageCreativePropertiesWithDefaults() *CustomImageCreativeProperties {
	this := CustomImageCreativeProperties{}
	return &this
}

// GetRectCustomImage returns the RectCustomImage field value if set, zero value otherwise.
func (o *CustomImageCreativeProperties) GetRectCustomImage() Image {
	if o == nil || IsNil(o.RectCustomImage) {
		var ret Image
		return ret
	}
	return *o.RectCustomImage
}

// GetRectCustomImageOk returns a tuple with the RectCustomImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomImageCreativeProperties) GetRectCustomImageOk() (*Image, bool) {
	if o == nil || IsNil(o.RectCustomImage) {
		return nil, false
	}
	return o.RectCustomImage, true
}

// HasRectCustomImage returns a boolean if a field has been set.
func (o *CustomImageCreativeProperties) HasRectCustomImage() bool {
	if o != nil && !IsNil(o.RectCustomImage) {
		return true
	}

	return false
}

// SetRectCustomImage gets a reference to the given Image and assigns it to the RectCustomImage field.
func (o *CustomImageCreativeProperties) SetRectCustomImage(v Image) {
	o.RectCustomImage = &v
}

// GetSquareCustomImage returns the SquareCustomImage field value if set, zero value otherwise.
func (o *CustomImageCreativeProperties) GetSquareCustomImage() Image {
	if o == nil || IsNil(o.SquareCustomImage) {
		var ret Image
		return ret
	}
	return *o.SquareCustomImage
}

// GetSquareCustomImageOk returns a tuple with the SquareCustomImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomImageCreativeProperties) GetSquareCustomImageOk() (*Image, bool) {
	if o == nil || IsNil(o.SquareCustomImage) {
		return nil, false
	}
	return o.SquareCustomImage, true
}

// HasSquareCustomImage returns a boolean if a field has been set.
func (o *CustomImageCreativeProperties) HasSquareCustomImage() bool {
	if o != nil && !IsNil(o.SquareCustomImage) {
		return true
	}

	return false
}

// SetSquareCustomImage gets a reference to the given Image and assigns it to the SquareCustomImage field.
func (o *CustomImageCreativeProperties) SetSquareCustomImage(v Image) {
	o.SquareCustomImage = &v
}

// GetSquareImages returns the SquareImages field value if set, zero value otherwise.
func (o *CustomImageCreativeProperties) GetSquareImages() []Image {
	if o == nil || IsNil(o.SquareImages) {
		var ret []Image
		return ret
	}
	return o.SquareImages
}

// GetSquareImagesOk returns a tuple with the SquareImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomImageCreativeProperties) GetSquareImagesOk() ([]Image, bool) {
	if o == nil || IsNil(o.SquareImages) {
		return nil, false
	}
	return o.SquareImages, true
}

// HasSquareImages returns a boolean if a field has been set.
func (o *CustomImageCreativeProperties) HasSquareImages() bool {
	if o != nil && !IsNil(o.SquareImages) {
		return true
	}

	return false
}

// SetSquareImages gets a reference to the given []Image and assigns it to the SquareImages field.
func (o *CustomImageCreativeProperties) SetSquareImages(v []Image) {
	o.SquareImages = v
}

// GetHorizontalImages returns the HorizontalImages field value if set, zero value otherwise.
func (o *CustomImageCreativeProperties) GetHorizontalImages() []Image {
	if o == nil || IsNil(o.HorizontalImages) {
		var ret []Image
		return ret
	}
	return o.HorizontalImages
}

// GetHorizontalImagesOk returns a tuple with the HorizontalImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomImageCreativeProperties) GetHorizontalImagesOk() ([]Image, bool) {
	if o == nil || IsNil(o.HorizontalImages) {
		return nil, false
	}
	return o.HorizontalImages, true
}

// HasHorizontalImages returns a boolean if a field has been set.
func (o *CustomImageCreativeProperties) HasHorizontalImages() bool {
	if o != nil && !IsNil(o.HorizontalImages) {
		return true
	}

	return false
}

// SetHorizontalImages gets a reference to the given []Image and assigns it to the HorizontalImages field.
func (o *CustomImageCreativeProperties) SetHorizontalImages(v []Image) {
	o.HorizontalImages = v
}

// GetVerticalImages returns the VerticalImages field value if set, zero value otherwise.
func (o *CustomImageCreativeProperties) GetVerticalImages() []Image {
	if o == nil || IsNil(o.VerticalImages) {
		var ret []Image
		return ret
	}
	return o.VerticalImages
}

// GetVerticalImagesOk returns a tuple with the VerticalImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomImageCreativeProperties) GetVerticalImagesOk() ([]Image, bool) {
	if o == nil || IsNil(o.VerticalImages) {
		return nil, false
	}
	return o.VerticalImages, true
}

// HasVerticalImages returns a boolean if a field has been set.
func (o *CustomImageCreativeProperties) HasVerticalImages() bool {
	if o != nil && !IsNil(o.VerticalImages) {
		return true
	}

	return false
}

// SetVerticalImages gets a reference to the given []Image and assigns it to the VerticalImages field.
func (o *CustomImageCreativeProperties) SetVerticalImages(v []Image) {
	o.VerticalImages = v
}

func (o CustomImageCreativeProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RectCustomImage) {
		toSerialize["rectCustomImage"] = o.RectCustomImage
	}
	if !IsNil(o.SquareCustomImage) {
		toSerialize["squareCustomImage"] = o.SquareCustomImage
	}
	if !IsNil(o.SquareImages) {
		toSerialize["squareImages"] = o.SquareImages
	}
	if !IsNil(o.HorizontalImages) {
		toSerialize["horizontalImages"] = o.HorizontalImages
	}
	if !IsNil(o.VerticalImages) {
		toSerialize["verticalImages"] = o.VerticalImages
	}
	return toSerialize, nil
}

type NullableCustomImageCreativeProperties struct {
	value *CustomImageCreativeProperties
	isSet bool
}

func (v NullableCustomImageCreativeProperties) Get() *CustomImageCreativeProperties {
	return v.value
}

func (v *NullableCustomImageCreativeProperties) Set(val *CustomImageCreativeProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomImageCreativeProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomImageCreativeProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomImageCreativeProperties(val *CustomImageCreativeProperties) *NullableCustomImageCreativeProperties {
	return &NullableCustomImageCreativeProperties{value: val, isSet: true}
}

func (v NullableCustomImageCreativeProperties) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCustomImageCreativeProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
