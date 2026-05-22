package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the Campaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Campaign{}

// Campaign struct for Campaign
type Campaign struct {
	BudgetType             BudgetType       `json:"budgetType"`
	RuleBasedBudget        *RuleBasedBudget `json:"ruleBasedBudget,omitempty"`
	BrandEntityId          *string          `json:"brandEntityId,omitempty"`
	IsMultiAdGroupsEnabled *bool            `json:"isMultiAdGroupsEnabled,omitempty"`
	// Goal will allow you to set goal type to help drive your campaign performance. If no goal is selected then it will default to PAGE_VISIT. The goal type of the campaign. - BRAND_IMPRESSION_SHARE - This goal will allow you grown your brand impression share on top of search placements - PAGE_VISIT [DEFAULT] - This goal drives traffic to your landing and detail pages through all placements - ACQUIRE_NEW_CUSTOMERS - This property is a PREVIEW ONLY and cannot be used as part of a request or response. This goal drives new customer acquisition for your brands. - AD_VIEWS - This property is a PREVIEW ONLY and cannot be used as part of a request or response. This goal maximizes view for your ads.
	Goal    *string  `json:"goal,omitempty"`
	Bidding *Bidding `json:"bidding,omitempty"`
	// The format of the date is YYYY-MM-DD.
	EndDate *string `json:"endDate,omitempty" validate:"regexp=^20[1-9][0-9]-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$"`
	// The identifier of the campaign.
	CampaignId      string           `json:"campaignId"`
	ProductLocation *ProductLocation `json:"productLocation,omitempty"`
	// A list of advertiser-specified custom identifiers for the campaign. Each customer identifier is a key-value pair. You can specify a maximum of 50 identifiers.
	Tags *map[string]string `json:"tags,omitempty"`
	// The identifier of an existing portfolio to which the campaign is associated.
	PortfolioId *string `json:"portfolioId,omitempty"`
	// The costType can be set to determines how the campaign will bid and charge. To view the bid maximums and minimums by geography and costType, see https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace - CPC [Default] - Cost per click. The performance of this campaign is measured by the clicks triggered by the ad. - VCPM - Cost per 1000 viewable impressions. The performance of this campaign is measured by the viewable impressions triggered by the ad.
	CostType *string `json:"costType,omitempty"`
	// The smartDefault specifies a list of the smart default options for the campaign.  `smartDefault` is optional for create campaign requests. `smartDefault` are applicable to all applicable child entities of the campaign and are not editable once the campaign is created. When using [\"TARGETING\"], targets will be automatically added based on the goal selected.  When [\"MANUAL\"] is selected, you will still be required to manually add targets.  If you don't specify `smartDefault`, default value will be applied based on `goal` . If campaign's `goal` is selected, `smartDefault` will be set to [\"TARGETING\"].  Otherwise, a campaign's `smartDefault` will be set to [\"MANUAL\"].  Each element in smartDefault can be set to determines which default strategy to be used - MANUAL - Manual settings, no smart default be applied to the campaign, if MANUAL is added in the list, no other items are allowed in the list (the list must contains only one item) - TARGETING - Smart Default Targeting creation, will automatically creating targetings when create ad group  Example: [\"TARGETING\"]
	SmartDefault []string `json:"smartDefault,omitempty"`
	// The name of the campaign.
	Name  string      `json:"name"`
	State EntityState `json:"state"`
	// The format of the date is YYYY-MM-DD.
	StartDate    *string               `json:"startDate,omitempty" validate:"regexp=^20[1-9][0-9]-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$"`
	Budget       float64               `json:"budget"`
	ExtendedData *CampaignExtendedData `json:"extendedData,omitempty"`
}

type _Campaign Campaign

// NewCampaign instantiates a new Campaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaign(budgetType BudgetType, campaignId string, name string, state EntityState, budget float64) *Campaign {
	this := Campaign{}
	this.BudgetType = budgetType
	this.CampaignId = campaignId
	this.Name = name
	this.State = state
	this.Budget = budget
	return &this
}

// NewCampaignWithDefaults instantiates a new Campaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignWithDefaults() *Campaign {
	this := Campaign{}
	return &this
}

// GetBudgetType returns the BudgetType field value
func (o *Campaign) GetBudgetType() BudgetType {
	if o == nil {
		var ret BudgetType
		return ret
	}

	return o.BudgetType
}

