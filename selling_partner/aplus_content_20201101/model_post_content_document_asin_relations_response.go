package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the PostContentDocumentAsinRelationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostContentDocumentAsinRelationsResponse{}

// PostContentDocumentAsinRelationsResponse struct for PostContentDocumentAsinRelationsResponse
type PostContentDocumentAsinRelationsResponse struct {
	// A set of messages to the user, such as warnings or comments.
	Warnings []Error `json:"warnings,omitempty"`
}

// NewPostContentDocumentAsinRelationsResponse instantiates a new PostContentDocumentAsinRelationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostContentDocumentAsinRelationsResponse() *PostContentDocumentAsinRelationsResponse {
	this := PostContentDocumentAsinRelationsResponse{}
	return &this
}

// NewPostContentDocumentAsinRelationsResponseWithDefaults instantiates a new PostContentDocumentAsinRelationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostContentDocumentAsinRelationsResponseWithDefaults() *PostContentDocumentAsinRelationsResponse {
	this := PostContentDocumentAsinRelationsResponse{}
	return &this
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostContentDocumentAsinRelationsResponse) GetWarnings() []Error {
	if o == nil || IsNil(o.Warnings) {
		var ret []Error
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostContentDocumentAsinRelationsResponse) GetWarningsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostContentDocumentAsinRelationsResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Error and assigns it to the Warnings field.
func (o *PostContentDocumentAsinRelationsResponse) SetWarnings(v []Error) {
	o.Warnings = v
}

func (o PostContentDocumentAsinRelationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostContentDocumentAsinRelationsResponse struct {
	value *PostContentDocumentAsinRelationsResponse
	isSet bool
}

func (v NullablePostContentDocumentAsinRelationsResponse) Get() *PostContentDocumentAsinRelationsResponse {
	return v.value
}

func (v *NullablePostContentDocumentAsinRelationsResponse) Set(val *PostContentDocumentAsinRelationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostContentDocumentAsinRelationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostContentDocumentAsinRelationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostContentDocumentAsinRelationsResponse(val *PostContentDocumentAsinRelationsResponse) *NullablePostContentDocumentAsinRelationsResponse {
	return &NullablePostContentDocumentAsinRelationsResponse{value: val, isSet: true}
}

func (v NullablePostContentDocumentAsinRelationsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePostContentDocumentAsinRelationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
