package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SPBudgetRulesRecommendationEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SPBudgetRulesRecommendationEvent{}

// SPBudgetRulesRecommendationEvent struct for SPBudgetRulesRecommendationEvent
type SPBudgetRulesRecommendationEvent struct {
	// The event identifier.
	EventId *string `json:"eventId,omitempty"`
	// The end date in YYYYMMDD format.
	EndDate *string `json:"endDate,omitempty"`
	// The suggested budget increase expressed as a percent.
	SuggestedBudgetIncreasePercent *float32 `json:"suggestedBudgetIncreasePercent,omitempty"`
	// The event name.
	EventName *string `json:"eventName,omitempty"`
	// The start date in YYYYMMDD format.
	StartDate *string `json:"startDate,omitempty"`
}

// NewSPBudgetRulesRecommendationEvent instantiates a new SPBudgetRulesRecommendationEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSPBudgetRulesRecommendationEvent() *SPBudgetRulesRecommendationEvent {
	this := SPBudgetRulesRecommendationEvent{}
	return &this
}

// NewSPBudgetRulesRecommendationEventWithDefaults instantiates a new SPBudgetRulesRecommendationEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSPBudgetRulesRecommendationEventWithDefaults() *SPBudgetRulesRecommendationEvent {
	this := SPBudgetRulesRecommendationEvent{}
	return &this
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *SPBudgetRulesRecommendationEvent) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPBudgetRulesRecommendationEvent) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *SPBudgetRulesRecommendationEvent) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *SPBudgetRulesRecommendationEvent) SetEventId(v string) {
	o.EventId = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SPBudgetRulesRecommendationEvent) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPBudgetRulesRecommendationEvent) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SPBudgetRulesRecommendationEvent) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *SPBudgetRulesRecommendationEvent) SetEndDate(v string) {
	o.EndDate = &v
}

// GetSuggestedBudgetIncreasePercent returns the SuggestedBudgetIncreasePercent field value if set, zero value otherwise.
func (o *SPBudgetRulesRecommendationEvent) GetSuggestedBudgetIncreasePercent() float32 {
	if o == nil || IsNil(o.SuggestedBudgetIncreasePercent) {
		var ret float32
		return ret
	}
	return *o.SuggestedBudgetIncreasePercent
}

// GetSuggestedBudgetIncreasePercentOk returns a tuple with the SuggestedBudgetIncreasePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPBudgetRulesRecommendationEvent) GetSuggestedBudgetIncreasePercentOk() (*float32, bool) {
	if o == nil || IsNil(o.SuggestedBudgetIncreasePercent) {
		return nil, false
	}
	return o.SuggestedBudgetIncreasePercent, true
}

// HasSuggestedBudgetIncreasePercent returns a boolean if a field has been set.
func (o *SPBudgetRulesRecommendationEvent) HasSuggestedBudgetIncreasePercent() bool {
	if o != nil && !IsNil(o.SuggestedBudgetIncreasePercent) {
		return true
	}

	return false
}

// SetSuggestedBudgetIncreasePercent gets a reference to the given float32 and assigns it to the SuggestedBudgetIncreasePercent field.
func (o *SPBudgetRulesRecommendationEvent) SetSuggestedBudgetIncreasePercent(v float32) {
	o.SuggestedBudgetIncreasePercent = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *SPBudgetRulesRecommendationEvent) GetEventName() string {
	if o == nil || IsNil(o.EventName) {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPBudgetRulesRecommendationEvent) GetEventNameOk() (*string, bool) {
	if o == nil || IsNil(o.EventName) {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *SPBudgetRulesRecommendationEvent) HasEventName() bool {
	if o != nil && !IsNil(o.EventName) {
		return true
	}

	return false
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *SPBudgetRulesRecommendationEvent) SetEventName(v string) {
	o.EventName = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SPBudgetRulesRecommendationEvent) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPBudgetRulesRecommendationEvent) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SPBudgetRulesRecommendationEvent) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SPBudgetRulesRecommendationEvent) SetStartDate(v string) {
	o.StartDate = &v
}

func (o SPBudgetRulesRecommendationEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventId) {
		toSerialize["eventId"] = o.EventId
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.SuggestedBudgetIncreasePercent) {
		toSerialize["suggestedBudgetIncreasePercent"] = o.SuggestedBudgetIncreasePercent
	}
	if !IsNil(o.EventName) {
		toSerialize["eventName"] = o.EventName
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	return toSerialize, nil
}

type NullableSPBudgetRulesRecommendationEvent struct {
	value *SPBudgetRulesRecommendationEvent
	isSet bool
}

func (v NullableSPBudgetRulesRecommendationEvent) Get() *SPBudgetRulesRecommendationEvent {
	return v.value
}

func (v *NullableSPBudgetRulesRecommendationEvent) Set(val *SPBudgetRulesRecommendationEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSPBudgetRulesRecommendationEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSPBudgetRulesRecommendationEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPBudgetRulesRecommendationEvent(val *SPBudgetRulesRecommendationEvent) *NullableSPBudgetRulesRecommendationEvent {
	return &NullableSPBudgetRulesRecommendationEvent{value: val, isSet: true}
}

func (v NullableSPBudgetRulesRecommendationEvent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPBudgetRulesRecommendationEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