// GetBudgetTypeOk returns a tuple with the BudgetType field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetBudgetTypeOk() (*BudgetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetType, true
}

// SetBudgetType sets field value
func (o *Campaign) SetBudgetType(v BudgetType) {
	o.BudgetType = v
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

// GetBrandEntityId returns the BrandEntityId field value if set, zero value otherwise.
func (o *Campaign) GetBrandEntityId() string {
	if o == nil || IsNil(o.BrandEntityId) {
		var ret string
		return ret
	}
	return *o.BrandEntityId
}

// GetBrandEntityIdOk returns a tuple with the BrandEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetBrandEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrandEntityId) {
		return nil, false
	}
	return o.BrandEntityId, true
}

// HasBrandEntityId returns a boolean if a field has been set.
func (o *Campaign) HasBrandEntityId() bool {
	if o != nil && !IsNil(o.BrandEntityId) {
		return true
	}

	return false
}

// SetBrandEntityId gets a reference to the given string and assigns it to the BrandEntityId field.
func (o *Campaign) SetBrandEntityId(v string) {
	o.BrandEntityId = &v
}

// GetIsMultiAdGroupsEnabled returns the IsMultiAdGroupsEnabled field value if set, zero value otherwise.
func (o *Campaign) GetIsMultiAdGroupsEnabled() bool {
	if o == nil || IsNil(o.IsMultiAdGroupsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsMultiAdGroupsEnabled
}

// GetIsMultiAdGroupsEnabledOk returns a tuple with the IsMultiAdGroupsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetIsMultiAdGroupsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMultiAdGroupsEnabled) {
		return nil, false
	}
	return o.IsMultiAdGroupsEnabled, true
}

// HasIsMultiAdGroupsEnabled returns a boolean if a field has been set.
func (o *Campaign) HasIsMultiAdGroupsEnabled() bool {
	if o != nil && !IsNil(o.IsMultiAdGroupsEnabled) {
		return true
	}

	return false
}

// SetIsMultiAdGroupsEnabled gets a reference to the given bool and assigns it to the IsMultiAdGroupsEnabled field.
func (o *Campaign) SetIsMultiAdGroupsEnabled(v bool) {
	o.IsMultiAdGroupsEnabled = &v
}

// GetGoal returns the Goal field value if set, zero value otherwise.
func (o *Campaign) GetGoal() string {
	if o == nil || IsNil(o.Goal) {
		var ret string
		return ret
	}
	return *o.Goal
}

// GetGoalOk returns a tuple with the Goal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetGoalOk() (*string, bool) {
	if o == nil || IsNil(o.Goal) {
		return nil, false
	}
	return o.Goal, true
}

// HasGoal returns a boolean if a field has been set.
func (o *Campaign) HasGoal() bool {
	if o != nil && !IsNil(o.Goal) {
		return true
	}

	return false
}

// SetGoal gets a reference to the given string and assigns it to the Goal field.
func (o *Campaign) SetGoal(v string) {
	o.Goal = &v
}

// GetBidding returns the Bidding field value if set, zero value otherwise.
func (o *Campaign) GetBidding() Bidding {
	if o == nil || IsNil(o.Bidding) {
		var ret Bidding
		return ret
	}
	return *o.Bidding
}

// GetBiddingOk returns a tuple with the Bidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetBiddingOk() (*Bidding, bool) {
	if o == nil || IsNil(o.Bidding) {
		return nil, false
	}
	return o.Bidding, true
}

// HasBidding returns a boolean if a field has been set.
func (o *Campaign) HasBidding() bool {
	if o != nil && !IsNil(o.Bidding) {
		return true
	}

	return false
}

// SetBidding gets a reference to the given Bidding and assigns it to the Bidding field.
func (o *Campaign) SetBidding(v Bidding) {
	o.Bidding = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *Campaign) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *Campaign) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *Campaign) SetEndDate(v string) {
	o.EndDate = &v
}

// GetCampaignId returns the CampaignId field value
func (o *Campaign) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *Campaign) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetProductLocation returns the ProductLocation field value if set, zero value otherwise.
func (o *Campaign) GetProductLocation() ProductLocation {
	if o == nil || IsNil(o.ProductLocation) {
		var ret ProductLocation
		return ret
	}
	return *o.ProductLocation
}

