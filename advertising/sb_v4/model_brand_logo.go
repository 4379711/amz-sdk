package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the BrandLogo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandLogo{}

// BrandLogo Properties associated with Brand Logo.
type BrandLogo struct {
	BrandLogoCrop *AssetCrop `json:"brandLogoCrop,omitempty"`
	BrandLogoUrl  *string    `json:"brandLogoUrl,omitempty"`
	// The identifier of image/video asset from the store's asset library
	BrandLogoAssetId *string `json:"brandLogoAssetId,omitempty"`
}

// NewBrandLogo instantiates a new BrandLogo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandLogo() *BrandLogo {
	this := BrandLogo{}
	return &this
}

// NewBrandLogoWithDefaults instantiates a new BrandLogo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandLogoWithDefaults() *BrandLogo {
	this := BrandLogo{}
	return &this
}

// GetBrandLogoCrop returns the BrandLogoCrop field value if set, zero value otherwise.
func (o *BrandLogo) GetBrandLogoCrop() AssetCrop {
	if o == nil || IsNil(o.BrandLogoCrop) {
		var ret AssetCrop
		return ret
	}
	return *o.BrandLogoCrop
}

// GetBrandLogoCropOk returns a tuple with the BrandLogoCrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandLogo) GetBrandLogoCropOk() (*AssetCrop, bool) {
	if o == nil || IsNil(o.BrandLogoCrop) {
		return nil, false
	}
	return o.BrandLogoCrop, true
}

// HasBrandLogoCrop returns a boolean if a field has been set.
func (o *BrandLogo) HasBrandLogoCrop() bool {
	if o != nil && !IsNil(o.BrandLogoCrop) {
		return true
	}

	return false
}

// SetBrandLogoCrop gets a reference to the given AssetCrop and assigns it to the BrandLogoCrop field.
func (o *BrandLogo) SetBrandLogoCrop(v AssetCrop) {
	o.BrandLogoCrop = &v
}

// GetBrandLogoUrl returns the BrandLogoUrl field value if set, zero value otherwise.
func (o *BrandLogo) GetBrandLogoUrl() string {
	if o == nil || IsNil(o.BrandLogoUrl) {
		var ret string
		return ret
	}
	return *o.BrandLogoUrl
}

// GetBrandLogoUrlOk returns a tuple with the BrandLogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandLogo) GetBrandLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BrandLogoUrl) {
		return nil, false
	}
	return o.BrandLogoUrl, true
}

// HasBrandLogoUrl returns a boolean if a field has been set.
func (o *BrandLogo) HasBrandLogoUrl() bool {
	if o != nil && !IsNil(o.BrandLogoUrl) {
		return true
	}

	return false
}

// SetBrandLogoUrl gets a reference to the given string and assigns it to the BrandLogoUrl field.
func (o *BrandLogo) SetBrandLogoUrl(v string) {
	o.BrandLogoUrl = &v
}

// GetBrandLogoAssetId returns the BrandLogoAssetId field value if set, zero value otherwise.
func (o *BrandLogo) GetBrandLogoAssetId() string {
	if o == nil || IsNil(o.BrandLogoAssetId) {
		var ret string
		return ret
	}
	return *o.BrandLogoAssetId
}

// GetBrandLogoAssetIdOk returns a tuple with the BrandLogoAssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandLogo) GetBrandLogoAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrandLogoAssetId) {
		return nil, false
	}
	return o.BrandLogoAssetId, true
}

// HasBrandLogoAssetId returns a boolean if a field has been set.
func (o *BrandLogo) HasBrandLogoAssetId() bool {
	if o != nil && !IsNil(o.BrandLogoAssetId) {
		return true
	}

	return false
}

// SetBrandLogoAssetId gets a reference to the given string and assigns it to the BrandLogoAssetId field.
func (o *BrandLogo) SetBrandLogoAssetId(v string) {
	o.BrandLogoAssetId = &v
}

func (o BrandLogo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandLogoCrop) {
		toSerialize["brandLogoCrop"] = o.BrandLogoCrop
	}
	if !IsNil(o.BrandLogoUrl) {
		toSerialize["brandLogoUrl"] = o.BrandLogoUrl
	}
	if !IsNil(o.BrandLogoAssetId) {
		toSerialize["brandLogoAssetId"] = o.BrandLogoAssetId
	}
	return toSerialize, nil
}

type NullableBrandLogo struct {
	value *BrandLogo
	isSet bool
}

func (v NullableBrandLogo) Get() *BrandLogo {
	return v.value
}

func (v *NullableBrandLogo) Set(val *BrandLogo) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandLogo) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandLogo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandLogo(val *BrandLogo) *NullableBrandLogo {
	return &NullableBrandLogo{value: val, isSet: true}
}

func (v NullableBrandLogo) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandLogo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
