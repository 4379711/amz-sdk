package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ImageEvidence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageEvidence{}

// ImageEvidence Structure of a image evidence
type ImageEvidence struct {
	// The top left Y-coordinate of the content that violates the specfied policy within the image.
	TopLeftY *int32 `json:"topLeftY,omitempty"`
	// The top left X-coordinate of the content that violates the specfied policy within the image.
	TopLeftX *int32 `json:"topLeftX,omitempty"`
	// The width of the content that violates the specfied policy within the image.
	Width *int32 `json:"width,omitempty"`
	// The height of the content that violates the specfied policy within the image.
	Height *int32 `json:"height,omitempty"`
}

// NewImageEvidence instantiates a new ImageEvidence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageEvidence() *ImageEvidence {
	this := ImageEvidence{}
	return &this
}

// NewImageEvidenceWithDefaults instantiates a new ImageEvidence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageEvidenceWithDefaults() *ImageEvidence {
	this := ImageEvidence{}
	return &this
}

// GetTopLeftY returns the TopLeftY field value if set, zero value otherwise.
func (o *ImageEvidence) GetTopLeftY() int32 {
	if o == nil || IsNil(o.TopLeftY) {
		var ret int32
		return ret
	}
	return *o.TopLeftY
}

// GetTopLeftYOk returns a tuple with the TopLeftY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageEvidence) GetTopLeftYOk() (*int32, bool) {
	if o == nil || IsNil(o.TopLeftY) {
		return nil, false
	}
	return o.TopLeftY, true
}

// HasTopLeftY returns a boolean if a field has been set.
func (o *ImageEvidence) HasTopLeftY() bool {
	if o != nil && !IsNil(o.TopLeftY) {
		return true
	}

	return false
}

// SetTopLeftY gets a reference to the given int32 and assigns it to the TopLeftY field.
func (o *ImageEvidence) SetTopLeftY(v int32) {
	o.TopLeftY = &v
}

// GetTopLeftX returns the TopLeftX field value if set, zero value otherwise.
func (o *ImageEvidence) GetTopLeftX() int32 {
	if o == nil || IsNil(o.TopLeftX) {
		var ret int32
		return ret
	}
	return *o.TopLeftX
}

// GetTopLeftXOk returns a tuple with the TopLeftX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageEvidence) GetTopLeftXOk() (*int32, bool) {
	if o == nil || IsNil(o.TopLeftX) {
		return nil, false
	}
	return o.TopLeftX, true
}

// HasTopLeftX returns a boolean if a field has been set.
func (o *ImageEvidence) HasTopLeftX() bool {
	if o != nil && !IsNil(o.TopLeftX) {
		return true
	}

	return false
}

// SetTopLeftX gets a reference to the given int32 and assigns it to the TopLeftX field.
func (o *ImageEvidence) SetTopLeftX(v int32) {
	o.TopLeftX = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *ImageEvidence) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageEvidence) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *ImageEvidence) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *ImageEvidence) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ImageEvidence) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageEvidence) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ImageEvidence) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *ImageEvidence) SetHeight(v int32) {
	o.Height = &v
}

func (o ImageEvidence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TopLeftY) {
		toSerialize["topLeftY"] = o.TopLeftY
	}
	if !IsNil(o.TopLeftX) {
		toSerialize["topLeftX"] = o.TopLeftX
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	return toSerialize, nil
}

type NullableImageEvidence struct {
	value *ImageEvidence
	isSet bool
}

func (v NullableImageEvidence) Get() *ImageEvidence {
	return v.value
}

func (v *NullableImageEvidence) Set(val *ImageEvidence) {
	v.value = val
	v.isSet = true
}

func (v NullableImageEvidence) IsSet() bool {
	return v.isSet
}

func (v *NullableImageEvidence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageEvidence(val *ImageEvidence) *NullableImageEvidence {
	return &NullableImageEvidence{value: val, isSet: true}
}

func (v NullableImageEvidence) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableImageEvidence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
