package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CustomImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomImage{}

// CustomImage struct for CustomImage
type CustomImage struct {
	AssetId *string          `json:"assetId,omitempty"`
	Crop    *CustomImageCrop `json:"crop,omitempty"`
	Url     *string          `json:"url,omitempty"`
}

// NewCustomImage instantiates a new CustomImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomImage() *CustomImage {
	this := CustomImage{}
	return &this
}

// NewCustomImageWithDefaults instantiates a new CustomImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomImageWithDefaults() *CustomImage {
	this := CustomImage{}
	return &this
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *CustomImage) GetAssetId() string {
	if o == nil || IsNil(o.AssetId) {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomImage) GetAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssetId) {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *CustomImage) HasAssetId() bool {
	if o != nil && !IsNil(o.AssetId) {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *CustomImage) SetAssetId(v string) {
	o.AssetId = &v
}

// GetCrop returns the Crop field value if set, zero value otherwise.
func (o *CustomImage) GetCrop() CustomImageCrop {
	if o == nil || IsNil(o.Crop) {
		var ret CustomImageCrop
		return ret
	}
	return *o.Crop
}

// GetCropOk returns a tuple with the Crop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomImage) GetCropOk() (*CustomImageCrop, bool) {
	if o == nil || IsNil(o.Crop) {
		return nil, false
	}
	return o.Crop, true
}

// HasCrop returns a boolean if a field has been set.
func (o *CustomImage) HasCrop() bool {
	if o != nil && !IsNil(o.Crop) {
		return true
	}

	return false
}

// SetCrop gets a reference to the given CustomImageCrop and assigns it to the Crop field.
func (o *CustomImage) SetCrop(v CustomImageCrop) {
	o.Crop = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CustomImage) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomImage) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CustomImage) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CustomImage) SetUrl(v string) {
	o.Url = &v
}

func (o CustomImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetId) {
		toSerialize["assetId"] = o.AssetId
	}
	if !IsNil(o.Crop) {
		toSerialize["crop"] = o.Crop
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableCustomImage struct {
	value *CustomImage
	isSet bool
}

func (v NullableCustomImage) Get() *CustomImage {
	return v.value
}

func (v *NullableCustomImage) Set(val *CustomImage) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomImage) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomImage(val *CustomImage) *NullableCustomImage {
	return &NullableCustomImage{value: val, isSet: true}
}

func (v NullableCustomImage) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCustomImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
