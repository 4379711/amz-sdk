package product_pricing_20220501

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// Condition The condition of the item.
type Condition string

// List of Condition
const (
	CONDITION_NEW         Condition = "New"
	CONDITION_USED        Condition = "Used"
	CONDITION_COLLECTIBLE Condition = "Collectible"
	CONDITION_REFURBISHED Condition = "Refurbished"
	CONDITION_CLUB        Condition = "Club"
)

// All allowed values of Condition enum
var AllowedConditionEnumValues = []Condition{
	CONDITION_NEW,
	CONDITION_USED,
	CONDITION_COLLECTIBLE,
	CONDITION_REFURBISHED,
	CONDITION_CLUB,
}

func (v *Condition) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = Condition(value)
	return nil
}

// NewConditionFromValue returns a pointer to a valid Condition
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConditionFromValue(v string) (*Condition, error) {
	ev := Condition(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Condition: valid values are %v", v, AllowedConditionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Condition) IsValid() bool {
	for _, existing := range AllowedConditionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Condition value
func (v Condition) Ptr() *Condition {
	return &v
}

type NullableCondition struct {
	value *Condition
	isSet bool
}

func (v NullableCondition) Get() *Condition {
	return v.value
}

func (v *NullableCondition) Set(val *Condition) {
	v.value = val
	v.isSet = true
}

func (v NullableCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondition(val *Condition) *NullableCondition {
	return &NullableCondition{value: val, isSet: true}
}

func (v NullableCondition) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
