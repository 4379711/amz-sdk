package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the LandingPageInternalServerExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LandingPageInternalServerExceptionResponseContent{}

// LandingPageInternalServerExceptionResponseContent struct for LandingPageInternalServerExceptionResponseContent
type LandingPageInternalServerExceptionResponseContent struct {
	Code LandingPageInternalErrorCode `json:"code"`
	// A human-readable description of the code field.
	Details string `json:"details"`
}

type _LandingPageInternalServerExceptionResponseContent LandingPageInternalServerExceptionResponseContent

// NewLandingPageInternalServerExceptionResponseContent instantiates a new LandingPageInternalServerExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLandingPageInternalServerExceptionResponseContent(code LandingPageInternalErrorCode, details string) *LandingPageInternalServerExceptionResponseContent {
	this := LandingPageInternalServerExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewLandingPageInternalServerExceptionResponseContentWithDefaults instantiates a new LandingPageInternalServerExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLandingPageInternalServerExceptionResponseContentWithDefaults() *LandingPageInternalServerExceptionResponseContent {
	this := LandingPageInternalServerExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *LandingPageInternalServerExceptionResponseContent) GetCode() LandingPageInternalErrorCode {
	if o == nil {
		var ret LandingPageInternalErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *LandingPageInternalServerExceptionResponseContent) GetCodeOk() (*LandingPageInternalErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *LandingPageInternalServerExceptionResponseContent) SetCode(v LandingPageInternalErrorCode) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *LandingPageInternalServerExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *LandingPageInternalServerExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *LandingPageInternalServerExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o LandingPageInternalServerExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableLandingPageInternalServerExceptionResponseContent struct {
	value *LandingPageInternalServerExceptionResponseContent
	isSet bool
}

func (v NullableLandingPageInternalServerExceptionResponseContent) Get() *LandingPageInternalServerExceptionResponseContent {
	return v.value
}

func (v *NullableLandingPageInternalServerExceptionResponseContent) Set(val *LandingPageInternalServerExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableLandingPageInternalServerExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableLandingPageInternalServerExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLandingPageInternalServerExceptionResponseContent(val *LandingPageInternalServerExceptionResponseContent) *NullableLandingPageInternalServerExceptionResponseContent {
	return &NullableLandingPageInternalServerExceptionResponseContent{value: val, isSet: true}
}

func (v NullableLandingPageInternalServerExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLandingPageInternalServerExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
