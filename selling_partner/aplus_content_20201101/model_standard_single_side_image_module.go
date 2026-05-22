package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the StandardSingleSideImageModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardSingleSideImageModule{}

// StandardSingleSideImageModule A standard headline and body text with an image on the side.
type StandardSingleSideImageModule struct {
	ImagePositionType PositionType            `json:"imagePositionType"`
	Block             *StandardImageTextBlock `json:"block,omitempty"`
}

type _StandardSingleSideImageModule StandardSingleSideImageModule

// NewStandardSingleSideImageModule instantiates a new StandardSingleSideImageModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardSingleSideImageModule(imagePositionType PositionType) *StandardSingleSideImageModule {
	this := StandardSingleSideImageModule{}
	this.ImagePositionType = imagePositionType
	return &this
}

// NewStandardSingleSideImageModuleWithDefaults instantiates a new StandardSingleSideImageModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardSingleSideImageModuleWithDefaults() *StandardSingleSideImageModule {
	this := StandardSingleSideImageModule{}
	return &this
}

// GetImagePositionType returns the ImagePositionType field value
func (o *StandardSingleSideImageModule) GetImagePositionType() PositionType {
	if o == nil {
		var ret PositionType
		return ret
	}

	return o.ImagePositionType
}

// GetImagePositionTypeOk returns a tuple with the ImagePositionType field value
// and a boolean to check if the value has been set.
func (o *StandardSingleSideImageModule) GetImagePositionTypeOk() (*PositionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImagePositionType, true
}

// SetImagePositionType sets field value
func (o *StandardSingleSideImageModule) SetImagePositionType(v PositionType) {
	o.ImagePositionType = v
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *StandardSingleSideImageModule) GetBlock() StandardImageTextBlock {
	if o == nil || IsNil(o.Block) {
		var ret StandardImageTextBlock
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardSingleSideImageModule) GetBlockOk() (*StandardImageTextBlock, bool) {
	if o == nil || IsNil(o.Block) {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *StandardSingleSideImageModule) HasBlock() bool {
	if o != nil && !IsNil(o.Block) {
		return true
	}

	return false
}

// SetBlock gets a reference to the given StandardImageTextBlock and assigns it to the Block field.
func (o *StandardSingleSideImageModule) SetBlock(v StandardImageTextBlock) {
	o.Block = &v
}

func (o StandardSingleSideImageModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["imagePositionType"] = o.ImagePositionType
	if !IsNil(o.Block) {
		toSerialize["block"] = o.Block
	}
	return toSerialize, nil
}

type NullableStandardSingleSideImageModule struct {
	value *StandardSingleSideImageModule
	isSet bool
}

func (v NullableStandardSingleSideImageModule) Get() *StandardSingleSideImageModule {
	return v.value
}

func (v *NullableStandardSingleSideImageModule) Set(val *StandardSingleSideImageModule) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardSingleSideImageModule) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardSingleSideImageModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardSingleSideImageModule(val *StandardSingleSideImageModule) *NullableStandardSingleSideImageModule {
	return &NullableStandardSingleSideImageModule{value: val, isSet: true}
}

func (v NullableStandardSingleSideImageModule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStandardSingleSideImageModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
