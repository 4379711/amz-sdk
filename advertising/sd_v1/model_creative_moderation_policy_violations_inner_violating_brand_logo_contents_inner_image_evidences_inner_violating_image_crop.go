package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop{}

// CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop struct for CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop
type CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop struct {
	// The top left X-coordinate of the content that violates the specfied policy within the image.
	TopLeftX *int32 `json:"topLeftX,omitempty"`
	// The top left Y-coordinate of the content that violates the specfied policy within the image.
	TopLeftY *int32 `json:"topLeftY,omitempty"`
	// The height of the content that violates the specfied policy within the image.
	Height *int32 `json:"height,omitempty"`
	// The width of the content that violates the specfied policy within the image.
	Width *int32 `json:"width,omitempty"`
}

// NewCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop instantiates a new CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop() *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop {
	this := CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop{}
	return &this
}

// NewCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCropWithDefaults instantiates a new CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCropWithDefaults() *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop {
	this := CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop{}
	return &this
}

// GetTopLeftX returns the TopLeftX field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) GetTopLeftX() int32 {
	if o == nil || IsNil(o.TopLeftX) {
		var ret int32
		return ret
	}
	return *o.TopLeftX
}

// GetTopLeftXOk returns a tuple with the TopLeftX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) GetTopLeftXOk() (*int32, bool) {
	if o == nil || IsNil(o.TopLeftX) {
		return nil, false
	}
	return o.TopLeftX, true
}

// HasTopLeftX returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) HasTopLeftX() bool {
	if o != nil && !IsNil(o.TopLeftX) {
		return true
	}

	return false
}

// SetTopLeftX gets a reference to the given int32 and assigns it to the TopLeftX field.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) SetTopLeftX(v int32) {
	o.TopLeftX = &v
}

// GetTopLeftY returns the TopLeftY field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) GetTopLeftY() int32 {
	if o == nil || IsNil(o.TopLeftY) {
		var ret int32
		return ret
	}
	return *o.TopLeftY
}

// GetTopLeftYOk returns a tuple with the TopLeftY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) GetTopLeftYOk() (*int32, bool) {
	if o == nil || IsNil(o.TopLeftY) {
		return nil, false
	}
	return o.TopLeftY, true
}

// HasTopLeftY returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) HasTopLeftY() bool {
	if o != nil && !IsNil(o.TopLeftY) {
		return true
	}

	return false
}

// SetTopLeftY gets a reference to the given int32 and assigns it to the TopLeftY field.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) SetTopLeftY(v int32) {
	o.TopLeftY = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) SetHeight(v int32) {
	o.Height = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) SetWidth(v int32) {
	o.Width = &v
}

func (o CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TopLeftX) {
		toSerialize["topLeftX"] = o.TopLeftX
	}
	if !IsNil(o.TopLeftY) {
		toSerialize["topLeftY"] = o.TopLeftY
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	return toSerialize, nil
}

type NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop struct {
	value *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop
	isSet bool
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) Get() *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop {
	return v.value
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) Set(val *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop(val *CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) *NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop {
	return &NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop{value: val, isSet: true}
}

func (v NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInnerImageEvidencesInnerViolatingImageCrop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
