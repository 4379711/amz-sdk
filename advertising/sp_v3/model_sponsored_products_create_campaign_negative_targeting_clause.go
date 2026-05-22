package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsCreateCampaignNegativeTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateCampaignNegativeTargetingClause{}

// SponsoredProductsCreateCampaignNegativeTargetingClause struct for SponsoredProductsCreateCampaignNegativeTargetingClause
type SponsoredProductsCreateCampaignNegativeTargetingClause struct {
	// The NegativeTargeting expression.
	Expression []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate `json:"expression"`
	// The identifier of the campaign to which this target is associated. CampaignNegativeTargetingClauses are only available for AUTO campaigns
	CampaignId string                                     `json:"campaignId"`
	State      SponsoredProductsCreateOrUpdateEntityState `json:"state"`
}

type _SponsoredProductsCreateCampaignNegativeTargetingClause SponsoredProductsCreateCampaignNegativeTargetingClause

// NewSponsoredProductsCreateCampaignNegativeTargetingClause instantiates a new SponsoredProductsCreateCampaignNegativeTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateCampaignNegativeTargetingClause(expression []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate, campaignId string, state SponsoredProductsCreateOrUpdateEntityState) *SponsoredProductsCreateCampaignNegativeTargetingClause {
	this := SponsoredProductsCreateCampaignNegativeTargetingClause{}
	this.Expression = expression
	this.CampaignId = campaignId
	this.State = state
	return &this
}

// NewSponsoredProductsCreateCampaignNegativeTargetingClauseWithDefaults instantiates a new SponsoredProductsCreateCampaignNegativeTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateCampaignNegativeTargetingClauseWithDefaults() *SponsoredProductsCreateCampaignNegativeTargetingClause {
	this := SponsoredProductsCreateCampaignNegativeTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value
func (o *SponsoredProductsCreateCampaignNegativeTargetingClause) GetExpression() []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateCampaignNegativeTargetingClause) GetExpressionOk() ([]SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *SponsoredProductsCreateCampaignNegativeTargetingClause) SetExpression(v []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate) {
	o.Expression = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateCampaignNegativeTargetingClause) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateCampaignNegativeTargetingClause) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateCampaignNegativeTargetingClause) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateCampaignNegativeTargetingClause) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateCampaignNegativeTargetingClause) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateCampaignNegativeTargetingClause) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = v
}

func (o SponsoredProductsCreateCampaignNegativeTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["state"] = o.State
	return toSerialize, nil
}

type NullableSponsoredProductsCreateCampaignNegativeTargetingClause struct {
	value *SponsoredProductsCreateCampaignNegativeTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsCreateCampaignNegativeTargetingClause) Get() *SponsoredProductsCreateCampaignNegativeTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsCreateCampaignNegativeTargetingClause) Set(val *SponsoredProductsCreateCampaignNegativeTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateCampaignNegativeTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateCampaignNegativeTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateCampaignNegativeTargetingClause(val *SponsoredProductsCreateCampaignNegativeTargetingClause) *NullableSponsoredProductsCreateCampaignNegativeTargetingClause {
	return &NullableSponsoredProductsCreateCampaignNegativeTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateCampaignNegativeTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateCampaignNegativeTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
