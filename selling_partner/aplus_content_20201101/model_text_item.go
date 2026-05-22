package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the TextItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextItem{}

// TextItem Rich positional text that is usually presented as a collection of bullet points.
type TextItem struct {
	// The rank or index of this text item within the collection. Different items cannot occupy the same position within a single collection.
	Position int32         `json:"position"`
	Text     TextComponent `json:"text"`
}

type _TextItem TextItem

// NewTextItem instantiates a new TextItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextItem(position int32, text TextComponent) *TextItem {
	this := TextItem{}
	this.Position = position
	this.Text = text
	return &this
}

// NewTextItemWithDefaults instantiates a new TextItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextItemWithDefaults() *TextItem {
	this := TextItem{}
	return &this
}

// GetPosition returns the Position field value
func (o *TextItem) GetPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *TextItem) GetPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *TextItem) SetPosition(v int32) {
	o.Position = v
}

// GetText returns the Text field value
func (o *TextItem) GetText() TextComponent {
	if o == nil {
		var ret TextComponent
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *TextItem) GetTextOk() (*TextComponent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *TextItem) SetText(v TextComponent) {
	o.Text = v
}

func (o TextItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["position"] = o.Position
	toSerialize["text"] = o.Text
	return toSerialize, nil
}

type NullableTextItem struct {
	value *TextItem
	isSet bool
}

func (v NullableTextItem) Get() *TextItem {
	return v.value
}

func (v *NullableTextItem) Set(val *TextItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTextItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTextItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextItem(val *TextItem) *NullableTextItem {
	return &NullableTextItem{value: val, isSet: true}
}

func (v NullableTextItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTextItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
