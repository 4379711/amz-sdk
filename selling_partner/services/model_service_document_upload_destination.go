package services

import (
	"github.com/bytedance/sonic"
)

// checks if the ServiceDocumentUploadDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDocumentUploadDestination{}

// ServiceDocumentUploadDestination Information about an upload destination.
type ServiceDocumentUploadDestination struct {
	// The unique identifier to be used by APIs that reference the upload destination.
	UploadDestinationId string `json:"uploadDestinationId"`
	// The URL to which to upload the file.
	Url               string            `json:"url"`
	EncryptionDetails EncryptionDetails `json:"encryptionDetails"`
	// The headers to include in the upload request.
	Headers map[string]interface{} `json:"headers,omitempty"`
}

type _ServiceDocumentUploadDestination ServiceDocumentUploadDestination

// NewServiceDocumentUploadDestination instantiates a new ServiceDocumentUploadDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDocumentUploadDestination(uploadDestinationId string, url string, encryptionDetails EncryptionDetails) *ServiceDocumentUploadDestination {
	this := ServiceDocumentUploadDestination{}
	this.UploadDestinationId = uploadDestinationId
	this.Url = url
	this.EncryptionDetails = encryptionDetails
	return &this
}

// NewServiceDocumentUploadDestinationWithDefaults instantiates a new ServiceDocumentUploadDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDocumentUploadDestinationWithDefaults() *ServiceDocumentUploadDestination {
	this := ServiceDocumentUploadDestination{}
	return &this
}

// GetUploadDestinationId returns the UploadDestinationId field value
func (o *ServiceDocumentUploadDestination) GetUploadDestinationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploadDestinationId
}

// GetUploadDestinationIdOk returns a tuple with the UploadDestinationId field value
// and a boolean to check if the value has been set.
func (o *ServiceDocumentUploadDestination) GetUploadDestinationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadDestinationId, true
}

// SetUploadDestinationId sets field value
func (o *ServiceDocumentUploadDestination) SetUploadDestinationId(v string) {
	o.UploadDestinationId = v
}

// GetUrl returns the Url field value
func (o *ServiceDocumentUploadDestination) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ServiceDocumentUploadDestination) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ServiceDocumentUploadDestination) SetUrl(v string) {
	o.Url = v
}

// GetEncryptionDetails returns the EncryptionDetails field value
func (o *ServiceDocumentUploadDestination) GetEncryptionDetails() EncryptionDetails {
	if o == nil {
		var ret EncryptionDetails
		return ret
	}

	return o.EncryptionDetails
}

// GetEncryptionDetailsOk returns a tuple with the EncryptionDetails field value
// and a boolean to check if the value has been set.
func (o *ServiceDocumentUploadDestination) GetEncryptionDetailsOk() (*EncryptionDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptionDetails, true
}

// SetEncryptionDetails sets field value
func (o *ServiceDocumentUploadDestination) SetEncryptionDetails(v EncryptionDetails) {
	o.EncryptionDetails = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *ServiceDocumentUploadDestination) GetHeaders() map[string]interface{} {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocumentUploadDestination) GetHeadersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Headers) {
		return map[string]interface{}{}, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *ServiceDocumentUploadDestination) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]interface{} and assigns it to the Headers field.
func (o *ServiceDocumentUploadDestination) SetHeaders(v map[string]interface{}) {
	o.Headers = v
}

func (o ServiceDocumentUploadDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uploadDestinationId"] = o.UploadDestinationId
	toSerialize["url"] = o.Url
	toSerialize["encryptionDetails"] = o.EncryptionDetails
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	return toSerialize, nil
}

type NullableServiceDocumentUploadDestination struct {
	value *ServiceDocumentUploadDestination
	isSet bool
}

func (v NullableServiceDocumentUploadDestination) Get() *ServiceDocumentUploadDestination {
	return v.value
}

func (v *NullableServiceDocumentUploadDestination) Set(val *ServiceDocumentUploadDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDocumentUploadDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDocumentUploadDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDocumentUploadDestination(val *ServiceDocumentUploadDestination) *NullableServiceDocumentUploadDestination {
	return &NullableServiceDocumentUploadDestination{value: val, isSet: true}
}

func (v NullableServiceDocumentUploadDestination) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableServiceDocumentUploadDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
