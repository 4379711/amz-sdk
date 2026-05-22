package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent{}

// SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent Exception resulting in mutating campaign management entities
type SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                               `json:"message"`
	Errors  []SponsoredProductsDraftNegativeKeywordMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent

// NewSponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent instantiates a new SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent {
	this := SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsDraftNegativeKeywordMutationExceptionResponseContentWithDefaults instantiates a new SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftNegativeKeywordMutationExceptionResponseContentWithDefaults() *SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent {
	this := SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) GetErrors() []SponsoredProductsDraftNegativeKeywordMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftNegativeKeywordMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsDraftNegativeKeywordMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftNegativeKeywordMutationError and assigns it to the Errors field.
func (o *SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) SetErrors(v []SponsoredProductsDraftNegativeKeywordMutationError) {
	o.Errors = v
}

func (o SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent struct {
	value *SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) Get() *SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) Set(val *SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent(val *SponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) *NullableSponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent {
	return &NullableSponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftNegativeKeywordMutationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