// GetProductLocationOk returns a tuple with the ProductLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetProductLocationOk() (*ProductLocation, bool) {
	if o == nil || IsNil(o.ProductLocation) {
		return nil, false
	}
	return o.ProductLocation, true
}

// HasProductLocation returns a boolean if a field has been set.
func (o *Campaign) HasProductLocation() bool {
	if o != nil && !IsNil(o.ProductLocation) {
		return true
	}

	return false
}

// SetProductLocation gets a reference to the given ProductLocation and assigns it to the ProductLocation field.
func (o *Campaign) SetProductLocation(v ProductLocation) {
	o.ProductLocation = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Campaign) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Campaign) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Campaign) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetPortfolioId returns the PortfolioId field value if set, zero value otherwise.
func (o *Campaign) GetPortfolioId() string {
	if o == nil || IsNil(o.PortfolioId) {
		var ret string
		return ret
	}
	return *o.PortfolioId
}

// GetPortfolioIdOk returns a tuple with the PortfolioId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetPortfolioIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortfolioId) {
		return nil, false
	}
	return o.PortfolioId, true
}

// HasPortfolioId returns a boolean if a field has been set.
func (o *Campaign) HasPortfolioId() bool {
	if o != nil && !IsNil(o.PortfolioId) {
		return true
	}

	return false
}

// SetPortfolioId gets a reference to the given string and assigns it to the PortfolioId field.
func (o *Campaign) SetPortfolioId(v string) {
	o.PortfolioId = &v
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

// GetSmartDefault returns the SmartDefault field value if set, zero value otherwise.
func (o *Campaign) GetSmartDefault() []string {
	if o == nil || IsNil(o.SmartDefault) {
		var ret []string
		return ret
	}
	return o.SmartDefault
}

// GetSmartDefaultOk returns a tuple with the SmartDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetSmartDefaultOk() ([]string, bool) {
	if o == nil || IsNil(o.SmartDefault) {
		return nil, false
	}
	return o.SmartDefault, true
}

// HasSmartDefault returns a boolean if a field has been set.
func (o *Campaign) HasSmartDefault() bool {
	if o != nil && !IsNil(o.SmartDefault) {
		return true
	}

	return false
}

// SetSmartDefault gets a reference to the given []string and assigns it to the SmartDefault field.
func (o *Campaign) SetSmartDefault(v []string) {
	o.SmartDefault = v
}

// GetName returns the Name field value
func (o *Campaign) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Campaign) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *Campaign) GetState() EntityState {
	if o == nil {
		var ret EntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetStateOk() (*EntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Campaign) SetState(v EntityState) {
	o.State = v
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

// GetBudget returns the Budget field value
func (o *Campaign) GetBudget() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetBudgetOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Budget, true
}

// SetBudget sets field value
func (o *Campaign) SetBudget(v float64) {
	o.Budget = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *Campaign) GetExtendedData() CampaignExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret CampaignExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetExtendedDataOk() (*CampaignExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *Campaign) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given CampaignExtendedData and assigns it to the ExtendedData field.
func (o *Campaign) SetExtendedData(v CampaignExtendedData) {
	o.ExtendedData = &v
}

func (o Campaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["budgetType"] = o.BudgetType
	if !IsNil(o.RuleBasedBudget) {
		toSerialize["ruleBasedBudget"] = o.RuleBasedBudget
	}
	if !IsNil(o.BrandEntityId) {
		toSerialize["brandEntityId"] = o.BrandEntityId
	}
	if !IsNil(o.IsMultiAdGroupsEnabled) {
		toSerialize["isMultiAdGroupsEnabled"] = o.IsMultiAdGroupsEnabled
	}
	if !IsNil(o.Goal) {
		toSerialize["goal"] = o.Goal
	}
	if !IsNil(o.Bidding) {
		toSerialize["bidding"] = o.Bidding
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	toSerialize["campaignId"] = o.CampaignId
	if !IsNil(o.ProductLocation) {
		toSerialize["productLocation"] = o.ProductLocation
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.PortfolioId) {
		toSerialize["portfolioId"] = o.PortfolioId
	}
	if !IsNil(o.CostType) {
		toSerialize["costType"] = o.CostType
	}
	if !IsNil(o.SmartDefault) {
		toSerialize["smartDefault"] = o.SmartDefault
	}
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	toSerialize["budget"] = o.Budget
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
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
