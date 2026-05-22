package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the StandardImageTextOverlayModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardImageTextOverlayModule{}

// StandardImageTextOverlayModule A standard background image with a floating text box.
type StandardImageTextOverlayModule struct {
	OverlayColorType ColorType               `json:"overlayColorType"`
	Block            *StandardImageTextBlock `json:"block,omitempty"`
}

type _StandardImageTextOverlayModule StandardImageTextOverlayModule

// NewStandardImageTextOverlayModule instantiates a new StandardImageTextOverlayModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardImageTextOverlayModule(overlayColorType ColorType) *StandardImageTextOverlayModule {
	this := StandardImageTextOverlayModule{}
	this.OverlayColorType = overlayColorType
	return &this
}

// NewStandardImageTextOverlayModuleWithDefaults instantiates a new StandardImageTextOverlayModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardImageTextOverlayModuleWithDefaults() *StandardImageTextOverlayModule {
	this := StandardImageTextOverlayModule{}
	return &this
}

// GetOverlayColorType returns the OverlayColorType field value
func (o *StandardImageTextOverlayModule) GetOverlayColorType() ColorType {
	if o == nil {
		var ret ColorType
		return ret
	}

	return o.OverlayColorType
}

// GetOverlayColorTypeOk returns a tuple with the OverlayColorType field value
// and a boolean to check if the value has been set.
func (o *StandardImageTextOverlayModule) GetOverlayColorTypeOk() (*ColorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverlayColorType, true
}

// SetOverlayColorType sets field value
func (o *StandardImageTextOverlayModule) SetOverlayColorType(v ColorType) {
	o.OverlayColorType = v
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *StandardImageTextOverlayModule) GetBlock() StandardImageTextBlock {
	if o == nil || IsNil(o.Block) {
		var ret StandardImageTextBlock
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardImageTextOverlayModule) GetBlockOk() (*StandardImageTextBlock, bool) {
	if o == nil || IsNil(o.Block) {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *StandardImageTextOverlayModule) HasBlock() bool {
	if o != nil && !IsNil(o.Block) {
		return true
	}

	return false
}

// SetBlock gets a reference to the given StandardImageTextBlock and assigns it to the Block field.
func (o *StandardImageTextOverlayModule) SetBlock(v StandardImageTextBlock) {
	o.Block = &v
}

func (o StandardImageTextOverlayModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["overlayColorType"] = o.OverlayColorType
	if !IsNil(o.Block) {
		toSerialize["block"] = o.Block
	}
	return toSerialize, nil
}

type NullableStandardImageTextOverlayModule struct {
	value *StandardImageTextOverlayModule
	isSet bool
}

func (v NullableStandardImageTextOverlayModule) Get() *StandardImageTextOverlayModule {
	return v.value
}

func (v *NullableStandardImageTextOverlayModule) Set(val *StandardImageTextOverlayModule) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardImageTextOverlayModule) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardImageTextOverlayModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardImageTextOverlayModule(val *StandardImageTextOverlayModule) *NullableStandardImageTextOverlayModule {
	return &NullableStandardImageTextOverlayModule{value: val, isSet: true}
}

func (v NullableStandardImageTextOverlayModule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStandardImageTextOverlayModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
