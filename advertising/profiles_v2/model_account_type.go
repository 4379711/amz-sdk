package profiles_v2

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// AccountType The `seller` and `vendor` account types are associated with Sponsored Ads APIs. The `agency` account type is associated with DSP and Data Provider APIs.
type AccountType string

// List of AccountType
const (
	ACCOUNTTYPE_VENDOR AccountType = "vendor"
	ACCOUNTTYPE_SELLER AccountType = "seller"
	ACCOUNTTYPE_AGENCY AccountType = "agency"
)

// All allowed values of AccountType enum
var AllowedAccountTypeEnumValues = []AccountType{
	"vendor",
	"seller",
	"agency",
}

func (v *AccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountType(value)
	for _, existing := range AllowedAccountTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountType", value)
}

// NewAccountTypeFromValue returns a pointer to a valid AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountTypeFromValue(v string) (*AccountType, error) {
	ev := AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountType: valid values are %v", v, AllowedAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountType) IsValid() bool {
	for _, existing := range AllowedAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountType value
func (v AccountType) Ptr() *AccountType {
	return &v
}

type NullableAccountType struct {
	value *AccountType
	isSet bool
}

func (v NullableAccountType) Get() *AccountType {
	return v.value
}

func (v *NullableAccountType) Set(val *AccountType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountType(val *AccountType) *NullableAccountType {
	return &NullableAccountType{value: val, isSet: true}
}

func (v NullableAccountType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAccountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
