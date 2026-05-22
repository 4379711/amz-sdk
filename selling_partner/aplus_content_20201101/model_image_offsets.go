package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the ImageOffsets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageOffsets{}

// ImageOffsets The top left corner of the cropped image, specified in the original image's coordinate space.
type ImageOffsets struct {
	X IntegerWithUnits `json:"x"`
	Y IntegerWithUnits `json:"y"`
}

type _ImageOffsets ImageOffsets

// NewImageOffsets instantiates a new ImageOffsets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageOffsets(x IntegerWithUnits, y IntegerWithUnits) *ImageOffsets {
	this := ImageOffsets{}
	this.X = x
	this.Y = y
	return &this
}

// NewImageOffsetsWithDefaults instantiates a new ImageOffsets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageOffsetsWithDefaults() *ImageOffsets {
	this := ImageOffsets{}
	return &this
}

// GetX returns the X field value
func (o *ImageOffsets) GetX() IntegerWithUnits {
	if o == nil {
		var ret IntegerWithUnits
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *ImageOffsets) GetXOk() (*IntegerWithUnits, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *ImageOffsets) SetX(v IntegerWithUnits) {
	o.X = v
}

// GetY returns the Y field value
func (o *ImageOffsets) GetY() IntegerWithUnits {
	if o == nil {
		var ret IntegerWithUnits
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *ImageOffsets) GetYOk() (*IntegerWithUnits, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *ImageOffsets) SetY(v IntegerWithUnits) {
	o.Y = v
}

func (o ImageOffsets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	return toSerialize, nil
}

type NullableImageOffsets struct {
	value *ImageOffsets
	isSet bool
}

func (v NullableImageOffsets) Get() *ImageOffsets {
	return v.value
}

func (v *NullableImageOffsets) Set(val *ImageOffsets) {
	v.value = val
	v.isSet = true
}

func (v NullableImageOffsets) IsSet() bool {
	return v.isSet
}

func (v *NullableImageOffsets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageOffsets(val *ImageOffsets) *NullableImageOffsets {
	return &NullableImageOffsets{value: val, isSet: true}
}

func (v NullableImageOffsets) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableImageOffsets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
