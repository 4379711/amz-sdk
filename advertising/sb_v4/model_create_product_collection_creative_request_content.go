package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateProductCollectionCreativeRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProductCollectionCreativeRequestContent{}

// CreateProductCollectionCreativeRequestContent struct for CreateProductCollectionCreativeRequestContent
type CreateProductCollectionCreativeRequestContent struct {
	// The unique ID of a Sponsored Brands ad.
	AdId     string                    `json:"adId"`
	Creative ProductCollectionCreative `json:"creative"`
}

type _CreateProductCollectionCreativeRequestContent CreateProductCollectionCreativeRequestContent

// NewCreateProductCollectionCreativeRequestContent instantiates a new CreateProductCollectionCreativeRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProductCollectionCreativeRequestContent(adId string, creative ProductCollectionCreative) *CreateProductCollectionCreativeRequestContent {
	this := CreateProductCollectionCreativeRequestContent{}
	this.AdId = adId
	this.Creative = creative
	return &this
}

// NewCreateProductCollectionCreativeRequestContentWithDefaults instantiates a new CreateProductCollectionCreativeRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProductCollectionCreativeRequestContentWithDefaults() *CreateProductCollectionCreativeRequestContent {
	this := CreateProductCollectionCreativeRequestContent{}
	return &this
}

// GetAdId returns the AdId field value
func (o *CreateProductCollectionCreativeRequestContent) GetAdId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value
// and a boolean to check if the value has been set.
func (o *CreateProductCollectionCreativeRequestContent) GetAdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdId, true
}

// SetAdId sets field value
func (o *CreateProductCollectionCreativeRequestContent) SetAdId(v string) {
	o.AdId = v
}

// GetCreative returns the Creative field value
func (o *CreateProductCollectionCreativeRequestContent) GetCreative() ProductCollectionCreative {
	if o == nil {
		var ret ProductCollectionCreative
		return ret
	}

	return o.Creative
}

// GetCreativeOk returns a tuple with the Creative field value
// and a boolean to check if the value has been set.
func (o *CreateProductCollectionCreativeRequestContent) GetCreativeOk() (*ProductCollectionCreative, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creative, true
}

// SetCreative sets field value
func (o *CreateProductCollectionCreativeRequestContent) SetCreative(v ProductCollectionCreative) {
	o.Creative = v
}

func (o CreateProductCollectionCreativeRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adId"] = o.AdId
	toSerialize["creative"] = o.Creative
	return toSerialize, nil
}

type NullableCreateProductCollectionCreativeRequestContent struct {
	value *CreateProductCollectionCreativeRequestContent
	isSet bool
}

func (v NullableCreateProductCollectionCreativeRequestContent) Get() *CreateProductCollectionCreativeRequestContent {
	return v.value
}

func (v *NullableCreateProductCollectionCreativeRequestContent) Set(val *CreateProductCollectionCreativeRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProductCollectionCreativeRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProductCollectionCreativeRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProductCollectionCreativeRequestContent(val *CreateProductCollectionCreativeRequestContent) *NullableCreateProductCollectionCreativeRequestContent {
	return &NullableCreateProductCollectionCreativeRequestContent{value: val, isSet: true}
}

func (v NullableCreateProductCollectionCreativeRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateProductCollectionCreativeRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
