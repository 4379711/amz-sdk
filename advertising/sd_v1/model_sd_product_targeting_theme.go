package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDProductTargetingTheme type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDProductTargetingTheme{}

// SDProductTargetingTheme Contextual targeting theme definitions.
type SDProductTargetingTheme struct {
	// This is the meaningful theme name which will be used as a unique identifier across various themes in the same request. This identifier will also be used to map the recommendations back to the theme in the response body. Note: the value for this field cannot be \"default\" as that's a reserved keyword in the system.
	Name string `json:"name" validate:"regexp=^(?!default$)"`
	// A list of expressions defining the contextual targeting theme. The list will define an AND operator on different expressions. For example, asinPriceGreaterThan and asinReviewRatingLessThan can be used to request product recommendations which are both with greater price and less review rating compared to the goal products. Note: Currently the service only supports one item in the array.
	Expression []SDProductTargetingThemeExpression `json:"expression"`
}

type _SDProductTargetingTheme SDProductTargetingTheme

// NewSDProductTargetingTheme instantiates a new SDProductTargetingTheme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDProductTargetingTheme(name string, expression []SDProductTargetingThemeExpression) *SDProductTargetingTheme {
	this := SDProductTargetingTheme{}
	this.Name = name
	this.Expression = expression
	return &this
}

// NewSDProductTargetingThemeWithDefaults instantiates a new SDProductTargetingTheme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDProductTargetingThemeWithDefaults() *SDProductTargetingTheme {
	this := SDProductTargetingTheme{}
	return &this
}

// GetName returns the Name field value
func (o *SDProductTargetingTheme) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SDProductTargetingTheme) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SDProductTargetingTheme) SetName(v string) {
	o.Name = v
}

// GetExpression returns the Expression field value
func (o *SDProductTargetingTheme) GetExpression() []SDProductTargetingThemeExpression {
	if o == nil {
		var ret []SDProductTargetingThemeExpression
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *SDProductTargetingTheme) GetExpressionOk() ([]SDProductTargetingThemeExpression, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *SDProductTargetingTheme) SetExpression(v []SDProductTargetingThemeExpression) {
	o.Expression = v
}

func (o SDProductTargetingTheme) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["expression"] = o.Expression
	return toSerialize, nil
}

type NullableSDProductTargetingTheme struct {
	value *SDProductTargetingTheme
	isSet bool
}

func (v NullableSDProductTargetingTheme) Get() *SDProductTargetingTheme {
	return v.value
}

func (v *NullableSDProductTargetingTheme) Set(val *SDProductTargetingTheme) {
	v.value = val
	v.isSet = true
}

func (v NullableSDProductTargetingTheme) IsSet() bool {
	return v.isSet
}

func (v *NullableSDProductTargetingTheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDProductTargetingTheme(val *SDProductTargetingTheme) *NullableSDProductTargetingTheme {
	return &NullableSDProductTargetingTheme{value: val, isSet: true}
}

func (v NullableSDProductTargetingTheme) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDProductTargetingTheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
