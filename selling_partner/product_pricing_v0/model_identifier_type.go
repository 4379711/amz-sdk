package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the IdentifierType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentifierType{}

// IdentifierType Specifies the identifiers used to uniquely identify an item.
type IdentifierType struct {
	MarketplaceASIN ASINIdentifier       `json:"MarketplaceASIN"`
	SKUIdentifier   *SellerSKUIdentifier `json:"SKUIdentifier,omitempty"`
}

type _IdentifierType IdentifierType

// NewIdentifierType instantiates a new IdentifierType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentifierType(marketplaceASIN ASINIdentifier) *IdentifierType {
	this := IdentifierType{}
	this.MarketplaceASIN = marketplaceASIN
	return &this
}

// NewIdentifierTypeWithDefaults instantiates a new IdentifierType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentifierTypeWithDefaults() *IdentifierType {
	this := IdentifierType{}
	return &this
}

// GetMarketplaceASIN returns the MarketplaceASIN field value
func (o *IdentifierType) GetMarketplaceASIN() ASINIdentifier {
	if o == nil {
		var ret ASINIdentifier
		return ret
	}

	return o.MarketplaceASIN
}

// GetMarketplaceASINOk returns a tuple with the MarketplaceASIN field value
// and a boolean to check if the value has been set.
func (o *IdentifierType) GetMarketplaceASINOk() (*ASINIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceASIN, true
}

// SetMarketplaceASIN sets field value
func (o *IdentifierType) SetMarketplaceASIN(v ASINIdentifier) {
	o.MarketplaceASIN = v
}

// GetSKUIdentifier returns the SKUIdentifier field value if set, zero value otherwise.
func (o *IdentifierType) GetSKUIdentifier() SellerSKUIdentifier {
	if o == nil || IsNil(o.SKUIdentifier) {
		var ret SellerSKUIdentifier
		return ret
	}
	return *o.SKUIdentifier
}

// GetSKUIdentifierOk returns a tuple with the SKUIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentifierType) GetSKUIdentifierOk() (*SellerSKUIdentifier, bool) {
	if o == nil || IsNil(o.SKUIdentifier) {
		return nil, false
	}
	return o.SKUIdentifier, true
}

// HasSKUIdentifier returns a boolean if a field has been set.
func (o *IdentifierType) HasSKUIdentifier() bool {
	if o != nil && !IsNil(o.SKUIdentifier) {
		return true
	}

	return false
}

// SetSKUIdentifier gets a reference to the given SellerSKUIdentifier and assigns it to the SKUIdentifier field.
func (o *IdentifierType) SetSKUIdentifier(v SellerSKUIdentifier) {
	o.SKUIdentifier = &v
}

func (o IdentifierType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["MarketplaceASIN"] = o.MarketplaceASIN
	if !IsNil(o.SKUIdentifier) {
		toSerialize["SKUIdentifier"] = o.SKUIdentifier
	}
	return toSerialize, nil
}

type NullableIdentifierType struct {
	value *IdentifierType
	isSet bool
}

func (v NullableIdentifierType) Get() *IdentifierType {
	return v.value
}

func (v *NullableIdentifierType) Set(val *IdentifierType) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentifierType) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentifierType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentifierType(val *IdentifierType) *NullableIdentifierType {
	return &NullableIdentifierType{value: val, isSet: true}
}

func (v NullableIdentifierType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableIdentifierType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
