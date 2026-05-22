package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent{}

// SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent Exception resulting in accessing campaign management entities
type SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                                `json:"message"`
	Errors  []SponsoredProductsCampaignNegativeKeywordAccessError `json:"errors,omitempty"`
}

type _SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent

// NewSponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent instantiates a new SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent {
	this := SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContentWithDefaults instantiates a new SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContentWithDefaults() *SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent {
	this := SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) GetErrors() []SponsoredProductsCampaignNegativeKeywordAccessError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsCampaignNegativeKeywordAccessError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsCampaignNegativeKeywordAccessError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsCampaignNegativeKeywordAccessError and assigns it to the Errors field.
func (o *SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) SetErrors(v []SponsoredProductsCampaignNegativeKeywordAccessError) {
	o.Errors = v
}

func (o SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent struct {
	value *SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) Get() *SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) Set(val *SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent(val *SponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) *NullableSponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent {
	return &NullableSponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignNegativeKeywordAccessExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
