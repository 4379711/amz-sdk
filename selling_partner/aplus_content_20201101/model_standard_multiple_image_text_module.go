package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the StandardMultipleImageTextModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardMultipleImageTextModule{}

// StandardMultipleImageTextModule Standard images with text, presented one at a time. The user clicks on thumbnails to view each block.
type StandardMultipleImageTextModule struct {
	Blocks []StandardImageTextCaptionBlock `json:"blocks,omitempty"`
}

// NewStandardMultipleImageTextModule instantiates a new StandardMultipleImageTextModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardMultipleImageTextModule() *StandardMultipleImageTextModule {
	this := StandardMultipleImageTextModule{}
	return &this
}

// NewStandardMultipleImageTextModuleWithDefaults instantiates a new StandardMultipleImageTextModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardMultipleImageTextModuleWithDefaults() *StandardMultipleImageTextModule {
	this := StandardMultipleImageTextModule{}
	return &this
}

// GetBlocks returns the Blocks field value if set, zero value otherwise.
func (o *StandardMultipleImageTextModule) GetBlocks() []StandardImageTextCaptionBlock {
	if o == nil || IsNil(o.Blocks) {
		var ret []StandardImageTextCaptionBlock
		return ret
	}
	return o.Blocks
}

// GetBlocksOk returns a tuple with the Blocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardMultipleImageTextModule) GetBlocksOk() ([]StandardImageTextCaptionBlock, bool) {
	if o == nil || IsNil(o.Blocks) {
		return nil, false
	}
	return o.Blocks, true
}

// HasBlocks returns a boolean if a field has been set.
func (o *StandardMultipleImageTextModule) HasBlocks() bool {
	if o != nil && !IsNil(o.Blocks) {
		return true
	}

	return false
}

// SetBlocks gets a reference to the given []StandardImageTextCaptionBlock and assigns it to the Blocks field.
func (o *StandardMultipleImageTextModule) SetBlocks(v []StandardImageTextCaptionBlock) {
	o.Blocks = v
}

func (o StandardMultipleImageTextModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Blocks) {
		toSerialize["blocks"] = o.Blocks
	}
	return toSerialize, nil
}

type NullableStandardMultipleImageTextModule struct {
	value *StandardMultipleImageTextModule
	isSet bool
}

func (v NullableStandardMultipleImageTextModule) Get() *StandardMultipleImageTextModule {
	return v.value
}

func (v *NullableStandardMultipleImageTextModule) Set(val *StandardMultipleImageTextModule) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardMultipleImageTextModule) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardMultipleImageTextModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardMultipleImageTextModule(val *StandardMultipleImageTextModule) *NullableStandardMultipleImageTextModule {
	return &NullableStandardMultipleImageTextModule{value: val, isSet: true}
}

func (v NullableStandardMultipleImageTextModule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStandardMultipleImageTextModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
