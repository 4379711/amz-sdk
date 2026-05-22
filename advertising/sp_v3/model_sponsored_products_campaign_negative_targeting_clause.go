package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignNegativeTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignNegativeTargetingClause{}

// SponsoredProductsCampaignNegativeTargetingClause struct for SponsoredProductsCampaignNegativeTargetingClause
type SponsoredProductsCampaignNegativeTargetingClause struct {
	// The CampaignNegativeTargetingClause expression.
	Expression []SponsoredProductsNegativeTargetingExpressionPredicate `json:"expression"`
	// The target identifier
	TargetId string `json:"targetId"`
	// The resolved CampaignNegativeTargetingClause expression.
	ResolvedExpression []SponsoredProductsNegativeTargetingExpressionPredicate `json:"resolvedExpression"`
	// The identifier of the campaign to which this target is associated.
	CampaignId   string                                                        `json:"campaignId"`
	State        SponsoredProductsEntityState                                  `json:"state"`
	ExtendedData *SponsoredProductsCampaignNegativeTargetingClauseExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsCampaignNegativeTargetingClause SponsoredProductsCampaignNegativeTargetingClause

// NewSponsoredProductsCampaignNegativeTargetingClause instantiates a new SponsoredProductsCampaignNegativeTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignNegativeTargetingClause(expression []SponsoredProductsNegativeTargetingExpressionPredicate, targetId string, resolvedExpression []SponsoredProductsNegativeTargetingExpressionPredicate, campaignId string, state SponsoredProductsEntityState) *SponsoredProductsCampaignNegativeTargetingClause {
	this := SponsoredProductsCampaignNegativeTargetingClause{}
	this.Expression = expression
	this.TargetId = targetId
	this.ResolvedExpression = resolvedExpression
	this.CampaignId = campaignId
	this.State = state
	return &this
}

// NewSponsoredProductsCampaignNegativeTargetingClauseWithDefaults instantiates a new SponsoredProductsCampaignNegativeTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignNegativeTargetingClauseWithDefaults() *SponsoredProductsCampaignNegativeTargetingClause {
	this := SponsoredProductsCampaignNegativeTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value
func (o *SponsoredProductsCampaignNegativeTargetingClause) GetExpression() []SponsoredProductsNegativeTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsNegativeTargetingExpressionPredicate
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetingClause) GetExpressionOk() ([]SponsoredProductsNegativeTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *SponsoredProductsCampaignNegativeTargetingClause) SetExpression(v []SponsoredProductsNegativeTargetingExpressionPredicate) {
	o.Expression = v
}

// GetTargetId returns the TargetId field value
func (o *SponsoredProductsCampaignNegativeTargetingClause) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetingClause) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *SponsoredProductsCampaignNegativeTargetingClause) SetTargetId(v string) {
	o.TargetId = v
}

// GetResolvedExpression returns the ResolvedExpression field value
func (o *SponsoredProductsCampaignNegativeTargetingClause) GetResolvedExpression() []SponsoredProductsNegativeTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsNegativeTargetingExpressionPredicate
		return ret
	}

	return o.ResolvedExpression
}

// GetResolvedExpressionOk returns a tuple with the ResolvedExpression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetingClause) GetResolvedExpressionOk() ([]SponsoredProductsNegativeTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResolvedExpression, true
}

// SetResolvedExpression sets field value
func (o *SponsoredProductsCampaignNegativeTargetingClause) SetResolvedExpression(v []SponsoredProductsNegativeTargetingExpressionPredicate) {
	o.ResolvedExpression = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCampaignNegativeTargetingClause) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetingClause) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCampaignNegativeTargetingClause) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetState returns the State field value
func (o *SponsoredProductsCampaignNegativeTargetingClause) GetState() SponsoredProductsEntityState {
	if o == nil {
		var ret SponsoredProductsEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetingClause) GetStateOk() (*SponsoredProductsEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCampaignNegativeTargetingClause) SetState(v SponsoredProductsEntityState) {
	o.State = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignNegativeTargetingClause) GetExtendedData() SponsoredProductsCampaignNegativeTargetingClauseExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsCampaignNegativeTargetingClauseExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetingClause) GetExtendedDataOk() (*SponsoredProductsCampaignNegativeTargetingClauseExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignNegativeTargetingClause) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsCampaignNegativeTargetingClauseExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsCampaignNegativeTargetingClause) SetExtendedData(v SponsoredProductsCampaignNegativeTargetingClauseExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsCampaignNegativeTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	toSerialize["targetId"] = o.TargetId
	toSerialize["resolvedExpression"] = o.ResolvedExpression
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["state"] = o.State
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignNegativeTargetingClause struct {
	value *SponsoredProductsCampaignNegativeTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsCampaignNegativeTargetingClause) Get() *SponsoredProductsCampaignNegativeTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsCampaignNegativeTargetingClause) Set(val *SponsoredProductsCampaignNegativeTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignNegativeTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignNegativeTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignNegativeTargetingClause(val *SponsoredProductsCampaignNegativeTargetingClause) *NullableSponsoredProductsCampaignNegativeTargetingClause {
	return &NullableSponsoredProductsCampaignNegativeTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignNegativeTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignNegativeTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
