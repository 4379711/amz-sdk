package sd_v1

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// BrandSafetyDenyListDomainState The state of the domain.
type BrandSafetyDenyListDomainState string

// List of BrandSafetyDenyListDomainState
const (
	BRANDSAFETYDENYLISTDOMAINSTATE_ENABLED  BrandSafetyDenyListDomainState = "ENABLED"
	BRANDSAFETYDENYLISTDOMAINSTATE_ARCHIVED BrandSafetyDenyListDomainState = "ARCHIVED"
)

// All allowed values of BrandSafetyDenyListDomainState enum
var AllowedBrandSafetyDenyListDomainStateEnumValues = []BrandSafetyDenyListDomainState{
	"ENABLED",
	"ARCHIVED",
}

func (v *BrandSafetyDenyListDomainState) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BrandSafetyDenyListDomainState(value)
	for _, existing := range AllowedBrandSafetyDenyListDomainStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BrandSafetyDenyListDomainState", value)
}

// NewBrandSafetyDenyListDomainStateFromValue returns a pointer to a valid BrandSafetyDenyListDomainState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandSafetyDenyListDomainStateFromValue(v string) (*BrandSafetyDenyListDomainState, error) {
	ev := BrandSafetyDenyListDomainState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandSafetyDenyListDomainState: valid values are %v", v, AllowedBrandSafetyDenyListDomainStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandSafetyDenyListDomainState) IsValid() bool {
	for _, existing := range AllowedBrandSafetyDenyListDomainStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BrandSafetyDenyListDomainState value
func (v BrandSafetyDenyListDomainState) Ptr() *BrandSafetyDenyListDomainState {
	return &v
}

type NullableBrandSafetyDenyListDomainState struct {
	value *BrandSafetyDenyListDomainState
	isSet bool
}

func (v NullableBrandSafetyDenyListDomainState) Get() *BrandSafetyDenyListDomainState {
	return v.value
}

func (v *NullableBrandSafetyDenyListDomainState) Set(val *BrandSafetyDenyListDomainState) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandSafetyDenyListDomainState) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandSafetyDenyListDomainState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandSafetyDenyListDomainState(val *BrandSafetyDenyListDomainState) *NullableBrandSafetyDenyListDomainState {
	return &NullableBrandSafetyDenyListDomainState{value: val, isSet: true}
}

func (v NullableBrandSafetyDenyListDomainState) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandSafetyDenyListDomainState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
