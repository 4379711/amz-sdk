package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the PostContentDocumentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostContentDocumentResponse{}

// PostContentDocumentResponse struct for PostContentDocumentResponse
type PostContentDocumentResponse struct {
	// A set of messages to the user, such as warnings or comments.
	Warnings []Error `json:"warnings,omitempty"`
	// A unique reference key for the A+ Content document. A content reference key cannot form a permalink and might change in the future. A content reference key is not guaranteed to match any A+ content identifier.
	ContentReferenceKey string `json:"contentReferenceKey"`
}

type _PostContentDocumentResponse PostContentDocumentResponse

// NewPostContentDocumentResponse instantiates a new PostContentDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostContentDocumentResponse(contentReferenceKey string) *PostContentDocumentResponse {
	this := PostContentDocumentResponse{}
	this.ContentReferenceKey = contentReferenceKey
	return &this
}

// NewPostContentDocumentResponseWithDefaults instantiates a new PostContentDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostContentDocumentResponseWithDefaults() *PostContentDocumentResponse {
	this := PostContentDocumentResponse{}
	return &this
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostContentDocumentResponse) GetWarnings() []Error {
	if o == nil || IsNil(o.Warnings) {
		var ret []Error
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostContentDocumentResponse) GetWarningsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostContentDocumentResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Error and assigns it to the Warnings field.
func (o *PostContentDocumentResponse) SetWarnings(v []Error) {
	o.Warnings = v
}

// GetContentReferenceKey returns the ContentReferenceKey field value
func (o *PostContentDocumentResponse) GetContentReferenceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentReferenceKey
}

// GetContentReferenceKeyOk returns a tuple with the ContentReferenceKey field value
// and a boolean to check if the value has been set.
func (o *PostContentDocumentResponse) GetContentReferenceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentReferenceKey, true
}

// SetContentReferenceKey sets field value
func (o *PostContentDocumentResponse) SetContentReferenceKey(v string) {
	o.ContentReferenceKey = v
}

func (o PostContentDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	toSerialize["contentReferenceKey"] = o.ContentReferenceKey
	return toSerialize, nil
}

type NullablePostContentDocumentResponse struct {
	value *PostContentDocumentResponse
	isSet bool
}

func (v NullablePostContentDocumentResponse) Get() *PostContentDocumentResponse {
	return v.value
}

func (v *NullablePostContentDocumentResponse) Set(val *PostContentDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostContentDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostContentDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostContentDocumentResponse(val *PostContentDocumentResponse) *NullablePostContentDocumentResponse {
	return &NullablePostContentDocumentResponse{value: val, isSet: true}
}

func (v NullablePostContentDocumentResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePostContentDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
