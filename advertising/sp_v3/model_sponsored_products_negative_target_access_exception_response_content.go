package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNegativeTargetAccessExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNegativeTargetAccessExceptionResponseContent{}

// SponsoredProductsNegativeTargetAccessExceptionResponseContent Exception resulting in accessing campaign management entities
type SponsoredProductsNegativeTargetAccessExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                       `json:"message"`
	Errors  []SponsoredProductsNegativeTargetAccessError `json:"errors,omitempty"`
}

type _SponsoredProductsNegativeTargetAccessExceptionResponseContent SponsoredProductsNegativeTargetAccessExceptionResponseContent

// NewSponsoredProductsNegativeTargetAccessExceptionResponseContent instantiates a new SponsoredProductsNegativeTargetAccessExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNegativeTargetAccessExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsNegativeTargetAccessExceptionResponseContent {
	this := SponsoredProductsNegativeTargetAccessExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsNegativeTargetAccessExceptionResponseContentWithDefaults instantiates a new SponsoredProductsNegativeTargetAccessExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNegativeTargetAccessExceptionResponseContentWithDefaults() *SponsoredProductsNegativeTargetAccessExceptionResponseContent {
	this := SponsoredProductsNegativeTargetAccessExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsNegativeTargetAccessExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetAccessExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsNegativeTargetAccessExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsNegativeTargetAccessExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetAccessExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsNegativeTargetAccessExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeTargetAccessExceptionResponseContent) GetErrors() []SponsoredProductsNegativeTargetAccessError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsNegativeTargetAccessError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetAccessExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsNegativeTargetAccessError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeTargetAccessExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsNegativeTargetAccessError and assigns it to the Errors field.
func (o *SponsoredProductsNegativeTargetAccessExceptionResponseContent) SetErrors(v []SponsoredProductsNegativeTargetAccessError) {
	o.Errors = v
}

func (o SponsoredProductsNegativeTargetAccessExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsNegativeTargetAccessExceptionResponseContent struct {
	value *SponsoredProductsNegativeTargetAccessExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsNegativeTargetAccessExceptionResponseContent) Get() *SponsoredProductsNegativeTargetAccessExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsNegativeTargetAccessExceptionResponseContent) Set(val *SponsoredProductsNegativeTargetAccessExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNegativeTargetAccessExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNegativeTargetAccessExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNegativeTargetAccessExceptionResponseContent(val *SponsoredProductsNegativeTargetAccessExceptionResponseContent) *NullableSponsoredProductsNegativeTargetAccessExceptionResponseContent {
	return &NullableSponsoredProductsNegativeTargetAccessExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsNegativeTargetAccessExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNegativeTargetAccessExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
