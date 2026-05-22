package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBatchResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBatchResponseError{}

// SponsoredProductsBatchResponseError struct for SponsoredProductsBatchResponseError
type SponsoredProductsBatchResponseError struct {
	Code    SponsoredProductsBatchErrorCode `json:"code"`
	Message string                          `json:"message"`
}

type _SponsoredProductsBatchResponseError SponsoredProductsBatchResponseError

// NewSponsoredProductsBatchResponseError instantiates a new SponsoredProductsBatchResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBatchResponseError(code SponsoredProductsBatchErrorCode, message string) *SponsoredProductsBatchResponseError {
	this := SponsoredProductsBatchResponseError{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsBatchResponseErrorWithDefaults instantiates a new SponsoredProductsBatchResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBatchResponseErrorWithDefaults() *SponsoredProductsBatchResponseError {
	this := SponsoredProductsBatchResponseError{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsBatchResponseError) GetCode() SponsoredProductsBatchErrorCode {
	if o == nil {
		var ret SponsoredProductsBatchErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBatchResponseError) GetCodeOk() (*SponsoredProductsBatchErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsBatchResponseError) SetCode(v SponsoredProductsBatchErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsBatchResponseError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBatchResponseError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsBatchResponseError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsBatchResponseError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsBatchResponseError struct {
	value *SponsoredProductsBatchResponseError
	isSet bool
}

func (v NullableSponsoredProductsBatchResponseError) Get() *SponsoredProductsBatchResponseError {
	return v.value
}

func (v *NullableSponsoredProductsBatchResponseError) Set(val *SponsoredProductsBatchResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBatchResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBatchResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBatchResponseError(val *SponsoredProductsBatchResponseError) *NullableSponsoredProductsBatchResponseError {
	return &NullableSponsoredProductsBatchResponseError{value: val, isSet: true}
}

func (v NullableSponsoredProductsBatchResponseError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBatchResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
