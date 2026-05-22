package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent{}

// SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent Exception resulting in mutating campaign management entities
type SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                                  `json:"message"`
	Errors  []SponsoredProductsCampaignNegativeTargetsMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent

// NewSponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent instantiates a new SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent {
	this := SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContentWithDefaults instantiates a new SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContentWithDefaults() *SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent {
	this := SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) GetErrors() []SponsoredProductsCampaignNegativeTargetsMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsCampaignNegativeTargetsMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsCampaignNegativeTargetsMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsCampaignNegativeTargetsMutationError and assigns it to the Errors field.
func (o *SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) SetErrors(v []SponsoredProductsCampaignNegativeTargetsMutationError) {
	o.Errors = v
}

func (o SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent struct {
	value *SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) Get() *SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) Set(val *SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent(val *SponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) *NullableSponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent {
	return &NullableSponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignNegativeTargetsMutationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
