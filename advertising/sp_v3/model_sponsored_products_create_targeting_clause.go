package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateTargetingClause{}

// SponsoredProductsCreateTargetingClause struct for SponsoredProductsCreateTargetingClause
type SponsoredProductsCreateTargetingClause struct {
	// The targeting expression.
	Expression []SponsoredProductsCreateTargetingExpressionPredicate `json:"expression"`
	// The identifier of the campaign to which this target is associated.
	CampaignId     string                                     `json:"campaignId"`
	ExpressionType SponsoredProductsCreateExpressionType      `json:"expressionType"`
	State          SponsoredProductsCreateOrUpdateEntityState `json:"state"`
	// The bid for ads sourced using the target. Targets that do not have bid values in listTargetingClauses will inherit the defaultBid from the adGroup level. For more information about bid constraints by marketplace, see [bid limits](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace).
	Bid NullableFloat64 `json:"bid,omitempty"`
	// The identifier of the ad group to which this target is associated.
	AdGroupId string `json:"adGroupId"`
}

type _SponsoredProductsCreateTargetingClause SponsoredProductsCreateTargetingClause

// NewSponsoredProductsCreateTargetingClause instantiates a new SponsoredProductsCreateTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateTargetingClause(expression []SponsoredProductsCreateTargetingExpressionPredicate, campaignId string, expressionType SponsoredProductsCreateExpressionType, state SponsoredProductsCreateOrUpdateEntityState, adGroupId string) *SponsoredProductsCreateTargetingClause {
	this := SponsoredProductsCreateTargetingClause{}
	this.Expression = expression
	this.CampaignId = campaignId
	this.ExpressionType = expressionType
	this.State = state
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsCreateTargetingClauseWithDefaults instantiates a new SponsoredProductsCreateTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateTargetingClauseWithDefaults() *SponsoredProductsCreateTargetingClause {
	this := SponsoredProductsCreateTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value
func (o *SponsoredProductsCreateTargetingClause) GetExpression() []SponsoredProductsCreateTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsCreateTargetingExpressionPredicate
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetingClause) GetExpressionOk() ([]SponsoredProductsCreateTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *SponsoredProductsCreateTargetingClause) SetExpression(v []SponsoredProductsCreateTargetingExpressionPredicate) {
	o.Expression = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateTargetingClause) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetingClause) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateTargetingClause) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetExpressionType returns the ExpressionType field value
func (o *SponsoredProductsCreateTargetingClause) GetExpressionType() SponsoredProductsCreateExpressionType {
	if o == nil {
		var ret SponsoredProductsCreateExpressionType
		return ret
	}

	return o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetingClause) GetExpressionTypeOk() (*SponsoredProductsCreateExpressionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpressionType, true
}

// SetExpressionType sets field value
func (o *SponsoredProductsCreateTargetingClause) SetExpressionType(v SponsoredProductsCreateExpressionType) {
	o.ExpressionType = v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateTargetingClause) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetingClause) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateTargetingClause) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = v
}

// GetBid returns the Bid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SponsoredProductsCreateTargetingClause) GetBid() float64 {
	if o == nil || IsNil(o.Bid.Get()) {
		var ret float64
		return ret
	}
	return *o.Bid.Get()
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SponsoredProductsCreateTargetingClause) GetBidOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bid.Get(), o.Bid.IsSet()
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetingClause) HasBid() bool {
	if o != nil && o.Bid.IsSet() {
		return true
	}

	return false
}

// SetBid gets a reference to the given NullableFloat64 and assigns it to the Bid field.
func (o *SponsoredProductsCreateTargetingClause) SetBid(v float64) {
	o.Bid.Set(&v)
}

// SetBidNil sets the value for Bid to be an explicit nil
func (o *SponsoredProductsCreateTargetingClause) SetBidNil() {
	o.Bid.Set(nil)
}

// UnsetBid ensures that no value is present for Bid, not even an explicit nil
func (o *SponsoredProductsCreateTargetingClause) UnsetBid() {
	o.Bid.Unset()
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsCreateTargetingClause) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetingClause) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsCreateTargetingClause) SetAdGroupId(v string) {
	o.AdGroupId = v
}

func (o SponsoredProductsCreateTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["expressionType"] = o.ExpressionType
	toSerialize["state"] = o.State
	if o.Bid.IsSet() {
		toSerialize["bid"] = o.Bid.Get()
	}
	toSerialize["adGroupId"] = o.AdGroupId
	return toSerialize, nil
}

type NullableSponsoredProductsCreateTargetingClause struct {
	value *SponsoredProductsCreateTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsCreateTargetingClause) Get() *SponsoredProductsCreateTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsCreateTargetingClause) Set(val *SponsoredProductsCreateTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateTargetingClause(val *SponsoredProductsCreateTargetingClause) *NullableSponsoredProductsCreateTargetingClause {
	return &NullableSponsoredProductsCreateTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
