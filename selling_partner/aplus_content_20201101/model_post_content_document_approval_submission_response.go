package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the PostContentDocumentApprovalSubmissionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostContentDocumentApprovalSubmissionResponse{}

// PostContentDocumentApprovalSubmissionResponse struct for PostContentDocumentApprovalSubmissionResponse
type PostContentDocumentApprovalSubmissionResponse struct {
	// A set of messages to the user, such as warnings or comments.
	Warnings []Error `json:"warnings,omitempty"`
}

// NewPostContentDocumentApprovalSubmissionResponse instantiates a new PostContentDocumentApprovalSubmissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostContentDocumentApprovalSubmissionResponse() *PostContentDocumentApprovalSubmissionResponse {
	this := PostContentDocumentApprovalSubmissionResponse{}
	return &this
}

// NewPostContentDocumentApprovalSubmissionResponseWithDefaults instantiates a new PostContentDocumentApprovalSubmissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostContentDocumentApprovalSubmissionResponseWithDefaults() *PostContentDocumentApprovalSubmissionResponse {
	this := PostContentDocumentApprovalSubmissionResponse{}
	return &this
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostContentDocumentApprovalSubmissionResponse) GetWarnings() []Error {
	if o == nil || IsNil(o.Warnings) {
		var ret []Error
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostContentDocumentApprovalSubmissionResponse) GetWarningsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostContentDocumentApprovalSubmissionResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Error and assigns it to the Warnings field.
func (o *PostContentDocumentApprovalSubmissionResponse) SetWarnings(v []Error) {
	o.Warnings = v
}

func (o PostContentDocumentApprovalSubmissionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostContentDocumentApprovalSubmissionResponse struct {
	value *PostContentDocumentApprovalSubmissionResponse
	isSet bool
}

func (v NullablePostContentDocumentApprovalSubmissionResponse) Get() *PostContentDocumentApprovalSubmissionResponse {
	return v.value
}

func (v *NullablePostContentDocumentApprovalSubmissionResponse) Set(val *PostContentDocumentApprovalSubmissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostContentDocumentApprovalSubmissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostContentDocumentApprovalSubmissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostContentDocumentApprovalSubmissionResponse(val *PostContentDocumentApprovalSubmissionResponse) *NullablePostContentDocumentApprovalSubmissionResponse {
	return &NullablePostContentDocumentApprovalSubmissionResponse{value: val, isSet: true}
}

func (v NullablePostContentDocumentApprovalSubmissionResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePostContentDocumentApprovalSubmissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
