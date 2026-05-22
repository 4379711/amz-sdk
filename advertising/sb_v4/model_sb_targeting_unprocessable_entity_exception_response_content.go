package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingUnprocessableEntityExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingUnprocessableEntityExceptionResponseContent{}

// SBTargetingUnprocessableEntityExceptionResponseContent Returns information about an UnprocessableEntityException.
type SBTargetingUnprocessableEntityExceptionResponseContent struct {
	// The HTTP status code of the response.
	Code string `json:"code"`
	// A human-readable description of the response.
	Details string `json:"details"`
}

type _SBTargetingUnprocessableEntityExceptionResponseContent SBTargetingUnprocessableEntityExceptionResponseContent

// NewSBTargetingUnprocessableEntityExceptionResponseContent instantiates a new SBTargetingUnprocessableEntityExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingUnprocessableEntityExceptionResponseContent(code string, details string) *SBTargetingUnprocessableEntityExceptionResponseContent {
	this := SBTargetingUnprocessableEntityExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewSBTargetingUnprocessableEntityExceptionResponseContentWithDefaults instantiates a new SBTargetingUnprocessableEntityExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingUnprocessableEntityExceptionResponseContentWithDefaults() *SBTargetingUnprocessableEntityExceptionResponseContent {
	this := SBTargetingUnprocessableEntityExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SBTargetingUnprocessableEntityExceptionResponseContent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SBTargetingUnprocessableEntityExceptionResponseContent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SBTargetingUnprocessableEntityExceptionResponseContent) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *SBTargetingUnprocessableEntityExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SBTargetingUnprocessableEntityExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SBTargetingUnprocessableEntityExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o SBTargetingUnprocessableEntityExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableSBTargetingUnprocessableEntityExceptionResponseContent struct {
	value *SBTargetingUnprocessableEntityExceptionResponseContent
	isSet bool
}

func (v NullableSBTargetingUnprocessableEntityExceptionResponseContent) Get() *SBTargetingUnprocessableEntityExceptionResponseContent {
	return v.value
}

func (v *NullableSBTargetingUnprocessableEntityExceptionResponseContent) Set(val *SBTargetingUnprocessableEntityExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingUnprocessableEntityExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingUnprocessableEntityExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingUnprocessableEntityExceptionResponseContent(val *SBTargetingUnprocessableEntityExceptionResponseContent) *NullableSBTargetingUnprocessableEntityExceptionResponseContent {
	return &NullableSBTargetingUnprocessableEntityExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSBTargetingUnprocessableEntityExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingUnprocessableEntityExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
