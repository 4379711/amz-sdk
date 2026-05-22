package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the Bidding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Bidding{}

// Bidding struct for Bidding
type Bidding struct {
	// Whether to use automatic placement level bid optimization. If set to true, Amazon will automatically set the right placement adjustment and the bidAdjustmentsByPlacement field is ignored. If set to false, the bidAdjustmentsByPlacement field will be used to adjust bid on different placements. If this field is changed from false to true, the bidAdjustmentsByPlacement field will be reset to null.
	BidOptimization *bool `json:"bidOptimization,omitempty"`
	// DEPRECATED [PLANNED SHUTOFF DATE 3/31/2024] **Note: This feature has been deprecated and planned to shutoff on 03/31/2024. After the shut off date, we will ignore this field in the request and treat it as null. You will not get an error if you supply this field in the request. Based on customer feedback, we are rethinking this feature in context to Goal based campaigns to help advertiser reach NTB customers at scale with transparent reporting. Meanwhile, if you have any feedback or suggestion related to this feature then please reach out to our customer support teams.  Shopper segment level bid adjustment. When both bidAdjustmentsByPlacement and bidAdjustmentsByShopperSegment are specified, the adjustment will be multiplicative.
	// Deprecated
	BidAdjustmentsByShopperSegment []BidAdjustmentByShopperSegment `json:"bidAdjustmentsByShopperSegment,omitempty"`
	// Shopper cohort based bid adjustments.
	ShopperCohortBidAdjustments []ShopperCohortBidAdjustment `json:"shopperCohortBidAdjustments,omitempty"`
	// Placement level bid adjustment. Note that this field can only be set when 'bidOptimization' is set to false.
	BidAdjustmentsByPlacement []BidAdjustmentByPlacement `json:"bidAdjustmentsByPlacement,omitempty"`
	// Deprecated
	BidOptimizationStrategy *BidOptimizationStrategy `json:"bidOptimizationStrategy,omitempty"`
}

// NewBidding instantiates a new Bidding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBidding() *Bidding {
	this := Bidding{}
	return &this
}

// NewBiddingWithDefaults instantiates a new Bidding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBiddingWithDefaults() *Bidding {
	this := Bidding{}
	return &this
}

// GetBidOptimization returns the BidOptimization field value if set, zero value otherwise.
func (o *Bidding) GetBidOptimization() bool {
	if o == nil || IsNil(o.BidOptimization) {
		var ret bool
		return ret
	}
	return *o.BidOptimization
}

// GetBidOptimizationOk returns a tuple with the BidOptimization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bidding) GetBidOptimizationOk() (*bool, bool) {
	if o == nil || IsNil(o.BidOptimization) {
		return nil, false
	}
	return o.BidOptimization, true
}

// HasBidOptimization returns a boolean if a field has been set.
func (o *Bidding) HasBidOptimization() bool {
	if o != nil && !IsNil(o.BidOptimization) {
		return true
	}

	return false
}

// SetBidOptimization gets a reference to the given bool and assigns it to the BidOptimization field.
func (o *Bidding) SetBidOptimization(v bool) {
	o.BidOptimization = &v
}

// GetBidAdjustmentsByShopperSegment returns the BidAdjustmentsByShopperSegment field value if set, zero value otherwise.
// Deprecated
func (o *Bidding) GetBidAdjustmentsByShopperSegment() []BidAdjustmentByShopperSegment {
	if o == nil || IsNil(o.BidAdjustmentsByShopperSegment) {
		var ret []BidAdjustmentByShopperSegment
		return ret
	}
	return o.BidAdjustmentsByShopperSegment
}

// GetBidAdjustmentsByShopperSegmentOk returns a tuple with the BidAdjustmentsByShopperSegment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Bidding) GetBidAdjustmentsByShopperSegmentOk() ([]BidAdjustmentByShopperSegment, bool) {
	if o == nil || IsNil(o.BidAdjustmentsByShopperSegment) {
		return nil, false
	}
	return o.BidAdjustmentsByShopperSegment, true
}

// HasBidAdjustmentsByShopperSegment returns a boolean if a field has been set.
func (o *Bidding) HasBidAdjustmentsByShopperSegment() bool {
	if o != nil && !IsNil(o.BidAdjustmentsByShopperSegment) {
		return true
	}

	return false
}

// SetBidAdjustmentsByShopperSegment gets a reference to the given []BidAdjustmentByShopperSegment and assigns it to the BidAdjustmentsByShopperSegment field.
// Deprecated
func (o *Bidding) SetBidAdjustmentsByShopperSegment(v []BidAdjustmentByShopperSegment) {
	o.BidAdjustmentsByShopperSegment = v
}

// GetShopperCohortBidAdjustments returns the ShopperCohortBidAdjustments field value if set, zero value otherwise.
func (o *Bidding) GetShopperCohortBidAdjustments() []ShopperCohortBidAdjustment {
	if o == nil || IsNil(o.ShopperCohortBidAdjustments) {
		var ret []ShopperCohortBidAdjustment
		return ret
	}
	return o.ShopperCohortBidAdjustments
}

