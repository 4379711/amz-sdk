package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingThrottlingExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingThrottlingExceptionResponseContent{}

// SBTargetingThrottlingExceptionResponseContent Returns information about a ThrottlingException.
type SBTargetingThrottlingExceptionResponseContent struct {
	// The HTTP status code of the response.
	Code string `json:"code"`
	// A human-readable description of the response.
	Details string `json:"details"`
}

type _SBTargetingThrottlingExceptionResponseContent SBTargetingThrottlingExceptionResponseContent

// NewSBTargetingThrottlingExceptionResponseContent instantiates a new SBTargetingThrottlingExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingThrottlingExceptionResponseContent(code string, details string) *SBTargetingThrottlingExceptionResponseContent {
	this := SBTargetingThrottlingExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewSBTargetingThrottlingExceptionResponseContentWithDefaults instantiates a new SBTargetingThrottlingExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingThrottlingExceptionResponseContentWithDefaults() *SBTargetingThrottlingExceptionResponseContent {
	this := SBTargetingThrottlingExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SBTargetingThrottlingExceptionResponseContent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SBTargetingThrottlingExceptionResponseContent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SBTargetingThrottlingExceptionResponseContent) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *SBTargetingThrottlingExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SBTargetingThrottlingExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SBTargetingThrottlingExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o SBTargetingThrottlingExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableSBTargetingThrottlingExceptionResponseContent struct {
	value *SBTargetingThrottlingExceptionResponseContent
	isSet bool
}

func (v NullableSBTargetingThrottlingExceptionResponseContent) Get() *SBTargetingThrottlingExceptionResponseContent {
	return v.value
}

func (v *NullableSBTargetingThrottlingExceptionResponseContent) Set(val *SBTargetingThrottlingExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingThrottlingExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingThrottlingExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingThrottlingExceptionResponseContent(val *SBTargetingThrottlingExceptionResponseContent) *NullableSBTargetingThrottlingExceptionResponseContent {
	return &NullableSBTargetingThrottlingExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSBTargetingThrottlingExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingThrottlingExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
