package listings_items_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the ListingsItemPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListingsItemPatchRequest{}

// ListingsItemPatchRequest The request body schema for the `patchListingsItem` operation.
type ListingsItemPatchRequest struct {
	// The Amazon product type of the listings item.
	ProductType string `json:"productType"`
	// One or more JSON Patch operations to perform on the listings item.
	Patches []PatchOperation `json:"patches"`
}

type _ListingsItemPatchRequest ListingsItemPatchRequest

// NewListingsItemPatchRequest instantiates a new ListingsItemPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListingsItemPatchRequest(productType string, patches []PatchOperation) *ListingsItemPatchRequest {
	this := ListingsItemPatchRequest{}
	this.ProductType = productType
	this.Patches = patches
	return &this
}

// NewListingsItemPatchRequestWithDefaults instantiates a new ListingsItemPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListingsItemPatchRequestWithDefaults() *ListingsItemPatchRequest {
	this := ListingsItemPatchRequest{}
	return &this
}

// GetProductType returns the ProductType field value
func (o *ListingsItemPatchRequest) GetProductType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *ListingsItemPatchRequest) GetProductTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *ListingsItemPatchRequest) SetProductType(v string) {
	o.ProductType = v
}

// GetPatches returns the Patches field value
func (o *ListingsItemPatchRequest) GetPatches() []PatchOperation {
	if o == nil {
		var ret []PatchOperation
		return ret
	}

	return o.Patches
}

// GetPatchesOk returns a tuple with the Patches field value
// and a boolean to check if the value has been set.
func (o *ListingsItemPatchRequest) GetPatchesOk() ([]PatchOperation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Patches, true
}

// SetPatches sets field value
func (o *ListingsItemPatchRequest) SetPatches(v []PatchOperation) {
	o.Patches = v
}

func (o ListingsItemPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productType"] = o.ProductType
	toSerialize["patches"] = o.Patches
	return toSerialize, nil
}

type NullableListingsItemPatchRequest struct {
	value *ListingsItemPatchRequest
	isSet bool
}

func (v NullableListingsItemPatchRequest) Get() *ListingsItemPatchRequest {
	return v.value
}

func (v *NullableListingsItemPatchRequest) Set(val *ListingsItemPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListingsItemPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListingsItemPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListingsItemPatchRequest(val *ListingsItemPatchRequest) *NullableListingsItemPatchRequest {
	return &NullableListingsItemPatchRequest{value: val, isSet: true}
}

func (v NullableListingsItemPatchRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListingsItemPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
