package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsMarketplaceBid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsMarketplaceBid{}

// SponsoredProductsMarketplaceBid struct for SponsoredProductsMarketplaceBid
type SponsoredProductsMarketplaceBid struct {
	Marketplace *SponsoredProductsMarketplace `json:"marketplace,omitempty"`
	// Monetary value
	Bid *float64 `json:"bid,omitempty"`
}

// NewSponsoredProductsMarketplaceBid instantiates a new SponsoredProductsMarketplaceBid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsMarketplaceBid() *SponsoredProductsMarketplaceBid {
	this := SponsoredProductsMarketplaceBid{}
	return &this
}

// NewSponsoredProductsMarketplaceBidWithDefaults instantiates a new SponsoredProductsMarketplaceBid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsMarketplaceBidWithDefaults() *SponsoredProductsMarketplaceBid {
	this := SponsoredProductsMarketplaceBid{}
	return &this
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsMarketplaceBid) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceBid) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsMarketplaceBid) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsMarketplaceBid) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *SponsoredProductsMarketplaceBid) GetBid() float64 {
	if o == nil || IsNil(o.Bid) {
		var ret float64
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceBid) GetBidOk() (*float64, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsMarketplaceBid) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given float64 and assigns it to the Bid field.
func (o *SponsoredProductsMarketplaceBid) SetBid(v float64) {
	o.Bid = &v
}

func (o SponsoredProductsMarketplaceBid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	return toSerialize, nil
}

type NullableSponsoredProductsMarketplaceBid struct {
	value *SponsoredProductsMarketplaceBid
	isSet bool
}

func (v NullableSponsoredProductsMarketplaceBid) Get() *SponsoredProductsMarketplaceBid {
	return v.value
}

func (v *NullableSponsoredProductsMarketplaceBid) Set(val *SponsoredProductsMarketplaceBid) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsMarketplaceBid) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsMarketplaceBid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsMarketplaceBid(val *SponsoredProductsMarketplaceBid) *NullableSponsoredProductsMarketplaceBid {
	return &NullableSponsoredProductsMarketplaceBid{value: val, isSet: true}
}

func (v NullableSponsoredProductsMarketplaceBid) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsMarketplaceBid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
