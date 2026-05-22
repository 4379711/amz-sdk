package portfolios_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// PortfolioEntityType the model 'PortfolioEntityType'
type PortfolioEntityType string

// List of PortfolioEntityType
const (
	PORTFOLIOENTITYTYPE_PORTFOLIO PortfolioEntityType = "PORTFOLIO"
)

// All allowed values of PortfolioEntityType enum
var AllowedPortfolioEntityTypeEnumValues = []PortfolioEntityType{
	"PORTFOLIO",
}

func (v *PortfolioEntityType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortfolioEntityType(value)
	for _, existing := range AllowedPortfolioEntityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortfolioEntityType", value)
}

// NewPortfolioEntityTypeFromValue returns a pointer to a valid PortfolioEntityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortfolioEntityTypeFromValue(v string) (*PortfolioEntityType, error) {
	ev := PortfolioEntityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortfolioEntityType: valid values are %v", v, AllowedPortfolioEntityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortfolioEntityType) IsValid() bool {
	for _, existing := range AllowedPortfolioEntityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortfolioEntityType value
func (v PortfolioEntityType) Ptr() *PortfolioEntityType {
	return &v
}

type NullablePortfolioEntityType struct {
	value *PortfolioEntityType
	isSet bool
}

func (v NullablePortfolioEntityType) Get() *PortfolioEntityType {
	return v.value
}

func (v *NullablePortfolioEntityType) Set(val *PortfolioEntityType) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioEntityType) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioEntityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioEntityType(val *PortfolioEntityType) *NullablePortfolioEntityType {
	return &NullablePortfolioEntityType{value: val, isSet: true}
}

func (v NullablePortfolioEntityType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioEntityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
