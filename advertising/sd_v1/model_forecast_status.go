package sd_v1

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// ForecastStatus It contains the forecast status. The IMPRESSION_TARGETING_TOO_NARROW field means the targeting  clauses are too narrow, and the IMPRESSION_TARGETING_TOO_BROAD field means the targeting clauses are too broad,  so our inventory impression forecast won't provide any useful information. The COMPLETE field means all the forecasts are complete.
type ForecastStatus string

// List of ForecastStatus
const (
	FORECASTSTATUS_IMPRESSION_TARGETING_TOO_NARROW ForecastStatus = "IMPRESSION_TARGETING_TOO_NARROW"
	FORECASTSTATUS_IMPRESSION_TARGETING_TOO_BROAD  ForecastStatus = "IMPRESSION_TARGETING_TOO_BROAD"
	FORECASTSTATUS_COMPLETE                        ForecastStatus = "COMPLETE"
)

// All allowed values of ForecastStatus enum
var AllowedForecastStatusEnumValues = []ForecastStatus{
	"IMPRESSION_TARGETING_TOO_NARROW",
	"IMPRESSION_TARGETING_TOO_BROAD",
	"COMPLETE",
}

func (v *ForecastStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ForecastStatus(value)
	for _, existing := range AllowedForecastStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ForecastStatus", value)
}

// NewForecastStatusFromValue returns a pointer to a valid ForecastStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewForecastStatusFromValue(v string) (*ForecastStatus, error) {
	ev := ForecastStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ForecastStatus: valid values are %v", v, AllowedForecastStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ForecastStatus) IsValid() bool {
	for _, existing := range AllowedForecastStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ForecastStatus value
func (v ForecastStatus) Ptr() *ForecastStatus {
	return &v
}

type NullableForecastStatus struct {
	value *ForecastStatus
	isSet bool
}

func (v NullableForecastStatus) Get() *ForecastStatus {
	return v.value
}

func (v *NullableForecastStatus) Set(val *ForecastStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableForecastStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableForecastStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForecastStatus(val *ForecastStatus) *NullableForecastStatus {
	return &NullableForecastStatus{value: val, isSet: true}
}

func (v NullableForecastStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableForecastStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
