package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTargetingClause{}

// CreateTargetingClause struct for CreateTargetingClause
type CreateTargetingClause struct {
	State *string `json:"state,omitempty"`
	// The bid will override the adGroup bid if specified. This field is not used for negative targeting clauses. The bid must be less than the maximum allowable bid for the campaign's marketplace; for a list of maximum allowable bids, find the [\"Bid constraints by marketplace\" table in our documentation overview](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace). You cannot manually set a bid when the targeting clause's adGroup has an enabled optimization rule.
	Bid NullableFloat32 `json:"bid,omitempty"`
	// The identifier of the ad group.
	AdGroupId int64 `json:"adGroupId"`
	// Tactic T00020 ad groups only allow manual targeting.
	ExpressionType string `json:"expressionType"`
	// The targeting expression to match against.  ------- Applicable to contextual targeting (T00020) ------- * A 'TargetingExpression' in a contextual targeting campaign can only contain 'TargetingPredicate' components. * Expressions must specify either a category predicate or an ASIN predicate, but never both. * Only one category may be specified per targeting expression. * Only one brand may be specified per targeting expression. * Only one asin may be specified per targeting expression. * To exclude a brand from a targeting expression you must create a negative targeting expression in the same ad group as the positive targeting expression.  ------- Applicable to audiences or contextual targeting (T00030) ------- * A 'TargetingExpression' in a audiences or contextual campaign can contain any target, including 'TargetingPredicate', 'ContentTargetingPredicate', or 'TargetingPredicateNested'.
	Expression []CreateTargetingExpressionInner `json:"expression"`
}

type _CreateTargetingClause CreateTargetingClause

// NewCreateTargetingClause instantiates a new CreateTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTargetingClause(adGroupId int64, expressionType string, expression []CreateTargetingExpressionInner) *CreateTargetingClause {
	this := CreateTargetingClause{}
	this.AdGroupId = adGroupId
	this.ExpressionType = expressionType
	this.Expression = expression
	return &this
}

// NewCreateTargetingClauseWithDefaults instantiates a new CreateTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTargetingClauseWithDefaults() *CreateTargetingClause {
	this := CreateTargetingClause{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CreateTargetingClause) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTargetingClause) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CreateTargetingClause) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CreateTargetingClause) SetState(v string) {
	o.State = &v
}

// GetBid returns the Bid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTargetingClause) GetBid() float32 {
	if o == nil || IsNil(o.Bid.Get()) {
		var ret float32
		return ret
	}
	return *o.Bid.Get()
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTargetingClause) GetBidOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bid.Get(), o.Bid.IsSet()
}

// HasBid returns a boolean if a field has been set.
func (o *CreateTargetingClause) HasBid() bool {
	if o != nil && o.Bid.IsSet() {
		return true
	}

	return false
}

// SetBid gets a reference to the given NullableFloat32 and assigns it to the Bid field.
func (o *CreateTargetingClause) SetBid(v float32) {
	o.Bid.Set(&v)
}

// SetBidNil sets the value for Bid to be an explicit nil
func (o *CreateTargetingClause) SetBidNil() {
	o.Bid.Set(nil)
}

// UnsetBid ensures that no value is present for Bid, not even an explicit nil
func (o *CreateTargetingClause) UnsetBid() {
	o.Bid.Unset()
}

// GetAdGroupId returns the AdGroupId field value
func (o *CreateTargetingClause) GetAdGroupId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *CreateTargetingClause) GetAdGroupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *CreateTargetingClause) SetAdGroupId(v int64) {
	o.AdGroupId = v
}

// GetExpressionType returns the ExpressionType field value
func (o *CreateTargetingClause) GetExpressionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value
// and a boolean to check if the value has been set.
func (o *CreateTargetingClause) GetExpressionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpressionType, true
}

// SetExpressionType sets field value
func (o *CreateTargetingClause) SetExpressionType(v string) {
	o.ExpressionType = v
}

// GetExpression returns the Expression field value
func (o *CreateTargetingClause) GetExpression() []CreateTargetingExpressionInner {
	if o == nil {
		var ret []CreateTargetingExpressionInner
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *CreateTargetingClause) GetExpressionOk() ([]CreateTargetingExpressionInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *CreateTargetingClause) SetExpression(v []CreateTargetingExpressionInner) {
	o.Expression = v
}

func (o CreateTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if o.Bid.IsSet() {
		toSerialize["bid"] = o.Bid.Get()
	}
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["expressionType"] = o.ExpressionType
	toSerialize["expression"] = o.Expression
	return toSerialize, nil
}

type NullableCreateTargetingClause struct {
	value *CreateTargetingClause
	isSet bool
}

func (v NullableCreateTargetingClause) Get() *CreateTargetingClause {
	return v.value
}

func (v *NullableCreateTargetingClause) Set(val *CreateTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTargetingClause(val *CreateTargetingClause) *NullableCreateTargetingClause {
	return &NullableCreateTargetingClause{value: val, isSet: true}
}

func (v NullableCreateTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
