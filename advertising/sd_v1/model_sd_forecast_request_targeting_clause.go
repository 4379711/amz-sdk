package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDForecastRequestTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDForecastRequestTargetingClause{}

// SDForecastRequestTargetingClause struct for SDForecastRequestTargetingClause
type SDForecastRequestTargetingClause struct {
	State *string `json:"state,omitempty"`
	// The bid will override the adGroup bid if specified. This field is not used for negative targeting clauses. The bid must be less than the maximum allowable bid for the campaign's marketplace; for a list of maximum allowable bids, find the [\"Bid constraints by marketplace\" table in our documentation overview](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace). You cannot manually set a bid when the targeting clause's adGroup has an enabled optimization rule.
	Bid      NullableFloat32 `json:"bid,omitempty"`
	TargetId *int64          `json:"targetId,omitempty"`
	// The identifier of the ad group.
	AdGroupId *int64 `json:"adGroupId,omitempty"`
	// Tactic T00020 & T00030 ad groups should use 'manual' targeting.
	ExpressionType *string `json:"expressionType,omitempty"`
	// The targeting expression to match against.  ------- Applicable to contextual or content targeting (T00020) ------- * A 'TargetingExpression' in a contextual targeting campaign can contain 'TargetingPredicate' or 'ContentTargetingPredicate' components. * Contextual expressions must specify either a category predicate or an ASIN predicate, but never both. * Only one category may be specified per targeting expression. * Only one brand may be specified per targeting expression. * Only one asin may be specified per targeting expression. * To exclude a brand from a targeting expression you must create a negative targeting expression in the same ad group as the positive targeting expression.  ------- Applicable to audiences or contextual targeting (T00030) ------- * A 'TargetingExpression' in a audiences or contextual campaign can contain any target, including 'TargetingPredicate', 'ContentTargetingPredicate', or 'TargetingPredicateNested'.
	Expression []TargetingExpressionInner `json:"expression,omitempty"`
	// The targeting expression to match against.  ------- Applicable to contextual or content targeting (T00020) ------- * A 'TargetingExpression' in a contextual targeting campaign can contain 'TargetingPredicate' or 'ContentTargetingPredicate' components. * Contextual expressions must specify either a category predicate or an ASIN predicate, but never both. * Only one category may be specified per targeting expression. * Only one brand may be specified per targeting expression. * Only one asin may be specified per targeting expression. * To exclude a brand from a targeting expression you must create a negative targeting expression in the same ad group as the positive targeting expression.  ------- Applicable to audiences or contextual targeting (T00030) ------- * A 'TargetingExpression' in a audiences or contextual campaign can contain any target, including 'TargetingPredicate', 'ContentTargetingPredicate', or 'TargetingPredicateNested'.
	ResolvedExpression []TargetingExpressionInner `json:"resolvedExpression,omitempty"`
}

// NewSDForecastRequestTargetingClause instantiates a new SDForecastRequestTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDForecastRequestTargetingClause() *SDForecastRequestTargetingClause {
	this := SDForecastRequestTargetingClause{}
	return &this
}

// NewSDForecastRequestTargetingClauseWithDefaults instantiates a new SDForecastRequestTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDForecastRequestTargetingClauseWithDefaults() *SDForecastRequestTargetingClause {
	this := SDForecastRequestTargetingClause{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SDForecastRequestTargetingClause) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDForecastRequestTargetingClause) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SDForecastRequestTargetingClause) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *SDForecastRequestTargetingClause) SetState(v string) {
	o.State = &v
}

// GetBid returns the Bid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SDForecastRequestTargetingClause) GetBid() float32 {
	if o == nil || IsNil(o.Bid.Get()) {
		var ret float32
		return ret
	}
	return *o.Bid.Get()
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SDForecastRequestTargetingClause) GetBidOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bid.Get(), o.Bid.IsSet()
}

// HasBid returns a boolean if a field has been set.
func (o *SDForecastRequestTargetingClause) HasBid() bool {
	if o != nil && o.Bid.IsSet() {
		return true
	}

	return false
}

// SetBid gets a reference to the given NullableFloat32 and assigns it to the Bid field.
func (o *SDForecastRequestTargetingClause) SetBid(v float32) {
	o.Bid.Set(&v)
}

// SetBidNil sets the value for Bid to be an explicit nil
func (o *SDForecastRequestTargetingClause) SetBidNil() {
	o.Bid.Set(nil)
}

