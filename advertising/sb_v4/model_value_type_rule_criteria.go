package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ValueTypeRuleCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueTypeRuleCriteria{}

// ValueTypeRuleCriteria struct for ValueTypeRuleCriteria
type ValueTypeRuleCriteria struct {
	// Enum: \"LESS_THAN_OR_EQUAL_TO\"  The comparison operator.
	ComparisonOperator *string `json:"comparisonOperator,omitempty"`
	// The value of the threshold associated with the attribute.
	Value *float64 `json:"value,omitempty"`
}

// NewValueTypeRuleCriteria instantiates a new ValueTypeRuleCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueTypeRuleCriteria() *ValueTypeRuleCriteria {
	this := ValueTypeRuleCriteria{}
	return &this
}

// NewValueTypeRuleCriteriaWithDefaults instantiates a new ValueTypeRuleCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueTypeRuleCriteriaWithDefaults() *ValueTypeRuleCriteria {
	this := ValueTypeRuleCriteria{}
	return &this
}

// GetComparisonOperator returns the ComparisonOperator field value if set, zero value otherwise.
func (o *ValueTypeRuleCriteria) GetComparisonOperator() string {
	if o == nil || IsNil(o.ComparisonOperator) {
		var ret string
		return ret
	}
	return *o.ComparisonOperator
}

// GetComparisonOperatorOk returns a tuple with the ComparisonOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueTypeRuleCriteria) GetComparisonOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.ComparisonOperator) {
		return nil, false
	}
	return o.ComparisonOperator, true
}

// HasComparisonOperator returns a boolean if a field has been set.
func (o *ValueTypeRuleCriteria) HasComparisonOperator() bool {
	if o != nil && !IsNil(o.ComparisonOperator) {
		return true
	}

	return false
}

// SetComparisonOperator gets a reference to the given string and assigns it to the ComparisonOperator field.
func (o *ValueTypeRuleCriteria) SetComparisonOperator(v string) {
	o.ComparisonOperator = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ValueTypeRuleCriteria) GetValue() float64 {
	if o == nil || IsNil(o.Value) {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueTypeRuleCriteria) GetValueOk() (*float64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ValueTypeRuleCriteria) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *ValueTypeRuleCriteria) SetValue(v float64) {
	o.Value = &v
}

func (o ValueTypeRuleCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComparisonOperator) {
		toSerialize["comparisonOperator"] = o.ComparisonOperator
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableValueTypeRuleCriteria struct {
	value *ValueTypeRuleCriteria
	isSet bool
}

func (v NullableValueTypeRuleCriteria) Get() *ValueTypeRuleCriteria {
	return v.value
}

func (v *NullableValueTypeRuleCriteria) Set(val *ValueTypeRuleCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableValueTypeRuleCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableValueTypeRuleCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueTypeRuleCriteria(val *ValueTypeRuleCriteria) *NullableValueTypeRuleCriteria {
	return &NullableValueTypeRuleCriteria{value: val, isSet: true}
}

func (v NullableValueTypeRuleCriteria) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableValueTypeRuleCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
