package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the TargetingRecommendationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetingRecommendationsRequest{}

// TargetingRecommendationsRequest Request for targeting recommendations
type TargetingRecommendationsRequest struct {
	Tactic Tactic `json:"tactic"`
	// A list of products for which to get targeting recommendations
	Products []GoalProduct `json:"products"`
	// A filter to indicate which types of recommendations to request. T00030 only allow \"CATEGORY\".
	TypeFilter []RecommendationType `json:"typeFilter"`
}

type _TargetingRecommendationsRequest TargetingRecommendationsRequest

// NewTargetingRecommendationsRequest instantiates a new TargetingRecommendationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetingRecommendationsRequest(tactic Tactic, products []GoalProduct, typeFilter []RecommendationType) *TargetingRecommendationsRequest {
	this := TargetingRecommendationsRequest{}
	this.Tactic = tactic
	this.Products = products
	this.TypeFilter = typeFilter
	return &this
}

// NewTargetingRecommendationsRequestWithDefaults instantiates a new TargetingRecommendationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetingRecommendationsRequestWithDefaults() *TargetingRecommendationsRequest {
	this := TargetingRecommendationsRequest{}
	return &this
}

// GetTactic returns the Tactic field value
func (o *TargetingRecommendationsRequest) GetTactic() Tactic {
	if o == nil {
		var ret Tactic
		return ret
	}

	return o.Tactic
}

// GetTacticOk returns a tuple with the Tactic field value
// and a boolean to check if the value has been set.
func (o *TargetingRecommendationsRequest) GetTacticOk() (*Tactic, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tactic, true
}

// SetTactic sets field value
func (o *TargetingRecommendationsRequest) SetTactic(v Tactic) {
	o.Tactic = v
}

// GetProducts returns the Products field value
func (o *TargetingRecommendationsRequest) GetProducts() []GoalProduct {
	if o == nil {
		var ret []GoalProduct
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *TargetingRecommendationsRequest) GetProductsOk() ([]GoalProduct, bool) {
	if o == nil {
		return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *TargetingRecommendationsRequest) SetProducts(v []GoalProduct) {
	o.Products = v
}

// GetTypeFilter returns the TypeFilter field value
func (o *TargetingRecommendationsRequest) GetTypeFilter() []RecommendationType {
	if o == nil {
		var ret []RecommendationType
		return ret
	}

	return o.TypeFilter
}

// GetTypeFilterOk returns a tuple with the TypeFilter field value
// and a boolean to check if the value has been set.
func (o *TargetingRecommendationsRequest) GetTypeFilterOk() ([]RecommendationType, bool) {
	if o == nil {
		return nil, false
	}
	return o.TypeFilter, true
}

// SetTypeFilter sets field value
func (o *TargetingRecommendationsRequest) SetTypeFilter(v []RecommendationType) {
	o.TypeFilter = v
}

func (o TargetingRecommendationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tactic"] = o.Tactic
	toSerialize["products"] = o.Products
	toSerialize["typeFilter"] = o.TypeFilter
	return toSerialize, nil
}

type NullableTargetingRecommendationsRequest struct {
	value *TargetingRecommendationsRequest
	isSet bool
}

func (v NullableTargetingRecommendationsRequest) Get() *TargetingRecommendationsRequest {
	return v.value
}

func (v *NullableTargetingRecommendationsRequest) Set(val *TargetingRecommendationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetingRecommendationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetingRecommendationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetingRecommendationsRequest(val *TargetingRecommendationsRequest) *NullableTargetingRecommendationsRequest {
	return &NullableTargetingRecommendationsRequest{value: val, isSet: true}
}

func (v NullableTargetingRecommendationsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTargetingRecommendationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
