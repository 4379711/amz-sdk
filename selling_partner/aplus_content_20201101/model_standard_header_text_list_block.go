package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the StandardHeaderTextListBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardHeaderTextListBlock{}

// StandardHeaderTextListBlock The A+ standard fixed-length list of text and a related headline.
type StandardHeaderTextListBlock struct {
	Headline *TextComponent         `json:"headline,omitempty"`
	Block    *StandardTextListBlock `json:"block,omitempty"`
}

// NewStandardHeaderTextListBlock instantiates a new StandardHeaderTextListBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardHeaderTextListBlock() *StandardHeaderTextListBlock {
	this := StandardHeaderTextListBlock{}
	return &this
}

// NewStandardHeaderTextListBlockWithDefaults instantiates a new StandardHeaderTextListBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardHeaderTextListBlockWithDefaults() *StandardHeaderTextListBlock {
	this := StandardHeaderTextListBlock{}
	return &this
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *StandardHeaderTextListBlock) GetHeadline() TextComponent {
	if o == nil || IsNil(o.Headline) {
		var ret TextComponent
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardHeaderTextListBlock) GetHeadlineOk() (*TextComponent, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *StandardHeaderTextListBlock) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given TextComponent and assigns it to the Headline field.
func (o *StandardHeaderTextListBlock) SetHeadline(v TextComponent) {
	o.Headline = &v
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *StandardHeaderTextListBlock) GetBlock() StandardTextListBlock {
	if o == nil || IsNil(o.Block) {
		var ret StandardTextListBlock
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardHeaderTextListBlock) GetBlockOk() (*StandardTextListBlock, bool) {
	if o == nil || IsNil(o.Block) {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *StandardHeaderTextListBlock) HasBlock() bool {
	if o != nil && !IsNil(o.Block) {
		return true
	}

	return false
}

// SetBlock gets a reference to the given StandardTextListBlock and assigns it to the Block field.
func (o *StandardHeaderTextListBlock) SetBlock(v StandardTextListBlock) {
	o.Block = &v
}

func (o StandardHeaderTextListBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	if !IsNil(o.Block) {
		toSerialize["block"] = o.Block
	}
	return toSerialize, nil
}

type NullableStandardHeaderTextListBlock struct {
	value *StandardHeaderTextListBlock
	isSet bool
}

func (v NullableStandardHeaderTextListBlock) Get() *StandardHeaderTextListBlock {
	return v.value
}

func (v *NullableStandardHeaderTextListBlock) Set(val *StandardHeaderTextListBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardHeaderTextListBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardHeaderTextListBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardHeaderTextListBlock(val *StandardHeaderTextListBlock) *NullableStandardHeaderTextListBlock {
	return &NullableStandardHeaderTextListBlock{value: val, isSet: true}
}

func (v NullableStandardHeaderTextListBlock) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStandardHeaderTextListBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
