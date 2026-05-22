package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CampaignResponseEx type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignResponseEx{}

// CampaignResponseEx struct for CampaignResponseEx
type CampaignResponseEx struct {
	// The identifier of the campaign.
	CampaignId *float32 `json:"campaignId,omitempty"`
	// The name of the campaign.
	Name   *string `json:"name,omitempty"`
	Tactic *Tactic `json:"tactic,omitempty"`
	// The time period over which the amount specified in the `budget` property is allocated.
	BudgetType *string `json:"budgetType,omitempty"`
	// The amount of the budget.
	Budget *float64 `json:"budget,omitempty"`
	// The YYYYMMDD start date of the campaign. The date must be today or in the future.
	StartDate *string `json:"startDate,omitempty"`
	// The YYYYMMDD end date of the campaign.
	EndDate *string `json:"endDate,omitempty"`
	// The state of the campaign.
	State *string `json:"state,omitempty"`
	// Identifier of the portfolio that will be associated with the campaign. If null then the campaign will be disassociated from existing portfolio. Campaigns with CPC and vCPM costType are supported.
	PortfolioId *int64 `json:"portfolioId,omitempty"`
	// The status of the campaign.
	ServingStatus *string `json:"servingStatus,omitempty"`
	// Determines how the campaign will bid and charge. |Name|Description| |----|----------|-----------| |cpc |[Default] The performance of this campaign is measured by the clicks triggered by the ad.| |vcpm|The performance of this campaign is measured by the viewed impressions triggered by the ad. $1 is the minimum bid for vCPM.|
	CostType *string `json:"costType,omitempty"`
	// Epoch date the campaign was created.
	CreationDate *int64 `json:"creationDate,omitempty"`
	// Epoch date of the last update to any property associated with the campaign.
	LastUpdatedDate *int64           `json:"lastUpdatedDate,omitempty"`
	RuleBasedBudget *RuleBasedBudget `json:"ruleBasedBudget,omitempty"`
}

// NewCampaignResponseEx instantiates a new CampaignResponseEx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignResponseEx() *CampaignResponseEx {
	this := CampaignResponseEx{}
	return &this
}

// NewCampaignResponseExWithDefaults instantiates a new CampaignResponseEx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignResponseExWithDefaults() *CampaignResponseEx {
	this := CampaignResponseEx{}
	return &this
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *CampaignResponseEx) GetCampaignId() float32 {
	if o == nil || IsNil(o.CampaignId) {
		var ret float32
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponseEx) GetCampaignIdOk() (*float32, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *CampaignResponseEx) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given float32 and assigns it to the CampaignId field.
func (o *CampaignResponseEx) SetCampaignId(v float32) {
	o.CampaignId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CampaignResponseEx) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponseEx) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CampaignResponseEx) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CampaignResponseEx) SetName(v string) {
	o.Name = &v
}

// GetTactic returns the Tactic field value if set, zero value otherwise.
func (o *CampaignResponseEx) GetTactic() Tactic {
	if o == nil || IsNil(o.Tactic) {
		var ret Tactic
		return ret
	}
	return *o.Tactic
}

// GetTacticOk returns a tuple with the Tactic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponseEx) GetTacticOk() (*Tactic, bool) {
	if o == nil || IsNil(o.Tactic) {
		return nil, false
	}
	return o.Tactic, true
}

// HasTactic returns a boolean if a field has been set.
func (o *CampaignResponseEx) HasTactic() bool {
	if o != nil && !IsNil(o.Tactic) {
		return true
	}

	return false
}

// SetTactic gets a reference to the given Tactic and assigns it to the Tactic field.
func (o *CampaignResponseEx) SetTactic(v Tactic) {
	o.Tactic = &v
}

// GetBudgetType returns the BudgetType field value if set, zero value otherwise.
func (o *CampaignResponseEx) GetBudgetType() string {
	if o == nil || IsNil(o.BudgetType) {
		var ret string
		return ret
	}
	return *o.BudgetType
}

// GetBudgetTypeOk returns a tuple with the BudgetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponseEx) GetBudgetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BudgetType) {
		return nil, false
	}
	return o.BudgetType, true
}

// HasBudgetType returns a boolean if a field has been set.
func (o *CampaignResponseEx) HasBudgetType() bool {
	if o != nil && !IsNil(o.BudgetType) {
		return true
	}

	return false
}

// SetBudgetType gets a reference to the given string and assigns it to the BudgetType field.
func (o *CampaignResponseEx) SetBudgetType(v string) {
	o.BudgetType = &v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *CampaignResponseEx) GetBudget() float64 {
	if o == nil || IsNil(o.Budget) {
		var ret float64
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponseEx) GetBudgetOk() (*float64, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *CampaignResponseEx) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given float64 and assigns it to the Budget field.
func (o *CampaignResponseEx) SetBudget(v float64) {
	o.Budget = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CampaignResponseEx) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponseEx) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CampaignResponseEx) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *CampaignResponseEx) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CampaignResponseEx) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponseEx) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CampaignResponseEx) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *CampaignResponseEx) SetEndDate(v string) {
	o.EndDate = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CampaignResponseEx) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponseEx) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CampaignResponseEx) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CampaignResponseEx) SetState(v string) {
	o.State = &v
}

