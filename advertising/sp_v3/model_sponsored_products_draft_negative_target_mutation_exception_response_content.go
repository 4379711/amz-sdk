package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent{}

// SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent Exception resulting in mutating campaign management entities
type SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                              `json:"message"`
	Errors  []SponsoredProductsDraftNegativeTargetMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent

// NewSponsoredProductsDraftNegativeTargetMutationExceptionResponseContent instantiates a new SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftNegativeTargetMutationExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent {
	this := SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsDraftNegativeTargetMutationExceptionResponseContentWithDefaults instantiates a new SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftNegativeTargetMutationExceptionResponseContentWithDefaults() *SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent {
	this := SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) GetErrors() []SponsoredProductsDraftNegativeTargetMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftNegativeTargetMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsDraftNegativeTargetMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftNegativeTargetMutationError and assigns it to the Errors field.
func (o *SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) SetErrors(v []SponsoredProductsDraftNegativeTargetMutationError) {
	o.Errors = v
}

func (o SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftNegativeTargetMutationExceptionResponseContent struct {
	value *SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) Get() *SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) Set(val *SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftNegativeTargetMutationExceptionResponseContent(val *SponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) *NullableSponsoredProductsDraftNegativeTargetMutationExceptionResponseContent {
	return &NullableSponsoredProductsDraftNegativeTargetMutationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftNegativeTargetMutationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
