package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingUnauthorizedExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingUnauthorizedExceptionResponseContent{}

// SBTargetingUnauthorizedExceptionResponseContent Returns information about an UnauthorizedException.
type SBTargetingUnauthorizedExceptionResponseContent struct {
	// The HTTP status code of the response.
	Code string `json:"code"`
	// A human-readable description of the response.
	Details string `json:"details"`
}

type _SBTargetingUnauthorizedExceptionResponseContent SBTargetingUnauthorizedExceptionResponseContent

// NewSBTargetingUnauthorizedExceptionResponseContent instantiates a new SBTargetingUnauthorizedExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingUnauthorizedExceptionResponseContent(code string, details string) *SBTargetingUnauthorizedExceptionResponseContent {
	this := SBTargetingUnauthorizedExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewSBTargetingUnauthorizedExceptionResponseContentWithDefaults instantiates a new SBTargetingUnauthorizedExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingUnauthorizedExceptionResponseContentWithDefaults() *SBTargetingUnauthorizedExceptionResponseContent {
	this := SBTargetingUnauthorizedExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SBTargetingUnauthorizedExceptionResponseContent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SBTargetingUnauthorizedExceptionResponseContent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SBTargetingUnauthorizedExceptionResponseContent) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *SBTargetingUnauthorizedExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SBTargetingUnauthorizedExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SBTargetingUnauthorizedExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o SBTargetingUnauthorizedExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableSBTargetingUnauthorizedExceptionResponseContent struct {
	value *SBTargetingUnauthorizedExceptionResponseContent
	isSet bool
}

func (v NullableSBTargetingUnauthorizedExceptionResponseContent) Get() *SBTargetingUnauthorizedExceptionResponseContent {
	return v.value
}

func (v *NullableSBTargetingUnauthorizedExceptionResponseContent) Set(val *SBTargetingUnauthorizedExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingUnauthorizedExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingUnauthorizedExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingUnauthorizedExceptionResponseContent(val *SBTargetingUnauthorizedExceptionResponseContent) *NullableSBTargetingUnauthorizedExceptionResponseContent {
	return &NullableSBTargetingUnauthorizedExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSBTargetingUnauthorizedExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingUnauthorizedExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
