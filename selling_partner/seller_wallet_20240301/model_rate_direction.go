package seller_wallet_20240301

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// RateDirection Whether the customer is buying or selling the source currency.
type RateDirection string

// List of RateDirection
const (
	RATEDIRECTION_BUY  RateDirection = "BUY"
	RATEDIRECTION_SELL RateDirection = "SELL"
)

// All allowed values of RateDirection enum
var AllowedRateDirectionEnumValues = []RateDirection{
	RATEDIRECTION_BUY,
	RATEDIRECTION_SELL,
}

func (v *RateDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = RateDirection(value)
	return nil
}

// NewRateDirectionFromValue returns a pointer to a valid RateDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRateDirectionFromValue(v string) (*RateDirection, error) {
	ev := RateDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RateDirection: valid values are %v", v, AllowedRateDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RateDirection) IsValid() bool {
	for _, existing := range AllowedRateDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RateDirection value
func (v RateDirection) Ptr() *RateDirection {
	return &v
}

type NullableRateDirection struct {
	value *RateDirection
	isSet bool
}

func (v NullableRateDirection) Get() *RateDirection {
	return v.value
}

func (v *NullableRateDirection) Set(val *RateDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableRateDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableRateDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateDirection(val *RateDirection) *NullableRateDirection {
	return &NullableRateDirection{value: val, isSet: true}
}

func (v NullableRateDirection) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRateDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
