package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingThrottlingExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingThrottlingExceptionResponseContent{}

// SBForecastingThrottlingExceptionResponseContent Returns information about a ThrottlingException.
type SBForecastingThrottlingExceptionResponseContent struct {
	// The HTTP status code of the response.
	Code string `json:"code"`
	// A human-readable description of the response.
	Details string `json:"details"`
}

type _SBForecastingThrottlingExceptionResponseContent SBForecastingThrottlingExceptionResponseContent

// NewSBForecastingThrottlingExceptionResponseContent instantiates a new SBForecastingThrottlingExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingThrottlingExceptionResponseContent(code string, details string) *SBForecastingThrottlingExceptionResponseContent {
	this := SBForecastingThrottlingExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewSBForecastingThrottlingExceptionResponseContentWithDefaults instantiates a new SBForecastingThrottlingExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingThrottlingExceptionResponseContentWithDefaults() *SBForecastingThrottlingExceptionResponseContent {
	this := SBForecastingThrottlingExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SBForecastingThrottlingExceptionResponseContent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SBForecastingThrottlingExceptionResponseContent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SBForecastingThrottlingExceptionResponseContent) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *SBForecastingThrottlingExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SBForecastingThrottlingExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SBForecastingThrottlingExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o SBForecastingThrottlingExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableSBForecastingThrottlingExceptionResponseContent struct {
	value *SBForecastingThrottlingExceptionResponseContent
	isSet bool
}

func (v NullableSBForecastingThrottlingExceptionResponseContent) Get() *SBForecastingThrottlingExceptionResponseContent {
	return v.value
}

func (v *NullableSBForecastingThrottlingExceptionResponseContent) Set(val *SBForecastingThrottlingExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingThrottlingExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingThrottlingExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingThrottlingExceptionResponseContent(val *SBForecastingThrottlingExceptionResponseContent) *NullableSBForecastingThrottlingExceptionResponseContent {
	return &NullableSBForecastingThrottlingExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSBForecastingThrottlingExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingThrottlingExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
