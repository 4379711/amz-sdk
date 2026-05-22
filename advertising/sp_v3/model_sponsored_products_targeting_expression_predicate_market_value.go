package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsTargetingExpressionPredicateMarketValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsTargetingExpressionPredicateMarketValue{}

// SponsoredProductsTargetingExpressionPredicateMarketValue struct for SponsoredProductsTargetingExpressionPredicateMarketValue
type SponsoredProductsTargetingExpressionPredicateMarketValue struct {
	Marketplace *SponsoredProductsMarketplace `json:"marketplace,omitempty"`
	// The expression value
	Value *string `json:"value,omitempty"`
}

// NewSponsoredProductsTargetingExpressionPredicateMarketValue instantiates a new SponsoredProductsTargetingExpressionPredicateMarketValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsTargetingExpressionPredicateMarketValue() *SponsoredProductsTargetingExpressionPredicateMarketValue {
	this := SponsoredProductsTargetingExpressionPredicateMarketValue{}
	return &this
}

// NewSponsoredProductsTargetingExpressionPredicateMarketValueWithDefaults instantiates a new SponsoredProductsTargetingExpressionPredicateMarketValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsTargetingExpressionPredicateMarketValueWithDefaults() *SponsoredProductsTargetingExpressionPredicateMarketValue {
	this := SponsoredProductsTargetingExpressionPredicateMarketValue{}
	return &this
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsTargetingExpressionPredicateMarketValue) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingExpressionPredicateMarketValue) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsTargetingExpressionPredicateMarketValue) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsTargetingExpressionPredicateMarketValue) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SponsoredProductsTargetingExpressionPredicateMarketValue) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingExpressionPredicateMarketValue) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SponsoredProductsTargetingExpressionPredicateMarketValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SponsoredProductsTargetingExpressionPredicateMarketValue) SetValue(v string) {
	o.Value = &v
}

func (o SponsoredProductsTargetingExpressionPredicateMarketValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSponsoredProductsTargetingExpressionPredicateMarketValue struct {
	value *SponsoredProductsTargetingExpressionPredicateMarketValue
	isSet bool
}

func (v NullableSponsoredProductsTargetingExpressionPredicateMarketValue) Get() *SponsoredProductsTargetingExpressionPredicateMarketValue {
	return v.value
}

func (v *NullableSponsoredProductsTargetingExpressionPredicateMarketValue) Set(val *SponsoredProductsTargetingExpressionPredicateMarketValue) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetingExpressionPredicateMarketValue) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetingExpressionPredicateMarketValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetingExpressionPredicateMarketValue(val *SponsoredProductsTargetingExpressionPredicateMarketValue) *NullableSponsoredProductsTargetingExpressionPredicateMarketValue {
	return &NullableSponsoredProductsTargetingExpressionPredicateMarketValue{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetingExpressionPredicateMarketValue) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetingExpressionPredicateMarketValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
