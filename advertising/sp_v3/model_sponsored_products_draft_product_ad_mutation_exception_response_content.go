package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftProductAdMutationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftProductAdMutationExceptionResponseContent{}

// SponsoredProductsDraftProductAdMutationExceptionResponseContent Exception resulting in mutating campaign management entities
type SponsoredProductsDraftProductAdMutationExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                         `json:"message"`
	Errors  []SponsoredProductsDraftProductAdMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftProductAdMutationExceptionResponseContent SponsoredProductsDraftProductAdMutationExceptionResponseContent

// NewSponsoredProductsDraftProductAdMutationExceptionResponseContent instantiates a new SponsoredProductsDraftProductAdMutationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftProductAdMutationExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsDraftProductAdMutationExceptionResponseContent {
	this := SponsoredProductsDraftProductAdMutationExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsDraftProductAdMutationExceptionResponseContentWithDefaults instantiates a new SponsoredProductsDraftProductAdMutationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftProductAdMutationExceptionResponseContentWithDefaults() *SponsoredProductsDraftProductAdMutationExceptionResponseContent {
	this := SponsoredProductsDraftProductAdMutationExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsDraftProductAdMutationExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdMutationExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsDraftProductAdMutationExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsDraftProductAdMutationExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdMutationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsDraftProductAdMutationExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdMutationExceptionResponseContent) GetErrors() []SponsoredProductsDraftProductAdMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftProductAdMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdMutationExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsDraftProductAdMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdMutationExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftProductAdMutationError and assigns it to the Errors field.
func (o *SponsoredProductsDraftProductAdMutationExceptionResponseContent) SetErrors(v []SponsoredProductsDraftProductAdMutationError) {
	o.Errors = v
}

func (o SponsoredProductsDraftProductAdMutationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftProductAdMutationExceptionResponseContent struct {
	value *SponsoredProductsDraftProductAdMutationExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDraftProductAdMutationExceptionResponseContent) Get() *SponsoredProductsDraftProductAdMutationExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDraftProductAdMutationExceptionResponseContent) Set(val *SponsoredProductsDraftProductAdMutationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftProductAdMutationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftProductAdMutationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftProductAdMutationExceptionResponseContent(val *SponsoredProductsDraftProductAdMutationExceptionResponseContent) *NullableSponsoredProductsDraftProductAdMutationExceptionResponseContent {
	return &NullableSponsoredProductsDraftProductAdMutationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftProductAdMutationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftProductAdMutationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
