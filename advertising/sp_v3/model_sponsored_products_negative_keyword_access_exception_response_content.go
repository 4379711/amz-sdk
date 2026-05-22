package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNegativeKeywordAccessExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNegativeKeywordAccessExceptionResponseContent{}

// SponsoredProductsNegativeKeywordAccessExceptionResponseContent Exception resulting in accessing campaign management entities
type SponsoredProductsNegativeKeywordAccessExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                        `json:"message"`
	Errors  []SponsoredProductsNegativeKeywordAccessError `json:"errors,omitempty"`
}

type _SponsoredProductsNegativeKeywordAccessExceptionResponseContent SponsoredProductsNegativeKeywordAccessExceptionResponseContent

// NewSponsoredProductsNegativeKeywordAccessExceptionResponseContent instantiates a new SponsoredProductsNegativeKeywordAccessExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNegativeKeywordAccessExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsNegativeKeywordAccessExceptionResponseContent {
	this := SponsoredProductsNegativeKeywordAccessExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsNegativeKeywordAccessExceptionResponseContentWithDefaults instantiates a new SponsoredProductsNegativeKeywordAccessExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNegativeKeywordAccessExceptionResponseContentWithDefaults() *SponsoredProductsNegativeKeywordAccessExceptionResponseContent {
	this := SponsoredProductsNegativeKeywordAccessExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsNegativeKeywordAccessExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordAccessExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsNegativeKeywordAccessExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsNegativeKeywordAccessExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordAccessExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsNegativeKeywordAccessExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordAccessExceptionResponseContent) GetErrors() []SponsoredProductsNegativeKeywordAccessError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsNegativeKeywordAccessError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordAccessExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsNegativeKeywordAccessError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordAccessExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsNegativeKeywordAccessError and assigns it to the Errors field.
func (o *SponsoredProductsNegativeKeywordAccessExceptionResponseContent) SetErrors(v []SponsoredProductsNegativeKeywordAccessError) {
	o.Errors = v
}

func (o SponsoredProductsNegativeKeywordAccessExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsNegativeKeywordAccessExceptionResponseContent struct {
	value *SponsoredProductsNegativeKeywordAccessExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsNegativeKeywordAccessExceptionResponseContent) Get() *SponsoredProductsNegativeKeywordAccessExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsNegativeKeywordAccessExceptionResponseContent) Set(val *SponsoredProductsNegativeKeywordAccessExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNegativeKeywordAccessExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNegativeKeywordAccessExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNegativeKeywordAccessExceptionResponseContent(val *SponsoredProductsNegativeKeywordAccessExceptionResponseContent) *NullableSponsoredProductsNegativeKeywordAccessExceptionResponseContent {
	return &NullableSponsoredProductsNegativeKeywordAccessExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsNegativeKeywordAccessExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNegativeKeywordAccessExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
