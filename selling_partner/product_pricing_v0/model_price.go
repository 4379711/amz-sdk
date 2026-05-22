package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the Price type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Price{}

// Price Schema for price info in `getPricing` response
type Price struct {
	// The status of the operation.
	Status string `json:"status"`
	// The seller stock keeping unit (SKU) of the item.
	SellerSKU *string `json:"SellerSKU,omitempty"`
	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN    *string  `json:"ASIN,omitempty"`
	Product *Product `json:"Product,omitempty"`
}

type _Price Price

// NewPrice instantiates a new Price object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrice(status string) *Price {
	this := Price{}
	this.Status = status
	return &this
}

// NewPriceWithDefaults instantiates a new Price object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceWithDefaults() *Price {
	this := Price{}
	return &this
}

// GetStatus returns the Status field value
func (o *Price) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Price) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Price) SetStatus(v string) {
	o.Status = v
}

// GetSellerSKU returns the SellerSKU field value if set, zero value otherwise.
func (o *Price) GetSellerSKU() string {
	if o == nil || IsNil(o.SellerSKU) {
		var ret string
		return ret
	}
	return *o.SellerSKU
}

// GetSellerSKUOk returns a tuple with the SellerSKU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetSellerSKUOk() (*string, bool) {
	if o == nil || IsNil(o.SellerSKU) {
		return nil, false
	}
	return o.SellerSKU, true
}

// HasSellerSKU returns a boolean if a field has been set.
func (o *Price) HasSellerSKU() bool {
	if o != nil && !IsNil(o.SellerSKU) {
		return true
	}

	return false
}

// SetSellerSKU gets a reference to the given string and assigns it to the SellerSKU field.
func (o *Price) SetSellerSKU(v string) {
	o.SellerSKU = &v
}

// GetASIN returns the ASIN field value if set, zero value otherwise.
func (o *Price) GetASIN() string {
	if o == nil || IsNil(o.ASIN) {
		var ret string
		return ret
	}
	return *o.ASIN
}

// GetASINOk returns a tuple with the ASIN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetASINOk() (*string, bool) {
	if o == nil || IsNil(o.ASIN) {
		return nil, false
	}
	return o.ASIN, true
}

// HasASIN returns a boolean if a field has been set.
func (o *Price) HasASIN() bool {
	if o != nil && !IsNil(o.ASIN) {
		return true
	}

	return false
}

// SetASIN gets a reference to the given string and assigns it to the ASIN field.
func (o *Price) SetASIN(v string) {
	o.ASIN = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *Price) GetProduct() Product {
	if o == nil || IsNil(o.Product) {
		var ret Product
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetProductOk() (*Product, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *Price) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given Product and assigns it to the Product field.
func (o *Price) SetProduct(v Product) {
	o.Product = &v
}

func (o Price) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if !IsNil(o.SellerSKU) {
		toSerialize["SellerSKU"] = o.SellerSKU
	}
	if !IsNil(o.ASIN) {
		toSerialize["ASIN"] = o.ASIN
	}
	if !IsNil(o.Product) {
		toSerialize["Product"] = o.Product
	}
	return toSerialize, nil
}

type NullablePrice struct {
	value *Price
	isSet bool
}

func (v NullablePrice) Get() *Price {
	return v.value
}

func (v *NullablePrice) Set(val *Price) {
	v.value = val
	v.isSet = true
}

func (v NullablePrice) IsSet() bool {
	return v.isSet
}

func (v *NullablePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrice(val *Price) *NullablePrice {
	return &NullablePrice{value: val, isSet: true}
}

func (v NullablePrice) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
