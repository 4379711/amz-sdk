package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the ReferencePrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReferencePrice{}

// ReferencePrice The reference price for the specified ASIN `marketplaceId` combination.
type ReferencePrice struct {
	// The name of the reference price, such as `CompetitivePriceThreshold` and `WasPrice`. For reference price definitions, refer to the [Use Case Guide](https://developer-docs.amazon.com/sp-api/docs/product-pricing-api-v2022-05-01-use-case-guide).
	Name  string    `json:"name"`
	Price MoneyType `json:"price"`
}

type _ReferencePrice ReferencePrice

// NewReferencePrice instantiates a new ReferencePrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferencePrice(name string, price MoneyType) *ReferencePrice {
	this := ReferencePrice{}
	this.Name = name
	this.Price = price
	return &this
}

// NewReferencePriceWithDefaults instantiates a new ReferencePrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferencePriceWithDefaults() *ReferencePrice {
	this := ReferencePrice{}
	return &this
}

// GetName returns the Name field value
func (o *ReferencePrice) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReferencePrice) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReferencePrice) SetName(v string) {
	o.Name = v
}

// GetPrice returns the Price field value
func (o *ReferencePrice) GetPrice() MoneyType {
	if o == nil {
		var ret MoneyType
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *ReferencePrice) GetPriceOk() (*MoneyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *ReferencePrice) SetPrice(v MoneyType) {
	o.Price = v
}

func (o ReferencePrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["price"] = o.Price
	return toSerialize, nil
}

type NullableReferencePrice struct {
	value *ReferencePrice
	isSet bool
}

func (v NullableReferencePrice) Get() *ReferencePrice {
	return v.value
}

func (v *NullableReferencePrice) Set(val *ReferencePrice) {
	v.value = val
	v.isSet = true
}

func (v NullableReferencePrice) IsSet() bool {
	return v.isSet
}

func (v *NullableReferencePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferencePrice(val *ReferencePrice) *NullableReferencePrice {
	return &NullableReferencePrice{value: val, isSet: true}
}

func (v NullableReferencePrice) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableReferencePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
