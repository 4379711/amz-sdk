package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the BidAdjustmentByPlacement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BidAdjustmentByPlacement{}

// BidAdjustmentByPlacement struct for BidAdjustmentByPlacement
type BidAdjustmentByPlacement struct {
	Percentage *float32   `json:"percentage,omitempty"`
	Placement  *Placement `json:"placement,omitempty"`
}

// NewBidAdjustmentByPlacement instantiates a new BidAdjustmentByPlacement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBidAdjustmentByPlacement() *BidAdjustmentByPlacement {
	this := BidAdjustmentByPlacement{}
	return &this
}

// NewBidAdjustmentByPlacementWithDefaults instantiates a new BidAdjustmentByPlacement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBidAdjustmentByPlacementWithDefaults() *BidAdjustmentByPlacement {
	this := BidAdjustmentByPlacement{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *BidAdjustmentByPlacement) GetPercentage() float32 {
	if o == nil || IsNil(o.Percentage) {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidAdjustmentByPlacement) GetPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *BidAdjustmentByPlacement) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *BidAdjustmentByPlacement) SetPercentage(v float32) {
	o.Percentage = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise.
func (o *BidAdjustmentByPlacement) GetPlacement() Placement {
	if o == nil || IsNil(o.Placement) {
		var ret Placement
		return ret
	}
	return *o.Placement
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidAdjustmentByPlacement) GetPlacementOk() (*Placement, bool) {
	if o == nil || IsNil(o.Placement) {
		return nil, false
	}
	return o.Placement, true
}

// HasPlacement returns a boolean if a field has been set.
func (o *BidAdjustmentByPlacement) HasPlacement() bool {
	if o != nil && !IsNil(o.Placement) {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given Placement and assigns it to the Placement field.
func (o *BidAdjustmentByPlacement) SetPlacement(v Placement) {
	o.Placement = &v
}

func (o BidAdjustmentByPlacement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	if !IsNil(o.Placement) {
		toSerialize["placement"] = o.Placement
	}
	return toSerialize, nil
}

type NullableBidAdjustmentByPlacement struct {
	value *BidAdjustmentByPlacement
	isSet bool
}

func (v NullableBidAdjustmentByPlacement) Get() *BidAdjustmentByPlacement {
	return v.value
}

func (v *NullableBidAdjustmentByPlacement) Set(val *BidAdjustmentByPlacement) {
	v.value = val
	v.isSet = true
}

func (v NullableBidAdjustmentByPlacement) IsSet() bool {
	return v.isSet
}

func (v *NullableBidAdjustmentByPlacement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBidAdjustmentByPlacement(val *BidAdjustmentByPlacement) *NullableBidAdjustmentByPlacement {
	return &NullableBidAdjustmentByPlacement{value: val, isSet: true}
}

func (v NullableBidAdjustmentByPlacement) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBidAdjustmentByPlacement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
