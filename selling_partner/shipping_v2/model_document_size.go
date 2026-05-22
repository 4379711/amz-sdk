package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the DocumentSize type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentSize{}

// DocumentSize The size dimensions of the label.
type DocumentSize struct {
	// The width of the document measured in the units specified.
	Width float32 `json:"width"`
	// The length of the document measured in the units specified.
	Length float32 `json:"length"`
	// The unit of measurement.
	Unit string `json:"unit"`
}

type _DocumentSize DocumentSize

// NewDocumentSize instantiates a new DocumentSize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentSize(width float32, length float32, unit string) *DocumentSize {
	this := DocumentSize{}
	this.Width = width
	this.Length = length
	this.Unit = unit
	return &this
}

// NewDocumentSizeWithDefaults instantiates a new DocumentSize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentSizeWithDefaults() *DocumentSize {
	this := DocumentSize{}
	return &this
}

// GetWidth returns the Width field value
func (o *DocumentSize) GetWidth() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *DocumentSize) GetWidthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *DocumentSize) SetWidth(v float32) {
	o.Width = v
}

// GetLength returns the Length field value
func (o *DocumentSize) GetLength() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *DocumentSize) GetLengthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *DocumentSize) SetLength(v float32) {
	o.Length = v
}

// GetUnit returns the Unit field value
func (o *DocumentSize) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *DocumentSize) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *DocumentSize) SetUnit(v string) {
	o.Unit = v
}

func (o DocumentSize) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["width"] = o.Width
	toSerialize["length"] = o.Length
	toSerialize["unit"] = o.Unit
	return toSerialize, nil
}

type NullableDocumentSize struct {
	value *DocumentSize
	isSet bool
}

func (v NullableDocumentSize) Get() *DocumentSize {
	return v.value
}

func (v *NullableDocumentSize) Set(val *DocumentSize) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentSize) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentSize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentSize(val *DocumentSize) *NullableDocumentSize {
	return &NullableDocumentSize{value: val, isSet: true}
}

func (v NullableDocumentSize) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDocumentSize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
