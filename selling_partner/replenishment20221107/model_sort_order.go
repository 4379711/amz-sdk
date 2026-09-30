package replenishment20221107

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SortOrder The sort order.
type SortOrder string

// List of SortOrder
const (
	SORTORDER_ASC  SortOrder = "ASC"
	SORTORDER_DESC SortOrder = "DESC"
)

// All allowed values of SortOrder enum
var AllowedSortOrderEnumValues = []SortOrder{
	SORTORDER_ASC,
	SORTORDER_DESC,
}

func (v *SortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SortOrder(value)
	return nil
}

// NewSortOrderFromValue returns a pointer to a valid SortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSortOrderFromValue(v string) (*SortOrder, error) {
	ev := SortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SortOrder: valid values are %v", v, AllowedSortOrderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SortOrder) IsValid() bool {
	for _, existing := range AllowedSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SortOrder value
func (v SortOrder) Ptr() *SortOrder {
	return &v
}

type NullableSortOrder struct {
	value *SortOrder
	isSet bool
}

func (v NullableSortOrder) Get() *SortOrder {
	return v.value
}

func (v *NullableSortOrder) Set(val *SortOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableSortOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableSortOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortOrder(val *SortOrder) *NullableSortOrder {
	return &NullableSortOrder{value: val, isSet: true}
}

func (v NullableSortOrder) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSortOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
