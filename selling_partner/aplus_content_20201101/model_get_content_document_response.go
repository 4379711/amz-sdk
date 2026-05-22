package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the GetContentDocumentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetContentDocumentResponse{}

// GetContentDocumentResponse struct for GetContentDocumentResponse
type GetContentDocumentResponse struct {
	// A set of messages to the user, such as warnings or comments.
	Warnings      []Error       `json:"warnings,omitempty"`
	ContentRecord ContentRecord `json:"contentRecord"`
}

type _GetContentDocumentResponse GetContentDocumentResponse

// NewGetContentDocumentResponse instantiates a new GetContentDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetContentDocumentResponse(contentRecord ContentRecord) *GetContentDocumentResponse {
	this := GetContentDocumentResponse{}
	this.ContentRecord = contentRecord
	return &this
}

// NewGetContentDocumentResponseWithDefaults instantiates a new GetContentDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetContentDocumentResponseWithDefaults() *GetContentDocumentResponse {
	this := GetContentDocumentResponse{}
	return &this
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *GetContentDocumentResponse) GetWarnings() []Error {
	if o == nil || IsNil(o.Warnings) {
		var ret []Error
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContentDocumentResponse) GetWarningsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *GetContentDocumentResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Error and assigns it to the Warnings field.
func (o *GetContentDocumentResponse) SetWarnings(v []Error) {
	o.Warnings = v
}

// GetContentRecord returns the ContentRecord field value
func (o *GetContentDocumentResponse) GetContentRecord() ContentRecord {
	if o == nil {
		var ret ContentRecord
		return ret
	}

	return o.ContentRecord
}

// GetContentRecordOk returns a tuple with the ContentRecord field value
// and a boolean to check if the value has been set.
func (o *GetContentDocumentResponse) GetContentRecordOk() (*ContentRecord, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentRecord, true
}

// SetContentRecord sets field value
func (o *GetContentDocumentResponse) SetContentRecord(v ContentRecord) {
	o.ContentRecord = v
}

func (o GetContentDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	toSerialize["contentRecord"] = o.ContentRecord
	return toSerialize, nil
}

type NullableGetContentDocumentResponse struct {
	value *GetContentDocumentResponse
	isSet bool
}

func (v NullableGetContentDocumentResponse) Get() *GetContentDocumentResponse {
	return v.value
}

func (v *NullableGetContentDocumentResponse) Set(val *GetContentDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetContentDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetContentDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetContentDocumentResponse(val *GetContentDocumentResponse) *NullableGetContentDocumentResponse {
	return &NullableGetContentDocumentResponse{value: val, isSet: true}
}

func (v NullableGetContentDocumentResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetContentDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
