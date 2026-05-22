package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SevenDaysMissedOpportunities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SevenDaysMissedOpportunities{}

// SevenDaysMissedOpportunities Missed Opportunities in the trailing seven days.
type SevenDaysMissedOpportunities struct {
	// Lower bound of the estimated Missed Sales. This will be in local currency.
	EstimatedMissedSalesLower *float64 `json:"estimatedMissedSalesLower,omitempty"`
	// Upper bound of the estimated Missed Sales. This will be in local currency.
	EstimatedMissedSalesUpper *float64 `json:"estimatedMissedSalesUpper,omitempty"`
	// End date of the Missed Opportunities date range (YYYY-MM-DD) in local time.
	EndDate *string `json:"endDate,omitempty" validate:"regexp=^20[1-9][0-9]-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$"`
	// Lower bound of the estimated Missed Impressions.
	EstimatedMissedImpressionsLower *float32 `json:"estimatedMissedImpressionsLower,omitempty"`
	// Lower bound of the estimated Missed Clicks.
	EstimatedMissedClicksLower *float32 `json:"estimatedMissedClicksLower,omitempty"`
	// Upper bound of the estimated Missed Clicks.
	EstimatedMissedClicksUpper *float32 `json:"estimatedMissedClicksUpper,omitempty"`
	// Upper bound of the estimated Missed Impressions.
	EstimatedMissedImpressionsUpper *float32 `json:"estimatedMissedImpressionsUpper,omitempty"`
	// Start date of the Missed Opportunities date range (YYYY-MM-DD) in local time.
	StartDate *string `json:"startDate,omitempty" validate:"regexp=^20[1-9][0-9]-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$"`
	// Percentage of time the campaign is active with a budget.
	PercentTimeInBudget *float64 `json:"percentTimeInBudget,omitempty"`
}

// NewSevenDaysMissedOpportunities instantiates a new SevenDaysMissedOpportunities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSevenDaysMissedOpportunities() *SevenDaysMissedOpportunities {
	this := SevenDaysMissedOpportunities{}
	return &this
}

// NewSevenDaysMissedOpportunitiesWithDefaults instantiates a new SevenDaysMissedOpportunities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSevenDaysMissedOpportunitiesWithDefaults() *SevenDaysMissedOpportunities {
	this := SevenDaysMissedOpportunities{}
	return &this
}

// GetEstimatedMissedSalesLower returns the EstimatedMissedSalesLower field value if set, zero value otherwise.
func (o *SevenDaysMissedOpportunities) GetEstimatedMissedSalesLower() float64 {
	if o == nil || IsNil(o.EstimatedMissedSalesLower) {
		var ret float64
		return ret
	}
	return *o.EstimatedMissedSalesLower
}

// GetEstimatedMissedSalesLowerOk returns a tuple with the EstimatedMissedSalesLower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SevenDaysMissedOpportunities) GetEstimatedMissedSalesLowerOk() (*float64, bool) {
	if o == nil || IsNil(o.EstimatedMissedSalesLower) {
		return nil, false
	}
	return o.EstimatedMissedSalesLower, true
}

// HasEstimatedMissedSalesLower returns a boolean if a field has been set.
func (o *SevenDaysMissedOpportunities) HasEstimatedMissedSalesLower() bool {
	if o != nil && !IsNil(o.EstimatedMissedSalesLower) {
		return true
	}

	return false
}

// SetEstimatedMissedSalesLower gets a reference to the given float64 and assigns it to the EstimatedMissedSalesLower field.
func (o *SevenDaysMissedOpportunities) SetEstimatedMissedSalesLower(v float64) {
	o.EstimatedMissedSalesLower = &v
}

// GetEstimatedMissedSalesUpper returns the EstimatedMissedSalesUpper field value if set, zero value otherwise.
func (o *SevenDaysMissedOpportunities) GetEstimatedMissedSalesUpper() float64 {
	if o == nil || IsNil(o.EstimatedMissedSalesUpper) {
		var ret float64
		return ret
	}
	return *o.EstimatedMissedSalesUpper
}

// GetEstimatedMissedSalesUpperOk returns a tuple with the EstimatedMissedSalesUpper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SevenDaysMissedOpportunities) GetEstimatedMissedSalesUpperOk() (*float64, bool) {
	if o == nil || IsNil(o.EstimatedMissedSalesUpper) {
		return nil, false
	}
	return o.EstimatedMissedSalesUpper, true
}

// HasEstimatedMissedSalesUpper returns a boolean if a field has been set.
func (o *SevenDaysMissedOpportunities) HasEstimatedMissedSalesUpper() bool {
	if o != nil && !IsNil(o.EstimatedMissedSalesUpper) {
		return true
	}

	return false
}

