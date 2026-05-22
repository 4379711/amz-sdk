package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the StandardImageCaptionBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardImageCaptionBlock{}

// StandardImageCaptionBlock The A+ Content standard image and caption block.
type StandardImageCaptionBlock struct {
	Image   *ImageComponent `json:"image,omitempty"`
	Caption *TextComponent  `json:"caption,omitempty"`
}

// NewStandardImageCaptionBlock instantiates a new StandardImageCaptionBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardImageCaptionBlock() *StandardImageCaptionBlock {
	this := StandardImageCaptionBlock{}
	return &this
}

// NewStandardImageCaptionBlockWithDefaults instantiates a new StandardImageCaptionBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardImageCaptionBlockWithDefaults() *StandardImageCaptionBlock {
	this := StandardImageCaptionBlock{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *StandardImageCaptionBlock) GetImage() ImageComponent {
	if o == nil || IsNil(o.Image) {
		var ret ImageComponent
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardImageCaptionBlock) GetImageOk() (*ImageComponent, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *StandardImageCaptionBlock) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given ImageComponent and assigns it to the Image field.
func (o *StandardImageCaptionBlock) SetImage(v ImageComponent) {
	o.Image = &v
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *StandardImageCaptionBlock) GetCaption() TextComponent {
	if o == nil || IsNil(o.Caption) {
		var ret TextComponent
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardImageCaptionBlock) GetCaptionOk() (*TextComponent, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *StandardImageCaptionBlock) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given TextComponent and assigns it to the Caption field.
func (o *StandardImageCaptionBlock) SetCaption(v TextComponent) {
	o.Caption = &v
}

func (o StandardImageCaptionBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Caption) {
		toSerialize["caption"] = o.Caption
	}
	return toSerialize, nil
}

type NullableStandardImageCaptionBlock struct {
	value *StandardImageCaptionBlock
	isSet bool
}

func (v NullableStandardImageCaptionBlock) Get() *StandardImageCaptionBlock {
	return v.value
}

func (v *NullableStandardImageCaptionBlock) Set(val *StandardImageCaptionBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardImageCaptionBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardImageCaptionBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardImageCaptionBlock(val *StandardImageCaptionBlock) *NullableStandardImageCaptionBlock {
	return &NullableStandardImageCaptionBlock{value: val, isSet: true}
}

func (v NullableStandardImageCaptionBlock) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStandardImageCaptionBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
