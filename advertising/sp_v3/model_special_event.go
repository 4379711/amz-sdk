package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SpecialEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpecialEvent{}

// SpecialEvent struct for SpecialEvent
type SpecialEvent struct {
	// Deprecated. The factor used to boost the recommended budget.
	BudgetModifier *float32 `json:"budgetModifier,omitempty"`
	// The end date of the special event in YYYYMMDD format.
	EndDate *string `json:"endDate,omitempty"`
	// Recommended daily budget for the new campaign during the special event period.
	DailyBudget *float32 `json:"dailyBudget,omitempty"`
	// The key of the special event.
	EventKey *string `json:"eventKey,omitempty"`
	// The name of the special event.
	EventName *string `json:"eventName,omitempty"`
	// The start date of the special event in YYYYMMDD format.
	StartDate *string    `json:"startDate,omitempty"`
	Benchmark *Benchmark `json:"benchmark,omitempty"`
}

// NewSpecialEvent instantiates a new SpecialEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecialEvent() *SpecialEvent {
	this := SpecialEvent{}
	return &this
}

// NewSpecialEventWithDefaults instantiates a new SpecialEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecialEventWithDefaults() *SpecialEvent {
	this := SpecialEvent{}
	return &this
}

// GetBudgetModifier returns the BudgetModifier field value if set, zero value otherwise.
func (o *SpecialEvent) GetBudgetModifier() float32 {
	if o == nil || IsNil(o.BudgetModifier) {
		var ret float32
		return ret
	}
	return *o.BudgetModifier
}

// GetBudgetModifierOk returns a tuple with the BudgetModifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialEvent) GetBudgetModifierOk() (*float32, bool) {
	if o == nil || IsNil(o.BudgetModifier) {
		return nil, false
	}
	return o.BudgetModifier, true
}

// HasBudgetModifier returns a boolean if a field has been set.
func (o *SpecialEvent) HasBudgetModifier() bool {
	if o != nil && !IsNil(o.BudgetModifier) {
		return true
	}

	return false
}

// SetBudgetModifier gets a reference to the given float32 and assigns it to the BudgetModifier field.
func (o *SpecialEvent) SetBudgetModifier(v float32) {
	o.BudgetModifier = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SpecialEvent) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialEvent) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SpecialEvent) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *SpecialEvent) SetEndDate(v string) {
	o.EndDate = &v
}

// GetDailyBudget returns the DailyBudget field value if set, zero value otherwise.
func (o *SpecialEvent) GetDailyBudget() float32 {
	if o == nil || IsNil(o.DailyBudget) {
		var ret float32
		return ret
	}
	return *o.DailyBudget
}

// GetDailyBudgetOk returns a tuple with the DailyBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialEvent) GetDailyBudgetOk() (*float32, bool) {
	if o == nil || IsNil(o.DailyBudget) {
		return nil, false
	}
	return o.DailyBudget, true
}

// HasDailyBudget returns a boolean if a field has been set.
func (o *SpecialEvent) HasDailyBudget() bool {
	if o != nil && !IsNil(o.DailyBudget) {
		return true
	}

	return false
}

// SetDailyBudget gets a reference to the given float32 and assigns it to the DailyBudget field.
func (o *SpecialEvent) SetDailyBudget(v float32) {
	o.DailyBudget = &v
}

// GetEventKey returns the EventKey field value if set, zero value otherwise.
func (o *SpecialEvent) GetEventKey() string {
	if o == nil || IsNil(o.EventKey) {
		var ret string
		return ret
	}
	return *o.EventKey
}

// GetEventKeyOk returns a tuple with the EventKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialEvent) GetEventKeyOk() (*string, bool) {
	if o == nil || IsNil(o.EventKey) {
		return nil, false
	}
	return o.EventKey, true
}

// HasEventKey returns a boolean if a field has been set.
func (o *SpecialEvent) HasEventKey() bool {
	if o != nil && !IsNil(o.EventKey) {
		return true
	}

	return false
}

// SetEventKey gets a reference to the given string and assigns it to the EventKey field.
func (o *SpecialEvent) SetEventKey(v string) {
	o.EventKey = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *SpecialEvent) GetEventName() string {
	if o == nil || IsNil(o.EventName) {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialEvent) GetEventNameOk() (*string, bool) {
	if o == nil || IsNil(o.EventName) {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *SpecialEvent) HasEventName() bool {
	if o != nil && !IsNil(o.EventName) {
		return true
	}

	return false
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *SpecialEvent) SetEventName(v string) {
	o.EventName = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SpecialEvent) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialEvent) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SpecialEvent) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SpecialEvent) SetStartDate(v string) {
	o.StartDate = &v
}

// GetBenchmark returns the Benchmark field value if set, zero value otherwise.
func (o *SpecialEvent) GetBenchmark() Benchmark {
	if o == nil || IsNil(o.Benchmark) {
		var ret Benchmark
		return ret
	}
	return *o.Benchmark
}

// GetBenchmarkOk returns a tuple with the Benchmark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialEvent) GetBenchmarkOk() (*Benchmark, bool) {
	if o == nil || IsNil(o.Benchmark) {
		return nil, false
	}
	return o.Benchmark, true
}

// HasBenchmark returns a boolean if a field has been set.
func (o *SpecialEvent) HasBenchmark() bool {
	if o != nil && !IsNil(o.Benchmark) {
		return true
	}

	return false
}

// SetBenchmark gets a reference to the given Benchmark and assigns it to the Benchmark field.
func (o *SpecialEvent) SetBenchmark(v Benchmark) {
	o.Benchmark = &v
}

func (o SpecialEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetModifier) {
		toSerialize["budgetModifier"] = o.BudgetModifier
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.DailyBudget) {
		toSerialize["dailyBudget"] = o.DailyBudget
	}
	if !IsNil(o.EventKey) {
		toSerialize["eventKey"] = o.EventKey
	}
	if !IsNil(o.EventName) {
		toSerialize["eventName"] = o.EventName
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.Benchmark) {
		toSerialize["benchmark"] = o.Benchmark
	}
	return toSerialize, nil
}

type NullableSpecialEvent struct {
	value *SpecialEvent
	isSet bool
}

func (v NullableSpecialEvent) Get() *SpecialEvent {
	return v.value
}

func (v *NullableSpecialEvent) Set(val *SpecialEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecialEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecialEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecialEvent(val *SpecialEvent) *NullableSpecialEvent {
	return &NullableSpecialEvent{value: val, isSet: true}
}

func (v NullableSpecialEvent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSpecialEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
