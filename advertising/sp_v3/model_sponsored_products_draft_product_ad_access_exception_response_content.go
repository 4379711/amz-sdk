package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftProductAdAccessExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftProductAdAccessExceptionResponseContent{}

// SponsoredProductsDraftProductAdAccessExceptionResponseContent Exception resulting in accessing campaign management entities
type SponsoredProductsDraftProductAdAccessExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                       `json:"message"`
	Errors  []SponsoredProductsDraftProductAdAccessError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftProductAdAccessExceptionResponseContent SponsoredProductsDraftProductAdAccessExceptionResponseContent

// NewSponsoredProductsDraftProductAdAccessExceptionResponseContent instantiates a new SponsoredProductsDraftProductAdAccessExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftProductAdAccessExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsDraftProductAdAccessExceptionResponseContent {
	this := SponsoredProductsDraftProductAdAccessExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsDraftProductAdAccessExceptionResponseContentWithDefaults instantiates a new SponsoredProductsDraftProductAdAccessExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftProductAdAccessExceptionResponseContentWithDefaults() *SponsoredProductsDraftProductAdAccessExceptionResponseContent {
	this := SponsoredProductsDraftProductAdAccessExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsDraftProductAdAccessExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdAccessExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsDraftProductAdAccessExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsDraftProductAdAccessExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdAccessExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsDraftProductAdAccessExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdAccessExceptionResponseContent) GetErrors() []SponsoredProductsDraftProductAdAccessError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftProductAdAccessError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdAccessExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsDraftProductAdAccessError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdAccessExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftProductAdAccessError and assigns it to the Errors field.
func (o *SponsoredProductsDraftProductAdAccessExceptionResponseContent) SetErrors(v []SponsoredProductsDraftProductAdAccessError) {
	o.Errors = v
}

func (o SponsoredProductsDraftProductAdAccessExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftProductAdAccessExceptionResponseContent struct {
	value *SponsoredProductsDraftProductAdAccessExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDraftProductAdAccessExceptionResponseContent) Get() *SponsoredProductsDraftProductAdAccessExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDraftProductAdAccessExceptionResponseContent) Set(val *SponsoredProductsDraftProductAdAccessExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftProductAdAccessExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftProductAdAccessExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftProductAdAccessExceptionResponseContent(val *SponsoredProductsDraftProductAdAccessExceptionResponseContent) *NullableSponsoredProductsDraftProductAdAccessExceptionResponseContent {
	return &NullableSponsoredProductsDraftProductAdAccessExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftProductAdAccessExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftProductAdAccessExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
