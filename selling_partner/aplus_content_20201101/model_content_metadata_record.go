package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the ContentMetadataRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentMetadataRecord{}

// ContentMetadataRecord The metadata for an A+ Content document, with additional information for content management.
type ContentMetadataRecord struct {
	// A unique reference key for the A+ Content document. A content reference key cannot form a permalink and might change in the future. A content reference key is not guaranteed to match any A+ content identifier.
	ContentReferenceKey string          `json:"contentReferenceKey"`
	ContentMetadata     ContentMetadata `json:"contentMetadata"`
}

type _ContentMetadataRecord ContentMetadataRecord

// NewContentMetadataRecord instantiates a new ContentMetadataRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentMetadataRecord(contentReferenceKey string, contentMetadata ContentMetadata) *ContentMetadataRecord {
	this := ContentMetadataRecord{}
	this.ContentReferenceKey = contentReferenceKey
	this.ContentMetadata = contentMetadata
	return &this
}

// NewContentMetadataRecordWithDefaults instantiates a new ContentMetadataRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentMetadataRecordWithDefaults() *ContentMetadataRecord {
	this := ContentMetadataRecord{}
	return &this
}

// GetContentReferenceKey returns the ContentReferenceKey field value
func (o *ContentMetadataRecord) GetContentReferenceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentReferenceKey
}

// GetContentReferenceKeyOk returns a tuple with the ContentReferenceKey field value
// and a boolean to check if the value has been set.
func (o *ContentMetadataRecord) GetContentReferenceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentReferenceKey, true
}

// SetContentReferenceKey sets field value
func (o *ContentMetadataRecord) SetContentReferenceKey(v string) {
	o.ContentReferenceKey = v
}

// GetContentMetadata returns the ContentMetadata field value
func (o *ContentMetadataRecord) GetContentMetadata() ContentMetadata {
	if o == nil {
		var ret ContentMetadata
		return ret
	}

	return o.ContentMetadata
}

// GetContentMetadataOk returns a tuple with the ContentMetadata field value
// and a boolean to check if the value has been set.
func (o *ContentMetadataRecord) GetContentMetadataOk() (*ContentMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentMetadata, true
}

// SetContentMetadata sets field value
func (o *ContentMetadataRecord) SetContentMetadata(v ContentMetadata) {
	o.ContentMetadata = v
}

func (o ContentMetadataRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contentReferenceKey"] = o.ContentReferenceKey
	toSerialize["contentMetadata"] = o.ContentMetadata
	return toSerialize, nil
}

type NullableContentMetadataRecord struct {
	value *ContentMetadataRecord
	isSet bool
}

func (v NullableContentMetadataRecord) Get() *ContentMetadataRecord {
	return v.value
}

func (v *NullableContentMetadataRecord) Set(val *ContentMetadataRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableContentMetadataRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableContentMetadataRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentMetadataRecord(val *ContentMetadataRecord) *NullableContentMetadataRecord {
	return &NullableContentMetadataRecord{value: val, isSet: true}
}

func (v NullableContentMetadataRecord) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableContentMetadataRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
