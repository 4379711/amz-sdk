package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SDHeadlineRecommendationIdentifierNotfoundException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDHeadlineRecommendationIdentifierNotfoundException{}

// SDHeadlineRecommendationIdentifierNotfoundException struct for SDHeadlineRecommendationIdentifierNotfoundException
type SDHeadlineRecommendationIdentifierNotfoundException struct {
	// IdentiferNotFoundErrorCode.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the error response.
	Details *string `json:"details,omitempty"`
}

// NewSDHeadlineRecommendationIdentifierNotfoundException instantiates a new SDHeadlineRecommendationIdentifierNotfoundException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDHeadlineRecommendationIdentifierNotfoundException() *SDHeadlineRecommendationIdentifierNotfoundException {
	this := SDHeadlineRecommendationIdentifierNotfoundException{}
	return &this
}

// NewSDHeadlineRecommendationIdentifierNotfoundExceptionWithDefaults instantiates a new SDHeadlineRecommendationIdentifierNotfoundException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDHeadlineRecommendationIdentifierNotfoundExceptionWithDefaults() *SDHeadlineRecommendationIdentifierNotfoundException {
	this := SDHeadlineRecommendationIdentifierNotfoundException{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SDHeadlineRecommendationIdentifierNotfoundException) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDHeadlineRecommendationIdentifierNotfoundException) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SDHeadlineRecommendationIdentifierNotfoundException) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SDHeadlineRecommendationIdentifierNotfoundException) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *SDHeadlineRecommendationIdentifierNotfoundException) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDHeadlineRecommendationIdentifierNotfoundException) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *SDHeadlineRecommendationIdentifierNotfoundException) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *SDHeadlineRecommendationIdentifierNotfoundException) SetDetails(v string) {
	o.Details = &v
}

func (o SDHeadlineRecommendationIdentifierNotfoundException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableSDHeadlineRecommendationIdentifierNotfoundException struct {
	value *SDHeadlineRecommendationIdentifierNotfoundException
	isSet bool
}

func (v NullableSDHeadlineRecommendationIdentifierNotfoundException) Get() *SDHeadlineRecommendationIdentifierNotfoundException {
	return v.value
}

func (v *NullableSDHeadlineRecommendationIdentifierNotfoundException) Set(val *SDHeadlineRecommendationIdentifierNotfoundException) {
	v.value = val
	v.isSet = true
}

func (v NullableSDHeadlineRecommendationIdentifierNotfoundException) IsSet() bool {
	return v.isSet
}

func (v *NullableSDHeadlineRecommendationIdentifierNotfoundException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDHeadlineRecommendationIdentifierNotfoundException(val *SDHeadlineRecommendationIdentifierNotfoundException) *NullableSDHeadlineRecommendationIdentifierNotfoundException {
	return &NullableSDHeadlineRecommendationIdentifierNotfoundException{value: val, isSet: true}
}

func (v NullableSDHeadlineRecommendationIdentifierNotfoundException) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDHeadlineRecommendationIdentifierNotfoundException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
