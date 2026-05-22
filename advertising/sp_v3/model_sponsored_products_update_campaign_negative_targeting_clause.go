package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateCampaignNegativeTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateCampaignNegativeTargetingClause{}

// SponsoredProductsUpdateCampaignNegativeTargetingClause struct for SponsoredProductsUpdateCampaignNegativeTargetingClause
type SponsoredProductsUpdateCampaignNegativeTargetingClause struct {
	// The NegativeTargeting expression.
	Expression []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate `json:"expression,omitempty"`
	// The target identifier
	TargetId string                                      `json:"targetId"`
	State    *SponsoredProductsCreateOrUpdateEntityState `json:"state,omitempty"`
}

type _SponsoredProductsUpdateCampaignNegativeTargetingClause SponsoredProductsUpdateCampaignNegativeTargetingClause

// NewSponsoredProductsUpdateCampaignNegativeTargetingClause instantiates a new SponsoredProductsUpdateCampaignNegativeTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateCampaignNegativeTargetingClause(targetId string) *SponsoredProductsUpdateCampaignNegativeTargetingClause {
	this := SponsoredProductsUpdateCampaignNegativeTargetingClause{}
	this.TargetId = targetId
	return &this
}

// NewSponsoredProductsUpdateCampaignNegativeTargetingClauseWithDefaults instantiates a new SponsoredProductsUpdateCampaignNegativeTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateCampaignNegativeTargetingClauseWithDefaults() *SponsoredProductsUpdateCampaignNegativeTargetingClause {
	this := SponsoredProductsUpdateCampaignNegativeTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateCampaignNegativeTargetingClause) GetExpression() []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate {
	if o == nil || IsNil(o.Expression) {
		var ret []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateCampaignNegativeTargetingClause) GetExpressionOk() ([]SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateCampaignNegativeTargetingClause) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate and assigns it to the Expression field.
func (o *SponsoredProductsUpdateCampaignNegativeTargetingClause) SetExpression(v []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate) {
	o.Expression = v
}

// GetTargetId returns the TargetId field value
func (o *SponsoredProductsUpdateCampaignNegativeTargetingClause) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateCampaignNegativeTargetingClause) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *SponsoredProductsUpdateCampaignNegativeTargetingClause) SetTargetId(v string) {
	o.TargetId = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateCampaignNegativeTargetingClause) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateCampaignNegativeTargetingClause) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateCampaignNegativeTargetingClause) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsCreateOrUpdateEntityState and assigns it to the State field.
func (o *SponsoredProductsUpdateCampaignNegativeTargetingClause) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = &v
}

func (o SponsoredProductsUpdateCampaignNegativeTargetingClause) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsUpdateCampaignNegativeTargetingClause struct {
	value *SponsoredProductsUpdateCampaignNegativeTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsUpdateCampaignNegativeTargetingClause) Get() *SponsoredProductsUpdateCampaignNegativeTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsUpdateCampaignNegativeTargetingClause) Set(val *SponsoredProductsUpdateCampaignNegativeTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateCampaignNegativeTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateCampaignNegativeTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateCampaignNegativeTargetingClause(val *SponsoredProductsUpdateCampaignNegativeTargetingClause) *NullableSponsoredProductsUpdateCampaignNegativeTargetingClause {
	return &NullableSponsoredProductsUpdateCampaignNegativeTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateCampaignNegativeTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateCampaignNegativeTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
