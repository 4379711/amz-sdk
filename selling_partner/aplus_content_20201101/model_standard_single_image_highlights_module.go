package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the StandardSingleImageHighlightsModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardSingleImageHighlightsModule{}

// StandardSingleImageHighlightsModule A standard image with several paragraphs and a bulleted list.
type StandardSingleImageHighlightsModule struct {
	Image             *ImageComponent              `json:"image,omitempty"`
	Headline          *TextComponent               `json:"headline,omitempty"`
	TextBlock1        *StandardTextBlock           `json:"textBlock1,omitempty"`
	TextBlock2        *StandardTextBlock           `json:"textBlock2,omitempty"`
	TextBlock3        *StandardTextBlock           `json:"textBlock3,omitempty"`
	BulletedListBlock *StandardHeaderTextListBlock `json:"bulletedListBlock,omitempty"`
}

// NewStandardSingleImageHighlightsModule instantiates a new StandardSingleImageHighlightsModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardSingleImageHighlightsModule() *StandardSingleImageHighlightsModule {
	this := StandardSingleImageHighlightsModule{}
	return &this
}

// NewStandardSingleImageHighlightsModuleWithDefaults instantiates a new StandardSingleImageHighlightsModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardSingleImageHighlightsModuleWithDefaults() *StandardSingleImageHighlightsModule {
	this := StandardSingleImageHighlightsModule{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *StandardSingleImageHighlightsModule) GetImage() ImageComponent {
	if o == nil || IsNil(o.Image) {
		var ret ImageComponent
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardSingleImageHighlightsModule) GetImageOk() (*ImageComponent, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *StandardSingleImageHighlightsModule) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given ImageComponent and assigns it to the Image field.
func (o *StandardSingleImageHighlightsModule) SetImage(v ImageComponent) {
	o.Image = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *StandardSingleImageHighlightsModule) GetHeadline() TextComponent {
	if o == nil || IsNil(o.Headline) {
		var ret TextComponent
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardSingleImageHighlightsModule) GetHeadlineOk() (*TextComponent, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *StandardSingleImageHighlightsModule) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given TextComponent and assigns it to the Headline field.
func (o *StandardSingleImageHighlightsModule) SetHeadline(v TextComponent) {
	o.Headline = &v
}

// GetTextBlock1 returns the TextBlock1 field value if set, zero value otherwise.
func (o *StandardSingleImageHighlightsModule) GetTextBlock1() StandardTextBlock {
	if o == nil || IsNil(o.TextBlock1) {
		var ret StandardTextBlock
		return ret
	}
	return *o.TextBlock1
}

// GetTextBlock1Ok returns a tuple with the TextBlock1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardSingleImageHighlightsModule) GetTextBlock1Ok() (*StandardTextBlock, bool) {
	if o == nil || IsNil(o.TextBlock1) {
		return nil, false
	}
	return o.TextBlock1, true
}

// HasTextBlock1 returns a boolean if a field has been set.
func (o *StandardSingleImageHighlightsModule) HasTextBlock1() bool {
	if o != nil && !IsNil(o.TextBlock1) {
		return true
	}

	return false
}

// SetTextBlock1 gets a reference to the given StandardTextBlock and assigns it to the TextBlock1 field.
func (o *StandardSingleImageHighlightsModule) SetTextBlock1(v StandardTextBlock) {
	o.TextBlock1 = &v
}

// GetTextBlock2 returns the TextBlock2 field value if set, zero value otherwise.
func (o *StandardSingleImageHighlightsModule) GetTextBlock2() StandardTextBlock {
	if o == nil || IsNil(o.TextBlock2) {
		var ret StandardTextBlock
		return ret
	}
	return *o.TextBlock2
}

// GetTextBlock2Ok returns a tuple with the TextBlock2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardSingleImageHighlightsModule) GetTextBlock2Ok() (*StandardTextBlock, bool) {
	if o == nil || IsNil(o.TextBlock2) {
		return nil, false
	}
	return o.TextBlock2, true
}

