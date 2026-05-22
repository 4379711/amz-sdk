package supply_sources_20200701

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// ParkingSpotIdentificationType The type of parking spot identification.
type ParkingSpotIdentificationType string

// List of ParkingSpotIdentificationType
const (
	PARKINGSPOTIDENTIFICATIONTYPE_NUMBERED ParkingSpotIdentificationType = "Numbered"
	PARKINGSPOTIDENTIFICATIONTYPE_OTHER    ParkingSpotIdentificationType = "Other"
)

// All allowed values of ParkingSpotIdentificationType enum
var AllowedParkingSpotIdentificationTypeEnumValues = []ParkingSpotIdentificationType{
	PARKINGSPOTIDENTIFICATIONTYPE_NUMBERED,
	PARKINGSPOTIDENTIFICATIONTYPE_OTHER,
}

func (v *ParkingSpotIdentificationType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ParkingSpotIdentificationType(value)
	for _, existing := range AllowedParkingSpotIdentificationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ParkingSpotIdentificationType", value)
}

// NewParkingSpotIdentificationTypeFromValue returns a pointer to a valid ParkingSpotIdentificationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewParkingSpotIdentificationTypeFromValue(v string) (*ParkingSpotIdentificationType, error) {
	ev := ParkingSpotIdentificationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ParkingSpotIdentificationType: valid values are %v", v, AllowedParkingSpotIdentificationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ParkingSpotIdentificationType) IsValid() bool {
	for _, existing := range AllowedParkingSpotIdentificationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ParkingSpotIdentificationType value
func (v ParkingSpotIdentificationType) Ptr() *ParkingSpotIdentificationType {
	return &v
}

type NullableParkingSpotIdentificationType struct {
	value *ParkingSpotIdentificationType
	isSet bool
}

func (v NullableParkingSpotIdentificationType) Get() *ParkingSpotIdentificationType {
	return v.value
}

func (v *NullableParkingSpotIdentificationType) Set(val *ParkingSpotIdentificationType) {
	v.value = val
	v.isSet = true
}

func (v NullableParkingSpotIdentificationType) IsSet() bool {
	return v.isSet
}

func (v *NullableParkingSpotIdentificationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParkingSpotIdentificationType(val *ParkingSpotIdentificationType) *NullableParkingSpotIdentificationType {
	return &NullableParkingSpotIdentificationType{value: val, isSet: true}
}

func (v NullableParkingSpotIdentificationType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableParkingSpotIdentificationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
