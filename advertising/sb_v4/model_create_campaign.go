package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCampaign{}

// CreateCampaign struct for CreateCampaign
type CreateCampaign struct {
	BudgetType BudgetType `json:"budgetType"`
	// Please note that brandEntityId is only required for sellers. You can get the brandEntityId by calling the [GET /brands](https://advertising.amazon.com/API/docs/en-us/sponsored-brands/3-0/openapi#tag/Brands/operation/getBrands) endpoint.
	BrandEntityId *string `json:"brandEntityId,omitempty"`
	// Goal will allow you to set goal type to help drive your campaign performance. If no goal is selected then it will default to PAGE_VISIT. The goal type of the campaign. - BRAND_IMPRESSION_SHARE - This goal will allow you grown your brand impression share on top of search placement - PAGE_VISIT [DEFAULT] - This goal drives traffic to your landing and detail pages through all placements. - ACQUIRE_NEW_CUSTOMERS - This property is a PREVIEW ONLY and cannot be used as part of a request or response. This goal drives new customer acquisition for your brands. - AD_VIEWS - This property is a PREVIEW ONLY and cannot be used as part of a request or response. This goal maximizes view for your ads.
	Goal    *string  `json:"goal,omitempty"`
	Bidding *Bidding `json:"bidding,omitempty"`
	// endDate is optional. If endDate is specified, startDate must be specified as well.
	EndDate         *string          `json:"endDate,omitempty" validate:"regexp=^20[1-9][0-9]-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$"`
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
	Name  string                    `json:"name"`
	State CreateOrUpdateEntityState `json:"state"`
	// startDate is optional. If startDate is not specified, current date will be used.
	StartDate *string `json:"startDate,omitempty" validate:"regexp=^20[1-9][0-9]-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$"`
	// The budget of the campaign.
	Budget float64 `json:"budget"`
}

type _CreateCampaign CreateCampaign

// NewCreateCampaign instantiates a new CreateCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCampaign(budgetType BudgetType, name string, state CreateOrUpdateEntityState, budget float64) *CreateCampaign {
	this := CreateCampaign{}
	this.BudgetType = budgetType
	this.Name = name
	this.State = state
	this.Budget = budget
	return &this
}

// NewCreateCampaignWithDefaults instantiates a new CreateCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCampaignWithDefaults() *CreateCampaign {
	this := CreateCampaign{}
	return &this
}

// GetBudgetType returns the BudgetType field value
func (o *CreateCampaign) GetBudgetType() BudgetType {
	if o == nil {
		var ret BudgetType
		return ret
	}

	return o.BudgetType
}

// GetBudgetTypeOk returns a tuple with the BudgetType field value
// and a boolean to check if the value has been set.
func (o *CreateCampaign) GetBudgetTypeOk() (*BudgetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetType, true
}

// SetBudgetType sets field value
func (o *CreateCampaign) SetBudgetType(v BudgetType) {
	o.BudgetType = v
}

// GetBrandEntityId returns the BrandEntityId field value if set, zero value otherwise.
func (o *CreateCampaign) GetBrandEntityId() string {
	if o == nil || IsNil(o.BrandEntityId) {
		var ret string
		return ret
	}
	return *o.BrandEntityId
}

// GetBrandEntityIdOk returns a tuple with the BrandEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCampaign) GetBrandEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrandEntityId) {
		return nil, false
	}
	return o.BrandEntityId, true
}

// HasBrandEntityId returns a boolean if a field has been set.
func (o *CreateCampaign) HasBrandEntityId() bool {
	if o != nil && !IsNil(o.BrandEntityId) {
		return true
	}

	return false
}

// SetBrandEntityId gets a reference to the given string and assigns it to the BrandEntityId field.
func (o *CreateCampaign) SetBrandEntityId(v string) {
	o.BrandEntityId = &v
}

// GetGoal returns the Goal field value if set, zero value otherwise.
func (o *CreateCampaign) GetGoal() string {
	if o == nil || IsNil(o.Goal) {
		var ret string
		return ret
	}
	return *o.Goal
}

// GetGoalOk returns a tuple with the Goal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCampaign) GetGoalOk() (*string, bool) {
	if o == nil || IsNil(o.Goal) {
		return nil, false
	}
	return o.Goal, true
}

// HasGoal returns a boolean if a field has been set.
func (o *CreateCampaign) HasGoal() bool {
	if o != nil && !IsNil(o.Goal) {
		return true
	}

	return false
}

// SetGoal gets a reference to the given string and assigns it to the Goal field.
func (o *CreateCampaign) SetGoal(v string) {
	o.Goal = &v
}

// GetBidding returns the Bidding field value if set, zero value otherwise.
func (o *CreateCampaign) GetBidding() Bidding {
	if o == nil || IsNil(o.Bidding) {
		var ret Bidding
		return ret
	}
	return *o.Bidding
}

// GetBiddingOk returns a tuple with the Bidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCampaign) GetBiddingOk() (*Bidding, bool) {
	if o == nil || IsNil(o.Bidding) {
		return nil, false
	}
	return o.Bidding, true
}

// HasBidding returns a boolean if a field has been set.
func (o *CreateCampaign) HasBidding() bool {
	if o != nil && !IsNil(o.Bidding) {
		return true
	}

	return false
}

