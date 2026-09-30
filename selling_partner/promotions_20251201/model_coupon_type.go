package promotions_20251201

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// CouponType The subtype of a `COUPON` promotion. Only applicable when `promotionType` is `COUPON`; not set for other promotion types. When omitted, the coupon is treated as `STANDARD`.
type CouponType string

// List of CouponType
const (
	COUPONTYPE_STANDARD           CouponType = "STANDARD"
	COUPONTYPE_SUBSCRIBE_AND_SAVE CouponType = "SUBSCRIBE_AND_SAVE"
	COUPONTYPE_REORDER_REWARDS    CouponType = "REORDER_REWARDS"
)

// All allowed values of CouponType enum
var AllowedCouponTypeEnumValues = []CouponType{
	"STANDARD",
	"SUBSCRIBE_AND_SAVE",
	"REORDER_REWARDS",
}

func (v *CouponType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = CouponType(value)
	return nil
}

// NewCouponTypeFromValue returns a pointer to a valid CouponType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCouponTypeFromValue(v string) (*CouponType, error) {
	ev := CouponType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CouponType: valid values are %v", v, AllowedCouponTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CouponType) IsValid() bool {
	for _, existing := range AllowedCouponTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CouponType value
func (v CouponType) Ptr() *CouponType {
	return &v
}

type NullableCouponType struct {
	value *CouponType
	isSet bool
}

func (v NullableCouponType) Get() *CouponType {
	return v.value
}

func (v *NullableCouponType) Set(val *CouponType) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponType) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponType(val *CouponType) *NullableCouponType {
	return &NullableCouponType{value: val, isSet: true}
}

func (v NullableCouponType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCouponType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