// GetShopperCohortBidAdjustmentsOk returns a tuple with the ShopperCohortBidAdjustments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bidding) GetShopperCohortBidAdjustmentsOk() ([]ShopperCohortBidAdjustment, bool) {
	if o == nil || IsNil(o.ShopperCohortBidAdjustments) {
		return nil, false
	}
	return o.ShopperCohortBidAdjustments, true
}

// HasShopperCohortBidAdjustments returns a boolean if a field has been set.
func (o *Bidding) HasShopperCohortBidAdjustments() bool {
	if o != nil && !IsNil(o.ShopperCohortBidAdjustments) {
		return true
	}

	return false
}

// SetShopperCohortBidAdjustments gets a reference to the given []ShopperCohortBidAdjustment and assigns it to the ShopperCohortBidAdjustments field.
func (o *Bidding) SetShopperCohortBidAdjustments(v []ShopperCohortBidAdjustment) {
	o.ShopperCohortBidAdjustments = v
}

// GetBidAdjustmentsByPlacement returns the BidAdjustmentsByPlacement field value if set, zero value otherwise.
func (o *Bidding) GetBidAdjustmentsByPlacement() []BidAdjustmentByPlacement {
	if o == nil || IsNil(o.BidAdjustmentsByPlacement) {
		var ret []BidAdjustmentByPlacement
		return ret
	}
	return o.BidAdjustmentsByPlacement
}

// GetBidAdjustmentsByPlacementOk returns a tuple with the BidAdjustmentsByPlacement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bidding) GetBidAdjustmentsByPlacementOk() ([]BidAdjustmentByPlacement, bool) {
	if o == nil || IsNil(o.BidAdjustmentsByPlacement) {
		return nil, false
	}
	return o.BidAdjustmentsByPlacement, true
}

// HasBidAdjustmentsByPlacement returns a boolean if a field has been set.
func (o *Bidding) HasBidAdjustmentsByPlacement() bool {
	if o != nil && !IsNil(o.BidAdjustmentsByPlacement) {
		return true
	}

	return false
}

// SetBidAdjustmentsByPlacement gets a reference to the given []BidAdjustmentByPlacement and assigns it to the BidAdjustmentsByPlacement field.
func (o *Bidding) SetBidAdjustmentsByPlacement(v []BidAdjustmentByPlacement) {
	o.BidAdjustmentsByPlacement = v
}

// GetBidOptimizationStrategy returns the BidOptimizationStrategy field value if set, zero value otherwise.
// Deprecated
func (o *Bidding) GetBidOptimizationStrategy() BidOptimizationStrategy {
	if o == nil || IsNil(o.BidOptimizationStrategy) {
		var ret BidOptimizationStrategy
		return ret
	}
	return *o.BidOptimizationStrategy
}

// GetBidOptimizationStrategyOk returns a tuple with the BidOptimizationStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Bidding) GetBidOptimizationStrategyOk() (*BidOptimizationStrategy, bool) {
	if o == nil || IsNil(o.BidOptimizationStrategy) {
		return nil, false
	}
	return o.BidOptimizationStrategy, true
}

// HasBidOptimizationStrategy returns a boolean if a field has been set.
func (o *Bidding) HasBidOptimizationStrategy() bool {
	if o != nil && !IsNil(o.BidOptimizationStrategy) {
		return true
	}

	return false
}

// SetBidOptimizationStrategy gets a reference to the given BidOptimizationStrategy and assigns it to the BidOptimizationStrategy field.
// Deprecated
func (o *Bidding) SetBidOptimizationStrategy(v BidOptimizationStrategy) {
	o.BidOptimizationStrategy = &v
}

func (o Bidding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BidOptimization) {
		toSerialize["bidOptimization"] = o.BidOptimization
	}
	if !IsNil(o.BidAdjustmentsByShopperSegment) {
		toSerialize["bidAdjustmentsByShopperSegment"] = o.BidAdjustmentsByShopperSegment
	}
	if !IsNil(o.ShopperCohortBidAdjustments) {
		toSerialize["shopperCohortBidAdjustments"] = o.ShopperCohortBidAdjustments
	}
	if !IsNil(o.BidAdjustmentsByPlacement) {
		toSerialize["bidAdjustmentsByPlacement"] = o.BidAdjustmentsByPlacement
	}
	if !IsNil(o.BidOptimizationStrategy) {
		toSerialize["bidOptimizationStrategy"] = o.BidOptimizationStrategy
	}
	return toSerialize, nil
}

type NullableBidding struct {
	value *Bidding
	isSet bool
}

func (v NullableBidding) Get() *Bidding {
	return v.value
}

func (v *NullableBidding) Set(val *Bidding) {
	v.value = val
	v.isSet = true
}

func (v NullableBidding) IsSet() bool {
	return v.isSet
}

func (v *NullableBidding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBidding(val *Bidding) *NullableBidding {
	return &NullableBidding{value: val, isSet: true}
}

func (v NullableBidding) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBidding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
