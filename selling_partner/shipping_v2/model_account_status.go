package shipping_v2

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// AccountStatus Account Status.
type AccountStatus string

// List of AccountStatus
const (
	ACCOUNTSTATUS_ACTIVE    AccountStatus = "ACTIVE"
	ACCOUNTSTATUS_INACTIVE  AccountStatus = "INACTIVE"
	ACCOUNTSTATUS_PENDING   AccountStatus = "PENDING"
	ACCOUNTSTATUS_SUSPENDED AccountStatus = "SUSPENDED"
)

// All allowed values of AccountStatus enum
var AllowedAccountStatusEnumValues = []AccountStatus{
	ACCOUNTSTATUS_ACTIVE,
	ACCOUNTSTATUS_INACTIVE,
	ACCOUNTSTATUS_PENDING,
	ACCOUNTSTATUS_SUSPENDED,
}

func (v *AccountStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountStatus(value)
	for _, existing := range AllowedAccountStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountStatus", value)
}

// NewAccountStatusFromValue returns a pointer to a valid AccountStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountStatusFromValue(v string) (*AccountStatus, error) {
	ev := AccountStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountStatus: valid values are %v", v, AllowedAccountStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountStatus) IsValid() bool {
	for _, existing := range AllowedAccountStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountStatus value
func (v AccountStatus) Ptr() *AccountStatus {
	return &v
}

type NullableAccountStatus struct {
	value *AccountStatus
	isSet bool
}

func (v NullableAccountStatus) Get() *AccountStatus {
	return v.value
}

func (v *NullableAccountStatus) Set(val *AccountStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountStatus(val *AccountStatus) *NullableAccountStatus {
	return &NullableAccountStatus{value: val, isSet: true}
}

func (v NullableAccountStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAccountStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
