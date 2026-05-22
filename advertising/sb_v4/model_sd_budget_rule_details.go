package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SDBudgetRuleDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDBudgetRuleDetails{}

// SDBudgetRuleDetails Object representing details of a budget rule for SD campaign
type SDBudgetRuleDetails struct {
	Duration         *RuleDuration     `json:"duration,omitempty"`
	Recurrence       *Recurrence       `json:"recurrence,omitempty"`
	RuleType         *SDRuleType       `json:"ruleType,omitempty"`
	BudgetIncreaseBy *BudgetIncreaseBy `json:"budgetIncreaseBy,omitempty"`
	// The budget rule name. Required to be unique within a campaign.
	Name                        *string                      `json:"name,omitempty"`
	PerformanceMeasureCondition *PerformanceMeasureCondition `json:"performanceMeasureCondition,omitempty"`
}

// NewSDBudgetRuleDetails instantiates a new SDBudgetRuleDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDBudgetRuleDetails() *SDBudgetRuleDetails {
	this := SDBudgetRuleDetails{}
	return &this
}

// NewSDBudgetRuleDetailsWithDefaults instantiates a new SDBudgetRuleDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDBudgetRuleDetailsWithDefaults() *SDBudgetRuleDetails {
	this := SDBudgetRuleDetails{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SDBudgetRuleDetails) GetDuration() RuleDuration {
	if o == nil || IsNil(o.Duration) {
		var ret RuleDuration
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDBudgetRuleDetails) GetDurationOk() (*RuleDuration, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SDBudgetRuleDetails) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given RuleDuration and assigns it to the Duration field.
func (o *SDBudgetRuleDetails) SetDuration(v RuleDuration) {
	o.Duration = &v
}

// GetRecurrence returns the Recurrence field value if set, zero value otherwise.
func (o *SDBudgetRuleDetails) GetRecurrence() Recurrence {
	if o == nil || IsNil(o.Recurrence) {
		var ret Recurrence
		return ret
	}
	return *o.Recurrence
}

// GetRecurrenceOk returns a tuple with the Recurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDBudgetRuleDetails) GetRecurrenceOk() (*Recurrence, bool) {
	if o == nil || IsNil(o.Recurrence) {
		return nil, false
	}
	return o.Recurrence, true
}

// HasRecurrence returns a boolean if a field has been set.
func (o *SDBudgetRuleDetails) HasRecurrence() bool {
	if o != nil && !IsNil(o.Recurrence) {
		return true
	}

	return false
}

// SetRecurrence gets a reference to the given Recurrence and assigns it to the Recurrence field.
func (o *SDBudgetRuleDetails) SetRecurrence(v Recurrence) {
	o.Recurrence = &v
}

// GetRuleType returns the RuleType field value if set, zero value otherwise.
func (o *SDBudgetRuleDetails) GetRuleType() SDRuleType {
	if o == nil || IsNil(o.RuleType) {
		var ret SDRuleType
		return ret
	}
	return *o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDBudgetRuleDetails) GetRuleTypeOk() (*SDRuleType, bool) {
	if o == nil || IsNil(o.RuleType) {
		return nil, false
	}
	return o.RuleType, true
}

// HasRuleType returns a boolean if a field has been set.
func (o *SDBudgetRuleDetails) HasRuleType() bool {
	if o != nil && !IsNil(o.RuleType) {
		return true
	}

	return false
}

// SetRuleType gets a reference to the given SDRuleType and assigns it to the RuleType field.
func (o *SDBudgetRuleDetails) SetRuleType(v SDRuleType) {
	o.RuleType = &v
}

// GetBudgetIncreaseBy returns the BudgetIncreaseBy field value if set, zero value otherwise.
func (o *SDBudgetRuleDetails) GetBudgetIncreaseBy() BudgetIncreaseBy {
	if o == nil || IsNil(o.BudgetIncreaseBy) {
		var ret BudgetIncreaseBy
		return ret
	}
	return *o.BudgetIncreaseBy
}

// GetBudgetIncreaseByOk returns a tuple with the BudgetIncreaseBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDBudgetRuleDetails) GetBudgetIncreaseByOk() (*BudgetIncreaseBy, bool) {
	if o == nil || IsNil(o.BudgetIncreaseBy) {
		return nil, false
	}
	return o.BudgetIncreaseBy, true
}

// HasBudgetIncreaseBy returns a boolean if a field has been set.
func (o *SDBudgetRuleDetails) HasBudgetIncreaseBy() bool {
	if o != nil && !IsNil(o.BudgetIncreaseBy) {
		return true
	}

	return false
}

// SetBudgetIncreaseBy gets a reference to the given BudgetIncreaseBy and assigns it to the BudgetIncreaseBy field.
func (o *SDBudgetRuleDetails) SetBudgetIncreaseBy(v BudgetIncreaseBy) {
	o.BudgetIncreaseBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SDBudgetRuleDetails) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDBudgetRuleDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SDBudgetRuleDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SDBudgetRuleDetails) SetName(v string) {
	o.Name = &v
}

// GetPerformanceMeasureCondition returns the PerformanceMeasureCondition field value if set, zero value otherwise.
func (o *SDBudgetRuleDetails) GetPerformanceMeasureCondition() PerformanceMeasureCondition {
	if o == nil || IsNil(o.PerformanceMeasureCondition) {
		var ret PerformanceMeasureCondition
		return ret
	}
	return *o.PerformanceMeasureCondition
}

// GetPerformanceMeasureConditionOk returns a tuple with the PerformanceMeasureCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDBudgetRuleDetails) GetPerformanceMeasureConditionOk() (*PerformanceMeasureCondition, bool) {
	if o == nil || IsNil(o.PerformanceMeasureCondition) {
		return nil, false
	}
	return o.PerformanceMeasureCondition, true
}

// HasPerformanceMeasureCondition returns a boolean if a field has been set.
func (o *SDBudgetRuleDetails) HasPerformanceMeasureCondition() bool {
	if o != nil && !IsNil(o.PerformanceMeasureCondition) {
		return true
	}

	return false
}

// SetPerformanceMeasureCondition gets a reference to the given PerformanceMeasureCondition and assigns it to the PerformanceMeasureCondition field.
func (o *SDBudgetRuleDetails) SetPerformanceMeasureCondition(v PerformanceMeasureCondition) {
	o.PerformanceMeasureCondition = &v
}

func (o SDBudgetRuleDetails) ToMap() (map[string]interface{}, error) {
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

type NullableSDBudgetRuleDetails struct {
	value *SDBudgetRuleDetails
	isSet bool
}

func (v NullableSDBudgetRuleDetails) Get() *SDBudgetRuleDetails {
	return v.value
}

func (v *NullableSDBudgetRuleDetails) Set(val *SDBudgetRuleDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSDBudgetRuleDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSDBudgetRuleDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDBudgetRuleDetails(val *SDBudgetRuleDetails) *NullableSDBudgetRuleDetails {
	return &NullableSDBudgetRuleDetails{value: val, isSet: true}
}

func (v NullableSDBudgetRuleDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDBudgetRuleDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
