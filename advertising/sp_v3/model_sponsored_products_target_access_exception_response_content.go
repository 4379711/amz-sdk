package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsTargetAccessExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsTargetAccessExceptionResponseContent{}

// SponsoredProductsTargetAccessExceptionResponseContent Exception resulting in accessing campaign management entities
type SponsoredProductsTargetAccessExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                               `json:"message"`
	Errors  []SponsoredProductsTargetAccessError `json:"errors,omitempty"`
}

type _SponsoredProductsTargetAccessExceptionResponseContent SponsoredProductsTargetAccessExceptionResponseContent

// NewSponsoredProductsTargetAccessExceptionResponseContent instantiates a new SponsoredProductsTargetAccessExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsTargetAccessExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsTargetAccessExceptionResponseContent {
	this := SponsoredProductsTargetAccessExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsTargetAccessExceptionResponseContentWithDefaults instantiates a new SponsoredProductsTargetAccessExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsTargetAccessExceptionResponseContentWithDefaults() *SponsoredProductsTargetAccessExceptionResponseContent {
	this := SponsoredProductsTargetAccessExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsTargetAccessExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetAccessExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsTargetAccessExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsTargetAccessExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetAccessExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsTargetAccessExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsTargetAccessExceptionResponseContent) GetErrors() []SponsoredProductsTargetAccessError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsTargetAccessError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetAccessExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsTargetAccessError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsTargetAccessExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsTargetAccessError and assigns it to the Errors field.
func (o *SponsoredProductsTargetAccessExceptionResponseContent) SetErrors(v []SponsoredProductsTargetAccessError) {
	o.Errors = v
}

func (o SponsoredProductsTargetAccessExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsTargetAccessExceptionResponseContent struct {
	value *SponsoredProductsTargetAccessExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsTargetAccessExceptionResponseContent) Get() *SponsoredProductsTargetAccessExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsTargetAccessExceptionResponseContent) Set(val *SponsoredProductsTargetAccessExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetAccessExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetAccessExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetAccessExceptionResponseContent(val *SponsoredProductsTargetAccessExceptionResponseContent) *NullableSponsoredProductsTargetAccessExceptionResponseContent {
	return &NullableSponsoredProductsTargetAccessExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetAccessExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetAccessExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
