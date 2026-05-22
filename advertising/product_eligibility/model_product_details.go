package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductDetails{}

// ProductDetails An Amazon product identifier, seller product identifier, or both.
type ProductDetails struct {
	GlobalStoreSetting *GlobalStoreSetting `json:"globalStoreSetting,omitempty"`
	// An Amazon product identifier.
	Asin string `json:"asin"`
	// A seller product identifier.
	Sku *string `json:"sku,omitempty"`
}

type _ProductDetails ProductDetails

// NewProductDetails instantiates a new ProductDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductDetails(asin string) *ProductDetails {
	this := ProductDetails{}
	this.Asin = asin
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
func (o *ProductDetails) GetGlobalStoreSetting() GlobalStoreSetting {
	if o == nil || IsNil(o.GlobalStoreSetting) {
		var ret GlobalStoreSetting
		return ret
	}
	return *o.GlobalStoreSetting
}

// GetGlobalStoreSettingOk returns a tuple with the GlobalStoreSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetails) GetGlobalStoreSettingOk() (*GlobalStoreSetting, bool) {
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

// SetGlobalStoreSetting gets a reference to the given GlobalStoreSetting and assigns it to the GlobalStoreSetting field.
func (o *ProductDetails) SetGlobalStoreSetting(v GlobalStoreSetting) {
	o.GlobalStoreSetting = &v
}

// GetAsin returns the Asin field value
func (o *ProductDetails) GetAsin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asin
}

// GetAsinOk returns a tuple with the Asin field value
// and a boolean to check if the value has been set.
func (o *ProductDetails) GetAsinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asin, true
}

// SetAsin sets field value
func (o *ProductDetails) SetAsin(v string) {
	o.Asin = v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *ProductDetails) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetails) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *ProductDetails) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *ProductDetails) SetSku(v string) {
	o.Sku = &v
}

func (o ProductDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GlobalStoreSetting) {
		toSerialize["globalStoreSetting"] = o.GlobalStoreSetting
	}
	toSerialize["asin"] = o.Asin
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
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
