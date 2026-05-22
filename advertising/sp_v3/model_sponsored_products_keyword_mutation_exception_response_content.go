package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsKeywordMutationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsKeywordMutationExceptionResponseContent{}

// SponsoredProductsKeywordMutationExceptionResponseContent Exception resulting in mutating campaign management entities
type SponsoredProductsKeywordMutationExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                  `json:"message"`
	Errors  []SponsoredProductsKeywordMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsKeywordMutationExceptionResponseContent SponsoredProductsKeywordMutationExceptionResponseContent

// NewSponsoredProductsKeywordMutationExceptionResponseContent instantiates a new SponsoredProductsKeywordMutationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsKeywordMutationExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsKeywordMutationExceptionResponseContent {
	this := SponsoredProductsKeywordMutationExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsKeywordMutationExceptionResponseContentWithDefaults instantiates a new SponsoredProductsKeywordMutationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsKeywordMutationExceptionResponseContentWithDefaults() *SponsoredProductsKeywordMutationExceptionResponseContent {
	this := SponsoredProductsKeywordMutationExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsKeywordMutationExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordMutationExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsKeywordMutationExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsKeywordMutationExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordMutationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsKeywordMutationExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsKeywordMutationExceptionResponseContent) GetErrors() []SponsoredProductsKeywordMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsKeywordMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordMutationExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsKeywordMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsKeywordMutationExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsKeywordMutationError and assigns it to the Errors field.
func (o *SponsoredProductsKeywordMutationExceptionResponseContent) SetErrors(v []SponsoredProductsKeywordMutationError) {
	o.Errors = v
}

func (o SponsoredProductsKeywordMutationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsKeywordMutationExceptionResponseContent struct {
	value *SponsoredProductsKeywordMutationExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsKeywordMutationExceptionResponseContent) Get() *SponsoredProductsKeywordMutationExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsKeywordMutationExceptionResponseContent) Set(val *SponsoredProductsKeywordMutationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsKeywordMutationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsKeywordMutationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsKeywordMutationExceptionResponseContent(val *SponsoredProductsKeywordMutationExceptionResponseContent) *NullableSponsoredProductsKeywordMutationExceptionResponseContent {
	return &NullableSponsoredProductsKeywordMutationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsKeywordMutationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsKeywordMutationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
