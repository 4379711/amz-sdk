package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateDraftTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateDraftTargetingClause{}

// SponsoredProductsCreateDraftTargetingClause struct for SponsoredProductsCreateDraftTargetingClause
type SponsoredProductsCreateDraftTargetingClause struct {
	// The targeting expression.
	Expression []SponsoredProductsTargetingExpressionPredicateWithoutOther `json:"expression"`
	// The identifier of the campaign to which this target is associated.
	CampaignId     string                                      `json:"campaignId"`
	ExpressionType SponsoredProductsExpressionTypeWithoutOther `json:"expressionType"`
	State          *SponsoredProductsEntityState               `json:"state,omitempty"`
	// The bid for ads sourced using the target. Targets that do not have bid values in listDraftTargetingClauses will inherit the defaultBid from the adGroup level. For more information about bid constraints by marketplace, see [bid limits](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace).
	Bid *float64 `json:"bid,omitempty"`
	// The identifier of the ad group to which this target is associated.
	AdGroupId string `json:"adGroupId"`
}

type _SponsoredProductsCreateDraftTargetingClause SponsoredProductsCreateDraftTargetingClause

// NewSponsoredProductsCreateDraftTargetingClause instantiates a new SponsoredProductsCreateDraftTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateDraftTargetingClause(expression []SponsoredProductsTargetingExpressionPredicateWithoutOther, campaignId string, expressionType SponsoredProductsExpressionTypeWithoutOther, adGroupId string) *SponsoredProductsCreateDraftTargetingClause {
	this := SponsoredProductsCreateDraftTargetingClause{}
	this.Expression = expression
	this.CampaignId = campaignId
	this.ExpressionType = expressionType
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsCreateDraftTargetingClauseWithDefaults instantiates a new SponsoredProductsCreateDraftTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateDraftTargetingClauseWithDefaults() *SponsoredProductsCreateDraftTargetingClause {
	this := SponsoredProductsCreateDraftTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value
func (o *SponsoredProductsCreateDraftTargetingClause) GetExpression() []SponsoredProductsTargetingExpressionPredicateWithoutOther {
	if o == nil {
		var ret []SponsoredProductsTargetingExpressionPredicateWithoutOther
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftTargetingClause) GetExpressionOk() ([]SponsoredProductsTargetingExpressionPredicateWithoutOther, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *SponsoredProductsCreateDraftTargetingClause) SetExpression(v []SponsoredProductsTargetingExpressionPredicateWithoutOther) {
	o.Expression = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateDraftTargetingClause) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftTargetingClause) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateDraftTargetingClause) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetExpressionType returns the ExpressionType field value
func (o *SponsoredProductsCreateDraftTargetingClause) GetExpressionType() SponsoredProductsExpressionTypeWithoutOther {
	if o == nil {
		var ret SponsoredProductsExpressionTypeWithoutOther
		return ret
	}

	return o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftTargetingClause) GetExpressionTypeOk() (*SponsoredProductsExpressionTypeWithoutOther, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpressionType, true
}

// SetExpressionType sets field value
func (o *SponsoredProductsCreateDraftTargetingClause) SetExpressionType(v SponsoredProductsExpressionTypeWithoutOther) {
	o.ExpressionType = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsCreateDraftTargetingClause) GetState() SponsoredProductsEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftTargetingClause) GetStateOk() (*SponsoredProductsEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftTargetingClause) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsEntityState and assigns it to the State field.
func (o *SponsoredProductsCreateDraftTargetingClause) SetState(v SponsoredProductsEntityState) {
	o.State = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *SponsoredProductsCreateDraftTargetingClause) GetBid() float64 {
	if o == nil || IsNil(o.Bid) {
		var ret float64
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftTargetingClause) GetBidOk() (*float64, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftTargetingClause) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given float64 and assigns it to the Bid field.
func (o *SponsoredProductsCreateDraftTargetingClause) SetBid(v float64) {
	o.Bid = &v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsCreateDraftTargetingClause) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftTargetingClause) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsCreateDraftTargetingClause) SetAdGroupId(v string) {
	o.AdGroupId = v
}

func (o SponsoredProductsCreateDraftTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["expressionType"] = o.ExpressionType
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	toSerialize["adGroupId"] = o.AdGroupId
	return toSerialize, nil
}

type NullableSponsoredProductsCreateDraftTargetingClause struct {
	value *SponsoredProductsCreateDraftTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsCreateDraftTargetingClause) Get() *SponsoredProductsCreateDraftTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsCreateDraftTargetingClause) Set(val *SponsoredProductsCreateDraftTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateDraftTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateDraftTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateDraftTargetingClause(val *SponsoredProductsCreateDraftTargetingClause) *NullableSponsoredProductsCreateDraftTargetingClause {
	return &NullableSponsoredProductsCreateDraftTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateDraftTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateDraftTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
