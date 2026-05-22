package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBInsightsUnprocessableEntityExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBInsightsUnprocessableEntityExceptionResponseContent{}

// SBInsightsUnprocessableEntityExceptionResponseContent Returns information about an UnprocessableEntityException.
type SBInsightsUnprocessableEntityExceptionResponseContent struct {
	// The HTTP status code of the response.
	Code string `json:"code"`
	// A human-readable description of the response.
	Details string `json:"details"`
}

type _SBInsightsUnprocessableEntityExceptionResponseContent SBInsightsUnprocessableEntityExceptionResponseContent

// NewSBInsightsUnprocessableEntityExceptionResponseContent instantiates a new SBInsightsUnprocessableEntityExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBInsightsUnprocessableEntityExceptionResponseContent(code string, details string) *SBInsightsUnprocessableEntityExceptionResponseContent {
	this := SBInsightsUnprocessableEntityExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewSBInsightsUnprocessableEntityExceptionResponseContentWithDefaults instantiates a new SBInsightsUnprocessableEntityExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBInsightsUnprocessableEntityExceptionResponseContentWithDefaults() *SBInsightsUnprocessableEntityExceptionResponseContent {
	this := SBInsightsUnprocessableEntityExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SBInsightsUnprocessableEntityExceptionResponseContent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SBInsightsUnprocessableEntityExceptionResponseContent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SBInsightsUnprocessableEntityExceptionResponseContent) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *SBInsightsUnprocessableEntityExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SBInsightsUnprocessableEntityExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SBInsightsUnprocessableEntityExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o SBInsightsUnprocessableEntityExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableSBInsightsUnprocessableEntityExceptionResponseContent struct {
	value *SBInsightsUnprocessableEntityExceptionResponseContent
	isSet bool
}

func (v NullableSBInsightsUnprocessableEntityExceptionResponseContent) Get() *SBInsightsUnprocessableEntityExceptionResponseContent {
	return v.value
}

func (v *NullableSBInsightsUnprocessableEntityExceptionResponseContent) Set(val *SBInsightsUnprocessableEntityExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBInsightsUnprocessableEntityExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBInsightsUnprocessableEntityExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBInsightsUnprocessableEntityExceptionResponseContent(val *SBInsightsUnprocessableEntityExceptionResponseContent) *NullableSBInsightsUnprocessableEntityExceptionResponseContent {
	return &NullableSBInsightsUnprocessableEntityExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSBInsightsUnprocessableEntityExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBInsightsUnprocessableEntityExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
