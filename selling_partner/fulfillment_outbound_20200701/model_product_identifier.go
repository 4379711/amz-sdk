package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductIdentifier{}

// ProductIdentifier Product identifier input that locates a product for MCF.
type ProductIdentifier struct {
	// The merchant SKU for the product.
	MerchantSku string `json:"merchantSku"`
}

type _ProductIdentifier ProductIdentifier

// NewProductIdentifier instantiates a new ProductIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductIdentifier(merchantSku string) *ProductIdentifier {
	this := ProductIdentifier{}
	this.MerchantSku = merchantSku
	return &this
}

// NewProductIdentifierWithDefaults instantiates a new ProductIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductIdentifierWithDefaults() *ProductIdentifier {
	this := ProductIdentifier{}
	return &this
}

// GetMerchantSku returns the MerchantSku field value
func (o *ProductIdentifier) GetMerchantSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantSku
}

// GetMerchantSkuOk returns a tuple with the MerchantSku field value
// and a boolean to check if the value has been set.
func (o *ProductIdentifier) GetMerchantSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantSku, true
}

// SetMerchantSku sets field value
func (o *ProductIdentifier) SetMerchantSku(v string) {
	o.MerchantSku = v
}

func (o ProductIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantSku"] = o.MerchantSku
	return toSerialize, nil
}

type NullableProductIdentifier struct {
	value *ProductIdentifier
	isSet bool
}

func (v NullableProductIdentifier) Get() *ProductIdentifier {
	return v.value
}

func (v *NullableProductIdentifier) Set(val *ProductIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableProductIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableProductIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductIdentifier(val *ProductIdentifier) *NullableProductIdentifier {
	return &NullableProductIdentifier{value: val, isSet: true}
}

func (v NullableProductIdentifier) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
