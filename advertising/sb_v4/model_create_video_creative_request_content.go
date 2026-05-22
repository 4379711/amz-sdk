package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateVideoCreativeRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVideoCreativeRequestContent{}

// CreateVideoCreativeRequestContent struct for CreateVideoCreativeRequestContent
type CreateVideoCreativeRequestContent struct {
	// The unique ID of a Sponsored Brands ad.
	AdId     string        `json:"adId"`
	Creative VideoCreative `json:"creative"`
}

type _CreateVideoCreativeRequestContent CreateVideoCreativeRequestContent

// NewCreateVideoCreativeRequestContent instantiates a new CreateVideoCreativeRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVideoCreativeRequestContent(adId string, creative VideoCreative) *CreateVideoCreativeRequestContent {
	this := CreateVideoCreativeRequestContent{}
	this.AdId = adId
	this.Creative = creative
	return &this
}

// NewCreateVideoCreativeRequestContentWithDefaults instantiates a new CreateVideoCreativeRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVideoCreativeRequestContentWithDefaults() *CreateVideoCreativeRequestContent {
	this := CreateVideoCreativeRequestContent{}
	return &this
}

// GetAdId returns the AdId field value
func (o *CreateVideoCreativeRequestContent) GetAdId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value
// and a boolean to check if the value has been set.
func (o *CreateVideoCreativeRequestContent) GetAdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdId, true
}

// SetAdId sets field value
func (o *CreateVideoCreativeRequestContent) SetAdId(v string) {
	o.AdId = v
}

// GetCreative returns the Creative field value
func (o *CreateVideoCreativeRequestContent) GetCreative() VideoCreative {
	if o == nil {
		var ret VideoCreative
		return ret
	}

	return o.Creative
}

// GetCreativeOk returns a tuple with the Creative field value
// and a boolean to check if the value has been set.
func (o *CreateVideoCreativeRequestContent) GetCreativeOk() (*VideoCreative, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creative, true
}

// SetCreative sets field value
func (o *CreateVideoCreativeRequestContent) SetCreative(v VideoCreative) {
	o.Creative = v
}

func (o CreateVideoCreativeRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adId"] = o.AdId
	toSerialize["creative"] = o.Creative
	return toSerialize, nil
}

type NullableCreateVideoCreativeRequestContent struct {
	value *CreateVideoCreativeRequestContent
	isSet bool
}

func (v NullableCreateVideoCreativeRequestContent) Get() *CreateVideoCreativeRequestContent {
	return v.value
}

func (v *NullableCreateVideoCreativeRequestContent) Set(val *CreateVideoCreativeRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVideoCreativeRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVideoCreativeRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVideoCreativeRequestContent(val *CreateVideoCreativeRequestContent) *NullableCreateVideoCreativeRequestContent {
	return &NullableCreateVideoCreativeRequestContent{value: val, isSet: true}
}

func (v NullableCreateVideoCreativeRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateVideoCreativeRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
