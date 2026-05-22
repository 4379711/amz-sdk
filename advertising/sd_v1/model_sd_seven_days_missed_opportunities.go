package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDSevenDaysMissedOpportunities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDSevenDaysMissedOpportunities{}

// SDSevenDaysMissedOpportunities struct for SDSevenDaysMissedOpportunities
type SDSevenDaysMissedOpportunities struct {
	// Start date of the missed opportunities date range (YYYY-MM-DD).
	StartDate *string `json:"startDate,omitempty"`
	// End date of the missed opportunities date range (YYYY-MM-DD).
	EndDate *string `json:"endDate,omitempty"`
	// Percentage of time the campaign is active with a budget.
	PercentTimeInBudget *float32 `json:"percentTimeInBudget,omitempty"`
	// Lower bound of the estimated missed sales. This will be in local currency.
	EstimatedMissedSalesLower *float32 `json:"estimatedMissedSalesLower,omitempty"`
	// Upper bound of the estimated missed sales. This will be in local currency.
	EstimatedMissedSalesUpper *float32 `json:"estimatedMissedSalesUpper,omitempty"`
	// Lower bound of the estimated missed clicks.
	EstimatedMissedClicksLower *int32 `json:"estimatedMissedClicksLower,omitempty"`
	// Upper bound of the estimated missed clicks.
	EstimatedMissedClicksUpper *int32 `json:"estimatedMissedClicksUpper,omitempty"`
	// Lower bound of the estimated missed impressions.
	EstimatedMissedImpressionsLower *int32 `json:"estimatedMissedImpressionsLower,omitempty"`
	// Upper bound of the estimated missed impressions.
	EstimatedMissedImpressionsUpper *int32 `json:"estimatedMissedImpressionsUpper,omitempty"`
	// Lower bound of the estimated missed viewable impressions for vCPM campaigns.
	EstimatedMissedViewableImpressionsLower *int32 `json:"estimatedMissedViewableImpressionsLower,omitempty"`
	// Upper bound of the estimated missed viewable impressions for vCPM campaigns.
	EstimatedMissedViewableImpressionsUpper *int32 `json:"estimatedMissedViewableImpressionsUpper,omitempty"`
}

// NewSDSevenDaysMissedOpportunities instantiates a new SDSevenDaysMissedOpportunities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDSevenDaysMissedOpportunities() *SDSevenDaysMissedOpportunities {
	this := SDSevenDaysMissedOpportunities{}
	return &this
}

// NewSDSevenDaysMissedOpportunitiesWithDefaults instantiates a new SDSevenDaysMissedOpportunities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDSevenDaysMissedOpportunitiesWithDefaults() *SDSevenDaysMissedOpportunities {
	this := SDSevenDaysMissedOpportunities{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SDSevenDaysMissedOpportunities) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDSevenDaysMissedOpportunities) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SDSevenDaysMissedOpportunities) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SDSevenDaysMissedOpportunities) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SDSevenDaysMissedOpportunities) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDSevenDaysMissedOpportunities) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SDSevenDaysMissedOpportunities) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *SDSevenDaysMissedOpportunities) SetEndDate(v string) {
	o.EndDate = &v
}

// GetPercentTimeInBudget returns the PercentTimeInBudget field value if set, zero value otherwise.
func (o *SDSevenDaysMissedOpportunities) GetPercentTimeInBudget() float32 {
	if o == nil || IsNil(o.PercentTimeInBudget) {
		var ret float32
		return ret
	}
	return *o.PercentTimeInBudget
}

// GetPercentTimeInBudgetOk returns a tuple with the PercentTimeInBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDSevenDaysMissedOpportunities) GetPercentTimeInBudgetOk() (*float32, bool) {
	if o == nil || IsNil(o.PercentTimeInBudget) {
		return nil, false
	}
	return o.PercentTimeInBudget, true
}

// HasPercentTimeInBudget returns a boolean if a field has been set.
func (o *SDSevenDaysMissedOpportunities) HasPercentTimeInBudget() bool {
	if o != nil && !IsNil(o.PercentTimeInBudget) {
		return true
	}

	return false
}

// SetPercentTimeInBudget gets a reference to the given float32 and assigns it to the PercentTimeInBudget field.
func (o *SDSevenDaysMissedOpportunities) SetPercentTimeInBudget(v float32) {
	o.PercentTimeInBudget = &v
}

// GetEstimatedMissedSalesLower returns the EstimatedMissedSalesLower field value if set, zero value otherwise.
func (o *SDSevenDaysMissedOpportunities) GetEstimatedMissedSalesLower() float32 {
	if o == nil || IsNil(o.EstimatedMissedSalesLower) {
		var ret float32
		return ret
	}
	return *o.EstimatedMissedSalesLower
}

