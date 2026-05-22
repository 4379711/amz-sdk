package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftTargetingClause{}

// SponsoredProductsDraftTargetingClause struct for SponsoredProductsDraftTargetingClause
type SponsoredProductsDraftTargetingClause struct {
	// The targeting expression.
	Expression []SponsoredProductsTargetingExpressionPredicate `json:"expression"`
	// The target identifier
	TargetId string `json:"targetId"`
	// The resolved targeting expression.
	ResolvedExpression []SponsoredProductsTargetingExpressionPredicate `json:"resolvedExpression"`
	// The identifier of the campaign to which this target is associated.
	CampaignId     string                          `json:"campaignId"`
	ExpressionType SponsoredProductsExpressionType `json:"expressionType"`
	State          *SponsoredProductsEntityState   `json:"state,omitempty"`
	// The bid for ads sourced using the target. Targets that do not have bid values in listTargetingClauses will inherit the defaultBid from the adGroup level. For more information about bid constraints by marketplace, see [bid limits](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace).
	Bid NullableFloat64 `json:"bid,omitempty"`
	// The identifier of the ad group to which this target is associated.
	AdGroupId    string                                             `json:"adGroupId"`
	ExtendedData *SponsoredProductsDraftTargetingClauseExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsDraftTargetingClause SponsoredProductsDraftTargetingClause

// NewSponsoredProductsDraftTargetingClause instantiates a new SponsoredProductsDraftTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftTargetingClause(expression []SponsoredProductsTargetingExpressionPredicate, targetId string, resolvedExpression []SponsoredProductsTargetingExpressionPredicate, campaignId string, expressionType SponsoredProductsExpressionType, adGroupId string) *SponsoredProductsDraftTargetingClause {
	this := SponsoredProductsDraftTargetingClause{}
	this.Expression = expression
	this.TargetId = targetId
	this.ResolvedExpression = resolvedExpression
	this.CampaignId = campaignId
	this.ExpressionType = expressionType
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsDraftTargetingClauseWithDefaults instantiates a new SponsoredProductsDraftTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftTargetingClauseWithDefaults() *SponsoredProductsDraftTargetingClause {
	this := SponsoredProductsDraftTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value
func (o *SponsoredProductsDraftTargetingClause) GetExpression() []SponsoredProductsTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsTargetingExpressionPredicate
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetingClause) GetExpressionOk() ([]SponsoredProductsTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *SponsoredProductsDraftTargetingClause) SetExpression(v []SponsoredProductsTargetingExpressionPredicate) {
	o.Expression = v
}

// GetTargetId returns the TargetId field value
func (o *SponsoredProductsDraftTargetingClause) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetingClause) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *SponsoredProductsDraftTargetingClause) SetTargetId(v string) {
	o.TargetId = v
}

// GetResolvedExpression returns the ResolvedExpression field value
func (o *SponsoredProductsDraftTargetingClause) GetResolvedExpression() []SponsoredProductsTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsTargetingExpressionPredicate
		return ret
	}

	return o.ResolvedExpression
}

// GetResolvedExpressionOk returns a tuple with the ResolvedExpression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetingClause) GetResolvedExpressionOk() ([]SponsoredProductsTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResolvedExpression, true
}

// SetResolvedExpression sets field value
func (o *SponsoredProductsDraftTargetingClause) SetResolvedExpression(v []SponsoredProductsTargetingExpressionPredicate) {
	o.ResolvedExpression = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsDraftTargetingClause) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetingClause) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsDraftTargetingClause) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetExpressionType returns the ExpressionType field value
func (o *SponsoredProductsDraftTargetingClause) GetExpressionType() SponsoredProductsExpressionType {
	if o == nil {
		var ret SponsoredProductsExpressionType
		return ret
	}

	return o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetingClause) GetExpressionTypeOk() (*SponsoredProductsExpressionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpressionType, true
}

// SetExpressionType sets field value
func (o *SponsoredProductsDraftTargetingClause) SetExpressionType(v SponsoredProductsExpressionType) {
	o.ExpressionType = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetingClause) GetState() SponsoredProductsEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetingClause) GetStateOk() (*SponsoredProductsEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetingClause) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsEntityState and assigns it to the State field.
func (o *SponsoredProductsDraftTargetingClause) SetState(v SponsoredProductsEntityState) {
	o.State = &v
}

// GetBid returns the Bid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SponsoredProductsDraftTargetingClause) GetBid() float64 {
	if o == nil || IsNil(o.Bid.Get()) {
		var ret float64
		return ret
	}
	return *o.Bid.Get()
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SponsoredProductsDraftTargetingClause) GetBidOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bid.Get(), o.Bid.IsSet()
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetingClause) HasBid() bool {
	if o != nil && o.Bid.IsSet() {
		return true
	}

	return false
}

// SetBid gets a reference to the given NullableFloat64 and assigns it to the Bid field.
func (o *SponsoredProductsDraftTargetingClause) SetBid(v float64) {
	o.Bid.Set(&v)
}

// SetBidNil sets the value for Bid to be an explicit nil
func (o *SponsoredProductsDraftTargetingClause) SetBidNil() {
	o.Bid.Set(nil)
}

// UnsetBid ensures that no value is present for Bid, not even an explicit nil
func (o *SponsoredProductsDraftTargetingClause) UnsetBid() {
	o.Bid.Unset()
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsDraftTargetingClause) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetingClause) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsDraftTargetingClause) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetingClause) GetExtendedData() SponsoredProductsDraftTargetingClauseExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsDraftTargetingClauseExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetingClause) GetExtendedDataOk() (*SponsoredProductsDraftTargetingClauseExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetingClause) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsDraftTargetingClauseExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsDraftTargetingClause) SetExtendedData(v SponsoredProductsDraftTargetingClauseExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsDraftTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	toSerialize["targetId"] = o.TargetId
	toSerialize["resolvedExpression"] = o.ResolvedExpression
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["expressionType"] = o.ExpressionType
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if o.Bid.IsSet() {
		toSerialize["bid"] = o.Bid.Get()
	}
	toSerialize["adGroupId"] = o.AdGroupId
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftTargetingClause struct {
	value *SponsoredProductsDraftTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsDraftTargetingClause) Get() *SponsoredProductsDraftTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsDraftTargetingClause) Set(val *SponsoredProductsDraftTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftTargetingClause(val *SponsoredProductsDraftTargetingClause) *NullableSponsoredProductsDraftTargetingClause {
	return &NullableSponsoredProductsDraftTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
