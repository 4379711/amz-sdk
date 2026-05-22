package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the Dimensions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dimensions{}

// Dimensions Measurement of a package's dimensions.
type Dimensions struct {
	// The height of a package.
	Height float32 `json:"height"`
	// The length of a package.
	Length            float32           `json:"length"`
	UnitOfMeasurement UnitOfMeasurement `json:"unitOfMeasurement"`
	// The width of a package.
	Width float32 `json:"width"`
}

type _Dimensions Dimensions

// NewDimensions instantiates a new Dimensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensions(height float32, length float32, unitOfMeasurement UnitOfMeasurement, width float32) *Dimensions {
	this := Dimensions{}
	this.Height = height
	this.Length = length
	this.UnitOfMeasurement = unitOfMeasurement
	this.Width = width
	return &this
}

// NewDimensionsWithDefaults instantiates a new Dimensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionsWithDefaults() *Dimensions {
	this := Dimensions{}
	return &this
}

// GetHeight returns the Height field value
func (o *Dimensions) GetHeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *Dimensions) GetHeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *Dimensions) SetHeight(v float32) {
	o.Height = v
}

// GetLength returns the Length field value
func (o *Dimensions) GetLength() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *Dimensions) GetLengthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *Dimensions) SetLength(v float32) {
	o.Length = v
}

// GetUnitOfMeasurement returns the UnitOfMeasurement field value
func (o *Dimensions) GetUnitOfMeasurement() UnitOfMeasurement {
	if o == nil {
		var ret UnitOfMeasurement
		return ret
	}

	return o.UnitOfMeasurement
}

// GetUnitOfMeasurementOk returns a tuple with the UnitOfMeasurement field value
// and a boolean to check if the value has been set.
func (o *Dimensions) GetUnitOfMeasurementOk() (*UnitOfMeasurement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitOfMeasurement, true
}

// SetUnitOfMeasurement sets field value
func (o *Dimensions) SetUnitOfMeasurement(v UnitOfMeasurement) {
	o.UnitOfMeasurement = v
}

// GetWidth returns the Width field value
func (o *Dimensions) GetWidth() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *Dimensions) GetWidthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *Dimensions) SetWidth(v float32) {
	o.Width = v
}

func (o Dimensions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["height"] = o.Height
	toSerialize["length"] = o.Length
	toSerialize["unitOfMeasurement"] = o.UnitOfMeasurement
	toSerialize["width"] = o.Width
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
