package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemPromotionBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemPromotionBreakdown{}

// ItemPromotionBreakdown Detailed information about a specific promotional offer applied to an order item.
type ItemPromotionBreakdown struct {
	// Unique identifier for the promotion applied to this item.
	PromotionId *string `json:"promotionId,omitempty"`
}

// NewItemPromotionBreakdown instantiates a new ItemPromotionBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemPromotionBreakdown() *ItemPromotionBreakdown {
	this := ItemPromotionBreakdown{}
	return &this
}

// NewItemPromotionBreakdownWithDefaults instantiates a new ItemPromotionBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemPromotionBreakdownWithDefaults() *ItemPromotionBreakdown {
	this := ItemPromotionBreakdown{}
	return &this
}

// GetPromotionId returns the PromotionId field value if set, zero value otherwise.
func (o *ItemPromotionBreakdown) GetPromotionId() string {
	if o == nil || IsNil(o.PromotionId) {
		var ret string
		return ret
	}
	return *o.PromotionId
}

// GetPromotionIdOk returns a tuple with the PromotionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemPromotionBreakdown) GetPromotionIdOk() (*string, bool) {
	if o == nil || IsNil(o.PromotionId) {
		return nil, false
	}
	return o.PromotionId, true
}

// HasPromotionId returns a boolean if a field has been set.
func (o *ItemPromotionBreakdown) HasPromotionId() bool {
	if o != nil && !IsNil(o.PromotionId) {
		return true
	}

	return false
}

// SetPromotionId gets a reference to the given string and assigns it to the PromotionId field.
func (o *ItemPromotionBreakdown) SetPromotionId(v string) {
	o.PromotionId = &v
}

func (o ItemPromotionBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PromotionId) {
		toSerialize["promotionId"] = o.PromotionId
	}
	return toSerialize, nil
}

type NullableItemPromotionBreakdown struct {
	value *ItemPromotionBreakdown
	isSet bool
}

func (v NullableItemPromotionBreakdown) Get() *ItemPromotionBreakdown {
	return v.value
}

func (v *NullableItemPromotionBreakdown) Set(val *ItemPromotionBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableItemPromotionBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableItemPromotionBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemPromotionBreakdown(val *ItemPromotionBreakdown) *NullableItemPromotionBreakdown {
	return &NullableItemPromotionBreakdown{value: val, isSet: true}
}

func (v NullableItemPromotionBreakdown) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemPromotionBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
