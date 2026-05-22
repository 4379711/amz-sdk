package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the VideoCreative type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoCreative{}

// VideoCreative struct for VideoCreative
type VideoCreative struct {
	// If set to true and the heaadline and/or video are not in the marketplace's default language, Amazon will attempt to translate them to the marketplace's default language. If Amazon is unable to translate them, the ad will be rejected by moderation. We only support translating headlines and videos from English to German, French, Italian, Spanish, Japanese, and Dutch. See developer notes for more information.
	ConsentToTranslate *bool `json:"consentToTranslate,omitempty"`
	// The assetIds of the original videos submitted by the advertiser. If 'consentToTranslate' is set to true and translation is SUCCESSFUL then 'videoAssetIds' will return translated video assetId whereas `originalVideoAssetIds` will return the original video assetId. In all other cases, `videoAssetIds` will return original video assetId.
	VideoAssetIds []string `json:"videoAssetIds"`
}

type _VideoCreative VideoCreative

// NewVideoCreative instantiates a new VideoCreative object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoCreative(videoAssetIds []string) *VideoCreative {
	this := VideoCreative{}
	this.VideoAssetIds = videoAssetIds
	return &this
}

// NewVideoCreativeWithDefaults instantiates a new VideoCreative object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoCreativeWithDefaults() *VideoCreative {
	this := VideoCreative{}
	return &this
}

// GetConsentToTranslate returns the ConsentToTranslate field value if set, zero value otherwise.
func (o *VideoCreative) GetConsentToTranslate() bool {
	if o == nil || IsNil(o.ConsentToTranslate) {
		var ret bool
		return ret
	}
	return *o.ConsentToTranslate
}

// GetConsentToTranslateOk returns a tuple with the ConsentToTranslate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoCreative) GetConsentToTranslateOk() (*bool, bool) {
	if o == nil || IsNil(o.ConsentToTranslate) {
		return nil, false
	}
	return o.ConsentToTranslate, true
}

// HasConsentToTranslate returns a boolean if a field has been set.
func (o *VideoCreative) HasConsentToTranslate() bool {
	if o != nil && !IsNil(o.ConsentToTranslate) {
		return true
	}

	return false
}

// SetConsentToTranslate gets a reference to the given bool and assigns it to the ConsentToTranslate field.
func (o *VideoCreative) SetConsentToTranslate(v bool) {
	o.ConsentToTranslate = &v
}

// GetVideoAssetIds returns the VideoAssetIds field value
func (o *VideoCreative) GetVideoAssetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.VideoAssetIds
}

// GetVideoAssetIdsOk returns a tuple with the VideoAssetIds field value
// and a boolean to check if the value has been set.
func (o *VideoCreative) GetVideoAssetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoAssetIds, true
}

// SetVideoAssetIds sets field value
func (o *VideoCreative) SetVideoAssetIds(v []string) {
	o.VideoAssetIds = v
}

func (o VideoCreative) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConsentToTranslate) {
		toSerialize["consentToTranslate"] = o.ConsentToTranslate
	}
	toSerialize["videoAssetIds"] = o.VideoAssetIds
	return toSerialize, nil
}

type NullableVideoCreative struct {
	value *VideoCreative
	isSet bool
}

func (v NullableVideoCreative) Get() *VideoCreative {
	return v.value
}

func (v *NullableVideoCreative) Set(val *VideoCreative) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoCreative) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoCreative) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoCreative(val *VideoCreative) *NullableVideoCreative {
	return &NullableVideoCreative{value: val, isSet: true}
}

func (v NullableVideoCreative) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableVideoCreative) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
