package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingInternalServerExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingInternalServerExceptionResponseContent{}

// SBTargetingInternalServerExceptionResponseContent Returns information about an InternalServerException.
type SBTargetingInternalServerExceptionResponseContent struct {
	// The HTTP status code of the response.
	Code string `json:"code"`
	// A human-readable description of the response.
	Details string `json:"details"`
}

type _SBTargetingInternalServerExceptionResponseContent SBTargetingInternalServerExceptionResponseContent

// NewSBTargetingInternalServerExceptionResponseContent instantiates a new SBTargetingInternalServerExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingInternalServerExceptionResponseContent(code string, details string) *SBTargetingInternalServerExceptionResponseContent {
	this := SBTargetingInternalServerExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewSBTargetingInternalServerExceptionResponseContentWithDefaults instantiates a new SBTargetingInternalServerExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingInternalServerExceptionResponseContentWithDefaults() *SBTargetingInternalServerExceptionResponseContent {
	this := SBTargetingInternalServerExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SBTargetingInternalServerExceptionResponseContent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SBTargetingInternalServerExceptionResponseContent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SBTargetingInternalServerExceptionResponseContent) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *SBTargetingInternalServerExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SBTargetingInternalServerExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SBTargetingInternalServerExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o SBTargetingInternalServerExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableSBTargetingInternalServerExceptionResponseContent struct {
	value *SBTargetingInternalServerExceptionResponseContent
	isSet bool
}

func (v NullableSBTargetingInternalServerExceptionResponseContent) Get() *SBTargetingInternalServerExceptionResponseContent {
	return v.value
}

func (v *NullableSBTargetingInternalServerExceptionResponseContent) Set(val *SBTargetingInternalServerExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingInternalServerExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingInternalServerExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingInternalServerExceptionResponseContent(val *SBTargetingInternalServerExceptionResponseContent) *NullableSBTargetingInternalServerExceptionResponseContent {
	return &NullableSBTargetingInternalServerExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSBTargetingInternalServerExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingInternalServerExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
