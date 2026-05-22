package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftKeywordAccessExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftKeywordAccessExceptionResponseContent{}

// SponsoredProductsDraftKeywordAccessExceptionResponseContent Exception resulting in accessing campaign management entities
type SponsoredProductsDraftKeywordAccessExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                     `json:"message"`
	Errors  []SponsoredProductsDraftKeywordAccessError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftKeywordAccessExceptionResponseContent SponsoredProductsDraftKeywordAccessExceptionResponseContent

// NewSponsoredProductsDraftKeywordAccessExceptionResponseContent instantiates a new SponsoredProductsDraftKeywordAccessExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftKeywordAccessExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsDraftKeywordAccessExceptionResponseContent {
	this := SponsoredProductsDraftKeywordAccessExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsDraftKeywordAccessExceptionResponseContentWithDefaults instantiates a new SponsoredProductsDraftKeywordAccessExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftKeywordAccessExceptionResponseContentWithDefaults() *SponsoredProductsDraftKeywordAccessExceptionResponseContent {
	this := SponsoredProductsDraftKeywordAccessExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsDraftKeywordAccessExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordAccessExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsDraftKeywordAccessExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsDraftKeywordAccessExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordAccessExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsDraftKeywordAccessExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeywordAccessExceptionResponseContent) GetErrors() []SponsoredProductsDraftKeywordAccessError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftKeywordAccessError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordAccessExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsDraftKeywordAccessError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeywordAccessExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftKeywordAccessError and assigns it to the Errors field.
func (o *SponsoredProductsDraftKeywordAccessExceptionResponseContent) SetErrors(v []SponsoredProductsDraftKeywordAccessError) {
	o.Errors = v
}

func (o SponsoredProductsDraftKeywordAccessExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftKeywordAccessExceptionResponseContent struct {
	value *SponsoredProductsDraftKeywordAccessExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDraftKeywordAccessExceptionResponseContent) Get() *SponsoredProductsDraftKeywordAccessExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDraftKeywordAccessExceptionResponseContent) Set(val *SponsoredProductsDraftKeywordAccessExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftKeywordAccessExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftKeywordAccessExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftKeywordAccessExceptionResponseContent(val *SponsoredProductsDraftKeywordAccessExceptionResponseContent) *NullableSponsoredProductsDraftKeywordAccessExceptionResponseContent {
	return &NullableSponsoredProductsDraftKeywordAccessExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftKeywordAccessExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftKeywordAccessExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
