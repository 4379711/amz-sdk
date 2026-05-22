package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SevenDaysEstimatedOpportunities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SevenDaysEstimatedOpportunities{}

// SevenDaysEstimatedOpportunities struct for SevenDaysEstimatedOpportunities
type SevenDaysEstimatedOpportunities struct {
	// Lower bound of the estimated incremental clicks that could be gained if all optimizations are made.
	EstimatedIncrementalClicksLower *int32 `json:"estimatedIncrementalClicksLower,omitempty"`
	// Upper bound of the estimated incremental clicks that could be gained if all optimizations are made.
	EstimatedIncrementalClicksUpper *int32 `json:"estimatedIncrementalClicksUpper,omitempty"`
	// End date of the opportunities date range in YYYY-MM-DDTHH:mm:ssZ format.
	EndDate *string `json:"endDate,omitempty"`
	// Start date of the opportunities date range in YYYY-MM-DDTHH:mm:ssZ format.
	StartDate *string `json:"startDate,omitempty"`
}

// NewSevenDaysEstimatedOpportunities instantiates a new SevenDaysEstimatedOpportunities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSevenDaysEstimatedOpportunities() *SevenDaysEstimatedOpportunities {
	this := SevenDaysEstimatedOpportunities{}
	return &this
}

// NewSevenDaysEstimatedOpportunitiesWithDefaults instantiates a new SevenDaysEstimatedOpportunities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSevenDaysEstimatedOpportunitiesWithDefaults() *SevenDaysEstimatedOpportunities {
	this := SevenDaysEstimatedOpportunities{}
	return &this
}

// GetEstimatedIncrementalClicksLower returns the EstimatedIncrementalClicksLower field value if set, zero value otherwise.
func (o *SevenDaysEstimatedOpportunities) GetEstimatedIncrementalClicksLower() int32 {
	if o == nil || IsNil(o.EstimatedIncrementalClicksLower) {
		var ret int32
		return ret
	}
	return *o.EstimatedIncrementalClicksLower
}

// GetEstimatedIncrementalClicksLowerOk returns a tuple with the EstimatedIncrementalClicksLower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SevenDaysEstimatedOpportunities) GetEstimatedIncrementalClicksLowerOk() (*int32, bool) {
	if o == nil || IsNil(o.EstimatedIncrementalClicksLower) {
		return nil, false
	}
	return o.EstimatedIncrementalClicksLower, true
}

// HasEstimatedIncrementalClicksLower returns a boolean if a field has been set.
func (o *SevenDaysEstimatedOpportunities) HasEstimatedIncrementalClicksLower() bool {
	if o != nil && !IsNil(o.EstimatedIncrementalClicksLower) {
		return true
	}

	return false
}

// SetEstimatedIncrementalClicksLower gets a reference to the given int32 and assigns it to the EstimatedIncrementalClicksLower field.
func (o *SevenDaysEstimatedOpportunities) SetEstimatedIncrementalClicksLower(v int32) {
	o.EstimatedIncrementalClicksLower = &v
}

// GetEstimatedIncrementalClicksUpper returns the EstimatedIncrementalClicksUpper field value if set, zero value otherwise.
func (o *SevenDaysEstimatedOpportunities) GetEstimatedIncrementalClicksUpper() int32 {
	if o == nil || IsNil(o.EstimatedIncrementalClicksUpper) {
		var ret int32
		return ret
	}
	return *o.EstimatedIncrementalClicksUpper
}

// GetEstimatedIncrementalClicksUpperOk returns a tuple with the EstimatedIncrementalClicksUpper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SevenDaysEstimatedOpportunities) GetEstimatedIncrementalClicksUpperOk() (*int32, bool) {
	if o == nil || IsNil(o.EstimatedIncrementalClicksUpper) {
		return nil, false
	}
	return o.EstimatedIncrementalClicksUpper, true
}

// HasEstimatedIncrementalClicksUpper returns a boolean if a field has been set.
func (o *SevenDaysEstimatedOpportunities) HasEstimatedIncrementalClicksUpper() bool {
	if o != nil && !IsNil(o.EstimatedIncrementalClicksUpper) {
		return true
	}

	return false
}

// SetEstimatedIncrementalClicksUpper gets a reference to the given int32 and assigns it to the EstimatedIncrementalClicksUpper field.
func (o *SevenDaysEstimatedOpportunities) SetEstimatedIncrementalClicksUpper(v int32) {
	o.EstimatedIncrementalClicksUpper = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SevenDaysEstimatedOpportunities) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SevenDaysEstimatedOpportunities) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SevenDaysEstimatedOpportunities) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *SevenDaysEstimatedOpportunities) SetEndDate(v string) {
	o.EndDate = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SevenDaysEstimatedOpportunities) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SevenDaysEstimatedOpportunities) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SevenDaysEstimatedOpportunities) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SevenDaysEstimatedOpportunities) SetStartDate(v string) {
	o.StartDate = &v
}

func (o SevenDaysEstimatedOpportunities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EstimatedIncrementalClicksLower) {
		toSerialize["estimatedIncrementalClicksLower"] = o.EstimatedIncrementalClicksLower
	}
	if !IsNil(o.EstimatedIncrementalClicksUpper) {
		toSerialize["estimatedIncrementalClicksUpper"] = o.EstimatedIncrementalClicksUpper
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	return toSerialize, nil
}

type NullableSevenDaysEstimatedOpportunities struct {
	value *SevenDaysEstimatedOpportunities
	isSet bool
}

func (v NullableSevenDaysEstimatedOpportunities) Get() *SevenDaysEstimatedOpportunities {
	return v.value
}

func (v *NullableSevenDaysEstimatedOpportunities) Set(val *SevenDaysEstimatedOpportunities) {
	v.value = val
	v.isSet = true
}

func (v NullableSevenDaysEstimatedOpportunities) IsSet() bool {
	return v.isSet
}

func (v *NullableSevenDaysEstimatedOpportunities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSevenDaysEstimatedOpportunities(val *SevenDaysEstimatedOpportunities) *NullableSevenDaysEstimatedOpportunities {
	return &NullableSevenDaysEstimatedOpportunities{value: val, isSet: true}
}

func (v NullableSevenDaysEstimatedOpportunities) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSevenDaysEstimatedOpportunities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
