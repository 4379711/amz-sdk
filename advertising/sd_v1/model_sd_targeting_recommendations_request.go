package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingRecommendationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingRecommendationsRequest{}

// SDTargetingRecommendationsRequest Request for targeting recommendations
type SDTargetingRecommendationsRequest struct {
	Tactic SDTactic `json:"tactic"`
	// A list of products for which to get targeting recommendations
	Products []SDGoalProduct `json:"products"`
	// A filter to indicate which types of recommendations to request.
	TypeFilter []SDRecommendationType `json:"typeFilter"`
}

type _SDTargetingRecommendationsRequest SDTargetingRecommendationsRequest

// NewSDTargetingRecommendationsRequest instantiates a new SDTargetingRecommendationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingRecommendationsRequest(tactic SDTactic, products []SDGoalProduct, typeFilter []SDRecommendationType) *SDTargetingRecommendationsRequest {
	this := SDTargetingRecommendationsRequest{}
	this.Tactic = tactic
	this.Products = products
	this.TypeFilter = typeFilter
	return &this
}

// NewSDTargetingRecommendationsRequestWithDefaults instantiates a new SDTargetingRecommendationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingRecommendationsRequestWithDefaults() *SDTargetingRecommendationsRequest {
	this := SDTargetingRecommendationsRequest{}
	return &this
}

// GetTactic returns the Tactic field value
func (o *SDTargetingRecommendationsRequest) GetTactic() SDTactic {
	if o == nil {
		var ret SDTactic
		return ret
	}

	return o.Tactic
}

// GetTacticOk returns a tuple with the Tactic field value
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsRequest) GetTacticOk() (*SDTactic, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tactic, true
}

// SetTactic sets field value
func (o *SDTargetingRecommendationsRequest) SetTactic(v SDTactic) {
	o.Tactic = v
}

// GetProducts returns the Products field value
func (o *SDTargetingRecommendationsRequest) GetProducts() []SDGoalProduct {
	if o == nil {
		var ret []SDGoalProduct
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsRequest) GetProductsOk() ([]SDGoalProduct, bool) {
	if o == nil {
		return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *SDTargetingRecommendationsRequest) SetProducts(v []SDGoalProduct) {
	o.Products = v
}

// GetTypeFilter returns the TypeFilter field value
func (o *SDTargetingRecommendationsRequest) GetTypeFilter() []SDRecommendationType {
	if o == nil {
		var ret []SDRecommendationType
		return ret
	}

	return o.TypeFilter
}

// GetTypeFilterOk returns a tuple with the TypeFilter field value
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsRequest) GetTypeFilterOk() ([]SDRecommendationType, bool) {
	if o == nil {
		return nil, false
	}
	return o.TypeFilter, true
}

// SetTypeFilter sets field value
func (o *SDTargetingRecommendationsRequest) SetTypeFilter(v []SDRecommendationType) {
	o.TypeFilter = v
}

func (o SDTargetingRecommendationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tactic"] = o.Tactic
	toSerialize["products"] = o.Products
	toSerialize["typeFilter"] = o.TypeFilter
	return toSerialize, nil
}

type NullableSDTargetingRecommendationsRequest struct {
	value *SDTargetingRecommendationsRequest
	isSet bool
}

func (v NullableSDTargetingRecommendationsRequest) Get() *SDTargetingRecommendationsRequest {
	return v.value
}

func (v *NullableSDTargetingRecommendationsRequest) Set(val *SDTargetingRecommendationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendationsRequest(val *SDTargetingRecommendationsRequest) *NullableSDTargetingRecommendationsRequest {
	return &NullableSDTargetingRecommendationsRequest{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendationsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
