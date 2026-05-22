package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingInternalServerExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingInternalServerExceptionResponseContent{}

// SBForecastingInternalServerExceptionResponseContent Returns information about a InternalServerException.
type SBForecastingInternalServerExceptionResponseContent struct {
	// The HTTP status code of the response.
	Code string `json:"code"`
	// A human-readable description of the response.
	Details string `json:"details"`
}

type _SBForecastingInternalServerExceptionResponseContent SBForecastingInternalServerExceptionResponseContent

// NewSBForecastingInternalServerExceptionResponseContent instantiates a new SBForecastingInternalServerExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingInternalServerExceptionResponseContent(code string, details string) *SBForecastingInternalServerExceptionResponseContent {
	this := SBForecastingInternalServerExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewSBForecastingInternalServerExceptionResponseContentWithDefaults instantiates a new SBForecastingInternalServerExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingInternalServerExceptionResponseContentWithDefaults() *SBForecastingInternalServerExceptionResponseContent {
	this := SBForecastingInternalServerExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SBForecastingInternalServerExceptionResponseContent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SBForecastingInternalServerExceptionResponseContent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SBForecastingInternalServerExceptionResponseContent) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *SBForecastingInternalServerExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SBForecastingInternalServerExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SBForecastingInternalServerExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o SBForecastingInternalServerExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableSBForecastingInternalServerExceptionResponseContent struct {
	value *SBForecastingInternalServerExceptionResponseContent
	isSet bool
}

func (v NullableSBForecastingInternalServerExceptionResponseContent) Get() *SBForecastingInternalServerExceptionResponseContent {
	return v.value
}

func (v *NullableSBForecastingInternalServerExceptionResponseContent) Set(val *SBForecastingInternalServerExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingInternalServerExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingInternalServerExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingInternalServerExceptionResponseContent(val *SBForecastingInternalServerExceptionResponseContent) *NullableSBForecastingInternalServerExceptionResponseContent {
	return &NullableSBForecastingInternalServerExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSBForecastingInternalServerExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingInternalServerExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
