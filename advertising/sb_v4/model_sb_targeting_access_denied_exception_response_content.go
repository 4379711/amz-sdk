package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingAccessDeniedExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingAccessDeniedExceptionResponseContent{}

// SBTargetingAccessDeniedExceptionResponseContent Returns information about an AccessDeniedException.
type SBTargetingAccessDeniedExceptionResponseContent struct {
	// The HTTP status code of the response.
	Code string `json:"code"`
	// A human-readable description of the response.
	Details string `json:"details"`
}

type _SBTargetingAccessDeniedExceptionResponseContent SBTargetingAccessDeniedExceptionResponseContent

// NewSBTargetingAccessDeniedExceptionResponseContent instantiates a new SBTargetingAccessDeniedExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingAccessDeniedExceptionResponseContent(code string, details string) *SBTargetingAccessDeniedExceptionResponseContent {
	this := SBTargetingAccessDeniedExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewSBTargetingAccessDeniedExceptionResponseContentWithDefaults instantiates a new SBTargetingAccessDeniedExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingAccessDeniedExceptionResponseContentWithDefaults() *SBTargetingAccessDeniedExceptionResponseContent {
	this := SBTargetingAccessDeniedExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SBTargetingAccessDeniedExceptionResponseContent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SBTargetingAccessDeniedExceptionResponseContent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SBTargetingAccessDeniedExceptionResponseContent) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *SBTargetingAccessDeniedExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SBTargetingAccessDeniedExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SBTargetingAccessDeniedExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o SBTargetingAccessDeniedExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableSBTargetingAccessDeniedExceptionResponseContent struct {
	value *SBTargetingAccessDeniedExceptionResponseContent
	isSet bool
}

func (v NullableSBTargetingAccessDeniedExceptionResponseContent) Get() *SBTargetingAccessDeniedExceptionResponseContent {
	return v.value
}

func (v *NullableSBTargetingAccessDeniedExceptionResponseContent) Set(val *SBTargetingAccessDeniedExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingAccessDeniedExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingAccessDeniedExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingAccessDeniedExceptionResponseContent(val *SBTargetingAccessDeniedExceptionResponseContent) *NullableSBTargetingAccessDeniedExceptionResponseContent {
	return &NullableSBTargetingAccessDeniedExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSBTargetingAccessDeniedExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingAccessDeniedExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
