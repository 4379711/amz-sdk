package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent{}

// SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent Exception resulting in accessing draft management entities
type SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                                     `json:"message"`
	Errors  []SponsoredProductsDraftCampaignNegativeKeywordAccessError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent

// NewSponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent instantiates a new SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent {
	this := SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContentWithDefaults instantiates a new SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContentWithDefaults() *SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent {
	this := SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) GetErrors() []SponsoredProductsDraftCampaignNegativeKeywordAccessError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftCampaignNegativeKeywordAccessError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsDraftCampaignNegativeKeywordAccessError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftCampaignNegativeKeywordAccessError and assigns it to the Errors field.
func (o *SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) SetErrors(v []SponsoredProductsDraftCampaignNegativeKeywordAccessError) {
	o.Errors = v
}

func (o SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent struct {
	value *SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) Get() *SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) Set(val *SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent(val *SponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) *NullableSponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent {
	return &NullableSponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordAccessExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