// GetPortfolioId returns the PortfolioId field value if set, zero value otherwise.
func (o *CampaignResponseEx) GetPortfolioId() int64 {
	if o == nil || IsNil(o.PortfolioId) {
		var ret int64
		return ret
	}
	return *o.PortfolioId
}

// GetPortfolioIdOk returns a tuple with the PortfolioId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponseEx) GetPortfolioIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PortfolioId) {
		return nil, false
	}
	return o.PortfolioId, true
}

// HasPortfolioId returns a boolean if a field has been set.
func (o *CampaignResponseEx) HasPortfolioId() bool {
	if o != nil && !IsNil(o.PortfolioId) {
		return true
	}

	return false
}

// SetPortfolioId gets a reference to the given int64 and assigns it to the PortfolioId field.
func (o *CampaignResponseEx) SetPortfolioId(v int64) {
	o.PortfolioId = &v
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *CampaignResponseEx) GetServingStatus() string {
	if o == nil || IsNil(o.ServingStatus) {
		var ret string
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponseEx) GetServingStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *CampaignResponseEx) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given string and assigns it to the ServingStatus field.
func (o *CampaignResponseEx) SetServingStatus(v string) {
	o.ServingStatus = &v
}

// GetCostType returns the CostType field value if set, zero value otherwise.
func (o *CampaignResponseEx) GetCostType() string {
	if o == nil || IsNil(o.CostType) {
		var ret string
		return ret
	}
	return *o.CostType
}

// GetCostTypeOk returns a tuple with the CostType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponseEx) GetCostTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CostType) {
		return nil, false
	}
	return o.CostType, true
}

// HasCostType returns a boolean if a field has been set.
func (o *CampaignResponseEx) HasCostType() bool {
	if o != nil && !IsNil(o.CostType) {
		return true
	}

	return false
}

// SetCostType gets a reference to the given string and assigns it to the CostType field.
func (o *CampaignResponseEx) SetCostType(v string) {
	o.CostType = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *CampaignResponseEx) GetCreationDate() int64 {
	if o == nil || IsNil(o.CreationDate) {
		var ret int64
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponseEx) GetCreationDateOk() (*int64, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *CampaignResponseEx) HasCreationDate() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given int64 and assigns it to the CreationDate field.
func (o *CampaignResponseEx) SetCreationDate(v int64) {
	o.CreationDate = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *CampaignResponseEx) GetLastUpdatedDate() int64 {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret int64
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponseEx) GetLastUpdatedDateOk() (*int64, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *CampaignResponseEx) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given int64 and assigns it to the LastUpdatedDate field.
func (o *CampaignResponseEx) SetLastUpdatedDate(v int64) {
	o.LastUpdatedDate = &v
}

// GetRuleBasedBudget returns the RuleBasedBudget field value if set, zero value otherwise.
func (o *CampaignResponseEx) GetRuleBasedBudget() RuleBasedBudget {
	if o == nil || IsNil(o.RuleBasedBudget) {
		var ret RuleBasedBudget
		return ret
	}
	return *o.RuleBasedBudget
}

// GetRuleBasedBudgetOk returns a tuple with the RuleBasedBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponseEx) GetRuleBasedBudgetOk() (*RuleBasedBudget, bool) {
	if o == nil || IsNil(o.RuleBasedBudget) {
		return nil, false
	}
	return o.RuleBasedBudget, true
}

// HasRuleBasedBudget returns a boolean if a field has been set.
func (o *CampaignResponseEx) HasRuleBasedBudget() bool {
	if o != nil && !IsNil(o.RuleBasedBudget) {
		return true
	}

	return false
}

// SetRuleBasedBudget gets a reference to the given RuleBasedBudget and assigns it to the RuleBasedBudget field.
func (o *CampaignResponseEx) SetRuleBasedBudget(v RuleBasedBudget) {
	o.RuleBasedBudget = &v
}

func (o CampaignResponseEx) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Tactic) {
		toSerialize["tactic"] = o.Tactic
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
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.PortfolioId) {
		toSerialize["portfolioId"] = o.PortfolioId
	}
	if !IsNil(o.ServingStatus) {
		toSerialize["servingStatus"] = o.ServingStatus
	}
	if !IsNil(o.CostType) {
		toSerialize["costType"] = o.CostType
	}
	if !IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !IsNil(o.LastUpdatedDate) {
		toSerialize["lastUpdatedDate"] = o.LastUpdatedDate
	}
	if !IsNil(o.RuleBasedBudget) {
		toSerialize["ruleBasedBudget"] = o.RuleBasedBudget
	}
	return toSerialize, nil
}

type NullableCampaignResponseEx struct {
	value *CampaignResponseEx
	isSet bool
}

func (v NullableCampaignResponseEx) Get() *CampaignResponseEx {
	return v.value
}

func (v *NullableCampaignResponseEx) Set(val *CampaignResponseEx) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignResponseEx) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignResponseEx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignResponseEx(val *CampaignResponseEx) *NullableCampaignResponseEx {
	return &NullableCampaignResponseEx{value: val, isSet: true}
}

func (v NullableCampaignResponseEx) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCampaignResponseEx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
