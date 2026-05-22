package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsThrottlingExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsThrottlingExceptionResponseContent{}

// SponsoredProductsThrottlingExceptionResponseContent struct for SponsoredProductsThrottlingExceptionResponseContent
type SponsoredProductsThrottlingExceptionResponseContent struct {
	Code SponsoredProductsThrottledErrorCode `json:"code"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsThrottlingExceptionResponseContent SponsoredProductsThrottlingExceptionResponseContent

// NewSponsoredProductsThrottlingExceptionResponseContent instantiates a new SponsoredProductsThrottlingExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsThrottlingExceptionResponseContent(code SponsoredProductsThrottledErrorCode, message string) *SponsoredProductsThrottlingExceptionResponseContent {
	this := SponsoredProductsThrottlingExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsThrottlingExceptionResponseContentWithDefaults instantiates a new SponsoredProductsThrottlingExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsThrottlingExceptionResponseContentWithDefaults() *SponsoredProductsThrottlingExceptionResponseContent {
	this := SponsoredProductsThrottlingExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsThrottlingExceptionResponseContent) GetCode() SponsoredProductsThrottledErrorCode {
	if o == nil {
		var ret SponsoredProductsThrottledErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsThrottlingExceptionResponseContent) GetCodeOk() (*SponsoredProductsThrottledErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsThrottlingExceptionResponseContent) SetCode(v SponsoredProductsThrottledErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsThrottlingExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsThrottlingExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsThrottlingExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsThrottlingExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsThrottlingExceptionResponseContent struct {
	value *SponsoredProductsThrottlingExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsThrottlingExceptionResponseContent) Get() *SponsoredProductsThrottlingExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsThrottlingExceptionResponseContent) Set(val *SponsoredProductsThrottlingExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsThrottlingExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsThrottlingExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsThrottlingExceptionResponseContent(val *SponsoredProductsThrottlingExceptionResponseContent) *NullableSponsoredProductsThrottlingExceptionResponseContent {
	return &NullableSponsoredProductsThrottlingExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsThrottlingExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsThrottlingExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
