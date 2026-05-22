package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsMarketplaceBudget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsMarketplaceBudget{}

// SponsoredProductsMarketplaceBudget struct for SponsoredProductsMarketplaceBudget
type SponsoredProductsMarketplaceBudget struct {
	Marketplace *SponsoredProductsMarketplace `json:"marketplace,omitempty"`
	// Monetary value
	Budget float64 `json:"budget"`
}

type _SponsoredProductsMarketplaceBudget SponsoredProductsMarketplaceBudget

// NewSponsoredProductsMarketplaceBudget instantiates a new SponsoredProductsMarketplaceBudget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsMarketplaceBudget(budget float64) *SponsoredProductsMarketplaceBudget {
	this := SponsoredProductsMarketplaceBudget{}
	this.Budget = budget
	return &this
}

// NewSponsoredProductsMarketplaceBudgetWithDefaults instantiates a new SponsoredProductsMarketplaceBudget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsMarketplaceBudgetWithDefaults() *SponsoredProductsMarketplaceBudget {
	this := SponsoredProductsMarketplaceBudget{}
	return &this
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsMarketplaceBudget) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceBudget) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsMarketplaceBudget) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsMarketplaceBudget) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

// GetBudget returns the Budget field value
func (o *SponsoredProductsMarketplaceBudget) GetBudget() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceBudget) GetBudgetOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Budget, true
}

// SetBudget sets field value
func (o *SponsoredProductsMarketplaceBudget) SetBudget(v float64) {
	o.Budget = v
}

func (o SponsoredProductsMarketplaceBudget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	toSerialize["budget"] = o.Budget
	return toSerialize, nil
}

type NullableSponsoredProductsMarketplaceBudget struct {
	value *SponsoredProductsMarketplaceBudget
	isSet bool
}

func (v NullableSponsoredProductsMarketplaceBudget) Get() *SponsoredProductsMarketplaceBudget {
	return v.value
}

func (v *NullableSponsoredProductsMarketplaceBudget) Set(val *SponsoredProductsMarketplaceBudget) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsMarketplaceBudget) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsMarketplaceBudget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsMarketplaceBudget(val *SponsoredProductsMarketplaceBudget) *NullableSponsoredProductsMarketplaceBudget {
	return &NullableSponsoredProductsMarketplaceBudget{value: val, isSet: true}
}

func (v NullableSponsoredProductsMarketplaceBudget) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsMarketplaceBudget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
