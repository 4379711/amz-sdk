package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignMutationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignMutationExceptionResponseContent{}

// SponsoredProductsDraftCampaignMutationExceptionResponseContent Exception resulting in mutating draft management entities
type SponsoredProductsDraftCampaignMutationExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                        `json:"message"`
	Errors  []SponsoredProductsDraftCampaignMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftCampaignMutationExceptionResponseContent SponsoredProductsDraftCampaignMutationExceptionResponseContent

// NewSponsoredProductsDraftCampaignMutationExceptionResponseContent instantiates a new SponsoredProductsDraftCampaignMutationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignMutationExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsDraftCampaignMutationExceptionResponseContent {
	this := SponsoredProductsDraftCampaignMutationExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsDraftCampaignMutationExceptionResponseContentWithDefaults instantiates a new SponsoredProductsDraftCampaignMutationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignMutationExceptionResponseContentWithDefaults() *SponsoredProductsDraftCampaignMutationExceptionResponseContent {
	this := SponsoredProductsDraftCampaignMutationExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsDraftCampaignMutationExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignMutationExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsDraftCampaignMutationExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsDraftCampaignMutationExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignMutationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsDraftCampaignMutationExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignMutationExceptionResponseContent) GetErrors() []SponsoredProductsDraftCampaignMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftCampaignMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignMutationExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsDraftCampaignMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignMutationExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftCampaignMutationError and assigns it to the Errors field.
func (o *SponsoredProductsDraftCampaignMutationExceptionResponseContent) SetErrors(v []SponsoredProductsDraftCampaignMutationError) {
	o.Errors = v
}

func (o SponsoredProductsDraftCampaignMutationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftCampaignMutationExceptionResponseContent struct {
	value *SponsoredProductsDraftCampaignMutationExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignMutationExceptionResponseContent) Get() *SponsoredProductsDraftCampaignMutationExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignMutationExceptionResponseContent) Set(val *SponsoredProductsDraftCampaignMutationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignMutationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignMutationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignMutationExceptionResponseContent(val *SponsoredProductsDraftCampaignMutationExceptionResponseContent) *NullableSponsoredProductsDraftCampaignMutationExceptionResponseContent {
	return &NullableSponsoredProductsDraftCampaignMutationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignMutationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignMutationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
