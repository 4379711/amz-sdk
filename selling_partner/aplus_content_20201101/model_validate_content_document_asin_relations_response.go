package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the ValidateContentDocumentAsinRelationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateContentDocumentAsinRelationsResponse{}

// ValidateContentDocumentAsinRelationsResponse struct for ValidateContentDocumentAsinRelationsResponse
type ValidateContentDocumentAsinRelationsResponse struct {
	// A set of messages to the user, such as warnings or comments.
	Warnings []Error `json:"warnings,omitempty"`
	// A list of error responses that are returned when a request is unsuccessful.
	Errors []Error `json:"errors"`
}

type _ValidateContentDocumentAsinRelationsResponse ValidateContentDocumentAsinRelationsResponse

// NewValidateContentDocumentAsinRelationsResponse instantiates a new ValidateContentDocumentAsinRelationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateContentDocumentAsinRelationsResponse(errors []Error) *ValidateContentDocumentAsinRelationsResponse {
	this := ValidateContentDocumentAsinRelationsResponse{}
	this.Errors = errors
	return &this
}

// NewValidateContentDocumentAsinRelationsResponseWithDefaults instantiates a new ValidateContentDocumentAsinRelationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateContentDocumentAsinRelationsResponseWithDefaults() *ValidateContentDocumentAsinRelationsResponse {
	this := ValidateContentDocumentAsinRelationsResponse{}
	return &this
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ValidateContentDocumentAsinRelationsResponse) GetWarnings() []Error {
	if o == nil || IsNil(o.Warnings) {
		var ret []Error
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateContentDocumentAsinRelationsResponse) GetWarningsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ValidateContentDocumentAsinRelationsResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Error and assigns it to the Warnings field.
func (o *ValidateContentDocumentAsinRelationsResponse) SetWarnings(v []Error) {
	o.Warnings = v
}

// GetErrors returns the Errors field value
func (o *ValidateContentDocumentAsinRelationsResponse) GetErrors() []Error {
	if o == nil {
		var ret []Error
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *ValidateContentDocumentAsinRelationsResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *ValidateContentDocumentAsinRelationsResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o ValidateContentDocumentAsinRelationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	toSerialize["errors"] = o.Errors
	return toSerialize, nil
}

type NullableValidateContentDocumentAsinRelationsResponse struct {
	value *ValidateContentDocumentAsinRelationsResponse
	isSet bool
}

func (v NullableValidateContentDocumentAsinRelationsResponse) Get() *ValidateContentDocumentAsinRelationsResponse {
	return v.value
}

func (v *NullableValidateContentDocumentAsinRelationsResponse) Set(val *ValidateContentDocumentAsinRelationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateContentDocumentAsinRelationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateContentDocumentAsinRelationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateContentDocumentAsinRelationsResponse(val *ValidateContentDocumentAsinRelationsResponse) *NullableValidateContentDocumentAsinRelationsResponse {
	return &NullableValidateContentDocumentAsinRelationsResponse{value: val, isSet: true}
}

func (v NullableValidateContentDocumentAsinRelationsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableValidateContentDocumentAsinRelationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
