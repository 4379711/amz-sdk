package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the StandardFourImageTextModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardFourImageTextModule{}

// StandardFourImageTextModule Four standard images with text, presented across a single row.
type StandardFourImageTextModule struct {
	Headline *TextComponent          `json:"headline,omitempty"`
	Block1   *StandardImageTextBlock `json:"block1,omitempty"`
	Block2   *StandardImageTextBlock `json:"block2,omitempty"`
	Block3   *StandardImageTextBlock `json:"block3,omitempty"`
	Block4   *StandardImageTextBlock `json:"block4,omitempty"`
}

// NewStandardFourImageTextModule instantiates a new StandardFourImageTextModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardFourImageTextModule() *StandardFourImageTextModule {
	this := StandardFourImageTextModule{}
	return &this
}

// NewStandardFourImageTextModuleWithDefaults instantiates a new StandardFourImageTextModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardFourImageTextModuleWithDefaults() *StandardFourImageTextModule {
	this := StandardFourImageTextModule{}
	return &this
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *StandardFourImageTextModule) GetHeadline() TextComponent {
	if o == nil || IsNil(o.Headline) {
		var ret TextComponent
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardFourImageTextModule) GetHeadlineOk() (*TextComponent, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *StandardFourImageTextModule) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given TextComponent and assigns it to the Headline field.
func (o *StandardFourImageTextModule) SetHeadline(v TextComponent) {
	o.Headline = &v
}

// GetBlock1 returns the Block1 field value if set, zero value otherwise.
func (o *StandardFourImageTextModule) GetBlock1() StandardImageTextBlock {
	if o == nil || IsNil(o.Block1) {
		var ret StandardImageTextBlock
		return ret
	}
	return *o.Block1
}

// GetBlock1Ok returns a tuple with the Block1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardFourImageTextModule) GetBlock1Ok() (*StandardImageTextBlock, bool) {
	if o == nil || IsNil(o.Block1) {
		return nil, false
	}
	return o.Block1, true
}

// HasBlock1 returns a boolean if a field has been set.
func (o *StandardFourImageTextModule) HasBlock1() bool {
	if o != nil && !IsNil(o.Block1) {
		return true
	}

	return false
}

// SetBlock1 gets a reference to the given StandardImageTextBlock and assigns it to the Block1 field.
func (o *StandardFourImageTextModule) SetBlock1(v StandardImageTextBlock) {
	o.Block1 = &v
}

// GetBlock2 returns the Block2 field value if set, zero value otherwise.
func (o *StandardFourImageTextModule) GetBlock2() StandardImageTextBlock {
	if o == nil || IsNil(o.Block2) {
		var ret StandardImageTextBlock
		return ret
	}
	return *o.Block2
}

// GetBlock2Ok returns a tuple with the Block2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardFourImageTextModule) GetBlock2Ok() (*StandardImageTextBlock, bool) {
	if o == nil || IsNil(o.Block2) {
		return nil, false
	}
	return o.Block2, true
}

// HasBlock2 returns a boolean if a field has been set.
func (o *StandardFourImageTextModule) HasBlock2() bool {
	if o != nil && !IsNil(o.Block2) {
		return true
	}

	return false
}

// SetBlock2 gets a reference to the given StandardImageTextBlock and assigns it to the Block2 field.
func (o *StandardFourImageTextModule) SetBlock2(v StandardImageTextBlock) {
	o.Block2 = &v
}

// GetBlock3 returns the Block3 field value if set, zero value otherwise.
func (o *StandardFourImageTextModule) GetBlock3() StandardImageTextBlock {
	if o == nil || IsNil(o.Block3) {
		var ret StandardImageTextBlock
		return ret
	}
	return *o.Block3
}

// GetBlock3Ok returns a tuple with the Block3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardFourImageTextModule) GetBlock3Ok() (*StandardImageTextBlock, bool) {
	if o == nil || IsNil(o.Block3) {
		return nil, false
	}
	return o.Block3, true
}

// HasBlock3 returns a boolean if a field has been set.
func (o *StandardFourImageTextModule) HasBlock3() bool {
	if o != nil && !IsNil(o.Block3) {
		return true
	}

	return false
}

// SetBlock3 gets a reference to the given StandardImageTextBlock and assigns it to the Block3 field.
func (o *StandardFourImageTextModule) SetBlock3(v StandardImageTextBlock) {
	o.Block3 = &v
}

// GetBlock4 returns the Block4 field value if set, zero value otherwise.
func (o *StandardFourImageTextModule) GetBlock4() StandardImageTextBlock {
	if o == nil || IsNil(o.Block4) {
		var ret StandardImageTextBlock
		return ret
	}
	return *o.Block4
}

// GetBlock4Ok returns a tuple with the Block4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardFourImageTextModule) GetBlock4Ok() (*StandardImageTextBlock, bool) {
	if o == nil || IsNil(o.Block4) {
		return nil, false
	}
	return o.Block4, true
}

// HasBlock4 returns a boolean if a field has been set.
func (o *StandardFourImageTextModule) HasBlock4() bool {
	if o != nil && !IsNil(o.Block4) {
		return true
	}

	return false
}

// SetBlock4 gets a reference to the given StandardImageTextBlock and assigns it to the Block4 field.
func (o *StandardFourImageTextModule) SetBlock4(v StandardImageTextBlock) {
	o.Block4 = &v
}

func (o StandardFourImageTextModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	if !IsNil(o.Block1) {
		toSerialize["block1"] = o.Block1
	}
	if !IsNil(o.Block2) {
		toSerialize["block2"] = o.Block2
	}
	if !IsNil(o.Block3) {
		toSerialize["block3"] = o.Block3
	}
	if !IsNil(o.Block4) {
		toSerialize["block4"] = o.Block4
	}
	return toSerialize, nil
}

type NullableStandardFourImageTextModule struct {
	value *StandardFourImageTextModule
	isSet bool
}

func (v NullableStandardFourImageTextModule) Get() *StandardFourImageTextModule {
	return v.value
}

func (v *NullableStandardFourImageTextModule) Set(val *StandardFourImageTextModule) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardFourImageTextModule) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardFourImageTextModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardFourImageTextModule(val *StandardFourImageTextModule) *NullableStandardFourImageTextModule {
	return &NullableStandardFourImageTextModule{value: val, isSet: true}
}

func (v NullableStandardFourImageTextModule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStandardFourImageTextModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
