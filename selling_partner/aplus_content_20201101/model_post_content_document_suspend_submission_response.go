package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the PostContentDocumentSuspendSubmissionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostContentDocumentSuspendSubmissionResponse{}

// PostContentDocumentSuspendSubmissionResponse struct for PostContentDocumentSuspendSubmissionResponse
type PostContentDocumentSuspendSubmissionResponse struct {
	// A set of messages to the user, such as warnings or comments.
	Warnings []Error `json:"warnings,omitempty"`
}

// NewPostContentDocumentSuspendSubmissionResponse instantiates a new PostContentDocumentSuspendSubmissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostContentDocumentSuspendSubmissionResponse() *PostContentDocumentSuspendSubmissionResponse {
	this := PostContentDocumentSuspendSubmissionResponse{}
	return &this
}

// NewPostContentDocumentSuspendSubmissionResponseWithDefaults instantiates a new PostContentDocumentSuspendSubmissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostContentDocumentSuspendSubmissionResponseWithDefaults() *PostContentDocumentSuspendSubmissionResponse {
	this := PostContentDocumentSuspendSubmissionResponse{}
	return &this
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostContentDocumentSuspendSubmissionResponse) GetWarnings() []Error {
	if o == nil || IsNil(o.Warnings) {
		var ret []Error
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostContentDocumentSuspendSubmissionResponse) GetWarningsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostContentDocumentSuspendSubmissionResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Error and assigns it to the Warnings field.
func (o *PostContentDocumentSuspendSubmissionResponse) SetWarnings(v []Error) {
	o.Warnings = v
}

func (o PostContentDocumentSuspendSubmissionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostContentDocumentSuspendSubmissionResponse struct {
	value *PostContentDocumentSuspendSubmissionResponse
	isSet bool
}

func (v NullablePostContentDocumentSuspendSubmissionResponse) Get() *PostContentDocumentSuspendSubmissionResponse {
	return v.value
}

func (v *NullablePostContentDocumentSuspendSubmissionResponse) Set(val *PostContentDocumentSuspendSubmissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostContentDocumentSuspendSubmissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostContentDocumentSuspendSubmissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostContentDocumentSuspendSubmissionResponse(val *PostContentDocumentSuspendSubmissionResponse) *NullablePostContentDocumentSuspendSubmissionResponse {
	return &NullablePostContentDocumentSuspendSubmissionResponse{value: val, isSet: true}
}

func (v NullablePostContentDocumentSuspendSubmissionResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePostContentDocumentSuspendSubmissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
