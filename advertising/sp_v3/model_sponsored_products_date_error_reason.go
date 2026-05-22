package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsDateErrorReason the model 'SponsoredProductsDateErrorReason'
type SponsoredProductsDateErrorReason string

// List of SponsoredProductsDateErrorReason
const (
	SPONSOREDPRODUCTSDATEERRORREASON_INVALID_DATE                              SponsoredProductsDateErrorReason = "INVALID_DATE"
	SPONSOREDPRODUCTSDATEERRORREASON_START_DATE_EARLIER_THAN_TODAY             SponsoredProductsDateErrorReason = "START_DATE_EARLIER_THAN_TODAY"
	SPONSOREDPRODUCTSDATEERRORREASON_END_DATE_EARLIER_THAN_TODAY               SponsoredProductsDateErrorReason = "END_DATE_EARLIER_THAN_TODAY"
	SPONSOREDPRODUCTSDATEERRORREASON_START_DATE_LATER_THAN_MAXIMUM             SponsoredProductsDateErrorReason = "START_DATE_LATER_THAN_MAXIMUM"
	SPONSOREDPRODUCTSDATEERRORREASON_END_DATE_LATER_THAN_MAXIMUM               SponsoredProductsDateErrorReason = "END_DATE_LATER_THAN_MAXIMUM"
	SPONSOREDPRODUCTSDATEERRORREASON_START_DATE_AFTER_END_DATE                 SponsoredProductsDateErrorReason = "START_DATE_AFTER_END_DATE"
	SPONSOREDPRODUCTSDATEERRORREASON_UPDATING_READ_ONLY_START_DATE             SponsoredProductsDateErrorReason = "UPDATING_READ_ONLY_START_DATE"
	SPONSOREDPRODUCTSDATEERRORREASON_UPDATING_READ_ONLY_END_DATE               SponsoredProductsDateErrorReason = "UPDATING_READ_ONLY_END_DATE"
	SPONSOREDPRODUCTSDATEERRORREASON_UPDATING_ENDED_CAMPAIGN_WITHOUT_EXTENSION SponsoredProductsDateErrorReason = "UPDATING_ENDED_CAMPAIGN_WITHOUT_EXTENSION"
)

// All allowed values of SponsoredProductsDateErrorReason enum
var AllowedSponsoredProductsDateErrorReasonEnumValues = []SponsoredProductsDateErrorReason{
	"INVALID_DATE",
	"START_DATE_EARLIER_THAN_TODAY",
	"END_DATE_EARLIER_THAN_TODAY",
	"START_DATE_LATER_THAN_MAXIMUM",
	"END_DATE_LATER_THAN_MAXIMUM",
	"START_DATE_AFTER_END_DATE",
	"UPDATING_READ_ONLY_START_DATE",
	"UPDATING_READ_ONLY_END_DATE",
	"UPDATING_ENDED_CAMPAIGN_WITHOUT_EXTENSION",
}

func (v *SponsoredProductsDateErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsDateErrorReason(value)
	for _, existing := range AllowedSponsoredProductsDateErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsDateErrorReason", value)
}

// NewSponsoredProductsDateErrorReasonFromValue returns a pointer to a valid SponsoredProductsDateErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsDateErrorReasonFromValue(v string) (*SponsoredProductsDateErrorReason, error) {
	ev := SponsoredProductsDateErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsDateErrorReason: valid values are %v", v, AllowedSponsoredProductsDateErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsDateErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsDateErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsDateErrorReason value
func (v SponsoredProductsDateErrorReason) Ptr() *SponsoredProductsDateErrorReason {
	return &v
}

type NullableSponsoredProductsDateErrorReason struct {
	value *SponsoredProductsDateErrorReason
	isSet bool
}

func (v NullableSponsoredProductsDateErrorReason) Get() *SponsoredProductsDateErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsDateErrorReason) Set(val *SponsoredProductsDateErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDateErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDateErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDateErrorReason(val *SponsoredProductsDateErrorReason) *NullableSponsoredProductsDateErrorReason {
	return &NullableSponsoredProductsDateErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsDateErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDateErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