// GetEstimatedMissedSalesLowerOk returns a tuple with the EstimatedMissedSalesLower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDSevenDaysMissedOpportunities) GetEstimatedMissedSalesLowerOk() (*float32, bool) {
	if o == nil || IsNil(o.EstimatedMissedSalesLower) {
		return nil, false
	}
	return o.EstimatedMissedSalesLower, true
}

// HasEstimatedMissedSalesLower returns a boolean if a field has been set.
func (o *SDSevenDaysMissedOpportunities) HasEstimatedMissedSalesLower() bool {
	if o != nil && !IsNil(o.EstimatedMissedSalesLower) {
		return true
	}

	return false
}

// SetEstimatedMissedSalesLower gets a reference to the given float32 and assigns it to the EstimatedMissedSalesLower field.
func (o *SDSevenDaysMissedOpportunities) SetEstimatedMissedSalesLower(v float32) {
	o.EstimatedMissedSalesLower = &v
}

// GetEstimatedMissedSalesUpper returns the EstimatedMissedSalesUpper field value if set, zero value otherwise.
func (o *SDSevenDaysMissedOpportunities) GetEstimatedMissedSalesUpper() float32 {
	if o == nil || IsNil(o.EstimatedMissedSalesUpper) {
		var ret float32
		return ret
	}
	return *o.EstimatedMissedSalesUpper
}

// GetEstimatedMissedSalesUpperOk returns a tuple with the EstimatedMissedSalesUpper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDSevenDaysMissedOpportunities) GetEstimatedMissedSalesUpperOk() (*float32, bool) {
	if o == nil || IsNil(o.EstimatedMissedSalesUpper) {
		return nil, false
	}
	return o.EstimatedMissedSalesUpper, true
}

// HasEstimatedMissedSalesUpper returns a boolean if a field has been set.
func (o *SDSevenDaysMissedOpportunities) HasEstimatedMissedSalesUpper() bool {
	if o != nil && !IsNil(o.EstimatedMissedSalesUpper) {
		return true
	}

	return false
}

// SetEstimatedMissedSalesUpper gets a reference to the given float32 and assigns it to the EstimatedMissedSalesUpper field.
func (o *SDSevenDaysMissedOpportunities) SetEstimatedMissedSalesUpper(v float32) {
	o.EstimatedMissedSalesUpper = &v
}

// GetEstimatedMissedClicksLower returns the EstimatedMissedClicksLower field value if set, zero value otherwise.
func (o *SDSevenDaysMissedOpportunities) GetEstimatedMissedClicksLower() int32 {
	if o == nil || IsNil(o.EstimatedMissedClicksLower) {
		var ret int32
		return ret
	}
	return *o.EstimatedMissedClicksLower
}

// GetEstimatedMissedClicksLowerOk returns a tuple with the EstimatedMissedClicksLower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDSevenDaysMissedOpportunities) GetEstimatedMissedClicksLowerOk() (*int32, bool) {
	if o == nil || IsNil(o.EstimatedMissedClicksLower) {
		return nil, false
	}
	return o.EstimatedMissedClicksLower, true
}

// HasEstimatedMissedClicksLower returns a boolean if a field has been set.
func (o *SDSevenDaysMissedOpportunities) HasEstimatedMissedClicksLower() bool {
	if o != nil && !IsNil(o.EstimatedMissedClicksLower) {
		return true
	}

	return false
}

// SetEstimatedMissedClicksLower gets a reference to the given int32 and assigns it to the EstimatedMissedClicksLower field.
func (o *SDSevenDaysMissedOpportunities) SetEstimatedMissedClicksLower(v int32) {
	o.EstimatedMissedClicksLower = &v
}

// GetEstimatedMissedClicksUpper returns the EstimatedMissedClicksUpper field value if set, zero value otherwise.
func (o *SDSevenDaysMissedOpportunities) GetEstimatedMissedClicksUpper() int32 {
	if o == nil || IsNil(o.EstimatedMissedClicksUpper) {
		var ret int32
		return ret
	}
	return *o.EstimatedMissedClicksUpper
}

// GetEstimatedMissedClicksUpperOk returns a tuple with the EstimatedMissedClicksUpper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDSevenDaysMissedOpportunities) GetEstimatedMissedClicksUpperOk() (*int32, bool) {
	if o == nil || IsNil(o.EstimatedMissedClicksUpper) {
		return nil, false
	}
	return o.EstimatedMissedClicksUpper, true
}

