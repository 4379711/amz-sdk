package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalNegativeTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalNegativeTargetingClause{}

// SponsoredProductsGlobalNegativeTargetingClause struct for SponsoredProductsGlobalNegativeTargetingClause
type SponsoredProductsGlobalNegativeTargetingClause struct {
	// The NegativeTargeting expression.
	Expression []SponsoredProductsGlobalTargetingExpressionPredicate `json:"expression"`
	// The target identifier
	TargetId string `json:"targetId"`
	// The resolved NegativeTargeting expression.
	ResolvedExpression []SponsoredProductsGlobalTargetingExpressionPredicate `json:"resolvedExpression"`
	// The identifier of the campaign to which this target is associated.
	CampaignId string `json:"campaignId"`
	// Name for the negative targeting clause
	Name  string                             `json:"name"`
	State SponsoredProductsGlobalEntityState `json:"state"`
	// The identifier of the ad group to which this target is associated.
	AdGroupId    string                                                      `json:"adGroupId"`
	ExtendedData *SponsoredProductsGlobalNegativeTargetingClauseExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsGlobalNegativeTargetingClause SponsoredProductsGlobalNegativeTargetingClause

// NewSponsoredProductsGlobalNegativeTargetingClause instantiates a new SponsoredProductsGlobalNegativeTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalNegativeTargetingClause(expression []SponsoredProductsGlobalTargetingExpressionPredicate, targetId string, resolvedExpression []SponsoredProductsGlobalTargetingExpressionPredicate, campaignId string, name string, state SponsoredProductsGlobalEntityState, adGroupId string) *SponsoredProductsGlobalNegativeTargetingClause {
	this := SponsoredProductsGlobalNegativeTargetingClause{}
	this.Expression = expression
	this.TargetId = targetId
	this.ResolvedExpression = resolvedExpression
	this.CampaignId = campaignId
	this.Name = name
	this.State = state
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsGlobalNegativeTargetingClauseWithDefaults instantiates a new SponsoredProductsGlobalNegativeTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalNegativeTargetingClauseWithDefaults() *SponsoredProductsGlobalNegativeTargetingClause {
	this := SponsoredProductsGlobalNegativeTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value
func (o *SponsoredProductsGlobalNegativeTargetingClause) GetExpression() []SponsoredProductsGlobalTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsGlobalTargetingExpressionPredicate
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClause) GetExpressionOk() ([]SponsoredProductsGlobalTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *SponsoredProductsGlobalNegativeTargetingClause) SetExpression(v []SponsoredProductsGlobalTargetingExpressionPredicate) {
	o.Expression = v
}

// GetTargetId returns the TargetId field value
func (o *SponsoredProductsGlobalNegativeTargetingClause) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClause) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *SponsoredProductsGlobalNegativeTargetingClause) SetTargetId(v string) {
	o.TargetId = v
}

// GetResolvedExpression returns the ResolvedExpression field value
func (o *SponsoredProductsGlobalNegativeTargetingClause) GetResolvedExpression() []SponsoredProductsGlobalTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsGlobalTargetingExpressionPredicate
		return ret
	}

	return o.ResolvedExpression
}

// GetResolvedExpressionOk returns a tuple with the ResolvedExpression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClause) GetResolvedExpressionOk() ([]SponsoredProductsGlobalTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResolvedExpression, true
}

// SetResolvedExpression sets field value
func (o *SponsoredProductsGlobalNegativeTargetingClause) SetResolvedExpression(v []SponsoredProductsGlobalTargetingExpressionPredicate) {
	o.ResolvedExpression = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsGlobalNegativeTargetingClause) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClause) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsGlobalNegativeTargetingClause) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetName returns the Name field value
func (o *SponsoredProductsGlobalNegativeTargetingClause) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClause) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SponsoredProductsGlobalNegativeTargetingClause) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *SponsoredProductsGlobalNegativeTargetingClause) GetState() SponsoredProductsGlobalEntityState {
	if o == nil {
		var ret SponsoredProductsGlobalEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClause) GetStateOk() (*SponsoredProductsGlobalEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsGlobalNegativeTargetingClause) SetState(v SponsoredProductsGlobalEntityState) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsGlobalNegativeTargetingClause) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClause) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsGlobalNegativeTargetingClause) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalNegativeTargetingClause) GetExtendedData() SponsoredProductsGlobalNegativeTargetingClauseExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsGlobalNegativeTargetingClauseExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClause) GetExtendedDataOk() (*SponsoredProductsGlobalNegativeTargetingClauseExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClause) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsGlobalNegativeTargetingClauseExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsGlobalNegativeTargetingClause) SetExtendedData(v SponsoredProductsGlobalNegativeTargetingClauseExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsGlobalNegativeTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	toSerialize["targetId"] = o.TargetId
	toSerialize["resolvedExpression"] = o.ResolvedExpression
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	toSerialize["adGroupId"] = o.AdGroupId
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalNegativeTargetingClause struct {
	value *SponsoredProductsGlobalNegativeTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsGlobalNegativeTargetingClause) Get() *SponsoredProductsGlobalNegativeTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsGlobalNegativeTargetingClause) Set(val *SponsoredProductsGlobalNegativeTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalNegativeTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalNegativeTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalNegativeTargetingClause(val *SponsoredProductsGlobalNegativeTargetingClause) *NullableSponsoredProductsGlobalNegativeTargetingClause {
	return &NullableSponsoredProductsGlobalNegativeTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalNegativeTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalNegativeTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
