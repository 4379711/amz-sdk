package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsInternalServerExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsInternalServerExceptionResponseContent{}

// SponsoredProductsInternalServerExceptionResponseContent struct for SponsoredProductsInternalServerExceptionResponseContent
type SponsoredProductsInternalServerExceptionResponseContent struct {
	Code SponsoredProductsInternalErrorErrorCode `json:"code"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsInternalServerExceptionResponseContent SponsoredProductsInternalServerExceptionResponseContent

// NewSponsoredProductsInternalServerExceptionResponseContent instantiates a new SponsoredProductsInternalServerExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsInternalServerExceptionResponseContent(code SponsoredProductsInternalErrorErrorCode, message string) *SponsoredProductsInternalServerExceptionResponseContent {
	this := SponsoredProductsInternalServerExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsInternalServerExceptionResponseContentWithDefaults instantiates a new SponsoredProductsInternalServerExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsInternalServerExceptionResponseContentWithDefaults() *SponsoredProductsInternalServerExceptionResponseContent {
	this := SponsoredProductsInternalServerExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsInternalServerExceptionResponseContent) GetCode() SponsoredProductsInternalErrorErrorCode {
	if o == nil {
		var ret SponsoredProductsInternalErrorErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsInternalServerExceptionResponseContent) GetCodeOk() (*SponsoredProductsInternalErrorErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsInternalServerExceptionResponseContent) SetCode(v SponsoredProductsInternalErrorErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsInternalServerExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsInternalServerExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsInternalServerExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsInternalServerExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsInternalServerExceptionResponseContent struct {
	value *SponsoredProductsInternalServerExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsInternalServerExceptionResponseContent) Get() *SponsoredProductsInternalServerExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsInternalServerExceptionResponseContent) Set(val *SponsoredProductsInternalServerExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsInternalServerExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsInternalServerExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsInternalServerExceptionResponseContent(val *SponsoredProductsInternalServerExceptionResponseContent) *NullableSponsoredProductsInternalServerExceptionResponseContent {
	return &NullableSponsoredProductsInternalServerExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsInternalServerExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsInternalServerExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
