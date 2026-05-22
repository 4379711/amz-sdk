package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the InitialGlobalBudgetRecommendationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InitialGlobalBudgetRecommendationRequest{}

// InitialGlobalBudgetRecommendationRequest struct for InitialGlobalBudgetRecommendationRequest
type InitialGlobalBudgetRecommendationRequest struct {
	// The budget value for each country of new campaign.
	CountryDailyBudgets *map[string]float32 `json:"countryDailyBudgets,omitempty"`
	Bidding             Bidding             `json:"bidding"`
	// The ad group information for this new campaign.
	AdGroups []GlobalAdGroup `json:"adGroups"`
	// The end date of the campaign in YYYYMMDD format.
	EndDate *string `json:"endDate,omitempty"`
	// Specifies the targeting type.
	TargetingType string `json:"targetingType"`
	// The start date of the campaign in YYYYMMDD format.
	StartDate *string `json:"startDate,omitempty"`
}

type _InitialGlobalBudgetRecommendationRequest InitialGlobalBudgetRecommendationRequest

// NewInitialGlobalBudgetRecommendationRequest instantiates a new InitialGlobalBudgetRecommendationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitialGlobalBudgetRecommendationRequest(bidding Bidding, adGroups []GlobalAdGroup, targetingType string) *InitialGlobalBudgetRecommendationRequest {
	this := InitialGlobalBudgetRecommendationRequest{}
	this.Bidding = bidding
	this.AdGroups = adGroups
	this.TargetingType = targetingType
	return &this
}

// NewInitialGlobalBudgetRecommendationRequestWithDefaults instantiates a new InitialGlobalBudgetRecommendationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitialGlobalBudgetRecommendationRequestWithDefaults() *InitialGlobalBudgetRecommendationRequest {
	this := InitialGlobalBudgetRecommendationRequest{}
	var targetingType string = "AUTO"
	this.TargetingType = targetingType
	return &this
}

// GetCountryDailyBudgets returns the CountryDailyBudgets field value if set, zero value otherwise.
func (o *InitialGlobalBudgetRecommendationRequest) GetCountryDailyBudgets() map[string]float32 {
	if o == nil || IsNil(o.CountryDailyBudgets) {
		var ret map[string]float32
		return ret
	}
	return *o.CountryDailyBudgets
}

// GetCountryDailyBudgetsOk returns a tuple with the CountryDailyBudgets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitialGlobalBudgetRecommendationRequest) GetCountryDailyBudgetsOk() (*map[string]float32, bool) {
	if o == nil || IsNil(o.CountryDailyBudgets) {
		return nil, false
	}
	return o.CountryDailyBudgets, true
}

// HasCountryDailyBudgets returns a boolean if a field has been set.
func (o *InitialGlobalBudgetRecommendationRequest) HasCountryDailyBudgets() bool {
	if o != nil && !IsNil(o.CountryDailyBudgets) {
		return true
	}

	return false
}

// SetCountryDailyBudgets gets a reference to the given map[string]float32 and assigns it to the CountryDailyBudgets field.
func (o *InitialGlobalBudgetRecommendationRequest) SetCountryDailyBudgets(v map[string]float32) {
	o.CountryDailyBudgets = &v
}

// GetBidding returns the Bidding field value
func (o *InitialGlobalBudgetRecommendationRequest) GetBidding() Bidding {
	if o == nil {
		var ret Bidding
		return ret
	}

	return o.Bidding
}

// GetBiddingOk returns a tuple with the Bidding field value
// and a boolean to check if the value has been set.
func (o *InitialGlobalBudgetRecommendationRequest) GetBiddingOk() (*Bidding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bidding, true
}

// SetBidding sets field value
func (o *InitialGlobalBudgetRecommendationRequest) SetBidding(v Bidding) {
	o.Bidding = v
}

// GetAdGroups returns the AdGroups field value
func (o *InitialGlobalBudgetRecommendationRequest) GetAdGroups() []GlobalAdGroup {
	if o == nil {
		var ret []GlobalAdGroup
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *InitialGlobalBudgetRecommendationRequest) GetAdGroupsOk() ([]GlobalAdGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdGroups, true
}

// SetAdGroups sets field value
func (o *InitialGlobalBudgetRecommendationRequest) SetAdGroups(v []GlobalAdGroup) {
	o.AdGroups = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *InitialGlobalBudgetRecommendationRequest) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitialGlobalBudgetRecommendationRequest) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *InitialGlobalBudgetRecommendationRequest) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *InitialGlobalBudgetRecommendationRequest) SetEndDate(v string) {
	o.EndDate = &v
}

// GetTargetingType returns the TargetingType field value
func (o *InitialGlobalBudgetRecommendationRequest) GetTargetingType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetingType
}

// GetTargetingTypeOk returns a tuple with the TargetingType field value
// and a boolean to check if the value has been set.
func (o *InitialGlobalBudgetRecommendationRequest) GetTargetingTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingType, true
}

// SetTargetingType sets field value
func (o *InitialGlobalBudgetRecommendationRequest) SetTargetingType(v string) {
	o.TargetingType = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *InitialGlobalBudgetRecommendationRequest) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitialGlobalBudgetRecommendationRequest) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *InitialGlobalBudgetRecommendationRequest) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *InitialGlobalBudgetRecommendationRequest) SetStartDate(v string) {
	o.StartDate = &v
}

func (o InitialGlobalBudgetRecommendationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryDailyBudgets) {
		toSerialize["countryDailyBudgets"] = o.CountryDailyBudgets
	}
	toSerialize["bidding"] = o.Bidding
	toSerialize["adGroups"] = o.AdGroups
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	toSerialize["targetingType"] = o.TargetingType
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	return toSerialize, nil
}

type NullableInitialGlobalBudgetRecommendationRequest struct {
	value *InitialGlobalBudgetRecommendationRequest
	isSet bool
}

func (v NullableInitialGlobalBudgetRecommendationRequest) Get() *InitialGlobalBudgetRecommendationRequest {
	return v.value
}

func (v *NullableInitialGlobalBudgetRecommendationRequest) Set(val *InitialGlobalBudgetRecommendationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInitialGlobalBudgetRecommendationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInitialGlobalBudgetRecommendationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitialGlobalBudgetRecommendationRequest(val *InitialGlobalBudgetRecommendationRequest) *NullableInitialGlobalBudgetRecommendationRequest {
	return &NullableInitialGlobalBudgetRecommendationRequest{value: val, isSet: true}
}

func (v NullableInitialGlobalBudgetRecommendationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInitialGlobalBudgetRecommendationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
