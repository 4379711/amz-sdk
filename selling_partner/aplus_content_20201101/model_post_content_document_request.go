package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the PostContentDocumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostContentDocumentRequest{}

// PostContentDocumentRequest struct for PostContentDocumentRequest
type PostContentDocumentRequest struct {
	ContentDocument ContentDocument `json:"contentDocument"`
}

type _PostContentDocumentRequest PostContentDocumentRequest

// NewPostContentDocumentRequest instantiates a new PostContentDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostContentDocumentRequest(contentDocument ContentDocument) *PostContentDocumentRequest {
	this := PostContentDocumentRequest{}
	this.ContentDocument = contentDocument
	return &this
}

// NewPostContentDocumentRequestWithDefaults instantiates a new PostContentDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostContentDocumentRequestWithDefaults() *PostContentDocumentRequest {
	this := PostContentDocumentRequest{}
	return &this
}

// GetContentDocument returns the ContentDocument field value
func (o *PostContentDocumentRequest) GetContentDocument() ContentDocument {
	if o == nil {
		var ret ContentDocument
		return ret
	}

	return o.ContentDocument
}

// GetContentDocumentOk returns a tuple with the ContentDocument field value
// and a boolean to check if the value has been set.
func (o *PostContentDocumentRequest) GetContentDocumentOk() (*ContentDocument, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentDocument, true
}

// SetContentDocument sets field value
func (o *PostContentDocumentRequest) SetContentDocument(v ContentDocument) {
	o.ContentDocument = v
}

func (o PostContentDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contentDocument"] = o.ContentDocument
	return toSerialize, nil
}

type NullablePostContentDocumentRequest struct {
	value *PostContentDocumentRequest
	isSet bool
}

func (v NullablePostContentDocumentRequest) Get() *PostContentDocumentRequest {
	return v.value
}

func (v *NullablePostContentDocumentRequest) Set(val *PostContentDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostContentDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostContentDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostContentDocumentRequest(val *PostContentDocumentRequest) *NullablePostContentDocumentRequest {
	return &NullablePostContentDocumentRequest{value: val, isSet: true}
}

func (v NullablePostContentDocumentRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePostContentDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
