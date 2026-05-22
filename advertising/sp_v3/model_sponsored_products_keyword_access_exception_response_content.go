package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsKeywordAccessExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsKeywordAccessExceptionResponseContent{}

// SponsoredProductsKeywordAccessExceptionResponseContent Exception resulting in accessing campaign management entities
type SponsoredProductsKeywordAccessExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                `json:"message"`
	Errors  []SponsoredProductsKeywordAccessError `json:"errors,omitempty"`
}

type _SponsoredProductsKeywordAccessExceptionResponseContent SponsoredProductsKeywordAccessExceptionResponseContent

// NewSponsoredProductsKeywordAccessExceptionResponseContent instantiates a new SponsoredProductsKeywordAccessExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsKeywordAccessExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsKeywordAccessExceptionResponseContent {
	this := SponsoredProductsKeywordAccessExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsKeywordAccessExceptionResponseContentWithDefaults instantiates a new SponsoredProductsKeywordAccessExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsKeywordAccessExceptionResponseContentWithDefaults() *SponsoredProductsKeywordAccessExceptionResponseContent {
	this := SponsoredProductsKeywordAccessExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsKeywordAccessExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordAccessExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsKeywordAccessExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsKeywordAccessExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordAccessExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsKeywordAccessExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsKeywordAccessExceptionResponseContent) GetErrors() []SponsoredProductsKeywordAccessError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsKeywordAccessError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordAccessExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsKeywordAccessError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsKeywordAccessExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsKeywordAccessError and assigns it to the Errors field.
func (o *SponsoredProductsKeywordAccessExceptionResponseContent) SetErrors(v []SponsoredProductsKeywordAccessError) {
	o.Errors = v
}

func (o SponsoredProductsKeywordAccessExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsKeywordAccessExceptionResponseContent struct {
	value *SponsoredProductsKeywordAccessExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsKeywordAccessExceptionResponseContent) Get() *SponsoredProductsKeywordAccessExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsKeywordAccessExceptionResponseContent) Set(val *SponsoredProductsKeywordAccessExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsKeywordAccessExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsKeywordAccessExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsKeywordAccessExceptionResponseContent(val *SponsoredProductsKeywordAccessExceptionResponseContent) *NullableSponsoredProductsKeywordAccessExceptionResponseContent {
	return &NullableSponsoredProductsKeywordAccessExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsKeywordAccessExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsKeywordAccessExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
