package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingRecommendationsRequestV34 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingRecommendationsRequestV34{}

// SDTargetingRecommendationsRequestV34 Request for targeting recommendations for API version 3.4.
type SDTargetingRecommendationsRequestV34 struct {
	Tactic SDTacticV31 `json:"tactic"`
	// A list of products for which to get targeting recommendations
	Products []SDGoalProduct `json:"products"`
	// A filter to indicate which types of recommendations to request.
	TypeFilter []SDRecommendationTypeV32         `json:"typeFilter"`
	Themes     *SDTargetingRecommendationsThemes `json:"themes,omitempty"`
}

type _SDTargetingRecommendationsRequestV34 SDTargetingRecommendationsRequestV34

// NewSDTargetingRecommendationsRequestV34 instantiates a new SDTargetingRecommendationsRequestV34 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingRecommendationsRequestV34(tactic SDTacticV31, products []SDGoalProduct, typeFilter []SDRecommendationTypeV32) *SDTargetingRecommendationsRequestV34 {
	this := SDTargetingRecommendationsRequestV34{}
	this.Tactic = tactic
	this.Products = products
	this.TypeFilter = typeFilter
	return &this
}

// NewSDTargetingRecommendationsRequestV34WithDefaults instantiates a new SDTargetingRecommendationsRequestV34 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingRecommendationsRequestV34WithDefaults() *SDTargetingRecommendationsRequestV34 {
	this := SDTargetingRecommendationsRequestV34{}
	return &this
}

// GetTactic returns the Tactic field value
func (o *SDTargetingRecommendationsRequestV34) GetTactic() SDTacticV31 {
	if o == nil {
		var ret SDTacticV31
		return ret
	}

	return o.Tactic
}

// GetTacticOk returns a tuple with the Tactic field value
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsRequestV34) GetTacticOk() (*SDTacticV31, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tactic, true
}

// SetTactic sets field value
func (o *SDTargetingRecommendationsRequestV34) SetTactic(v SDTacticV31) {
	o.Tactic = v
}

// GetProducts returns the Products field value
func (o *SDTargetingRecommendationsRequestV34) GetProducts() []SDGoalProduct {
	if o == nil {
		var ret []SDGoalProduct
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsRequestV34) GetProductsOk() ([]SDGoalProduct, bool) {
	if o == nil {
		return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *SDTargetingRecommendationsRequestV34) SetProducts(v []SDGoalProduct) {
	o.Products = v
}

// GetTypeFilter returns the TypeFilter field value
func (o *SDTargetingRecommendationsRequestV34) GetTypeFilter() []SDRecommendationTypeV32 {
	if o == nil {
		var ret []SDRecommendationTypeV32
		return ret
	}

	return o.TypeFilter
}

// GetTypeFilterOk returns a tuple with the TypeFilter field value
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsRequestV34) GetTypeFilterOk() ([]SDRecommendationTypeV32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TypeFilter, true
}

// SetTypeFilter sets field value
func (o *SDTargetingRecommendationsRequestV34) SetTypeFilter(v []SDRecommendationTypeV32) {
	o.TypeFilter = v
}

// GetThemes returns the Themes field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsRequestV34) GetThemes() SDTargetingRecommendationsThemes {
	if o == nil || IsNil(o.Themes) {
		var ret SDTargetingRecommendationsThemes
		return ret
	}
	return *o.Themes
}

// GetThemesOk returns a tuple with the Themes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsRequestV34) GetThemesOk() (*SDTargetingRecommendationsThemes, bool) {
	if o == nil || IsNil(o.Themes) {
		return nil, false
	}
	return o.Themes, true
}

// HasThemes returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsRequestV34) HasThemes() bool {
	if o != nil && !IsNil(o.Themes) {
		return true
	}

	return false
}

// SetThemes gets a reference to the given SDTargetingRecommendationsThemes and assigns it to the Themes field.
func (o *SDTargetingRecommendationsRequestV34) SetThemes(v SDTargetingRecommendationsThemes) {
	o.Themes = &v
}

func (o SDTargetingRecommendationsRequestV34) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tactic"] = o.Tactic
	toSerialize["products"] = o.Products
	toSerialize["typeFilter"] = o.TypeFilter
	if !IsNil(o.Themes) {
		toSerialize["themes"] = o.Themes
	}
	return toSerialize, nil
}

type NullableSDTargetingRecommendationsRequestV34 struct {
	value *SDTargetingRecommendationsRequestV34
	isSet bool
}

func (v NullableSDTargetingRecommendationsRequestV34) Get() *SDTargetingRecommendationsRequestV34 {
	return v.value
}

func (v *NullableSDTargetingRecommendationsRequestV34) Set(val *SDTargetingRecommendationsRequestV34) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendationsRequestV34) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendationsRequestV34) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendationsRequestV34(val *SDTargetingRecommendationsRequestV34) *NullableSDTargetingRecommendationsRequestV34 {
	return &NullableSDTargetingRecommendationsRequestV34{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendationsRequestV34) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendationsRequestV34) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
