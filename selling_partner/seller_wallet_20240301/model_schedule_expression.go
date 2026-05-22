package seller_wallet_20240301

import (
	"github.com/bytedance/sonic"
)

// checks if the ScheduleExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleExpression{}

// ScheduleExpression The configuration of the schedule.
type ScheduleExpression struct {
	ScheduleExpressionType ScheduleExpressionType `json:"scheduleExpressionType"`
	RecurringFrequency     *RecurringFrequency    `json:"recurringFrequency,omitempty"`
}

type _ScheduleExpression ScheduleExpression

// NewScheduleExpression instantiates a new ScheduleExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleExpression(scheduleExpressionType ScheduleExpressionType) *ScheduleExpression {
	this := ScheduleExpression{}
	this.ScheduleExpressionType = scheduleExpressionType
	return &this
}

// NewScheduleExpressionWithDefaults instantiates a new ScheduleExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleExpressionWithDefaults() *ScheduleExpression {
	this := ScheduleExpression{}
	return &this
}

// GetScheduleExpressionType returns the ScheduleExpressionType field value
func (o *ScheduleExpression) GetScheduleExpressionType() ScheduleExpressionType {
	if o == nil {
		var ret ScheduleExpressionType
		return ret
	}

	return o.ScheduleExpressionType
}

// GetScheduleExpressionTypeOk returns a tuple with the ScheduleExpressionType field value
// and a boolean to check if the value has been set.
func (o *ScheduleExpression) GetScheduleExpressionTypeOk() (*ScheduleExpressionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleExpressionType, true
}

// SetScheduleExpressionType sets field value
func (o *ScheduleExpression) SetScheduleExpressionType(v ScheduleExpressionType) {
	o.ScheduleExpressionType = v
}

// GetRecurringFrequency returns the RecurringFrequency field value if set, zero value otherwise.
func (o *ScheduleExpression) GetRecurringFrequency() RecurringFrequency {
	if o == nil || IsNil(o.RecurringFrequency) {
		var ret RecurringFrequency
		return ret
	}
	return *o.RecurringFrequency
}

// GetRecurringFrequencyOk returns a tuple with the RecurringFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleExpression) GetRecurringFrequencyOk() (*RecurringFrequency, bool) {
	if o == nil || IsNil(o.RecurringFrequency) {
		return nil, false
	}
	return o.RecurringFrequency, true
}

// HasRecurringFrequency returns a boolean if a field has been set.
func (o *ScheduleExpression) HasRecurringFrequency() bool {
	if o != nil && !IsNil(o.RecurringFrequency) {
		return true
	}

	return false
}

// SetRecurringFrequency gets a reference to the given RecurringFrequency and assigns it to the RecurringFrequency field.
func (o *ScheduleExpression) SetRecurringFrequency(v RecurringFrequency) {
	o.RecurringFrequency = &v
}

func (o ScheduleExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["scheduleExpressionType"] = o.ScheduleExpressionType
	if !IsNil(o.RecurringFrequency) {
		toSerialize["recurringFrequency"] = o.RecurringFrequency
	}
	return toSerialize, nil
}

type NullableScheduleExpression struct {
	value *ScheduleExpression
	isSet bool
}

func (v NullableScheduleExpression) Get() *ScheduleExpression {
	return v.value
}

func (v *NullableScheduleExpression) Set(val *ScheduleExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleExpression(val *ScheduleExpression) *NullableScheduleExpression {
	return &NullableScheduleExpression{value: val, isSet: true}
}

func (v NullableScheduleExpression) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableScheduleExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
