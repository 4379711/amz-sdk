package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateNegativeTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateNegativeTargetingClause{}

// SponsoredProductsUpdateNegativeTargetingClause struct for SponsoredProductsUpdateNegativeTargetingClause
type SponsoredProductsUpdateNegativeTargetingClause struct {
	// The NegativeTargeting expression.
	Expression []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate `json:"expression,omitempty"`
	// The target identifier
	TargetId string                                      `json:"targetId"`
	State    *SponsoredProductsCreateOrUpdateEntityState `json:"state,omitempty"`
}

type _SponsoredProductsUpdateNegativeTargetingClause SponsoredProductsUpdateNegativeTargetingClause

// NewSponsoredProductsUpdateNegativeTargetingClause instantiates a new SponsoredProductsUpdateNegativeTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateNegativeTargetingClause(targetId string) *SponsoredProductsUpdateNegativeTargetingClause {
	this := SponsoredProductsUpdateNegativeTargetingClause{}
	this.TargetId = targetId
	return &this
}

// NewSponsoredProductsUpdateNegativeTargetingClauseWithDefaults instantiates a new SponsoredProductsUpdateNegativeTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateNegativeTargetingClauseWithDefaults() *SponsoredProductsUpdateNegativeTargetingClause {
	this := SponsoredProductsUpdateNegativeTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateNegativeTargetingClause) GetExpression() []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate {
	if o == nil || IsNil(o.Expression) {
		var ret []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateNegativeTargetingClause) GetExpressionOk() ([]SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateNegativeTargetingClause) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate and assigns it to the Expression field.
func (o *SponsoredProductsUpdateNegativeTargetingClause) SetExpression(v []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate) {
	o.Expression = v
}

// GetTargetId returns the TargetId field value
func (o *SponsoredProductsUpdateNegativeTargetingClause) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateNegativeTargetingClause) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *SponsoredProductsUpdateNegativeTargetingClause) SetTargetId(v string) {
	o.TargetId = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateNegativeTargetingClause) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateNegativeTargetingClause) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateNegativeTargetingClause) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsCreateOrUpdateEntityState and assigns it to the State field.
func (o *SponsoredProductsUpdateNegativeTargetingClause) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = &v
}

func (o SponsoredProductsUpdateNegativeTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	toSerialize["targetId"] = o.TargetId
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateNegativeTargetingClause struct {
	value *SponsoredProductsUpdateNegativeTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsUpdateNegativeTargetingClause) Get() *SponsoredProductsUpdateNegativeTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsUpdateNegativeTargetingClause) Set(val *SponsoredProductsUpdateNegativeTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateNegativeTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateNegativeTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateNegativeTargetingClause(val *SponsoredProductsUpdateNegativeTargetingClause) *NullableSponsoredProductsUpdateNegativeTargetingClause {
	return &NullableSponsoredProductsUpdateNegativeTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateNegativeTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateNegativeTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
