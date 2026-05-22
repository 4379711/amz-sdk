package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateGlobalNegativeTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateGlobalNegativeTargetingClause{}

// SponsoredProductsUpdateGlobalNegativeTargetingClause struct for SponsoredProductsUpdateGlobalNegativeTargetingClause
type SponsoredProductsUpdateGlobalNegativeTargetingClause struct {
	// The NegativeTargeting expression.
	Expression []SponsoredProductsGlobalTargetingExpressionPredicate `json:"expression,omitempty"`
	// The target identifier
	TargetId string `json:"targetId"`
	// Name for the negative targeting clause
	Name  *string                                           `json:"name,omitempty"`
	State *SponsoredProductsCreateOrUpdateGlobalEntityState `json:"state,omitempty"`
}

type _SponsoredProductsUpdateGlobalNegativeTargetingClause SponsoredProductsUpdateGlobalNegativeTargetingClause

// NewSponsoredProductsUpdateGlobalNegativeTargetingClause instantiates a new SponsoredProductsUpdateGlobalNegativeTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateGlobalNegativeTargetingClause(targetId string) *SponsoredProductsUpdateGlobalNegativeTargetingClause {
	this := SponsoredProductsUpdateGlobalNegativeTargetingClause{}
	this.TargetId = targetId
	return &this
}

// NewSponsoredProductsUpdateGlobalNegativeTargetingClauseWithDefaults instantiates a new SponsoredProductsUpdateGlobalNegativeTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateGlobalNegativeTargetingClauseWithDefaults() *SponsoredProductsUpdateGlobalNegativeTargetingClause {
	this := SponsoredProductsUpdateGlobalNegativeTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalNegativeTargetingClause) GetExpression() []SponsoredProductsGlobalTargetingExpressionPredicate {
	if o == nil || IsNil(o.Expression) {
		var ret []SponsoredProductsGlobalTargetingExpressionPredicate
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalNegativeTargetingClause) GetExpressionOk() ([]SponsoredProductsGlobalTargetingExpressionPredicate, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalNegativeTargetingClause) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given []SponsoredProductsGlobalTargetingExpressionPredicate and assigns it to the Expression field.
func (o *SponsoredProductsUpdateGlobalNegativeTargetingClause) SetExpression(v []SponsoredProductsGlobalTargetingExpressionPredicate) {
	o.Expression = v
}

// GetTargetId returns the TargetId field value
func (o *SponsoredProductsUpdateGlobalNegativeTargetingClause) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalNegativeTargetingClause) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *SponsoredProductsUpdateGlobalNegativeTargetingClause) SetTargetId(v string) {
	o.TargetId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalNegativeTargetingClause) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalNegativeTargetingClause) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalNegativeTargetingClause) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsUpdateGlobalNegativeTargetingClause) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalNegativeTargetingClause) GetState() SponsoredProductsCreateOrUpdateGlobalEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsCreateOrUpdateGlobalEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalNegativeTargetingClause) GetStateOk() (*SponsoredProductsCreateOrUpdateGlobalEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalNegativeTargetingClause) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsCreateOrUpdateGlobalEntityState and assigns it to the State field.
func (o *SponsoredProductsUpdateGlobalNegativeTargetingClause) SetState(v SponsoredProductsCreateOrUpdateGlobalEntityState) {
	o.State = &v
}

func (o SponsoredProductsUpdateGlobalNegativeTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	toSerialize["targetId"] = o.TargetId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateGlobalNegativeTargetingClause struct {
	value *SponsoredProductsUpdateGlobalNegativeTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsUpdateGlobalNegativeTargetingClause) Get() *SponsoredProductsUpdateGlobalNegativeTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsUpdateGlobalNegativeTargetingClause) Set(val *SponsoredProductsUpdateGlobalNegativeTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateGlobalNegativeTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateGlobalNegativeTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateGlobalNegativeTargetingClause(val *SponsoredProductsUpdateGlobalNegativeTargetingClause) *NullableSponsoredProductsUpdateGlobalNegativeTargetingClause {
	return &NullableSponsoredProductsUpdateGlobalNegativeTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateGlobalNegativeTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateGlobalNegativeTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
