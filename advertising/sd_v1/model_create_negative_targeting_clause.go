package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateNegativeTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNegativeTargetingClause{}

// CreateNegativeTargetingClause struct for CreateNegativeTargetingClause
type CreateNegativeTargetingClause struct {
	State string `json:"state"`
	// The identifier of the ad group.
	AdGroupId int64 `json:"adGroupId"`
	// The expression to negatively match against. * Only one brand may be specified per targeting expression. * Only one asin may be specified per targeting expression. * To exclude a brand from a targeting expression, you must create a negative targeting expression in the same ad group as the positive targeting expression.
	Expression     []NegativeTargetingExpression `json:"expression"`
	ExpressionType string                        `json:"expressionType"`
}

type _CreateNegativeTargetingClause CreateNegativeTargetingClause

// NewCreateNegativeTargetingClause instantiates a new CreateNegativeTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNegativeTargetingClause(state string, adGroupId int64, expression []NegativeTargetingExpression, expressionType string) *CreateNegativeTargetingClause {
	this := CreateNegativeTargetingClause{}
	this.State = state
	this.AdGroupId = adGroupId
	this.Expression = expression
	this.ExpressionType = expressionType
	return &this
}

// NewCreateNegativeTargetingClauseWithDefaults instantiates a new CreateNegativeTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNegativeTargetingClauseWithDefaults() *CreateNegativeTargetingClause {
	this := CreateNegativeTargetingClause{}
	return &this
}

// GetState returns the State field value
func (o *CreateNegativeTargetingClause) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreateNegativeTargetingClause) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreateNegativeTargetingClause) SetState(v string) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *CreateNegativeTargetingClause) GetAdGroupId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *CreateNegativeTargetingClause) GetAdGroupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *CreateNegativeTargetingClause) SetAdGroupId(v int64) {
	o.AdGroupId = v
}

// GetExpression returns the Expression field value
func (o *CreateNegativeTargetingClause) GetExpression() []NegativeTargetingExpression {
	if o == nil {
		var ret []NegativeTargetingExpression
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *CreateNegativeTargetingClause) GetExpressionOk() ([]NegativeTargetingExpression, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *CreateNegativeTargetingClause) SetExpression(v []NegativeTargetingExpression) {
	o.Expression = v
}

// GetExpressionType returns the ExpressionType field value
func (o *CreateNegativeTargetingClause) GetExpressionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value
// and a boolean to check if the value has been set.
func (o *CreateNegativeTargetingClause) GetExpressionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpressionType, true
}

// SetExpressionType sets field value
func (o *CreateNegativeTargetingClause) SetExpressionType(v string) {
	o.ExpressionType = v
}

func (o CreateNegativeTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["expression"] = o.Expression
	toSerialize["expressionType"] = o.ExpressionType
	return toSerialize, nil
}

type NullableCreateNegativeTargetingClause struct {
	value *CreateNegativeTargetingClause
	isSet bool
}

func (v NullableCreateNegativeTargetingClause) Get() *CreateNegativeTargetingClause {
	return v.value
}

func (v *NullableCreateNegativeTargetingClause) Set(val *CreateNegativeTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNegativeTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNegativeTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNegativeTargetingClause(val *CreateNegativeTargetingClause) *NullableCreateNegativeTargetingClause {
	return &NullableCreateNegativeTargetingClause{value: val, isSet: true}
}

func (v NullableCreateNegativeTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateNegativeTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
