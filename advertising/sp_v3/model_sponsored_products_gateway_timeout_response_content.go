package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGatewayTimeoutResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGatewayTimeoutResponseContent{}

// SponsoredProductsGatewayTimeoutResponseContent struct for SponsoredProductsGatewayTimeoutResponseContent
type SponsoredProductsGatewayTimeoutResponseContent struct {
	Code           SponsoredProductsErrorCode `json:"code"`
	Message        string                     `json:"message"`
	HttpStatusCode *string                    `json:"httpStatusCode,omitempty"`
}

type _SponsoredProductsGatewayTimeoutResponseContent SponsoredProductsGatewayTimeoutResponseContent

// NewSponsoredProductsGatewayTimeoutResponseContent instantiates a new SponsoredProductsGatewayTimeoutResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGatewayTimeoutResponseContent(code SponsoredProductsErrorCode, message string) *SponsoredProductsGatewayTimeoutResponseContent {
	this := SponsoredProductsGatewayTimeoutResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsGatewayTimeoutResponseContentWithDefaults instantiates a new SponsoredProductsGatewayTimeoutResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGatewayTimeoutResponseContentWithDefaults() *SponsoredProductsGatewayTimeoutResponseContent {
	this := SponsoredProductsGatewayTimeoutResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsGatewayTimeoutResponseContent) GetCode() SponsoredProductsErrorCode {
	if o == nil {
		var ret SponsoredProductsErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGatewayTimeoutResponseContent) GetCodeOk() (*SponsoredProductsErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsGatewayTimeoutResponseContent) SetCode(v SponsoredProductsErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsGatewayTimeoutResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGatewayTimeoutResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsGatewayTimeoutResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetHttpStatusCode returns the HttpStatusCode field value if set, zero value otherwise.
func (o *SponsoredProductsGatewayTimeoutResponseContent) GetHttpStatusCode() string {
	if o == nil || IsNil(o.HttpStatusCode) {
		var ret string
		return ret
	}
	return *o.HttpStatusCode
}

// GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGatewayTimeoutResponseContent) GetHttpStatusCodeOk() (*string, bool) {
	if o == nil || IsNil(o.HttpStatusCode) {
		return nil, false
	}
	return o.HttpStatusCode, true
}

// HasHttpStatusCode returns a boolean if a field has been set.
func (o *SponsoredProductsGatewayTimeoutResponseContent) HasHttpStatusCode() bool {
	if o != nil && !IsNil(o.HttpStatusCode) {
		return true
	}

	return false
}

// SetHttpStatusCode gets a reference to the given string and assigns it to the HttpStatusCode field.
func (o *SponsoredProductsGatewayTimeoutResponseContent) SetHttpStatusCode(v string) {
	o.HttpStatusCode = &v
}

func (o SponsoredProductsGatewayTimeoutResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.HttpStatusCode) {
		toSerialize["httpStatusCode"] = o.HttpStatusCode
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGatewayTimeoutResponseContent struct {
	value *SponsoredProductsGatewayTimeoutResponseContent
	isSet bool
}

func (v NullableSponsoredProductsGatewayTimeoutResponseContent) Get() *SponsoredProductsGatewayTimeoutResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsGatewayTimeoutResponseContent) Set(val *SponsoredProductsGatewayTimeoutResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGatewayTimeoutResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGatewayTimeoutResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGatewayTimeoutResponseContent(val *SponsoredProductsGatewayTimeoutResponseContent) *NullableSponsoredProductsGatewayTimeoutResponseContent {
	return &NullableSponsoredProductsGatewayTimeoutResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsGatewayTimeoutResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGatewayTimeoutResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
