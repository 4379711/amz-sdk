package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsTargetingClause{}

// SponsoredProductsTargetingClause struct for SponsoredProductsTargetingClause
type SponsoredProductsTargetingClause struct {
	// The targeting expression.
	Expression []SponsoredProductsTargetingExpressionPredicate `json:"expression"`
	// The target identifier
	TargetId string `json:"targetId"`
	// The resolved targeting expression.
	ResolvedExpression []SponsoredProductsTargetingExpressionPredicate `json:"resolvedExpression"`
	// The identifier of the campaign to which this target is associated.
	CampaignId     string                          `json:"campaignId"`
	ExpressionType SponsoredProductsExpressionType `json:"expressionType"`
	State          SponsoredProductsEntityState    `json:"state"`
	// The bid for ads sourced using the target. Targets that do not have bid values in listTargetingClauses will inherit the defaultBid from the adGroup level. For more information about bid constraints by marketplace, see [bid limits](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace).
	Bid *float64 `json:"bid,omitempty"`
	// The identifier of the ad group to which this target is associated.
	AdGroupId    string                                        `json:"adGroupId"`
	ExtendedData *SponsoredProductsTargetingClauseExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsTargetingClause SponsoredProductsTargetingClause

// NewSponsoredProductsTargetingClause instantiates a new SponsoredProductsTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsTargetingClause(expression []SponsoredProductsTargetingExpressionPredicate, targetId string, resolvedExpression []SponsoredProductsTargetingExpressionPredicate, campaignId string, expressionType SponsoredProductsExpressionType, state SponsoredProductsEntityState, adGroupId string) *SponsoredProductsTargetingClause {
	this := SponsoredProductsTargetingClause{}
	this.Expression = expression
	this.TargetId = targetId
	this.ResolvedExpression = resolvedExpression
	this.CampaignId = campaignId
	this.ExpressionType = expressionType
	this.State = state
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsTargetingClauseWithDefaults instantiates a new SponsoredProductsTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsTargetingClauseWithDefaults() *SponsoredProductsTargetingClause {
	this := SponsoredProductsTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value
func (o *SponsoredProductsTargetingClause) GetExpression() []SponsoredProductsTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsTargetingExpressionPredicate
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingClause) GetExpressionOk() ([]SponsoredProductsTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *SponsoredProductsTargetingClause) SetExpression(v []SponsoredProductsTargetingExpressionPredicate) {
	o.Expression = v
}

// GetTargetId returns the TargetId field value
func (o *SponsoredProductsTargetingClause) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingClause) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *SponsoredProductsTargetingClause) SetTargetId(v string) {
	o.TargetId = v
}

// GetResolvedExpression returns the ResolvedExpression field value
func (o *SponsoredProductsTargetingClause) GetResolvedExpression() []SponsoredProductsTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsTargetingExpressionPredicate
		return ret
	}

	return o.ResolvedExpression
}

// GetResolvedExpressionOk returns a tuple with the ResolvedExpression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingClause) GetResolvedExpressionOk() ([]SponsoredProductsTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResolvedExpression, true
}

// SetResolvedExpression sets field value
func (o *SponsoredProductsTargetingClause) SetResolvedExpression(v []SponsoredProductsTargetingExpressionPredicate) {
	o.ResolvedExpression = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsTargetingClause) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingClause) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsTargetingClause) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetExpressionType returns the ExpressionType field value
func (o *SponsoredProductsTargetingClause) GetExpressionType() SponsoredProductsExpressionType {
	if o == nil {
		var ret SponsoredProductsExpressionType
		return ret
	}

	return o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingClause) GetExpressionTypeOk() (*SponsoredProductsExpressionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpressionType, true
}

// SetExpressionType sets field value
func (o *SponsoredProductsTargetingClause) SetExpressionType(v SponsoredProductsExpressionType) {
	o.ExpressionType = v
}

// GetState returns the State field value
func (o *SponsoredProductsTargetingClause) GetState() SponsoredProductsEntityState {
	if o == nil {
		var ret SponsoredProductsEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingClause) GetStateOk() (*SponsoredProductsEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsTargetingClause) SetState(v SponsoredProductsEntityState) {
	o.State = v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *SponsoredProductsTargetingClause) GetBid() float64 {
	if o == nil || IsNil(o.Bid) {
		var ret float64
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingClause) GetBidOk() (*float64, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsTargetingClause) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given float64 and assigns it to the Bid field.
func (o *SponsoredProductsTargetingClause) SetBid(v float64) {
	o.Bid = &v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsTargetingClause) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingClause) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsTargetingClause) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsTargetingClause) GetExtendedData() SponsoredProductsTargetingClauseExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsTargetingClauseExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingClause) GetExtendedDataOk() (*SponsoredProductsTargetingClauseExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsTargetingClause) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsTargetingClauseExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsTargetingClause) SetExtendedData(v SponsoredProductsTargetingClauseExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	toSerialize["targetId"] = o.TargetId
	toSerialize["resolvedExpression"] = o.ResolvedExpression
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["expressionType"] = o.ExpressionType
	toSerialize["state"] = o.State
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	toSerialize["adGroupId"] = o.AdGroupId
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsTargetingClause struct {
	value *SponsoredProductsTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsTargetingClause) Get() *SponsoredProductsTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsTargetingClause) Set(val *SponsoredProductsTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetingClause(val *SponsoredProductsTargetingClause) *NullableSponsoredProductsTargetingClause {
	return &NullableSponsoredProductsTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
