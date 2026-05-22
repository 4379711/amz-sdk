package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingClauseV31 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingClauseV31{}

// SDTargetingClauseV31 The targeting clause
type SDTargetingClauseV31 struct {
	// Tactic T00020 ad groups only allow manual targeting.
	ExpressionType string `json:"expressionType"`
	// The targeting expression to match against.  ------- Applicable to contextual targeting (T00020) ------- * A 'TargetingExpression' in a contextual targeting campaign can only contain 'TargetingPredicate' components. * Expressions must specify either a category predicate or an ASIN predicate, but never both. * Only one category may be specified per targeting expression. * Only one brand may be specified per targeting expression. * Only one asin may be specified per targeting expression. * To exclude a brand from a targeting expression you must create a negative targeting expression in the same ad group as the positive targeting expression.  ------- Applicable to audiences or contextual targeting (T00030) ------- * A 'TargetingExpression' in a audiences or contextual campaign can contain any target, including 'TargetingPredicate', 'ContentTargetingPredicate', or 'TargetingPredicateNested'.
	Expression []SDTargetExpressionV31 `json:"expression"`
}

type _SDTargetingClauseV31 SDTargetingClauseV31

// NewSDTargetingClauseV31 instantiates a new SDTargetingClauseV31 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingClauseV31(expressionType string, expression []SDTargetExpressionV31) *SDTargetingClauseV31 {
	this := SDTargetingClauseV31{}
	this.ExpressionType = expressionType
	this.Expression = expression
	return &this
}

// NewSDTargetingClauseV31WithDefaults instantiates a new SDTargetingClauseV31 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingClauseV31WithDefaults() *SDTargetingClauseV31 {
	this := SDTargetingClauseV31{}
	return &this
}

// GetExpressionType returns the ExpressionType field value
func (o *SDTargetingClauseV31) GetExpressionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value
// and a boolean to check if the value has been set.
func (o *SDTargetingClauseV31) GetExpressionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpressionType, true
}

// SetExpressionType sets field value
func (o *SDTargetingClauseV31) SetExpressionType(v string) {
	o.ExpressionType = v
}

// GetExpression returns the Expression field value
func (o *SDTargetingClauseV31) GetExpression() []SDTargetExpressionV31 {
	if o == nil {
		var ret []SDTargetExpressionV31
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *SDTargetingClauseV31) GetExpressionOk() ([]SDTargetExpressionV31, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *SDTargetingClauseV31) SetExpression(v []SDTargetExpressionV31) {
	o.Expression = v
}

func (o SDTargetingClauseV31) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expressionType"] = o.ExpressionType
	toSerialize["expression"] = o.Expression
	return toSerialize, nil
}

type NullableSDTargetingClauseV31 struct {
	value *SDTargetingClauseV31
	isSet bool
}

func (v NullableSDTargetingClauseV31) Get() *SDTargetingClauseV31 {
	return v.value
}

func (v *NullableSDTargetingClauseV31) Set(val *SDTargetingClauseV31) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingClauseV31) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingClauseV31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingClauseV31(val *SDTargetingClauseV31) *NullableSDTargetingClauseV31 {
	return &NullableSDTargetingClauseV31{value: val, isSet: true}
}

func (v NullableSDTargetingClauseV31) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingClauseV31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
