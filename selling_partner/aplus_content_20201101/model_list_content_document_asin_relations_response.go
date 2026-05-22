package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the ListContentDocumentAsinRelationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListContentDocumentAsinRelationsResponse{}

// ListContentDocumentAsinRelationsResponse struct for ListContentDocumentAsinRelationsResponse
type ListContentDocumentAsinRelationsResponse struct {
	// A set of messages to the user, such as warnings or comments.
	Warnings []Error `json:"warnings,omitempty"`
	// A token that you use to fetch a specific page when there are multiple pages of results.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// The set of ASIN metadata.
	AsinMetadataSet []AsinMetadata `json:"asinMetadataSet"`
}

type _ListContentDocumentAsinRelationsResponse ListContentDocumentAsinRelationsResponse

// NewListContentDocumentAsinRelationsResponse instantiates a new ListContentDocumentAsinRelationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListContentDocumentAsinRelationsResponse(asinMetadataSet []AsinMetadata) *ListContentDocumentAsinRelationsResponse {
	this := ListContentDocumentAsinRelationsResponse{}
	this.AsinMetadataSet = asinMetadataSet
	return &this
}

// NewListContentDocumentAsinRelationsResponseWithDefaults instantiates a new ListContentDocumentAsinRelationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListContentDocumentAsinRelationsResponseWithDefaults() *ListContentDocumentAsinRelationsResponse {
	this := ListContentDocumentAsinRelationsResponse{}
	return &this
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ListContentDocumentAsinRelationsResponse) GetWarnings() []Error {
	if o == nil || IsNil(o.Warnings) {
		var ret []Error
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListContentDocumentAsinRelationsResponse) GetWarningsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ListContentDocumentAsinRelationsResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Error and assigns it to the Warnings field.
func (o *ListContentDocumentAsinRelationsResponse) SetWarnings(v []Error) {
	o.Warnings = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListContentDocumentAsinRelationsResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListContentDocumentAsinRelationsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ListContentDocumentAsinRelationsResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListContentDocumentAsinRelationsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetAsinMetadataSet returns the AsinMetadataSet field value
func (o *ListContentDocumentAsinRelationsResponse) GetAsinMetadataSet() []AsinMetadata {
	if o == nil {
		var ret []AsinMetadata
		return ret
	}

	return o.AsinMetadataSet
}

// GetAsinMetadataSetOk returns a tuple with the AsinMetadataSet field value
// and a boolean to check if the value has been set.
func (o *ListContentDocumentAsinRelationsResponse) GetAsinMetadataSetOk() ([]AsinMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.AsinMetadataSet, true
}

// SetAsinMetadataSet sets field value
func (o *ListContentDocumentAsinRelationsResponse) SetAsinMetadataSet(v []AsinMetadata) {
	o.AsinMetadataSet = v
}

func (o ListContentDocumentAsinRelationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	toSerialize["asinMetadataSet"] = o.AsinMetadataSet
	return toSerialize, nil
}

type NullableListContentDocumentAsinRelationsResponse struct {
	value *ListContentDocumentAsinRelationsResponse
	isSet bool
}

func (v NullableListContentDocumentAsinRelationsResponse) Get() *ListContentDocumentAsinRelationsResponse {
	return v.value
}

func (v *NullableListContentDocumentAsinRelationsResponse) Set(val *ListContentDocumentAsinRelationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListContentDocumentAsinRelationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListContentDocumentAsinRelationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListContentDocumentAsinRelationsResponse(val *ListContentDocumentAsinRelationsResponse) *NullableListContentDocumentAsinRelationsResponse {
	return &NullableListContentDocumentAsinRelationsResponse{value: val, isSet: true}
}

func (v NullableListContentDocumentAsinRelationsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListContentDocumentAsinRelationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
