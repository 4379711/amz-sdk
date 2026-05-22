package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftKeywordMutationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftKeywordMutationExceptionResponseContent{}

// SponsoredProductsDraftKeywordMutationExceptionResponseContent Exception resulting in mutating campaign management entities
type SponsoredProductsDraftKeywordMutationExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                       `json:"message"`
	Errors  []SponsoredProductsDraftKeywordMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftKeywordMutationExceptionResponseContent SponsoredProductsDraftKeywordMutationExceptionResponseContent

// NewSponsoredProductsDraftKeywordMutationExceptionResponseContent instantiates a new SponsoredProductsDraftKeywordMutationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftKeywordMutationExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsDraftKeywordMutationExceptionResponseContent {
	this := SponsoredProductsDraftKeywordMutationExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsDraftKeywordMutationExceptionResponseContentWithDefaults instantiates a new SponsoredProductsDraftKeywordMutationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftKeywordMutationExceptionResponseContentWithDefaults() *SponsoredProductsDraftKeywordMutationExceptionResponseContent {
	this := SponsoredProductsDraftKeywordMutationExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsDraftKeywordMutationExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordMutationExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsDraftKeywordMutationExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsDraftKeywordMutationExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordMutationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsDraftKeywordMutationExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeywordMutationExceptionResponseContent) GetErrors() []SponsoredProductsDraftKeywordMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftKeywordMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordMutationExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsDraftKeywordMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeywordMutationExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftKeywordMutationError and assigns it to the Errors field.
func (o *SponsoredProductsDraftKeywordMutationExceptionResponseContent) SetErrors(v []SponsoredProductsDraftKeywordMutationError) {
	o.Errors = v
}

func (o SponsoredProductsDraftKeywordMutationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftKeywordMutationExceptionResponseContent struct {
	value *SponsoredProductsDraftKeywordMutationExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDraftKeywordMutationExceptionResponseContent) Get() *SponsoredProductsDraftKeywordMutationExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDraftKeywordMutationExceptionResponseContent) Set(val *SponsoredProductsDraftKeywordMutationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftKeywordMutationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftKeywordMutationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftKeywordMutationExceptionResponseContent(val *SponsoredProductsDraftKeywordMutationExceptionResponseContent) *NullableSponsoredProductsDraftKeywordMutationExceptionResponseContent {
	return &NullableSponsoredProductsDraftKeywordMutationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftKeywordMutationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftKeywordMutationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
