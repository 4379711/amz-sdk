package services

import (
	"github.com/bytedance/sonic"
)

// checks if the ServiceUploadDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceUploadDocument{}

// ServiceUploadDocument Input for to be uploaded document.
type ServiceUploadDocument struct {
	// The content type of the to-be-uploaded file
	ContentType string `json:"contentType"`
	// The content length of the to-be-uploaded file
	ContentLength float32 `json:"contentLength"`
	// An MD5 hash of the content to be submitted to the upload destination. This value is used to determine if the data has been corrupted or tampered with during transit.
	ContentMD5 *string `json:"contentMD5,omitempty" validate:"regexp=^[A-Za-z0-9\\\\\\\\+\\/]{22}={2}$"`
}

type _ServiceUploadDocument ServiceUploadDocument

// NewServiceUploadDocument instantiates a new ServiceUploadDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceUploadDocument(contentType string, contentLength float32) *ServiceUploadDocument {
	this := ServiceUploadDocument{}
	this.ContentType = contentType
	this.ContentLength = contentLength
	return &this
}

// NewServiceUploadDocumentWithDefaults instantiates a new ServiceUploadDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceUploadDocumentWithDefaults() *ServiceUploadDocument {
	this := ServiceUploadDocument{}
	return &this
}

// GetContentType returns the ContentType field value
func (o *ServiceUploadDocument) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *ServiceUploadDocument) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *ServiceUploadDocument) SetContentType(v string) {
	o.ContentType = v
}

// GetContentLength returns the ContentLength field value
func (o *ServiceUploadDocument) GetContentLength() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ContentLength
}

// GetContentLengthOk returns a tuple with the ContentLength field value
// and a boolean to check if the value has been set.
func (o *ServiceUploadDocument) GetContentLengthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentLength, true
}

// SetContentLength sets field value
func (o *ServiceUploadDocument) SetContentLength(v float32) {
	o.ContentLength = v
}

// GetContentMD5 returns the ContentMD5 field value if set, zero value otherwise.
func (o *ServiceUploadDocument) GetContentMD5() string {
	if o == nil || IsNil(o.ContentMD5) {
		var ret string
		return ret
	}
	return *o.ContentMD5
}

// GetContentMD5Ok returns a tuple with the ContentMD5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUploadDocument) GetContentMD5Ok() (*string, bool) {
	if o == nil || IsNil(o.ContentMD5) {
		return nil, false
	}
	return o.ContentMD5, true
}

// HasContentMD5 returns a boolean if a field has been set.
func (o *ServiceUploadDocument) HasContentMD5() bool {
	if o != nil && !IsNil(o.ContentMD5) {
		return true
	}

	return false
}

// SetContentMD5 gets a reference to the given string and assigns it to the ContentMD5 field.
func (o *ServiceUploadDocument) SetContentMD5(v string) {
	o.ContentMD5 = &v
}

func (o ServiceUploadDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contentType"] = o.ContentType
	toSerialize["contentLength"] = o.ContentLength
	if !IsNil(o.ContentMD5) {
		toSerialize["contentMD5"] = o.ContentMD5
	}
	return toSerialize, nil
}

type NullableServiceUploadDocument struct {
	value *ServiceUploadDocument
	isSet bool
}

func (v NullableServiceUploadDocument) Get() *ServiceUploadDocument {
	return v.value
}

func (v *NullableServiceUploadDocument) Set(val *ServiceUploadDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceUploadDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceUploadDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceUploadDocument(val *ServiceUploadDocument) *NullableServiceUploadDocument {
	return &NullableServiceUploadDocument{value: val, isSet: true}
}

func (v NullableServiceUploadDocument) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableServiceUploadDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