// HasTextBlock2 returns a boolean if a field has been set.
func (o *StandardSingleImageHighlightsModule) HasTextBlock2() bool {
	if o != nil && !IsNil(o.TextBlock2) {
		return true
	}

	return false
}

// SetTextBlock2 gets a reference to the given StandardTextBlock and assigns it to the TextBlock2 field.
func (o *StandardSingleImageHighlightsModule) SetTextBlock2(v StandardTextBlock) {
	o.TextBlock2 = &v
}

// GetTextBlock3 returns the TextBlock3 field value if set, zero value otherwise.
func (o *StandardSingleImageHighlightsModule) GetTextBlock3() StandardTextBlock {
	if o == nil || IsNil(o.TextBlock3) {
		var ret StandardTextBlock
		return ret
	}
	return *o.TextBlock3
}

// GetTextBlock3Ok returns a tuple with the TextBlock3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardSingleImageHighlightsModule) GetTextBlock3Ok() (*StandardTextBlock, bool) {
	if o == nil || IsNil(o.TextBlock3) {
		return nil, false
	}
	return o.TextBlock3, true
}

// HasTextBlock3 returns a boolean if a field has been set.
func (o *StandardSingleImageHighlightsModule) HasTextBlock3() bool {
	if o != nil && !IsNil(o.TextBlock3) {
		return true
	}

	return false
}

// SetTextBlock3 gets a reference to the given StandardTextBlock and assigns it to the TextBlock3 field.
func (o *StandardSingleImageHighlightsModule) SetTextBlock3(v StandardTextBlock) {
	o.TextBlock3 = &v
}

// GetBulletedListBlock returns the BulletedListBlock field value if set, zero value otherwise.
func (o *StandardSingleImageHighlightsModule) GetBulletedListBlock() StandardHeaderTextListBlock {
	if o == nil || IsNil(o.BulletedListBlock) {
		var ret StandardHeaderTextListBlock
		return ret
	}
	return *o.BulletedListBlock
}

// GetBulletedListBlockOk returns a tuple with the BulletedListBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardSingleImageHighlightsModule) GetBulletedListBlockOk() (*StandardHeaderTextListBlock, bool) {
	if o == nil || IsNil(o.BulletedListBlock) {
		return nil, false
	}
	return o.BulletedListBlock, true
}

// HasBulletedListBlock returns a boolean if a field has been set.
func (o *StandardSingleImageHighlightsModule) HasBulletedListBlock() bool {
	if o != nil && !IsNil(o.BulletedListBlock) {
		return true
	}

	return false
}

// SetBulletedListBlock gets a reference to the given StandardHeaderTextListBlock and assigns it to the BulletedListBlock field.
func (o *StandardSingleImageHighlightsModule) SetBulletedListBlock(v StandardHeaderTextListBlock) {
	o.BulletedListBlock = &v
}

func (o StandardSingleImageHighlightsModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	if !IsNil(o.TextBlock1) {
		toSerialize["textBlock1"] = o.TextBlock1
	}
	if !IsNil(o.TextBlock2) {
		toSerialize["textBlock2"] = o.TextBlock2
	}
	if !IsNil(o.TextBlock3) {
		toSerialize["textBlock3"] = o.TextBlock3
	}
	if !IsNil(o.BulletedListBlock) {
		toSerialize["bulletedListBlock"] = o.BulletedListBlock
	}
	return toSerialize, nil
}

type NullableStandardSingleImageHighlightsModule struct {
	value *StandardSingleImageHighlightsModule
	isSet bool
}

func (v NullableStandardSingleImageHighlightsModule) Get() *StandardSingleImageHighlightsModule {
	return v.value
}

func (v *NullableStandardSingleImageHighlightsModule) Set(val *StandardSingleImageHighlightsModule) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardSingleImageHighlightsModule) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardSingleImageHighlightsModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardSingleImageHighlightsModule(val *StandardSingleImageHighlightsModule) *NullableStandardSingleImageHighlightsModule {
	return &NullableStandardSingleImageHighlightsModule{value: val, isSet: true}
}

func (v NullableStandardSingleImageHighlightsModule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStandardSingleImageHighlightsModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
