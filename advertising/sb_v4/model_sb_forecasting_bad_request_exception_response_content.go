package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingBadRequestExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingBadRequestExceptionResponseContent{}

// SBForecastingBadRequestExceptionResponseContent Returns information about a BadRequestException.
type SBForecastingBadRequestExceptionResponseContent struct {
	// The HTTP status code of the response.
	Code string `json:"code"`
	// A human-readable description of the response.
	Details string `json:"details"`
}

type _SBForecastingBadRequestExceptionResponseContent SBForecastingBadRequestExceptionResponseContent

// NewSBForecastingBadRequestExceptionResponseContent instantiates a new SBForecastingBadRequestExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingBadRequestExceptionResponseContent(code string, details string) *SBForecastingBadRequestExceptionResponseContent {
	this := SBForecastingBadRequestExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewSBForecastingBadRequestExceptionResponseContentWithDefaults instantiates a new SBForecastingBadRequestExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingBadRequestExceptionResponseContentWithDefaults() *SBForecastingBadRequestExceptionResponseContent {
	this := SBForecastingBadRequestExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SBForecastingBadRequestExceptionResponseContent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SBForecastingBadRequestExceptionResponseContent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SBForecastingBadRequestExceptionResponseContent) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *SBForecastingBadRequestExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SBForecastingBadRequestExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SBForecastingBadRequestExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o SBForecastingBadRequestExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableSBForecastingBadRequestExceptionResponseContent struct {
	value *SBForecastingBadRequestExceptionResponseContent
	isSet bool
}

func (v NullableSBForecastingBadRequestExceptionResponseContent) Get() *SBForecastingBadRequestExceptionResponseContent {
	return v.value
}

func (v *NullableSBForecastingBadRequestExceptionResponseContent) Set(val *SBForecastingBadRequestExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingBadRequestExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingBadRequestExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingBadRequestExceptionResponseContent(val *SBForecastingBadRequestExceptionResponseContent) *NullableSBForecastingBadRequestExceptionResponseContent {
	return &NullableSBForecastingBadRequestExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSBForecastingBadRequestExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingBadRequestExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
