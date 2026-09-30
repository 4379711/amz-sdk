package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ProgramType Asset program type
type ProgramType string

// List of ProgramType
const (
	PROGRAMTYPE_A_PLUS                   ProgramType = "A_PLUS"
	PROGRAMTYPE_SB                       ProgramType = "SB"
	PROGRAMTYPE_POSTS                    ProgramType = "POSTS"
	PROGRAMTYPE_STORES                   ProgramType = "STORES"
	PROGRAMTYPE_BBB_STORES               ProgramType = "BBB_STORES"
	PROGRAMTYPE_AMAZON_DSP               ProgramType = "AMAZON_DSP"
	PROGRAMTYPE_AMAZON_CREATIVE_SERVICES ProgramType = "AMAZON_CREATIVE_SERVICES"
)

// All allowed values of ProgramType enum
var AllowedProgramTypeEnumValues = []ProgramType{
	"A_PLUS",
	"SB",
	"POSTS",
	"STORES",
	"BBB_STORES",
	"AMAZON_DSP",
	"AMAZON_CREATIVE_SERVICES",
}

func (v *ProgramType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = ProgramType(value)
	return nil
}

// NewProgramTypeFromValue returns a pointer to a valid ProgramType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProgramTypeFromValue(v string) (*ProgramType, error) {
	ev := ProgramType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProgramType: valid values are %v", v, AllowedProgramTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProgramType) IsValid() bool {
	for _, existing := range AllowedProgramTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProgramType value
func (v ProgramType) Ptr() *ProgramType {
	return &v
}

type NullableProgramType struct {
	value *ProgramType
	isSet bool
}

func (v NullableProgramType) Get() *ProgramType {
	return v.value
}

func (v *NullableProgramType) Set(val *ProgramType) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramType) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramType(val *ProgramType) *NullableProgramType {
	return &NullableProgramType{value: val, isSet: true}
}

func (v NullableProgramType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProgramType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
