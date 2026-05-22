package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingRecommendationsRequestV35 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingRecommendationsRequestV35{}

// SDTargetingRecommendationsRequestV35 Request for targeting recommendations for API version 3.5.
type SDTargetingRecommendationsRequestV35 struct {
	Tactic SDTacticV31 `json:"tactic"`
	// A list of products for which to get targeting recommendations. This array can only contain either asins or landing pages. If landingPageUrl is used, there can only be one item in the array for each request.
	Products []SDAdvertisedProduct `json:"products"`
	// A filter to indicate which types of recommendations to request.
	TypeFilter []SDRecommendationTypeV33         `json:"typeFilter"`
	Themes     *SDTargetingRecommendationsThemes `json:"themes,omitempty"`
	// This field is optional unless the field locationExpression is present in the request. It is used for category audience targeting to specify if the audience is for views (re-marketing) or purchases (re-purchasing). The specified categories will be returned accordingly.
	CategoryType *string `json:"categoryType,omitempty"`
	// This optional field is used to specify the locations used in SD location targeting for non-Amazon sellers only at the moment. Therefore it's only supported if the product is a landing page url.
	LocationExpression []LocationExpression `json:"locationExpression,omitempty"`
}

type _SDTargetingRecommendationsRequestV35 SDTargetingRecommendationsRequestV35

// NewSDTargetingRecommendationsRequestV35 instantiates a new SDTargetingRecommendationsRequestV35 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingRecommendationsRequestV35(tactic SDTacticV31, products []SDAdvertisedProduct, typeFilter []SDRecommendationTypeV33) *SDTargetingRecommendationsRequestV35 {
	this := SDTargetingRecommendationsRequestV35{}
	this.Tactic = tactic
	this.Products = products
	this.TypeFilter = typeFilter
	return &this
}

// NewSDTargetingRecommendationsRequestV35WithDefaults instantiates a new SDTargetingRecommendationsRequestV35 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingRecommendationsRequestV35WithDefaults() *SDTargetingRecommendationsRequestV35 {
	this := SDTargetingRecommendationsRequestV35{}
	return &this
}

// GetTactic returns the Tactic field value
func (o *SDTargetingRecommendationsRequestV35) GetTactic() SDTacticV31 {
	if o == nil {
		var ret SDTacticV31
		return ret
	}

	return o.Tactic
}

// GetTacticOk returns a tuple with the Tactic field value
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsRequestV35) GetTacticOk() (*SDTacticV31, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tactic, true
}

// SetTactic sets field value
func (o *SDTargetingRecommendationsRequestV35) SetTactic(v SDTacticV31) {
	o.Tactic = v
}

// GetProducts returns the Products field value
func (o *SDTargetingRecommendationsRequestV35) GetProducts() []SDAdvertisedProduct {
	if o == nil {
		var ret []SDAdvertisedProduct
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsRequestV35) GetProductsOk() ([]SDAdvertisedProduct, bool) {
	if o == nil {
		return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *SDTargetingRecommendationsRequestV35) SetProducts(v []SDAdvertisedProduct) {
	o.Products = v
}

// GetTypeFilter returns the TypeFilter field value
func (o *SDTargetingRecommendationsRequestV35) GetTypeFilter() []SDRecommendationTypeV33 {
	if o == nil {
		var ret []SDRecommendationTypeV33
		return ret
	}

	return o.TypeFilter
}

// GetTypeFilterOk returns a tuple with the TypeFilter field value
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsRequestV35) GetTypeFilterOk() ([]SDRecommendationTypeV33, bool) {
	if o == nil {
		return nil, false
	}
	return o.TypeFilter, true
}

// SetTypeFilter sets field value
func (o *SDTargetingRecommendationsRequestV35) SetTypeFilter(v []SDRecommendationTypeV33) {
	o.TypeFilter = v
}

// GetThemes returns the Themes field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsRequestV35) GetThemes() SDTargetingRecommendationsThemes {
	if o == nil || IsNil(o.Themes) {
		var ret SDTargetingRecommendationsThemes
		return ret
	}
	return *o.Themes
}

// GetThemesOk returns a tuple with the Themes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsRequestV35) GetThemesOk() (*SDTargetingRecommendationsThemes, bool) {
	if o == nil || IsNil(o.Themes) {
		return nil, false
	}
	return o.Themes, true
}

// HasThemes returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsRequestV35) HasThemes() bool {
	if o != nil && !IsNil(o.Themes) {
		return true
	}

	return false
}

// SetThemes gets a reference to the given SDTargetingRecommendationsThemes and assigns it to the Themes field.
func (o *SDTargetingRecommendationsRequestV35) SetThemes(v SDTargetingRecommendationsThemes) {
	o.Themes = &v
}

// GetCategoryType returns the CategoryType field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsRequestV35) GetCategoryType() string {
	if o == nil || IsNil(o.CategoryType) {
		var ret string
		return ret
	}
	return *o.CategoryType
}

// GetCategoryTypeOk returns a tuple with the CategoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsRequestV35) GetCategoryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryType) {
		return nil, false
	}
	return o.CategoryType, true
}

// HasCategoryType returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsRequestV35) HasCategoryType() bool {
	if o != nil && !IsNil(o.CategoryType) {
		return true
	}

	return false
}

// SetCategoryType gets a reference to the given string and assigns it to the CategoryType field.
func (o *SDTargetingRecommendationsRequestV35) SetCategoryType(v string) {
	o.CategoryType = &v
}

// GetLocationExpression returns the LocationExpression field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsRequestV35) GetLocationExpression() []LocationExpression {
	if o == nil || IsNil(o.LocationExpression) {
		var ret []LocationExpression
		return ret
	}
	return o.LocationExpression
}

// GetLocationExpressionOk returns a tuple with the LocationExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsRequestV35) GetLocationExpressionOk() ([]LocationExpression, bool) {
	if o == nil || IsNil(o.LocationExpression) {
		return nil, false
	}
	return o.LocationExpression, true
}

// HasLocationExpression returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsRequestV35) HasLocationExpression() bool {
	if o != nil && !IsNil(o.LocationExpression) {
		return true
	}

	return false
}

// SetLocationExpression gets a reference to the given []LocationExpression and assigns it to the LocationExpression field.
func (o *SDTargetingRecommendationsRequestV35) SetLocationExpression(v []LocationExpression) {
	o.LocationExpression = v
}

func (o SDTargetingRecommendationsRequestV35) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tactic"] = o.Tactic
	toSerialize["products"] = o.Products
	toSerialize["typeFilter"] = o.TypeFilter
	if !IsNil(o.Themes) {
		toSerialize["themes"] = o.Themes
	}
	if !IsNil(o.CategoryType) {
		toSerialize["categoryType"] = o.CategoryType
	}
	if !IsNil(o.LocationExpression) {
		toSerialize["locationExpression"] = o.LocationExpression
	}
	return toSerialize, nil
}

type NullableSDTargetingRecommendationsRequestV35 struct {
	value *SDTargetingRecommendationsRequestV35
	isSet bool
}

func (v NullableSDTargetingRecommendationsRequestV35) Get() *SDTargetingRecommendationsRequestV35 {
	return v.value
}

func (v *NullableSDTargetingRecommendationsRequestV35) Set(val *SDTargetingRecommendationsRequestV35) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendationsRequestV35) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendationsRequestV35) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendationsRequestV35(val *SDTargetingRecommendationsRequestV35) *NullableSDTargetingRecommendationsRequestV35 {
	return &NullableSDTargetingRecommendationsRequestV35{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendationsRequestV35) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendationsRequestV35) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
