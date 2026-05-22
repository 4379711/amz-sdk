package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsCreateGlobalNegativeTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateGlobalNegativeTargetingClause{}

// SponsoredProductsCreateGlobalNegativeTargetingClause struct for SponsoredProductsCreateGlobalNegativeTargetingClause
type SponsoredProductsCreateGlobalNegativeTargetingClause struct {
	// The NegativeTargeting expression.
	Expression []SponsoredProductsGlobalTargetingExpressionPredicate `json:"expression"`
	// The identifier of the campaign to which this target is associated.
	CampaignId string `json:"campaignId"`
	// Name for the negative targeting clause
	Name  string                                           `json:"name"`
	State SponsoredProductsCreateOrUpdateGlobalEntityState `json:"state"`
	// The identifier of the ad group to which this target is associated.
	AdGroupId string `json:"adGroupId"`
}

type _SponsoredProductsCreateGlobalNegativeTargetingClause SponsoredProductsCreateGlobalNegativeTargetingClause

// NewSponsoredProductsCreateGlobalNegativeTargetingClause instantiates a new SponsoredProductsCreateGlobalNegativeTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateGlobalNegativeTargetingClause(expression []SponsoredProductsGlobalTargetingExpressionPredicate, campaignId string, name string, state SponsoredProductsCreateOrUpdateGlobalEntityState, adGroupId string) *SponsoredProductsCreateGlobalNegativeTargetingClause {
	this := SponsoredProductsCreateGlobalNegativeTargetingClause{}
	this.Expression = expression
	this.CampaignId = campaignId
	this.Name = name
	this.State = state
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsCreateGlobalNegativeTargetingClauseWithDefaults instantiates a new SponsoredProductsCreateGlobalNegativeTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateGlobalNegativeTargetingClauseWithDefaults() *SponsoredProductsCreateGlobalNegativeTargetingClause {
	this := SponsoredProductsCreateGlobalNegativeTargetingClause{}
	return &this
}

// GetExpression returns the Expression field value
func (o *SponsoredProductsCreateGlobalNegativeTargetingClause) GetExpression() []SponsoredProductsGlobalTargetingExpressionPredicate {
	if o == nil {
		var ret []SponsoredProductsGlobalTargetingExpressionPredicate
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalNegativeTargetingClause) GetExpressionOk() ([]SponsoredProductsGlobalTargetingExpressionPredicate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *SponsoredProductsCreateGlobalNegativeTargetingClause) SetExpression(v []SponsoredProductsGlobalTargetingExpressionPredicate) {
	o.Expression = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateGlobalNegativeTargetingClause) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalNegativeTargetingClause) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateGlobalNegativeTargetingClause) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetName returns the Name field value
func (o *SponsoredProductsCreateGlobalNegativeTargetingClause) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalNegativeTargetingClause) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SponsoredProductsCreateGlobalNegativeTargetingClause) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateGlobalNegativeTargetingClause) GetState() SponsoredProductsCreateOrUpdateGlobalEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateGlobalEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalNegativeTargetingClause) GetStateOk() (*SponsoredProductsCreateOrUpdateGlobalEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateGlobalNegativeTargetingClause) SetState(v SponsoredProductsCreateOrUpdateGlobalEntityState) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsCreateGlobalNegativeTargetingClause) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalNegativeTargetingClause) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsCreateGlobalNegativeTargetingClause) SetAdGroupId(v string) {
	o.AdGroupId = v
}

func (o SponsoredProductsCreateGlobalNegativeTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	toSerialize["adGroupId"] = o.AdGroupId
	return toSerialize, nil
}

type NullableSponsoredProductsCreateGlobalNegativeTargetingClause struct {
	value *SponsoredProductsCreateGlobalNegativeTargetingClause
	isSet bool
}

func (v NullableSponsoredProductsCreateGlobalNegativeTargetingClause) Get() *SponsoredProductsCreateGlobalNegativeTargetingClause {
	return v.value
}

func (v *NullableSponsoredProductsCreateGlobalNegativeTargetingClause) Set(val *SponsoredProductsCreateGlobalNegativeTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateGlobalNegativeTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateGlobalNegativeTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateGlobalNegativeTargetingClause(val *SponsoredProductsCreateGlobalNegativeTargetingClause) *NullableSponsoredProductsCreateGlobalNegativeTargetingClause {
	return &NullableSponsoredProductsCreateGlobalNegativeTargetingClause{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateGlobalNegativeTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateGlobalNegativeTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
