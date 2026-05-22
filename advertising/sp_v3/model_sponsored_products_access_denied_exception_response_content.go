package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsAccessDeniedExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsAccessDeniedExceptionResponseContent{}

// SponsoredProductsAccessDeniedExceptionResponseContent struct for SponsoredProductsAccessDeniedExceptionResponseContent
type SponsoredProductsAccessDeniedExceptionResponseContent struct {
	Code SponsoredProductsAccessDeniedErrorCode `json:"code"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsAccessDeniedExceptionResponseContent SponsoredProductsAccessDeniedExceptionResponseContent

// NewSponsoredProductsAccessDeniedExceptionResponseContent instantiates a new SponsoredProductsAccessDeniedExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsAccessDeniedExceptionResponseContent(code SponsoredProductsAccessDeniedErrorCode, message string) *SponsoredProductsAccessDeniedExceptionResponseContent {
	this := SponsoredProductsAccessDeniedExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsAccessDeniedExceptionResponseContentWithDefaults instantiates a new SponsoredProductsAccessDeniedExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsAccessDeniedExceptionResponseContentWithDefaults() *SponsoredProductsAccessDeniedExceptionResponseContent {
	this := SponsoredProductsAccessDeniedExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsAccessDeniedExceptionResponseContent) GetCode() SponsoredProductsAccessDeniedErrorCode {
	if o == nil {
		var ret SponsoredProductsAccessDeniedErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAccessDeniedExceptionResponseContent) GetCodeOk() (*SponsoredProductsAccessDeniedErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsAccessDeniedExceptionResponseContent) SetCode(v SponsoredProductsAccessDeniedErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsAccessDeniedExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAccessDeniedExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsAccessDeniedExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsAccessDeniedExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsAccessDeniedExceptionResponseContent struct {
	value *SponsoredProductsAccessDeniedExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsAccessDeniedExceptionResponseContent) Get() *SponsoredProductsAccessDeniedExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsAccessDeniedExceptionResponseContent) Set(val *SponsoredProductsAccessDeniedExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAccessDeniedExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAccessDeniedExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAccessDeniedExceptionResponseContent(val *SponsoredProductsAccessDeniedExceptionResponseContent) *NullableSponsoredProductsAccessDeniedExceptionResponseContent {
	return &NullableSponsoredProductsAccessDeniedExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsAccessDeniedExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAccessDeniedExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
