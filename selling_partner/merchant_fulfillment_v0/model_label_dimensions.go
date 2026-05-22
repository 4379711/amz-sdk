package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the LabelDimensions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelDimensions{}

// LabelDimensions Dimensions for printing a shipping label.
type LabelDimensions struct {
	// A label dimension.
	Length float32 `json:"Length"`
	// A label dimension.
	Width float32      `json:"Width"`
	Unit  UnitOfLength `json:"Unit"`
}

type _LabelDimensions LabelDimensions

// NewLabelDimensions instantiates a new LabelDimensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelDimensions(length float32, width float32, unit UnitOfLength) *LabelDimensions {
	this := LabelDimensions{}
	this.Length = length
	this.Width = width
	this.Unit = unit
	return &this
}

// NewLabelDimensionsWithDefaults instantiates a new LabelDimensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelDimensionsWithDefaults() *LabelDimensions {
	this := LabelDimensions{}
	return &this
}

// GetLength returns the Length field value
func (o *LabelDimensions) GetLength() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *LabelDimensions) GetLengthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *LabelDimensions) SetLength(v float32) {
	o.Length = v
}

// GetWidth returns the Width field value
func (o *LabelDimensions) GetWidth() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *LabelDimensions) GetWidthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *LabelDimensions) SetWidth(v float32) {
	o.Width = v
}

// GetUnit returns the Unit field value
func (o *LabelDimensions) GetUnit() UnitOfLength {
	if o == nil {
		var ret UnitOfLength
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *LabelDimensions) GetUnitOk() (*UnitOfLength, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *LabelDimensions) SetUnit(v UnitOfLength) {
	o.Unit = v
}

func (o LabelDimensions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Length"] = o.Length
	toSerialize["Width"] = o.Width
	toSerialize["Unit"] = o.Unit
	return toSerialize, nil
}

type NullableLabelDimensions struct {
	value *LabelDimensions
	isSet bool
}

func (v NullableLabelDimensions) Get() *LabelDimensions {
	return v.value
}

func (v *NullableLabelDimensions) Set(val *LabelDimensions) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelDimensions) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelDimensions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelDimensions(val *LabelDimensions) *NullableLabelDimensions {
	return &NullableLabelDimensions{value: val, isSet: true}
}

func (v NullableLabelDimensions) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLabelDimensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
