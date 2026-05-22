package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsBatchErrorCode Per-item batch error codes. | Error Code |  description | |-----------|------------| | `BAD_REQUEST` | Non-retryable error that indicates the requester must take some corrective action before the request will succeed. | | `TOO_MANY_REQUESTS` | The requester has sent too many requests. Please retry later and distribute requests over larger time intervals. | | `INTERNAL_ERROR` | An internal error has occurred. Please contact customer support if the issue persists. |
type SponsoredProductsBatchErrorCode string

// List of SponsoredProductsBatchErrorCode
const (
	SPONSOREDPRODUCTSBATCHERRORCODE_BAD_REQUEST           SponsoredProductsBatchErrorCode = "BAD_REQUEST"
	SPONSOREDPRODUCTSBATCHERRORCODE_TOO_MANY_REQUESTS     SponsoredProductsBatchErrorCode = "TOO_MANY_REQUESTS"
	SPONSOREDPRODUCTSBATCHERRORCODE_INTERNAL_SERVER_ERROR SponsoredProductsBatchErrorCode = "INTERNAL_SERVER_ERROR"
)

// All allowed values of SponsoredProductsBatchErrorCode enum
var AllowedSponsoredProductsBatchErrorCodeEnumValues = []SponsoredProductsBatchErrorCode{
	"BAD_REQUEST",
	"TOO_MANY_REQUESTS",
	"INTERNAL_SERVER_ERROR",
}

func (v *SponsoredProductsBatchErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsBatchErrorCode(value)
	for _, existing := range AllowedSponsoredProductsBatchErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsBatchErrorCode", value)
}

// NewSponsoredProductsBatchErrorCodeFromValue returns a pointer to a valid SponsoredProductsBatchErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsBatchErrorCodeFromValue(v string) (*SponsoredProductsBatchErrorCode, error) {
	ev := SponsoredProductsBatchErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsBatchErrorCode: valid values are %v", v, AllowedSponsoredProductsBatchErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsBatchErrorCode) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsBatchErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsBatchErrorCode value
func (v SponsoredProductsBatchErrorCode) Ptr() *SponsoredProductsBatchErrorCode {
	return &v
}

type NullableSponsoredProductsBatchErrorCode struct {
	value *SponsoredProductsBatchErrorCode
	isSet bool
}

func (v NullableSponsoredProductsBatchErrorCode) Get() *SponsoredProductsBatchErrorCode {
	return v.value
}

func (v *NullableSponsoredProductsBatchErrorCode) Set(val *SponsoredProductsBatchErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBatchErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBatchErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBatchErrorCode(val *SponsoredProductsBatchErrorCode) *NullableSponsoredProductsBatchErrorCode {
	return &NullableSponsoredProductsBatchErrorCode{value: val, isSet: true}
}

func (v NullableSponsoredProductsBatchErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBatchErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
