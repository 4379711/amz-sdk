package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateVideoCreative type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVideoCreative{}

// CreateVideoCreative struct for CreateVideoCreative
type CreateVideoCreative struct {
	Asins []string `json:"asins,omitempty"`
	// If set to true and the headline and/or video are not in the marketplace's default language, Amazon will attempt to translate them to the marketplace's default language. If Amazon is unable to translate them, the ad will be rejected by moderation. We only support translating headlines and videos from English to German, French, Italian, Spanish, Japanese, and Dutch. See developer notes for more information.
	ConsentToTranslate *bool `json:"consentToTranslate,omitempty"`
	// In SB API V4, `videoMediaIds` is replaced by `videoAssetIds`. `videoAssetIds` will only allow Asset Library identifiers for ad creation, but responses can include mediaIds for v1 campaigns and API V3 operations. At a future state, existing mediaIds will be added to Asset library for use in SB campaigns.
	VideoAssetIds []string `json:"videoAssetIds,omitempty"`
}

// NewCreateVideoCreative instantiates a new CreateVideoCreative object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVideoCreative() *CreateVideoCreative {
	this := CreateVideoCreative{}
	return &this
}

// NewCreateVideoCreativeWithDefaults instantiates a new CreateVideoCreative object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVideoCreativeWithDefaults() *CreateVideoCreative {
	this := CreateVideoCreative{}
	return &this
}

// GetAsins returns the Asins field value if set, zero value otherwise.
func (o *CreateVideoCreative) GetAsins() []string {
	if o == nil || IsNil(o.Asins) {
		var ret []string
		return ret
	}
	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVideoCreative) GetAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.Asins) {
		return nil, false
	}
	return o.Asins, true
}

// HasAsins returns a boolean if a field has been set.
func (o *CreateVideoCreative) HasAsins() bool {
	if o != nil && !IsNil(o.Asins) {
		return true
	}

	return false
}

// SetAsins gets a reference to the given []string and assigns it to the Asins field.
func (o *CreateVideoCreative) SetAsins(v []string) {
	o.Asins = v
}

// GetConsentToTranslate returns the ConsentToTranslate field value if set, zero value otherwise.
func (o *CreateVideoCreative) GetConsentToTranslate() bool {
	if o == nil || IsNil(o.ConsentToTranslate) {
		var ret bool
		return ret
	}
	return *o.ConsentToTranslate
}

// GetConsentToTranslateOk returns a tuple with the ConsentToTranslate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVideoCreative) GetConsentToTranslateOk() (*bool, bool) {
	if o == nil || IsNil(o.ConsentToTranslate) {
		return nil, false
	}
	return o.ConsentToTranslate, true
}

// HasConsentToTranslate returns a boolean if a field has been set.
func (o *CreateVideoCreative) HasConsentToTranslate() bool {
	if o != nil && !IsNil(o.ConsentToTranslate) {
		return true
	}

	return false
}

// SetConsentToTranslate gets a reference to the given bool and assigns it to the ConsentToTranslate field.
func (o *CreateVideoCreative) SetConsentToTranslate(v bool) {
	o.ConsentToTranslate = &v
}

// GetVideoAssetIds returns the VideoAssetIds field value if set, zero value otherwise.
func (o *CreateVideoCreative) GetVideoAssetIds() []string {
	if o == nil || IsNil(o.VideoAssetIds) {
		var ret []string
		return ret
	}
	return o.VideoAssetIds
}

// GetVideoAssetIdsOk returns a tuple with the VideoAssetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVideoCreative) GetVideoAssetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.VideoAssetIds) {
		return nil, false
	}
	return o.VideoAssetIds, true
}

// HasVideoAssetIds returns a boolean if a field has been set.
func (o *CreateVideoCreative) HasVideoAssetIds() bool {
	if o != nil && !IsNil(o.VideoAssetIds) {
		return true
	}

	return false
}

// SetVideoAssetIds gets a reference to the given []string and assigns it to the VideoAssetIds field.
func (o *CreateVideoCreative) SetVideoAssetIds(v []string) {
	o.VideoAssetIds = v
}

func (o CreateVideoCreative) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asins) {
		toSerialize["asins"] = o.Asins
	}
	if !IsNil(o.ConsentToTranslate) {
		toSerialize["consentToTranslate"] = o.ConsentToTranslate
	}
	if !IsNil(o.VideoAssetIds) {
		toSerialize["videoAssetIds"] = o.VideoAssetIds
	}
	return toSerialize, nil
}

type NullableCreateVideoCreative struct {
	value *CreateVideoCreative
	isSet bool
}

func (v NullableCreateVideoCreative) Get() *CreateVideoCreative {
	return v.value
}

func (v *NullableCreateVideoCreative) Set(val *CreateVideoCreative) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVideoCreative) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVideoCreative) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVideoCreative(val *CreateVideoCreative) *NullableCreateVideoCreative {
	return &NullableCreateVideoCreative{value: val, isSet: true}
}

func (v NullableCreateVideoCreative) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateVideoCreative) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
