package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBInsightsUnauthorizedExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBInsightsUnauthorizedExceptionResponseContent{}

// SBInsightsUnauthorizedExceptionResponseContent Returns information about an UnauthorizedException.
type SBInsightsUnauthorizedExceptionResponseContent struct {
	// The HTTP status code of the response.
	Code string `json:"code"`
	// A human-readable description of the response.
	Details string `json:"details"`
}

type _SBInsightsUnauthorizedExceptionResponseContent SBInsightsUnauthorizedExceptionResponseContent

// NewSBInsightsUnauthorizedExceptionResponseContent instantiates a new SBInsightsUnauthorizedExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBInsightsUnauthorizedExceptionResponseContent(code string, details string) *SBInsightsUnauthorizedExceptionResponseContent {
	this := SBInsightsUnauthorizedExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewSBInsightsUnauthorizedExceptionResponseContentWithDefaults instantiates a new SBInsightsUnauthorizedExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBInsightsUnauthorizedExceptionResponseContentWithDefaults() *SBInsightsUnauthorizedExceptionResponseContent {
	this := SBInsightsUnauthorizedExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SBInsightsUnauthorizedExceptionResponseContent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SBInsightsUnauthorizedExceptionResponseContent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SBInsightsUnauthorizedExceptionResponseContent) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *SBInsightsUnauthorizedExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SBInsightsUnauthorizedExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SBInsightsUnauthorizedExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o SBInsightsUnauthorizedExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableSBInsightsUnauthorizedExceptionResponseContent struct {
	value *SBInsightsUnauthorizedExceptionResponseContent
	isSet bool
}

func (v NullableSBInsightsUnauthorizedExceptionResponseContent) Get() *SBInsightsUnauthorizedExceptionResponseContent {
	return v.value
}

func (v *NullableSBInsightsUnauthorizedExceptionResponseContent) Set(val *SBInsightsUnauthorizedExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBInsightsUnauthorizedExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBInsightsUnauthorizedExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBInsightsUnauthorizedExceptionResponseContent(val *SBInsightsUnauthorizedExceptionResponseContent) *NullableSBInsightsUnauthorizedExceptionResponseContent {
	return &NullableSBInsightsUnauthorizedExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSBInsightsUnauthorizedExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBInsightsUnauthorizedExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
