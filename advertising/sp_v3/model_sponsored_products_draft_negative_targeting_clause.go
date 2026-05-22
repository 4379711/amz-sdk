package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftNegativeTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftNegativeTargetingClause{}

// SponsoredProductsDraftNegativeTargetingClause struct for SponsoredProductsDraftNegativeTargetingClause
type SponsoredProductsDraftNegativeTargetingClause struct {
	// The DraftNegativeTargeting expression.
	Expression []SponsoredProductsNegativeTargetingExpressionPredicate `json:"expression"`
	// The target identifier
	TargetId string `json:"targetId"`
	// The resolved DraftNegativeTargeting expression.
	ResolvedExpression []SponsoredProductsNegativeTargetingExpressionPredicate `json:"resolvedExpression"`
	// The identifier of the campaign to which this target is associated.
	CampaignId string `json:"campaignId"`
	// The identifier of the ad group to which this target is associated.
	AdGroupId    string                                                     `json:"adGroupId"`
	ExtendedData *SponsoredProductsDraftNegativeTargetingClauseExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsDraftNegativeTargetingClause SponsoredProductsDraftNegativeTargetingClause

// NewSponsoredProductsDraftNegativeTargetingClause instantiates a new SponsoredProductsDraftNegativeTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftNegativeTargetingClause(expression []SponsoredProductsNegativeTargetingExpressionPredicate, targetId string, resolvedExpression []SponsoredProductsNegativeTargetingExpressionPredicate, campaignId string, adGroupId string) *SponsoredProductsDraftNegativeTargetingClause {
	this := SponsoredProductsDraftNegativeTargetingClause{}
	this.Expression = expression
	this.TargetId = targetId
	this.ResolvedExpression = resolvedExpression
	this.CampaignId = campaignId
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsDraftNegativeTargetingClauseWithDefaults instantiates a new SponsoredProductsDraftNegativeTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftNegativeTargetingClauseWithDefaults() *SponsoredProductsDraftNegativeTargetingClause {
	this := SponsoredProductsDraftNegativeTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value
func (o *SponsoredProductsDraftNegativeTargetingClause) GetExpression() []SponsoredProductsNegativeTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsNegativeTargetingExpressionPredicate
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetingClause) GetExpressionOk() ([]SponsoredProductsNegativeTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *SponsoredProductsDraftNegativeTargetingClause) SetExpression(v []SponsoredProductsNegativeTargetingExpressionPredicate) {
	o.Expression = v
}

// GetTargetId returns the TargetId field value
func (o *SponsoredProductsDraftNegativeTargetingClause) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetingClause) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *SponsoredProductsDraftNegativeTargetingClause) SetTargetId(v string) {
	o.TargetId = v
}

// GetResolvedExpression returns the ResolvedExpression field value
func (o *SponsoredProductsDraftNegativeTargetingClause) GetResolvedExpression() []SponsoredProductsNegativeTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsNegativeTargetingExpressionPredicate
		return ret
	}

	return o.ResolvedExpression
}

// GetResolvedExpressionOk returns a tuple with the ResolvedExpression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetingClause) GetResolvedExpressionOk() ([]SponsoredProductsNegativeTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResolvedExpression, true
}

// SetResolvedExpression sets field value
func (o *SponsoredProductsDraftNegativeTargetingClause) SetResolvedExpression(v []SponsoredProductsNegativeTargetingExpressionPredicate) {
	o.ResolvedExpression = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsDraftNegativeTargetingClause) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetingClause) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsDraftNegativeTargetingClause) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsDraftNegativeTargetingClause) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetingClause) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsDraftNegativeTargetingClause) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsDraftNegativeTargetingClause) GetExtendedData() SponsoredProductsDraftNegativeTargetingClauseExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsDraftNegativeTargetingClauseExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetingClause) GetExtendedDataOk() (*SponsoredProductsDraftNegativeTargetingClauseExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsDraftNegativeTargetingClause) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsDraftNegativeTargetingClauseExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsDraftNegativeTargetingClause) SetExtendedData(v SponsoredProductsDraftNegativeTargetingClauseExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsDraftNegativeTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	toSerialize["targetId"] = o.TargetId
	toSerialize["resolvedExpression"] = o.ResolvedExpression
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["adGroupId"] = o.AdGroupId
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftNegativeTargetingClause struct {
	value *SponsoredProductsDraftNegativeTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsDraftNegativeTargetingClause) Get() *SponsoredProductsDraftNegativeTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsDraftNegativeTargetingClause) Set(val *SponsoredProductsDraftNegativeTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftNegativeTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftNegativeTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftNegativeTargetingClause(val *SponsoredProductsDraftNegativeTargetingClause) *NullableSponsoredProductsDraftNegativeTargetingClause {
	return &NullableSponsoredProductsDraftNegativeTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftNegativeTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftNegativeTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
