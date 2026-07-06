package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemPromotion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemPromotion{}

// ItemPromotion Details about any discounts, coupons, or promotional offers applied to this item.
type ItemPromotion struct {
	// A list of promotions applied to the order item.
	Breakdowns []ItemPromotionBreakdown `json:"breakdowns,omitempty"`
}

// NewItemPromotion instantiates a new ItemPromotion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemPromotion() *ItemPromotion {
	this := ItemPromotion{}
	return &this
}

// NewItemPromotionWithDefaults instantiates a new ItemPromotion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemPromotionWithDefaults() *ItemPromotion {
	this := ItemPromotion{}
	return &this
}

// GetBreakdowns returns the Breakdowns field value if set, zero value otherwise.
func (o *ItemPromotion) GetBreakdowns() []ItemPromotionBreakdown {
	if o == nil || IsNil(o.Breakdowns) {
		var ret []ItemPromotionBreakdown
		return ret
	}
	return o.Breakdowns
}

// GetBreakdownsOk returns a tuple with the Breakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemPromotion) GetBreakdownsOk() ([]ItemPromotionBreakdown, bool) {
	if o == nil || IsNil(o.Breakdowns) {
		return nil, false
	}
	return o.Breakdowns, true
}

// HasBreakdowns returns a boolean if a field has been set.
func (o *ItemPromotion) HasBreakdowns() bool {
	if o != nil && !IsNil(o.Breakdowns) {
		return true
	}

	return false
}

// SetBreakdowns gets a reference to the given []ItemPromotionBreakdown and assigns it to the Breakdowns field.
func (o *ItemPromotion) SetBreakdowns(v []ItemPromotionBreakdown) {
	o.Breakdowns = v
}

func (o ItemPromotion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Breakdowns) {
		toSerialize["breakdowns"] = o.Breakdowns
	}
	return toSerialize, nil
}

type NullableItemPromotion struct {
	value *ItemPromotion
	isSet bool
}

func (v NullableItemPromotion) Get() *ItemPromotion {
	return v.value
}

func (v *NullableItemPromotion) Set(val *ItemPromotion) {
	v.value = val
	v.isSet = true
}

func (v NullableItemPromotion) IsSet() bool {
	return v.isSet
}

func (v *NullableItemPromotion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemPromotion(val *ItemPromotion) *NullableItemPromotion {
	return &NullableItemPromotion{value: val, isSet: true}
}

func (v NullableItemPromotion) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemPromotion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
