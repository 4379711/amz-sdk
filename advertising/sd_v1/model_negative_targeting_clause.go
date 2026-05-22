package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the NegativeTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NegativeTargetingClause{}

// NegativeTargetingClause struct for NegativeTargetingClause
type NegativeTargetingClause struct {
	State    *string `json:"state,omitempty"`
	TargetId *int64  `json:"targetId,omitempty"`
	// The identifier of the ad group.
	AdGroupId      *int64  `json:"adGroupId,omitempty"`
	ExpressionType *string `json:"expressionType,omitempty"`
	// The expression to negatively match against. * Only one brand may be specified per targeting expression. * Only one asin may be specified per targeting expression. * To exclude a brand from a targeting expression, you must create a negative targeting expression in the same ad group as the positive targeting expression.
	Expression []NegativeTargetingExpression `json:"expression,omitempty"`
	// The resolved negative targeting expression.
	ResolvedExpression []NegativeTargetingExpression `json:"resolvedExpression,omitempty"`
}

// NewNegativeTargetingClause instantiates a new NegativeTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNegativeTargetingClause() *NegativeTargetingClause {
	this := NegativeTargetingClause{}
	return &this
}

// NewNegativeTargetingClauseWithDefaults instantiates a new NegativeTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNegativeTargetingClauseWithDefaults() *NegativeTargetingClause {
	this := NegativeTargetingClause{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NegativeTargetingClause) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeTargetingClause) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NegativeTargetingClause) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NegativeTargetingClause) SetState(v string) {
	o.State = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *NegativeTargetingClause) GetTargetId() int64 {
	if o == nil || IsNil(o.TargetId) {
		var ret int64
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeTargetingClause) GetTargetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *NegativeTargetingClause) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given int64 and assigns it to the TargetId field.
func (o *NegativeTargetingClause) SetTargetId(v int64) {
	o.TargetId = &v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *NegativeTargetingClause) GetAdGroupId() int64 {
	if o == nil || IsNil(o.AdGroupId) {
		var ret int64
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeTargetingClause) GetAdGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *NegativeTargetingClause) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given int64 and assigns it to the AdGroupId field.
func (o *NegativeTargetingClause) SetAdGroupId(v int64) {
	o.AdGroupId = &v
}

// GetExpressionType returns the ExpressionType field value if set, zero value otherwise.
func (o *NegativeTargetingClause) GetExpressionType() string {
	if o == nil || IsNil(o.ExpressionType) {
		var ret string
		return ret
	}
	return *o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeTargetingClause) GetExpressionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpressionType) {
		return nil, false
	}
	return o.ExpressionType, true
}

// HasExpressionType returns a boolean if a field has been set.
func (o *NegativeTargetingClause) HasExpressionType() bool {
	if o != nil && !IsNil(o.ExpressionType) {
		return true
	}

	return false
}

// SetExpressionType gets a reference to the given string and assigns it to the ExpressionType field.
func (o *NegativeTargetingClause) SetExpressionType(v string) {
	o.ExpressionType = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *NegativeTargetingClause) GetExpression() []NegativeTargetingExpression {
	if o == nil || IsNil(o.Expression) {
		var ret []NegativeTargetingExpression
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeTargetingClause) GetExpressionOk() ([]NegativeTargetingExpression, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *NegativeTargetingClause) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given []NegativeTargetingExpression and assigns it to the Expression field.
func (o *NegativeTargetingClause) SetExpression(v []NegativeTargetingExpression) {
	o.Expression = v
}

// GetResolvedExpression returns the ResolvedExpression field value if set, zero value otherwise.
func (o *NegativeTargetingClause) GetResolvedExpression() []NegativeTargetingExpression {
	if o == nil || IsNil(o.ResolvedExpression) {
		var ret []NegativeTargetingExpression
		return ret
	}
	return o.ResolvedExpression
}

// GetResolvedExpressionOk returns a tuple with the ResolvedExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeTargetingClause) GetResolvedExpressionOk() ([]NegativeTargetingExpression, bool) {
	if o == nil || IsNil(o.ResolvedExpression) {
		return nil, false
	}
	return o.ResolvedExpression, true
}

// HasResolvedExpression returns a boolean if a field has been set.
func (o *NegativeTargetingClause) HasResolvedExpression() bool {
	if o != nil && !IsNil(o.ResolvedExpression) {
		return true
	}

	return false
}

// SetResolvedExpression gets a reference to the given []NegativeTargetingExpression and assigns it to the ResolvedExpression field.
func (o *NegativeTargetingClause) SetResolvedExpression(v []NegativeTargetingExpression) {
	o.ResolvedExpression = v
}

func (o NegativeTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.TargetId) {
		toSerialize["targetId"] = o.TargetId
	}
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	if !IsNil(o.ExpressionType) {
		toSerialize["expressionType"] = o.ExpressionType
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !IsNil(o.ResolvedExpression) {
		toSerialize["resolvedExpression"] = o.ResolvedExpression
	}
	return toSerialize, nil
}

type NullableNegativeTargetingClause struct {
	value *NegativeTargetingClause
	isSet bool
}

func (v NullableNegativeTargetingClause) Get() *NegativeTargetingClause {
	return v.value
}

func (v *NullableNegativeTargetingClause) Set(val *NegativeTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableNegativeTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableNegativeTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNegativeTargetingClause(val *NegativeTargetingClause) *NullableNegativeTargetingClause {
	return &NullableNegativeTargetingClause{value: val, isSet: true}
}

func (v NullableNegativeTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableNegativeTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
