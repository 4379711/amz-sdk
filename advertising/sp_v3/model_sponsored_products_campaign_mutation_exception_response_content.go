package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignMutationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignMutationExceptionResponseContent{}

// SponsoredProductsCampaignMutationExceptionResponseContent Exception resulting in mutating campaign management entities
type SponsoredProductsCampaignMutationExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                   `json:"message"`
	Errors  []SponsoredProductsCampaignMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsCampaignMutationExceptionResponseContent SponsoredProductsCampaignMutationExceptionResponseContent

// NewSponsoredProductsCampaignMutationExceptionResponseContent instantiates a new SponsoredProductsCampaignMutationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignMutationExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsCampaignMutationExceptionResponseContent {
	this := SponsoredProductsCampaignMutationExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsCampaignMutationExceptionResponseContentWithDefaults instantiates a new SponsoredProductsCampaignMutationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignMutationExceptionResponseContentWithDefaults() *SponsoredProductsCampaignMutationExceptionResponseContent {
	this := SponsoredProductsCampaignMutationExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsCampaignMutationExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsCampaignMutationExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsCampaignMutationExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsCampaignMutationExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationExceptionResponseContent) GetErrors() []SponsoredProductsCampaignMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsCampaignMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsCampaignMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsCampaignMutationError and assigns it to the Errors field.
func (o *SponsoredProductsCampaignMutationExceptionResponseContent) SetErrors(v []SponsoredProductsCampaignMutationError) {
	o.Errors = v
}

func (o SponsoredProductsCampaignMutationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignMutationExceptionResponseContent struct {
	value *SponsoredProductsCampaignMutationExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCampaignMutationExceptionResponseContent) Get() *SponsoredProductsCampaignMutationExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCampaignMutationExceptionResponseContent) Set(val *SponsoredProductsCampaignMutationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignMutationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignMutationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignMutationExceptionResponseContent(val *SponsoredProductsCampaignMutationExceptionResponseContent) *NullableSponsoredProductsCampaignMutationExceptionResponseContent {
	return &NullableSponsoredProductsCampaignMutationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignMutationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignMutationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
