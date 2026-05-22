package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SDHeadlineRecommendationInternalServerException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDHeadlineRecommendationInternalServerException{}

// SDHeadlineRecommendationInternalServerException struct for SDHeadlineRecommendationInternalServerException
type SDHeadlineRecommendationInternalServerException struct {
	// InternalErrorCode.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the error response.
	Details *string `json:"details,omitempty"`
}

// NewSDHeadlineRecommendationInternalServerException instantiates a new SDHeadlineRecommendationInternalServerException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDHeadlineRecommendationInternalServerException() *SDHeadlineRecommendationInternalServerException {
	this := SDHeadlineRecommendationInternalServerException{}
	return &this
}

// NewSDHeadlineRecommendationInternalServerExceptionWithDefaults instantiates a new SDHeadlineRecommendationInternalServerException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDHeadlineRecommendationInternalServerExceptionWithDefaults() *SDHeadlineRecommendationInternalServerException {
	this := SDHeadlineRecommendationInternalServerException{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SDHeadlineRecommendationInternalServerException) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDHeadlineRecommendationInternalServerException) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SDHeadlineRecommendationInternalServerException) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SDHeadlineRecommendationInternalServerException) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *SDHeadlineRecommendationInternalServerException) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDHeadlineRecommendationInternalServerException) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *SDHeadlineRecommendationInternalServerException) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *SDHeadlineRecommendationInternalServerException) SetDetails(v string) {
	o.Details = &v
}

func (o SDHeadlineRecommendationInternalServerException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableSDHeadlineRecommendationInternalServerException struct {
	value *SDHeadlineRecommendationInternalServerException
	isSet bool
}

func (v NullableSDHeadlineRecommendationInternalServerException) Get() *SDHeadlineRecommendationInternalServerException {
	return v.value
}

func (v *NullableSDHeadlineRecommendationInternalServerException) Set(val *SDHeadlineRecommendationInternalServerException) {
	v.value = val
	v.isSet = true
}

func (v NullableSDHeadlineRecommendationInternalServerException) IsSet() bool {
	return v.isSet
}

func (v *NullableSDHeadlineRecommendationInternalServerException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDHeadlineRecommendationInternalServerException(val *SDHeadlineRecommendationInternalServerException) *NullableSDHeadlineRecommendationInternalServerException {
	return &NullableSDHeadlineRecommendationInternalServerException{value: val, isSet: true}
}

func (v NullableSDHeadlineRecommendationInternalServerException) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDHeadlineRecommendationInternalServerException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
