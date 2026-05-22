package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the ImageComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageComponent{}

// ImageComponent A reference to an image, hosted in the A+ Content media library.
type ImageComponent struct {
	// This identifier is provided by the [Uploads API](https://developer-docs.amazon.com/sp-api/reference/welcome-to-api-references).
	UploadDestinationId    string                 `json:"uploadDestinationId"`
	ImageCropSpecification ImageCropSpecification `json:"imageCropSpecification"`
	// The alternative text for the image.
	AltText string `json:"altText"`
}

type _ImageComponent ImageComponent

// NewImageComponent instantiates a new ImageComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageComponent(uploadDestinationId string, imageCropSpecification ImageCropSpecification, altText string) *ImageComponent {
	this := ImageComponent{}
	this.UploadDestinationId = uploadDestinationId
	this.ImageCropSpecification = imageCropSpecification
	this.AltText = altText
	return &this
}

// NewImageComponentWithDefaults instantiates a new ImageComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageComponentWithDefaults() *ImageComponent {
	this := ImageComponent{}
	return &this
}

// GetUploadDestinationId returns the UploadDestinationId field value
func (o *ImageComponent) GetUploadDestinationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploadDestinationId
}

// GetUploadDestinationIdOk returns a tuple with the UploadDestinationId field value
// and a boolean to check if the value has been set.
func (o *ImageComponent) GetUploadDestinationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadDestinationId, true
}

// SetUploadDestinationId sets field value
func (o *ImageComponent) SetUploadDestinationId(v string) {
	o.UploadDestinationId = v
}

// GetImageCropSpecification returns the ImageCropSpecification field value
func (o *ImageComponent) GetImageCropSpecification() ImageCropSpecification {
	if o == nil {
		var ret ImageCropSpecification
		return ret
	}

	return o.ImageCropSpecification
}

// GetImageCropSpecificationOk returns a tuple with the ImageCropSpecification field value
// and a boolean to check if the value has been set.
func (o *ImageComponent) GetImageCropSpecificationOk() (*ImageCropSpecification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageCropSpecification, true
}

// SetImageCropSpecification sets field value
func (o *ImageComponent) SetImageCropSpecification(v ImageCropSpecification) {
	o.ImageCropSpecification = v
}

// GetAltText returns the AltText field value
func (o *ImageComponent) GetAltText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AltText
}

// GetAltTextOk returns a tuple with the AltText field value
// and a boolean to check if the value has been set.
func (o *ImageComponent) GetAltTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AltText, true
}

// SetAltText sets field value
func (o *ImageComponent) SetAltText(v string) {
	o.AltText = v
}

func (o ImageComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uploadDestinationId"] = o.UploadDestinationId
	toSerialize["imageCropSpecification"] = o.ImageCropSpecification
	toSerialize["altText"] = o.AltText
	return toSerialize, nil
}

type NullableImageComponent struct {
	value *ImageComponent
	isSet bool
}

func (v NullableImageComponent) Get() *ImageComponent {
	return v.value
}

func (v *NullableImageComponent) Set(val *ImageComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableImageComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableImageComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageComponent(val *ImageComponent) *NullableImageComponent {
	return &NullableImageComponent{value: val, isSet: true}
}

func (v NullableImageComponent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableImageComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
