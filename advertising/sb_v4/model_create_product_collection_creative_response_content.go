package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateProductCollectionCreativeResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProductCollectionCreativeResponseContent{}

// CreateProductCollectionCreativeResponseContent Create creative response
type CreateProductCollectionCreativeResponseContent struct {
	// The unique ID of a Sponsored Brands ad.
	AdId *string `json:"adId,omitempty"`
	// The version identifier that helps you keep track of multiple versions of a submitted (non-draft) Sponsored Brands creative.
	CreativeVersion *string `json:"creativeVersion,omitempty"`
}

// NewCreateProductCollectionCreativeResponseContent instantiates a new CreateProductCollectionCreativeResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProductCollectionCreativeResponseContent() *CreateProductCollectionCreativeResponseContent {
	this := CreateProductCollectionCreativeResponseContent{}
	return &this
}

// NewCreateProductCollectionCreativeResponseContentWithDefaults instantiates a new CreateProductCollectionCreativeResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProductCollectionCreativeResponseContentWithDefaults() *CreateProductCollectionCreativeResponseContent {
	this := CreateProductCollectionCreativeResponseContent{}
	return &this
}

// GetAdId returns the AdId field value if set, zero value otherwise.
func (o *CreateProductCollectionCreativeResponseContent) GetAdId() string {
	if o == nil || IsNil(o.AdId) {
		var ret string
		return ret
	}
	return *o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductCollectionCreativeResponseContent) GetAdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdId) {
		return nil, false
	}
	return o.AdId, true
}

// HasAdId returns a boolean if a field has been set.
func (o *CreateProductCollectionCreativeResponseContent) HasAdId() bool {
	if o != nil && !IsNil(o.AdId) {
		return true
	}

	return false
}

// SetAdId gets a reference to the given string and assigns it to the AdId field.
func (o *CreateProductCollectionCreativeResponseContent) SetAdId(v string) {
	o.AdId = &v
}

// GetCreativeVersion returns the CreativeVersion field value if set, zero value otherwise.
func (o *CreateProductCollectionCreativeResponseContent) GetCreativeVersion() string {
	if o == nil || IsNil(o.CreativeVersion) {
		var ret string
		return ret
	}
	return *o.CreativeVersion
}

// GetCreativeVersionOk returns a tuple with the CreativeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductCollectionCreativeResponseContent) GetCreativeVersionOk() (*string, bool) {
	if o == nil || IsNil(o.CreativeVersion) {
		return nil, false
	}
	return o.CreativeVersion, true
}

// HasCreativeVersion returns a boolean if a field has been set.
func (o *CreateProductCollectionCreativeResponseContent) HasCreativeVersion() bool {
	if o != nil && !IsNil(o.CreativeVersion) {
		return true
	}

	return false
}

// SetCreativeVersion gets a reference to the given string and assigns it to the CreativeVersion field.
func (o *CreateProductCollectionCreativeResponseContent) SetCreativeVersion(v string) {
	o.CreativeVersion = &v
}

func (o CreateProductCollectionCreativeResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdId) {
		toSerialize["adId"] = o.AdId
	}
	if !IsNil(o.CreativeVersion) {
		toSerialize["creativeVersion"] = o.CreativeVersion
	}
	return toSerialize, nil
}

type NullableCreateProductCollectionCreativeResponseContent struct {
	value *CreateProductCollectionCreativeResponseContent
	isSet bool
}

func (v NullableCreateProductCollectionCreativeResponseContent) Get() *CreateProductCollectionCreativeResponseContent {
	return v.value
}

func (v *NullableCreateProductCollectionCreativeResponseContent) Set(val *CreateProductCollectionCreativeResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProductCollectionCreativeResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProductCollectionCreativeResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProductCollectionCreativeResponseContent(val *CreateProductCollectionCreativeResponseContent) *NullableCreateProductCollectionCreativeResponseContent {
	return &NullableCreateProductCollectionCreativeResponseContent{value: val, isSet: true}
}

func (v NullableCreateProductCollectionCreativeResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateProductCollectionCreativeResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
