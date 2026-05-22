package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignAccessExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignAccessExceptionResponseContent{}

// SponsoredProductsCampaignAccessExceptionResponseContent Exception resulting in accessing campaign management entities
type SponsoredProductsCampaignAccessExceptionResponseContent struct {
	Code SponsoredProductsInvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                                 `json:"message"`
	Errors  []SponsoredProductsCampaignAccessError `json:"errors,omitempty"`
}

type _SponsoredProductsCampaignAccessExceptionResponseContent SponsoredProductsCampaignAccessExceptionResponseContent

// NewSponsoredProductsCampaignAccessExceptionResponseContent instantiates a new SponsoredProductsCampaignAccessExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignAccessExceptionResponseContent(code SponsoredProductsInvalidArgumentErrorCode, message string) *SponsoredProductsCampaignAccessExceptionResponseContent {
	this := SponsoredProductsCampaignAccessExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsCampaignAccessExceptionResponseContentWithDefaults instantiates a new SponsoredProductsCampaignAccessExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignAccessExceptionResponseContentWithDefaults() *SponsoredProductsCampaignAccessExceptionResponseContent {
	this := SponsoredProductsCampaignAccessExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsCampaignAccessExceptionResponseContent) GetCode() SponsoredProductsInvalidArgumentErrorCode {
	if o == nil {
		var ret SponsoredProductsInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignAccessExceptionResponseContent) GetCodeOk() (*SponsoredProductsInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsCampaignAccessExceptionResponseContent) SetCode(v SponsoredProductsInvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsCampaignAccessExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignAccessExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsCampaignAccessExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignAccessExceptionResponseContent) GetErrors() []SponsoredProductsCampaignAccessError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsCampaignAccessError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignAccessExceptionResponseContent) GetErrorsOk() ([]SponsoredProductsCampaignAccessError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignAccessExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsCampaignAccessError and assigns it to the Errors field.
func (o *SponsoredProductsCampaignAccessExceptionResponseContent) SetErrors(v []SponsoredProductsCampaignAccessError) {
	o.Errors = v
}

func (o SponsoredProductsCampaignAccessExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignAccessExceptionResponseContent struct {
	value *SponsoredProductsCampaignAccessExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCampaignAccessExceptionResponseContent) Get() *SponsoredProductsCampaignAccessExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCampaignAccessExceptionResponseContent) Set(val *SponsoredProductsCampaignAccessExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignAccessExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignAccessExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignAccessExceptionResponseContent(val *SponsoredProductsCampaignAccessExceptionResponseContent) *NullableSponsoredProductsCampaignAccessExceptionResponseContent {
	return &NullableSponsoredProductsCampaignAccessExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignAccessExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignAccessExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
