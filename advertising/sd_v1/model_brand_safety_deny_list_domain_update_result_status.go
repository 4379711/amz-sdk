package sd_v1

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// BrandSafetyDenyListDomainUpdateResultStatus The state of the domain.
type BrandSafetyDenyListDomainUpdateResultStatus string

// List of BrandSafetyDenyListDomainUpdateResultStatus
const (
	BRANDSAFETYDENYLISTDOMAINUPDATERESULTSTATUS_SUCCESS BrandSafetyDenyListDomainUpdateResultStatus = "SUCCESS"
	BRANDSAFETYDENYLISTDOMAINUPDATERESULTSTATUS_FAILURE BrandSafetyDenyListDomainUpdateResultStatus = "FAILURE"
)

// All allowed values of BrandSafetyDenyListDomainUpdateResultStatus enum
var AllowedBrandSafetyDenyListDomainUpdateResultStatusEnumValues = []BrandSafetyDenyListDomainUpdateResultStatus{
	"SUCCESS",
	"FAILURE",
}

func (v *BrandSafetyDenyListDomainUpdateResultStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BrandSafetyDenyListDomainUpdateResultStatus(value)
	for _, existing := range AllowedBrandSafetyDenyListDomainUpdateResultStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BrandSafetyDenyListDomainUpdateResultStatus", value)
}

// NewBrandSafetyDenyListDomainUpdateResultStatusFromValue returns a pointer to a valid BrandSafetyDenyListDomainUpdateResultStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandSafetyDenyListDomainUpdateResultStatusFromValue(v string) (*BrandSafetyDenyListDomainUpdateResultStatus, error) {
	ev := BrandSafetyDenyListDomainUpdateResultStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandSafetyDenyListDomainUpdateResultStatus: valid values are %v", v, AllowedBrandSafetyDenyListDomainUpdateResultStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandSafetyDenyListDomainUpdateResultStatus) IsValid() bool {
	for _, existing := range AllowedBrandSafetyDenyListDomainUpdateResultStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BrandSafetyDenyListDomainUpdateResultStatus value
func (v BrandSafetyDenyListDomainUpdateResultStatus) Ptr() *BrandSafetyDenyListDomainUpdateResultStatus {
	return &v
}

type NullableBrandSafetyDenyListDomainUpdateResultStatus struct {
	value *BrandSafetyDenyListDomainUpdateResultStatus
	isSet bool
}

func (v NullableBrandSafetyDenyListDomainUpdateResultStatus) Get() *BrandSafetyDenyListDomainUpdateResultStatus {
	return v.value
}

func (v *NullableBrandSafetyDenyListDomainUpdateResultStatus) Set(val *BrandSafetyDenyListDomainUpdateResultStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandSafetyDenyListDomainUpdateResultStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandSafetyDenyListDomainUpdateResultStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandSafetyDenyListDomainUpdateResultStatus(val *BrandSafetyDenyListDomainUpdateResultStatus) *NullableBrandSafetyDenyListDomainUpdateResultStatus {
	return &NullableBrandSafetyDenyListDomainUpdateResultStatus{value: val, isSet: true}
}

func (v NullableBrandSafetyDenyListDomainUpdateResultStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandSafetyDenyListDomainUpdateResultStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
