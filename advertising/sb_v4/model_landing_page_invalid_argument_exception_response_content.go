package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the LandingPageInvalidArgumentExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LandingPageInvalidArgumentExceptionResponseContent{}

// LandingPageInvalidArgumentExceptionResponseContent struct for LandingPageInvalidArgumentExceptionResponseContent
type LandingPageInvalidArgumentExceptionResponseContent struct {
	Code LandingPageInvalidArgumentErrorCode `json:"code"`
	// A human-readable description of the code field.
	Details string `json:"details"`
}

type _LandingPageInvalidArgumentExceptionResponseContent LandingPageInvalidArgumentExceptionResponseContent

// NewLandingPageInvalidArgumentExceptionResponseContent instantiates a new LandingPageInvalidArgumentExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLandingPageInvalidArgumentExceptionResponseContent(code LandingPageInvalidArgumentErrorCode, details string) *LandingPageInvalidArgumentExceptionResponseContent {
	this := LandingPageInvalidArgumentExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewLandingPageInvalidArgumentExceptionResponseContentWithDefaults instantiates a new LandingPageInvalidArgumentExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLandingPageInvalidArgumentExceptionResponseContentWithDefaults() *LandingPageInvalidArgumentExceptionResponseContent {
	this := LandingPageInvalidArgumentExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *LandingPageInvalidArgumentExceptionResponseContent) GetCode() LandingPageInvalidArgumentErrorCode {
	if o == nil {
		var ret LandingPageInvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *LandingPageInvalidArgumentExceptionResponseContent) GetCodeOk() (*LandingPageInvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *LandingPageInvalidArgumentExceptionResponseContent) SetCode(v LandingPageInvalidArgumentErrorCode) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *LandingPageInvalidArgumentExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *LandingPageInvalidArgumentExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *LandingPageInvalidArgumentExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o LandingPageInvalidArgumentExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableLandingPageInvalidArgumentExceptionResponseContent struct {
	value *LandingPageInvalidArgumentExceptionResponseContent
	isSet bool
}

func (v NullableLandingPageInvalidArgumentExceptionResponseContent) Get() *LandingPageInvalidArgumentExceptionResponseContent {
	return v.value
}

func (v *NullableLandingPageInvalidArgumentExceptionResponseContent) Set(val *LandingPageInvalidArgumentExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableLandingPageInvalidArgumentExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableLandingPageInvalidArgumentExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLandingPageInvalidArgumentExceptionResponseContent(val *LandingPageInvalidArgumentExceptionResponseContent) *NullableLandingPageInvalidArgumentExceptionResponseContent {
	return &NullableLandingPageInvalidArgumentExceptionResponseContent{value: val, isSet: true}
}

func (v NullableLandingPageInvalidArgumentExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLandingPageInvalidArgumentExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
