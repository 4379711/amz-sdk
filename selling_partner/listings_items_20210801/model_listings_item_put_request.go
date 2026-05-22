package listings_items_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the ListingsItemPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListingsItemPutRequest{}

// ListingsItemPutRequest The request body schema for the `putListingsItem` operation.
type ListingsItemPutRequest struct {
	// The Amazon product type of the listings item.
	ProductType string `json:"productType"`
	// The name of the requirements set for the provided data.
	Requirements *string `json:"requirements,omitempty"`
	// A JSON object containing structured listings item attribute data keyed by attribute name.
	Attributes map[string]interface{} `json:"attributes"`
}

type _ListingsItemPutRequest ListingsItemPutRequest

// NewListingsItemPutRequest instantiates a new ListingsItemPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListingsItemPutRequest(productType string, attributes map[string]interface{}) *ListingsItemPutRequest {
	this := ListingsItemPutRequest{}
	this.ProductType = productType
	this.Attributes = attributes
	return &this
}

// NewListingsItemPutRequestWithDefaults instantiates a new ListingsItemPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListingsItemPutRequestWithDefaults() *ListingsItemPutRequest {
	this := ListingsItemPutRequest{}
	return &this
}

// GetProductType returns the ProductType field value
func (o *ListingsItemPutRequest) GetProductType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *ListingsItemPutRequest) GetProductTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *ListingsItemPutRequest) SetProductType(v string) {
	o.ProductType = v
}

// GetRequirements returns the Requirements field value if set, zero value otherwise.
func (o *ListingsItemPutRequest) GetRequirements() string {
	if o == nil || IsNil(o.Requirements) {
		var ret string
		return ret
	}
	return *o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingsItemPutRequest) GetRequirementsOk() (*string, bool) {
	if o == nil || IsNil(o.Requirements) {
		return nil, false
	}
	return o.Requirements, true
}

// HasRequirements returns a boolean if a field has been set.
func (o *ListingsItemPutRequest) HasRequirements() bool {
	if o != nil && !IsNil(o.Requirements) {
		return true
	}

	return false
}

// SetRequirements gets a reference to the given string and assigns it to the Requirements field.
func (o *ListingsItemPutRequest) SetRequirements(v string) {
	o.Requirements = &v
}

// GetAttributes returns the Attributes field value
func (o *ListingsItemPutRequest) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ListingsItemPutRequest) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *ListingsItemPutRequest) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o ListingsItemPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productType"] = o.ProductType
	if !IsNil(o.Requirements) {
		toSerialize["requirements"] = o.Requirements
	}
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

type NullableListingsItemPutRequest struct {
	value *ListingsItemPutRequest
	isSet bool
}

func (v NullableListingsItemPutRequest) Get() *ListingsItemPutRequest {
	return v.value
}

func (v *NullableListingsItemPutRequest) Set(val *ListingsItemPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListingsItemPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListingsItemPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListingsItemPutRequest(val *ListingsItemPutRequest) *NullableListingsItemPutRequest {
	return &NullableListingsItemPutRequest{value: val, isSet: true}
}

func (v NullableListingsItemPutRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListingsItemPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
