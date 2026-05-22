package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the GenerateCollectionFormResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateCollectionFormResponse{}

// GenerateCollectionFormResponse The Response  for the GenerateCollectionFormResponse operation.
type GenerateCollectionFormResponse struct {
	CollectionsFormDocument *CollectionsFormDocument `json:"collectionsFormDocument,omitempty"`
}

// NewGenerateCollectionFormResponse instantiates a new GenerateCollectionFormResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateCollectionFormResponse() *GenerateCollectionFormResponse {
	this := GenerateCollectionFormResponse{}
	return &this
}

// NewGenerateCollectionFormResponseWithDefaults instantiates a new GenerateCollectionFormResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateCollectionFormResponseWithDefaults() *GenerateCollectionFormResponse {
	this := GenerateCollectionFormResponse{}
	return &this
}

// GetCollectionsFormDocument returns the CollectionsFormDocument field value if set, zero value otherwise.
func (o *GenerateCollectionFormResponse) GetCollectionsFormDocument() CollectionsFormDocument {
	if o == nil || IsNil(o.CollectionsFormDocument) {
		var ret CollectionsFormDocument
		return ret
	}
	return *o.CollectionsFormDocument
}

// GetCollectionsFormDocumentOk returns a tuple with the CollectionsFormDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateCollectionFormResponse) GetCollectionsFormDocumentOk() (*CollectionsFormDocument, bool) {
	if o == nil || IsNil(o.CollectionsFormDocument) {
		return nil, false
	}
	return o.CollectionsFormDocument, true
}

// HasCollectionsFormDocument returns a boolean if a field has been set.
func (o *GenerateCollectionFormResponse) HasCollectionsFormDocument() bool {
	if o != nil && !IsNil(o.CollectionsFormDocument) {
		return true
	}

	return false
}

// SetCollectionsFormDocument gets a reference to the given CollectionsFormDocument and assigns it to the CollectionsFormDocument field.
func (o *GenerateCollectionFormResponse) SetCollectionsFormDocument(v CollectionsFormDocument) {
	o.CollectionsFormDocument = &v
}

func (o GenerateCollectionFormResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CollectionsFormDocument) {
		toSerialize["collectionsFormDocument"] = o.CollectionsFormDocument
	}
	return toSerialize, nil
}

type NullableGenerateCollectionFormResponse struct {
	value *GenerateCollectionFormResponse
	isSet bool
}

func (v NullableGenerateCollectionFormResponse) Get() *GenerateCollectionFormResponse {
	return v.value
}

func (v *NullableGenerateCollectionFormResponse) Set(val *GenerateCollectionFormResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateCollectionFormResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateCollectionFormResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateCollectionFormResponse(val *GenerateCollectionFormResponse) *NullableGenerateCollectionFormResponse {
	return &NullableGenerateCollectionFormResponse{value: val, isSet: true}
}

func (v NullableGenerateCollectionFormResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGenerateCollectionFormResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
