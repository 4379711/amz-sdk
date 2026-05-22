package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsUnsupportedMediaTypeExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUnsupportedMediaTypeExceptionResponseContent{}

// SponsoredProductsUnsupportedMediaTypeExceptionResponseContent struct for SponsoredProductsUnsupportedMediaTypeExceptionResponseContent
type SponsoredProductsUnsupportedMediaTypeExceptionResponseContent struct {
	Code SponsoredProductsUnsupportedMediaTypeErrorCode `json:"code"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsUnsupportedMediaTypeExceptionResponseContent SponsoredProductsUnsupportedMediaTypeExceptionResponseContent

// NewSponsoredProductsUnsupportedMediaTypeExceptionResponseContent instantiates a new SponsoredProductsUnsupportedMediaTypeExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUnsupportedMediaTypeExceptionResponseContent(code SponsoredProductsUnsupportedMediaTypeErrorCode, message string) *SponsoredProductsUnsupportedMediaTypeExceptionResponseContent {
	this := SponsoredProductsUnsupportedMediaTypeExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSponsoredProductsUnsupportedMediaTypeExceptionResponseContentWithDefaults instantiates a new SponsoredProductsUnsupportedMediaTypeExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUnsupportedMediaTypeExceptionResponseContentWithDefaults() *SponsoredProductsUnsupportedMediaTypeExceptionResponseContent {
	this := SponsoredProductsUnsupportedMediaTypeExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SponsoredProductsUnsupportedMediaTypeExceptionResponseContent) GetCode() SponsoredProductsUnsupportedMediaTypeErrorCode {
	if o == nil {
		var ret SponsoredProductsUnsupportedMediaTypeErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUnsupportedMediaTypeExceptionResponseContent) GetCodeOk() (*SponsoredProductsUnsupportedMediaTypeErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SponsoredProductsUnsupportedMediaTypeExceptionResponseContent) SetCode(v SponsoredProductsUnsupportedMediaTypeErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsUnsupportedMediaTypeExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUnsupportedMediaTypeExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsUnsupportedMediaTypeExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsUnsupportedMediaTypeExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsUnsupportedMediaTypeExceptionResponseContent struct {
	value *SponsoredProductsUnsupportedMediaTypeExceptionResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUnsupportedMediaTypeExceptionResponseContent) Get() *SponsoredProductsUnsupportedMediaTypeExceptionResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUnsupportedMediaTypeExceptionResponseContent) Set(val *SponsoredProductsUnsupportedMediaTypeExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUnsupportedMediaTypeExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUnsupportedMediaTypeExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUnsupportedMediaTypeExceptionResponseContent(val *SponsoredProductsUnsupportedMediaTypeExceptionResponseContent) *NullableSponsoredProductsUnsupportedMediaTypeExceptionResponseContent {
	return &NullableSponsoredProductsUnsupportedMediaTypeExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUnsupportedMediaTypeExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUnsupportedMediaTypeExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
