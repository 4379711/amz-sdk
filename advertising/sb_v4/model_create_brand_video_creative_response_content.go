package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateBrandVideoCreativeResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBrandVideoCreativeResponseContent{}

// CreateBrandVideoCreativeResponseContent Create creative response
type CreateBrandVideoCreativeResponseContent struct {
	// The unique ID of a Sponsored Brands ad.
	AdId *string `json:"adId,omitempty"`
	// The version identifier that helps you keep track of multiple versions of a submitted (non-draft) Sponsored Brands creative.
	CreativeVersion *string `json:"creativeVersion,omitempty"`
}

// NewCreateBrandVideoCreativeResponseContent instantiates a new CreateBrandVideoCreativeResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBrandVideoCreativeResponseContent() *CreateBrandVideoCreativeResponseContent {
	this := CreateBrandVideoCreativeResponseContent{}
	return &this
}

// NewCreateBrandVideoCreativeResponseContentWithDefaults instantiates a new CreateBrandVideoCreativeResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBrandVideoCreativeResponseContentWithDefaults() *CreateBrandVideoCreativeResponseContent {
	this := CreateBrandVideoCreativeResponseContent{}
	return &this
}

// GetAdId returns the AdId field value if set, zero value otherwise.
func (o *CreateBrandVideoCreativeResponseContent) GetAdId() string {
	if o == nil || IsNil(o.AdId) {
		var ret string
		return ret
	}
	return *o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBrandVideoCreativeResponseContent) GetAdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdId) {
		return nil, false
	}
	return o.AdId, true
}

// HasAdId returns a boolean if a field has been set.
func (o *CreateBrandVideoCreativeResponseContent) HasAdId() bool {
	if o != nil && !IsNil(o.AdId) {
		return true
	}

	return false
}

// SetAdId gets a reference to the given string and assigns it to the AdId field.
func (o *CreateBrandVideoCreativeResponseContent) SetAdId(v string) {
	o.AdId = &v
}

// GetCreativeVersion returns the CreativeVersion field value if set, zero value otherwise.
func (o *CreateBrandVideoCreativeResponseContent) GetCreativeVersion() string {
	if o == nil || IsNil(o.CreativeVersion) {
		var ret string
		return ret
	}
	return *o.CreativeVersion
}

// GetCreativeVersionOk returns a tuple with the CreativeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBrandVideoCreativeResponseContent) GetCreativeVersionOk() (*string, bool) {
	if o == nil || IsNil(o.CreativeVersion) {
		return nil, false
	}
	return o.CreativeVersion, true
}

// HasCreativeVersion returns a boolean if a field has been set.
func (o *CreateBrandVideoCreativeResponseContent) HasCreativeVersion() bool {
	if o != nil && !IsNil(o.CreativeVersion) {
		return true
	}

	return false
}

// SetCreativeVersion gets a reference to the given string and assigns it to the CreativeVersion field.
func (o *CreateBrandVideoCreativeResponseContent) SetCreativeVersion(v string) {
	o.CreativeVersion = &v
}

func (o CreateBrandVideoCreativeResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdId) {
		toSerialize["adId"] = o.AdId
	}
	if !IsNil(o.CreativeVersion) {
		toSerialize["creativeVersion"] = o.CreativeVersion
	}
	return toSerialize, nil
}

type NullableCreateBrandVideoCreativeResponseContent struct {
	value *CreateBrandVideoCreativeResponseContent
	isSet bool
}

func (v NullableCreateBrandVideoCreativeResponseContent) Get() *CreateBrandVideoCreativeResponseContent {
	return v.value
}

func (v *NullableCreateBrandVideoCreativeResponseContent) Set(val *CreateBrandVideoCreativeResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBrandVideoCreativeResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBrandVideoCreativeResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBrandVideoCreativeResponseContent(val *CreateBrandVideoCreativeResponseContent) *NullableCreateBrandVideoCreativeResponseContent {
	return &NullableCreateBrandVideoCreativeResponseContent{value: val, isSet: true}
}

func (v NullableCreateBrandVideoCreativeResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateBrandVideoCreativeResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