// SetEstimatedMissedSalesUpper gets a reference to the given float64 and assigns it to the EstimatedMissedSalesUpper field.
func (o *SevenDaysMissedOpportunities) SetEstimatedMissedSalesUpper(v float64) {
	o.EstimatedMissedSalesUpper = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SevenDaysMissedOpportunities) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SevenDaysMissedOpportunities) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SevenDaysMissedOpportunities) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *SevenDaysMissedOpportunities) SetEndDate(v string) {
	o.EndDate = &v
}

// GetEstimatedMissedImpressionsLower returns the EstimatedMissedImpressionsLower field value if set, zero value otherwise.
func (o *SevenDaysMissedOpportunities) GetEstimatedMissedImpressionsLower() float32 {
	if o == nil || IsNil(o.EstimatedMissedImpressionsLower) {
		var ret float32
		return ret
	}
	return *o.EstimatedMissedImpressionsLower
}

// GetEstimatedMissedImpressionsLowerOk returns a tuple with the EstimatedMissedImpressionsLower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SevenDaysMissedOpportunities) GetEstimatedMissedImpressionsLowerOk() (*float32, bool) {
	if o == nil || IsNil(o.EstimatedMissedImpressionsLower) {
		return nil, false
	}
	return o.EstimatedMissedImpressionsLower, true
}

// HasEstimatedMissedImpressionsLower returns a boolean if a field has been set.
func (o *SevenDaysMissedOpportunities) HasEstimatedMissedImpressionsLower() bool {
	if o != nil && !IsNil(o.EstimatedMissedImpressionsLower) {
		return true
	}

	return false
}

// SetEstimatedMissedImpressionsLower gets a reference to the given float32 and assigns it to the EstimatedMissedImpressionsLower field.
func (o *SevenDaysMissedOpportunities) SetEstimatedMissedImpressionsLower(v float32) {
	o.EstimatedMissedImpressionsLower = &v
}

// GetEstimatedMissedClicksLower returns the EstimatedMissedClicksLower field value if set, zero value otherwise.
func (o *SevenDaysMissedOpportunities) GetEstimatedMissedClicksLower() float32 {
	if o == nil || IsNil(o.EstimatedMissedClicksLower) {
		var ret float32
		return ret
	}
	return *o.EstimatedMissedClicksLower
}

// GetEstimatedMissedClicksLowerOk returns a tuple with the EstimatedMissedClicksLower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SevenDaysMissedOpportunities) GetEstimatedMissedClicksLowerOk() (*float32, bool) {
	if o == nil || IsNil(o.EstimatedMissedClicksLower) {
		return nil, false
	}
	return o.EstimatedMissedClicksLower, true
}

// HasEstimatedMissedClicksLower returns a boolean if a field has been set.
func (o *SevenDaysMissedOpportunities) HasEstimatedMissedClicksLower() bool {
	if o != nil && !IsNil(o.EstimatedMissedClicksLower) {
		return true
	}

	return false
}

// SetEstimatedMissedClicksLower gets a reference to the given float32 and assigns it to the EstimatedMissedClicksLower field.
func (o *SevenDaysMissedOpportunities) SetEstimatedMissedClicksLower(v float32) {
	o.EstimatedMissedClicksLower = &v
}

// GetEstimatedMissedClicksUpper returns the EstimatedMissedClicksUpper field value if set, zero value otherwise.
func (o *SevenDaysMissedOpportunities) GetEstimatedMissedClicksUpper() float32 {
	if o == nil || IsNil(o.EstimatedMissedClicksUpper) {
		var ret float32
		return ret
	}
	return *o.EstimatedMissedClicksUpper
}

// GetEstimatedMissedClicksUpperOk returns a tuple with the EstimatedMissedClicksUpper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SevenDaysMissedOpportunities) GetEstimatedMissedClicksUpperOk() (*float32, bool) {
	if o == nil || IsNil(o.EstimatedMissedClicksUpper) {
		return nil, false
	}
	return o.EstimatedMissedClicksUpper, true
}

// HasEstimatedMissedClicksUpper returns a boolean if a field has been set.
func (o *SevenDaysMissedOpportunities) HasEstimatedMissedClicksUpper() bool {
	if o != nil && !IsNil(o.EstimatedMissedClicksUpper) {
		return true
	}

	return false
}

// SetEstimatedMissedClicksUpper gets a reference to the given float32 and assigns it to the EstimatedMissedClicksUpper field.
func (o *SevenDaysMissedOpportunities) SetEstimatedMissedClicksUpper(v float32) {
	o.EstimatedMissedClicksUpper = &v
}

