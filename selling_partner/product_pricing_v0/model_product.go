package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the Product type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Product{}

// Product An item.
type Product struct {
	Identifiers IdentifierType `json:"Identifiers"`
	// A list of product attributes if they are applicable to the product that is returned.
	AttributeSets []map[string]interface{} `json:"AttributeSets,omitempty"`
	// A list that contains product variation information, if applicable.
	Relationships      []map[string]interface{} `json:"Relationships,omitempty"`
	CompetitivePricing *CompetitivePricingType  `json:"CompetitivePricing,omitempty"`
	// A list of sales rank information for the item, by category.
	SalesRankings []SalesRankType `json:"SalesRankings,omitempty"`
	// A list of offers.
	Offers []OfferType `json:"Offers,omitempty"`
}

type _Product Product

// NewProduct instantiates a new Product object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct(identifiers IdentifierType) *Product {
	this := Product{}
	this.Identifiers = identifiers
	return &this
}

// NewProductWithDefaults instantiates a new Product object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithDefaults() *Product {
	this := Product{}
	return &this
}

// GetIdentifiers returns the Identifiers field value
func (o *Product) GetIdentifiers() IdentifierType {
	if o == nil {
		var ret IdentifierType
		return ret
	}

	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value
// and a boolean to check if the value has been set.
func (o *Product) GetIdentifiersOk() (*IdentifierType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifiers, true
}

// SetIdentifiers sets field value
func (o *Product) SetIdentifiers(v IdentifierType) {
	o.Identifiers = v
}

// GetAttributeSets returns the AttributeSets field value if set, zero value otherwise.
func (o *Product) GetAttributeSets() []map[string]interface{} {
	if o == nil || IsNil(o.AttributeSets) {
		var ret []map[string]interface{}
		return ret
	}
	return o.AttributeSets
}

// GetAttributeSetsOk returns a tuple with the AttributeSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetAttributeSetsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.AttributeSets) {
		return nil, false
	}
	return o.AttributeSets, true
}

// HasAttributeSets returns a boolean if a field has been set.
func (o *Product) HasAttributeSets() bool {
	if o != nil && !IsNil(o.AttributeSets) {
		return true
	}

	return false
}

// SetAttributeSets gets a reference to the given []map[string]interface{} and assigns it to the AttributeSets field.
func (o *Product) SetAttributeSets(v []map[string]interface{}) {
	o.AttributeSets = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *Product) GetRelationships() []map[string]interface{} {
	if o == nil || IsNil(o.Relationships) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetRelationshipsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *Product) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given []map[string]interface{} and assigns it to the Relationships field.
func (o *Product) SetRelationships(v []map[string]interface{}) {
	o.Relationships = v
}

// GetCompetitivePricing returns the CompetitivePricing field value if set, zero value otherwise.
func (o *Product) GetCompetitivePricing() CompetitivePricingType {
	if o == nil || IsNil(o.CompetitivePricing) {
		var ret CompetitivePricingType
		return ret
	}
	return *o.CompetitivePricing
}

// GetCompetitivePricingOk returns a tuple with the CompetitivePricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetCompetitivePricingOk() (*CompetitivePricingType, bool) {
	if o == nil || IsNil(o.CompetitivePricing) {
		return nil, false
	}
	return o.CompetitivePricing, true
}

// HasCompetitivePricing returns a boolean if a field has been set.
func (o *Product) HasCompetitivePricing() bool {
	if o != nil && !IsNil(o.CompetitivePricing) {
		return true
	}

	return false
}

// SetCompetitivePricing gets a reference to the given CompetitivePricingType and assigns it to the CompetitivePricing field.
func (o *Product) SetCompetitivePricing(v CompetitivePricingType) {
	o.CompetitivePricing = &v
}

// GetSalesRankings returns the SalesRankings field value if set, zero value otherwise.
func (o *Product) GetSalesRankings() []SalesRankType {
	if o == nil || IsNil(o.SalesRankings) {
		var ret []SalesRankType
		return ret
	}
	return o.SalesRankings
}

// GetSalesRankingsOk returns a tuple with the SalesRankings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetSalesRankingsOk() ([]SalesRankType, bool) {
	if o == nil || IsNil(o.SalesRankings) {
		return nil, false
	}
	return o.SalesRankings, true
}

// HasSalesRankings returns a boolean if a field has been set.
func (o *Product) HasSalesRankings() bool {
	if o != nil && !IsNil(o.SalesRankings) {
		return true
	}

	return false
}

// SetSalesRankings gets a reference to the given []SalesRankType and assigns it to the SalesRankings field.
func (o *Product) SetSalesRankings(v []SalesRankType) {
	o.SalesRankings = v
}

// GetOffers returns the Offers field value if set, zero value otherwise.
func (o *Product) GetOffers() []OfferType {
	if o == nil || IsNil(o.Offers) {
		var ret []OfferType
		return ret
	}
	return o.Offers
}

// GetOffersOk returns a tuple with the Offers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetOffersOk() ([]OfferType, bool) {
	if o == nil || IsNil(o.Offers) {
		return nil, false
	}
	return o.Offers, true
}

// HasOffers returns a boolean if a field has been set.
func (o *Product) HasOffers() bool {
	if o != nil && !IsNil(o.Offers) {
		return true
	}

	return false
}

// SetOffers gets a reference to the given []OfferType and assigns it to the Offers field.
func (o *Product) SetOffers(v []OfferType) {
	o.Offers = v
}

func (o Product) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Identifiers"] = o.Identifiers
	if !IsNil(o.AttributeSets) {
		toSerialize["AttributeSets"] = o.AttributeSets
	}
	if !IsNil(o.Relationships) {
		toSerialize["Relationships"] = o.Relationships
	}
	if !IsNil(o.CompetitivePricing) {
		toSerialize["CompetitivePricing"] = o.CompetitivePricing
	}
	if !IsNil(o.SalesRankings) {
		toSerialize["SalesRankings"] = o.SalesRankings
	}
	if !IsNil(o.Offers) {
		toSerialize["Offers"] = o.Offers
	}
	return toSerialize, nil
}

type NullableProduct struct {
	value *Product
	isSet bool
}

func (v NullableProduct) Get() *Product {
	return v.value
}

func (v *NullableProduct) Set(val *Product) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct(val *Product) *NullableProduct {
	return &NullableProduct{value: val, isSet: true}
}

func (v NullableProduct) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
