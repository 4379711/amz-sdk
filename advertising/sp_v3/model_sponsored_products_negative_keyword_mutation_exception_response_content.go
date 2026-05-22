package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNegativeKeywordMutationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNegativeKeywordMutationExceptionResponseContent{}

// SponsoredProductsNegativeKeywordMutationExceptionResponseContent Exception resulting in mutating campaign management entities
type SponsoredProductsNegativeKeywordMutationExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                          `json:"message"`
	Errors  []SponsoredProductsNegativeKeywordMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsNegativeKeywordMutationExceptionResponseContent SponsoredProductsNegativeKeywordMutationExceptionResponseContent

// NewSponsoredProductsNegativeKeywordMutationExceptionResponseContent instantiates a new SponsoredProductsNegativeKeywordMutationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNegativeKeywordMutationExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsNegativeKeywordMutationExceptionResponseContent {
	this := SponsoredProductsNegativeKeywordMutationExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsNegativeKeywordMutationExceptionResponseContentWithDefaults instantiates a new SponsoredProductsNegativeKeywordMutationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNegativeKeywordMutationExceptionResponseContentWithDefaults() *SponsoredProductsNegativeKeywordMutationExceptionResponseContent {
	this := SponsoredProductsNegativeKeywordMutationExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsNegativeKeywordMutationExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordMutationExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsNegativeKeywordMutationExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsNegativeKeywordMutationExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordMutationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsNegativeKeywordMutationExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordMutationExceptionResponseContent) GetErrors() []SponsoredProductsNegativeKeywordMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsNegativeKeywordMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordMutationExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsNegativeKeywordMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordMutationExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsNegativeKeywordMutationError and assigns it to the Errors field.
func (o *SponsoredProductsNegativeKeywordMutationExceptionResponseContent) SetErrors(v []SponsoredProductsNegativeKeywordMutationError) {
	o.Errors = v
}

func (o SponsoredProductsNegativeKeywordMutationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsNegativeKeywordMutationExceptionResponseContent struct {
	value *SponsoredProductsNegativeKeywordMutationExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsNegativeKeywordMutationExceptionResponseContent) Get() *SponsoredProductsNegativeKeywordMutationExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsNegativeKeywordMutationExceptionResponseContent) Set(val *SponsoredProductsNegativeKeywordMutationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNegativeKeywordMutationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNegativeKeywordMutationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNegativeKeywordMutationExceptionResponseContent(val *SponsoredProductsNegativeKeywordMutationExceptionResponseContent) *NullableSponsoredProductsNegativeKeywordMutationExceptionResponseContent {
	return &NullableSponsoredProductsNegativeKeywordMutationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsNegativeKeywordMutationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNegativeKeywordMutationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
