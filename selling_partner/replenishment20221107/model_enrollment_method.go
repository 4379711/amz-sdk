package replenishment20221107

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// EnrollmentMethod The enrollment method used to enroll the offer into the program.
type EnrollmentMethod string

// List of EnrollmentMethod
const (
	ENROLLMENTMETHOD_MANUAL    EnrollmentMethod = "MANUAL"
	ENROLLMENTMETHOD_AUTOMATIC EnrollmentMethod = "AUTOMATIC"
)

// All allowed values of EnrollmentMethod enum
var AllowedEnrollmentMethodEnumValues = []EnrollmentMethod{
	ENROLLMENTMETHOD_MANUAL,
	ENROLLMENTMETHOD_AUTOMATIC,
}

func (v *EnrollmentMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = EnrollmentMethod(value)
	return nil
}

// NewEnrollmentMethodFromValue returns a pointer to a valid EnrollmentMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnrollmentMethodFromValue(v string) (*EnrollmentMethod, error) {
	ev := EnrollmentMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnrollmentMethod: valid values are %v", v, AllowedEnrollmentMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnrollmentMethod) IsValid() bool {
	for _, existing := range AllowedEnrollmentMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnrollmentMethod value
func (v EnrollmentMethod) Ptr() *EnrollmentMethod {
	return &v
}

type NullableEnrollmentMethod struct {
	value *EnrollmentMethod
	isSet bool
}

func (v NullableEnrollmentMethod) Get() *EnrollmentMethod {
	return v.value
}

func (v *NullableEnrollmentMethod) Set(val *EnrollmentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentMethod(val *EnrollmentMethod) *NullableEnrollmentMethod {
	return &NullableEnrollmentMethod{value: val, isSet: true}
}

func (v NullableEnrollmentMethod) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableEnrollmentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
