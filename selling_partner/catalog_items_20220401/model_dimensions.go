package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the Dimensions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dimensions{}

// Dimensions Dimensions of an Amazon catalog item or item in its packaging.
type Dimensions struct {
	Height *Dimension `json:"height,omitempty"`
	Length *Dimension `json:"length,omitempty"`
	Weight *Dimension `json:"weight,omitempty"`
	Width  *Dimension `json:"width,omitempty"`
}

// NewDimensions instantiates a new Dimensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensions() *Dimensions {
	this := Dimensions{}
	return &this
}

// NewDimensionsWithDefaults instantiates a new Dimensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionsWithDefaults() *Dimensions {
	this := Dimensions{}
	return &this
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *Dimensions) GetHeight() Dimension {
	if o == nil || IsNil(o.Height) {
		var ret Dimension
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dimensions) GetHeightOk() (*Dimension, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *Dimensions) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given Dimension and assigns it to the Height field.
func (o *Dimensions) SetHeight(v Dimension) {
	o.Height = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *Dimensions) GetLength() Dimension {
	if o == nil || IsNil(o.Length) {
		var ret Dimension
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dimensions) GetLengthOk() (*Dimension, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *Dimensions) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given Dimension and assigns it to the Length field.
func (o *Dimensions) SetLength(v Dimension) {
	o.Length = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *Dimensions) GetWeight() Dimension {
	if o == nil || IsNil(o.Weight) {
		var ret Dimension
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dimensions) GetWeightOk() (*Dimension, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *Dimensions) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given Dimension and assigns it to the Weight field.
func (o *Dimensions) SetWeight(v Dimension) {
	o.Weight = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *Dimensions) GetWidth() Dimension {
	if o == nil || IsNil(o.Width) {
		var ret Dimension
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dimensions) GetWidthOk() (*Dimension, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *Dimensions) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given Dimension and assigns it to the Width field.
func (o *Dimensions) SetWidth(v Dimension) {
	o.Width = &v
}

func (o Dimensions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
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
