package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the Image type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Image{}

// Image This field denotes image which is displayed on the ad. This can either be a brand logo or a custom image. This field is optional and mutable. For custom image, both rectCustomImage and squareCustomImage should use the same asset id and asset version. Specific restrictions based on the Image type are listed in the following table. |Image type|Maximum file size|Minimum width|Minimum height|Accepted file formats| |------|-----------|-----------|-----------|-----------| |Custom Image|5MB|1200|628|JPEG, JPG, PNG, GIF| |Brand Logo|1MB|600|100|JPEG, JPG, PNG| Note: For square custom images the cropped image should be 628x628 at minimum.
type Image struct {
	// The unique identifier of the image asset. This assetId comes from the Creative Asset Library.
	AssetId string `json:"assetId"`
	// The identifier of the particular image assetversion.
	AssetVersion        string                    `json:"assetVersion"`
	CroppingCoordinates *ImageCroppingCoordinates `json:"croppingCoordinates,omitempty"`
}

type _Image Image

// NewImage instantiates a new Image object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImage(assetId string, assetVersion string) *Image {
	this := Image{}
	this.AssetId = assetId
	this.AssetVersion = assetVersion
	return &this
}

// NewImageWithDefaults instantiates a new Image object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageWithDefaults() *Image {
	this := Image{}
	return &this
}

// GetAssetId returns the AssetId field value
func (o *Image) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *Image) GetAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *Image) SetAssetId(v string) {
	o.AssetId = v
}

// GetAssetVersion returns the AssetVersion field value
func (o *Image) GetAssetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetVersion
}

// GetAssetVersionOk returns a tuple with the AssetVersion field value
// and a boolean to check if the value has been set.
func (o *Image) GetAssetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetVersion, true
}

// SetAssetVersion sets field value
func (o *Image) SetAssetVersion(v string) {
	o.AssetVersion = v
}

// GetCroppingCoordinates returns the CroppingCoordinates field value if set, zero value otherwise.
func (o *Image) GetCroppingCoordinates() ImageCroppingCoordinates {
	if o == nil || IsNil(o.CroppingCoordinates) {
		var ret ImageCroppingCoordinates
		return ret
	}
	return *o.CroppingCoordinates
}

// GetCroppingCoordinatesOk returns a tuple with the CroppingCoordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetCroppingCoordinatesOk() (*ImageCroppingCoordinates, bool) {
	if o == nil || IsNil(o.CroppingCoordinates) {
		return nil, false
	}
	return o.CroppingCoordinates, true
}

// HasCroppingCoordinates returns a boolean if a field has been set.
func (o *Image) HasCroppingCoordinates() bool {
	if o != nil && !IsNil(o.CroppingCoordinates) {
		return true
	}

	return false
}

// SetCroppingCoordinates gets a reference to the given ImageCroppingCoordinates and assigns it to the CroppingCoordinates field.
func (o *Image) SetCroppingCoordinates(v ImageCroppingCoordinates) {
	o.CroppingCoordinates = &v
}

func (o Image) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assetId"] = o.AssetId
	toSerialize["assetVersion"] = o.AssetVersion
	if !IsNil(o.CroppingCoordinates) {
		toSerialize["croppingCoordinates"] = o.CroppingCoordinates
	}
	return toSerialize, nil
}

type NullableImage struct {
	value *Image
	isSet bool
}

func (v NullableImage) Get() *Image {
	return v.value
}

func (v *NullableImage) Set(val *Image) {
	v.value = val
	v.isSet = true
}

func (v NullableImage) IsSet() bool {
	return v.isSet
}

func (v *NullableImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImage(val *Image) *NullableImage {
	return &NullableImage{value: val, isSet: true}
}

func (v NullableImage) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
