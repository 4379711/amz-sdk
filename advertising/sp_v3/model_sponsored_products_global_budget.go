package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalBudget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalBudget{}

// SponsoredProductsGlobalBudget struct for SponsoredProductsGlobalBudget
type SponsoredProductsGlobalBudget struct {
	BudgetType SponsoredProductsBudgetType `json:"budgetType"`
	// marketplace budget settings.
	MarketplaceSettings []SponsoredProductsMarketplaceBudget `json:"marketplaceSettings,omitempty"`
	Currency            SponsoredProductsCurrency            `json:"currency"`
}

type _SponsoredProductsGlobalBudget SponsoredProductsGlobalBudget

// NewSponsoredProductsGlobalBudget instantiates a new SponsoredProductsGlobalBudget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalBudget(budgetType SponsoredProductsBudgetType, currency SponsoredProductsCurrency) *SponsoredProductsGlobalBudget {
	this := SponsoredProductsGlobalBudget{}
	this.BudgetType = budgetType
	this.Currency = currency
	return &this
}

// NewSponsoredProductsGlobalBudgetWithDefaults instantiates a new SponsoredProductsGlobalBudget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalBudgetWithDefaults() *SponsoredProductsGlobalBudget {
	this := SponsoredProductsGlobalBudget{}
	return &this
}

// GetBudgetType returns the BudgetType field value
func (o *SponsoredProductsGlobalBudget) GetBudgetType() SponsoredProductsBudgetType {
	if o == nil {
		var ret SponsoredProductsBudgetType
		return ret
	}

	return o.BudgetType
}

// GetBudgetTypeOk returns a tuple with the BudgetType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalBudget) GetBudgetTypeOk() (*SponsoredProductsBudgetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetType, true
}

// SetBudgetType sets field value
func (o *SponsoredProductsGlobalBudget) SetBudgetType(v SponsoredProductsBudgetType) {
	o.BudgetType = v
}

// GetMarketplaceSettings returns the MarketplaceSettings field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalBudget) GetMarketplaceSettings() []SponsoredProductsMarketplaceBudget {
	if o == nil || IsNil(o.MarketplaceSettings) {
		var ret []SponsoredProductsMarketplaceBudget
		return ret
	}
	return o.MarketplaceSettings
}

// GetMarketplaceSettingsOk returns a tuple with the MarketplaceSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalBudget) GetMarketplaceSettingsOk() ([]SponsoredProductsMarketplaceBudget, bool) {
	if o == nil || IsNil(o.MarketplaceSettings) {
		return nil, false
	}
	return o.MarketplaceSettings, true
}

// HasMarketplaceSettings returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalBudget) HasMarketplaceSettings() bool {
	if o != nil && !IsNil(o.MarketplaceSettings) {
		return true
	}

	return false
}

// SetMarketplaceSettings gets a reference to the given []SponsoredProductsMarketplaceBudget and assigns it to the MarketplaceSettings field.
func (o *SponsoredProductsGlobalBudget) SetMarketplaceSettings(v []SponsoredProductsMarketplaceBudget) {
	o.MarketplaceSettings = v
}

// GetCurrency returns the Currency field value
func (o *SponsoredProductsGlobalBudget) GetCurrency() SponsoredProductsCurrency {
	if o == nil {
		var ret SponsoredProductsCurrency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalBudget) GetCurrencyOk() (*SponsoredProductsCurrency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *SponsoredProductsGlobalBudget) SetCurrency(v SponsoredProductsCurrency) {
	o.Currency = v
}

func (o SponsoredProductsGlobalBudget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["budgetType"] = o.BudgetType
	if !IsNil(o.MarketplaceSettings) {
		toSerialize["marketplaceSettings"] = o.MarketplaceSettings
	}
	toSerialize["currency"] = o.Currency
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalBudget struct {
	value *SponsoredProductsGlobalBudget
	isSet bool
}

func (v NullableSponsoredProductsGlobalBudget) Get() *SponsoredProductsGlobalBudget {
	return v.value
}

func (v *NullableSponsoredProductsGlobalBudget) Set(val *SponsoredProductsGlobalBudget) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalBudget) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalBudget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalBudget(val *SponsoredProductsGlobalBudget) *NullableSponsoredProductsGlobalBudget {
	return &NullableSponsoredProductsGlobalBudget{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalBudget) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalBudget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
