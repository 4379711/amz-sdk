package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsProductAdAccessExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsProductAdAccessExceptionResponseContent{}

// SponsoredProductsProductAdAccessExceptionResponseContent Exception resulting in accessing campaign management entities
type SponsoredProductsProductAdAccessExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                  `json:"message"`
	Errors  []SponsoredProductsProductAdAccessError `json:"errors,omitempty"`
}

type _SponsoredProductsProductAdAccessExceptionResponseContent SponsoredProductsProductAdAccessExceptionResponseContent

// NewSponsoredProductsProductAdAccessExceptionResponseContent instantiates a new SponsoredProductsProductAdAccessExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsProductAdAccessExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsProductAdAccessExceptionResponseContent {
	this := SponsoredProductsProductAdAccessExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsProductAdAccessExceptionResponseContentWithDefaults instantiates a new SponsoredProductsProductAdAccessExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsProductAdAccessExceptionResponseContentWithDefaults() *SponsoredProductsProductAdAccessExceptionResponseContent {
	this := SponsoredProductsProductAdAccessExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsProductAdAccessExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdAccessExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsProductAdAccessExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsProductAdAccessExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdAccessExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsProductAdAccessExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsProductAdAccessExceptionResponseContent) GetErrors() []SponsoredProductsProductAdAccessError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsProductAdAccessError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdAccessExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsProductAdAccessError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsProductAdAccessExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsProductAdAccessError and assigns it to the Errors field.
func (o *SponsoredProductsProductAdAccessExceptionResponseContent) SetErrors(v []SponsoredProductsProductAdAccessError) {
	o.Errors = v
}

func (o SponsoredProductsProductAdAccessExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsProductAdAccessExceptionResponseContent struct {
	value *SponsoredProductsProductAdAccessExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsProductAdAccessExceptionResponseContent) Get() *SponsoredProductsProductAdAccessExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsProductAdAccessExceptionResponseContent) Set(val *SponsoredProductsProductAdAccessExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsProductAdAccessExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsProductAdAccessExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsProductAdAccessExceptionResponseContent(val *SponsoredProductsProductAdAccessExceptionResponseContent) *NullableSponsoredProductsProductAdAccessExceptionResponseContent {
	return &NullableSponsoredProductsProductAdAccessExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsProductAdAccessExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsProductAdAccessExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
