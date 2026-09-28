package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the PromotionBenefit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotionBenefit{}

// PromotionBenefit Promotion-level benefit configuration. Applicable to `COUPON` and `BASKET_BUILDING` promotion types, where the benefit applies uniformly across all items in the promotion. For `DEAL` and `PRICE_DISCOUNT` promotion types, benefits are configured at the item level within each item in the selection.
type PromotionBenefit struct {
	Discount *Discount `json:"discount,omitempty"`
	// The quantity of items from the benefit selection that receive the discount after the customer satisfies purchase conditions. This property is specific to `BASKET_BUILDING` promotions.
	BenefitQuantity *int32 `json:"benefitQuantity,omitempty"`
	// The maximum number of uses per customer for this promotion.
	PerCustomerUses *int32 `json:"perCustomerUses,omitempty"`
	// Whether this benefit can be stacked with other promotions.
	Stacking *string `json:"stacking,omitempty"`
	// Progressive discount tiers offering increased benefits as customers purchase more. For example: buy 2 get 10% off, buy 3 get 15% off. For multi-tier `BASKET_BUILDING` promotions, each tier specifies additional purchase conditions and corresponding discounts beyond the first tier (defined in `benefit.discount` and `purchaseRequirements.condition`).
	AdditionalTiers []BenefitTier `json:"additionalTiers,omitempty"`
}

// NewPromotionBenefit instantiates a new PromotionBenefit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotionBenefit() *PromotionBenefit {
	this := PromotionBenefit{}
	return &this
}

// NewPromotionBenefitWithDefaults instantiates a new PromotionBenefit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionBenefitWithDefaults() *PromotionBenefit {
	this := PromotionBenefit{}
	return &this
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *PromotionBenefit) GetDiscount() Discount {
	if o == nil || IsNil(o.Discount) {
		var ret Discount
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionBenefit) GetDiscountOk() (*Discount, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *PromotionBenefit) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given Discount and assigns it to the Discount field.
func (o *PromotionBenefit) SetDiscount(v Discount) {
	o.Discount = &v
}

// GetBenefitQuantity returns the BenefitQuantity field value if set, zero value otherwise.
func (o *PromotionBenefit) GetBenefitQuantity() int32 {
	if o == nil || IsNil(o.BenefitQuantity) {
		var ret int32
		return ret
	}
	return *o.BenefitQuantity
}

// GetBenefitQuantityOk returns a tuple with the BenefitQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionBenefit) GetBenefitQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.BenefitQuantity) {
		return nil, false
	}
	return o.BenefitQuantity, true
}

// HasBenefitQuantity returns a boolean if a field has been set.
func (o *PromotionBenefit) HasBenefitQuantity() bool {
	if o != nil && !IsNil(o.BenefitQuantity) {
		return true
	}

	return false
}

// SetBenefitQuantity gets a reference to the given int32 and assigns it to the BenefitQuantity field.
func (o *PromotionBenefit) SetBenefitQuantity(v int32) {
	o.BenefitQuantity = &v
}

// GetPerCustomerUses returns the PerCustomerUses field value if set, zero value otherwise.
func (o *PromotionBenefit) GetPerCustomerUses() int32 {
	if o == nil || IsNil(o.PerCustomerUses) {
		var ret int32
		return ret
	}
	return *o.PerCustomerUses
}

// GetPerCustomerUsesOk returns a tuple with the PerCustomerUses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionBenefit) GetPerCustomerUsesOk() (*int32, bool) {
	if o == nil || IsNil(o.PerCustomerUses) {
		return nil, false
	}
	return o.PerCustomerUses, true
}

// HasPerCustomerUses returns a boolean if a field has been set.
func (o *PromotionBenefit) HasPerCustomerUses() bool {
	if o != nil && !IsNil(o.PerCustomerUses) {
		return true
	}

	return false
}

// SetPerCustomerUses gets a reference to the given int32 and assigns it to the PerCustomerUses field.
func (o *PromotionBenefit) SetPerCustomerUses(v int32) {
	o.PerCustomerUses = &v
}

// GetStacking returns the Stacking field value if set, zero value otherwise.
func (o *PromotionBenefit) GetStacking() string {
	if o == nil || IsNil(o.Stacking) {
		var ret string
		return ret
	}
	return *o.Stacking
}

// GetStackingOk returns a tuple with the Stacking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionBenefit) GetStackingOk() (*string, bool) {
	if o == nil || IsNil(o.Stacking) {
		return nil, false
	}
	return o.Stacking, true
}

// HasStacking returns a boolean if a field has been set.
func (o *PromotionBenefit) HasStacking() bool {
	if o != nil && !IsNil(o.Stacking) {
		return true
	}

	return false
}

// SetStacking gets a reference to the given string and assigns it to the Stacking field.
func (o *PromotionBenefit) SetStacking(v string) {
	o.Stacking = &v
}

// GetAdditionalTiers returns the AdditionalTiers field value if set, zero value otherwise.
func (o *PromotionBenefit) GetAdditionalTiers() []BenefitTier {
	if o == nil || IsNil(o.AdditionalTiers) {
		var ret []BenefitTier
		return ret
	}
	return o.AdditionalTiers
}

// GetAdditionalTiersOk returns a tuple with the AdditionalTiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionBenefit) GetAdditionalTiersOk() ([]BenefitTier, bool) {
	if o == nil || IsNil(o.AdditionalTiers) {
		return nil, false
	}
	return o.AdditionalTiers, true
}

// HasAdditionalTiers returns a boolean if a field has been set.
func (o *PromotionBenefit) HasAdditionalTiers() bool {
	if o != nil && !IsNil(o.AdditionalTiers) {
		return true
	}

	return false
}

// SetAdditionalTiers gets a reference to the given []BenefitTier and assigns it to the AdditionalTiers field.
func (o *PromotionBenefit) SetAdditionalTiers(v []BenefitTier) {
	o.AdditionalTiers = v
}

func (o PromotionBenefit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.BenefitQuantity) {
		toSerialize["benefitQuantity"] = o.BenefitQuantity
	}
	if !IsNil(o.PerCustomerUses) {
		toSerialize["perCustomerUses"] = o.PerCustomerUses
	}
	if !IsNil(o.Stacking) {
		toSerialize["stacking"] = o.Stacking
	}
	if !IsNil(o.AdditionalTiers) {
		toSerialize["additionalTiers"] = o.AdditionalTiers
	}
	return toSerialize, nil
}

type NullablePromotionBenefit struct {
	value *PromotionBenefit
	isSet bool
}

func (v NullablePromotionBenefit) Get() *PromotionBenefit {
	return v.value
}

func (v *NullablePromotionBenefit) Set(val *PromotionBenefit) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionBenefit) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionBenefit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionBenefit(val *PromotionBenefit) *NullablePromotionBenefit {
	return &NullablePromotionBenefit{value: val, isSet: true}
}

func (v NullablePromotionBenefit) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePromotionBenefit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
