package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalTargetingClause{}

// SponsoredProductsGlobalTargetingClause struct for SponsoredProductsGlobalTargetingClause
type SponsoredProductsGlobalTargetingClause struct {
	// The targeting expression.
	Expression []SponsoredProductsGlobalTargetingExpressionPredicate `json:"expression"`
	// The target identifier
	TargetId string `json:"targetId"`
	// The resolved targeting expression.
	ResolvedExpression []SponsoredProductsGlobalTargetingExpressionPredicate `json:"resolvedExpression"`
	// The identifier of the campaign to which this target is associated.
	CampaignId string `json:"campaignId"`
	// Name for the targeting clause
	Name           *string                            `json:"name,omitempty"`
	ExpressionType SponsoredProductsExpressionType    `json:"expressionType"`
	State          SponsoredProductsGlobalEntityState `json:"state"`
	Bid            SponsoredProductsGlobalBid         `json:"bid"`
	// The identifier of the ad group to which this target is associated.
	AdGroupId    string                                              `json:"adGroupId"`
	ExtendedData *SponsoredProductsGlobalTargetingClauseExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsGlobalTargetingClause SponsoredProductsGlobalTargetingClause

// NewSponsoredProductsGlobalTargetingClause instantiates a new SponsoredProductsGlobalTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalTargetingClause(expression []SponsoredProductsGlobalTargetingExpressionPredicate, targetId string, resolvedExpression []SponsoredProductsGlobalTargetingExpressionPredicate, campaignId string, expressionType SponsoredProductsExpressionType, state SponsoredProductsGlobalEntityState, bid SponsoredProductsGlobalBid, adGroupId string) *SponsoredProductsGlobalTargetingClause {
	this := SponsoredProductsGlobalTargetingClause{}
	this.Expression = expression
	this.TargetId = targetId
	this.ResolvedExpression = resolvedExpression
	this.CampaignId = campaignId
	this.ExpressionType = expressionType
	this.State = state
	this.Bid = bid
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsGlobalTargetingClauseWithDefaults instantiates a new SponsoredProductsGlobalTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalTargetingClauseWithDefaults() *SponsoredProductsGlobalTargetingClause {
	this := SponsoredProductsGlobalTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value
func (o *SponsoredProductsGlobalTargetingClause) GetExpression() []SponsoredProductsGlobalTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsGlobalTargetingExpressionPredicate
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClause) GetExpressionOk() ([]SponsoredProductsGlobalTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *SponsoredProductsGlobalTargetingClause) SetExpression(v []SponsoredProductsGlobalTargetingExpressionPredicate) {
	o.Expression = v
}

// GetTargetId returns the TargetId field value
func (o *SponsoredProductsGlobalTargetingClause) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClause) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *SponsoredProductsGlobalTargetingClause) SetTargetId(v string) {
	o.TargetId = v
}

// GetResolvedExpression returns the ResolvedExpression field value
func (o *SponsoredProductsGlobalTargetingClause) GetResolvedExpression() []SponsoredProductsGlobalTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsGlobalTargetingExpressionPredicate
		return ret
	}

	return o.ResolvedExpression
}

// GetResolvedExpressionOk returns a tuple with the ResolvedExpression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClause) GetResolvedExpressionOk() ([]SponsoredProductsGlobalTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResolvedExpression, true
}

// SetResolvedExpression sets field value
func (o *SponsoredProductsGlobalTargetingClause) SetResolvedExpression(v []SponsoredProductsGlobalTargetingExpressionPredicate) {
	o.ResolvedExpression = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsGlobalTargetingClause) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClause) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsGlobalTargetingClause) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalTargetingClause) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClause) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalTargetingClause) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsGlobalTargetingClause) SetName(v string) {
	o.Name = &v
}

// GetExpressionType returns the ExpressionType field value
func (o *SponsoredProductsGlobalTargetingClause) GetExpressionType() SponsoredProductsExpressionType {
	if o == nil {
		var ret SponsoredProductsExpressionType
		return ret
	}

	return o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClause) GetExpressionTypeOk() (*SponsoredProductsExpressionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpressionType, true
}

// SetExpressionType sets field value
func (o *SponsoredProductsGlobalTargetingClause) SetExpressionType(v SponsoredProductsExpressionType) {
	o.ExpressionType = v
}

// GetState returns the State field value
func (o *SponsoredProductsGlobalTargetingClause) GetState() SponsoredProductsGlobalEntityState {
	if o == nil {
		var ret SponsoredProductsGlobalEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClause) GetStateOk() (*SponsoredProductsGlobalEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsGlobalTargetingClause) SetState(v SponsoredProductsGlobalEntityState) {
	o.State = v
}

// GetBid returns the Bid field value
func (o *SponsoredProductsGlobalTargetingClause) GetBid() SponsoredProductsGlobalBid {
	if o == nil {
		var ret SponsoredProductsGlobalBid
		return ret
	}

	return o.Bid
}

// GetBidOk returns a tuple with the Bid field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClause) GetBidOk() (*SponsoredProductsGlobalBid, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bid, true
}

// SetBid sets field value
func (o *SponsoredProductsGlobalTargetingClause) SetBid(v SponsoredProductsGlobalBid) {
	o.Bid = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsGlobalTargetingClause) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClause) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsGlobalTargetingClause) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalTargetingClause) GetExtendedData() SponsoredProductsGlobalTargetingClauseExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsGlobalTargetingClauseExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClause) GetExtendedDataOk() (*SponsoredProductsGlobalTargetingClauseExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalTargetingClause) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsGlobalTargetingClauseExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsGlobalTargetingClause) SetExtendedData(v SponsoredProductsGlobalTargetingClauseExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsGlobalTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	toSerialize["targetId"] = o.TargetId
	toSerialize["resolvedExpression"] = o.ResolvedExpression
	toSerialize["campaignId"] = o.CampaignId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["expressionType"] = o.ExpressionType
	toSerialize["state"] = o.State
	toSerialize["bid"] = o.Bid
	toSerialize["adGroupId"] = o.AdGroupId
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalTargetingClause struct {
	value *SponsoredProductsGlobalTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsGlobalTargetingClause) Get() *SponsoredProductsGlobalTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsGlobalTargetingClause) Set(val *SponsoredProductsGlobalTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalTargetingClause(val *SponsoredProductsGlobalTargetingClause) *NullableSponsoredProductsGlobalTargetingClause {
	return &NullableSponsoredProductsGlobalTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
