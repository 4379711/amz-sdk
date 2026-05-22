package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsErrorCode Error codes. | Error Code |  description | |-----------|------------| | `BAD_REQUEST` | Received malformed request. | | `UNAUTHORIZED` | | | `FORBIDDEN` | | | `NOT_FOUND` | | | `CONTENT_TOO_LARGE` | | | `TOO_MANY_REQUESTS` | | | `INTERNAL_ERROR` | | | `BAD_GATEWAY` | | | `SERVICE_UNAVAILABLE` | | | `GATEWAY_TIMEOUT` | |
type SponsoredProductsErrorCode string

// List of SponsoredProductsErrorCode
const (
	SPONSOREDPRODUCTSERRORCODE_BAD_REQUEST           SponsoredProductsErrorCode = "BAD_REQUEST"
	SPONSOREDPRODUCTSERRORCODE_UNAUTHORIZED          SponsoredProductsErrorCode = "UNAUTHORIZED"
	SPONSOREDPRODUCTSERRORCODE_FORBIDDEN             SponsoredProductsErrorCode = "FORBIDDEN"
	SPONSOREDPRODUCTSERRORCODE_NOT_FOUND             SponsoredProductsErrorCode = "NOT_FOUND"
	SPONSOREDPRODUCTSERRORCODE_CONTENT_TOO_LARGE     SponsoredProductsErrorCode = "CONTENT_TOO_LARGE"
	SPONSOREDPRODUCTSERRORCODE_TOO_MANY_REQUESTS     SponsoredProductsErrorCode = "TOO_MANY_REQUESTS"
	SPONSOREDPRODUCTSERRORCODE_INTERNAL_SERVER_ERROR SponsoredProductsErrorCode = "INTERNAL_SERVER_ERROR"
	SPONSOREDPRODUCTSERRORCODE_BAD_GATEWAY           SponsoredProductsErrorCode = "BAD_GATEWAY"
	SPONSOREDPRODUCTSERRORCODE_SERVICE_UNAVAILABLE   SponsoredProductsErrorCode = "SERVICE_UNAVAILABLE"
	SPONSOREDPRODUCTSERRORCODE_GATEWAY_TIMEOUT       SponsoredProductsErrorCode = "GATEWAY_TIMEOUT"
)

// All allowed values of SponsoredProductsErrorCode enum
var AllowedSponsoredProductsErrorCodeEnumValues = []SponsoredProductsErrorCode{
	"BAD_REQUEST",
	"UNAUTHORIZED",
	"FORBIDDEN",
	"NOT_FOUND",
	"CONTENT_TOO_LARGE",
	"TOO_MANY_REQUESTS",
	"INTERNAL_SERVER_ERROR",
	"BAD_GATEWAY",
	"SERVICE_UNAVAILABLE",
	"GATEWAY_TIMEOUT",
}

func (v *SponsoredProductsErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsErrorCode(value)
	for _, existing := range AllowedSponsoredProductsErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsErrorCode", value)
}

// NewSponsoredProductsErrorCodeFromValue returns a pointer to a valid SponsoredProductsErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsErrorCodeFromValue(v string) (*SponsoredProductsErrorCode, error) {
	ev := SponsoredProductsErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsErrorCode: valid values are %v", v, AllowedSponsoredProductsErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsErrorCode) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsErrorCode value
func (v SponsoredProductsErrorCode) Ptr() *SponsoredProductsErrorCode {
	return &v
}

type NullableSponsoredProductsErrorCode struct {
	value *SponsoredProductsErrorCode
	isSet bool
}

func (v NullableSponsoredProductsErrorCode) Get() *SponsoredProductsErrorCode {
	return v.value
}

func (v *NullableSponsoredProductsErrorCode) Set(val *SponsoredProductsErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsErrorCode(val *SponsoredProductsErrorCode) *NullableSponsoredProductsErrorCode {
	return &NullableSponsoredProductsErrorCode{value: val, isSet: true}
}

func (v NullableSponsoredProductsErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