// HasEstimatedMissedClicksUpper returns a boolean if a field has been set.
func (o *SDSevenDaysMissedOpportunities) HasEstimatedMissedClicksUpper() bool {
	if o != nil && !IsNil(o.EstimatedMissedClicksUpper) {
		return true
	}

	return false
}

// SetEstimatedMissedClicksUpper gets a reference to the given int32 and assigns it to the EstimatedMissedClicksUpper field.
func (o *SDSevenDaysMissedOpportunities) SetEstimatedMissedClicksUpper(v int32) {
	o.EstimatedMissedClicksUpper = &v
}

// GetEstimatedMissedImpressionsLower returns the EstimatedMissedImpressionsLower field value if set, zero value otherwise.
func (o *SDSevenDaysMissedOpportunities) GetEstimatedMissedImpressionsLower() int32 {
	if o == nil || IsNil(o.EstimatedMissedImpressionsLower) {
		var ret int32
		return ret
	}
	return *o.EstimatedMissedImpressionsLower
}

// GetEstimatedMissedImpressionsLowerOk returns a tuple with the EstimatedMissedImpressionsLower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDSevenDaysMissedOpportunities) GetEstimatedMissedImpressionsLowerOk() (*int32, bool) {
	if o == nil || IsNil(o.EstimatedMissedImpressionsLower) {
		return nil, false
	}
	return o.EstimatedMissedImpressionsLower, true
}

// HasEstimatedMissedImpressionsLower returns a boolean if a field has been set.
func (o *SDSevenDaysMissedOpportunities) HasEstimatedMissedImpressionsLower() bool {
	if o != nil && !IsNil(o.EstimatedMissedImpressionsLower) {
		return true
	}

	return false
}

// SetEstimatedMissedImpressionsLower gets a reference to the given int32 and assigns it to the EstimatedMissedImpressionsLower field.
func (o *SDSevenDaysMissedOpportunities) SetEstimatedMissedImpressionsLower(v int32) {
	o.EstimatedMissedImpressionsLower = &v
}

// GetEstimatedMissedImpressionsUpper returns the EstimatedMissedImpressionsUpper field value if set, zero value otherwise.
func (o *SDSevenDaysMissedOpportunities) GetEstimatedMissedImpressionsUpper() int32 {
	if o == nil || IsNil(o.EstimatedMissedImpressionsUpper) {
		var ret int32
		return ret
	}
	return *o.EstimatedMissedImpressionsUpper
}

// GetEstimatedMissedImpressionsUpperOk returns a tuple with the EstimatedMissedImpressionsUpper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDSevenDaysMissedOpportunities) GetEstimatedMissedImpressionsUpperOk() (*int32, bool) {
	if o == nil || IsNil(o.EstimatedMissedImpressionsUpper) {
		return nil, false
	}
	return o.EstimatedMissedImpressionsUpper, true
}

// HasEstimatedMissedImpressionsUpper returns a boolean if a field has been set.
func (o *SDSevenDaysMissedOpportunities) HasEstimatedMissedImpressionsUpper() bool {
	if o != nil && !IsNil(o.EstimatedMissedImpressionsUpper) {
		return true
	}

	return false
}

// SetEstimatedMissedImpressionsUpper gets a reference to the given int32 and assigns it to the EstimatedMissedImpressionsUpper field.
func (o *SDSevenDaysMissedOpportunities) SetEstimatedMissedImpressionsUpper(v int32) {
	o.EstimatedMissedImpressionsUpper = &v
}

// GetEstimatedMissedViewableImpressionsLower returns the EstimatedMissedViewableImpressionsLower field value if set, zero value otherwise.
func (o *SDSevenDaysMissedOpportunities) GetEstimatedMissedViewableImpressionsLower() int32 {
	if o == nil || IsNil(o.EstimatedMissedViewableImpressionsLower) {
		var ret int32
		return ret
	}
	return *o.EstimatedMissedViewableImpressionsLower
}

// GetEstimatedMissedViewableImpressionsLowerOk returns a tuple with the EstimatedMissedViewableImpressionsLower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDSevenDaysMissedOpportunities) GetEstimatedMissedViewableImpressionsLowerOk() (*int32, bool) {
	if o == nil || IsNil(o.EstimatedMissedViewableImpressionsLower) {
		return nil, false
	}
	return o.EstimatedMissedViewableImpressionsLower, true
}

// HasEstimatedMissedViewableImpressionsLower returns a boolean if a field has been set.
func (o *SDSevenDaysMissedOpportunities) HasEstimatedMissedViewableImpressionsLower() bool {
	if o != nil && !IsNil(o.EstimatedMissedViewableImpressionsLower) {
		return true
	}

	return false
}

