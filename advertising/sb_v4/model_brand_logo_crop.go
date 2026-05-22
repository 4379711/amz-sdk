package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the BrandLogoCrop type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandLogoCrop{}

// BrandLogoCrop The crop to apply to the selected Brand logo. A Brand logo must have minimum dimensions of 400x400. If a brandLogoAssetID is supplied but a crop is not, the crop will be defaulted to the whole image.
type BrandLogoCrop struct {
	Top    *float32 `json:"top,omitempty"`
	Left   *float32 `json:"left,omitempty"`
	Width  *float32 `json:"width,omitempty"`
	Height *float32 `json:"height,omitempty"`
}

// NewBrandLogoCrop instantiates a new BrandLogoCrop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandLogoCrop() *BrandLogoCrop {
	this := BrandLogoCrop{}
	return &this
}

// NewBrandLogoCropWithDefaults instantiates a new BrandLogoCrop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandLogoCropWithDefaults() *BrandLogoCrop {
	this := BrandLogoCrop{}
	return &this
}

// GetTop returns the Top field value if set, zero value otherwise.
func (o *BrandLogoCrop) GetTop() float32 {
	if o == nil || IsNil(o.Top) {
		var ret float32
		return ret
	}
	return *o.Top
}

// GetTopOk returns a tuple with the Top field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandLogoCrop) GetTopOk() (*float32, bool) {
	if o == nil || IsNil(o.Top) {
		return nil, false
	}
	return o.Top, true
}

// HasTop returns a boolean if a field has been set.
func (o *BrandLogoCrop) HasTop() bool {
	if o != nil && !IsNil(o.Top) {
		return true
	}

	return false
}

// SetTop gets a reference to the given float32 and assigns it to the Top field.
func (o *BrandLogoCrop) SetTop(v float32) {
	o.Top = &v
}

// GetLeft returns the Left field value if set, zero value otherwise.
func (o *BrandLogoCrop) GetLeft() float32 {
	if o == nil || IsNil(o.Left) {
		var ret float32
		return ret
	}
	return *o.Left
}

// GetLeftOk returns a tuple with the Left field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandLogoCrop) GetLeftOk() (*float32, bool) {
	if o == nil || IsNil(o.Left) {
		return nil, false
	}
	return o.Left, true
}

// HasLeft returns a boolean if a field has been set.
func (o *BrandLogoCrop) HasLeft() bool {
	if o != nil && !IsNil(o.Left) {
		return true
	}

	return false
}

// SetLeft gets a reference to the given float32 and assigns it to the Left field.
func (o *BrandLogoCrop) SetLeft(v float32) {
	o.Left = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *BrandLogoCrop) GetWidth() float32 {
	if o == nil || IsNil(o.Width) {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandLogoCrop) GetWidthOk() (*float32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *BrandLogoCrop) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *BrandLogoCrop) SetWidth(v float32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *BrandLogoCrop) GetHeight() float32 {
	if o == nil || IsNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandLogoCrop) GetHeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *BrandLogoCrop) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *BrandLogoCrop) SetHeight(v float32) {
	o.Height = &v
}

func (o BrandLogoCrop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Top) {
		toSerialize["top"] = o.Top
	}
	if !IsNil(o.Left) {
		toSerialize["left"] = o.Left
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	return toSerialize, nil
}

type NullableBrandLogoCrop struct {
	value *BrandLogoCrop
	isSet bool
}

func (v NullableBrandLogoCrop) Get() *BrandLogoCrop {
	return v.value
}

func (v *NullableBrandLogoCrop) Set(val *BrandLogoCrop) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandLogoCrop) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandLogoCrop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandLogoCrop(val *BrandLogoCrop) *NullableBrandLogoCrop {
	return &NullableBrandLogoCrop{value: val, isSet: true}
}

func (v NullableBrandLogoCrop) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandLogoCrop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
