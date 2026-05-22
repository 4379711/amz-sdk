package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the Video type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Video{}

// Video This field denotes video which is displayed on the ad. This field is optional and mutable. A video asset must be provided for a VIDEO creative. Specific restrictions based on the video are listed in the following table. ||Specifications| |------------------|------------------| |Maximum file size|500MB| |Aspect ratio|16:9| |Minimum duration|6s| |Maximum duration|45s| |Minimum frame size|1920x1080| |Minimum video bitrate|4mbps| |Video frame rate(fps)|23.976(recommended), 24, 25, or 29.97| |Video frame rate mode|Constant| |Minimum audio bitrate|192kbps| |Audio sample rate|44.1kHz or 48kHz| |Supported Formats|Video: H.264, MPEG-2, or MPEG-4; Audio: PCM or AAC| |Audio Channel|Audio format needs to be stereo or mono.| |Recommended video bitrate|8mbps| |Recommended duration|A duration of exactly 6s, 15s, 20s, or 30s is recommended. Use of videos outside of these durations may negatively impact your campaign performance. Shorter lengths will drive higher VCR (although scale on 6s may be limited).|
type Video struct {
	// The unique identifier of the video asset. This assetId comes from the Creative Asset Library.
	AssetId string `json:"assetId"`
	// The identifier of the particular video assetversion.
	AssetVersion string `json:"assetVersion"`
}

type _Video Video

// NewVideo instantiates a new Video object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideo(assetId string, assetVersion string) *Video {
	this := Video{}
	this.AssetId = assetId
	this.AssetVersion = assetVersion
	return &this
}

// NewVideoWithDefaults instantiates a new Video object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoWithDefaults() *Video {
	this := Video{}
	return &this
}

// GetAssetId returns the AssetId field value
func (o *Video) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *Video) GetAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *Video) SetAssetId(v string) {
	o.AssetId = v
}

// GetAssetVersion returns the AssetVersion field value
func (o *Video) GetAssetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetVersion
}

// GetAssetVersionOk returns a tuple with the AssetVersion field value
// and a boolean to check if the value has been set.
func (o *Video) GetAssetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetVersion, true
}

// SetAssetVersion sets field value
func (o *Video) SetAssetVersion(v string) {
	o.AssetVersion = v
}

func (o Video) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assetId"] = o.AssetId
	toSerialize["assetVersion"] = o.AssetVersion
	return toSerialize, nil
}

type NullableVideo struct {
	value *Video
	isSet bool
}

func (v NullableVideo) Get() *Video {
	return v.value
}

func (v *NullableVideo) Set(val *Video) {
	v.value = val
	v.isSet = true
}

func (v NullableVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideo(val *Video) *NullableVideo {
	return &NullableVideo{value: val, isSet: true}
}

func (v NullableVideo) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
