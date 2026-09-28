package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the BenefitTier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BenefitTier{}

// BenefitTier Defines a progressive (multi-tier) benefit tier that applies when customers exceed a higher purchase quantity for basket building promotions. Each tier specifies a purchase condition and corresponding enhanced discount.
type BenefitTier struct {
	PurchaseCondition PurchaseCondition `json:"purchaseCondition"`
	Discount          Discount          `json:"discount"`
}

// NewBenefitTier instantiates a new BenefitTier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBenefitTier(purchaseCondition PurchaseCondition, discount Discount) *BenefitTier {
	this := BenefitTier{}
	this.PurchaseCondition = purchaseCondition
	this.Discount = discount
	return &this
}

// NewBenefitTierWithDefaults instantiates a new BenefitTier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBenefitTierWithDefaults() *BenefitTier {
	this := BenefitTier{}
	return &this
}

// GetPurchaseCondition returns the PurchaseCondition field value
func (o *BenefitTier) GetPurchaseCondition() PurchaseCondition {
	if o == nil {
		var ret PurchaseCondition
		return ret
	}

	return o.PurchaseCondition
}

// GetPurchaseConditionOk returns a tuple with the PurchaseCondition field value
// and a boolean to check if the value has been set.
func (o *BenefitTier) GetPurchaseConditionOk() (*PurchaseCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseCondition, true
}

// SetPurchaseCondition sets field value
func (o *BenefitTier) SetPurchaseCondition(v PurchaseCondition) {
	o.PurchaseCondition = v
}

// GetDiscount returns the Discount field value
func (o *BenefitTier) GetDiscount() Discount {
	if o == nil {
		var ret Discount
		return ret
	}

	return o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value
// and a boolean to check if the value has been set.
func (o *BenefitTier) GetDiscountOk() (*Discount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Discount, true
}

// SetDiscount sets field value
func (o *BenefitTier) SetDiscount(v Discount) {
	o.Discount = v
}

func (o BenefitTier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["purchaseCondition"] = o.PurchaseCondition
	toSerialize["discount"] = o.Discount
	return toSerialize, nil
}

type NullableBenefitTier struct {
	value *BenefitTier
	isSet bool
}

func (v NullableBenefitTier) Get() *BenefitTier {
	return v.value
}

func (v *NullableBenefitTier) Set(val *BenefitTier) {
	v.value = val
	v.isSet = true
}

func (v NullableBenefitTier) IsSet() bool {
	return v.isSet
}

func (v *NullableBenefitTier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBenefitTier(val *BenefitTier) *NullableBenefitTier {
	return &NullableBenefitTier{value: val, isSet: true}
}

func (v NullableBenefitTier) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBenefitTier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
