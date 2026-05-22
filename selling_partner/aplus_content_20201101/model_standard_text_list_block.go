package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the StandardTextListBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardTextListBlock{}

// StandardTextListBlock The A+ Content standard fixed-length list of text, usually presented as bullet points.
type StandardTextListBlock struct {
	TextList []TextItem `json:"textList"`
}

type _StandardTextListBlock StandardTextListBlock

// NewStandardTextListBlock instantiates a new StandardTextListBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardTextListBlock(textList []TextItem) *StandardTextListBlock {
	this := StandardTextListBlock{}
	this.TextList = textList
	return &this
}

// NewStandardTextListBlockWithDefaults instantiates a new StandardTextListBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardTextListBlockWithDefaults() *StandardTextListBlock {
	this := StandardTextListBlock{}
	return &this
}

// GetTextList returns the TextList field value
func (o *StandardTextListBlock) GetTextList() []TextItem {
	if o == nil {
		var ret []TextItem
		return ret
	}

	return o.TextList
}

// GetTextListOk returns a tuple with the TextList field value
// and a boolean to check if the value has been set.
func (o *StandardTextListBlock) GetTextListOk() ([]TextItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.TextList, true
}

// SetTextList sets field value
func (o *StandardTextListBlock) SetTextList(v []TextItem) {
	o.TextList = v
}

func (o StandardTextListBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["textList"] = o.TextList
	return toSerialize, nil
}

type NullableStandardTextListBlock struct {
	value *StandardTextListBlock
	isSet bool
}

func (v NullableStandardTextListBlock) Get() *StandardTextListBlock {
	return v.value
}

func (v *NullableStandardTextListBlock) Set(val *StandardTextListBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardTextListBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardTextListBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardTextListBlock(val *StandardTextListBlock) *NullableStandardTextListBlock {
	return &NullableStandardTextListBlock{value: val, isSet: true}
}

func (v NullableStandardTextListBlock) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStandardTextListBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
