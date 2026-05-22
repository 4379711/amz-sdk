package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsEntityStateErrorReason the model 'SponsoredProductsEntityStateErrorReason'
type SponsoredProductsEntityStateErrorReason string

// List of SponsoredProductsEntityStateErrorReason
const (
	SPONSOREDPRODUCTSENTITYSTATEERRORREASON_INVALID_TARGET_STATE                              SponsoredProductsEntityStateErrorReason = "INVALID_TARGET_STATE"
	SPONSOREDPRODUCTSENTITYSTATEERRORREASON_PARENT_ENTITY_FORBIDS_CREATION                    SponsoredProductsEntityStateErrorReason = "PARENT_ENTITY_FORBIDS_CREATION"
	SPONSOREDPRODUCTSENTITYSTATEERRORREASON_MARKETPLACE_STATE_CANNOT_BE_ARCHIVED              SponsoredProductsEntityStateErrorReason = "MARKETPLACE_STATE_CANNOT_BE_ARCHIVED"
	SPONSOREDPRODUCTSENTITYSTATEERRORREASON_AUTO_TARGETING_CLAUSE_CANNOT_BE_ARCHIVED_MANUALLY SponsoredProductsEntityStateErrorReason = "AUTO_TARGETING_CLAUSE_CANNOT_BE_ARCHIVED_MANUALLY"
	SPONSOREDPRODUCTSENTITYSTATEERRORREASON_ARCHIVED_ENTITY_CANNOT_BE_MODIFIED                SponsoredProductsEntityStateErrorReason = "ARCHIVED_ENTITY_CANNOT_BE_MODIFIED"
	SPONSOREDPRODUCTSENTITYSTATEERRORREASON_PARENT_ARCHIVED_FORBIDS_UPDATES                   SponsoredProductsEntityStateErrorReason = "PARENT_ARCHIVED_FORBIDS_UPDATES"
	SPONSOREDPRODUCTSENTITYSTATEERRORREASON_PARENT_STATUS_FORBIDS_UPDATES_AND_CREATES         SponsoredProductsEntityStateErrorReason = "PARENT_STATUS_FORBIDS_UPDATES_AND_CREATES"
)

// All allowed values of SponsoredProductsEntityStateErrorReason enum
var AllowedSponsoredProductsEntityStateErrorReasonEnumValues = []SponsoredProductsEntityStateErrorReason{
	"INVALID_TARGET_STATE",
	"PARENT_ENTITY_FORBIDS_CREATION",
	"MARKETPLACE_STATE_CANNOT_BE_ARCHIVED",
	"AUTO_TARGETING_CLAUSE_CANNOT_BE_ARCHIVED_MANUALLY",
	"ARCHIVED_ENTITY_CANNOT_BE_MODIFIED",
	"PARENT_ARCHIVED_FORBIDS_UPDATES",
	"PARENT_STATUS_FORBIDS_UPDATES_AND_CREATES",
}

func (v *SponsoredProductsEntityStateErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsEntityStateErrorReason(value)
	for _, existing := range AllowedSponsoredProductsEntityStateErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsEntityStateErrorReason", value)
}

// NewSponsoredProductsEntityStateErrorReasonFromValue returns a pointer to a valid SponsoredProductsEntityStateErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsEntityStateErrorReasonFromValue(v string) (*SponsoredProductsEntityStateErrorReason, error) {
	ev := SponsoredProductsEntityStateErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsEntityStateErrorReason: valid values are %v", v, AllowedSponsoredProductsEntityStateErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsEntityStateErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsEntityStateErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsEntityStateErrorReason value
func (v SponsoredProductsEntityStateErrorReason) Ptr() *SponsoredProductsEntityStateErrorReason {
	return &v
}

type NullableSponsoredProductsEntityStateErrorReason struct {
	value *SponsoredProductsEntityStateErrorReason
	isSet bool
}

func (v NullableSponsoredProductsEntityStateErrorReason) Get() *SponsoredProductsEntityStateErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsEntityStateErrorReason) Set(val *SponsoredProductsEntityStateErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsEntityStateErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsEntityStateErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsEntityStateErrorReason(val *SponsoredProductsEntityStateErrorReason) *NullableSponsoredProductsEntityStateErrorReason {
	return &NullableSponsoredProductsEntityStateErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsEntityStateErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsEntityStateErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
