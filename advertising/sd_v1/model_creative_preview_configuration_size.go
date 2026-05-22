package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativePreviewConfigurationSize type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativePreviewConfigurationSize{}

// CreativePreviewConfigurationSize The slot dimension to render the creative. Sponsored Display creatives are responsive to a limited list of width and height pairs, including 300x250, 650x130, 245x250, 414x125, 600x160, 600x300, 728x90, 980x55, 320x50, 970x250 and 270x150.
type CreativePreviewConfigurationSize struct {
	Width  *int32 `json:"width,omitempty"`
	Height *int32 `json:"height,omitempty"`
}

// NewCreativePreviewConfigurationSize instantiates a new CreativePreviewConfigurationSize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativePreviewConfigurationSize() *CreativePreviewConfigurationSize {
	this := CreativePreviewConfigurationSize{}
	return &this
}

// NewCreativePreviewConfigurationSizeWithDefaults instantiates a new CreativePreviewConfigurationSize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativePreviewConfigurationSizeWithDefaults() *CreativePreviewConfigurationSize {
	this := CreativePreviewConfigurationSize{}
	return &this
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *CreativePreviewConfigurationSize) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativePreviewConfigurationSize) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *CreativePreviewConfigurationSize) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *CreativePreviewConfigurationSize) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *CreativePreviewConfigurationSize) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativePreviewConfigurationSize) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *CreativePreviewConfigurationSize) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *CreativePreviewConfigurationSize) SetHeight(v int32) {
	o.Height = &v
}

func (o CreativePreviewConfigurationSize) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	return toSerialize, nil
}

type NullableCreativePreviewConfigurationSize struct {
	value *CreativePreviewConfigurationSize
	isSet bool
}

func (v NullableCreativePreviewConfigurationSize) Get() *CreativePreviewConfigurationSize {
	return v.value
}

func (v *NullableCreativePreviewConfigurationSize) Set(val *CreativePreviewConfigurationSize) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativePreviewConfigurationSize) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativePreviewConfigurationSize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativePreviewConfigurationSize(val *CreativePreviewConfigurationSize) *NullableCreativePreviewConfigurationSize {
	return &NullableCreativePreviewConfigurationSize{value: val, isSet: true}
}

func (v NullableCreativePreviewConfigurationSize) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativePreviewConfigurationSize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
