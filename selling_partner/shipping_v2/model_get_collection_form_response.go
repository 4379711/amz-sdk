package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the GetCollectionFormResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCollectionFormResponse{}

// GetCollectionFormResponse The Response  for the GetCollectionFormResponse operation.
type GetCollectionFormResponse struct {
	CollectionsFormDocument *CollectionsFormDocument `json:"collectionsFormDocument,omitempty"`
}

// NewGetCollectionFormResponse instantiates a new GetCollectionFormResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCollectionFormResponse() *GetCollectionFormResponse {
	this := GetCollectionFormResponse{}
	return &this
}

// NewGetCollectionFormResponseWithDefaults instantiates a new GetCollectionFormResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCollectionFormResponseWithDefaults() *GetCollectionFormResponse {
	this := GetCollectionFormResponse{}
	return &this
}

// GetCollectionsFormDocument returns the CollectionsFormDocument field value if set, zero value otherwise.
func (o *GetCollectionFormResponse) GetCollectionsFormDocument() CollectionsFormDocument {
	if o == nil || IsNil(o.CollectionsFormDocument) {
		var ret CollectionsFormDocument
		return ret
	}
	return *o.CollectionsFormDocument
}

// GetCollectionsFormDocumentOk returns a tuple with the CollectionsFormDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionFormResponse) GetCollectionsFormDocumentOk() (*CollectionsFormDocument, bool) {
	if o == nil || IsNil(o.CollectionsFormDocument) {
		return nil, false
	}
	return o.CollectionsFormDocument, true
}

// HasCollectionsFormDocument returns a boolean if a field has been set.
func (o *GetCollectionFormResponse) HasCollectionsFormDocument() bool {
	if o != nil && !IsNil(o.CollectionsFormDocument) {
		return true
	}

	return false
}

// SetCollectionsFormDocument gets a reference to the given CollectionsFormDocument and assigns it to the CollectionsFormDocument field.
func (o *GetCollectionFormResponse) SetCollectionsFormDocument(v CollectionsFormDocument) {
	o.CollectionsFormDocument = &v
}

func (o GetCollectionFormResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CollectionsFormDocument) {
		toSerialize["collectionsFormDocument"] = o.CollectionsFormDocument
	}
	return toSerialize, nil
}

type NullableGetCollectionFormResponse struct {
	value *GetCollectionFormResponse
	isSet bool
}

func (v NullableGetCollectionFormResponse) Get() *GetCollectionFormResponse {
	return v.value
}

func (v *NullableGetCollectionFormResponse) Set(val *GetCollectionFormResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCollectionFormResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCollectionFormResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCollectionFormResponse(val *GetCollectionFormResponse) *NullableGetCollectionFormResponse {
	return &NullableGetCollectionFormResponse{value: val, isSet: true}
}

func (v NullableGetCollectionFormResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetCollectionFormResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
