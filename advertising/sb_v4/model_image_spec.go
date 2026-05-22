package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ImageSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageSpec{}

// ImageSpec Structure for Image specification
type ImageSpec struct {
	// Image resolution, default is 1200x628. New values will be added later. |   Resolution  |   Value       | |---------------|---------------| |   1200x628  |   1200x628  |
	Resolution *string `json:"resolution,omitempty"`
	// Valid values are PNG and JPEG, default is PNG. New values will be added later. |   File Format  |   Value       | |---------------|---------------| |   PNG          |   PNG         | |   JPEG         |   JPEG        |
	FileFormat *string `json:"fileFormat,omitempty"`
}

// NewImageSpec instantiates a new ImageSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageSpec() *ImageSpec {
	this := ImageSpec{}
	return &this
}

// NewImageSpecWithDefaults instantiates a new ImageSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageSpecWithDefaults() *ImageSpec {
	this := ImageSpec{}
	return &this
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *ImageSpec) GetResolution() string {
	if o == nil || IsNil(o.Resolution) {
		var ret string
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageSpec) GetResolutionOk() (*string, bool) {
	if o == nil || IsNil(o.Resolution) {
		return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *ImageSpec) HasResolution() bool {
	if o != nil && !IsNil(o.Resolution) {
		return true
	}

	return false
}

// SetResolution gets a reference to the given string and assigns it to the Resolution field.
func (o *ImageSpec) SetResolution(v string) {
	o.Resolution = &v
}

// GetFileFormat returns the FileFormat field value if set, zero value otherwise.
func (o *ImageSpec) GetFileFormat() string {
	if o == nil || IsNil(o.FileFormat) {
		var ret string
		return ret
	}
	return *o.FileFormat
}

// GetFileFormatOk returns a tuple with the FileFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageSpec) GetFileFormatOk() (*string, bool) {
	if o == nil || IsNil(o.FileFormat) {
		return nil, false
	}
	return o.FileFormat, true
}

// HasFileFormat returns a boolean if a field has been set.
func (o *ImageSpec) HasFileFormat() bool {
	if o != nil && !IsNil(o.FileFormat) {
		return true
	}

	return false
}

// SetFileFormat gets a reference to the given string and assigns it to the FileFormat field.
func (o *ImageSpec) SetFileFormat(v string) {
	o.FileFormat = &v
}

func (o ImageSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resolution) {
		toSerialize["resolution"] = o.Resolution
	}
	if !IsNil(o.FileFormat) {
		toSerialize["fileFormat"] = o.FileFormat
	}
	return toSerialize, nil
}

type NullableImageSpec struct {
	value *ImageSpec
	isSet bool
}

func (v NullableImageSpec) Get() *ImageSpec {
	return v.value
}

func (v *NullableImageSpec) Set(val *ImageSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableImageSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableImageSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageSpec(val *ImageSpec) *NullableImageSpec {
	return &NullableImageSpec{value: val, isSet: true}
}

func (v NullableImageSpec) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableImageSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
