package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUnauthorizedExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUnauthorizedExceptionResponseContent{}

// SponsoredProductsUnauthorizedExceptionResponseContent struct for SponsoredProductsUnauthorizedExceptionResponseContent
type SponsoredProductsUnauthorizedExceptionResponseContent struct {
	Code SponsoredProductsUnauthorizedErrorCode `json:"code"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsUnauthorizedExceptionResponseContent SponsoredProductsUnauthorizedExceptionResponseContent

// NewSponsoredProductsUnauthorizedExceptionResponseContent instantiates a new SponsoredProductsUnauthorizedExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUnauthorizedExceptionResponseContent(code SponsoredProductsUnauthorizedErrorCode, message string) *SponsoredProductsUnauthorizedExceptionResponseContent {
	this := SponsoredProductsUnauthorizedExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsUnauthorizedExceptionResponseContentWithDefaults instantiates a new SponsoredProductsUnauthorizedExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUnauthorizedExceptionResponseContentWithDefaults() *SponsoredProductsUnauthorizedExceptionResponseContent {
	this := SponsoredProductsUnauthorizedExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsUnauthorizedExceptionResponseContent) GetCode() SponsoredProductsUnauthorizedErrorCode {
	if o == nil {
		var ret SponsoredProductsUnauthorizedErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUnauthorizedExceptionResponseContent) GetCodeOk() (*SponsoredProductsUnauthorizedErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsUnauthorizedExceptionResponseContent) SetCode(v SponsoredProductsUnauthorizedErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsUnauthorizedExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUnauthorizedExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsUnauthorizedExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsUnauthorizedExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsUnauthorizedExceptionResponseContent struct {
	value *SponsoredProductsUnauthorizedExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUnauthorizedExceptionResponseContent) Get() *SponsoredProductsUnauthorizedExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUnauthorizedExceptionResponseContent) Set(val *SponsoredProductsUnauthorizedExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUnauthorizedExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUnauthorizedExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUnauthorizedExceptionResponseContent(val *SponsoredProductsUnauthorizedExceptionResponseContent) *NullableSponsoredProductsUnauthorizedExceptionResponseContent {
	return &NullableSponsoredProductsUnauthorizedExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUnauthorizedExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUnauthorizedExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
