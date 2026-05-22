package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingBidRecommendationsRequestV31 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingBidRecommendationsRequestV31{}

// SDTargetingBidRecommendationsRequestV31 Request for targeting bid recommendations.
type SDTargetingBidRecommendationsRequestV31 struct {
	// A list of products to tailor bid recommendations for category and audience based targeting clauses.
	Products []SDGoalProduct `json:"products,omitempty"`
	// A list of targeting clauses to receive bid recommendations for.
	TargetingClauses []SDTargetingBidRecommendationsRequestV31TargetingClausesInner `json:"targetingClauses"`
}

type _SDTargetingBidRecommendationsRequestV31 SDTargetingBidRecommendationsRequestV31

// NewSDTargetingBidRecommendationsRequestV31 instantiates a new SDTargetingBidRecommendationsRequestV31 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingBidRecommendationsRequestV31(targetingClauses []SDTargetingBidRecommendationsRequestV31TargetingClausesInner) *SDTargetingBidRecommendationsRequestV31 {
	this := SDTargetingBidRecommendationsRequestV31{}
	this.TargetingClauses = targetingClauses
	return &this
}

// NewSDTargetingBidRecommendationsRequestV31WithDefaults instantiates a new SDTargetingBidRecommendationsRequestV31 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingBidRecommendationsRequestV31WithDefaults() *SDTargetingBidRecommendationsRequestV31 {
	this := SDTargetingBidRecommendationsRequestV31{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *SDTargetingBidRecommendationsRequestV31) GetProducts() []SDGoalProduct {
	if o == nil || IsNil(o.Products) {
		var ret []SDGoalProduct
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsRequestV31) GetProductsOk() ([]SDGoalProduct, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *SDTargetingBidRecommendationsRequestV31) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []SDGoalProduct and assigns it to the Products field.
func (o *SDTargetingBidRecommendationsRequestV31) SetProducts(v []SDGoalProduct) {
	o.Products = v
}

// GetTargetingClauses returns the TargetingClauses field value
func (o *SDTargetingBidRecommendationsRequestV31) GetTargetingClauses() []SDTargetingBidRecommendationsRequestV31TargetingClausesInner {
	if o == nil {
		var ret []SDTargetingBidRecommendationsRequestV31TargetingClausesInner
		return ret
	}

	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsRequestV31) GetTargetingClausesOk() ([]SDTargetingBidRecommendationsRequestV31TargetingClausesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetingClauses, true
}

// SetTargetingClauses sets field value
func (o *SDTargetingBidRecommendationsRequestV31) SetTargetingClauses(v []SDTargetingBidRecommendationsRequestV31TargetingClausesInner) {
	o.TargetingClauses = v
}

func (o SDTargetingBidRecommendationsRequestV31) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	toSerialize["targetingClauses"] = o.TargetingClauses
	return toSerialize, nil
}

type NullableSDTargetingBidRecommendationsRequestV31 struct {
	value *SDTargetingBidRecommendationsRequestV31
	isSet bool
}

func (v NullableSDTargetingBidRecommendationsRequestV31) Get() *SDTargetingBidRecommendationsRequestV31 {
	return v.value
}

func (v *NullableSDTargetingBidRecommendationsRequestV31) Set(val *SDTargetingBidRecommendationsRequestV31) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingBidRecommendationsRequestV31) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingBidRecommendationsRequestV31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingBidRecommendationsRequestV31(val *SDTargetingBidRecommendationsRequestV31) *NullableSDTargetingBidRecommendationsRequestV31 {
	return &NullableSDTargetingBidRecommendationsRequestV31{value: val, isSet: true}
}

func (v NullableSDTargetingBidRecommendationsRequestV31) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingBidRecommendationsRequestV31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
