package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDBudgetHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDBudgetHistory{}

// SDBudgetHistory struct for SDBudgetHistory
type SDBudgetHistory struct {
	History []SDRuleBasedBudget `json:"history,omitempty"`
}

// NewSDBudgetHistory instantiates a new SDBudgetHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDBudgetHistory() *SDBudgetHistory {
	this := SDBudgetHistory{}
	return &this
}

// NewSDBudgetHistoryWithDefaults instantiates a new SDBudgetHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDBudgetHistoryWithDefaults() *SDBudgetHistory {
	this := SDBudgetHistory{}
	return &this
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *SDBudgetHistory) GetHistory() []SDRuleBasedBudget {
	if o == nil || IsNil(o.History) {
		var ret []SDRuleBasedBudget
		return ret
	}
	return o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDBudgetHistory) GetHistoryOk() ([]SDRuleBasedBudget, bool) {
	if o == nil || IsNil(o.History) {
		return nil, false
	}
	return o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *SDBudgetHistory) HasHistory() bool {
	if o != nil && !IsNil(o.History) {
		return true
	}

	return false
}

// SetHistory gets a reference to the given []SDRuleBasedBudget and assigns it to the History field.
func (o *SDBudgetHistory) SetHistory(v []SDRuleBasedBudget) {
	o.History = v
}

func (o SDBudgetHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.History) {
		toSerialize["history"] = o.History
	}
	return toSerialize, nil
}

type NullableSDBudgetHistory struct {
	value *SDBudgetHistory
	isSet bool
}

func (v NullableSDBudgetHistory) Get() *SDBudgetHistory {
	return v.value
}

func (v *NullableSDBudgetHistory) Set(val *SDBudgetHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableSDBudgetHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableSDBudgetHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDBudgetHistory(val *SDBudgetHistory) *NullableSDBudgetHistory {
	return &NullableSDBudgetHistory{value: val, isSet: true}
}

func (v NullableSDBudgetHistory) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDBudgetHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
