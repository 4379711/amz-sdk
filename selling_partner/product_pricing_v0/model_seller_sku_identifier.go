package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the SellerSKUIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SellerSKUIdentifier{}

// SellerSKUIdentifier Schema to identify an item by MarketPlaceId, SellerId, and SellerSKU.
type SellerSKUIdentifier struct {
	// A marketplace identifier.
	MarketplaceId string `json:"MarketplaceId"`
	// The seller identifier submitted for the operation.
	SellerId string `json:"SellerId"`
	// The seller stock keeping unit (SKU) of the item.
	SellerSKU string `json:"SellerSKU"`
}

type _SellerSKUIdentifier SellerSKUIdentifier

// NewSellerSKUIdentifier instantiates a new SellerSKUIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSellerSKUIdentifier(marketplaceId string, sellerId string, sellerSKU string) *SellerSKUIdentifier {
	this := SellerSKUIdentifier{}
	this.MarketplaceId = marketplaceId
	this.SellerId = sellerId
	this.SellerSKU = sellerSKU
	return &this
}

// NewSellerSKUIdentifierWithDefaults instantiates a new SellerSKUIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSellerSKUIdentifierWithDefaults() *SellerSKUIdentifier {
	this := SellerSKUIdentifier{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *SellerSKUIdentifier) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *SellerSKUIdentifier) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *SellerSKUIdentifier) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetSellerId returns the SellerId field value
func (o *SellerSKUIdentifier) GetSellerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerId
}

// GetSellerIdOk returns a tuple with the SellerId field value
// and a boolean to check if the value has been set.
func (o *SellerSKUIdentifier) GetSellerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerId, true
}

// SetSellerId sets field value
func (o *SellerSKUIdentifier) SetSellerId(v string) {
	o.SellerId = v
}

// GetSellerSKU returns the SellerSKU field value
func (o *SellerSKUIdentifier) GetSellerSKU() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerSKU
}

// GetSellerSKUOk returns a tuple with the SellerSKU field value
// and a boolean to check if the value has been set.
func (o *SellerSKUIdentifier) GetSellerSKUOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerSKU, true
}

// SetSellerSKU sets field value
func (o *SellerSKUIdentifier) SetSellerSKU(v string) {
	o.SellerSKU = v
}

func (o SellerSKUIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["MarketplaceId"] = o.MarketplaceId
	toSerialize["SellerId"] = o.SellerId
	toSerialize["SellerSKU"] = o.SellerSKU
	return toSerialize, nil
}

type NullableSellerSKUIdentifier struct {
	value *SellerSKUIdentifier
	isSet bool
}

func (v NullableSellerSKUIdentifier) Get() *SellerSKUIdentifier {
	return v.value
}

func (v *NullableSellerSKUIdentifier) Set(val *SellerSKUIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableSellerSKUIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableSellerSKUIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSellerSKUIdentifier(val *SellerSKUIdentifier) *NullableSellerSKUIdentifier {
	return &NullableSellerSKUIdentifier{value: val, isSet: true}
}

func (v NullableSellerSKUIdentifier) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSellerSKUIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
