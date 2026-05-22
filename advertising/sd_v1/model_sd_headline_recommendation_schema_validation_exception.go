package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDHeadlineRecommendationSchemaValidationException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDHeadlineRecommendationSchemaValidationException{}

// SDHeadlineRecommendationSchemaValidationException struct for SDHeadlineRecommendationSchemaValidationException
type SDHeadlineRecommendationSchemaValidationException struct {
	// InvalidArgumentErrorCode.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the error response.
	Details *string `json:"details,omitempty"`
}

// NewSDHeadlineRecommendationSchemaValidationException instantiates a new SDHeadlineRecommendationSchemaValidationException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDHeadlineRecommendationSchemaValidationException() *SDHeadlineRecommendationSchemaValidationException {
	this := SDHeadlineRecommendationSchemaValidationException{}
	return &this
}

// NewSDHeadlineRecommendationSchemaValidationExceptionWithDefaults instantiates a new SDHeadlineRecommendationSchemaValidationException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDHeadlineRecommendationSchemaValidationExceptionWithDefaults() *SDHeadlineRecommendationSchemaValidationException {
	this := SDHeadlineRecommendationSchemaValidationException{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SDHeadlineRecommendationSchemaValidationException) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDHeadlineRecommendationSchemaValidationException) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SDHeadlineRecommendationSchemaValidationException) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SDHeadlineRecommendationSchemaValidationException) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *SDHeadlineRecommendationSchemaValidationException) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDHeadlineRecommendationSchemaValidationException) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *SDHeadlineRecommendationSchemaValidationException) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *SDHeadlineRecommendationSchemaValidationException) SetDetails(v string) {
	o.Details = &v
}

func (o SDHeadlineRecommendationSchemaValidationException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableSDHeadlineRecommendationSchemaValidationException struct {
	value *SDHeadlineRecommendationSchemaValidationException
	isSet bool
}

func (v NullableSDHeadlineRecommendationSchemaValidationException) Get() *SDHeadlineRecommendationSchemaValidationException {
	return v.value
}

func (v *NullableSDHeadlineRecommendationSchemaValidationException) Set(val *SDHeadlineRecommendationSchemaValidationException) {
	v.value = val
	v.isSet = true
}

func (v NullableSDHeadlineRecommendationSchemaValidationException) IsSet() bool {
	return v.isSet
}

func (v *NullableSDHeadlineRecommendationSchemaValidationException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDHeadlineRecommendationSchemaValidationException(val *SDHeadlineRecommendationSchemaValidationException) *NullableSDHeadlineRecommendationSchemaValidationException {
	return &NullableSDHeadlineRecommendationSchemaValidationException{value: val, isSet: true}
}

func (v NullableSDHeadlineRecommendationSchemaValidationException) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDHeadlineRecommendationSchemaValidationException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
