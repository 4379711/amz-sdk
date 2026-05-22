package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent{}

// SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent Exception resulting in mutating campaign management entities
type SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                                  `json:"message"`
	Errors  []SponsoredProductsCampaignNegativeKeywordMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent

// NewSponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent instantiates a new SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent {
	this := SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContentWithDefaults instantiates a new SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContentWithDefaults() *SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent {
	this := SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) GetErrors() []SponsoredProductsCampaignNegativeKeywordMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsCampaignNegativeKeywordMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsCampaignNegativeKeywordMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsCampaignNegativeKeywordMutationError and assigns it to the Errors field.
func (o *SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) SetErrors(v []SponsoredProductsCampaignNegativeKeywordMutationError) {
	o.Errors = v
}

func (o SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent struct {
	value *SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) Get() *SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) Set(val *SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent(val *SponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) *NullableSponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent {
	return &NullableSponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignNegativeKeywordMutationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
