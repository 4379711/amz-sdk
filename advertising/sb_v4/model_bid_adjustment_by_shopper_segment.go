package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the BidAdjustmentByShopperSegment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BidAdjustmentByShopperSegment{}

// BidAdjustmentByShopperSegment struct for BidAdjustmentByShopperSegment
type BidAdjustmentByShopperSegment struct {
	Percentage     *float32        `json:"percentage,omitempty"`
	ShopperSegment *ShopperSegment `json:"shopperSegment,omitempty"`
}

// NewBidAdjustmentByShopperSegment instantiates a new BidAdjustmentByShopperSegment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBidAdjustmentByShopperSegment() *BidAdjustmentByShopperSegment {
	this := BidAdjustmentByShopperSegment{}
	return &this
}

// NewBidAdjustmentByShopperSegmentWithDefaults instantiates a new BidAdjustmentByShopperSegment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBidAdjustmentByShopperSegmentWithDefaults() *BidAdjustmentByShopperSegment {
	this := BidAdjustmentByShopperSegment{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *BidAdjustmentByShopperSegment) GetPercentage() float32 {
	if o == nil || IsNil(o.Percentage) {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidAdjustmentByShopperSegment) GetPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *BidAdjustmentByShopperSegment) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *BidAdjustmentByShopperSegment) SetPercentage(v float32) {
	o.Percentage = &v
}

// GetShopperSegment returns the ShopperSegment field value if set, zero value otherwise.
func (o *BidAdjustmentByShopperSegment) GetShopperSegment() ShopperSegment {
	if o == nil || IsNil(o.ShopperSegment) {
		var ret ShopperSegment
		return ret
	}
	return *o.ShopperSegment
}

// GetShopperSegmentOk returns a tuple with the ShopperSegment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidAdjustmentByShopperSegment) GetShopperSegmentOk() (*ShopperSegment, bool) {
	if o == nil || IsNil(o.ShopperSegment) {
		return nil, false
	}
	return o.ShopperSegment, true
}

// HasShopperSegment returns a boolean if a field has been set.
func (o *BidAdjustmentByShopperSegment) HasShopperSegment() bool {
	if o != nil && !IsNil(o.ShopperSegment) {
		return true
	}

	return false
}

// SetShopperSegment gets a reference to the given ShopperSegment and assigns it to the ShopperSegment field.
func (o *BidAdjustmentByShopperSegment) SetShopperSegment(v ShopperSegment) {
	o.ShopperSegment = &v
}

func (o BidAdjustmentByShopperSegment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	if !IsNil(o.ShopperSegment) {
		toSerialize["shopperSegment"] = o.ShopperSegment
	}
	return toSerialize, nil
}

type NullableBidAdjustmentByShopperSegment struct {
	value *BidAdjustmentByShopperSegment
	isSet bool
}

func (v NullableBidAdjustmentByShopperSegment) Get() *BidAdjustmentByShopperSegment {
	return v.value
}

func (v *NullableBidAdjustmentByShopperSegment) Set(val *BidAdjustmentByShopperSegment) {
	v.value = val
	v.isSet = true
}

func (v NullableBidAdjustmentByShopperSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableBidAdjustmentByShopperSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBidAdjustmentByShopperSegment(val *BidAdjustmentByShopperSegment) *NullableBidAdjustmentByShopperSegment {
	return &NullableBidAdjustmentByShopperSegment{value: val, isSet: true}
}

func (v NullableBidAdjustmentByShopperSegment) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBidAdjustmentByShopperSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
