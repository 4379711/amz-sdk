package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent{}

// SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent Exception resulting in accessing campaign management entities
type SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                                `json:"message"`
	Errors  []SponsoredProductsCampaignNegativeTargetsAccessError `json:"errors,omitempty"`
}

type _SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent

// NewSponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent instantiates a new SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent {
	this := SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContentWithDefaults instantiates a new SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContentWithDefaults() *SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent {
	this := SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) GetErrors() []SponsoredProductsCampaignNegativeTargetsAccessError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsCampaignNegativeTargetsAccessError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsCampaignNegativeTargetsAccessError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsCampaignNegativeTargetsAccessError and assigns it to the Errors field.
func (o *SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) SetErrors(v []SponsoredProductsCampaignNegativeTargetsAccessError) {
	o.Errors = v
}

func (o SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent struct {
	value *SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) Get() *SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) Set(val *SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent(val *SponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) *NullableSponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent {
	return &NullableSponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignNegativeTargetsAccessExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
