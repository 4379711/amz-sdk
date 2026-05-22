package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SPBudgetRuleDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SPBudgetRuleDetails{}

// SPBudgetRuleDetails Object representing details of a budget rule for SP campaign
type SPBudgetRuleDetails struct {
	Duration         *RuleDuration     `json:"duration,omitempty"`
	Recurrence       *Recurrence       `json:"recurrence,omitempty"`
	RuleType         *SPRuleType       `json:"ruleType,omitempty"`
	BudgetIncreaseBy *BudgetIncreaseBy `json:"budgetIncreaseBy,omitempty"`
	// The budget rule name. Required to be unique within a campaign.
	Name                        *string                      `json:"name,omitempty"`
	PerformanceMeasureCondition *PerformanceMeasureCondition `json:"performanceMeasureCondition,omitempty"`
}

// NewSPBudgetRuleDetails instantiates a new SPBudgetRuleDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSPBudgetRuleDetails() *SPBudgetRuleDetails {
	this := SPBudgetRuleDetails{}
	return &this
}

// NewSPBudgetRuleDetailsWithDefaults instantiates a new SPBudgetRuleDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSPBudgetRuleDetailsWithDefaults() *SPBudgetRuleDetails {
	this := SPBudgetRuleDetails{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SPBudgetRuleDetails) GetDuration() RuleDuration {
	if o == nil || IsNil(o.Duration) {
		var ret RuleDuration
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPBudgetRuleDetails) GetDurationOk() (*RuleDuration, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SPBudgetRuleDetails) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given RuleDuration and assigns it to the Duration field.
func (o *SPBudgetRuleDetails) SetDuration(v RuleDuration) {
	o.Duration = &v
}

// GetRecurrence returns the Recurrence field value if set, zero value otherwise.
func (o *SPBudgetRuleDetails) GetRecurrence() Recurrence {
	if o == nil || IsNil(o.Recurrence) {
		var ret Recurrence
		return ret
	}
	return *o.Recurrence
}

// GetRecurrenceOk returns a tuple with the Recurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPBudgetRuleDetails) GetRecurrenceOk() (*Recurrence, bool) {
	if o == nil || IsNil(o.Recurrence) {
		return nil, false
	}
	return o.Recurrence, true
}

// HasRecurrence returns a boolean if a field has been set.
func (o *SPBudgetRuleDetails) HasRecurrence() bool {
	if o != nil && !IsNil(o.Recurrence) {
		return true
	}

	return false
}

// SetRecurrence gets a reference to the given Recurrence and assigns it to the Recurrence field.
func (o *SPBudgetRuleDetails) SetRecurrence(v Recurrence) {
	o.Recurrence = &v
}

// GetRuleType returns the RuleType field value if set, zero value otherwise.
func (o *SPBudgetRuleDetails) GetRuleType() SPRuleType {
	if o == nil || IsNil(o.RuleType) {
		var ret SPRuleType
		return ret
	}
	return *o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPBudgetRuleDetails) GetRuleTypeOk() (*SPRuleType, bool) {
	if o == nil || IsNil(o.RuleType) {
		return nil, false
	}
	return o.RuleType, true
}

// HasRuleType returns a boolean if a field has been set.
func (o *SPBudgetRuleDetails) HasRuleType() bool {
	if o != nil && !IsNil(o.RuleType) {
		return true
	}

	return false
}

// SetRuleType gets a reference to the given SPRuleType and assigns it to the RuleType field.
func (o *SPBudgetRuleDetails) SetRuleType(v SPRuleType) {
	o.RuleType = &v
}

// GetBudgetIncreaseBy returns the BudgetIncreaseBy field value if set, zero value otherwise.
func (o *SPBudgetRuleDetails) GetBudgetIncreaseBy() BudgetIncreaseBy {
	if o == nil || IsNil(o.BudgetIncreaseBy) {
		var ret BudgetIncreaseBy
		return ret
	}
	return *o.BudgetIncreaseBy
}

// GetBudgetIncreaseByOk returns a tuple with the BudgetIncreaseBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPBudgetRuleDetails) GetBudgetIncreaseByOk() (*BudgetIncreaseBy, bool) {
	if o == nil || IsNil(o.BudgetIncreaseBy) {
		return nil, false
	}
	return o.BudgetIncreaseBy, true
}

// HasBudgetIncreaseBy returns a boolean if a field has been set.
func (o *SPBudgetRuleDetails) HasBudgetIncreaseBy() bool {
	if o != nil && !IsNil(o.BudgetIncreaseBy) {
		return true
	}

	return false
}

// SetBudgetIncreaseBy gets a reference to the given BudgetIncreaseBy and assigns it to the BudgetIncreaseBy field.
func (o *SPBudgetRuleDetails) SetBudgetIncreaseBy(v BudgetIncreaseBy) {
	o.BudgetIncreaseBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SPBudgetRuleDetails) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPBudgetRuleDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SPBudgetRuleDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SPBudgetRuleDetails) SetName(v string) {
	o.Name = &v
}

// GetPerformanceMeasureCondition returns the PerformanceMeasureCondition field value if set, zero value otherwise.
func (o *SPBudgetRuleDetails) GetPerformanceMeasureCondition() PerformanceMeasureCondition {
	if o == nil || IsNil(o.PerformanceMeasureCondition) {
		var ret PerformanceMeasureCondition
		return ret
	}
	return *o.PerformanceMeasureCondition
}

// GetPerformanceMeasureConditionOk returns a tuple with the PerformanceMeasureCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPBudgetRuleDetails) GetPerformanceMeasureConditionOk() (*PerformanceMeasureCondition, bool) {
	if o == nil || IsNil(o.PerformanceMeasureCondition) {
		return nil, false
	}
	return o.PerformanceMeasureCondition, true
}

// HasPerformanceMeasureCondition returns a boolean if a field has been set.
func (o *SPBudgetRuleDetails) HasPerformanceMeasureCondition() bool {
	if o != nil && !IsNil(o.PerformanceMeasureCondition) {
		return true
	}

	return false
}

// SetPerformanceMeasureCondition gets a reference to the given PerformanceMeasureCondition and assigns it to the PerformanceMeasureCondition field.
func (o *SPBudgetRuleDetails) SetPerformanceMeasureCondition(v PerformanceMeasureCondition) {
	o.PerformanceMeasureCondition = &v
}

func (o SPBudgetRuleDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Recurrence) {
		toSerialize["recurrence"] = o.Recurrence
	}
	if !IsNil(o.RuleType) {
		toSerialize["ruleType"] = o.RuleType
	}
	if !IsNil(o.BudgetIncreaseBy) {
		toSerialize["budgetIncreaseBy"] = o.BudgetIncreaseBy
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PerformanceMeasureCondition) {
		toSerialize["performanceMeasureCondition"] = o.PerformanceMeasureCondition
	}
	return toSerialize, nil
}

type NullableSPBudgetRuleDetails struct {
	value *SPBudgetRuleDetails
	isSet bool
}

func (v NullableSPBudgetRuleDetails) Get() *SPBudgetRuleDetails {
	return v.value
}

func (v *NullableSPBudgetRuleDetails) Set(val *SPBudgetRuleDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSPBudgetRuleDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSPBudgetRuleDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPBudgetRuleDetails(val *SPBudgetRuleDetails) *NullableSPBudgetRuleDetails {
	return &NullableSPBudgetRuleDetails{value: val, isSet: true}
}

func (v NullableSPBudgetRuleDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPBudgetRuleDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
