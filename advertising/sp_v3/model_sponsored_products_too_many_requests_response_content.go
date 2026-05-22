package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsTooManyRequestsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsTooManyRequestsResponseContent{}

// SponsoredProductsTooManyRequestsResponseContent struct for SponsoredProductsTooManyRequestsResponseContent
type SponsoredProductsTooManyRequestsResponseContent struct {
	Code           SponsoredProductsErrorCode `json:"code"`
	Message        string                     `json:"message"`
	HttpStatusCode *string                    `json:"httpStatusCode,omitempty"`
}

type _SponsoredProductsTooManyRequestsResponseContent SponsoredProductsTooManyRequestsResponseContent

// NewSponsoredProductsTooManyRequestsResponseContent instantiates a new SponsoredProductsTooManyRequestsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsTooManyRequestsResponseContent(code SponsoredProductsErrorCode, message string) *SponsoredProductsTooManyRequestsResponseContent {
	this := SponsoredProductsTooManyRequestsResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsTooManyRequestsResponseContentWithDefaults instantiates a new SponsoredProductsTooManyRequestsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsTooManyRequestsResponseContentWithDefaults() *SponsoredProductsTooManyRequestsResponseContent {
	this := SponsoredProductsTooManyRequestsResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsTooManyRequestsResponseContent) GetCode() SponsoredProductsErrorCode {
	if o == nil {
		var ret SponsoredProductsErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTooManyRequestsResponseContent) GetCodeOk() (*SponsoredProductsErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsTooManyRequestsResponseContent) SetCode(v SponsoredProductsErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsTooManyRequestsResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTooManyRequestsResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsTooManyRequestsResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetHttpStatusCode returns the HttpStatusCode field value if set, zero value otherwise.
func (o *SponsoredProductsTooManyRequestsResponseContent) GetHttpStatusCode() string {
	if o == nil || IsNil(o.HttpStatusCode) {
		var ret string
		return ret
	}
	return *o.HttpStatusCode
}

// GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTooManyRequestsResponseContent) GetHttpStatusCodeOk() (*string, bool) {
	if o == nil || IsNil(o.HttpStatusCode) {
		return nil, false
	}
	return o.HttpStatusCode, true
}

// HasHttpStatusCode returns a boolean if a field has been set.
func (o *SponsoredProductsTooManyRequestsResponseContent) HasHttpStatusCode() bool {
	if o != nil && !IsNil(o.HttpStatusCode) {
		return true
	}

	return false
}

// SetHttpStatusCode gets a reference to the given string and assigns it to the HttpStatusCode field.
func (o *SponsoredProductsTooManyRequestsResponseContent) SetHttpStatusCode(v string) {
	o.HttpStatusCode = &v
}

func (o SponsoredProductsTooManyRequestsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.HttpStatusCode) {
		toSerialize["httpStatusCode"] = o.HttpStatusCode
	}
	return toSerialize, nil
}

type NullableSponsoredProductsTooManyRequestsResponseContent struct {
	value *SponsoredProductsTooManyRequestsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsTooManyRequestsResponseContent) Get() *SponsoredProductsTooManyRequestsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsTooManyRequestsResponseContent) Set(val *SponsoredProductsTooManyRequestsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTooManyRequestsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTooManyRequestsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTooManyRequestsResponseContent(val *SponsoredProductsTooManyRequestsResponseContent) *NullableSponsoredProductsTooManyRequestsResponseContent {
	return &NullableSponsoredProductsTooManyRequestsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsTooManyRequestsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTooManyRequestsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
