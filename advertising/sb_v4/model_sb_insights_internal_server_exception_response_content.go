package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBInsightsInternalServerExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBInsightsInternalServerExceptionResponseContent{}

// SBInsightsInternalServerExceptionResponseContent Returns information about an InternalServerException.
type SBInsightsInternalServerExceptionResponseContent struct {
	// The HTTP status code of the response.
	Code string `json:"code"`
	// A human-readable description of the response.
	Details string `json:"details"`
}

type _SBInsightsInternalServerExceptionResponseContent SBInsightsInternalServerExceptionResponseContent

// NewSBInsightsInternalServerExceptionResponseContent instantiates a new SBInsightsInternalServerExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBInsightsInternalServerExceptionResponseContent(code string, details string) *SBInsightsInternalServerExceptionResponseContent {
	this := SBInsightsInternalServerExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewSBInsightsInternalServerExceptionResponseContentWithDefaults instantiates a new SBInsightsInternalServerExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBInsightsInternalServerExceptionResponseContentWithDefaults() *SBInsightsInternalServerExceptionResponseContent {
	this := SBInsightsInternalServerExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SBInsightsInternalServerExceptionResponseContent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SBInsightsInternalServerExceptionResponseContent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SBInsightsInternalServerExceptionResponseContent) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *SBInsightsInternalServerExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SBInsightsInternalServerExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SBInsightsInternalServerExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o SBInsightsInternalServerExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableSBInsightsInternalServerExceptionResponseContent struct {
	value *SBInsightsInternalServerExceptionResponseContent
	isSet bool
}

func (v NullableSBInsightsInternalServerExceptionResponseContent) Get() *SBInsightsInternalServerExceptionResponseContent {
	return v.value
}

func (v *NullableSBInsightsInternalServerExceptionResponseContent) Set(val *SBInsightsInternalServerExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBInsightsInternalServerExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBInsightsInternalServerExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBInsightsInternalServerExceptionResponseContent(val *SBInsightsInternalServerExceptionResponseContent) *NullableSBInsightsInternalServerExceptionResponseContent {
	return &NullableSBInsightsInternalServerExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSBInsightsInternalServerExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBInsightsInternalServerExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
