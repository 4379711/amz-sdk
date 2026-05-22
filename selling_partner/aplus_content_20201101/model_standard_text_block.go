package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the StandardTextBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardTextBlock{}

// StandardTextBlock The A+ Content standard text box block, which contains a paragraph and a headline.
type StandardTextBlock struct {
	Headline *TextComponent      `json:"headline,omitempty"`
	Body     *ParagraphComponent `json:"body,omitempty"`
}

// NewStandardTextBlock instantiates a new StandardTextBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardTextBlock() *StandardTextBlock {
	this := StandardTextBlock{}
	return &this
}

// NewStandardTextBlockWithDefaults instantiates a new StandardTextBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardTextBlockWithDefaults() *StandardTextBlock {
	this := StandardTextBlock{}
	return &this
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *StandardTextBlock) GetHeadline() TextComponent {
	if o == nil || IsNil(o.Headline) {
		var ret TextComponent
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardTextBlock) GetHeadlineOk() (*TextComponent, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *StandardTextBlock) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given TextComponent and assigns it to the Headline field.
func (o *StandardTextBlock) SetHeadline(v TextComponent) {
	o.Headline = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *StandardTextBlock) GetBody() ParagraphComponent {
	if o == nil || IsNil(o.Body) {
		var ret ParagraphComponent
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardTextBlock) GetBodyOk() (*ParagraphComponent, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *StandardTextBlock) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given ParagraphComponent and assigns it to the Body field.
func (o *StandardTextBlock) SetBody(v ParagraphComponent) {
	o.Body = &v
}

func (o StandardTextBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	return toSerialize, nil
}

type NullableStandardTextBlock struct {
	value *StandardTextBlock
	isSet bool
}

func (v NullableStandardTextBlock) Get() *StandardTextBlock {
	return v.value
}

func (v *NullableStandardTextBlock) Set(val *StandardTextBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardTextBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardTextBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardTextBlock(val *StandardTextBlock) *NullableStandardTextBlock {
	return &NullableStandardTextBlock{value: val, isSet: true}
}

func (v NullableStandardTextBlock) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStandardTextBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
