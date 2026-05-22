package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsDraftAdGroupAccessExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftAdGroupAccessExceptionResponseContent{}

// SponsoredProductsDraftAdGroupAccessExceptionResponseContent Exception resulting in accessing campaign management entities
type SponsoredProductsDraftAdGroupAccessExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                     `json:"message"`
	Errors  []SponsoredProductsDraftAdGroupAccessError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftAdGroupAccessExceptionResponseContent SponsoredProductsDraftAdGroupAccessExceptionResponseContent

// NewSponsoredProductsDraftAdGroupAccessExceptionResponseContent instantiates a new SponsoredProductsDraftAdGroupAccessExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftAdGroupAccessExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsDraftAdGroupAccessExceptionResponseContent {
	this := SponsoredProductsDraftAdGroupAccessExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsDraftAdGroupAccessExceptionResponseContentWithDefaults instantiates a new SponsoredProductsDraftAdGroupAccessExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftAdGroupAccessExceptionResponseContentWithDefaults() *SponsoredProductsDraftAdGroupAccessExceptionResponseContent {
	this := SponsoredProductsDraftAdGroupAccessExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsDraftAdGroupAccessExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroupAccessExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsDraftAdGroupAccessExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsDraftAdGroupAccessExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroupAccessExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsDraftAdGroupAccessExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftAdGroupAccessExceptionResponseContent) GetErrors() []SponsoredProductsDraftAdGroupAccessError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftAdGroupAccessError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroupAccessExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsDraftAdGroupAccessError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftAdGroupAccessExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftAdGroupAccessError and assigns it to the Errors field.
func (o *SponsoredProductsDraftAdGroupAccessExceptionResponseContent) SetErrors(v []SponsoredProductsDraftAdGroupAccessError) {
	o.Errors = v
}

func (o SponsoredProductsDraftAdGroupAccessExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftAdGroupAccessExceptionResponseContent struct {
	value *SponsoredProductsDraftAdGroupAccessExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDraftAdGroupAccessExceptionResponseContent) Get() *SponsoredProductsDraftAdGroupAccessExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDraftAdGroupAccessExceptionResponseContent) Set(val *SponsoredProductsDraftAdGroupAccessExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftAdGroupAccessExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftAdGroupAccessExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftAdGroupAccessExceptionResponseContent(val *SponsoredProductsDraftAdGroupAccessExceptionResponseContent) *NullableSponsoredProductsDraftAdGroupAccessExceptionResponseContent {
	return &NullableSponsoredProductsDraftAdGroupAccessExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftAdGroupAccessExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftAdGroupAccessExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
