package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsAudienceSegmentType This field specifies the type of audience used. | Value               | Description| |---------------------|------------| | `SPONSORED_ADS_AMC` | This type refers to the Audience Segments created in AMC for Sponsored Ads. See [AMC API](https://advertising.amazon.com/API/docs/en-us/amc-rba#tag/Rule-based-audience) for details on how to create AMC Audiences. Once the AMC Audiences are created, the Audience Ids can be retrieved using the Discovery API [ListTargetableEntities](https://advertising.amazon.com/API/docs/en-us/targetable-entities#operation/ListTargetableEntities) with parameters; `adProduct`=`SPONSORED_PRODUCTS`, `targetTypeFilter`=`AUDIENCE` and `pathsFilter` = `[[\"Audience Category\", \"Custom-built\", \"AMC\"]]`. Only the audiences retrieved using these filters are usable.|
type SponsoredProductsAudienceSegmentType string

// List of SponsoredProductsAudienceSegmentType
const (
	SPONSOREDPRODUCTSAUDIENCESEGMENTTYPE_SPONSORED_ADS_AMC SponsoredProductsAudienceSegmentType = "SPONSORED_ADS_AMC"
)

// All allowed values of SponsoredProductsAudienceSegmentType enum
var AllowedSponsoredProductsAudienceSegmentTypeEnumValues = []SponsoredProductsAudienceSegmentType{
	"SPONSORED_ADS_AMC",
	"BEHAVIOR_DYNAMIC",
}

func (v *SponsoredProductsAudienceSegmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsAudienceSegmentType(value)
	for _, existing := range AllowedSponsoredProductsAudienceSegmentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsAudienceSegmentType", value)
}

// NewSponsoredProductsAudienceSegmentTypeFromValue returns a pointer to a valid SponsoredProductsAudienceSegmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsAudienceSegmentTypeFromValue(v string) (*SponsoredProductsAudienceSegmentType, error) {
	ev := SponsoredProductsAudienceSegmentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsAudienceSegmentType: valid values are %v", v, AllowedSponsoredProductsAudienceSegmentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsAudienceSegmentType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsAudienceSegmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsAudienceSegmentType value
func (v SponsoredProductsAudienceSegmentType) Ptr() *SponsoredProductsAudienceSegmentType {
	return &v
}

type NullableSponsoredProductsAudienceSegmentType struct {
	value *SponsoredProductsAudienceSegmentType
	isSet bool
}

func (v NullableSponsoredProductsAudienceSegmentType) Get() *SponsoredProductsAudienceSegmentType {
	return v.value
}

func (v *NullableSponsoredProductsAudienceSegmentType) Set(val *SponsoredProductsAudienceSegmentType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAudienceSegmentType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAudienceSegmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAudienceSegmentType(val *SponsoredProductsAudienceSegmentType) *NullableSponsoredProductsAudienceSegmentType {
	return &NullableSponsoredProductsAudienceSegmentType{value: val, isSet: true}
}

func (v NullableSponsoredProductsAudienceSegmentType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAudienceSegmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
