package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SPBudgetRulesRecommendationEventResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SPBudgetRulesRecommendationEventResponse{}

// SPBudgetRulesRecommendationEventResponse Special events with date range and suggested budget increase.
type SPBudgetRulesRecommendationEventResponse struct {
	// A list of recommended special events with date range and suggested budget increase.
	RecommendedBudgetRuleEvents []SPBudgetRulesRecommendationEvent `json:"recommendedBudgetRuleEvents,omitempty"`
}

// NewSPBudgetRulesRecommendationEventResponse instantiates a new SPBudgetRulesRecommendationEventResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSPBudgetRulesRecommendationEventResponse() *SPBudgetRulesRecommendationEventResponse {
	this := SPBudgetRulesRecommendationEventResponse{}
	return &this
}

// NewSPBudgetRulesRecommendationEventResponseWithDefaults instantiates a new SPBudgetRulesRecommendationEventResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSPBudgetRulesRecommendationEventResponseWithDefaults() *SPBudgetRulesRecommendationEventResponse {
	this := SPBudgetRulesRecommendationEventResponse{}
	return &this
}

// GetRecommendedBudgetRuleEvents returns the RecommendedBudgetRuleEvents field value if set, zero value otherwise.
func (o *SPBudgetRulesRecommendationEventResponse) GetRecommendedBudgetRuleEvents() []SPBudgetRulesRecommendationEvent {
	if o == nil || IsNil(o.RecommendedBudgetRuleEvents) {
		var ret []SPBudgetRulesRecommendationEvent
		return ret
	}
	return o.RecommendedBudgetRuleEvents
}

// GetRecommendedBudgetRuleEventsOk returns a tuple with the RecommendedBudgetRuleEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPBudgetRulesRecommendationEventResponse) GetRecommendedBudgetRuleEventsOk() ([]SPBudgetRulesRecommendationEvent, bool) {
	if o == nil || IsNil(o.RecommendedBudgetRuleEvents) {
		return nil, false
	}
	return o.RecommendedBudgetRuleEvents, true
}

// HasRecommendedBudgetRuleEvents returns a boolean if a field has been set.
func (o *SPBudgetRulesRecommendationEventResponse) HasRecommendedBudgetRuleEvents() bool {
	if o != nil && !IsNil(o.RecommendedBudgetRuleEvents) {
		return true
	}

	return false
}

// SetRecommendedBudgetRuleEvents gets a reference to the given []SPBudgetRulesRecommendationEvent and assigns it to the RecommendedBudgetRuleEvents field.
func (o *SPBudgetRulesRecommendationEventResponse) SetRecommendedBudgetRuleEvents(v []SPBudgetRulesRecommendationEvent) {
	o.RecommendedBudgetRuleEvents = v
}

func (o SPBudgetRulesRecommendationEventResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecommendedBudgetRuleEvents) {
		toSerialize["recommendedBudgetRuleEvents"] = o.RecommendedBudgetRuleEvents
	}
	return toSerialize, nil
}

type NullableSPBudgetRulesRecommendationEventResponse struct {
	value *SPBudgetRulesRecommendationEventResponse
	isSet bool
}

func (v NullableSPBudgetRulesRecommendationEventResponse) Get() *SPBudgetRulesRecommendationEventResponse {
	return v.value
}

func (v *NullableSPBudgetRulesRecommendationEventResponse) Set(val *SPBudgetRulesRecommendationEventResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSPBudgetRulesRecommendationEventResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSPBudgetRulesRecommendationEventResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPBudgetRulesRecommendationEventResponse(val *SPBudgetRulesRecommendationEventResponse) *NullableSPBudgetRulesRecommendationEventResponse {
	return &NullableSPBudgetRulesRecommendationEventResponse{value: val, isSet: true}
}

func (v NullableSPBudgetRulesRecommendationEventResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPBudgetRulesRecommendationEventResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
