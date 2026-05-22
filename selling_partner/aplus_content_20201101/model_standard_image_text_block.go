package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the StandardImageTextBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardImageTextBlock{}

// StandardImageTextBlock The A+ Content standard image and text box block.
type StandardImageTextBlock struct {
	Image    *ImageComponent     `json:"image,omitempty"`
	Headline *TextComponent      `json:"headline,omitempty"`
	Body     *ParagraphComponent `json:"body,omitempty"`
}

// NewStandardImageTextBlock instantiates a new StandardImageTextBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardImageTextBlock() *StandardImageTextBlock {
	this := StandardImageTextBlock{}
	return &this
}

// NewStandardImageTextBlockWithDefaults instantiates a new StandardImageTextBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardImageTextBlockWithDefaults() *StandardImageTextBlock {
	this := StandardImageTextBlock{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *StandardImageTextBlock) GetImage() ImageComponent {
	if o == nil || IsNil(o.Image) {
		var ret ImageComponent
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardImageTextBlock) GetImageOk() (*ImageComponent, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *StandardImageTextBlock) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given ImageComponent and assigns it to the Image field.
func (o *StandardImageTextBlock) SetImage(v ImageComponent) {
	o.Image = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *StandardImageTextBlock) GetHeadline() TextComponent {
	if o == nil || IsNil(o.Headline) {
		var ret TextComponent
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardImageTextBlock) GetHeadlineOk() (*TextComponent, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *StandardImageTextBlock) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given TextComponent and assigns it to the Headline field.
func (o *StandardImageTextBlock) SetHeadline(v TextComponent) {
	o.Headline = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *StandardImageTextBlock) GetBody() ParagraphComponent {
	if o == nil || IsNil(o.Body) {
		var ret ParagraphComponent
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardImageTextBlock) GetBodyOk() (*ParagraphComponent, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *StandardImageTextBlock) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given ParagraphComponent and assigns it to the Body field.
func (o *StandardImageTextBlock) SetBody(v ParagraphComponent) {
	o.Body = &v
}

func (o StandardImageTextBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	return toSerialize, nil
}

type NullableStandardImageTextBlock struct {
	value *StandardImageTextBlock
	isSet bool
}

func (v NullableStandardImageTextBlock) Get() *StandardImageTextBlock {
	return v.value
}

func (v *NullableStandardImageTextBlock) Set(val *StandardImageTextBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardImageTextBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardImageTextBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardImageTextBlock(val *StandardImageTextBlock) *NullableStandardImageTextBlock {
	return &NullableStandardImageTextBlock{value: val, isSet: true}
}

func (v NullableStandardImageTextBlock) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStandardImageTextBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
