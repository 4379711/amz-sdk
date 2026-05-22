package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the LandingPageThrottlingExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LandingPageThrottlingExceptionResponseContent{}

// LandingPageThrottlingExceptionResponseContent struct for LandingPageThrottlingExceptionResponseContent
type LandingPageThrottlingExceptionResponseContent struct {
	Code LandingPageThrottledErrorCode `json:"code"`
	// A human-readable description of the code field.
	Details string `json:"details"`
}

type _LandingPageThrottlingExceptionResponseContent LandingPageThrottlingExceptionResponseContent

// NewLandingPageThrottlingExceptionResponseContent instantiates a new LandingPageThrottlingExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLandingPageThrottlingExceptionResponseContent(code LandingPageThrottledErrorCode, details string) *LandingPageThrottlingExceptionResponseContent {
	this := LandingPageThrottlingExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewLandingPageThrottlingExceptionResponseContentWithDefaults instantiates a new LandingPageThrottlingExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLandingPageThrottlingExceptionResponseContentWithDefaults() *LandingPageThrottlingExceptionResponseContent {
	this := LandingPageThrottlingExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *LandingPageThrottlingExceptionResponseContent) GetCode() LandingPageThrottledErrorCode {
	if o == nil {
		var ret LandingPageThrottledErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *LandingPageThrottlingExceptionResponseContent) GetCodeOk() (*LandingPageThrottledErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *LandingPageThrottlingExceptionResponseContent) SetCode(v LandingPageThrottledErrorCode) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *LandingPageThrottlingExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *LandingPageThrottlingExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *LandingPageThrottlingExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o LandingPageThrottlingExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableLandingPageThrottlingExceptionResponseContent struct {
	value *LandingPageThrottlingExceptionResponseContent
	isSet bool
}

func (v NullableLandingPageThrottlingExceptionResponseContent) Get() *LandingPageThrottlingExceptionResponseContent {
	return v.value
}

func (v *NullableLandingPageThrottlingExceptionResponseContent) Set(val *LandingPageThrottlingExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableLandingPageThrottlingExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableLandingPageThrottlingExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLandingPageThrottlingExceptionResponseContent(val *LandingPageThrottlingExceptionResponseContent) *NullableLandingPageThrottlingExceptionResponseContent {
	return &NullableLandingPageThrottlingExceptionResponseContent{value: val, isSet: true}
}

func (v NullableLandingPageThrottlingExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLandingPageThrottlingExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
