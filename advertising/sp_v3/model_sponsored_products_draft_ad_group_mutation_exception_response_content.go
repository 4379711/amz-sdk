package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftAdGroupMutationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftAdGroupMutationExceptionResponseContent{}

// SponsoredProductsDraftAdGroupMutationExceptionResponseContent Exception resulting in mutating campaign management entities
type SponsoredProductsDraftAdGroupMutationExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                       `json:"message"`
	Errors  []SponsoredProductsDraftAdGroupMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftAdGroupMutationExceptionResponseContent SponsoredProductsDraftAdGroupMutationExceptionResponseContent

// NewSponsoredProductsDraftAdGroupMutationExceptionResponseContent instantiates a new SponsoredProductsDraftAdGroupMutationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftAdGroupMutationExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsDraftAdGroupMutationExceptionResponseContent {
	this := SponsoredProductsDraftAdGroupMutationExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsDraftAdGroupMutationExceptionResponseContentWithDefaults instantiates a new SponsoredProductsDraftAdGroupMutationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftAdGroupMutationExceptionResponseContentWithDefaults() *SponsoredProductsDraftAdGroupMutationExceptionResponseContent {
	this := SponsoredProductsDraftAdGroupMutationExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsDraftAdGroupMutationExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroupMutationExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsDraftAdGroupMutationExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsDraftAdGroupMutationExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroupMutationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsDraftAdGroupMutationExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftAdGroupMutationExceptionResponseContent) GetErrors() []SponsoredProductsDraftAdGroupMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftAdGroupMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroupMutationExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsDraftAdGroupMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftAdGroupMutationExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftAdGroupMutationError and assigns it to the Errors field.
func (o *SponsoredProductsDraftAdGroupMutationExceptionResponseContent) SetErrors(v []SponsoredProductsDraftAdGroupMutationError) {
	o.Errors = v
}

func (o SponsoredProductsDraftAdGroupMutationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftAdGroupMutationExceptionResponseContent struct {
	value *SponsoredProductsDraftAdGroupMutationExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDraftAdGroupMutationExceptionResponseContent) Get() *SponsoredProductsDraftAdGroupMutationExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDraftAdGroupMutationExceptionResponseContent) Set(val *SponsoredProductsDraftAdGroupMutationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftAdGroupMutationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftAdGroupMutationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftAdGroupMutationExceptionResponseContent(val *SponsoredProductsDraftAdGroupMutationExceptionResponseContent) *NullableSponsoredProductsDraftAdGroupMutationExceptionResponseContent {
	return &NullableSponsoredProductsDraftAdGroupMutationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftAdGroupMutationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftAdGroupMutationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
