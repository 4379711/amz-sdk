package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsEntityType the model 'SponsoredProductsEntityType'
type SponsoredProductsEntityType string

// List of SponsoredProductsEntityType
const (
	SPONSOREDPRODUCTSENTITYTYPE_CAMPAIGN                           SponsoredProductsEntityType = "CAMPAIGN"
	SPONSOREDPRODUCTSENTITYTYPE_AD_GROUP                           SponsoredProductsEntityType = "AD_GROUP"
	SPONSOREDPRODUCTSENTITYTYPE_KEYWORD                            SponsoredProductsEntityType = "KEYWORD"
	SPONSOREDPRODUCTSENTITYTYPE_PRODUCT_AD                         SponsoredProductsEntityType = "PRODUCT_AD"
	SPONSOREDPRODUCTSENTITYTYPE_CAMPAIGN_NEGATIVE_KEYWORD          SponsoredProductsEntityType = "CAMPAIGN_NEGATIVE_KEYWORD"
	SPONSOREDPRODUCTSENTITYTYPE_NEGATIVE_KEYWORD                   SponsoredProductsEntityType = "NEGATIVE_KEYWORD"
	SPONSOREDPRODUCTSENTITYTYPE_TARGETING_CLAUSE                   SponsoredProductsEntityType = "TARGETING_CLAUSE"
	SPONSOREDPRODUCTSENTITYTYPE_NEGATIVE_TARGETING_CLAUSE          SponsoredProductsEntityType = "NEGATIVE_TARGETING_CLAUSE"
	SPONSOREDPRODUCTSENTITYTYPE_CAMPAIGN_NEGATIVE_TARGETING_CLAUSE SponsoredProductsEntityType = "CAMPAIGN_NEGATIVE_TARGETING_CLAUSE"
)

// All allowed values of SponsoredProductsEntityType enum
var AllowedSponsoredProductsEntityTypeEnumValues = []SponsoredProductsEntityType{
	"CAMPAIGN",
	"AD_GROUP",
	"KEYWORD",
	"PRODUCT_AD",
	"CAMPAIGN_NEGATIVE_KEYWORD",
	"NEGATIVE_KEYWORD",
	"TARGETING_CLAUSE",
	"NEGATIVE_TARGETING_CLAUSE",
	"CAMPAIGN_NEGATIVE_TARGETING_CLAUSE",
}

func (v *SponsoredProductsEntityType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsEntityType(value)
	for _, existing := range AllowedSponsoredProductsEntityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsEntityType", value)
}

// NewSponsoredProductsEntityTypeFromValue returns a pointer to a valid SponsoredProductsEntityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsEntityTypeFromValue(v string) (*SponsoredProductsEntityType, error) {
	ev := SponsoredProductsEntityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsEntityType: valid values are %v", v, AllowedSponsoredProductsEntityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsEntityType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsEntityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsEntityType value
func (v SponsoredProductsEntityType) Ptr() *SponsoredProductsEntityType {
	return &v
}

type NullableSponsoredProductsEntityType struct {
	value *SponsoredProductsEntityType
	isSet bool
}

func (v NullableSponsoredProductsEntityType) Get() *SponsoredProductsEntityType {
	return v.value
}

func (v *NullableSponsoredProductsEntityType) Set(val *SponsoredProductsEntityType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsEntityType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsEntityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsEntityType(val *SponsoredProductsEntityType) *NullableSponsoredProductsEntityType {
	return &NullableSponsoredProductsEntityType{value: val, isSet: true}
}

func (v NullableSponsoredProductsEntityType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsEntityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
