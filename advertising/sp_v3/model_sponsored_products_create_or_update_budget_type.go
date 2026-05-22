package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsCreateOrUpdateBudgetType the model 'SponsoredProductsCreateOrUpdateBudgetType'
type SponsoredProductsCreateOrUpdateBudgetType string

// List of SponsoredProductsCreateOrUpdateBudgetType
const (
	SPONSOREDPRODUCTSCREATEORUPDATEBUDGETTYPE_DAILY SponsoredProductsCreateOrUpdateBudgetType = "DAILY"
)

// All allowed values of SponsoredProductsCreateOrUpdateBudgetType enum
var AllowedSponsoredProductsCreateOrUpdateBudgetTypeEnumValues = []SponsoredProductsCreateOrUpdateBudgetType{
	"DAILY",
}

func (v *SponsoredProductsCreateOrUpdateBudgetType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsCreateOrUpdateBudgetType(value)
	for _, existing := range AllowedSponsoredProductsCreateOrUpdateBudgetTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsCreateOrUpdateBudgetType", value)
}

// NewSponsoredProductsCreateOrUpdateBudgetTypeFromValue returns a pointer to a valid SponsoredProductsCreateOrUpdateBudgetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsCreateOrUpdateBudgetTypeFromValue(v string) (*SponsoredProductsCreateOrUpdateBudgetType, error) {
	ev := SponsoredProductsCreateOrUpdateBudgetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsCreateOrUpdateBudgetType: valid values are %v", v, AllowedSponsoredProductsCreateOrUpdateBudgetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsCreateOrUpdateBudgetType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsCreateOrUpdateBudgetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsCreateOrUpdateBudgetType value
func (v SponsoredProductsCreateOrUpdateBudgetType) Ptr() *SponsoredProductsCreateOrUpdateBudgetType {
	return &v
}

type NullableSponsoredProductsCreateOrUpdateBudgetType struct {
	value *SponsoredProductsCreateOrUpdateBudgetType
	isSet bool
}

func (v NullableSponsoredProductsCreateOrUpdateBudgetType) Get() *SponsoredProductsCreateOrUpdateBudgetType {
	return v.value
}

func (v *NullableSponsoredProductsCreateOrUpdateBudgetType) Set(val *SponsoredProductsCreateOrUpdateBudgetType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateOrUpdateBudgetType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateOrUpdateBudgetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateOrUpdateBudgetType(val *SponsoredProductsCreateOrUpdateBudgetType) *NullableSponsoredProductsCreateOrUpdateBudgetType {
	return &NullableSponsoredProductsCreateOrUpdateBudgetType{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateOrUpdateBudgetType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateOrUpdateBudgetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
