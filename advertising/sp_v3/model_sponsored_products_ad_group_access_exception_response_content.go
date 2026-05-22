package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsAdGroupAccessExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsAdGroupAccessExceptionResponseContent{}

// SponsoredProductsAdGroupAccessExceptionResponseContent Exception resulting in accessing campaign management entities
type SponsoredProductsAdGroupAccessExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                `json:"message"`
	Errors  []SponsoredProductsAdGroupAccessError `json:"errors,omitempty"`
}

type _SponsoredProductsAdGroupAccessExceptionResponseContent SponsoredProductsAdGroupAccessExceptionResponseContent

// NewSponsoredProductsAdGroupAccessExceptionResponseContent instantiates a new SponsoredProductsAdGroupAccessExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsAdGroupAccessExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsAdGroupAccessExceptionResponseContent {
	this := SponsoredProductsAdGroupAccessExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsAdGroupAccessExceptionResponseContentWithDefaults instantiates a new SponsoredProductsAdGroupAccessExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsAdGroupAccessExceptionResponseContentWithDefaults() *SponsoredProductsAdGroupAccessExceptionResponseContent {
	this := SponsoredProductsAdGroupAccessExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsAdGroupAccessExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupAccessExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsAdGroupAccessExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsAdGroupAccessExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupAccessExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsAdGroupAccessExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupAccessExceptionResponseContent) GetErrors() []SponsoredProductsAdGroupAccessError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsAdGroupAccessError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupAccessExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsAdGroupAccessError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupAccessExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsAdGroupAccessError and assigns it to the Errors field.
func (o *SponsoredProductsAdGroupAccessExceptionResponseContent) SetErrors(v []SponsoredProductsAdGroupAccessError) {
	o.Errors = v
}

func (o SponsoredProductsAdGroupAccessExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsAdGroupAccessExceptionResponseContent struct {
	value *SponsoredProductsAdGroupAccessExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsAdGroupAccessExceptionResponseContent) Get() *SponsoredProductsAdGroupAccessExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsAdGroupAccessExceptionResponseContent) Set(val *SponsoredProductsAdGroupAccessExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAdGroupAccessExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAdGroupAccessExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAdGroupAccessExceptionResponseContent(val *SponsoredProductsAdGroupAccessExceptionResponseContent) *NullableSponsoredProductsAdGroupAccessExceptionResponseContent {
	return &NullableSponsoredProductsAdGroupAccessExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsAdGroupAccessExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAdGroupAccessExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
