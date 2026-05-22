package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalBudgetRulesRecommendationEventResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalBudgetRulesRecommendationEventResponse{}

// GlobalBudgetRulesRecommendationEventResponse Globalised Special events with date range and suggested budget increase for existing campaigns.
type GlobalBudgetRulesRecommendationEventResponse struct {
	// A list of recommended special events with date range and suggested budget increase for each country targeted by the campaign.
	CountryBudgetRuleEventsRecommendations map[string][]SPBudgetRulesRecommendationEvent `json:"countryBudgetRuleEventsRecommendations"`
	Errors                                 []GlobalBudgetRulesRecommendationError        `json:"errors,omitempty"`
}

type _GlobalBudgetRulesRecommendationEventResponse GlobalBudgetRulesRecommendationEventResponse

// NewGlobalBudgetRulesRecommendationEventResponse instantiates a new GlobalBudgetRulesRecommendationEventResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalBudgetRulesRecommendationEventResponse(countryBudgetRuleEventsRecommendations map[string][]SPBudgetRulesRecommendationEvent) *GlobalBudgetRulesRecommendationEventResponse {
	this := GlobalBudgetRulesRecommendationEventResponse{}
	this.CountryBudgetRuleEventsRecommendations = countryBudgetRuleEventsRecommendations
	return &this
}

// NewGlobalBudgetRulesRecommendationEventResponseWithDefaults instantiates a new GlobalBudgetRulesRecommendationEventResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalBudgetRulesRecommendationEventResponseWithDefaults() *GlobalBudgetRulesRecommendationEventResponse {
	this := GlobalBudgetRulesRecommendationEventResponse{}
	return &this
}

// GetCountryBudgetRuleEventsRecommendations returns the CountryBudgetRuleEventsRecommendations field value
func (o *GlobalBudgetRulesRecommendationEventResponse) GetCountryBudgetRuleEventsRecommendations() map[string][]SPBudgetRulesRecommendationEvent {
	if o == nil {
		var ret map[string][]SPBudgetRulesRecommendationEvent
		return ret
	}

	return o.CountryBudgetRuleEventsRecommendations
}

// GetCountryBudgetRuleEventsRecommendationsOk returns a tuple with the CountryBudgetRuleEventsRecommendations field value
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRulesRecommendationEventResponse) GetCountryBudgetRuleEventsRecommendationsOk() (*map[string][]SPBudgetRulesRecommendationEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryBudgetRuleEventsRecommendations, true
}

// SetCountryBudgetRuleEventsRecommendations sets field value
func (o *GlobalBudgetRulesRecommendationEventResponse) SetCountryBudgetRuleEventsRecommendations(v map[string][]SPBudgetRulesRecommendationEvent) {
	o.CountryBudgetRuleEventsRecommendations = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GlobalBudgetRulesRecommendationEventResponse) GetErrors() []GlobalBudgetRulesRecommendationError {
	if o == nil || IsNil(o.Errors) {
		var ret []GlobalBudgetRulesRecommendationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRulesRecommendationEventResponse) GetErrorsOk() ([]GlobalBudgetRulesRecommendationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GlobalBudgetRulesRecommendationEventResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []GlobalBudgetRulesRecommendationError and assigns it to the Errors field.
func (o *GlobalBudgetRulesRecommendationEventResponse) SetErrors(v []GlobalBudgetRulesRecommendationError) {
	o.Errors = v
}

func (o GlobalBudgetRulesRecommendationEventResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["countryBudgetRuleEventsRecommendations"] = o.CountryBudgetRuleEventsRecommendations
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGlobalBudgetRulesRecommendationEventResponse struct {
	value *GlobalBudgetRulesRecommendationEventResponse
	isSet bool
}

func (v NullableGlobalBudgetRulesRecommendationEventResponse) Get() *GlobalBudgetRulesRecommendationEventResponse {
	return v.value
}

func (v *NullableGlobalBudgetRulesRecommendationEventResponse) Set(val *GlobalBudgetRulesRecommendationEventResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalBudgetRulesRecommendationEventResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalBudgetRulesRecommendationEventResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalBudgetRulesRecommendationEventResponse(val *GlobalBudgetRulesRecommendationEventResponse) *NullableGlobalBudgetRulesRecommendationEventResponse {
	return &NullableGlobalBudgetRulesRecommendationEventResponse{value: val, isSet: true}
}

func (v NullableGlobalBudgetRulesRecommendationEventResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalBudgetRulesRecommendationEventResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
