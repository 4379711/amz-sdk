package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ImageTaskMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageTaskMetadata{}

// ImageTaskMetadata struct for ImageTaskMetadata
type ImageTaskMetadata struct {
	ImageSpec *ImageSpec `json:"imageSpec,omitempty"`
	// Optional. An upper bound for number of image results for this set of metadata. Default value is 4.
	MaxResults *float32 `json:"maxResults,omitempty"`
	// Optional.
	ThemeId *string `json:"themeId,omitempty"`
	// Required. The product that is shown in AI image.
	Asin string `json:"asin"`
	// Optional. Open text prompt
	Prompt *string `json:"prompt,omitempty"`
	// Optional. Source image provided by advertiser and they are registered in Asset Library
	ProductImageAssetId *string `json:"productImageAssetId,omitempty"`
}

type _ImageTaskMetadata ImageTaskMetadata

// NewImageTaskMetadata instantiates a new ImageTaskMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageTaskMetadata(asin string) *ImageTaskMetadata {
	this := ImageTaskMetadata{}
	this.Asin = asin
	return &this
}

// NewImageTaskMetadataWithDefaults instantiates a new ImageTaskMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageTaskMetadataWithDefaults() *ImageTaskMetadata {
	this := ImageTaskMetadata{}
	return &this
}

// GetImageSpec returns the ImageSpec field value if set, zero value otherwise.
func (o *ImageTaskMetadata) GetImageSpec() ImageSpec {
	if o == nil || IsNil(o.ImageSpec) {
		var ret ImageSpec
		return ret
	}
	return *o.ImageSpec
}

// GetImageSpecOk returns a tuple with the ImageSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageTaskMetadata) GetImageSpecOk() (*ImageSpec, bool) {
	if o == nil || IsNil(o.ImageSpec) {
		return nil, false
	}
	return o.ImageSpec, true
}

// HasImageSpec returns a boolean if a field has been set.
func (o *ImageTaskMetadata) HasImageSpec() bool {
	if o != nil && !IsNil(o.ImageSpec) {
		return true
	}

	return false
}

// SetImageSpec gets a reference to the given ImageSpec and assigns it to the ImageSpec field.
func (o *ImageTaskMetadata) SetImageSpec(v ImageSpec) {
	o.ImageSpec = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *ImageTaskMetadata) GetMaxResults() float32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret float32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageTaskMetadata) GetMaxResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *ImageTaskMetadata) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given float32 and assigns it to the MaxResults field.
func (o *ImageTaskMetadata) SetMaxResults(v float32) {
	o.MaxResults = &v
}

// GetThemeId returns the ThemeId field value if set, zero value otherwise.
func (o *ImageTaskMetadata) GetThemeId() string {
	if o == nil || IsNil(o.ThemeId) {
		var ret string
		return ret
	}
	return *o.ThemeId
}

// GetThemeIdOk returns a tuple with the ThemeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageTaskMetadata) GetThemeIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThemeId) {
		return nil, false
	}
	return o.ThemeId, true
}

// HasThemeId returns a boolean if a field has been set.
func (o *ImageTaskMetadata) HasThemeId() bool {
	if o != nil && !IsNil(o.ThemeId) {
		return true
	}

	return false
}

// SetThemeId gets a reference to the given string and assigns it to the ThemeId field.
func (o *ImageTaskMetadata) SetThemeId(v string) {
	o.ThemeId = &v
}

// GetAsin returns the Asin field value
func (o *ImageTaskMetadata) GetAsin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asin
}

// GetAsinOk returns a tuple with the Asin field value
// and a boolean to check if the value has been set.
func (o *ImageTaskMetadata) GetAsinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asin, true
}

// SetAsin sets field value
func (o *ImageTaskMetadata) SetAsin(v string) {
	o.Asin = v
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *ImageTaskMetadata) GetPrompt() string {
	if o == nil || IsNil(o.Prompt) {
		var ret string
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageTaskMetadata) GetPromptOk() (*string, bool) {
	if o == nil || IsNil(o.Prompt) {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *ImageTaskMetadata) HasPrompt() bool {
	if o != nil && !IsNil(o.Prompt) {
		return true
	}

	return false
}

// SetPrompt gets a reference to the given string and assigns it to the Prompt field.
func (o *ImageTaskMetadata) SetPrompt(v string) {
	o.Prompt = &v
}

// GetProductImageAssetId returns the ProductImageAssetId field value if set, zero value otherwise.
func (o *ImageTaskMetadata) GetProductImageAssetId() string {
	if o == nil || IsNil(o.ProductImageAssetId) {
		var ret string
		return ret
	}
	return *o.ProductImageAssetId
}

// GetProductImageAssetIdOk returns a tuple with the ProductImageAssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageTaskMetadata) GetProductImageAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductImageAssetId) {
		return nil, false
	}
	return o.ProductImageAssetId, true
}

// HasProductImageAssetId returns a boolean if a field has been set.
func (o *ImageTaskMetadata) HasProductImageAssetId() bool {
	if o != nil && !IsNil(o.ProductImageAssetId) {
		return true
	}

	return false
}

// SetProductImageAssetId gets a reference to the given string and assigns it to the ProductImageAssetId field.
func (o *ImageTaskMetadata) SetProductImageAssetId(v string) {
	o.ProductImageAssetId = &v
}

func (o ImageTaskMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageSpec) {
		toSerialize["imageSpec"] = o.ImageSpec
	}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.ThemeId) {
		toSerialize["themeId"] = o.ThemeId
	}
	toSerialize["asin"] = o.Asin
	if !IsNil(o.Prompt) {
		toSerialize["prompt"] = o.Prompt
	}
	if !IsNil(o.ProductImageAssetId) {
		toSerialize["productImageAssetId"] = o.ProductImageAssetId
	}
	return toSerialize, nil
}

type NullableImageTaskMetadata struct {
	value *ImageTaskMetadata
	isSet bool
}

func (v NullableImageTaskMetadata) Get() *ImageTaskMetadata {
	return v.value
}

func (v *NullableImageTaskMetadata) Set(val *ImageTaskMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableImageTaskMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableImageTaskMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageTaskMetadata(val *ImageTaskMetadata) *NullableImageTaskMetadata {
	return &NullableImageTaskMetadata{value: val, isSet: true}
}

func (v NullableImageTaskMetadata) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableImageTaskMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
