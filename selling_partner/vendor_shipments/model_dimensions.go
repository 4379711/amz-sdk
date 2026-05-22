package vendor_shipments

import (
	"github.com/bytedance/sonic"
)

// checks if the Dimensions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dimensions{}

// Dimensions Physical dimensional measurements of a container.
type Dimensions struct {
	// A decimal number with no loss of precision. Useful when precision loss is unacceptable, as with currencies. Follows RFC7159 for number representation. <br>**Pattern** : `^-?(0|([1-9]\\d*))(\\.\\d+)?([eE][+-]?\\d+)?$`.
	Length string `json:"length"`
	// A decimal number with no loss of precision. Useful when precision loss is unacceptable, as with currencies. Follows RFC7159 for number representation. <br>**Pattern** : `^-?(0|([1-9]\\d*))(\\.\\d+)?([eE][+-]?\\d+)?$`.
	Width string `json:"width"`
	// A decimal number with no loss of precision. Useful when precision loss is unacceptable, as with currencies. Follows RFC7159 for number representation. <br>**Pattern** : `^-?(0|([1-9]\\d*))(\\.\\d+)?([eE][+-]?\\d+)?$`.
	Height string `json:"height"`
	// The unit of measure for dimensions.
	UnitOfMeasure string `json:"unitOfMeasure"`
}

type _Dimensions Dimensions

// NewDimensions instantiates a new Dimensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensions(length string, width string, height string, unitOfMeasure string) *Dimensions {
	this := Dimensions{}
	this.Length = length
	this.Width = width
	this.Height = height
	this.UnitOfMeasure = unitOfMeasure
	return &this
}

// NewDimensionsWithDefaults instantiates a new Dimensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionsWithDefaults() *Dimensions {
	this := Dimensions{}
	return &this
}

// GetLength returns the Length field value
func (o *Dimensions) GetLength() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *Dimensions) GetLengthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *Dimensions) SetLength(v string) {
	o.Length = v
}

// GetWidth returns the Width field value
func (o *Dimensions) GetWidth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *Dimensions) GetWidthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *Dimensions) SetWidth(v string) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *Dimensions) GetHeight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *Dimensions) GetHeightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *Dimensions) SetHeight(v string) {
	o.Height = v
}

// GetUnitOfMeasure returns the UnitOfMeasure field value
func (o *Dimensions) GetUnitOfMeasure() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value
// and a boolean to check if the value has been set.
func (o *Dimensions) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitOfMeasure, true
}

// SetUnitOfMeasure sets field value
func (o *Dimensions) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = v
}

func (o Dimensions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["length"] = o.Length
	toSerialize["width"] = o.Width
	toSerialize["height"] = o.Height
	toSerialize["unitOfMeasure"] = o.UnitOfMeasure
	return toSerialize, nil
}

type NullableDimensions struct {
	value *Dimensions
	isSet bool
}

func (v NullableDimensions) Get() *Dimensions {
	return v.value
}

func (v *NullableDimensions) Set(val *Dimensions) {
	v.value = val
	v.isSet = true
}

func (v NullableDimensions) IsSet() bool {
	return v.isSet
}

func (v *NullableDimensions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDimensions(val *Dimensions) *NullableDimensions {
	return &NullableDimensions{value: val, isSet: true}
}

func (v NullableDimensions) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDimensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
