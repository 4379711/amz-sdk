package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the Campaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Campaign{}

// Campaign struct for Campaign
type Campaign struct {
	// The name of the campaign.
	Name *string `json:"name,omitempty"`
	// The time period over which the amount specified in the `budget` property is allocated.
	BudgetType *string `json:"budgetType,omitempty"`
	// The amount of the budget.
	Budget *float64 `json:"budget,omitempty"`
	// The YYYYMMDD start date of the campaign. The date must be today or in the future.
	StartDate *string `json:"startDate,omitempty"`
	// The YYYYMMDD end date of the campaign.
	EndDate NullableString `json:"endDate,omitempty"`
	// Determines how the campaign will bid and charge. |Name|Description| |----|----------| |cpc |[Default] The performance of this campaign is measured by the clicks triggered by the ad.| |vcpm |The performance of this campaign is measured by the viewed impressions triggered by the ad. |  To view minimum and maximum bids based on the costType, see [Limits](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace).
	CostType *string `json:"costType,omitempty"`
	// The state of the campaign.
	State *string `json:"state,omitempty"`
	// Identifier of the portfolio that will be associated with the campaign. If null then the campaign will be disassociated from existing portfolio. Campaigns with CPC and vCPM costType are supported.
	PortfolioId NullableInt64 `json:"portfolioId,omitempty"`
	// The identifier of the campaign.
	CampaignId      *int64           `json:"campaignId,omitempty"`
	Tactic          *Tactic          `json:"tactic,omitempty"`
	DeliveryProfile *string          `json:"deliveryProfile,omitempty"`
	RuleBasedBudget *RuleBasedBudget `json:"ruleBasedBudget,omitempty"`
}

// NewCampaign instantiates a new Campaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaign() *Campaign {
	this := Campaign{}
	return &this
}

// NewCampaignWithDefaults instantiates a new Campaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignWithDefaults() *Campaign {
	this := Campaign{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Campaign) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Campaign) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Campaign) SetName(v string) {
	o.Name = &v
}

// GetBudgetType returns the BudgetType field value if set, zero value otherwise.
func (o *Campaign) GetBudgetType() string {
	if o == nil || IsNil(o.BudgetType) {
		var ret string
		return ret
	}
	return *o.BudgetType
}

// GetBudgetTypeOk returns a tuple with the BudgetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetBudgetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BudgetType) {
		return nil, false
	}
	return o.BudgetType, true
}

// HasBudgetType returns a boolean if a field has been set.
func (o *Campaign) HasBudgetType() bool {
	if o != nil && !IsNil(o.BudgetType) {
		return true
	}

	return false
}

// SetBudgetType gets a reference to the given string and assigns it to the BudgetType field.
func (o *Campaign) SetBudgetType(v string) {
	o.BudgetType = &v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *Campaign) GetBudget() float64 {
	if o == nil || IsNil(o.Budget) {
		var ret float64
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetBudgetOk() (*float64, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *Campaign) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given float64 and assigns it to the Budget field.
func (o *Campaign) SetBudget(v float64) {
	o.Budget = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Campaign) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Campaign) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *Campaign) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Campaign) GetEndDate() string {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Campaign) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *Campaign) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *Campaign) SetEndDate(v string) {
	o.EndDate.Set(&v)
}

// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *Campaign) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *Campaign) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetCostType returns the CostType field value if set, zero value otherwise.
func (o *Campaign) GetCostType() string {
	if o == nil || IsNil(o.CostType) {
		var ret string
		return ret
	}
	return *o.CostType
}

// GetCostTypeOk returns a tuple with the CostType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetCostTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CostType) {
		return nil, false
	}
	return o.CostType, true
}

// HasCostType returns a boolean if a field has been set.
func (o *Campaign) HasCostType() bool {
	if o != nil && !IsNil(o.CostType) {
		return true
	}

	return false
}

// SetCostType gets a reference to the given string and assigns it to the CostType field.
func (o *Campaign) SetCostType(v string) {
	o.CostType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Campaign) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Campaign) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Campaign) SetState(v string) {
	o.State = &v
}

// GetPortfolioId returns the PortfolioId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Campaign) GetPortfolioId() int64 {
	if o == nil || IsNil(o.PortfolioId.Get()) {
		var ret int64
		return ret
	}
	return *o.PortfolioId.Get()
}

// GetPortfolioIdOk returns a tuple with the PortfolioId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Campaign) GetPortfolioIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortfolioId.Get(), o.PortfolioId.IsSet()
}

// HasPortfolioId returns a boolean if a field has been set.
func (o *Campaign) HasPortfolioId() bool {
	if o != nil && o.PortfolioId.IsSet() {
		return true
	}

	return false
}

