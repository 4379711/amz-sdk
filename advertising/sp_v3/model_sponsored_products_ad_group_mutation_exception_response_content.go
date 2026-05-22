package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsAdGroupMutationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsAdGroupMutationExceptionResponseContent{}

// SponsoredProductsAdGroupMutationExceptionResponseContent Exception resulting in mutating campaign management entities
type SponsoredProductsAdGroupMutationExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                  `json:"message"`
	Errors  []SponsoredProductsAdGroupMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsAdGroupMutationExceptionResponseContent SponsoredProductsAdGroupMutationExceptionResponseContent

// NewSponsoredProductsAdGroupMutationExceptionResponseContent instantiates a new SponsoredProductsAdGroupMutationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsAdGroupMutationExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsAdGroupMutationExceptionResponseContent {
	this := SponsoredProductsAdGroupMutationExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsAdGroupMutationExceptionResponseContentWithDefaults instantiates a new SponsoredProductsAdGroupMutationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsAdGroupMutationExceptionResponseContentWithDefaults() *SponsoredProductsAdGroupMutationExceptionResponseContent {
	this := SponsoredProductsAdGroupMutationExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsAdGroupMutationExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsAdGroupMutationExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsAdGroupMutationExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsAdGroupMutationExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupMutationExceptionResponseContent) GetErrors() []SponsoredProductsAdGroupMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsAdGroupMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsAdGroupMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupMutationExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsAdGroupMutationError and assigns it to the Errors field.
func (o *SponsoredProductsAdGroupMutationExceptionResponseContent) SetErrors(v []SponsoredProductsAdGroupMutationError) {
	o.Errors = v
}

func (o SponsoredProductsAdGroupMutationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsAdGroupMutationExceptionResponseContent struct {
	value *SponsoredProductsAdGroupMutationExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsAdGroupMutationExceptionResponseContent) Get() *SponsoredProductsAdGroupMutationExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsAdGroupMutationExceptionResponseContent) Set(val *SponsoredProductsAdGroupMutationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAdGroupMutationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAdGroupMutationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAdGroupMutationExceptionResponseContent(val *SponsoredProductsAdGroupMutationExceptionResponseContent) *NullableSponsoredProductsAdGroupMutationExceptionResponseContent {
	return &NullableSponsoredProductsAdGroupMutationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsAdGroupMutationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAdGroupMutationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