// UnsetBid ensures that no value is present for Bid, not even an explicit nil
func (o *SDForecastRequestTargetingClause) UnsetBid() {
	o.Bid.Unset()
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *SDForecastRequestTargetingClause) GetTargetId() int64 {
	if o == nil || IsNil(o.TargetId) {
		var ret int64
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDForecastRequestTargetingClause) GetTargetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *SDForecastRequestTargetingClause) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given int64 and assigns it to the TargetId field.
func (o *SDForecastRequestTargetingClause) SetTargetId(v int64) {
	o.TargetId = &v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *SDForecastRequestTargetingClause) GetAdGroupId() int64 {
	if o == nil || IsNil(o.AdGroupId) {
		var ret int64
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDForecastRequestTargetingClause) GetAdGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *SDForecastRequestTargetingClause) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given int64 and assigns it to the AdGroupId field.
func (o *SDForecastRequestTargetingClause) SetAdGroupId(v int64) {
	o.AdGroupId = &v
}

// GetExpressionType returns the ExpressionType field value if set, zero value otherwise.
func (o *SDForecastRequestTargetingClause) GetExpressionType() string {
	if o == nil || IsNil(o.ExpressionType) {
		var ret string
		return ret
	}
	return *o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDForecastRequestTargetingClause) GetExpressionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpressionType) {
		return nil, false
	}
	return o.ExpressionType, true
}

// HasExpressionType returns a boolean if a field has been set.
func (o *SDForecastRequestTargetingClause) HasExpressionType() bool {
	if o != nil && !IsNil(o.ExpressionType) {
		return true
	}

	return false
}

// SetExpressionType gets a reference to the given string and assigns it to the ExpressionType field.
func (o *SDForecastRequestTargetingClause) SetExpressionType(v string) {
	o.ExpressionType = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *SDForecastRequestTargetingClause) GetExpression() []TargetingExpressionInner {
	if o == nil || IsNil(o.Expression) {
		var ret []TargetingExpressionInner
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDForecastRequestTargetingClause) GetExpressionOk() ([]TargetingExpressionInner, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *SDForecastRequestTargetingClause) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given []TargetingExpressionInner and assigns it to the Expression field.
func (o *SDForecastRequestTargetingClause) SetExpression(v []TargetingExpressionInner) {
	o.Expression = v
}

// GetResolvedExpression returns the ResolvedExpression field value if set, zero value otherwise.
func (o *SDForecastRequestTargetingClause) GetResolvedExpression() []TargetingExpressionInner {
	if o == nil || IsNil(o.ResolvedExpression) {
		var ret []TargetingExpressionInner
		return ret
	}
	return o.ResolvedExpression
}

// GetResolvedExpressionOk returns a tuple with the ResolvedExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDForecastRequestTargetingClause) GetResolvedExpressionOk() ([]TargetingExpressionInner, bool) {
	if o == nil || IsNil(o.ResolvedExpression) {
		return nil, false
	}
	return o.ResolvedExpression, true
}

// HasResolvedExpression returns a boolean if a field has been set.
func (o *SDForecastRequestTargetingClause) HasResolvedExpression() bool {
	if o != nil && !IsNil(o.ResolvedExpression) {
		return true
	}

	return false
}

// SetResolvedExpression gets a reference to the given []TargetingExpressionInner and assigns it to the ResolvedExpression field.
func (o *SDForecastRequestTargetingClause) SetResolvedExpression(v []TargetingExpressionInner) {
	o.ResolvedExpression = v
}

func (o SDForecastRequestTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if o.Bid.IsSet() {
		toSerialize["bid"] = o.Bid.Get()
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

type NullableSDForecastRequestTargetingClause struct {
	value *SDForecastRequestTargetingClause
	isSet bool
}

func (v NullableSDForecastRequestTargetingClause) Get() *SDForecastRequestTargetingClause {
	return v.value
}

func (v *NullableSDForecastRequestTargetingClause) Set(val *SDForecastRequestTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSDForecastRequestTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSDForecastRequestTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDForecastRequestTargetingClause(val *SDForecastRequestTargetingClause) *NullableSDForecastRequestTargetingClause {
	return &NullableSDForecastRequestTargetingClause{value: val, isSet: true}
}

func (v NullableSDForecastRequestTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDForecastRequestTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
