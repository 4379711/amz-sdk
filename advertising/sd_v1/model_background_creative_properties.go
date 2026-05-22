package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BackgroundCreativeProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackgroundCreativeProperties{}

// BackgroundCreativeProperties User-customizable properties of a creative with background. Only supported for productAds with landingPageType of OFF_AMAZON_LINK.
type BackgroundCreativeProperties struct {
	// An optional collection of backgrounds which are displayed on the ad.
	Backgrounds []Background `json:"backgrounds,omitempty"`
}

// NewBackgroundCreativeProperties instantiates a new BackgroundCreativeProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackgroundCreativeProperties() *BackgroundCreativeProperties {
	this := BackgroundCreativeProperties{}
	return &this
}

// NewBackgroundCreativePropertiesWithDefaults instantiates a new BackgroundCreativeProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackgroundCreativePropertiesWithDefaults() *BackgroundCreativeProperties {
	this := BackgroundCreativeProperties{}
	return &this
}

// GetBackgrounds returns the Backgrounds field value if set, zero value otherwise.
func (o *BackgroundCreativeProperties) GetBackgrounds() []Background {
	if o == nil || IsNil(o.Backgrounds) {
		var ret []Background
		return ret
	}
	return o.Backgrounds
}

// GetBackgroundsOk returns a tuple with the Backgrounds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackgroundCreativeProperties) GetBackgroundsOk() ([]Background, bool) {
	if o == nil || IsNil(o.Backgrounds) {
		return nil, false
	}
	return o.Backgrounds, true
}

// HasBackgrounds returns a boolean if a field has been set.
func (o *BackgroundCreativeProperties) HasBackgrounds() bool {
	if o != nil && !IsNil(o.Backgrounds) {
		return true
	}

	return false
}

// SetBackgrounds gets a reference to the given []Background and assigns it to the Backgrounds field.
func (o *BackgroundCreativeProperties) SetBackgrounds(v []Background) {
	o.Backgrounds = v
}

func (o BackgroundCreativeProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Backgrounds) {
		toSerialize["backgrounds"] = o.Backgrounds
	}
	return toSerialize, nil
}

type NullableBackgroundCreativeProperties struct {
	value *BackgroundCreativeProperties
	isSet bool
}

func (v NullableBackgroundCreativeProperties) Get() *BackgroundCreativeProperties {
	return v.value
}

func (v *NullableBackgroundCreativeProperties) Set(val *BackgroundCreativeProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundCreativeProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundCreativeProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundCreativeProperties(val *BackgroundCreativeProperties) *NullableBackgroundCreativeProperties {
	return &NullableBackgroundCreativeProperties{value: val, isSet: true}
}

func (v NullableBackgroundCreativeProperties) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBackgroundCreativeProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
