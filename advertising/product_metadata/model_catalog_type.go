package product_metadata

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// CatalogType Returned listing under AMAZON, BRAND, KDP or GLOBAL_STORE catalog
type CatalogType string

// List of catalogType
const (
	CATALOGTYPE_AMAZON       CatalogType = "AMAZON"
	CATALOGTYPE_BRAND        CatalogType = "BRAND"
	CATALOGTYPE_KDP          CatalogType = "KDP"
	CATALOGTYPE_GLOBAL_STORE CatalogType = "GLOBAL_STORE"
)

// All allowed values of CatalogType enum
var AllowedCatalogTypeEnumValues = []CatalogType{
	"AMAZON",
	"BRAND",
	"KDP",
	"GLOBAL_STORE",
}

func (v *CatalogType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CatalogType(value)
	for _, existing := range AllowedCatalogTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CatalogType", value)
}

// NewCatalogTypeFromValue returns a pointer to a valid CatalogType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCatalogTypeFromValue(v string) (*CatalogType, error) {
	ev := CatalogType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CatalogType: valid values are %v", v, AllowedCatalogTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CatalogType) IsValid() bool {
	for _, existing := range AllowedCatalogTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to catalogType value
func (v CatalogType) Ptr() *CatalogType {
	return &v
}

type NullableCatalogType struct {
	value *CatalogType
	isSet bool
}

func (v NullableCatalogType) Get() *CatalogType {
	return v.value
}

func (v *NullableCatalogType) Set(val *CatalogType) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogType) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogType(val *CatalogType) *NullableCatalogType {
	return &NullableCatalogType{value: val, isSet: true}
}

func (v NullableCatalogType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCatalogType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
