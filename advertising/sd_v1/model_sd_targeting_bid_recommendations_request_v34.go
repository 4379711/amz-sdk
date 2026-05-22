package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingBidRecommendationsRequestV34 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingBidRecommendationsRequestV34{}

// SDTargetingBidRecommendationsRequestV34 Request for targeting bid recommendations.
type SDTargetingBidRecommendationsRequestV34 struct {
	// A list of products to tailor bid recommendations for category and audience based targeting clauses. This array must contain consistent fields of either asins or landing pages (when linking to other pages), these cannot be mixed for any given request. If landingPageUrl is used, only one item is allowed for the list.
	Products        []SDGoalProduct        `json:"products,omitempty"`
	BidOptimization SDBidOptimizationV32   `json:"bidOptimization"`
	CostType        SDCostTypeV31          `json:"costType"`
	CreativeType    NullableSDCreativeType `json:"creativeType,omitempty"`
	// A list of targeting clauses to receive bid recommendations for.
	TargetingClauses []SDTargetingBidRecommendationsRequestV34TargetingClausesInner `json:"targetingClauses"`
}

type _SDTargetingBidRecommendationsRequestV34 SDTargetingBidRecommendationsRequestV34

// NewSDTargetingBidRecommendationsRequestV34 instantiates a new SDTargetingBidRecommendationsRequestV34 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingBidRecommendationsRequestV34(bidOptimization SDBidOptimizationV32, costType SDCostTypeV31, targetingClauses []SDTargetingBidRecommendationsRequestV34TargetingClausesInner) *SDTargetingBidRecommendationsRequestV34 {
	this := SDTargetingBidRecommendationsRequestV34{}
	this.BidOptimization = bidOptimization
	this.CostType = costType
	this.TargetingClauses = targetingClauses
	return &this
}

// NewSDTargetingBidRecommendationsRequestV34WithDefaults instantiates a new SDTargetingBidRecommendationsRequestV34 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingBidRecommendationsRequestV34WithDefaults() *SDTargetingBidRecommendationsRequestV34 {
	this := SDTargetingBidRecommendationsRequestV34{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *SDTargetingBidRecommendationsRequestV34) GetProducts() []SDGoalProduct {
	if o == nil || IsNil(o.Products) {
		var ret []SDGoalProduct
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsRequestV34) GetProductsOk() ([]SDGoalProduct, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *SDTargetingBidRecommendationsRequestV34) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []SDGoalProduct and assigns it to the Products field.
func (o *SDTargetingBidRecommendationsRequestV34) SetProducts(v []SDGoalProduct) {
	o.Products = v
}

// GetBidOptimization returns the BidOptimization field value
func (o *SDTargetingBidRecommendationsRequestV34) GetBidOptimization() SDBidOptimizationV32 {
	if o == nil {
		var ret SDBidOptimizationV32
		return ret
	}

	return o.BidOptimization
}

// GetBidOptimizationOk returns a tuple with the BidOptimization field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsRequestV34) GetBidOptimizationOk() (*SDBidOptimizationV32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BidOptimization, true
}

// SetBidOptimization sets field value
func (o *SDTargetingBidRecommendationsRequestV34) SetBidOptimization(v SDBidOptimizationV32) {
	o.BidOptimization = v
}

// GetCostType returns the CostType field value
func (o *SDTargetingBidRecommendationsRequestV34) GetCostType() SDCostTypeV31 {
	if o == nil {
		var ret SDCostTypeV31
		return ret
	}

	return o.CostType
}

// GetCostTypeOk returns a tuple with the CostType field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsRequestV34) GetCostTypeOk() (*SDCostTypeV31, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CostType, true
}

// SetCostType sets field value
func (o *SDTargetingBidRecommendationsRequestV34) SetCostType(v SDCostTypeV31) {
	o.CostType = v
}

// GetCreativeType returns the CreativeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SDTargetingBidRecommendationsRequestV34) GetCreativeType() SDCreativeType {
	if o == nil || IsNil(o.CreativeType.Get()) {
		var ret SDCreativeType
		return ret
	}
	return *o.CreativeType.Get()
}

// GetCreativeTypeOk returns a tuple with the CreativeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SDTargetingBidRecommendationsRequestV34) GetCreativeTypeOk() (*SDCreativeType, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreativeType.Get(), o.CreativeType.IsSet()
}

// HasCreativeType returns a boolean if a field has been set.
func (o *SDTargetingBidRecommendationsRequestV34) HasCreativeType() bool {
	if o != nil && o.CreativeType.IsSet() {
		return true
	}

	return false
}

// SetCreativeType gets a reference to the given NullableSDCreativeType and assigns it to the CreativeType field.
func (o *SDTargetingBidRecommendationsRequestV34) SetCreativeType(v SDCreativeType) {
	o.CreativeType.Set(&v)
}

// SetCreativeTypeNil sets the value for CreativeType to be an explicit nil
func (o *SDTargetingBidRecommendationsRequestV34) SetCreativeTypeNil() {
	o.CreativeType.Set(nil)
}

// UnsetCreativeType ensures that no value is present for CreativeType, not even an explicit nil
func (o *SDTargetingBidRecommendationsRequestV34) UnsetCreativeType() {
	o.CreativeType.Unset()
}

// GetTargetingClauses returns the TargetingClauses field value
func (o *SDTargetingBidRecommendationsRequestV34) GetTargetingClauses() []SDTargetingBidRecommendationsRequestV34TargetingClausesInner {
	if o == nil {
		var ret []SDTargetingBidRecommendationsRequestV34TargetingClausesInner
		return ret
	}

	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsRequestV34) GetTargetingClausesOk() ([]SDTargetingBidRecommendationsRequestV34TargetingClausesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetingClauses, true
}

// SetTargetingClauses sets field value
func (o *SDTargetingBidRecommendationsRequestV34) SetTargetingClauses(v []SDTargetingBidRecommendationsRequestV34TargetingClausesInner) {
	o.TargetingClauses = v
}

func (o SDTargetingBidRecommendationsRequestV34) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	toSerialize["bidOptimization"] = o.BidOptimization
	toSerialize["costType"] = o.CostType
	if o.CreativeType.IsSet() {
		toSerialize["creativeType"] = o.CreativeType.Get()
	}
	toSerialize["targetingClauses"] = o.TargetingClauses
	return toSerialize, nil
}

type NullableSDTargetingBidRecommendationsRequestV34 struct {
	value *SDTargetingBidRecommendationsRequestV34
	isSet bool
}

func (v NullableSDTargetingBidRecommendationsRequestV34) Get() *SDTargetingBidRecommendationsRequestV34 {
	return v.value
}

func (v *NullableSDTargetingBidRecommendationsRequestV34) Set(val *SDTargetingBidRecommendationsRequestV34) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingBidRecommendationsRequestV34) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingBidRecommendationsRequestV34) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingBidRecommendationsRequestV34(val *SDTargetingBidRecommendationsRequestV34) *NullableSDTargetingBidRecommendationsRequestV34 {
	return &NullableSDTargetingBidRecommendationsRequestV34{value: val, isSet: true}
}

func (v NullableSDTargetingBidRecommendationsRequestV34) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingBidRecommendationsRequestV34) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
