package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateStoreSpotlightCreativeRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateStoreSpotlightCreativeRequestContent{}

// CreateStoreSpotlightCreativeRequestContent struct for CreateStoreSpotlightCreativeRequestContent
type CreateStoreSpotlightCreativeRequestContent struct {
	// The unique ID of a Sponsored Brands ad.
	AdId     string                 `json:"adId"`
	Creative StoreSpotlightCreative `json:"creative"`
}

type _CreateStoreSpotlightCreativeRequestContent CreateStoreSpotlightCreativeRequestContent

// NewCreateStoreSpotlightCreativeRequestContent instantiates a new CreateStoreSpotlightCreativeRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStoreSpotlightCreativeRequestContent(adId string, creative StoreSpotlightCreative) *CreateStoreSpotlightCreativeRequestContent {
	this := CreateStoreSpotlightCreativeRequestContent{}
	this.AdId = adId
	this.Creative = creative
	return &this
}

// NewCreateStoreSpotlightCreativeRequestContentWithDefaults instantiates a new CreateStoreSpotlightCreativeRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStoreSpotlightCreativeRequestContentWithDefaults() *CreateStoreSpotlightCreativeRequestContent {
	this := CreateStoreSpotlightCreativeRequestContent{}
	return &this
}

// GetAdId returns the AdId field value
func (o *CreateStoreSpotlightCreativeRequestContent) GetAdId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value
// and a boolean to check if the value has been set.
func (o *CreateStoreSpotlightCreativeRequestContent) GetAdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdId, true
}

// SetAdId sets field value
func (o *CreateStoreSpotlightCreativeRequestContent) SetAdId(v string) {
	o.AdId = v
}

// GetCreative returns the Creative field value
func (o *CreateStoreSpotlightCreativeRequestContent) GetCreative() StoreSpotlightCreative {
	if o == nil {
		var ret StoreSpotlightCreative
		return ret
	}

	return o.Creative
}

// GetCreativeOk returns a tuple with the Creative field value
// and a boolean to check if the value has been set.
func (o *CreateStoreSpotlightCreativeRequestContent) GetCreativeOk() (*StoreSpotlightCreative, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creative, true
}

// SetCreative sets field value
func (o *CreateStoreSpotlightCreativeRequestContent) SetCreative(v StoreSpotlightCreative) {
	o.Creative = v
}

func (o CreateStoreSpotlightCreativeRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adId"] = o.AdId
	toSerialize["creative"] = o.Creative
	return toSerialize, nil
}

type NullableCreateStoreSpotlightCreativeRequestContent struct {
	value *CreateStoreSpotlightCreativeRequestContent
	isSet bool
}

func (v NullableCreateStoreSpotlightCreativeRequestContent) Get() *CreateStoreSpotlightCreativeRequestContent {
	return v.value
}

func (v *NullableCreateStoreSpotlightCreativeRequestContent) Set(val *CreateStoreSpotlightCreativeRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStoreSpotlightCreativeRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStoreSpotlightCreativeRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStoreSpotlightCreativeRequestContent(val *CreateStoreSpotlightCreativeRequestContent) *NullableCreateStoreSpotlightCreativeRequestContent {
	return &NullableCreateStoreSpotlightCreativeRequestContent{value: val, isSet: true}
}

func (v NullableCreateStoreSpotlightCreativeRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateStoreSpotlightCreativeRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
