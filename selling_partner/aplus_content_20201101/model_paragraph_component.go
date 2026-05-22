package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the ParagraphComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParagraphComponent{}

// ParagraphComponent A list of rich text content that is typically presented in a text box.
type ParagraphComponent struct {
	TextList []TextComponent `json:"textList"`
}

type _ParagraphComponent ParagraphComponent

// NewParagraphComponent instantiates a new ParagraphComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParagraphComponent(textList []TextComponent) *ParagraphComponent {
	this := ParagraphComponent{}
	this.TextList = textList
	return &this
}

// NewParagraphComponentWithDefaults instantiates a new ParagraphComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParagraphComponentWithDefaults() *ParagraphComponent {
	this := ParagraphComponent{}
	return &this
}

// GetTextList returns the TextList field value
func (o *ParagraphComponent) GetTextList() []TextComponent {
	if o == nil {
		var ret []TextComponent
		return ret
	}

	return o.TextList
}

// GetTextListOk returns a tuple with the TextList field value
// and a boolean to check if the value has been set.
func (o *ParagraphComponent) GetTextListOk() ([]TextComponent, bool) {
	if o == nil {
		return nil, false
	}
	return o.TextList, true
}

// SetTextList sets field value
func (o *ParagraphComponent) SetTextList(v []TextComponent) {
	o.TextList = v
}

func (o ParagraphComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["textList"] = o.TextList
	return toSerialize, nil
}

type NullableParagraphComponent struct {
	value *ParagraphComponent
	isSet bool
}

func (v NullableParagraphComponent) Get() *ParagraphComponent {
	return v.value
}

func (v *NullableParagraphComponent) Set(val *ParagraphComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableParagraphComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableParagraphComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParagraphComponent(val *ParagraphComponent) *NullableParagraphComponent {
	return &NullableParagraphComponent{value: val, isSet: true}
}

func (v NullableParagraphComponent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableParagraphComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
