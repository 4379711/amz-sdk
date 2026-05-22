package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsProductAdMutationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsProductAdMutationExceptionResponseContent{}

// SponsoredProductsProductAdMutationExceptionResponseContent Exception resulting in mutating campaign management entities
type SponsoredProductsProductAdMutationExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                    `json:"message"`
	Errors  []SponsoredProductsProductAdMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsProductAdMutationExceptionResponseContent SponsoredProductsProductAdMutationExceptionResponseContent

// NewSponsoredProductsProductAdMutationExceptionResponseContent instantiates a new SponsoredProductsProductAdMutationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsProductAdMutationExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsProductAdMutationExceptionResponseContent {
	this := SponsoredProductsProductAdMutationExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsProductAdMutationExceptionResponseContentWithDefaults instantiates a new SponsoredProductsProductAdMutationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsProductAdMutationExceptionResponseContentWithDefaults() *SponsoredProductsProductAdMutationExceptionResponseContent {
	this := SponsoredProductsProductAdMutationExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsProductAdMutationExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdMutationExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsProductAdMutationExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsProductAdMutationExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdMutationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsProductAdMutationExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsProductAdMutationExceptionResponseContent) GetErrors() []SponsoredProductsProductAdMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsProductAdMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdMutationExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsProductAdMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsProductAdMutationExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsProductAdMutationError and assigns it to the Errors field.
func (o *SponsoredProductsProductAdMutationExceptionResponseContent) SetErrors(v []SponsoredProductsProductAdMutationError) {
	o.Errors = v
}

func (o SponsoredProductsProductAdMutationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsProductAdMutationExceptionResponseContent struct {
	value *SponsoredProductsProductAdMutationExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsProductAdMutationExceptionResponseContent) Get() *SponsoredProductsProductAdMutationExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsProductAdMutationExceptionResponseContent) Set(val *SponsoredProductsProductAdMutationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsProductAdMutationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsProductAdMutationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsProductAdMutationExceptionResponseContent(val *SponsoredProductsProductAdMutationExceptionResponseContent) *NullableSponsoredProductsProductAdMutationExceptionResponseContent {
	return &NullableSponsoredProductsProductAdMutationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsProductAdMutationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsProductAdMutationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
