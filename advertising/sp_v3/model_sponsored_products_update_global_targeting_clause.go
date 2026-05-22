package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateGlobalTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateGlobalTargetingClause{}

// SponsoredProductsUpdateGlobalTargetingClause struct for SponsoredProductsUpdateGlobalTargetingClause
type SponsoredProductsUpdateGlobalTargetingClause struct {
	// The targeting expression.
	Expression []SponsoredProductsGlobalTargetingExpressionPredicate `json:"expression,omitempty"`
	// The target identifier
	TargetId string `json:"targetId"`
	// Name for the targeting clause
	Name           *string                                           `json:"name,omitempty"`
	ExpressionType *SponsoredProductsExpressionTypeWithoutOther      `json:"expressionType,omitempty"`
	State          *SponsoredProductsCreateOrUpdateGlobalEntityState `json:"state,omitempty"`
	Bid            *SponsoredProductsGlobalBid                       `json:"bid,omitempty"`
}

type _SponsoredProductsUpdateGlobalTargetingClause SponsoredProductsUpdateGlobalTargetingClause

// NewSponsoredProductsUpdateGlobalTargetingClause instantiates a new SponsoredProductsUpdateGlobalTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateGlobalTargetingClause(targetId string) *SponsoredProductsUpdateGlobalTargetingClause {
	this := SponsoredProductsUpdateGlobalTargetingClause{}
	this.TargetId = targetId
	return &this
}

// NewSponsoredProductsUpdateGlobalTargetingClauseWithDefaults instantiates a new SponsoredProductsUpdateGlobalTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateGlobalTargetingClauseWithDefaults() *SponsoredProductsUpdateGlobalTargetingClause {
	this := SponsoredProductsUpdateGlobalTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalTargetingClause) GetExpression() []SponsoredProductsGlobalTargetingExpressionPredicate {
	if o == nil || IsNil(o.Expression) {
		var ret []SponsoredProductsGlobalTargetingExpressionPredicate
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalTargetingClause) GetExpressionOk() ([]SponsoredProductsGlobalTargetingExpressionPredicate, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalTargetingClause) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given []SponsoredProductsGlobalTargetingExpressionPredicate and assigns it to the Expression field.
func (o *SponsoredProductsUpdateGlobalTargetingClause) SetExpression(v []SponsoredProductsGlobalTargetingExpressionPredicate) {
	o.Expression = v
}

// GetTargetId returns the TargetId field value
func (o *SponsoredProductsUpdateGlobalTargetingClause) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalTargetingClause) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *SponsoredProductsUpdateGlobalTargetingClause) SetTargetId(v string) {
	o.TargetId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalTargetingClause) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalTargetingClause) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalTargetingClause) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsUpdateGlobalTargetingClause) SetName(v string) {
	o.Name = &v
}

// GetExpressionType returns the ExpressionType field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalTargetingClause) GetExpressionType() SponsoredProductsExpressionTypeWithoutOther {
	if o == nil || IsNil(o.ExpressionType) {
		var ret SponsoredProductsExpressionTypeWithoutOther
		return ret
	}
	return *o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalTargetingClause) GetExpressionTypeOk() (*SponsoredProductsExpressionTypeWithoutOther, bool) {
	if o == nil || IsNil(o.ExpressionType) {
		return nil, false
	}
	return o.ExpressionType, true
}

// HasExpressionType returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalTargetingClause) HasExpressionType() bool {
	if o != nil && !IsNil(o.ExpressionType) {
		return true
	}

	return false
}

// SetExpressionType gets a reference to the given SponsoredProductsExpressionTypeWithoutOther and assigns it to the ExpressionType field.
func (o *SponsoredProductsUpdateGlobalTargetingClause) SetExpressionType(v SponsoredProductsExpressionTypeWithoutOther) {
	o.ExpressionType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalTargetingClause) GetState() SponsoredProductsCreateOrUpdateGlobalEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsCreateOrUpdateGlobalEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalTargetingClause) GetStateOk() (*SponsoredProductsCreateOrUpdateGlobalEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalTargetingClause) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsCreateOrUpdateGlobalEntityState and assigns it to the State field.
func (o *SponsoredProductsUpdateGlobalTargetingClause) SetState(v SponsoredProductsCreateOrUpdateGlobalEntityState) {
	o.State = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalTargetingClause) GetBid() SponsoredProductsGlobalBid {
	if o == nil || IsNil(o.Bid) {
		var ret SponsoredProductsGlobalBid
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalTargetingClause) GetBidOk() (*SponsoredProductsGlobalBid, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalTargetingClause) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given SponsoredProductsGlobalBid and assigns it to the Bid field.
func (o *SponsoredProductsUpdateGlobalTargetingClause) SetBid(v SponsoredProductsGlobalBid) {
	o.Bid = &v
}

func (o SponsoredProductsUpdateGlobalTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	toSerialize["targetId"] = o.TargetId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ExpressionType) {
		toSerialize["expressionType"] = o.ExpressionType
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateGlobalTargetingClause struct {
	value *SponsoredProductsUpdateGlobalTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsUpdateGlobalTargetingClause) Get() *SponsoredProductsUpdateGlobalTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsUpdateGlobalTargetingClause) Set(val *SponsoredProductsUpdateGlobalTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateGlobalTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateGlobalTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateGlobalTargetingClause(val *SponsoredProductsUpdateGlobalTargetingClause) *NullableSponsoredProductsUpdateGlobalTargetingClause {
	return &NullableSponsoredProductsUpdateGlobalTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateGlobalTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateGlobalTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
