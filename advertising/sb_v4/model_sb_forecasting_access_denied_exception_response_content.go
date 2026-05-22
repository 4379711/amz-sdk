package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingAccessDeniedExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingAccessDeniedExceptionResponseContent{}

// SBForecastingAccessDeniedExceptionResponseContent Returns information about an AccessDeniedException.
type SBForecastingAccessDeniedExceptionResponseContent struct {
	// The HTTP status code of the response.
	Code string `json:"code"`
	// A human-readable description of the response.
	Details string `json:"details"`
}

type _SBForecastingAccessDeniedExceptionResponseContent SBForecastingAccessDeniedExceptionResponseContent

// NewSBForecastingAccessDeniedExceptionResponseContent instantiates a new SBForecastingAccessDeniedExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingAccessDeniedExceptionResponseContent(code string, details string) *SBForecastingAccessDeniedExceptionResponseContent {
	this := SBForecastingAccessDeniedExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewSBForecastingAccessDeniedExceptionResponseContentWithDefaults instantiates a new SBForecastingAccessDeniedExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingAccessDeniedExceptionResponseContentWithDefaults() *SBForecastingAccessDeniedExceptionResponseContent {
	this := SBForecastingAccessDeniedExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SBForecastingAccessDeniedExceptionResponseContent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SBForecastingAccessDeniedExceptionResponseContent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SBForecastingAccessDeniedExceptionResponseContent) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *SBForecastingAccessDeniedExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SBForecastingAccessDeniedExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SBForecastingAccessDeniedExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o SBForecastingAccessDeniedExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableSBForecastingAccessDeniedExceptionResponseContent struct {
	value *SBForecastingAccessDeniedExceptionResponseContent
	isSet bool
}

func (v NullableSBForecastingAccessDeniedExceptionResponseContent) Get() *SBForecastingAccessDeniedExceptionResponseContent {
	return v.value
}

func (v *NullableSBForecastingAccessDeniedExceptionResponseContent) Set(val *SBForecastingAccessDeniedExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingAccessDeniedExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingAccessDeniedExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingAccessDeniedExceptionResponseContent(val *SBForecastingAccessDeniedExceptionResponseContent) *NullableSBForecastingAccessDeniedExceptionResponseContent {
	return &NullableSBForecastingAccessDeniedExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSBForecastingAccessDeniedExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingAccessDeniedExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
