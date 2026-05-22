package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the StandardFourImageTextQuadrantModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardFourImageTextQuadrantModule{}

// StandardFourImageTextQuadrantModule Four standard images with text, presented on a grid of four quadrants.
type StandardFourImageTextQuadrantModule struct {
	Block1 StandardImageTextBlock `json:"block1"`
	Block2 StandardImageTextBlock `json:"block2"`
	Block3 StandardImageTextBlock `json:"block3"`
	Block4 StandardImageTextBlock `json:"block4"`
}

type _StandardFourImageTextQuadrantModule StandardFourImageTextQuadrantModule

// NewStandardFourImageTextQuadrantModule instantiates a new StandardFourImageTextQuadrantModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardFourImageTextQuadrantModule(block1 StandardImageTextBlock, block2 StandardImageTextBlock, block3 StandardImageTextBlock, block4 StandardImageTextBlock) *StandardFourImageTextQuadrantModule {
	this := StandardFourImageTextQuadrantModule{}
	this.Block1 = block1
	this.Block2 = block2
	this.Block3 = block3
	this.Block4 = block4
	return &this
}

// NewStandardFourImageTextQuadrantModuleWithDefaults instantiates a new StandardFourImageTextQuadrantModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardFourImageTextQuadrantModuleWithDefaults() *StandardFourImageTextQuadrantModule {
	this := StandardFourImageTextQuadrantModule{}
	return &this
}

// GetBlock1 returns the Block1 field value
func (o *StandardFourImageTextQuadrantModule) GetBlock1() StandardImageTextBlock {
	if o == nil {
		var ret StandardImageTextBlock
		return ret
	}

	return o.Block1
}

// GetBlock1Ok returns a tuple with the Block1 field value
// and a boolean to check if the value has been set.
func (o *StandardFourImageTextQuadrantModule) GetBlock1Ok() (*StandardImageTextBlock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Block1, true
}

// SetBlock1 sets field value
func (o *StandardFourImageTextQuadrantModule) SetBlock1(v StandardImageTextBlock) {
	o.Block1 = v
}

// GetBlock2 returns the Block2 field value
func (o *StandardFourImageTextQuadrantModule) GetBlock2() StandardImageTextBlock {
	if o == nil {
		var ret StandardImageTextBlock
		return ret
	}

	return o.Block2
}

// GetBlock2Ok returns a tuple with the Block2 field value
// and a boolean to check if the value has been set.
func (o *StandardFourImageTextQuadrantModule) GetBlock2Ok() (*StandardImageTextBlock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Block2, true
}

// SetBlock2 sets field value
func (o *StandardFourImageTextQuadrantModule) SetBlock2(v StandardImageTextBlock) {
	o.Block2 = v
}

// GetBlock3 returns the Block3 field value
func (o *StandardFourImageTextQuadrantModule) GetBlock3() StandardImageTextBlock {
	if o == nil {
		var ret StandardImageTextBlock
		return ret
	}

	return o.Block3
}

// GetBlock3Ok returns a tuple with the Block3 field value
// and a boolean to check if the value has been set.
func (o *StandardFourImageTextQuadrantModule) GetBlock3Ok() (*StandardImageTextBlock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Block3, true
}

// SetBlock3 sets field value
func (o *StandardFourImageTextQuadrantModule) SetBlock3(v StandardImageTextBlock) {
	o.Block3 = v
}

// GetBlock4 returns the Block4 field value
func (o *StandardFourImageTextQuadrantModule) GetBlock4() StandardImageTextBlock {
	if o == nil {
		var ret StandardImageTextBlock
		return ret
	}

	return o.Block4
}

// GetBlock4Ok returns a tuple with the Block4 field value
// and a boolean to check if the value has been set.
func (o *StandardFourImageTextQuadrantModule) GetBlock4Ok() (*StandardImageTextBlock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Block4, true
}

// SetBlock4 sets field value
func (o *StandardFourImageTextQuadrantModule) SetBlock4(v StandardImageTextBlock) {
	o.Block4 = v
}

func (o StandardFourImageTextQuadrantModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["block1"] = o.Block1
	toSerialize["block2"] = o.Block2
	toSerialize["block3"] = o.Block3
	toSerialize["block4"] = o.Block4
	return toSerialize, nil
}

type NullableStandardFourImageTextQuadrantModule struct {
	value *StandardFourImageTextQuadrantModule
	isSet bool
}

func (v NullableStandardFourImageTextQuadrantModule) Get() *StandardFourImageTextQuadrantModule {
	return v.value
}

func (v *NullableStandardFourImageTextQuadrantModule) Set(val *StandardFourImageTextQuadrantModule) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardFourImageTextQuadrantModule) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardFourImageTextQuadrantModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardFourImageTextQuadrantModule(val *StandardFourImageTextQuadrantModule) *NullableStandardFourImageTextQuadrantModule {
	return &NullableStandardFourImageTextQuadrantModule{value: val, isSet: true}
}

func (v NullableStandardFourImageTextQuadrantModule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStandardFourImageTextQuadrantModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
