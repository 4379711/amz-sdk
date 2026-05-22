package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the PostContentDocumentAsinRelationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostContentDocumentAsinRelationsRequest{}

// PostContentDocumentAsinRelationsRequest struct for PostContentDocumentAsinRelationsRequest
type PostContentDocumentAsinRelationsRequest struct {
	// The set of ASINs.
	AsinSet []string `json:"asinSet"`
}

type _PostContentDocumentAsinRelationsRequest PostContentDocumentAsinRelationsRequest

// NewPostContentDocumentAsinRelationsRequest instantiates a new PostContentDocumentAsinRelationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostContentDocumentAsinRelationsRequest(asinSet []string) *PostContentDocumentAsinRelationsRequest {
	this := PostContentDocumentAsinRelationsRequest{}
	this.AsinSet = asinSet
	return &this
}

// NewPostContentDocumentAsinRelationsRequestWithDefaults instantiates a new PostContentDocumentAsinRelationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostContentDocumentAsinRelationsRequestWithDefaults() *PostContentDocumentAsinRelationsRequest {
	this := PostContentDocumentAsinRelationsRequest{}
	return &this
}

// GetAsinSet returns the AsinSet field value
func (o *PostContentDocumentAsinRelationsRequest) GetAsinSet() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AsinSet
}

// GetAsinSetOk returns a tuple with the AsinSet field value
// and a boolean to check if the value has been set.
func (o *PostContentDocumentAsinRelationsRequest) GetAsinSetOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AsinSet, true
}

// SetAsinSet sets field value
func (o *PostContentDocumentAsinRelationsRequest) SetAsinSet(v []string) {
	o.AsinSet = v
}

func (o PostContentDocumentAsinRelationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asinSet"] = o.AsinSet
	return toSerialize, nil
}

type NullablePostContentDocumentAsinRelationsRequest struct {
	value *PostContentDocumentAsinRelationsRequest
	isSet bool
}

func (v NullablePostContentDocumentAsinRelationsRequest) Get() *PostContentDocumentAsinRelationsRequest {
	return v.value
}

func (v *NullablePostContentDocumentAsinRelationsRequest) Set(val *PostContentDocumentAsinRelationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostContentDocumentAsinRelationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostContentDocumentAsinRelationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostContentDocumentAsinRelationsRequest(val *PostContentDocumentAsinRelationsRequest) *NullablePostContentDocumentAsinRelationsRequest {
	return &NullablePostContentDocumentAsinRelationsRequest{value: val, isSet: true}
}

func (v NullablePostContentDocumentAsinRelationsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePostContentDocumentAsinRelationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
