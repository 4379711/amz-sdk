package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CountryBudgetRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountryBudgetRecommendation{}

// CountryBudgetRecommendation struct for CountryBudgetRecommendation
type CountryBudgetRecommendation struct {
	// recommended budget for the country in country's local currency.
	SuggestedBudget              float32                      `json:"suggestedBudget"`
	SevenDaysMissedOpportunities SevenDaysMissedOpportunities `json:"sevenDaysMissedOpportunities"`
}

type _CountryBudgetRecommendation CountryBudgetRecommendation

// NewCountryBudgetRecommendation instantiates a new CountryBudgetRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryBudgetRecommendation(suggestedBudget float32, sevenDaysMissedOpportunities SevenDaysMissedOpportunities) *CountryBudgetRecommendation {
	this := CountryBudgetRecommendation{}
	this.SuggestedBudget = suggestedBudget
	this.SevenDaysMissedOpportunities = sevenDaysMissedOpportunities
	return &this
}

// NewCountryBudgetRecommendationWithDefaults instantiates a new CountryBudgetRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryBudgetRecommendationWithDefaults() *CountryBudgetRecommendation {
	this := CountryBudgetRecommendation{}
	return &this
}

// GetSuggestedBudget returns the SuggestedBudget field value
func (o *CountryBudgetRecommendation) GetSuggestedBudget() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SuggestedBudget
}

// GetSuggestedBudgetOk returns a tuple with the SuggestedBudget field value
// and a boolean to check if the value has been set.
func (o *CountryBudgetRecommendation) GetSuggestedBudgetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuggestedBudget, true
}

// SetSuggestedBudget sets field value
func (o *CountryBudgetRecommendation) SetSuggestedBudget(v float32) {
	o.SuggestedBudget = v
}

// GetSevenDaysMissedOpportunities returns the SevenDaysMissedOpportunities field value
func (o *CountryBudgetRecommendation) GetSevenDaysMissedOpportunities() SevenDaysMissedOpportunities {
	if o == nil {
		var ret SevenDaysMissedOpportunities
		return ret
	}

	return o.SevenDaysMissedOpportunities
}

// GetSevenDaysMissedOpportunitiesOk returns a tuple with the SevenDaysMissedOpportunities field value
// and a boolean to check if the value has been set.
func (o *CountryBudgetRecommendation) GetSevenDaysMissedOpportunitiesOk() (*SevenDaysMissedOpportunities, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SevenDaysMissedOpportunities, true
}

// SetSevenDaysMissedOpportunities sets field value
func (o *CountryBudgetRecommendation) SetSevenDaysMissedOpportunities(v SevenDaysMissedOpportunities) {
	o.SevenDaysMissedOpportunities = v
}

func (o CountryBudgetRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["suggestedBudget"] = o.SuggestedBudget
	toSerialize["sevenDaysMissedOpportunities"] = o.SevenDaysMissedOpportunities
	return toSerialize, nil
}

type NullableCountryBudgetRecommendation struct {
	value *CountryBudgetRecommendation
	isSet bool
}

func (v NullableCountryBudgetRecommendation) Get() *CountryBudgetRecommendation {
	return v.value
}

func (v *NullableCountryBudgetRecommendation) Set(val *CountryBudgetRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryBudgetRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryBudgetRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryBudgetRecommendation(val *CountryBudgetRecommendation) *NullableCountryBudgetRecommendation {
	return &NullableCountryBudgetRecommendation{value: val, isSet: true}
}

func (v NullableCountryBudgetRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCountryBudgetRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
