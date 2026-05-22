package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent{}

// SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent Exception resulting in accessing campaign management entities
type SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                             `json:"message"`
	Errors  []SponsoredProductsDraftNegativeKeywordAccessError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent

// NewSponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent instantiates a new SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent {
	this := SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsDraftNegativeKeywordAccessExceptionResponseContentWithDefaults instantiates a new SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftNegativeKeywordAccessExceptionResponseContentWithDefaults() *SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent {
	this := SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) GetErrors() []SponsoredProductsDraftNegativeKeywordAccessError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftNegativeKeywordAccessError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsDraftNegativeKeywordAccessError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftNegativeKeywordAccessError and assigns it to the Errors field.
func (o *SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) SetErrors(v []SponsoredProductsDraftNegativeKeywordAccessError) {
	o.Errors = v
}

func (o SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent struct {
	value *SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) Get() *SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) Set(val *SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent(val *SponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) *NullableSponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent {
	return &NullableSponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftNegativeKeywordAccessExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
