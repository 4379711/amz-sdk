package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingBadRequestExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingBadRequestExceptionResponseContent{}

// SBTargetingBadRequestExceptionResponseContent Returns information about a BadRequestException.
type SBTargetingBadRequestExceptionResponseContent struct {
	// The HTTP status code of the response.
	Code string `json:"code"`
	// A human-readable description of the response.
	Details string `json:"details"`
}

type _SBTargetingBadRequestExceptionResponseContent SBTargetingBadRequestExceptionResponseContent

// NewSBTargetingBadRequestExceptionResponseContent instantiates a new SBTargetingBadRequestExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingBadRequestExceptionResponseContent(code string, details string) *SBTargetingBadRequestExceptionResponseContent {
	this := SBTargetingBadRequestExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewSBTargetingBadRequestExceptionResponseContentWithDefaults instantiates a new SBTargetingBadRequestExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingBadRequestExceptionResponseContentWithDefaults() *SBTargetingBadRequestExceptionResponseContent {
	this := SBTargetingBadRequestExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SBTargetingBadRequestExceptionResponseContent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SBTargetingBadRequestExceptionResponseContent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SBTargetingBadRequestExceptionResponseContent) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *SBTargetingBadRequestExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SBTargetingBadRequestExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SBTargetingBadRequestExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o SBTargetingBadRequestExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableSBTargetingBadRequestExceptionResponseContent struct {
	value *SBTargetingBadRequestExceptionResponseContent
	isSet bool
}

func (v NullableSBTargetingBadRequestExceptionResponseContent) Get() *SBTargetingBadRequestExceptionResponseContent {
	return v.value
}

func (v *NullableSBTargetingBadRequestExceptionResponseContent) Set(val *SBTargetingBadRequestExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingBadRequestExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingBadRequestExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingBadRequestExceptionResponseContent(val *SBTargetingBadRequestExceptionResponseContent) *NullableSBTargetingBadRequestExceptionResponseContent {
	return &NullableSBTargetingBadRequestExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSBTargetingBadRequestExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingBadRequestExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
