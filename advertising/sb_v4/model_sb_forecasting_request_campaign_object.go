package sb_v4

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the SBForecastingRequestCampaignObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingRequestCampaignObject{}

// SBForecastingRequestCampaignObject The campaign settings.
type SBForecastingRequestCampaignObject struct {
	// The amount of the budget.
	Budget float64 `json:"budget"`
	// Budget can be set to DAILY or LIFETIME.   |BudgetType|Description| |-----------|-----------| |DAILY| The amount that you're willing to spend on this campaign each day. If the campaign spends less than your daily budget, the unspent amount can be used to increase your daily budget on the other days of the calendar month.| |LIFETIME| The total amount that you are willing to spend on this campaign.|
	BudgetType string `json:"budgetType"`
	// The forecast type. can be set to WEEKLY or MONTHLY.   **If have not set the forecastType during campaign creation then use WEEKLY as goal value.**
	ForecastType string `json:"forecastType"`
	// The YYYY-MM-DD start date for the campaign. If this field is not set to a value, the current date is used.
	StartDate *time.Time `json:"startDate,omitempty"`
	// The YYYY-MM-DD end date for the campaign. Must be greater than the value for `startDate`. If not specified, the campaign has no end date and runs continuously.
	EndDate *time.Time `json:"endDate,omitempty"`
	// Goal will allow you to set goal type to help drive your campaign performance.   **If have not set the goal during campaign creation then use PAGE_VISIT as goal type.**    The goal type of the campaign. Initial launch only supports PAGE_VISIT.   BRAND_IMPRESSION_SHARE - This goal is a PREVIEW ONLY and cannot be set currently. It will allow you grown your brand impression share on top of search placement.   PAGE_VISIT - This goal drives traffic to your landing and detail pages through all placements.   ACQUIRE_NEW_CUSTOMERS - This property is a PREVIEW ONLY and cannot be used as part of a request or response. This goal drives new customer acquisition for your brands.   AD_VIEWS - This property is a PREVIEW ONLY and cannot be used as part of a request or response. This goal maximizes view for your ads.
	Goal     *string                `json:"goal,omitempty"`
	AdGroups []SBForecastingAdGroup `json:"adGroups"`
}

type _SBForecastingRequestCampaignObject SBForecastingRequestCampaignObject

// NewSBForecastingRequestCampaignObject instantiates a new SBForecastingRequestCampaignObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingRequestCampaignObject(budget float64, budgetType string, forecastType string, adGroups []SBForecastingAdGroup) *SBForecastingRequestCampaignObject {
	this := SBForecastingRequestCampaignObject{}
	this.Budget = budget
	this.BudgetType = budgetType
	this.ForecastType = forecastType
	this.AdGroups = adGroups
	return &this
}

// NewSBForecastingRequestCampaignObjectWithDefaults instantiates a new SBForecastingRequestCampaignObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingRequestCampaignObjectWithDefaults() *SBForecastingRequestCampaignObject {
	this := SBForecastingRequestCampaignObject{}
	return &this
}

// GetBudget returns the Budget field value
func (o *SBForecastingRequestCampaignObject) GetBudget() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value
// and a boolean to check if the value has been set.
func (o *SBForecastingRequestCampaignObject) GetBudgetOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Budget, true
}

// SetBudget sets field value
func (o *SBForecastingRequestCampaignObject) SetBudget(v float64) {
	o.Budget = v
}

// GetBudgetType returns the BudgetType field value
func (o *SBForecastingRequestCampaignObject) GetBudgetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BudgetType
}

// GetBudgetTypeOk returns a tuple with the BudgetType field value
// and a boolean to check if the value has been set.
func (o *SBForecastingRequestCampaignObject) GetBudgetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetType, true
}

// SetBudgetType sets field value
func (o *SBForecastingRequestCampaignObject) SetBudgetType(v string) {
	o.BudgetType = v
}

// GetForecastType returns the ForecastType field value
func (o *SBForecastingRequestCampaignObject) GetForecastType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ForecastType
}

// GetForecastTypeOk returns a tuple with the ForecastType field value
// and a boolean to check if the value has been set.
func (o *SBForecastingRequestCampaignObject) GetForecastTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForecastType, true
}

// SetForecastType sets field value
func (o *SBForecastingRequestCampaignObject) SetForecastType(v string) {
	o.ForecastType = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SBForecastingRequestCampaignObject) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingRequestCampaignObject) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SBForecastingRequestCampaignObject) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *SBForecastingRequestCampaignObject) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SBForecastingRequestCampaignObject) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingRequestCampaignObject) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SBForecastingRequestCampaignObject) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *SBForecastingRequestCampaignObject) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetGoal returns the Goal field value if set, zero value otherwise.
func (o *SBForecastingRequestCampaignObject) GetGoal() string {
	if o == nil || IsNil(o.Goal) {
		var ret string
		return ret
	}
	return *o.Goal
}

// GetGoalOk returns a tuple with the Goal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingRequestCampaignObject) GetGoalOk() (*string, bool) {
	if o == nil || IsNil(o.Goal) {
		return nil, false
	}
	return o.Goal, true
}

// HasGoal returns a boolean if a field has been set.
func (o *SBForecastingRequestCampaignObject) HasGoal() bool {
	if o != nil && !IsNil(o.Goal) {
		return true
	}

	return false
}

// SetGoal gets a reference to the given string and assigns it to the Goal field.
func (o *SBForecastingRequestCampaignObject) SetGoal(v string) {
	o.Goal = &v
}

// GetAdGroups returns the AdGroups field value
func (o *SBForecastingRequestCampaignObject) GetAdGroups() []SBForecastingAdGroup {
	if o == nil {
		var ret []SBForecastingAdGroup
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *SBForecastingRequestCampaignObject) GetAdGroupsOk() ([]SBForecastingAdGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdGroups, true
}

// SetAdGroups sets field value
func (o *SBForecastingRequestCampaignObject) SetAdGroups(v []SBForecastingAdGroup) {
	o.AdGroups = v
}

func (o SBForecastingRequestCampaignObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["budget"] = o.Budget
	toSerialize["budgetType"] = o.BudgetType
	toSerialize["forecastType"] = o.ForecastType
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Goal) {
		toSerialize["goal"] = o.Goal
	}
	toSerialize["adGroups"] = o.AdGroups
	return toSerialize, nil
}

type NullableSBForecastingRequestCampaignObject struct {
	value *SBForecastingRequestCampaignObject
	isSet bool
}

func (v NullableSBForecastingRequestCampaignObject) Get() *SBForecastingRequestCampaignObject {
	return v.value
}

func (v *NullableSBForecastingRequestCampaignObject) Set(val *SBForecastingRequestCampaignObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingRequestCampaignObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingRequestCampaignObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingRequestCampaignObject(val *SBForecastingRequestCampaignObject) *NullableSBForecastingRequestCampaignObject {
	return &NullableSBForecastingRequestCampaignObject{value: val, isSet: true}
}

func (v NullableSBForecastingRequestCampaignObject) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingRequestCampaignObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
