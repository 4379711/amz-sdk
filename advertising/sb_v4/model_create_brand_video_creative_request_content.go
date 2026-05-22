package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateBrandVideoCreativeRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBrandVideoCreativeRequestContent{}

// CreateBrandVideoCreativeRequestContent struct for CreateBrandVideoCreativeRequestContent
type CreateBrandVideoCreativeRequestContent struct {
	// The unique ID of a Sponsored Brands ad.
	AdId     string             `json:"adId"`
	Creative BrandVideoCreative `json:"creative"`
}

type _CreateBrandVideoCreativeRequestContent CreateBrandVideoCreativeRequestContent

// NewCreateBrandVideoCreativeRequestContent instantiates a new CreateBrandVideoCreativeRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBrandVideoCreativeRequestContent(adId string, creative BrandVideoCreative) *CreateBrandVideoCreativeRequestContent {
	this := CreateBrandVideoCreativeRequestContent{}
	this.AdId = adId
	this.Creative = creative
	return &this
}

// NewCreateBrandVideoCreativeRequestContentWithDefaults instantiates a new CreateBrandVideoCreativeRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBrandVideoCreativeRequestContentWithDefaults() *CreateBrandVideoCreativeRequestContent {
	this := CreateBrandVideoCreativeRequestContent{}
	return &this
}

// GetAdId returns the AdId field value
func (o *CreateBrandVideoCreativeRequestContent) GetAdId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value
// and a boolean to check if the value has been set.
func (o *CreateBrandVideoCreativeRequestContent) GetAdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdId, true
}

// SetAdId sets field value
func (o *CreateBrandVideoCreativeRequestContent) SetAdId(v string) {
	o.AdId = v
}

// GetCreative returns the Creative field value
func (o *CreateBrandVideoCreativeRequestContent) GetCreative() BrandVideoCreative {
	if o == nil {
		var ret BrandVideoCreative
		return ret
	}

	return o.Creative
}

// GetCreativeOk returns a tuple with the Creative field value
// and a boolean to check if the value has been set.
func (o *CreateBrandVideoCreativeRequestContent) GetCreativeOk() (*BrandVideoCreative, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creative, true
}

// SetCreative sets field value
func (o *CreateBrandVideoCreativeRequestContent) SetCreative(v BrandVideoCreative) {
	o.Creative = v
}

func (o CreateBrandVideoCreativeRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adId"] = o.AdId
	toSerialize["creative"] = o.Creative
	return toSerialize, nil
}

type NullableCreateBrandVideoCreativeRequestContent struct {
	value *CreateBrandVideoCreativeRequestContent
	isSet bool
}

func (v NullableCreateBrandVideoCreativeRequestContent) Get() *CreateBrandVideoCreativeRequestContent {
	return v.value
}

func (v *NullableCreateBrandVideoCreativeRequestContent) Set(val *CreateBrandVideoCreativeRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBrandVideoCreativeRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBrandVideoCreativeRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBrandVideoCreativeRequestContent(val *CreateBrandVideoCreativeRequestContent) *NullableCreateBrandVideoCreativeRequestContent {
	return &NullableCreateBrandVideoCreativeRequestContent{value: val, isSet: true}
}

func (v NullableCreateBrandVideoCreativeRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateBrandVideoCreativeRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
