package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateNegativeTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateNegativeTargetingClause{}

// SponsoredProductsCreateNegativeTargetingClause struct for SponsoredProductsCreateNegativeTargetingClause
type SponsoredProductsCreateNegativeTargetingClause struct {
	// The NegativeTargeting expression.
	Expression []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate `json:"expression"`
	// The identifier of the campaign to which this target is associated.
	CampaignId string                                     `json:"campaignId"`
	State      SponsoredProductsCreateOrUpdateEntityState `json:"state"`
	// The identifier of the ad group to which this target is associated.
	AdGroupId string `json:"adGroupId"`
}

type _SponsoredProductsCreateNegativeTargetingClause SponsoredProductsCreateNegativeTargetingClause

// NewSponsoredProductsCreateNegativeTargetingClause instantiates a new SponsoredProductsCreateNegativeTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateNegativeTargetingClause(expression []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate, campaignId string, state SponsoredProductsCreateOrUpdateEntityState, adGroupId string) *SponsoredProductsCreateNegativeTargetingClause {
	this := SponsoredProductsCreateNegativeTargetingClause{}
	this.Expression = expression
	this.CampaignId = campaignId
	this.State = state
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsCreateNegativeTargetingClauseWithDefaults instantiates a new SponsoredProductsCreateNegativeTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateNegativeTargetingClauseWithDefaults() *SponsoredProductsCreateNegativeTargetingClause {
	this := SponsoredProductsCreateNegativeTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value
func (o *SponsoredProductsCreateNegativeTargetingClause) GetExpression() []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateNegativeTargetingClause) GetExpressionOk() ([]SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *SponsoredProductsCreateNegativeTargetingClause) SetExpression(v []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate) {
	o.Expression = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateNegativeTargetingClause) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateNegativeTargetingClause) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateNegativeTargetingClause) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateNegativeTargetingClause) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateNegativeTargetingClause) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateNegativeTargetingClause) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsCreateNegativeTargetingClause) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateNegativeTargetingClause) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsCreateNegativeTargetingClause) SetAdGroupId(v string) {
	o.AdGroupId = v
}

func (o SponsoredProductsCreateNegativeTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["state"] = o.State
	toSerialize["adGroupId"] = o.AdGroupId
	return toSerialize, nil
}

type NullableSponsoredProductsCreateNegativeTargetingClause struct {
	value *SponsoredProductsCreateNegativeTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsCreateNegativeTargetingClause) Get() *SponsoredProductsCreateNegativeTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsCreateNegativeTargetingClause) Set(val *SponsoredProductsCreateNegativeTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateNegativeTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateNegativeTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateNegativeTargetingClause(val *SponsoredProductsCreateNegativeTargetingClause) *NullableSponsoredProductsCreateNegativeTargetingClause {
	return &NullableSponsoredProductsCreateNegativeTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateNegativeTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateNegativeTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