// SetPortfolioId gets a reference to the given NullableInt64 and assigns it to the PortfolioId field.
func (o *Campaign) SetPortfolioId(v int64) {
	o.PortfolioId.Set(&v)
}

// SetPortfolioIdNil sets the value for PortfolioId to be an explicit nil
func (o *Campaign) SetPortfolioIdNil() {
	o.PortfolioId.Set(nil)
}

// UnsetPortfolioId ensures that no value is present for PortfolioId, not even an explicit nil
func (o *Campaign) UnsetPortfolioId() {
	o.PortfolioId.Unset()
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *Campaign) GetCampaignId() int64 {
	if o == nil || IsNil(o.CampaignId) {
		var ret int64
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetCampaignIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *Campaign) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given int64 and assigns it to the CampaignId field.
func (o *Campaign) SetCampaignId(v int64) {
	o.CampaignId = &v
}

// GetTactic returns the Tactic field value if set, zero value otherwise.
func (o *Campaign) GetTactic() Tactic {
	if o == nil || IsNil(o.Tactic) {
		var ret Tactic
		return ret
	}
	return *o.Tactic
}

// GetTacticOk returns a tuple with the Tactic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetTacticOk() (*Tactic, bool) {
	if o == nil || IsNil(o.Tactic) {
		return nil, false
	}
	return o.Tactic, true
}

// HasTactic returns a boolean if a field has been set.
func (o *Campaign) HasTactic() bool {
	if o != nil && !IsNil(o.Tactic) {
		return true
	}

	return false
}

// SetTactic gets a reference to the given Tactic and assigns it to the Tactic field.
func (o *Campaign) SetTactic(v Tactic) {
	o.Tactic = &v
}

// GetDeliveryProfile returns the DeliveryProfile field value if set, zero value otherwise.
func (o *Campaign) GetDeliveryProfile() string {
	if o == nil || IsNil(o.DeliveryProfile) {
		var ret string
		return ret
	}
	return *o.DeliveryProfile
}

// GetDeliveryProfileOk returns a tuple with the DeliveryProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetDeliveryProfileOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryProfile) {
		return nil, false
	}
	return o.DeliveryProfile, true
}

// HasDeliveryProfile returns a boolean if a field has been set.
func (o *Campaign) HasDeliveryProfile() bool {
	if o != nil && !IsNil(o.DeliveryProfile) {
		return true
	}

	return false
}

// SetDeliveryProfile gets a reference to the given string and assigns it to the DeliveryProfile field.
func (o *Campaign) SetDeliveryProfile(v string) {
	o.DeliveryProfile = &v
}

// GetRuleBasedBudget returns the RuleBasedBudget field value if set, zero value otherwise.
func (o *Campaign) GetRuleBasedBudget() RuleBasedBudget {
	if o == nil || IsNil(o.RuleBasedBudget) {
		var ret RuleBasedBudget
		return ret
	}
	return *o.RuleBasedBudget
}

// GetRuleBasedBudgetOk returns a tuple with the RuleBasedBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetRuleBasedBudgetOk() (*RuleBasedBudget, bool) {
	if o == nil || IsNil(o.RuleBasedBudget) {
		return nil, false
	}
	return o.RuleBasedBudget, true
}

// HasRuleBasedBudget returns a boolean if a field has been set.
func (o *Campaign) HasRuleBasedBudget() bool {
	if o != nil && !IsNil(o.RuleBasedBudget) {
		return true
	}

	return false
}

// SetRuleBasedBudget gets a reference to the given RuleBasedBudget and assigns it to the RuleBasedBudget field.
func (o *Campaign) SetRuleBasedBudget(v RuleBasedBudget) {
	o.RuleBasedBudget = &v
}

func (o Campaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.BudgetType) {
		toSerialize["budgetType"] = o.BudgetType
	}
	if !IsNil(o.Budget) {
		toSerialize["budget"] = o.Budget
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	if !IsNil(o.CostType) {
		toSerialize["costType"] = o.CostType
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if o.PortfolioId.IsSet() {
		toSerialize["portfolioId"] = o.PortfolioId.Get()
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	if !IsNil(o.Tactic) {
		toSerialize["tactic"] = o.Tactic
	}
	if !IsNil(o.DeliveryProfile) {
		toSerialize["deliveryProfile"] = o.DeliveryProfile
	}
	if !IsNil(o.RuleBasedBudget) {
		toSerialize["ruleBasedBudget"] = o.RuleBasedBudget
	}
	return toSerialize, nil
}

type NullableCampaign struct {
	value *Campaign
	isSet bool
}

func (v NullableCampaign) Get() *Campaign {
	return v.value
}

func (v *NullableCampaign) Set(val *Campaign) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaign(val *Campaign) *NullableCampaign {
	return &NullableCampaign{value: val, isSet: true}
}

func (v NullableCampaign) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
