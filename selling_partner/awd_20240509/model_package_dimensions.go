package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the PackageDimensions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageDimensions{}

// PackageDimensions Dimensions of the package.
type PackageDimensions struct {
	// Height of the package.
	Height float64 `json:"height"`
	// Length of the package.
	Length            float64                    `json:"length"`
	UnitOfMeasurement DimensionUnitOfMeasurement `json:"unitOfMeasurement"`
	// Width of the package.
	Width float64 `json:"width"`
}

type _PackageDimensions PackageDimensions

// NewPackageDimensions instantiates a new PackageDimensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageDimensions(height float64, length float64, unitOfMeasurement DimensionUnitOfMeasurement, width float64) *PackageDimensions {
	this := PackageDimensions{}
	this.Height = height
	this.Length = length
	this.UnitOfMeasurement = unitOfMeasurement
	this.Width = width
	return &this
}

// NewPackageDimensionsWithDefaults instantiates a new PackageDimensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageDimensionsWithDefaults() *PackageDimensions {
	this := PackageDimensions{}
	return &this
}

// GetHeight returns the Height field value
func (o *PackageDimensions) GetHeight() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *PackageDimensions) GetHeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *PackageDimensions) SetHeight(v float64) {
	o.Height = v
}

// GetLength returns the Length field value
func (o *PackageDimensions) GetLength() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *PackageDimensions) GetLengthOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *PackageDimensions) SetLength(v float64) {
	o.Length = v
}

// GetUnitOfMeasurement returns the UnitOfMeasurement field value
func (o *PackageDimensions) GetUnitOfMeasurement() DimensionUnitOfMeasurement {
	if o == nil {
		var ret DimensionUnitOfMeasurement
		return ret
	}

	return o.UnitOfMeasurement
}

// GetUnitOfMeasurementOk returns a tuple with the UnitOfMeasurement field value
// and a boolean to check if the value has been set.
func (o *PackageDimensions) GetUnitOfMeasurementOk() (*DimensionUnitOfMeasurement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitOfMeasurement, true
}

// SetUnitOfMeasurement sets field value
func (o *PackageDimensions) SetUnitOfMeasurement(v DimensionUnitOfMeasurement) {
	o.UnitOfMeasurement = v
}

// GetWidth returns the Width field value
func (o *PackageDimensions) GetWidth() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *PackageDimensions) GetWidthOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *PackageDimensions) SetWidth(v float64) {
	o.Width = v
}

func (o PackageDimensions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["height"] = o.Height
	toSerialize["length"] = o.Length
	toSerialize["unitOfMeasurement"] = o.UnitOfMeasurement
	toSerialize["width"] = o.Width
	return toSerialize, nil
}

type NullablePackageDimensions struct {
	value *PackageDimensions
	isSet bool
}

func (v NullablePackageDimensions) Get() *PackageDimensions {
	return v.value
}

func (v *NullablePackageDimensions) Set(val *PackageDimensions) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageDimensions) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageDimensions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageDimensions(val *PackageDimensions) *NullablePackageDimensions {
	return &NullablePackageDimensions{value: val, isSet: true}
}

func (v NullablePackageDimensions) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackageDimensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
