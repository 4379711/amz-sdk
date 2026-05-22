package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingUnprocessableEntityExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingUnprocessableEntityExceptionResponseContent{}

// SBForecastingUnprocessableEntityExceptionResponseContent Returns information about an UnprocessableEntityException.
type SBForecastingUnprocessableEntityExceptionResponseContent struct {
	// The HTTP status code of the response.
	Code string `json:"code"`
	// A human-readable description of the response.
	Details string `json:"details"`
}

type _SBForecastingUnprocessableEntityExceptionResponseContent SBForecastingUnprocessableEntityExceptionResponseContent

// NewSBForecastingUnprocessableEntityExceptionResponseContent instantiates a new SBForecastingUnprocessableEntityExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingUnprocessableEntityExceptionResponseContent(code string, details string) *SBForecastingUnprocessableEntityExceptionResponseContent {
	this := SBForecastingUnprocessableEntityExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewSBForecastingUnprocessableEntityExceptionResponseContentWithDefaults instantiates a new SBForecastingUnprocessableEntityExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingUnprocessableEntityExceptionResponseContentWithDefaults() *SBForecastingUnprocessableEntityExceptionResponseContent {
	this := SBForecastingUnprocessableEntityExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SBForecastingUnprocessableEntityExceptionResponseContent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SBForecastingUnprocessableEntityExceptionResponseContent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SBForecastingUnprocessableEntityExceptionResponseContent) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *SBForecastingUnprocessableEntityExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SBForecastingUnprocessableEntityExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SBForecastingUnprocessableEntityExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o SBForecastingUnprocessableEntityExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableSBForecastingUnprocessableEntityExceptionResponseContent struct {
	value *SBForecastingUnprocessableEntityExceptionResponseContent
	isSet bool
}

func (v NullableSBForecastingUnprocessableEntityExceptionResponseContent) Get() *SBForecastingUnprocessableEntityExceptionResponseContent {
	return v.value
}

func (v *NullableSBForecastingUnprocessableEntityExceptionResponseContent) Set(val *SBForecastingUnprocessableEntityExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingUnprocessableEntityExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingUnprocessableEntityExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingUnprocessableEntityExceptionResponseContent(val *SBForecastingUnprocessableEntityExceptionResponseContent) *NullableSBForecastingUnprocessableEntityExceptionResponseContent {
	return &NullableSBForecastingUnprocessableEntityExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSBForecastingUnprocessableEntityExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingUnprocessableEntityExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
