package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingRecommendationsRequestV31 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingRecommendationsRequestV31{}

// SDTargetingRecommendationsRequestV31 Request for targeting recommendations
type SDTargetingRecommendationsRequestV31 struct {
	Tactic SDTacticV31 `json:"tactic"`
	// A list of products for which to get targeting recommendations
	Products []SDGoalProduct `json:"products"`
	// A filter to indicate which types of recommendations to request.
	TypeFilter []SDRecommendationTypeV31 `json:"typeFilter"`
}

type _SDTargetingRecommendationsRequestV31 SDTargetingRecommendationsRequestV31

// NewSDTargetingRecommendationsRequestV31 instantiates a new SDTargetingRecommendationsRequestV31 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingRecommendationsRequestV31(tactic SDTacticV31, products []SDGoalProduct, typeFilter []SDRecommendationTypeV31) *SDTargetingRecommendationsRequestV31 {
	this := SDTargetingRecommendationsRequestV31{}
	this.Tactic = tactic
	this.Products = products
	this.TypeFilter = typeFilter
	return &this
}

// NewSDTargetingRecommendationsRequestV31WithDefaults instantiates a new SDTargetingRecommendationsRequestV31 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingRecommendationsRequestV31WithDefaults() *SDTargetingRecommendationsRequestV31 {
	this := SDTargetingRecommendationsRequestV31{}
	return &this
}

// GetTactic returns the Tactic field value
func (o *SDTargetingRecommendationsRequestV31) GetTactic() SDTacticV31 {
	if o == nil {
		var ret SDTacticV31
		return ret
	}

	return o.Tactic
}

// GetTacticOk returns a tuple with the Tactic field value
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsRequestV31) GetTacticOk() (*SDTacticV31, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tactic, true
}

// SetTactic sets field value
func (o *SDTargetingRecommendationsRequestV31) SetTactic(v SDTacticV31) {
	o.Tactic = v
}

// GetProducts returns the Products field value
func (o *SDTargetingRecommendationsRequestV31) GetProducts() []SDGoalProduct {
	if o == nil {
		var ret []SDGoalProduct
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsRequestV31) GetProductsOk() ([]SDGoalProduct, bool) {
	if o == nil {
		return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *SDTargetingRecommendationsRequestV31) SetProducts(v []SDGoalProduct) {
	o.Products = v
}

// GetTypeFilter returns the TypeFilter field value
func (o *SDTargetingRecommendationsRequestV31) GetTypeFilter() []SDRecommendationTypeV31 {
	if o == nil {
		var ret []SDRecommendationTypeV31
		return ret
	}

	return o.TypeFilter
}

// GetTypeFilterOk returns a tuple with the TypeFilter field value
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsRequestV31) GetTypeFilterOk() ([]SDRecommendationTypeV31, bool) {
	if o == nil {
		return nil, false
	}
	return o.TypeFilter, true
}

// SetTypeFilter sets field value
func (o *SDTargetingRecommendationsRequestV31) SetTypeFilter(v []SDRecommendationTypeV31) {
	o.TypeFilter = v
}

func (o SDTargetingRecommendationsRequestV31) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tactic"] = o.Tactic
	toSerialize["products"] = o.Products
	toSerialize["typeFilter"] = o.TypeFilter
	return toSerialize, nil
}

type NullableSDTargetingRecommendationsRequestV31 struct {
	value *SDTargetingRecommendationsRequestV31
	isSet bool
}

func (v NullableSDTargetingRecommendationsRequestV31) Get() *SDTargetingRecommendationsRequestV31 {
	return v.value
}

func (v *NullableSDTargetingRecommendationsRequestV31) Set(val *SDTargetingRecommendationsRequestV31) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendationsRequestV31) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendationsRequestV31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendationsRequestV31(val *SDTargetingRecommendationsRequestV31) *NullableSDTargetingRecommendationsRequestV31 {
	return &NullableSDTargetingRecommendationsRequestV31{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendationsRequestV31) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendationsRequestV31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
