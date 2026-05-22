package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsTargetingClauseSetupErrorReason the model 'SponsoredProductsTargetingClauseSetupErrorReason'
type SponsoredProductsTargetingClauseSetupErrorReason string

// List of SponsoredProductsTargetingClauseSetupErrorReason
const (
	SPONSOREDPRODUCTSTARGETINGCLAUSESETUPERRORREASON_TARGETING_TYPE_NOT_ALLOWED_FOR_AUTO_TARGETING_CAMPAIGN SponsoredProductsTargetingClauseSetupErrorReason = "TARGETING_TYPE_NOT_ALLOWED_FOR_AUTO_TARGETING_CAMPAIGN"
	SPONSOREDPRODUCTSTARGETINGCLAUSESETUPERRORREASON_TYPE_CONFLICT_IN_AD_GROUP                              SponsoredProductsTargetingClauseSetupErrorReason = "TYPE_CONFLICT_IN_AD_GROUP"
	SPONSOREDPRODUCTSTARGETINGCLAUSESETUPERRORREASON_AUTO_TARGETING_CLAUSE_CANNOT_BE_CREATED_MANUALLY       SponsoredProductsTargetingClauseSetupErrorReason = "AUTO_TARGETING_CLAUSE_CANNOT_BE_CREATED_MANUALLY"
	SPONSOREDPRODUCTSTARGETINGCLAUSESETUPERRORREASON_TARGETING_EXPRESSION_INVALID_VALUE                     SponsoredProductsTargetingClauseSetupErrorReason = "TARGETING_EXPRESSION_INVALID_VALUE"
)

// All allowed values of SponsoredProductsTargetingClauseSetupErrorReason enum
var AllowedSponsoredProductsTargetingClauseSetupErrorReasonEnumValues = []SponsoredProductsTargetingClauseSetupErrorReason{
	"TARGETING_TYPE_NOT_ALLOWED_FOR_AUTO_TARGETING_CAMPAIGN",
	"TYPE_CONFLICT_IN_AD_GROUP",
	"AUTO_TARGETING_CLAUSE_CANNOT_BE_CREATED_MANUALLY",
	"TARGETING_EXPRESSION_INVALID_VALUE",
}

func (v *SponsoredProductsTargetingClauseSetupErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsTargetingClauseSetupErrorReason(value)
	for _, existing := range AllowedSponsoredProductsTargetingClauseSetupErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsTargetingClauseSetupErrorReason", value)
}

// NewSponsoredProductsTargetingClauseSetupErrorReasonFromValue returns a pointer to a valid SponsoredProductsTargetingClauseSetupErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsTargetingClauseSetupErrorReasonFromValue(v string) (*SponsoredProductsTargetingClauseSetupErrorReason, error) {
	ev := SponsoredProductsTargetingClauseSetupErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsTargetingClauseSetupErrorReason: valid values are %v", v, AllowedSponsoredProductsTargetingClauseSetupErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsTargetingClauseSetupErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsTargetingClauseSetupErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsTargetingClauseSetupErrorReason value
func (v SponsoredProductsTargetingClauseSetupErrorReason) Ptr() *SponsoredProductsTargetingClauseSetupErrorReason {
	return &v
}

type NullableSponsoredProductsTargetingClauseSetupErrorReason struct {
	value *SponsoredProductsTargetingClauseSetupErrorReason
	isSet bool
}

func (v NullableSponsoredProductsTargetingClauseSetupErrorReason) Get() *SponsoredProductsTargetingClauseSetupErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsTargetingClauseSetupErrorReason) Set(val *SponsoredProductsTargetingClauseSetupErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetingClauseSetupErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetingClauseSetupErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetingClauseSetupErrorReason(val *SponsoredProductsTargetingClauseSetupErrorReason) *NullableSponsoredProductsTargetingClauseSetupErrorReason {
	return &NullableSponsoredProductsTargetingClauseSetupErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetingClauseSetupErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetingClauseSetupErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
