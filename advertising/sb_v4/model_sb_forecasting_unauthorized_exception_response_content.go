package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingUnauthorizedExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingUnauthorizedExceptionResponseContent{}

// SBForecastingUnauthorizedExceptionResponseContent Returns information about an UnauthorizedException.
type SBForecastingUnauthorizedExceptionResponseContent struct {
	// The HTTP status code of the response.
	Code string `json:"code"`
	// A human-readable description of the response.
	Details string `json:"details"`
}

type _SBForecastingUnauthorizedExceptionResponseContent SBForecastingUnauthorizedExceptionResponseContent

// NewSBForecastingUnauthorizedExceptionResponseContent instantiates a new SBForecastingUnauthorizedExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingUnauthorizedExceptionResponseContent(code string, details string) *SBForecastingUnauthorizedExceptionResponseContent {
	this := SBForecastingUnauthorizedExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewSBForecastingUnauthorizedExceptionResponseContentWithDefaults instantiates a new SBForecastingUnauthorizedExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingUnauthorizedExceptionResponseContentWithDefaults() *SBForecastingUnauthorizedExceptionResponseContent {
	this := SBForecastingUnauthorizedExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SBForecastingUnauthorizedExceptionResponseContent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SBForecastingUnauthorizedExceptionResponseContent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SBForecastingUnauthorizedExceptionResponseContent) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *SBForecastingUnauthorizedExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SBForecastingUnauthorizedExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SBForecastingUnauthorizedExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o SBForecastingUnauthorizedExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableSBForecastingUnauthorizedExceptionResponseContent struct {
	value *SBForecastingUnauthorizedExceptionResponseContent
	isSet bool
}

func (v NullableSBForecastingUnauthorizedExceptionResponseContent) Get() *SBForecastingUnauthorizedExceptionResponseContent {
	return v.value
}

func (v *NullableSBForecastingUnauthorizedExceptionResponseContent) Set(val *SBForecastingUnauthorizedExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingUnauthorizedExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingUnauthorizedExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingUnauthorizedExceptionResponseContent(val *SBForecastingUnauthorizedExceptionResponseContent) *NullableSBForecastingUnauthorizedExceptionResponseContent {
	return &NullableSBForecastingUnauthorizedExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSBForecastingUnauthorizedExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingUnauthorizedExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
