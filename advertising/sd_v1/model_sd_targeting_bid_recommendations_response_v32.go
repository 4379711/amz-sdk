package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingBidRecommendationsResponseV32 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingBidRecommendationsResponseV32{}

// SDTargetingBidRecommendationsResponseV32 Response to a request for targeting bid recommendations.
type SDTargetingBidRecommendationsResponseV32 struct {
	BidOptimization    SDBidOptimizationV32                                              `json:"bidOptimization"`
	CostType           SDCostTypeV31                                                     `json:"costType"`
	BidRecommendations []SDTargetingBidRecommendationsResponseV31BidRecommendationsInner `json:"bidRecommendations"`
}

type _SDTargetingBidRecommendationsResponseV32 SDTargetingBidRecommendationsResponseV32

// NewSDTargetingBidRecommendationsResponseV32 instantiates a new SDTargetingBidRecommendationsResponseV32 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingBidRecommendationsResponseV32(bidOptimization SDBidOptimizationV32, costType SDCostTypeV31, bidRecommendations []SDTargetingBidRecommendationsResponseV31BidRecommendationsInner) *SDTargetingBidRecommendationsResponseV32 {
	this := SDTargetingBidRecommendationsResponseV32{}
	this.BidOptimization = bidOptimization
	this.CostType = costType
	this.BidRecommendations = bidRecommendations
	return &this
}

// NewSDTargetingBidRecommendationsResponseV32WithDefaults instantiates a new SDTargetingBidRecommendationsResponseV32 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingBidRecommendationsResponseV32WithDefaults() *SDTargetingBidRecommendationsResponseV32 {
	this := SDTargetingBidRecommendationsResponseV32{}
	return &this
}

// GetBidOptimization returns the BidOptimization field value
func (o *SDTargetingBidRecommendationsResponseV32) GetBidOptimization() SDBidOptimizationV32 {
	if o == nil {
		var ret SDBidOptimizationV32
		return ret
	}

	return o.BidOptimization
}

// GetBidOptimizationOk returns a tuple with the BidOptimization field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsResponseV32) GetBidOptimizationOk() (*SDBidOptimizationV32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BidOptimization, true
}

// SetBidOptimization sets field value
func (o *SDTargetingBidRecommendationsResponseV32) SetBidOptimization(v SDBidOptimizationV32) {
	o.BidOptimization = v
}

// GetCostType returns the CostType field value
func (o *SDTargetingBidRecommendationsResponseV32) GetCostType() SDCostTypeV31 {
	if o == nil {
		var ret SDCostTypeV31
		return ret
	}

	return o.CostType
}

// GetCostTypeOk returns a tuple with the CostType field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsResponseV32) GetCostTypeOk() (*SDCostTypeV31, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CostType, true
}

// SetCostType sets field value
func (o *SDTargetingBidRecommendationsResponseV32) SetCostType(v SDCostTypeV31) {
	o.CostType = v
}

// GetBidRecommendations returns the BidRecommendations field value
func (o *SDTargetingBidRecommendationsResponseV32) GetBidRecommendations() []SDTargetingBidRecommendationsResponseV31BidRecommendationsInner {
	if o == nil {
		var ret []SDTargetingBidRecommendationsResponseV31BidRecommendationsInner
		return ret
	}

	return o.BidRecommendations
}

// GetBidRecommendationsOk returns a tuple with the BidRecommendations field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsResponseV32) GetBidRecommendationsOk() ([]SDTargetingBidRecommendationsResponseV31BidRecommendationsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.BidRecommendations, true
}

// SetBidRecommendations sets field value
func (o *SDTargetingBidRecommendationsResponseV32) SetBidRecommendations(v []SDTargetingBidRecommendationsResponseV31BidRecommendationsInner) {
	o.BidRecommendations = v
}

func (o SDTargetingBidRecommendationsResponseV32) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bidOptimization"] = o.BidOptimization
	toSerialize["costType"] = o.CostType
	toSerialize["bidRecommendations"] = o.BidRecommendations
	return toSerialize, nil
}

type NullableSDTargetingBidRecommendationsResponseV32 struct {
	value *SDTargetingBidRecommendationsResponseV32
	isSet bool
}

func (v NullableSDTargetingBidRecommendationsResponseV32) Get() *SDTargetingBidRecommendationsResponseV32 {
	return v.value
}

func (v *NullableSDTargetingBidRecommendationsResponseV32) Set(val *SDTargetingBidRecommendationsResponseV32) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingBidRecommendationsResponseV32) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingBidRecommendationsResponseV32) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingBidRecommendationsResponseV32(val *SDTargetingBidRecommendationsResponseV32) *NullableSDTargetingBidRecommendationsResponseV32 {
	return &NullableSDTargetingBidRecommendationsResponseV32{value: val, isSet: true}
}

func (v NullableSDTargetingBidRecommendationsResponseV32) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingBidRecommendationsResponseV32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