// SetEstimatedMissedViewableImpressionsLower gets a reference to the given int32 and assigns it to the EstimatedMissedViewableImpressionsLower field.
func (o *SDSevenDaysMissedOpportunities) SetEstimatedMissedViewableImpressionsLower(v int32) {
	o.EstimatedMissedViewableImpressionsLower = &v
}

// GetEstimatedMissedViewableImpressionsUpper returns the EstimatedMissedViewableImpressionsUpper field value if set, zero value otherwise.
func (o *SDSevenDaysMissedOpportunities) GetEstimatedMissedViewableImpressionsUpper() int32 {
	if o == nil || IsNil(o.EstimatedMissedViewableImpressionsUpper) {
		var ret int32
		return ret
	}
	return *o.EstimatedMissedViewableImpressionsUpper
}

// GetEstimatedMissedViewableImpressionsUpperOk returns a tuple with the EstimatedMissedViewableImpressionsUpper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDSevenDaysMissedOpportunities) GetEstimatedMissedViewableImpressionsUpperOk() (*int32, bool) {
	if o == nil || IsNil(o.EstimatedMissedViewableImpressionsUpper) {
		return nil, false
	}
	return o.EstimatedMissedViewableImpressionsUpper, true
}

// HasEstimatedMissedViewableImpressionsUpper returns a boolean if a field has been set.
func (o *SDSevenDaysMissedOpportunities) HasEstimatedMissedViewableImpressionsUpper() bool {
	if o != nil && !IsNil(o.EstimatedMissedViewableImpressionsUpper) {
		return true
	}

	return false
}

// SetEstimatedMissedViewableImpressionsUpper gets a reference to the given int32 and assigns it to the EstimatedMissedViewableImpressionsUpper field.
func (o *SDSevenDaysMissedOpportunities) SetEstimatedMissedViewableImpressionsUpper(v int32) {
	o.EstimatedMissedViewableImpressionsUpper = &v
}

func (o SDSevenDaysMissedOpportunities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.PercentTimeInBudget) {
		toSerialize["percentTimeInBudget"] = o.PercentTimeInBudget
	}
	if !IsNil(o.EstimatedMissedSalesLower) {
		toSerialize["estimatedMissedSalesLower"] = o.EstimatedMissedSalesLower
	}
	if !IsNil(o.EstimatedMissedSalesUpper) {
		toSerialize["estimatedMissedSalesUpper"] = o.EstimatedMissedSalesUpper
	}
	if !IsNil(o.EstimatedMissedClicksLower) {
		toSerialize["estimatedMissedClicksLower"] = o.EstimatedMissedClicksLower
	}
	if !IsNil(o.EstimatedMissedClicksUpper) {
		toSerialize["estimatedMissedClicksUpper"] = o.EstimatedMissedClicksUpper
	}
	if !IsNil(o.EstimatedMissedImpressionsLower) {
		toSerialize["estimatedMissedImpressionsLower"] = o.EstimatedMissedImpressionsLower
	}
	if !IsNil(o.EstimatedMissedImpressionsUpper) {
		toSerialize["estimatedMissedImpressionsUpper"] = o.EstimatedMissedImpressionsUpper
	}
	if !IsNil(o.EstimatedMissedViewableImpressionsLower) {
		toSerialize["estimatedMissedViewableImpressionsLower"] = o.EstimatedMissedViewableImpressionsLower
	}
	if !IsNil(o.EstimatedMissedViewableImpressionsUpper) {
		toSerialize["estimatedMissedViewableImpressionsUpper"] = o.EstimatedMissedViewableImpressionsUpper
	}
	return toSerialize, nil
}

type NullableSDSevenDaysMissedOpportunities struct {
	value *SDSevenDaysMissedOpportunities
	isSet bool
}

func (v NullableSDSevenDaysMissedOpportunities) Get() *SDSevenDaysMissedOpportunities {
	return v.value
}

func (v *NullableSDSevenDaysMissedOpportunities) Set(val *SDSevenDaysMissedOpportunities) {
	v.value = val
	v.isSet = true
}

func (v NullableSDSevenDaysMissedOpportunities) IsSet() bool {
	return v.isSet
}

func (v *NullableSDSevenDaysMissedOpportunities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDSevenDaysMissedOpportunities(val *SDSevenDaysMissedOpportunities) *NullableSDSevenDaysMissedOpportunities {
	return &NullableSDSevenDaysMissedOpportunities{value: val, isSet: true}
}

func (v NullableSDSevenDaysMissedOpportunities) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDSevenDaysMissedOpportunities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
