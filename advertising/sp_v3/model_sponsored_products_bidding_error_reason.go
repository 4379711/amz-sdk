package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsBiddingErrorReason the model 'SponsoredProductsBiddingErrorReason'
type SponsoredProductsBiddingErrorReason string

// List of SponsoredProductsBiddingErrorReason
const (
	SPONSOREDPRODUCTSBIDDINGERRORREASON_GT_BUDGET                         SponsoredProductsBiddingErrorReason = "BID_GT_BUDGET"
	SPONSOREDPRODUCTSBIDDINGERRORREASON_OUT_OF_MARKET_PLACE_RANGE         SponsoredProductsBiddingErrorReason = "BID_OUT_OF_MARKET_PLACE_RANGE"
	SPONSOREDPRODUCTSBIDDINGERRORREASON_INVALID_PLACEMENT                 SponsoredProductsBiddingErrorReason = "BID_INVALID_PLACEMENT"
	SPONSOREDPRODUCTSBIDDINGERRORREASON_INVALID_SHOPPER_COHORT_TYPE       SponsoredProductsBiddingErrorReason = "BID_INVALID_SHOPPER_COHORT_TYPE"
	SPONSOREDPRODUCTSBIDDINGERRORREASON_SHOPPER_COHORTS_MORE_THAN_ALLOWED SponsoredProductsBiddingErrorReason = "BID_SHOPPER_COHORTS_MORE_THAN_ALLOWED"
	SPONSOREDPRODUCTSBIDDINGERRORREASON_MISSING_AUDIENCES                 SponsoredProductsBiddingErrorReason = "BID_MISSING_AUDIENCES"
	SPONSOREDPRODUCTSBIDDINGERRORREASON_AUDIENCES_MORE_THAN_ALLOWED       SponsoredProductsBiddingErrorReason = "BID_AUDIENCES_MORE_THAN_ALLOWED"
	SPONSOREDPRODUCTSBIDDINGERRORREASON_INVALID_AUDIENCE_SEGMENT_TYPE     SponsoredProductsBiddingErrorReason = "BID_INVALID_AUDIENCE_SEGMENT_TYPE"
	SPONSOREDPRODUCTSBIDDINGERRORREASON_INVALID_AUDIENCE_ID               SponsoredProductsBiddingErrorReason = "BID_INVALID_AUDIENCE_ID"
)

// All allowed values of SponsoredProductsBiddingErrorReason enum
var AllowedSponsoredProductsBiddingErrorReasonEnumValues = []SponsoredProductsBiddingErrorReason{
	"BID_GT_BUDGET",
	"BID_OUT_OF_MARKET_PLACE_RANGE",
	"BID_INVALID_PLACEMENT",
	"BID_INVALID_SHOPPER_COHORT_TYPE",
	"BID_SHOPPER_COHORTS_MORE_THAN_ALLOWED",
	"BID_MISSING_AUDIENCES",
	"BID_AUDIENCES_MORE_THAN_ALLOWED",
	"BID_INVALID_AUDIENCE_SEGMENT_TYPE",
	"BID_INVALID_AUDIENCE_ID",
}

func (v *SponsoredProductsBiddingErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsBiddingErrorReason(value)
	for _, existing := range AllowedSponsoredProductsBiddingErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsBiddingErrorReason", value)
}

// NewSponsoredProductsBiddingErrorReasonFromValue returns a pointer to a valid SponsoredProductsBiddingErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsBiddingErrorReasonFromValue(v string) (*SponsoredProductsBiddingErrorReason, error) {
	ev := SponsoredProductsBiddingErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsBiddingErrorReason: valid values are %v", v, AllowedSponsoredProductsBiddingErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsBiddingErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsBiddingErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsBiddingErrorReason value
func (v SponsoredProductsBiddingErrorReason) Ptr() *SponsoredProductsBiddingErrorReason {
	return &v
}

type NullableSponsoredProductsBiddingErrorReason struct {
	value *SponsoredProductsBiddingErrorReason
	isSet bool
}

func (v NullableSponsoredProductsBiddingErrorReason) Get() *SponsoredProductsBiddingErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsBiddingErrorReason) Set(val *SponsoredProductsBiddingErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBiddingErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBiddingErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBiddingErrorReason(val *SponsoredProductsBiddingErrorReason) *NullableSponsoredProductsBiddingErrorReason {
	return &NullableSponsoredProductsBiddingErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsBiddingErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBiddingErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