// GetEstimatedMissedImpressionsUpper returns the EstimatedMissedImpressionsUpper field value if set, zero value otherwise.
func (o *SevenDaysMissedOpportunities) GetEstimatedMissedImpressionsUpper() float32 {
	if o == nil || IsNil(o.EstimatedMissedImpressionsUpper) {
		var ret float32
		return ret
	}
	return *o.EstimatedMissedImpressionsUpper
}

// GetEstimatedMissedImpressionsUpperOk returns a tuple with the EstimatedMissedImpressionsUpper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SevenDaysMissedOpportunities) GetEstimatedMissedImpressionsUpperOk() (*float32, bool) {
	if o == nil || IsNil(o.EstimatedMissedImpressionsUpper) {
		return nil, false
	}
	return o.EstimatedMissedImpressionsUpper, true
}

// HasEstimatedMissedImpressionsUpper returns a boolean if a field has been set.
func (o *SevenDaysMissedOpportunities) HasEstimatedMissedImpressionsUpper() bool {
	if o != nil && !IsNil(o.EstimatedMissedImpressionsUpper) {
		return true
	}

	return false
}

// SetEstimatedMissedImpressionsUpper gets a reference to the given float32 and assigns it to the EstimatedMissedImpressionsUpper field.
func (o *SevenDaysMissedOpportunities) SetEstimatedMissedImpressionsUpper(v float32) {
	o.EstimatedMissedImpressionsUpper = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SevenDaysMissedOpportunities) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SevenDaysMissedOpportunities) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SevenDaysMissedOpportunities) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SevenDaysMissedOpportunities) SetStartDate(v string) {
	o.StartDate = &v
}

// GetPercentTimeInBudget returns the PercentTimeInBudget field value if set, zero value otherwise.
func (o *SevenDaysMissedOpportunities) GetPercentTimeInBudget() float64 {
	if o == nil || IsNil(o.PercentTimeInBudget) {
		var ret float64
		return ret
	}
	return *o.PercentTimeInBudget
}

// GetPercentTimeInBudgetOk returns a tuple with the PercentTimeInBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SevenDaysMissedOpportunities) GetPercentTimeInBudgetOk() (*float64, bool) {
	if o == nil || IsNil(o.PercentTimeInBudget) {
		return nil, false
	}
	return o.PercentTimeInBudget, true
}

// HasPercentTimeInBudget returns a boolean if a field has been set.
func (o *SevenDaysMissedOpportunities) HasPercentTimeInBudget() bool {
	if o != nil && !IsNil(o.PercentTimeInBudget) {
		return true
	}

	return false
}

// SetPercentTimeInBudget gets a reference to the given float64 and assigns it to the PercentTimeInBudget field.
func (o *SevenDaysMissedOpportunities) SetPercentTimeInBudget(v float64) {
	o.PercentTimeInBudget = &v
}

func (o SevenDaysMissedOpportunities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EstimatedMissedSalesLower) {
		toSerialize["estimatedMissedSalesLower"] = o.EstimatedMissedSalesLower
	}
	if !IsNil(o.EstimatedMissedSalesUpper) {
		toSerialize["estimatedMissedSalesUpper"] = o.EstimatedMissedSalesUpper
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.EstimatedMissedImpressionsLower) {
		toSerialize["estimatedMissedImpressionsLower"] = o.EstimatedMissedImpressionsLower
	}
	if !IsNil(o.EstimatedMissedClicksLower) {
		toSerialize["estimatedMissedClicksLower"] = o.EstimatedMissedClicksLower
	}
	if !IsNil(o.EstimatedMissedClicksUpper) {
		toSerialize["estimatedMissedClicksUpper"] = o.EstimatedMissedClicksUpper
	}
	if !IsNil(o.EstimatedMissedImpressionsUpper) {
		toSerialize["estimatedMissedImpressionsUpper"] = o.EstimatedMissedImpressionsUpper
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.PercentTimeInBudget) {
		toSerialize["percentTimeInBudget"] = o.PercentTimeInBudget
	}
	return toSerialize, nil
}

type NullableSevenDaysMissedOpportunities struct {
	value *SevenDaysMissedOpportunities
	isSet bool
}

func (v NullableSevenDaysMissedOpportunities) Get() *SevenDaysMissedOpportunities {
	return v.value
}

func (v *NullableSevenDaysMissedOpportunities) Set(val *SevenDaysMissedOpportunities) {
	v.value = val
	v.isSet = true
}

func (v NullableSevenDaysMissedOpportunities) IsSet() bool {
	return v.isSet
}

func (v *NullableSevenDaysMissedOpportunities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSevenDaysMissedOpportunities(val *SevenDaysMissedOpportunities) *NullableSevenDaysMissedOpportunities {
	return &NullableSevenDaysMissedOpportunities{value: val, isSet: true}
}

func (v NullableSevenDaysMissedOpportunities) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSevenDaysMissedOpportunities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
