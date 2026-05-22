package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the PlainTextItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlainTextItem{}

// PlainTextItem Plain positional text that is used in collections of brief labels and descriptors.
type PlainTextItem struct {
	// The rank or index of this text item within the collection. Different items cannot occupy the same position within a single collection.
	Position int32 `json:"position"`
	// The actual plain text.
	Value string `json:"value"`
}

type _PlainTextItem PlainTextItem

// NewPlainTextItem instantiates a new PlainTextItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlainTextItem(position int32, value string) *PlainTextItem {
	this := PlainTextItem{}
	this.Position = position
	this.Value = value
	return &this
}

// NewPlainTextItemWithDefaults instantiates a new PlainTextItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlainTextItemWithDefaults() *PlainTextItem {
	this := PlainTextItem{}
	return &this
}

// GetPosition returns the Position field value
func (o *PlainTextItem) GetPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *PlainTextItem) GetPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *PlainTextItem) SetPosition(v int32) {
	o.Position = v
}

// GetValue returns the Value field value
func (o *PlainTextItem) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PlainTextItem) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PlainTextItem) SetValue(v string) {
	o.Value = v
}

func (o PlainTextItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["position"] = o.Position
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullablePlainTextItem struct {
	value *PlainTextItem
	isSet bool
}

func (v NullablePlainTextItem) Get() *PlainTextItem {
	return v.value
}

func (v *NullablePlainTextItem) Set(val *PlainTextItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePlainTextItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePlainTextItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlainTextItem(val *PlainTextItem) *NullablePlainTextItem {
	return &NullablePlainTextItem{value: val, isSet: true}
}

func (v NullablePlainTextItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePlainTextItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
