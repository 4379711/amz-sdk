package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent{}

// SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent Exception resulting in accessing campaign management entities
type SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                            `json:"message"`
	Errors  []SponsoredProductsDraftNegativeTargetAccessError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent

// NewSponsoredProductsDraftNegativeTargetAccessExceptionResponseContent instantiates a new SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftNegativeTargetAccessExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent {
	this := SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsDraftNegativeTargetAccessExceptionResponseContentWithDefaults instantiates a new SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftNegativeTargetAccessExceptionResponseContentWithDefaults() *SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent {
	this := SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) GetErrors() []SponsoredProductsDraftNegativeTargetAccessError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftNegativeTargetAccessError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsDraftNegativeTargetAccessError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftNegativeTargetAccessError and assigns it to the Errors field.
func (o *SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) SetErrors(v []SponsoredProductsDraftNegativeTargetAccessError) {
	o.Errors = v
}

func (o SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftNegativeTargetAccessExceptionResponseContent struct {
	value *SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) Get() *SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) Set(val *SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftNegativeTargetAccessExceptionResponseContent(val *SponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) *NullableSponsoredProductsDraftNegativeTargetAccessExceptionResponseContent {
	return &NullableSponsoredProductsDraftNegativeTargetAccessExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftNegativeTargetAccessExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
