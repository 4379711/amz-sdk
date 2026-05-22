package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the RuleBasedBudget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleBasedBudget{}

// RuleBasedBudget struct for RuleBasedBudget
type RuleBasedBudget struct {
	IsProcessing       *bool    `json:"isProcessing,omitempty"`
	ApplicableRuleName *string  `json:"applicableRuleName,omitempty"`
	Value              *float64 `json:"value,omitempty"`
	ApplicableRuleId   *string  `json:"applicableRuleId,omitempty"`
}

// NewRuleBasedBudget instantiates a new RuleBasedBudget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleBasedBudget() *RuleBasedBudget {
	this := RuleBasedBudget{}
	return &this
}

// NewRuleBasedBudgetWithDefaults instantiates a new RuleBasedBudget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleBasedBudgetWithDefaults() *RuleBasedBudget {
	this := RuleBasedBudget{}
	return &this
}

// GetIsProcessing returns the IsProcessing field value if set, zero value otherwise.
func (o *RuleBasedBudget) GetIsProcessing() bool {
	if o == nil || IsNil(o.IsProcessing) {
		var ret bool
		return ret
	}
	return *o.IsProcessing
}

// GetIsProcessingOk returns a tuple with the IsProcessing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleBasedBudget) GetIsProcessingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsProcessing) {
		return nil, false
	}
	return o.IsProcessing, true
}

// HasIsProcessing returns a boolean if a field has been set.
func (o *RuleBasedBudget) HasIsProcessing() bool {
	if o != nil && !IsNil(o.IsProcessing) {
		return true
	}

	return false
}

// SetIsProcessing gets a reference to the given bool and assigns it to the IsProcessing field.
func (o *RuleBasedBudget) SetIsProcessing(v bool) {
	o.IsProcessing = &v
}

// GetApplicableRuleName returns the ApplicableRuleName field value if set, zero value otherwise.
func (o *RuleBasedBudget) GetApplicableRuleName() string {
	if o == nil || IsNil(o.ApplicableRuleName) {
		var ret string
		return ret
	}
	return *o.ApplicableRuleName
}

// GetApplicableRuleNameOk returns a tuple with the ApplicableRuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleBasedBudget) GetApplicableRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicableRuleName) {
		return nil, false
	}
	return o.ApplicableRuleName, true
}

// HasApplicableRuleName returns a boolean if a field has been set.
func (o *RuleBasedBudget) HasApplicableRuleName() bool {
	if o != nil && !IsNil(o.ApplicableRuleName) {
		return true
	}

	return false
}

// SetApplicableRuleName gets a reference to the given string and assigns it to the ApplicableRuleName field.
func (o *RuleBasedBudget) SetApplicableRuleName(v string) {
	o.ApplicableRuleName = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RuleBasedBudget) GetValue() float64 {
	if o == nil || IsNil(o.Value) {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleBasedBudget) GetValueOk() (*float64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RuleBasedBudget) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *RuleBasedBudget) SetValue(v float64) {
	o.Value = &v
}

// GetApplicableRuleId returns the ApplicableRuleId field value if set, zero value otherwise.
func (o *RuleBasedBudget) GetApplicableRuleId() string {
	if o == nil || IsNil(o.ApplicableRuleId) {
		var ret string
		return ret
	}
	return *o.ApplicableRuleId
}

// GetApplicableRuleIdOk returns a tuple with the ApplicableRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleBasedBudget) GetApplicableRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicableRuleId) {
		return nil, false
	}
	return o.ApplicableRuleId, true
}

// HasApplicableRuleId returns a boolean if a field has been set.
func (o *RuleBasedBudget) HasApplicableRuleId() bool {
	if o != nil && !IsNil(o.ApplicableRuleId) {
		return true
	}

	return false
}

// SetApplicableRuleId gets a reference to the given string and assigns it to the ApplicableRuleId field.
func (o *RuleBasedBudget) SetApplicableRuleId(v string) {
	o.ApplicableRuleId = &v
}

func (o RuleBasedBudget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsProcessing) {
		toSerialize["isProcessing"] = o.IsProcessing
	}
	if !IsNil(o.ApplicableRuleName) {
		toSerialize["applicableRuleName"] = o.ApplicableRuleName
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.ApplicableRuleId) {
		toSerialize["applicableRuleId"] = o.ApplicableRuleId
	}
	return toSerialize, nil
}

type NullableRuleBasedBudget struct {
	value *RuleBasedBudget
	isSet bool
}

func (v NullableRuleBasedBudget) Get() *RuleBasedBudget {
	return v.value
}

func (v *NullableRuleBasedBudget) Set(val *RuleBasedBudget) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleBasedBudget) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleBasedBudget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleBasedBudget(val *RuleBasedBudget) *NullableRuleBasedBudget {
	return &NullableRuleBasedBudget{value: val, isSet: true}
}

func (v NullableRuleBasedBudget) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRuleBasedBudget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
