package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent{}

// SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent Exception resulting in mutating draft management entities
type SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                                       `json:"message"`
	Errors  []SponsoredProductsDraftCampaignNegativeKeywordMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent

// NewSponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent instantiates a new SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent {
	this := SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContentWithDefaults instantiates a new SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContentWithDefaults() *SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent {
	this := SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) GetErrors() []SponsoredProductsDraftCampaignNegativeKeywordMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftCampaignNegativeKeywordMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsDraftCampaignNegativeKeywordMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftCampaignNegativeKeywordMutationError and assigns it to the Errors field.
func (o *SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) SetErrors(v []SponsoredProductsDraftCampaignNegativeKeywordMutationError) {
	o.Errors = v
}

func (o SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent struct {
	value *SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) Get() *SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) Set(val *SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent(val *SponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) *NullableSponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent {
	return &NullableSponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordMutationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
