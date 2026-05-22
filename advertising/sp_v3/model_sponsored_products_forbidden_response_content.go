package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsForbiddenResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsForbiddenResponseContent{}

// SponsoredProductsForbiddenResponseContent struct for SponsoredProductsForbiddenResponseContent
type SponsoredProductsForbiddenResponseContent struct {
	Code           SponsoredProductsErrorCode `json:"code"`
	Message        string                     `json:"message"`
	HttpStatusCode *string                    `json:"httpStatusCode,omitempty"`
}

type _SponsoredProductsForbiddenResponseContent SponsoredProductsForbiddenResponseContent

// NewSponsoredProductsForbiddenResponseContent instantiates a new SponsoredProductsForbiddenResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsForbiddenResponseContent(code SponsoredProductsErrorCode, message string) *SponsoredProductsForbiddenResponseContent {
	this := SponsoredProductsForbiddenResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsForbiddenResponseContentWithDefaults instantiates a new SponsoredProductsForbiddenResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsForbiddenResponseContentWithDefaults() *SponsoredProductsForbiddenResponseContent {
	this := SponsoredProductsForbiddenResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsForbiddenResponseContent) GetCode() SponsoredProductsErrorCode {
	if o == nil {
		var ret SponsoredProductsErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsForbiddenResponseContent) GetCodeOk() (*SponsoredProductsErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsForbiddenResponseContent) SetCode(v SponsoredProductsErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsForbiddenResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsForbiddenResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsForbiddenResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetHttpStatusCode returns the HttpStatusCode field value if set, zero value otherwise.
func (o *SponsoredProductsForbiddenResponseContent) GetHttpStatusCode() string {
	if o == nil || IsNil(o.HttpStatusCode) {
		var ret string
		return ret
	}
	return *o.HttpStatusCode
}

// GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsForbiddenResponseContent) GetHttpStatusCodeOk() (*string, bool) {
	if o == nil || IsNil(o.HttpStatusCode) {
		return nil, false
	}
	return o.HttpStatusCode, true
}

// HasHttpStatusCode returns a boolean if a field has been set.
func (o *SponsoredProductsForbiddenResponseContent) HasHttpStatusCode() bool {
	if o != nil && !IsNil(o.HttpStatusCode) {
		return true
	}

	return false
}

// SetHttpStatusCode gets a reference to the given string and assigns it to the HttpStatusCode field.
func (o *SponsoredProductsForbiddenResponseContent) SetHttpStatusCode(v string) {
	o.HttpStatusCode = &v
}

func (o SponsoredProductsForbiddenResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.HttpStatusCode) {
		toSerialize["httpStatusCode"] = o.HttpStatusCode
	}
	return toSerialize, nil
}

type NullableSponsoredProductsForbiddenResponseContent struct {
	value *SponsoredProductsForbiddenResponseContent
	isSet bool
}

func (v NullableSponsoredProductsForbiddenResponseContent) Get() *SponsoredProductsForbiddenResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsForbiddenResponseContent) Set(val *SponsoredProductsForbiddenResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsForbiddenResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsForbiddenResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsForbiddenResponseContent(val *SponsoredProductsForbiddenResponseContent) *NullableSponsoredProductsForbiddenResponseContent {
	return &NullableSponsoredProductsForbiddenResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsForbiddenResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsForbiddenResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
