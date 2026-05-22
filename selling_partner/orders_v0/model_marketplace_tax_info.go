package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the MarketplaceTaxInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketplaceTaxInfo{}

// MarketplaceTaxInfo Tax information about the marketplace.
type MarketplaceTaxInfo struct {
	// A list of tax classifications that apply to the order.
	TaxClassifications []TaxClassification `json:"TaxClassifications,omitempty"`
}

// NewMarketplaceTaxInfo instantiates a new MarketplaceTaxInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceTaxInfo() *MarketplaceTaxInfo {
	this := MarketplaceTaxInfo{}
	return &this
}

// NewMarketplaceTaxInfoWithDefaults instantiates a new MarketplaceTaxInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceTaxInfoWithDefaults() *MarketplaceTaxInfo {
	this := MarketplaceTaxInfo{}
	return &this
}

// GetTaxClassifications returns the TaxClassifications field value if set, zero value otherwise.
func (o *MarketplaceTaxInfo) GetTaxClassifications() []TaxClassification {
	if o == nil || IsNil(o.TaxClassifications) {
		var ret []TaxClassification
		return ret
	}
	return o.TaxClassifications
}

// GetTaxClassificationsOk returns a tuple with the TaxClassifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceTaxInfo) GetTaxClassificationsOk() ([]TaxClassification, bool) {
	if o == nil || IsNil(o.TaxClassifications) {
		return nil, false
	}
	return o.TaxClassifications, true
}

// HasTaxClassifications returns a boolean if a field has been set.
func (o *MarketplaceTaxInfo) HasTaxClassifications() bool {
	if o != nil && !IsNil(o.TaxClassifications) {
		return true
	}

	return false
}

// SetTaxClassifications gets a reference to the given []TaxClassification and assigns it to the TaxClassifications field.
func (o *MarketplaceTaxInfo) SetTaxClassifications(v []TaxClassification) {
	o.TaxClassifications = v
}

func (o MarketplaceTaxInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaxClassifications) {
		toSerialize["TaxClassifications"] = o.TaxClassifications
	}
	return toSerialize, nil
}

type NullableMarketplaceTaxInfo struct {
	value *MarketplaceTaxInfo
	isSet bool
}

func (v NullableMarketplaceTaxInfo) Get() *MarketplaceTaxInfo {
	return v.value
}

func (v *NullableMarketplaceTaxInfo) Set(val *MarketplaceTaxInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceTaxInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceTaxInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceTaxInfo(val *MarketplaceTaxInfo) *NullableMarketplaceTaxInfo {
	return &NullableMarketplaceTaxInfo{value: val, isSet: true}
}

func (v NullableMarketplaceTaxInfo) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMarketplaceTaxInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