// SetBidding gets a reference to the given Bidding and assigns it to the Bidding field.
func (o *CreateCampaign) SetBidding(v Bidding) {
	o.Bidding = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CreateCampaign) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCampaign) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CreateCampaign) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *CreateCampaign) SetEndDate(v string) {
	o.EndDate = &v
}

// GetProductLocation returns the ProductLocation field value if set, zero value otherwise.
func (o *CreateCampaign) GetProductLocation() ProductLocation {
	if o == nil || IsNil(o.ProductLocation) {
		var ret ProductLocation
		return ret
	}
	return *o.ProductLocation
}

// GetProductLocationOk returns a tuple with the ProductLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCampaign) GetProductLocationOk() (*ProductLocation, bool) {
	if o == nil || IsNil(o.ProductLocation) {
		return nil, false
	}
	return o.ProductLocation, true
}

// HasProductLocation returns a boolean if a field has been set.
func (o *CreateCampaign) HasProductLocation() bool {
	if o != nil && !IsNil(o.ProductLocation) {
		return true
	}

	return false
}

// SetProductLocation gets a reference to the given ProductLocation and assigns it to the ProductLocation field.
func (o *CreateCampaign) SetProductLocation(v ProductLocation) {
	o.ProductLocation = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateCampaign) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCampaign) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateCampaign) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *CreateCampaign) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetPortfolioId returns the PortfolioId field value if set, zero value otherwise.
func (o *CreateCampaign) GetPortfolioId() string {
	if o == nil || IsNil(o.PortfolioId) {
		var ret string
		return ret
	}
	return *o.PortfolioId
}

// GetPortfolioIdOk returns a tuple with the PortfolioId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCampaign) GetPortfolioIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortfolioId) {
		return nil, false
	}
	return o.PortfolioId, true
}

// HasPortfolioId returns a boolean if a field has been set.
func (o *CreateCampaign) HasPortfolioId() bool {
	if o != nil && !IsNil(o.PortfolioId) {
		return true
	}

	return false
}

// SetPortfolioId gets a reference to the given string and assigns it to the PortfolioId field.
func (o *CreateCampaign) SetPortfolioId(v string) {
	o.PortfolioId = &v
}

// GetCostType returns the CostType field value if set, zero value otherwise.
func (o *CreateCampaign) GetCostType() string {
	if o == nil || IsNil(o.CostType) {
		var ret string
		return ret
	}
	return *o.CostType
}

// GetCostTypeOk returns a tuple with the CostType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCampaign) GetCostTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CostType) {
		return nil, false
	}
	return o.CostType, true
}

// HasCostType returns a boolean if a field has been set.
func (o *CreateCampaign) HasCostType() bool {
	if o != nil && !IsNil(o.CostType) {
		return true
	}

	return false
}

// SetCostType gets a reference to the given string and assigns it to the CostType field.
func (o *CreateCampaign) SetCostType(v string) {
	o.CostType = &v
}

// GetSmartDefault returns the SmartDefault field value if set, zero value otherwise.
func (o *CreateCampaign) GetSmartDefault() []string {
	if o == nil || IsNil(o.SmartDefault) {
		var ret []string
		return ret
	}
	return o.SmartDefault
}

// GetSmartDefaultOk returns a tuple with the SmartDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCampaign) GetSmartDefaultOk() ([]string, bool) {
	if o == nil || IsNil(o.SmartDefault) {
		return nil, false
	}
	return o.SmartDefault, true
}

// HasSmartDefault returns a boolean if a field has been set.
func (o *CreateCampaign) HasSmartDefault() bool {
	if o != nil && !IsNil(o.SmartDefault) {
		return true
	}

	return false
}

// SetSmartDefault gets a reference to the given []string and assigns it to the SmartDefault field.
func (o *CreateCampaign) SetSmartDefault(v []string) {
	o.SmartDefault = v
}

// GetName returns the Name field value
func (o *CreateCampaign) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateCampaign) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateCampaign) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *CreateCampaign) GetState() CreateOrUpdateEntityState {
	if o == nil {
		var ret CreateOrUpdateEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreateCampaign) GetStateOk() (*CreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreateCampaign) SetState(v CreateOrUpdateEntityState) {
	o.State = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CreateCampaign) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCampaign) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CreateCampaign) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *CreateCampaign) SetStartDate(v string) {
	o.StartDate = &v
}

// GetBudget returns the Budget field value
func (o *CreateCampaign) GetBudget() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value
// and a boolean to check if the value has been set.
func (o *CreateCampaign) GetBudgetOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Budget, true
}

// SetBudget sets field value
func (o *CreateCampaign) SetBudget(v float64) {
	o.Budget = v
}

func (o CreateCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["budgetType"] = o.BudgetType
	if !IsNil(o.BrandEntityId) {
		toSerialize["brandEntityId"] = o.BrandEntityId
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
	return toSerialize, nil
}

type NullableCreateCampaign struct {
	value *CreateCampaign
	isSet bool
}

func (v NullableCreateCampaign) Get() *CreateCampaign {
	return v.value
}

func (v *NullableCreateCampaign) Set(val *CreateCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCampaign(val *CreateCampaign) *NullableCreateCampaign {
	return &NullableCreateCampaign{value: val, isSet: true}
}

func (v NullableCreateCampaign) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
