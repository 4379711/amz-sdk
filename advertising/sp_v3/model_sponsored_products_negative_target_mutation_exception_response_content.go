package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsNegativeTargetMutationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNegativeTargetMutationExceptionResponseContent{}

// SponsoredProductsNegativeTargetMutationExceptionResponseContent Exception resulting in mutating campaign management entities
type SponsoredProductsNegativeTargetMutationExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                         `json:"message"`
	Errors  []SponsoredProductsNegativeTargetMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsNegativeTargetMutationExceptionResponseContent SponsoredProductsNegativeTargetMutationExceptionResponseContent

// NewSponsoredProductsNegativeTargetMutationExceptionResponseContent instantiates a new SponsoredProductsNegativeTargetMutationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNegativeTargetMutationExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsNegativeTargetMutationExceptionResponseContent {
	this := SponsoredProductsNegativeTargetMutationExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsNegativeTargetMutationExceptionResponseContentWithDefaults instantiates a new SponsoredProductsNegativeTargetMutationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNegativeTargetMutationExceptionResponseContentWithDefaults() *SponsoredProductsNegativeTargetMutationExceptionResponseContent {
	this := SponsoredProductsNegativeTargetMutationExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsNegativeTargetMutationExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetMutationExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsNegativeTargetMutationExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsNegativeTargetMutationExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetMutationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsNegativeTargetMutationExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeTargetMutationExceptionResponseContent) GetErrors() []SponsoredProductsNegativeTargetMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsNegativeTargetMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetMutationExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsNegativeTargetMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeTargetMutationExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsNegativeTargetMutationError and assigns it to the Errors field.
func (o *SponsoredProductsNegativeTargetMutationExceptionResponseContent) SetErrors(v []SponsoredProductsNegativeTargetMutationError) {
	o.Errors = v
}

func (o SponsoredProductsNegativeTargetMutationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsNegativeTargetMutationExceptionResponseContent struct {
	value *SponsoredProductsNegativeTargetMutationExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsNegativeTargetMutationExceptionResponseContent) Get() *SponsoredProductsNegativeTargetMutationExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsNegativeTargetMutationExceptionResponseContent) Set(val *SponsoredProductsNegativeTargetMutationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNegativeTargetMutationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNegativeTargetMutationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNegativeTargetMutationExceptionResponseContent(val *SponsoredProductsNegativeTargetMutationExceptionResponseContent) *NullableSponsoredProductsNegativeTargetMutationExceptionResponseContent {
	return &NullableSponsoredProductsNegativeTargetMutationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsNegativeTargetMutationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNegativeTargetMutationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
