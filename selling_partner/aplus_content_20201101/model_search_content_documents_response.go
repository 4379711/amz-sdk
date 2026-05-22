package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the SearchContentDocumentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchContentDocumentsResponse{}

// SearchContentDocumentsResponse struct for SearchContentDocumentsResponse
type SearchContentDocumentsResponse struct {
	// A set of messages to the user, such as warnings or comments.
	Warnings []Error `json:"warnings,omitempty"`
	// A token that you use to fetch a specific page when there are multiple pages of results.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// A list of A+ Content metadata records.
	ContentMetadataRecords []ContentMetadataRecord `json:"contentMetadataRecords"`
}

type _SearchContentDocumentsResponse SearchContentDocumentsResponse

// NewSearchContentDocumentsResponse instantiates a new SearchContentDocumentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchContentDocumentsResponse(contentMetadataRecords []ContentMetadataRecord) *SearchContentDocumentsResponse {
	this := SearchContentDocumentsResponse{}
	this.ContentMetadataRecords = contentMetadataRecords
	return &this
}

// NewSearchContentDocumentsResponseWithDefaults instantiates a new SearchContentDocumentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchContentDocumentsResponseWithDefaults() *SearchContentDocumentsResponse {
	this := SearchContentDocumentsResponse{}
	return &this
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *SearchContentDocumentsResponse) GetWarnings() []Error {
	if o == nil || IsNil(o.Warnings) {
		var ret []Error
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchContentDocumentsResponse) GetWarningsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *SearchContentDocumentsResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Error and assigns it to the Warnings field.
func (o *SearchContentDocumentsResponse) SetWarnings(v []Error) {
	o.Warnings = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *SearchContentDocumentsResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchContentDocumentsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *SearchContentDocumentsResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *SearchContentDocumentsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetContentMetadataRecords returns the ContentMetadataRecords field value
func (o *SearchContentDocumentsResponse) GetContentMetadataRecords() []ContentMetadataRecord {
	if o == nil {
		var ret []ContentMetadataRecord
		return ret
	}

	return o.ContentMetadataRecords
}

// GetContentMetadataRecordsOk returns a tuple with the ContentMetadataRecords field value
// and a boolean to check if the value has been set.
func (o *SearchContentDocumentsResponse) GetContentMetadataRecordsOk() ([]ContentMetadataRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentMetadataRecords, true
}

// SetContentMetadataRecords sets field value
func (o *SearchContentDocumentsResponse) SetContentMetadataRecords(v []ContentMetadataRecord) {
	o.ContentMetadataRecords = v
}

func (o SearchContentDocumentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	toSerialize["contentMetadataRecords"] = o.ContentMetadataRecords
	return toSerialize, nil
}

type NullableSearchContentDocumentsResponse struct {
	value *SearchContentDocumentsResponse
	isSet bool
}

func (v NullableSearchContentDocumentsResponse) Get() *SearchContentDocumentsResponse {
	return v.value
}

func (v *NullableSearchContentDocumentsResponse) Set(val *SearchContentDocumentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchContentDocumentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchContentDocumentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchContentDocumentsResponse(val *SearchContentDocumentsResponse) *NullableSearchContentDocumentsResponse {
	return &NullableSearchContentDocumentsResponse{value: val, isSet: true}
}

func (v NullableSearchContentDocumentsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSearchContentDocumentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
