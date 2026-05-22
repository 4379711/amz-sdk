package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductDetails{}

// ProductDetails struct for ProductDetails
type ProductDetails struct {
	GlobalStoreSetting *ProductDetailsGlobalStoreSetting `json:"globalStoreSetting,omitempty"`
	// The identifier of the product.
	Asin *string `json:"asin,omitempty"`
}

// NewProductDetails instantiates a new ProductDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductDetails() *ProductDetails {
	this := ProductDetails{}
	return &this
}

// NewProductDetailsWithDefaults instantiates a new ProductDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductDetailsWithDefaults() *ProductDetails {
	this := ProductDetails{}
	return &this
}

// GetGlobalStoreSetting returns the GlobalStoreSetting field value if set, zero value otherwise.
func (o *ProductDetails) GetGlobalStoreSetting() ProductDetailsGlobalStoreSetting {
	if o == nil || IsNil(o.GlobalStoreSetting) {
		var ret ProductDetailsGlobalStoreSetting
		return ret
	}
	return *o.GlobalStoreSetting
}

// GetGlobalStoreSettingOk returns a tuple with the GlobalStoreSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetails) GetGlobalStoreSettingOk() (*ProductDetailsGlobalStoreSetting, bool) {
	if o == nil || IsNil(o.GlobalStoreSetting) {
		return nil, false
	}
	return o.GlobalStoreSetting, true
}

// HasGlobalStoreSetting returns a boolean if a field has been set.
func (o *ProductDetails) HasGlobalStoreSetting() bool {
	if o != nil && !IsNil(o.GlobalStoreSetting) {
		return true
	}

	return false
}

// SetGlobalStoreSetting gets a reference to the given ProductDetailsGlobalStoreSetting and assigns it to the GlobalStoreSetting field.
func (o *ProductDetails) SetGlobalStoreSetting(v ProductDetailsGlobalStoreSetting) {
	o.GlobalStoreSetting = &v
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *ProductDetails) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetails) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *ProductDetails) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *ProductDetails) SetAsin(v string) {
	o.Asin = &v
}

func (o ProductDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GlobalStoreSetting) {
		toSerialize["globalStoreSetting"] = o.GlobalStoreSetting
	}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	return toSerialize, nil
}

type NullableProductDetails struct {
	value *ProductDetails
	isSet bool
}

func (v NullableProductDetails) Get() *ProductDetails {
	return v.value
}

func (v *NullableProductDetails) Set(val *ProductDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableProductDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableProductDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductDetails(val *ProductDetails) *NullableProductDetails {
	return &NullableProductDetails{value: val, isSet: true}
}

func (v NullableProductDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
