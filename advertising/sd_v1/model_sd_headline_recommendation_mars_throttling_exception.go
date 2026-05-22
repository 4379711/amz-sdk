package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDHeadlineRecommendationMarsThrottlingException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDHeadlineRecommendationMarsThrottlingException{}

// SDHeadlineRecommendationMarsThrottlingException struct for SDHeadlineRecommendationMarsThrottlingException
type SDHeadlineRecommendationMarsThrottlingException struct {
	// ThrottledErrorCode.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the error response.
	Details *string `json:"details,omitempty"`
}

// NewSDHeadlineRecommendationMarsThrottlingException instantiates a new SDHeadlineRecommendationMarsThrottlingException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDHeadlineRecommendationMarsThrottlingException() *SDHeadlineRecommendationMarsThrottlingException {
	this := SDHeadlineRecommendationMarsThrottlingException{}
	return &this
}

// NewSDHeadlineRecommendationMarsThrottlingExceptionWithDefaults instantiates a new SDHeadlineRecommendationMarsThrottlingException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDHeadlineRecommendationMarsThrottlingExceptionWithDefaults() *SDHeadlineRecommendationMarsThrottlingException {
	this := SDHeadlineRecommendationMarsThrottlingException{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SDHeadlineRecommendationMarsThrottlingException) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDHeadlineRecommendationMarsThrottlingException) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SDHeadlineRecommendationMarsThrottlingException) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SDHeadlineRecommendationMarsThrottlingException) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *SDHeadlineRecommendationMarsThrottlingException) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDHeadlineRecommendationMarsThrottlingException) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *SDHeadlineRecommendationMarsThrottlingException) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *SDHeadlineRecommendationMarsThrottlingException) SetDetails(v string) {
	o.Details = &v
}

func (o SDHeadlineRecommendationMarsThrottlingException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableSDHeadlineRecommendationMarsThrottlingException struct {
	value *SDHeadlineRecommendationMarsThrottlingException
	isSet bool
}

func (v NullableSDHeadlineRecommendationMarsThrottlingException) Get() *SDHeadlineRecommendationMarsThrottlingException {
	return v.value
}

func (v *NullableSDHeadlineRecommendationMarsThrottlingException) Set(val *SDHeadlineRecommendationMarsThrottlingException) {
	v.value = val
	v.isSet = true
}

func (v NullableSDHeadlineRecommendationMarsThrottlingException) IsSet() bool {
	return v.isSet
}

func (v *NullableSDHeadlineRecommendationMarsThrottlingException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDHeadlineRecommendationMarsThrottlingException(val *SDHeadlineRecommendationMarsThrottlingException) *NullableSDHeadlineRecommendationMarsThrottlingException {
	return &NullableSDHeadlineRecommendationMarsThrottlingException{value: val, isSet: true}
}

func (v NullableSDHeadlineRecommendationMarsThrottlingException) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDHeadlineRecommendationMarsThrottlingException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
