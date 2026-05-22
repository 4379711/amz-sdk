package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateTargetingClause{}

// SponsoredProductsUpdateTargetingClause struct for SponsoredProductsUpdateTargetingClause
type SponsoredProductsUpdateTargetingClause struct {
	// The targeting expression.
	Expression []SponsoredProductsTargetingExpressionPredicateWithoutOther `json:"expression,omitempty"`
	// The target identifier
	TargetId       string                                       `json:"targetId"`
	ExpressionType *SponsoredProductsExpressionTypeWithoutOther `json:"expressionType,omitempty"`
	State          *SponsoredProductsCreateOrUpdateEntityState  `json:"state,omitempty"`
	// The bid for ads sourced using the target. Targets that do not have bid values in listTargetingClauses will inherit the defaultBid from the adGroup level. For more information about bid constraints by marketplace, see [bid limits](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace).
	Bid NullableFloat64 `json:"bid,omitempty"`
}

type _SponsoredProductsUpdateTargetingClause SponsoredProductsUpdateTargetingClause

// NewSponsoredProductsUpdateTargetingClause instantiates a new SponsoredProductsUpdateTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateTargetingClause(targetId string) *SponsoredProductsUpdateTargetingClause {
	this := SponsoredProductsUpdateTargetingClause{}
	this.TargetId = targetId
	return &this
}

// NewSponsoredProductsUpdateTargetingClauseWithDefaults instantiates a new SponsoredProductsUpdateTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateTargetingClauseWithDefaults() *SponsoredProductsUpdateTargetingClause {
	this := SponsoredProductsUpdateTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateTargetingClause) GetExpression() []SponsoredProductsTargetingExpressionPredicateWithoutOther {
	if o == nil || IsNil(o.Expression) {
		var ret []SponsoredProductsTargetingExpressionPredicateWithoutOther
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateTargetingClause) GetExpressionOk() ([]SponsoredProductsTargetingExpressionPredicateWithoutOther, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateTargetingClause) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given []SponsoredProductsTargetingExpressionPredicateWithoutOther and assigns it to the Expression field.
func (o *SponsoredProductsUpdateTargetingClause) SetExpression(v []SponsoredProductsTargetingExpressionPredicateWithoutOther) {
	o.Expression = v
}

// GetTargetId returns the TargetId field value
func (o *SponsoredProductsUpdateTargetingClause) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateTargetingClause) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *SponsoredProductsUpdateTargetingClause) SetTargetId(v string) {
	o.TargetId = v
}

// GetExpressionType returns the ExpressionType field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateTargetingClause) GetExpressionType() SponsoredProductsExpressionTypeWithoutOther {
	if o == nil || IsNil(o.ExpressionType) {
		var ret SponsoredProductsExpressionTypeWithoutOther
		return ret
	}
	return *o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateTargetingClause) GetExpressionTypeOk() (*SponsoredProductsExpressionTypeWithoutOther, bool) {
	if o == nil || IsNil(o.ExpressionType) {
		return nil, false
	}
	return o.ExpressionType, true
}

// HasExpressionType returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateTargetingClause) HasExpressionType() bool {
	if o != nil && !IsNil(o.ExpressionType) {
		return true
	}

	return false
}

// SetExpressionType gets a reference to the given SponsoredProductsExpressionTypeWithoutOther and assigns it to the ExpressionType field.
func (o *SponsoredProductsUpdateTargetingClause) SetExpressionType(v SponsoredProductsExpressionTypeWithoutOther) {
	o.ExpressionType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateTargetingClause) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateTargetingClause) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateTargetingClause) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsCreateOrUpdateEntityState and assigns it to the State field.
func (o *SponsoredProductsUpdateTargetingClause) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = &v
}

// GetBid returns the Bid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SponsoredProductsUpdateTargetingClause) GetBid() float64 {
	if o == nil || IsNil(o.Bid.Get()) {
		var ret float64
		return ret
	}
	return *o.Bid.Get()
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SponsoredProductsUpdateTargetingClause) GetBidOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bid.Get(), o.Bid.IsSet()
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateTargetingClause) HasBid() bool {
	if o != nil && o.Bid.IsSet() {
		return true
	}

	return false
}

// SetBid gets a reference to the given NullableFloat64 and assigns it to the Bid field.
func (o *SponsoredProductsUpdateTargetingClause) SetBid(v float64) {
	o.Bid.Set(&v)
}

// SetBidNil sets the value for Bid to be an explicit nil
func (o *SponsoredProductsUpdateTargetingClause) SetBidNil() {
	o.Bid.Set(nil)
}

// UnsetBid ensures that no value is present for Bid, not even an explicit nil
func (o *SponsoredProductsUpdateTargetingClause) UnsetBid() {
	o.Bid.Unset()
}

func (o SponsoredProductsUpdateTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	toSerialize["targetId"] = o.TargetId
	if !IsNil(o.ExpressionType) {
		toSerialize["expressionType"] = o.ExpressionType
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if o.Bid.IsSet() {
		toSerialize["bid"] = o.Bid.Get()
	}
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateTargetingClause struct {
	value *SponsoredProductsUpdateTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsUpdateTargetingClause) Get() *SponsoredProductsUpdateTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsUpdateTargetingClause) Set(val *SponsoredProductsUpdateTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateTargetingClause(val *SponsoredProductsUpdateTargetingClause) *NullableSponsoredProductsUpdateTargetingClause {
	return &NullableSponsoredProductsUpdateTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
