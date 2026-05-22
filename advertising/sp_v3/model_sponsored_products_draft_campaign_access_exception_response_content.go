package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignAccessExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignAccessExceptionResponseContent{}

// SponsoredProductsDraftCampaignAccessExceptionResponseContent Exception resulting in accessing draft management entities
type SponsoredProductsDraftCampaignAccessExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                      `json:"message"`
	Errors  []SponsoredProductsDraftCampaignAccessError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftCampaignAccessExceptionResponseContent SponsoredProductsDraftCampaignAccessExceptionResponseContent

// NewSponsoredProductsDraftCampaignAccessExceptionResponseContent instantiates a new SponsoredProductsDraftCampaignAccessExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignAccessExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsDraftCampaignAccessExceptionResponseContent {
	this := SponsoredProductsDraftCampaignAccessExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsDraftCampaignAccessExceptionResponseContentWithDefaults instantiates a new SponsoredProductsDraftCampaignAccessExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignAccessExceptionResponseContentWithDefaults() *SponsoredProductsDraftCampaignAccessExceptionResponseContent {
	this := SponsoredProductsDraftCampaignAccessExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsDraftCampaignAccessExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignAccessExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsDraftCampaignAccessExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsDraftCampaignAccessExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignAccessExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsDraftCampaignAccessExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignAccessExceptionResponseContent) GetErrors() []SponsoredProductsDraftCampaignAccessError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftCampaignAccessError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignAccessExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsDraftCampaignAccessError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignAccessExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftCampaignAccessError and assigns it to the Errors field.
func (o *SponsoredProductsDraftCampaignAccessExceptionResponseContent) SetErrors(v []SponsoredProductsDraftCampaignAccessError) {
	o.Errors = v
}

func (o SponsoredProductsDraftCampaignAccessExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftCampaignAccessExceptionResponseContent struct {
	value *SponsoredProductsDraftCampaignAccessExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignAccessExceptionResponseContent) Get() *SponsoredProductsDraftCampaignAccessExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignAccessExceptionResponseContent) Set(val *SponsoredProductsDraftCampaignAccessExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignAccessExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignAccessExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignAccessExceptionResponseContent(val *SponsoredProductsDraftCampaignAccessExceptionResponseContent) *NullableSponsoredProductsDraftCampaignAccessExceptionResponseContent {
	return &NullableSponsoredProductsDraftCampaignAccessExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignAccessExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignAccessExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
