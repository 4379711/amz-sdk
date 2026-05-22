package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsTargetMutationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsTargetMutationExceptionResponseContent{}

// SponsoredProductsTargetMutationExceptionResponseContent Exception resulting in mutating campaign management entities
type SponsoredProductsTargetMutationExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                 `json:"message"`
	Errors  []SponsoredProductsTargetMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsTargetMutationExceptionResponseContent SponsoredProductsTargetMutationExceptionResponseContent

// NewSponsoredProductsTargetMutationExceptionResponseContent instantiates a new SponsoredProductsTargetMutationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsTargetMutationExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsTargetMutationExceptionResponseContent {
	this := SponsoredProductsTargetMutationExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsTargetMutationExceptionResponseContentWithDefaults instantiates a new SponsoredProductsTargetMutationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsTargetMutationExceptionResponseContentWithDefaults() *SponsoredProductsTargetMutationExceptionResponseContent {
	this := SponsoredProductsTargetMutationExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsTargetMutationExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetMutationExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsTargetMutationExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsTargetMutationExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetMutationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsTargetMutationExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsTargetMutationExceptionResponseContent) GetErrors() []SponsoredProductsTargetMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsTargetMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetMutationExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsTargetMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsTargetMutationExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsTargetMutationError and assigns it to the Errors field.
func (o *SponsoredProductsTargetMutationExceptionResponseContent) SetErrors(v []SponsoredProductsTargetMutationError) {
	o.Errors = v
}

func (o SponsoredProductsTargetMutationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsTargetMutationExceptionResponseContent struct {
	value *SponsoredProductsTargetMutationExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsTargetMutationExceptionResponseContent) Get() *SponsoredProductsTargetMutationExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsTargetMutationExceptionResponseContent) Set(val *SponsoredProductsTargetMutationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetMutationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetMutationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetMutationExceptionResponseContent(val *SponsoredProductsTargetMutationExceptionResponseContent) *NullableSponsoredProductsTargetMutationExceptionResponseContent {
	return &NullableSponsoredProductsTargetMutationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetMutationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetMutationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
