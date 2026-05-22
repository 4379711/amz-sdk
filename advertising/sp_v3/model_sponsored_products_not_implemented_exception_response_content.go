package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNotImplementedExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNotImplementedExceptionResponseContent{}

// SponsoredProductsNotImplementedExceptionResponseContent Operation is not implemented.
type SponsoredProductsNotImplementedExceptionResponseContent struct {
	Code *SponsoredProductsNotImplementedExceptionCode `json:"code,omitempty"`
	// A human-readable description of the response.
	Message *string `json:"message,omitempty"`
}

// NewSponsoredProductsNotImplementedExceptionResponseContent instantiates a new SponsoredProductsNotImplementedExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNotImplementedExceptionResponseContent() *SponsoredProductsNotImplementedExceptionResponseContent {
	this := SponsoredProductsNotImplementedExceptionResponseContent{}
	return &this
}

// NewSponsoredProductsNotImplementedExceptionResponseContentWithDefaults instantiates a new SponsoredProductsNotImplementedExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNotImplementedExceptionResponseContentWithDefaults() *SponsoredProductsNotImplementedExceptionResponseContent {
	this := SponsoredProductsNotImplementedExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SponsoredProductsNotImplementedExceptionResponseContent) GetCode() SponsoredProductsNotImplementedExceptionCode {
	if o == nil || IsNil(o.Code) {
		var ret SponsoredProductsNotImplementedExceptionCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNotImplementedExceptionResponseContent) GetCodeOk() (*SponsoredProductsNotImplementedExceptionCode, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SponsoredProductsNotImplementedExceptionResponseContent) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given SponsoredProductsNotImplementedExceptionCode and assigns it to the Code field.
func (o *SponsoredProductsNotImplementedExceptionResponseContent) SetCode(v SponsoredProductsNotImplementedExceptionCode) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SponsoredProductsNotImplementedExceptionResponseContent) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNotImplementedExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SponsoredProductsNotImplementedExceptionResponseContent) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SponsoredProductsNotImplementedExceptionResponseContent) SetMessage(v string) {
	o.Message = &v
}

func (o SponsoredProductsNotImplementedExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableSponsoredProductsNotImplementedExceptionResponseContent struct {
	value *SponsoredProductsNotImplementedExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsNotImplementedExceptionResponseContent) Get() *SponsoredProductsNotImplementedExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsNotImplementedExceptionResponseContent) Set(val *SponsoredProductsNotImplementedExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNotImplementedExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNotImplementedExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNotImplementedExceptionResponseContent(val *SponsoredProductsNotImplementedExceptionResponseContent) *NullableSponsoredProductsNotImplementedExceptionResponseContent {
	return &NullableSponsoredProductsNotImplementedExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsNotImplementedExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNotImplementedExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
