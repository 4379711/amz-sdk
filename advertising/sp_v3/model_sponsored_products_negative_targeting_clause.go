package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNegativeTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNegativeTargetingClause{}

// SponsoredProductsNegativeTargetingClause struct for SponsoredProductsNegativeTargetingClause
type SponsoredProductsNegativeTargetingClause struct {
	// The NegativeTargeting expression.
	Expression []SponsoredProductsNegativeTargetingExpressionPredicate `json:"expression"`
	// The target identifier
	TargetId string `json:"targetId"`
	// The resolved NegativeTargeting expression.
	ResolvedExpression []SponsoredProductsNegativeTargetingExpressionPredicate `json:"resolvedExpression"`
	// The identifier of the campaign to which this target is associated.
	CampaignId string                       `json:"campaignId"`
	State      SponsoredProductsEntityState `json:"state"`
	// The identifier of the ad group to which this target is associated.
	AdGroupId    string                                                `json:"adGroupId"`
	ExtendedData *SponsoredProductsNegativeTargetingClauseExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsNegativeTargetingClause SponsoredProductsNegativeTargetingClause

// NewSponsoredProductsNegativeTargetingClause instantiates a new SponsoredProductsNegativeTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNegativeTargetingClause(expression []SponsoredProductsNegativeTargetingExpressionPredicate, targetId string, resolvedExpression []SponsoredProductsNegativeTargetingExpressionPredicate, campaignId string, state SponsoredProductsEntityState, adGroupId string) *SponsoredProductsNegativeTargetingClause {
	this := SponsoredProductsNegativeTargetingClause{}
	this.Expression = expression
	this.TargetId = targetId
	this.ResolvedExpression = resolvedExpression
	this.CampaignId = campaignId
	this.State = state
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsNegativeTargetingClauseWithDefaults instantiates a new SponsoredProductsNegativeTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNegativeTargetingClauseWithDefaults() *SponsoredProductsNegativeTargetingClause {
	this := SponsoredProductsNegativeTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value
func (o *SponsoredProductsNegativeTargetingClause) GetExpression() []SponsoredProductsNegativeTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsNegativeTargetingExpressionPredicate
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetingClause) GetExpressionOk() ([]SponsoredProductsNegativeTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *SponsoredProductsNegativeTargetingClause) SetExpression(v []SponsoredProductsNegativeTargetingExpressionPredicate) {
	o.Expression = v
}

// GetTargetId returns the TargetId field value
func (o *SponsoredProductsNegativeTargetingClause) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetingClause) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *SponsoredProductsNegativeTargetingClause) SetTargetId(v string) {
	o.TargetId = v
}

// GetResolvedExpression returns the ResolvedExpression field value
func (o *SponsoredProductsNegativeTargetingClause) GetResolvedExpression() []SponsoredProductsNegativeTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsNegativeTargetingExpressionPredicate
		return ret
	}

	return o.ResolvedExpression
}

// GetResolvedExpressionOk returns a tuple with the ResolvedExpression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetingClause) GetResolvedExpressionOk() ([]SponsoredProductsNegativeTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResolvedExpression, true
}

// SetResolvedExpression sets field value
func (o *SponsoredProductsNegativeTargetingClause) SetResolvedExpression(v []SponsoredProductsNegativeTargetingExpressionPredicate) {
	o.ResolvedExpression = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsNegativeTargetingClause) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetingClause) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsNegativeTargetingClause) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetState returns the State field value
func (o *SponsoredProductsNegativeTargetingClause) GetState() SponsoredProductsEntityState {
	if o == nil {
		var ret SponsoredProductsEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetingClause) GetStateOk() (*SponsoredProductsEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsNegativeTargetingClause) SetState(v SponsoredProductsEntityState) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsNegativeTargetingClause) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetingClause) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsNegativeTargetingClause) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeTargetingClause) GetExtendedData() SponsoredProductsNegativeTargetingClauseExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsNegativeTargetingClauseExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetingClause) GetExtendedDataOk() (*SponsoredProductsNegativeTargetingClauseExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeTargetingClause) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsNegativeTargetingClauseExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsNegativeTargetingClause) SetExtendedData(v SponsoredProductsNegativeTargetingClauseExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsNegativeTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	toSerialize["targetId"] = o.TargetId
	toSerialize["resolvedExpression"] = o.ResolvedExpression
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["state"] = o.State
	toSerialize["adGroupId"] = o.AdGroupId
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsNegativeTargetingClause struct {
	value *SponsoredProductsNegativeTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsNegativeTargetingClause) Get() *SponsoredProductsNegativeTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsNegativeTargetingClause) Set(val *SponsoredProductsNegativeTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNegativeTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNegativeTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNegativeTargetingClause(val *SponsoredProductsNegativeTargetingClause) *NullableSponsoredProductsNegativeTargetingClause {
	return &NullableSponsoredProductsNegativeTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsNegativeTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNegativeTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
