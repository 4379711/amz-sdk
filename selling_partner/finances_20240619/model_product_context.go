package finances_20240619

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductContext{}

// ProductContext Additional information related to the product.
type ProductContext struct {
	// Amazon Standard Identification Number (ASIN) of the item.
	Asin *string `json:"asin,omitempty"`
	// Stock keeping unit (SKU) of the item.
	Sku *string `json:"sku,omitempty"`
	// Quantity of the item shipped.
	QuantityShipped *int32 `json:"quantityShipped,omitempty"`
	// Fulfillment network of the item.
	FulfillmentNetwork *string `json:"fulfillmentNetwork,omitempty"`
}

// NewProductContext instantiates a new ProductContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductContext() *ProductContext {
	this := ProductContext{}
	return &this
}

// NewProductContextWithDefaults instantiates a new ProductContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductContextWithDefaults() *ProductContext {
	this := ProductContext{}
	return &this
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *ProductContext) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductContext) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *ProductContext) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *ProductContext) SetAsin(v string) {
	o.Asin = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *ProductContext) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductContext) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *ProductContext) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *ProductContext) SetSku(v string) {
	o.Sku = &v
}

// GetQuantityShipped returns the QuantityShipped field value if set, zero value otherwise.
func (o *ProductContext) GetQuantityShipped() int32 {
	if o == nil || IsNil(o.QuantityShipped) {
		var ret int32
		return ret
	}
	return *o.QuantityShipped
}

// GetQuantityShippedOk returns a tuple with the QuantityShipped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductContext) GetQuantityShippedOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityShipped) {
		return nil, false
	}
	return o.QuantityShipped, true
}

// HasQuantityShipped returns a boolean if a field has been set.
func (o *ProductContext) HasQuantityShipped() bool {
	if o != nil && !IsNil(o.QuantityShipped) {
		return true
	}

	return false
}

// SetQuantityShipped gets a reference to the given int32 and assigns it to the QuantityShipped field.
func (o *ProductContext) SetQuantityShipped(v int32) {
	o.QuantityShipped = &v
}

// GetFulfillmentNetwork returns the FulfillmentNetwork field value if set, zero value otherwise.
func (o *ProductContext) GetFulfillmentNetwork() string {
	if o == nil || IsNil(o.FulfillmentNetwork) {
		var ret string
		return ret
	}
	return *o.FulfillmentNetwork
}

// GetFulfillmentNetworkOk returns a tuple with the FulfillmentNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductContext) GetFulfillmentNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.FulfillmentNetwork) {
		return nil, false
	}
	return o.FulfillmentNetwork, true
}

// HasFulfillmentNetwork returns a boolean if a field has been set.
func (o *ProductContext) HasFulfillmentNetwork() bool {
	if o != nil && !IsNil(o.FulfillmentNetwork) {
		return true
	}

	return false
}

// SetFulfillmentNetwork gets a reference to the given string and assigns it to the FulfillmentNetwork field.
func (o *ProductContext) SetFulfillmentNetwork(v string) {
	o.FulfillmentNetwork = &v
}

func (o ProductContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.QuantityShipped) {
		toSerialize["quantityShipped"] = o.QuantityShipped
	}
	if !IsNil(o.FulfillmentNetwork) {
		toSerialize["fulfillmentNetwork"] = o.FulfillmentNetwork
	}
	return toSerialize, nil
}

type NullableProductContext struct {
	value *ProductContext
	isSet bool
}

func (v NullableProductContext) Get() *ProductContext {
	return v.value
}

func (v *NullableProductContext) Set(val *ProductContext) {
	v.value = val
	v.isSet = true
}

func (v NullableProductContext) IsSet() bool {
	return v.isSet
}

func (v *NullableProductContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductContext(val *ProductContext) *NullableProductContext {
	return &NullableProductContext{value: val, isSet: true}
}

func (v NullableProductContext) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
