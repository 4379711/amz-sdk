package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetIncreaseBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetIncreaseBy{}

// BudgetIncreaseBy struct for BudgetIncreaseBy
type BudgetIncreaseBy struct {
	Type BudgetChangeType `json:"type"`
	// The budget value.
	Value float64 `json:"value"`
}

type _BudgetIncreaseBy BudgetIncreaseBy

// NewBudgetIncreaseBy instantiates a new BudgetIncreaseBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetIncreaseBy(type_ BudgetChangeType, value float64) *BudgetIncreaseBy {
	this := BudgetIncreaseBy{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewBudgetIncreaseByWithDefaults instantiates a new BudgetIncreaseBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetIncreaseByWithDefaults() *BudgetIncreaseBy {
	this := BudgetIncreaseBy{}
	return &this
}

// GetType returns the Type field value
func (o *BudgetIncreaseBy) GetType() BudgetChangeType {
	if o == nil {
		var ret BudgetChangeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BudgetIncreaseBy) GetTypeOk() (*BudgetChangeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BudgetIncreaseBy) SetType(v BudgetChangeType) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *BudgetIncreaseBy) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *BudgetIncreaseBy) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *BudgetIncreaseBy) SetValue(v float64) {
	o.Value = v
}

func (o BudgetIncreaseBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableBudgetIncreaseBy struct {
	value *BudgetIncreaseBy
	isSet bool
}

func (v NullableBudgetIncreaseBy) Get() *BudgetIncreaseBy {
	return v.value
}

func (v *NullableBudgetIncreaseBy) Set(val *BudgetIncreaseBy) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetIncreaseBy) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetIncreaseBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetIncreaseBy(val *BudgetIncreaseBy) *NullableBudgetIncreaseBy {
	return &NullableBudgetIncreaseBy{value: val, isSet: true}
}

func (v NullableBudgetIncreaseBy) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetIncreaseBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
