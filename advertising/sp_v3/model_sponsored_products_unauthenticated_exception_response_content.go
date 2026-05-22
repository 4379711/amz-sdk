package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUnauthenticatedExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUnauthenticatedExceptionResponseContent{}

// SponsoredProductsUnauthenticatedExceptionResponseContent Unauthenticated. Request failed because user is not authenticated.
type SponsoredProductsUnauthenticatedExceptionResponseContent struct {
	Code *SponsoredProductsUnauthenticatedExceptionCode `json:"code,omitempty"`
	// A human-readable description of the response.
	Message *string `json:"message,omitempty"`
}

// NewSponsoredProductsUnauthenticatedExceptionResponseContent instantiates a new SponsoredProductsUnauthenticatedExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUnauthenticatedExceptionResponseContent() *SponsoredProductsUnauthenticatedExceptionResponseContent {
	this := SponsoredProductsUnauthenticatedExceptionResponseContent{}
	return &this
}

// NewSponsoredProductsUnauthenticatedExceptionResponseContentWithDefaults instantiates a new SponsoredProductsUnauthenticatedExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUnauthenticatedExceptionResponseContentWithDefaults() *SponsoredProductsUnauthenticatedExceptionResponseContent {
	this := SponsoredProductsUnauthenticatedExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SponsoredProductsUnauthenticatedExceptionResponseContent) GetCode() SponsoredProductsUnauthenticatedExceptionCode {
	if o == nil || IsNil(o.Code) {
		var ret SponsoredProductsUnauthenticatedExceptionCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUnauthenticatedExceptionResponseContent) GetCodeOk() (*SponsoredProductsUnauthenticatedExceptionCode, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SponsoredProductsUnauthenticatedExceptionResponseContent) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given SponsoredProductsUnauthenticatedExceptionCode and assigns it to the Code field.
func (o *SponsoredProductsUnauthenticatedExceptionResponseContent) SetCode(v SponsoredProductsUnauthenticatedExceptionCode) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SponsoredProductsUnauthenticatedExceptionResponseContent) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUnauthenticatedExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SponsoredProductsUnauthenticatedExceptionResponseContent) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SponsoredProductsUnauthenticatedExceptionResponseContent) SetMessage(v string) {
	o.Message = &v
}

func (o SponsoredProductsUnauthenticatedExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableSponsoredProductsUnauthenticatedExceptionResponseContent struct {
	value *SponsoredProductsUnauthenticatedExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUnauthenticatedExceptionResponseContent) Get() *SponsoredProductsUnauthenticatedExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUnauthenticatedExceptionResponseContent) Set(val *SponsoredProductsUnauthenticatedExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUnauthenticatedExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUnauthenticatedExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUnauthenticatedExceptionResponseContent(val *SponsoredProductsUnauthenticatedExceptionResponseContent) *NullableSponsoredProductsUnauthenticatedExceptionResponseContent {
	return &NullableSponsoredProductsUnauthenticatedExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUnauthenticatedExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUnauthenticatedExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
