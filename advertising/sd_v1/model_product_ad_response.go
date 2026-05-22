package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductAdResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAdResponse{}

// ProductAdResponse struct for ProductAdResponse
type ProductAdResponse struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Description *string `json:"description,omitempty"`
	// The identifier of the ad.
	AdId *float32 `json:"adId,omitempty"`
}

// NewProductAdResponse instantiates a new ProductAdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAdResponse() *ProductAdResponse {
	this := ProductAdResponse{}
	return &this
}

// NewProductAdResponseWithDefaults instantiates a new ProductAdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAdResponseWithDefaults() *ProductAdResponse {
	this := ProductAdResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProductAdResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAdResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProductAdResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProductAdResponse) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProductAdResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAdResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductAdResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProductAdResponse) SetDescription(v string) {
	o.Description = &v
}

// GetAdId returns the AdId field value if set, zero value otherwise.
func (o *ProductAdResponse) GetAdId() float32 {
	if o == nil || IsNil(o.AdId) {
		var ret float32
		return ret
	}
	return *o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAdResponse) GetAdIdOk() (*float32, bool) {
	if o == nil || IsNil(o.AdId) {
		return nil, false
	}
	return o.AdId, true
}

// HasAdId returns a boolean if a field has been set.
func (o *ProductAdResponse) HasAdId() bool {
	if o != nil && !IsNil(o.AdId) {
		return true
	}

	return false
}

// SetAdId gets a reference to the given float32 and assigns it to the AdId field.
func (o *ProductAdResponse) SetAdId(v float32) {
	o.AdId = &v
}

func (o ProductAdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AdId) {
		toSerialize["adId"] = o.AdId
	}
	return toSerialize, nil
}

type NullableProductAdResponse struct {
	value *ProductAdResponse
	isSet bool
}

func (v NullableProductAdResponse) Get() *ProductAdResponse {
	return v.value
}

func (v *NullableProductAdResponse) Set(val *ProductAdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAdResponse(val *ProductAdResponse) *NullableProductAdResponse {
	return &NullableProductAdResponse{value: val, isSet: true}
}

func (v NullableProductAdResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductAdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
