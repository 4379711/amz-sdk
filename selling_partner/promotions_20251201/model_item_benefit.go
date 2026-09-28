package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemBenefit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemBenefit{}

// ItemBenefit Item-level benefit configuration with pricing options. Applicable to `DEAL` and `PRICE_DISCOUNT` promotion types, where each item can have its own distinct benefit. For `COUPON` and `BASKET_BUILDING` promotion types, benefits are configured at the promotion level.
type ItemBenefit struct {
	// The benefit type for item-level pricing.
	Type     string    `json:"type"`
	Price    *Currency `json:"price,omitempty"`
	Discount *Discount `json:"discount,omitempty"`
	// The maximum number of uses per customer for this item benefit.
	PerCustomerUses *int32 `json:"perCustomerUses,omitempty"`
}

// NewItemBenefit instantiates a new ItemBenefit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemBenefit(type_ string) *ItemBenefit {
	this := ItemBenefit{}
	this.Type = type_
	return &this
}

// NewItemBenefitWithDefaults instantiates a new ItemBenefit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemBenefitWithDefaults() *ItemBenefit {
	this := ItemBenefit{}
	return &this
}

// GetType returns the Type field value
func (o *ItemBenefit) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ItemBenefit) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ItemBenefit) SetType(v string) {
	o.Type = v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ItemBenefit) GetPrice() Currency {
	if o == nil || IsNil(o.Price) {
		var ret Currency
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemBenefit) GetPriceOk() (*Currency, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ItemBenefit) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given Currency and assigns it to the Price field.
func (o *ItemBenefit) SetPrice(v Currency) {
	o.Price = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *ItemBenefit) GetDiscount() Discount {
	if o == nil || IsNil(o.Discount) {
		var ret Discount
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemBenefit) GetDiscountOk() (*Discount, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *ItemBenefit) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given Discount and assigns it to the Discount field.
func (o *ItemBenefit) SetDiscount(v Discount) {
	o.Discount = &v
}

// GetPerCustomerUses returns the PerCustomerUses field value if set, zero value otherwise.
func (o *ItemBenefit) GetPerCustomerUses() int32 {
	if o == nil || IsNil(o.PerCustomerUses) {
		var ret int32
		return ret
	}
	return *o.PerCustomerUses
}

// GetPerCustomerUsesOk returns a tuple with the PerCustomerUses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemBenefit) GetPerCustomerUsesOk() (*int32, bool) {
	if o == nil || IsNil(o.PerCustomerUses) {
		return nil, false
	}
	return o.PerCustomerUses, true
}

// HasPerCustomerUses returns a boolean if a field has been set.
func (o *ItemBenefit) HasPerCustomerUses() bool {
	if o != nil && !IsNil(o.PerCustomerUses) {
		return true
	}

	return false
}

// SetPerCustomerUses gets a reference to the given int32 and assigns it to the PerCustomerUses field.
func (o *ItemBenefit) SetPerCustomerUses(v int32) {
	o.PerCustomerUses = &v
}

func (o ItemBenefit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.PerCustomerUses) {
		toSerialize["perCustomerUses"] = o.PerCustomerUses
	}
	return toSerialize, nil
}

type NullableItemBenefit struct {
	value *ItemBenefit
	isSet bool
}

func (v NullableItemBenefit) Get() *ItemBenefit {
	return v.value
}

func (v *NullableItemBenefit) Set(val *ItemBenefit) {
	v.value = val
	v.isSet = true
}

func (v NullableItemBenefit) IsSet() bool {
	return v.isSet
}

func (v *NullableItemBenefit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemBenefit(val *ItemBenefit) *NullableItemBenefit {
	return &NullableItemBenefit{value: val, isSet: true}
}

func (v NullableItemBenefit) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemBenefit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
