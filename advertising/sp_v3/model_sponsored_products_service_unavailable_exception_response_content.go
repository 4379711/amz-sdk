package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsServiceUnavailableExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsServiceUnavailableExceptionResponseContent{}

// SponsoredProductsServiceUnavailableExceptionResponseContent Server unable to process request. Please retry later.
type SponsoredProductsServiceUnavailableExceptionResponseContent struct {
	Code    SponsoredProductsServiceUnavailableExceptionErrorCode `json:"code"`
	Message string                                                `json:"message"`
}

type _SponsoredProductsServiceUnavailableExceptionResponseContent SponsoredProductsServiceUnavailableExceptionResponseContent

// NewSponsoredProductsServiceUnavailableExceptionResponseContent instantiates a new SponsoredProductsServiceUnavailableExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsServiceUnavailableExceptionResponseContent(code SponsoredProductsServiceUnavailableExceptionErrorCode, message string) *SponsoredProductsServiceUnavailableExceptionResponseContent {
	this := SponsoredProductsServiceUnavailableExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsServiceUnavailableExceptionResponseContentWithDefaults instantiates a new SponsoredProductsServiceUnavailableExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsServiceUnavailableExceptionResponseContentWithDefaults() *SponsoredProductsServiceUnavailableExceptionResponseContent {
	this := SponsoredProductsServiceUnavailableExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsServiceUnavailableExceptionResponseContent) GetCode() SponsoredProductsServiceUnavailableExceptionErrorCode {
	if o == nil {
		var ret SponsoredProductsServiceUnavailableExceptionErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsServiceUnavailableExceptionResponseContent) GetCodeOk() (*SponsoredProductsServiceUnavailableExceptionErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsServiceUnavailableExceptionResponseContent) SetCode(v SponsoredProductsServiceUnavailableExceptionErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsServiceUnavailableExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsServiceUnavailableExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsServiceUnavailableExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsServiceUnavailableExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsServiceUnavailableExceptionResponseContent struct {
	value *SponsoredProductsServiceUnavailableExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsServiceUnavailableExceptionResponseContent) Get() *SponsoredProductsServiceUnavailableExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsServiceUnavailableExceptionResponseContent) Set(val *SponsoredProductsServiceUnavailableExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsServiceUnavailableExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsServiceUnavailableExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsServiceUnavailableExceptionResponseContent(val *SponsoredProductsServiceUnavailableExceptionResponseContent) *NullableSponsoredProductsServiceUnavailableExceptionResponseContent {
	return &NullableSponsoredProductsServiceUnavailableExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsServiceUnavailableExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsServiceUnavailableExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
