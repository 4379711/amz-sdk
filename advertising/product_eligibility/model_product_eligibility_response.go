package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductEligibilityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductEligibilityResponse{}

// ProductEligibilityResponse A product advertising eligibility response object.
type ProductEligibilityResponse struct {
	// A list of product advertising eligibility responses.
	ProductResponseList []ProductResponse `json:"productResponseList,omitempty"`
}

// NewProductEligibilityResponse instantiates a new ProductEligibilityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductEligibilityResponse() *ProductEligibilityResponse {
	this := ProductEligibilityResponse{}
	return &this
}

// NewProductEligibilityResponseWithDefaults instantiates a new ProductEligibilityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductEligibilityResponseWithDefaults() *ProductEligibilityResponse {
	this := ProductEligibilityResponse{}
	return &this
}

// GetProductResponseList returns the ProductResponseList field value if set, zero value otherwise.
func (o *ProductEligibilityResponse) GetProductResponseList() []ProductResponse {
	if o == nil || IsNil(o.ProductResponseList) {
		var ret []ProductResponse
		return ret
	}
	return o.ProductResponseList
}

// GetProductResponseListOk returns a tuple with the ProductResponseList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductEligibilityResponse) GetProductResponseListOk() ([]ProductResponse, bool) {
	if o == nil || IsNil(o.ProductResponseList) {
		return nil, false
	}
	return o.ProductResponseList, true
}

// HasProductResponseList returns a boolean if a field has been set.
func (o *ProductEligibilityResponse) HasProductResponseList() bool {
	if o != nil && !IsNil(o.ProductResponseList) {
		return true
	}

	return false
}

// SetProductResponseList gets a reference to the given []ProductResponse and assigns it to the ProductResponseList field.
func (o *ProductEligibilityResponse) SetProductResponseList(v []ProductResponse) {
	o.ProductResponseList = v
}

func (o ProductEligibilityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductResponseList) {
		toSerialize["productResponseList"] = o.ProductResponseList
	}
	return toSerialize, nil
}

type NullableProductEligibilityResponse struct {
	value *ProductEligibilityResponse
	isSet bool
}

func (v NullableProductEligibilityResponse) Get() *ProductEligibilityResponse {
	return v.value
}

func (v *NullableProductEligibilityResponse) Set(val *ProductEligibilityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductEligibilityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductEligibilityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductEligibilityResponse(val *ProductEligibilityResponse) *NullableProductEligibilityResponse {
	return &NullableProductEligibilityResponse{value: val, isSet: true}
}

func (v NullableProductEligibilityResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductEligibilityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
