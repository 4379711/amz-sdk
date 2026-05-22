package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeImageRecommendationEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeImageRecommendationEntry{}

// CreativeImageRecommendationEntry struct for CreativeImageRecommendationEntry
type CreativeImageRecommendationEntry struct {
	// Recommendations with higher values are more relevant
	Score *float64 `json:"score,omitempty"`
	// The asset size in bytes
	SizeInBytes *float32 `json:"sizeInBytes,omitempty"`
	// The identifier of image/video asset from the store's asset library
	AssetId *string `json:"assetId,omitempty"`
	// The URL of the asset
	ImageUrl *string `json:"imageUrl,omitempty"`
	// The width of the asset in pixels
	Width *float32 `json:"width,omitempty"`
	// The fileName of the asset
	Name        *string    `json:"name,omitempty"`
	ContentType *MediaType `json:"contentType,omitempty"`
	// The height of the asset in pixels
	Height *float32 `json:"height,omitempty"`
}

// NewCreativeImageRecommendationEntry instantiates a new CreativeImageRecommendationEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeImageRecommendationEntry() *CreativeImageRecommendationEntry {
	this := CreativeImageRecommendationEntry{}
	return &this
}

// NewCreativeImageRecommendationEntryWithDefaults instantiates a new CreativeImageRecommendationEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeImageRecommendationEntryWithDefaults() *CreativeImageRecommendationEntry {
	this := CreativeImageRecommendationEntry{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *CreativeImageRecommendationEntry) GetScore() float64 {
	if o == nil || IsNil(o.Score) {
		var ret float64
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeImageRecommendationEntry) GetScoreOk() (*float64, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *CreativeImageRecommendationEntry) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float64 and assigns it to the Score field.
func (o *CreativeImageRecommendationEntry) SetScore(v float64) {
	o.Score = &v
}

// GetSizeInBytes returns the SizeInBytes field value if set, zero value otherwise.
func (o *CreativeImageRecommendationEntry) GetSizeInBytes() float32 {
	if o == nil || IsNil(o.SizeInBytes) {
		var ret float32
		return ret
	}
	return *o.SizeInBytes
}

// GetSizeInBytesOk returns a tuple with the SizeInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeImageRecommendationEntry) GetSizeInBytesOk() (*float32, bool) {
	if o == nil || IsNil(o.SizeInBytes) {
		return nil, false
	}
	return o.SizeInBytes, true
}

// HasSizeInBytes returns a boolean if a field has been set.
func (o *CreativeImageRecommendationEntry) HasSizeInBytes() bool {
	if o != nil && !IsNil(o.SizeInBytes) {
		return true
	}

	return false
}

// SetSizeInBytes gets a reference to the given float32 and assigns it to the SizeInBytes field.
func (o *CreativeImageRecommendationEntry) SetSizeInBytes(v float32) {
	o.SizeInBytes = &v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *CreativeImageRecommendationEntry) GetAssetId() string {
	if o == nil || IsNil(o.AssetId) {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeImageRecommendationEntry) GetAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssetId) {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *CreativeImageRecommendationEntry) HasAssetId() bool {
	if o != nil && !IsNil(o.AssetId) {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *CreativeImageRecommendationEntry) SetAssetId(v string) {
	o.AssetId = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *CreativeImageRecommendationEntry) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeImageRecommendationEntry) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *CreativeImageRecommendationEntry) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *CreativeImageRecommendationEntry) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *CreativeImageRecommendationEntry) GetWidth() float32 {
	if o == nil || IsNil(o.Width) {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeImageRecommendationEntry) GetWidthOk() (*float32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *CreativeImageRecommendationEntry) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *CreativeImageRecommendationEntry) SetWidth(v float32) {
	o.Width = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreativeImageRecommendationEntry) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeImageRecommendationEntry) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreativeImageRecommendationEntry) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreativeImageRecommendationEntry) SetName(v string) {
	o.Name = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *CreativeImageRecommendationEntry) GetContentType() MediaType {
	if o == nil || IsNil(o.ContentType) {
		var ret MediaType
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeImageRecommendationEntry) GetContentTypeOk() (*MediaType, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *CreativeImageRecommendationEntry) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given MediaType and assigns it to the ContentType field.
func (o *CreativeImageRecommendationEntry) SetContentType(v MediaType) {
	o.ContentType = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *CreativeImageRecommendationEntry) GetHeight() float32 {
	if o == nil || IsNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeImageRecommendationEntry) GetHeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *CreativeImageRecommendationEntry) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *CreativeImageRecommendationEntry) SetHeight(v float32) {
	o.Height = &v
}

func (o CreativeImageRecommendationEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.SizeInBytes) {
		toSerialize["sizeInBytes"] = o.SizeInBytes
	}
	if !IsNil(o.AssetId) {
		toSerialize["assetId"] = o.AssetId
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ContentType) {
		toSerialize["contentType"] = o.ContentType
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	return toSerialize, nil
}

type NullableCreativeImageRecommendationEntry struct {
	value *CreativeImageRecommendationEntry
	isSet bool
}

func (v NullableCreativeImageRecommendationEntry) Get() *CreativeImageRecommendationEntry {
	return v.value
}

func (v *NullableCreativeImageRecommendationEntry) Set(val *CreativeImageRecommendationEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeImageRecommendationEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeImageRecommendationEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeImageRecommendationEntry(val *CreativeImageRecommendationEntry) *NullableCreativeImageRecommendationEntry {
	return &NullableCreativeImageRecommendationEntry{value: val, isSet: true}
}

func (v NullableCreativeImageRecommendationEntry) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeImageRecommendationEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
