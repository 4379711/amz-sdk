package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBInsightsBadRequestExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBInsightsBadRequestExceptionResponseContent{}

// SBInsightsBadRequestExceptionResponseContent Returns information about a BadRequestException.
type SBInsightsBadRequestExceptionResponseContent struct {
	// The HTTP status code of the response.
	Code string `json:"code"`
	// A human-readable description of the response.
	Details string `json:"details"`
}

type _SBInsightsBadRequestExceptionResponseContent SBInsightsBadRequestExceptionResponseContent

// NewSBInsightsBadRequestExceptionResponseContent instantiates a new SBInsightsBadRequestExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBInsightsBadRequestExceptionResponseContent(code string, details string) *SBInsightsBadRequestExceptionResponseContent {
	this := SBInsightsBadRequestExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewSBInsightsBadRequestExceptionResponseContentWithDefaults instantiates a new SBInsightsBadRequestExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBInsightsBadRequestExceptionResponseContentWithDefaults() *SBInsightsBadRequestExceptionResponseContent {
	this := SBInsightsBadRequestExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SBInsightsBadRequestExceptionResponseContent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SBInsightsBadRequestExceptionResponseContent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SBInsightsBadRequestExceptionResponseContent) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *SBInsightsBadRequestExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SBInsightsBadRequestExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SBInsightsBadRequestExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o SBInsightsBadRequestExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableSBInsightsBadRequestExceptionResponseContent struct {
	value *SBInsightsBadRequestExceptionResponseContent
	isSet bool
}

func (v NullableSBInsightsBadRequestExceptionResponseContent) Get() *SBInsightsBadRequestExceptionResponseContent {
	return v.value
}

func (v *NullableSBInsightsBadRequestExceptionResponseContent) Set(val *SBInsightsBadRequestExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBInsightsBadRequestExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBInsightsBadRequestExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBInsightsBadRequestExceptionResponseContent(val *SBInsightsBadRequestExceptionResponseContent) *NullableSBInsightsBadRequestExceptionResponseContent {
	return &NullableSBInsightsBadRequestExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSBInsightsBadRequestExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBInsightsBadRequestExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
