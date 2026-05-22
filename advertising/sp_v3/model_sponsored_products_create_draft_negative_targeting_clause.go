package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateDraftNegativeTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateDraftNegativeTargetingClause{}

// SponsoredProductsCreateDraftNegativeTargetingClause struct for SponsoredProductsCreateDraftNegativeTargetingClause
type SponsoredProductsCreateDraftNegativeTargetingClause struct {
	// The DraftNegativeTargeting expression.
	Expression []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate `json:"expression"`
	// The identifier of the campaign to which this target is associated.
	CampaignId string `json:"campaignId"`
	// The identifier of the ad group to which this target is associated.
	AdGroupId string `json:"adGroupId"`
}

type _SponsoredProductsCreateDraftNegativeTargetingClause SponsoredProductsCreateDraftNegativeTargetingClause

// NewSponsoredProductsCreateDraftNegativeTargetingClause instantiates a new SponsoredProductsCreateDraftNegativeTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateDraftNegativeTargetingClause(expression []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate, campaignId string, adGroupId string) *SponsoredProductsCreateDraftNegativeTargetingClause {
	this := SponsoredProductsCreateDraftNegativeTargetingClause{}
	this.Expression = expression
	this.CampaignId = campaignId
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsCreateDraftNegativeTargetingClauseWithDefaults instantiates a new SponsoredProductsCreateDraftNegativeTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateDraftNegativeTargetingClauseWithDefaults() *SponsoredProductsCreateDraftNegativeTargetingClause {
	this := SponsoredProductsCreateDraftNegativeTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value
func (o *SponsoredProductsCreateDraftNegativeTargetingClause) GetExpression() []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftNegativeTargetingClause) GetExpressionOk() ([]SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *SponsoredProductsCreateDraftNegativeTargetingClause) SetExpression(v []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicate) {
	o.Expression = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateDraftNegativeTargetingClause) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftNegativeTargetingClause) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateDraftNegativeTargetingClause) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsCreateDraftNegativeTargetingClause) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftNegativeTargetingClause) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsCreateDraftNegativeTargetingClause) SetAdGroupId(v string) {
	o.AdGroupId = v
}

func (o SponsoredProductsCreateDraftNegativeTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["adGroupId"] = o.AdGroupId
	return toSerialize, nil
}

type NullableSponsoredProductsCreateDraftNegativeTargetingClause struct {
	value *SponsoredProductsCreateDraftNegativeTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsCreateDraftNegativeTargetingClause) Get() *SponsoredProductsCreateDraftNegativeTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsCreateDraftNegativeTargetingClause) Set(val *SponsoredProductsCreateDraftNegativeTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateDraftNegativeTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateDraftNegativeTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateDraftNegativeTargetingClause(val *SponsoredProductsCreateDraftNegativeTargetingClause) *NullableSponsoredProductsCreateDraftNegativeTargetingClause {
	return &NullableSponsoredProductsCreateDraftNegativeTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateDraftNegativeTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateDraftNegativeTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
