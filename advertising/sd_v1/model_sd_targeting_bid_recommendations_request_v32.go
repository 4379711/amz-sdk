package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingBidRecommendationsRequestV32 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingBidRecommendationsRequestV32{}

// SDTargetingBidRecommendationsRequestV32 Request for targeting bid recommendations.
type SDTargetingBidRecommendationsRequestV32 struct {
	// A list of products to tailor bid recommendations for category and audience based targeting clauses.
	Products        []SDGoalProduct      `json:"products,omitempty"`
	BidOptimization SDBidOptimizationV32 `json:"bidOptimization"`
	CostType        SDCostTypeV31        `json:"costType"`
	// A list of targeting clauses to receive bid recommendations for.
	TargetingClauses []SDTargetingBidRecommendationsRequestV31TargetingClausesInner `json:"targetingClauses"`
}

type _SDTargetingBidRecommendationsRequestV32 SDTargetingBidRecommendationsRequestV32

// NewSDTargetingBidRecommendationsRequestV32 instantiates a new SDTargetingBidRecommendationsRequestV32 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingBidRecommendationsRequestV32(bidOptimization SDBidOptimizationV32, costType SDCostTypeV31, targetingClauses []SDTargetingBidRecommendationsRequestV31TargetingClausesInner) *SDTargetingBidRecommendationsRequestV32 {
	this := SDTargetingBidRecommendationsRequestV32{}
	this.BidOptimization = bidOptimization
	this.CostType = costType
	this.TargetingClauses = targetingClauses
	return &this
}

// NewSDTargetingBidRecommendationsRequestV32WithDefaults instantiates a new SDTargetingBidRecommendationsRequestV32 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingBidRecommendationsRequestV32WithDefaults() *SDTargetingBidRecommendationsRequestV32 {
	this := SDTargetingBidRecommendationsRequestV32{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *SDTargetingBidRecommendationsRequestV32) GetProducts() []SDGoalProduct {
	if o == nil || IsNil(o.Products) {
		var ret []SDGoalProduct
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsRequestV32) GetProductsOk() ([]SDGoalProduct, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *SDTargetingBidRecommendationsRequestV32) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []SDGoalProduct and assigns it to the Products field.
func (o *SDTargetingBidRecommendationsRequestV32) SetProducts(v []SDGoalProduct) {
	o.Products = v
}

// GetBidOptimization returns the BidOptimization field value
func (o *SDTargetingBidRecommendationsRequestV32) GetBidOptimization() SDBidOptimizationV32 {
	if o == nil {
		var ret SDBidOptimizationV32
		return ret
	}

	return o.BidOptimization
}

// GetBidOptimizationOk returns a tuple with the BidOptimization field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsRequestV32) GetBidOptimizationOk() (*SDBidOptimizationV32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BidOptimization, true
}

// SetBidOptimization sets field value
func (o *SDTargetingBidRecommendationsRequestV32) SetBidOptimization(v SDBidOptimizationV32) {
	o.BidOptimization = v
}

// GetCostType returns the CostType field value
func (o *SDTargetingBidRecommendationsRequestV32) GetCostType() SDCostTypeV31 {
	if o == nil {
		var ret SDCostTypeV31
		return ret
	}

	return o.CostType
}

// GetCostTypeOk returns a tuple with the CostType field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsRequestV32) GetCostTypeOk() (*SDCostTypeV31, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CostType, true
}

// SetCostType sets field value
func (o *SDTargetingBidRecommendationsRequestV32) SetCostType(v SDCostTypeV31) {
	o.CostType = v
}

// GetTargetingClauses returns the TargetingClauses field value
func (o *SDTargetingBidRecommendationsRequestV32) GetTargetingClauses() []SDTargetingBidRecommendationsRequestV31TargetingClausesInner {
	if o == nil {
		var ret []SDTargetingBidRecommendationsRequestV31TargetingClausesInner
		return ret
	}

	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsRequestV32) GetTargetingClausesOk() ([]SDTargetingBidRecommendationsRequestV31TargetingClausesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetingClauses, true
}

// SetTargetingClauses sets field value
func (o *SDTargetingBidRecommendationsRequestV32) SetTargetingClauses(v []SDTargetingBidRecommendationsRequestV31TargetingClausesInner) {
	o.TargetingClauses = v
}

func (o SDTargetingBidRecommendationsRequestV32) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	toSerialize["bidOptimization"] = o.BidOptimization
	toSerialize["costType"] = o.CostType
	toSerialize["targetingClauses"] = o.TargetingClauses
	return toSerialize, nil
}

type NullableSDTargetingBidRecommendationsRequestV32 struct {
	value *SDTargetingBidRecommendationsRequestV32
	isSet bool
}

func (v NullableSDTargetingBidRecommendationsRequestV32) Get() *SDTargetingBidRecommendationsRequestV32 {
	return v.value
}

func (v *NullableSDTargetingBidRecommendationsRequestV32) Set(val *SDTargetingBidRecommendationsRequestV32) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingBidRecommendationsRequestV32) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingBidRecommendationsRequestV32) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingBidRecommendationsRequestV32(val *SDTargetingBidRecommendationsRequestV32) *NullableSDTargetingBidRecommendationsRequestV32 {
	return &NullableSDTargetingBidRecommendationsRequestV32{value: val, isSet: true}
}

func (v NullableSDTargetingBidRecommendationsRequestV32) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingBidRecommendationsRequestV32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
