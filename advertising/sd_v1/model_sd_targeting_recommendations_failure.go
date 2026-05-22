package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingRecommendationsFailure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingRecommendationsFailure{}

// SDTargetingRecommendationsFailure A targeting recommendation failure record.
type SDTargetingRecommendationsFailure struct {
	// HTTP status code indicating a failure response for targeting recomendations.
	Code *string `json:"code,omitempty"`
	// The theme name specified in the request. If the themes field is not provided in the request, the value of this field will be set to default.
	Name *string `json:"name,omitempty"`
	// A human friendly error message indicating the failure reasons.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// NewSDTargetingRecommendationsFailure instantiates a new SDTargetingRecommendationsFailure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingRecommendationsFailure() *SDTargetingRecommendationsFailure {
	this := SDTargetingRecommendationsFailure{}
	return &this
}

// NewSDTargetingRecommendationsFailureWithDefaults instantiates a new SDTargetingRecommendationsFailure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingRecommendationsFailureWithDefaults() *SDTargetingRecommendationsFailure {
	this := SDTargetingRecommendationsFailure{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsFailure) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsFailure) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsFailure) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SDTargetingRecommendationsFailure) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsFailure) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsFailure) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsFailure) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SDTargetingRecommendationsFailure) SetName(v string) {
	o.Name = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsFailure) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsFailure) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsFailure) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *SDTargetingRecommendationsFailure) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o SDTargetingRecommendationsFailure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	return toSerialize, nil
}

type NullableSDTargetingRecommendationsFailure struct {
	value *SDTargetingRecommendationsFailure
	isSet bool
}

func (v NullableSDTargetingRecommendationsFailure) Get() *SDTargetingRecommendationsFailure {
	return v.value
}

func (v *NullableSDTargetingRecommendationsFailure) Set(val *SDTargetingRecommendationsFailure) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendationsFailure) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendationsFailure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendationsFailure(val *SDTargetingRecommendationsFailure) *NullableSDTargetingRecommendationsFailure {
	return &NullableSDTargetingRecommendationsFailure{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendationsFailure) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendationsFailure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
