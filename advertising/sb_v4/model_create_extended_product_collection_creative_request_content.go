package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateExtendedProductCollectionCreativeRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateExtendedProductCollectionCreativeRequestContent{}

// CreateExtendedProductCollectionCreativeRequestContent struct for CreateExtendedProductCollectionCreativeRequestContent
type CreateExtendedProductCollectionCreativeRequestContent struct {
	// The unique ID of a Sponsored Brands ad.
	AdId     string                            `json:"adId"`
	Creative ExtendedProductCollectionCreative `json:"creative"`
}

type _CreateExtendedProductCollectionCreativeRequestContent CreateExtendedProductCollectionCreativeRequestContent

// NewCreateExtendedProductCollectionCreativeRequestContent instantiates a new CreateExtendedProductCollectionCreativeRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateExtendedProductCollectionCreativeRequestContent(adId string, creative ExtendedProductCollectionCreative) *CreateExtendedProductCollectionCreativeRequestContent {
	this := CreateExtendedProductCollectionCreativeRequestContent{}
	this.AdId = adId
	this.Creative = creative
	return &this
}

// NewCreateExtendedProductCollectionCreativeRequestContentWithDefaults instantiates a new CreateExtendedProductCollectionCreativeRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateExtendedProductCollectionCreativeRequestContentWithDefaults() *CreateExtendedProductCollectionCreativeRequestContent {
	this := CreateExtendedProductCollectionCreativeRequestContent{}
	return &this
}

// GetAdId returns the AdId field value
func (o *CreateExtendedProductCollectionCreativeRequestContent) GetAdId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value
// and a boolean to check if the value has been set.
func (o *CreateExtendedProductCollectionCreativeRequestContent) GetAdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdId, true
}

// SetAdId sets field value
func (o *CreateExtendedProductCollectionCreativeRequestContent) SetAdId(v string) {
	o.AdId = v
}

// GetCreative returns the Creative field value
func (o *CreateExtendedProductCollectionCreativeRequestContent) GetCreative() ExtendedProductCollectionCreative {
	if o == nil {
		var ret ExtendedProductCollectionCreative
		return ret
	}

	return o.Creative
}

// GetCreativeOk returns a tuple with the Creative field value
// and a boolean to check if the value has been set.
func (o *CreateExtendedProductCollectionCreativeRequestContent) GetCreativeOk() (*ExtendedProductCollectionCreative, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creative, true
}

// SetCreative sets field value
func (o *CreateExtendedProductCollectionCreativeRequestContent) SetCreative(v ExtendedProductCollectionCreative) {
	o.Creative = v
}

func (o CreateExtendedProductCollectionCreativeRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adId"] = o.AdId
	toSerialize["creative"] = o.Creative
	return toSerialize, nil
}

type NullableCreateExtendedProductCollectionCreativeRequestContent struct {
	value *CreateExtendedProductCollectionCreativeRequestContent
	isSet bool
}

func (v NullableCreateExtendedProductCollectionCreativeRequestContent) Get() *CreateExtendedProductCollectionCreativeRequestContent {
	return v.value
}

func (v *NullableCreateExtendedProductCollectionCreativeRequestContent) Set(val *CreateExtendedProductCollectionCreativeRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateExtendedProductCollectionCreativeRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateExtendedProductCollectionCreativeRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateExtendedProductCollectionCreativeRequestContent(val *CreateExtendedProductCollectionCreativeRequestContent) *NullableCreateExtendedProductCollectionCreativeRequestContent {
	return &NullableCreateExtendedProductCollectionCreativeRequestContent{value: val, isSet: true}
}

func (v NullableCreateExtendedProductCollectionCreativeRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateExtendedProductCollectionCreativeRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
