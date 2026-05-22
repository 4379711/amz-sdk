package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateGlobalTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateGlobalTargetingClause{}

// SponsoredProductsCreateGlobalTargetingClause struct for SponsoredProductsCreateGlobalTargetingClause
type SponsoredProductsCreateGlobalTargetingClause struct {
	// The targeting expression.
	Expression []SponsoredProductsGlobalTargetingExpressionPredicate `json:"expression"`
	// The identifier of the campaign to which this target is associated.
	CampaignId string `json:"campaignId"`
	// Name for the targeting clause
	Name           *string                                          `json:"name,omitempty"`
	ExpressionType SponsoredProductsCreateExpressionType            `json:"expressionType"`
	State          SponsoredProductsCreateOrUpdateGlobalEntityState `json:"state"`
	Bid            *SponsoredProductsGlobalBid                      `json:"bid,omitempty"`
	// The identifier of the ad group to which this target is associated.
	AdGroupId string `json:"adGroupId"`
}

type _SponsoredProductsCreateGlobalTargetingClause SponsoredProductsCreateGlobalTargetingClause

// NewSponsoredProductsCreateGlobalTargetingClause instantiates a new SponsoredProductsCreateGlobalTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateGlobalTargetingClause(expression []SponsoredProductsGlobalTargetingExpressionPredicate, campaignId string, expressionType SponsoredProductsCreateExpressionType, state SponsoredProductsCreateOrUpdateGlobalEntityState, adGroupId string) *SponsoredProductsCreateGlobalTargetingClause {
	this := SponsoredProductsCreateGlobalTargetingClause{}
	this.Expression = expression
	this.CampaignId = campaignId
	this.ExpressionType = expressionType
	this.State = state
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsCreateGlobalTargetingClauseWithDefaults instantiates a new SponsoredProductsCreateGlobalTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateGlobalTargetingClauseWithDefaults() *SponsoredProductsCreateGlobalTargetingClause {
	this := SponsoredProductsCreateGlobalTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value
func (o *SponsoredProductsCreateGlobalTargetingClause) GetExpression() []SponsoredProductsGlobalTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsGlobalTargetingExpressionPredicate
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalTargetingClause) GetExpressionOk() ([]SponsoredProductsGlobalTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *SponsoredProductsCreateGlobalTargetingClause) SetExpression(v []SponsoredProductsGlobalTargetingExpressionPredicate) {
	o.Expression = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateGlobalTargetingClause) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalTargetingClause) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateGlobalTargetingClause) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsCreateGlobalTargetingClause) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalTargetingClause) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsCreateGlobalTargetingClause) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsCreateGlobalTargetingClause) SetName(v string) {
	o.Name = &v
}

// GetExpressionType returns the ExpressionType field value
func (o *SponsoredProductsCreateGlobalTargetingClause) GetExpressionType() SponsoredProductsCreateExpressionType {
	if o == nil {
		var ret SponsoredProductsCreateExpressionType
		return ret
	}

	return o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalTargetingClause) GetExpressionTypeOk() (*SponsoredProductsCreateExpressionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpressionType, true
}

// SetExpressionType sets field value
func (o *SponsoredProductsCreateGlobalTargetingClause) SetExpressionType(v SponsoredProductsCreateExpressionType) {
	o.ExpressionType = v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateGlobalTargetingClause) GetState() SponsoredProductsCreateOrUpdateGlobalEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateGlobalEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalTargetingClause) GetStateOk() (*SponsoredProductsCreateOrUpdateGlobalEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateGlobalTargetingClause) SetState(v SponsoredProductsCreateOrUpdateGlobalEntityState) {
	o.State = v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *SponsoredProductsCreateGlobalTargetingClause) GetBid() SponsoredProductsGlobalBid {
	if o == nil || IsNil(o.Bid) {
		var ret SponsoredProductsGlobalBid
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalTargetingClause) GetBidOk() (*SponsoredProductsGlobalBid, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsCreateGlobalTargetingClause) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given SponsoredProductsGlobalBid and assigns it to the Bid field.
func (o *SponsoredProductsCreateGlobalTargetingClause) SetBid(v SponsoredProductsGlobalBid) {
	o.Bid = &v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsCreateGlobalTargetingClause) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalTargetingClause) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsCreateGlobalTargetingClause) SetAdGroupId(v string) {
	o.AdGroupId = v
}

func (o SponsoredProductsCreateGlobalTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	toSerialize["campaignId"] = o.CampaignId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["expressionType"] = o.ExpressionType
	toSerialize["state"] = o.State
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	toSerialize["adGroupId"] = o.AdGroupId
	return toSerialize, nil
}

type NullableSponsoredProductsCreateGlobalTargetingClause struct {
	value *SponsoredProductsCreateGlobalTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsCreateGlobalTargetingClause) Get() *SponsoredProductsCreateGlobalTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsCreateGlobalTargetingClause) Set(val *SponsoredProductsCreateGlobalTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateGlobalTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateGlobalTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateGlobalTargetingClause(val *SponsoredProductsCreateGlobalTargetingClause) *NullableSponsoredProductsCreateGlobalTargetingClause {
	return &NullableSponsoredProductsCreateGlobalTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateGlobalTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateGlobalTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
