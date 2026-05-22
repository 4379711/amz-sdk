package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductInfoDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductInfoDetail{}

// ProductInfoDetail Product information on the number of items.
type ProductInfoDetail struct {
	// The total number of items that are included in the ASIN.
	NumberOfItems *string `json:"NumberOfItems,omitempty"`
}

// NewProductInfoDetail instantiates a new ProductInfoDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductInfoDetail() *ProductInfoDetail {
	this := ProductInfoDetail{}
	return &this
}

// NewProductInfoDetailWithDefaults instantiates a new ProductInfoDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductInfoDetailWithDefaults() *ProductInfoDetail {
	this := ProductInfoDetail{}
	return &this
}

// GetNumberOfItems returns the NumberOfItems field value if set, zero value otherwise.
func (o *ProductInfoDetail) GetNumberOfItems() string {
	if o == nil || IsNil(o.NumberOfItems) {
		var ret string
		return ret
	}
	return *o.NumberOfItems
}

// GetNumberOfItemsOk returns a tuple with the NumberOfItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoDetail) GetNumberOfItemsOk() (*string, bool) {
	if o == nil || IsNil(o.NumberOfItems) {
		return nil, false
	}
	return o.NumberOfItems, true
}

// HasNumberOfItems returns a boolean if a field has been set.
func (o *ProductInfoDetail) HasNumberOfItems() bool {
	if o != nil && !IsNil(o.NumberOfItems) {
		return true
	}

	return false
}

// SetNumberOfItems gets a reference to the given string and assigns it to the NumberOfItems field.
func (o *ProductInfoDetail) SetNumberOfItems(v string) {
	o.NumberOfItems = &v
}

func (o ProductInfoDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumberOfItems) {
		toSerialize["NumberOfItems"] = o.NumberOfItems
	}
	return toSerialize, nil
}

type NullableProductInfoDetail struct {
	value *ProductInfoDetail
	isSet bool
}

func (v NullableProductInfoDetail) Get() *ProductInfoDetail {
	return v.value
}

func (v *NullableProductInfoDetail) Set(val *ProductInfoDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableProductInfoDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableProductInfoDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductInfoDetail(val *ProductInfoDetail) *NullableProductInfoDetail {
	return &NullableProductInfoDetail{value: val, isSet: true}
}

func (v NullableProductInfoDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductInfoDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
