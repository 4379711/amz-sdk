package data_kiosk_20231115

import (
	"github.com/bytedance/sonic"
)

// checks if the GetDocumentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDocumentResponse{}

// GetDocumentResponse The response for the `getDocument` operation.
type GetDocumentResponse struct {
	// The identifier for the Data Kiosk document. This identifier is unique only in combination with a selling partner account ID.
	DocumentId string `json:"documentId"`
	// A presigned URL that can be used to retrieve the Data Kiosk document. This URL expires after 5 minutes. If the Data Kiosk document is compressed, the `Content-Encoding` header will indicate the compression algorithm.  **Note:** Most HTTP clients are capable of automatically decompressing downloaded files based on the `Content-Encoding` header.
	DocumentUrl string `json:"documentUrl"`
}

type _GetDocumentResponse GetDocumentResponse

// NewGetDocumentResponse instantiates a new GetDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDocumentResponse(documentId string, documentUrl string) *GetDocumentResponse {
	this := GetDocumentResponse{}
	this.DocumentId = documentId
	this.DocumentUrl = documentUrl
	return &this
}

// NewGetDocumentResponseWithDefaults instantiates a new GetDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDocumentResponseWithDefaults() *GetDocumentResponse {
	this := GetDocumentResponse{}
	return &this
}

// GetDocumentId returns the DocumentId field value
func (o *GetDocumentResponse) GetDocumentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value
// and a boolean to check if the value has been set.
func (o *GetDocumentResponse) GetDocumentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentId, true
}

// SetDocumentId sets field value
func (o *GetDocumentResponse) SetDocumentId(v string) {
	o.DocumentId = v
}

// GetDocumentUrl returns the DocumentUrl field value
func (o *GetDocumentResponse) GetDocumentUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentUrl
}

// GetDocumentUrlOk returns a tuple with the DocumentUrl field value
// and a boolean to check if the value has been set.
func (o *GetDocumentResponse) GetDocumentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentUrl, true
}

// SetDocumentUrl sets field value
func (o *GetDocumentResponse) SetDocumentUrl(v string) {
	o.DocumentUrl = v
}

func (o GetDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["documentId"] = o.DocumentId
	toSerialize["documentUrl"] = o.DocumentUrl
	return toSerialize, nil
}

type NullableGetDocumentResponse struct {
	value *GetDocumentResponse
	isSet bool
}

func (v NullableGetDocumentResponse) Get() *GetDocumentResponse {
	return v.value
}

func (v *NullableGetDocumentResponse) Set(val *GetDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDocumentResponse(val *GetDocumentResponse) *NullableGetDocumentResponse {
	return &NullableGetDocumentResponse{value: val, isSet: true}
}

func (v NullableGetDocumentResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
