package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDProductTargetingRecommendationsSuccessV34 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDProductTargetingRecommendationsSuccessV34{}

// SDProductTargetingRecommendationsSuccessV34 Recommendation results for contextual targeting.
type SDProductTargetingRecommendationsSuccessV34 struct {
	// HTTP status code 200 indicating a successful response for product recommendations.
	Code *string `json:"code,omitempty"`
	// The theme name specified in the request.
	Name *string `json:"name,omitempty"`
	// A list of expressions defining the product targeting theme. The list will define an AND operator on different expressions. For example, asinPriceGreaterThan and asinReviewRatingLessThan can be used to request product recommendations which are both with greater price and less review rating compared to the goal products. Note: currently the service only support one item in the array.
	Expression []SDProductTargetingThemeExpression `json:"expression,omitempty"`
	// A list of recommended products.
	Recommendations []SDProductRecommendationV32 `json:"recommendations,omitempty"`
}

// NewSDProductTargetingRecommendationsSuccessV34 instantiates a new SDProductTargetingRecommendationsSuccessV34 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDProductTargetingRecommendationsSuccessV34() *SDProductTargetingRecommendationsSuccessV34 {
	this := SDProductTargetingRecommendationsSuccessV34{}
	return &this
}

// NewSDProductTargetingRecommendationsSuccessV34WithDefaults instantiates a new SDProductTargetingRecommendationsSuccessV34 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDProductTargetingRecommendationsSuccessV34WithDefaults() *SDProductTargetingRecommendationsSuccessV34 {
	this := SDProductTargetingRecommendationsSuccessV34{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SDProductTargetingRecommendationsSuccessV34) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDProductTargetingRecommendationsSuccessV34) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SDProductTargetingRecommendationsSuccessV34) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SDProductTargetingRecommendationsSuccessV34) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SDProductTargetingRecommendationsSuccessV34) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDProductTargetingRecommendationsSuccessV34) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SDProductTargetingRecommendationsSuccessV34) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SDProductTargetingRecommendationsSuccessV34) SetName(v string) {
	o.Name = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *SDProductTargetingRecommendationsSuccessV34) GetExpression() []SDProductTargetingThemeExpression {
	if o == nil || IsNil(o.Expression) {
		var ret []SDProductTargetingThemeExpression
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDProductTargetingRecommendationsSuccessV34) GetExpressionOk() ([]SDProductTargetingThemeExpression, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *SDProductTargetingRecommendationsSuccessV34) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given []SDProductTargetingThemeExpression and assigns it to the Expression field.
func (o *SDProductTargetingRecommendationsSuccessV34) SetExpression(v []SDProductTargetingThemeExpression) {
	o.Expression = v
}

// GetRecommendations returns the Recommendations field value if set, zero value otherwise.
func (o *SDProductTargetingRecommendationsSuccessV34) GetRecommendations() []SDProductRecommendationV32 {
	if o == nil || IsNil(o.Recommendations) {
		var ret []SDProductRecommendationV32
		return ret
	}
	return o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDProductTargetingRecommendationsSuccessV34) GetRecommendationsOk() ([]SDProductRecommendationV32, bool) {
	if o == nil || IsNil(o.Recommendations) {
		return nil, false
	}
	return o.Recommendations, true
}

// HasRecommendations returns a boolean if a field has been set.
func (o *SDProductTargetingRecommendationsSuccessV34) HasRecommendations() bool {
	if o != nil && !IsNil(o.Recommendations) {
		return true
	}

	return false
}

// SetRecommendations gets a reference to the given []SDProductRecommendationV32 and assigns it to the Recommendations field.
func (o *SDProductTargetingRecommendationsSuccessV34) SetRecommendations(v []SDProductRecommendationV32) {
	o.Recommendations = v
}

func (o SDProductTargetingRecommendationsSuccessV34) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !IsNil(o.Recommendations) {
		toSerialize["recommendations"] = o.Recommendations
	}
	return toSerialize, nil
}

type NullableSDProductTargetingRecommendationsSuccessV34 struct {
	value *SDProductTargetingRecommendationsSuccessV34
	isSet bool
}

func (v NullableSDProductTargetingRecommendationsSuccessV34) Get() *SDProductTargetingRecommendationsSuccessV34 {
	return v.value
}

func (v *NullableSDProductTargetingRecommendationsSuccessV34) Set(val *SDProductTargetingRecommendationsSuccessV34) {
	v.value = val
	v.isSet = true
}

func (v NullableSDProductTargetingRecommendationsSuccessV34) IsSet() bool {
	return v.isSet
}

func (v *NullableSDProductTargetingRecommendationsSuccessV34) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDProductTargetingRecommendationsSuccessV34(val *SDProductTargetingRecommendationsSuccessV34) *NullableSDProductTargetingRecommendationsSuccessV34 {
	return &NullableSDProductTargetingRecommendationsSuccessV34{value: val, isSet: true}
}

func (v NullableSDProductTargetingRecommendationsSuccessV34) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDProductTargetingRecommendationsSuccessV34) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
