package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the VideoCreativeProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoCreativeProperties{}

// VideoCreativeProperties User-customizable properties of a video creative.
type VideoCreativeProperties struct {
	Video *Video `json:"video,omitempty"`
}

// NewVideoCreativeProperties instantiates a new VideoCreativeProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoCreativeProperties() *VideoCreativeProperties {
	this := VideoCreativeProperties{}
	return &this
}

// NewVideoCreativePropertiesWithDefaults instantiates a new VideoCreativeProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoCreativePropertiesWithDefaults() *VideoCreativeProperties {
	this := VideoCreativeProperties{}
	return &this
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *VideoCreativeProperties) GetVideo() Video {
	if o == nil || IsNil(o.Video) {
		var ret Video
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoCreativeProperties) GetVideoOk() (*Video, bool) {
	if o == nil || IsNil(o.Video) {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *VideoCreativeProperties) HasVideo() bool {
	if o != nil && !IsNil(o.Video) {
		return true
	}

	return false
}

// SetVideo gets a reference to the given Video and assigns it to the Video field.
func (o *VideoCreativeProperties) SetVideo(v Video) {
	o.Video = &v
}

func (o VideoCreativeProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Video) {
		toSerialize["video"] = o.Video
	}
	return toSerialize, nil
}

type NullableVideoCreativeProperties struct {
	value *VideoCreativeProperties
	isSet bool
}

func (v NullableVideoCreativeProperties) Get() *VideoCreativeProperties {
	return v.value
}

func (v *NullableVideoCreativeProperties) Set(val *VideoCreativeProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoCreativeProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoCreativeProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoCreativeProperties(val *VideoCreativeProperties) *NullableVideoCreativeProperties {
	return &NullableVideoCreativeProperties{value: val, isSet: true}
}

func (v NullableVideoCreativeProperties) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableVideoCreativeProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
