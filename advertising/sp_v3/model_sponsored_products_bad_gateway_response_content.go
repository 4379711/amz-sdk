package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBadGatewayResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBadGatewayResponseContent{}

// SponsoredProductsBadGatewayResponseContent struct for SponsoredProductsBadGatewayResponseContent
type SponsoredProductsBadGatewayResponseContent struct {
	Code           SponsoredProductsErrorCode `json:"code"`
	Message        string                     `json:"message"`
	HttpStatusCode *string                    `json:"httpStatusCode,omitempty"`
}

type _SponsoredProductsBadGatewayResponseContent SponsoredProductsBadGatewayResponseContent

// NewSponsoredProductsBadGatewayResponseContent instantiates a new SponsoredProductsBadGatewayResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBadGatewayResponseContent(code SponsoredProductsErrorCode, message string) *SponsoredProductsBadGatewayResponseContent {
	this := SponsoredProductsBadGatewayResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsBadGatewayResponseContentWithDefaults instantiates a new SponsoredProductsBadGatewayResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBadGatewayResponseContentWithDefaults() *SponsoredProductsBadGatewayResponseContent {
	this := SponsoredProductsBadGatewayResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsBadGatewayResponseContent) GetCode() SponsoredProductsErrorCode {
	if o == nil {
		var ret SponsoredProductsErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBadGatewayResponseContent) GetCodeOk() (*SponsoredProductsErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsBadGatewayResponseContent) SetCode(v SponsoredProductsErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsBadGatewayResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBadGatewayResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsBadGatewayResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetHttpStatusCode returns the HttpStatusCode field value if set, zero value otherwise.
func (o *SponsoredProductsBadGatewayResponseContent) GetHttpStatusCode() string {
	if o == nil || IsNil(o.HttpStatusCode) {
		var ret string
		return ret
	}
	return *o.HttpStatusCode
}

// GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBadGatewayResponseContent) GetHttpStatusCodeOk() (*string, bool) {
	if o == nil || IsNil(o.HttpStatusCode) {
		return nil, false
	}
	return o.HttpStatusCode, true
}

// HasHttpStatusCode returns a boolean if a field has been set.
func (o *SponsoredProductsBadGatewayResponseContent) HasHttpStatusCode() bool {
	if o != nil && !IsNil(o.HttpStatusCode) {
		return true
	}

	return false
}

// SetHttpStatusCode gets a reference to the given string and assigns it to the HttpStatusCode field.
func (o *SponsoredProductsBadGatewayResponseContent) SetHttpStatusCode(v string) {
	o.HttpStatusCode = &v
}

func (o SponsoredProductsBadGatewayResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.HttpStatusCode) {
		toSerialize["httpStatusCode"] = o.HttpStatusCode
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBadGatewayResponseContent struct {
	value *SponsoredProductsBadGatewayResponseContent
	isSet bool
}

func (v NullableSponsoredProductsBadGatewayResponseContent) Get() *SponsoredProductsBadGatewayResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsBadGatewayResponseContent) Set(val *SponsoredProductsBadGatewayResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBadGatewayResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBadGatewayResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBadGatewayResponseContent(val *SponsoredProductsBadGatewayResponseContent) *NullableSponsoredProductsBadGatewayResponseContent {
	return &NullableSponsoredProductsBadGatewayResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsBadGatewayResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBadGatewayResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
