package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsSchemaValidationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsSchemaValidationExceptionResponseContent{}

// SponsoredProductsSchemaValidationExceptionResponseContent struct for SponsoredProductsSchemaValidationExceptionResponseContent
type SponsoredProductsSchemaValidationExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsSchemaValidationExceptionResponseContent SponsoredProductsSchemaValidationExceptionResponseContent

// NewSponsoredProductsSchemaValidationExceptionResponseContent instantiates a new SponsoredProductsSchemaValidationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsSchemaValidationExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsSchemaValidationExceptionResponseContent {
	this := SponsoredProductsSchemaValidationExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsSchemaValidationExceptionResponseContentWithDefaults instantiates a new SponsoredProductsSchemaValidationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsSchemaValidationExceptionResponseContentWithDefaults() *SponsoredProductsSchemaValidationExceptionResponseContent {
	this := SponsoredProductsSchemaValidationExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsSchemaValidationExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSchemaValidationExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsSchemaValidationExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsSchemaValidationExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSchemaValidationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsSchemaValidationExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsSchemaValidationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsSchemaValidationExceptionResponseContent struct {
	value *SponsoredProductsSchemaValidationExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsSchemaValidationExceptionResponseContent) Get() *SponsoredProductsSchemaValidationExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsSchemaValidationExceptionResponseContent) Set(val *SponsoredProductsSchemaValidationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsSchemaValidationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsSchemaValidationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsSchemaValidationExceptionResponseContent(val *SponsoredProductsSchemaValidationExceptionResponseContent) *NullableSponsoredProductsSchemaValidationExceptionResponseContent {
	return &NullableSponsoredProductsSchemaValidationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsSchemaValidationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsSchemaValidationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
