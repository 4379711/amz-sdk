package fulfillment_inbound_v0

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// PrepGuidance Item preparation instructions.
type PrepGuidance string

// List of PrepGuidance
const (
	PREPGUIDANCE_CONSULT_HELP_DOCUMENTS      PrepGuidance = "ConsultHelpDocuments"
	PREPGUIDANCE_NO_ADDITIONAL_PREP_REQUIRED PrepGuidance = "NoAdditionalPrepRequired"
	PREPGUIDANCE_SEE_PREP_INSTRUCTIONS_LIST  PrepGuidance = "SeePrepInstructionsList"
)

// All allowed values of PrepGuidance enum
var AllowedPrepGuidanceEnumValues = []PrepGuidance{
	PREPGUIDANCE_CONSULT_HELP_DOCUMENTS,
	PREPGUIDANCE_NO_ADDITIONAL_PREP_REQUIRED,
	PREPGUIDANCE_SEE_PREP_INSTRUCTIONS_LIST,
}

func (v *PrepGuidance) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrepGuidance(value)
	for _, existing := range AllowedPrepGuidanceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrepGuidance", value)
}

// NewPrepGuidanceFromValue returns a pointer to a valid PrepGuidance
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrepGuidanceFromValue(v string) (*PrepGuidance, error) {
	ev := PrepGuidance(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrepGuidance: valid values are %v", v, AllowedPrepGuidanceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrepGuidance) IsValid() bool {
	for _, existing := range AllowedPrepGuidanceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PrepGuidance value
func (v PrepGuidance) Ptr() *PrepGuidance {
	return &v
}

type NullablePrepGuidance struct {
	value *PrepGuidance
	isSet bool
}

func (v NullablePrepGuidance) Get() *PrepGuidance {
	return v.value
}

func (v *NullablePrepGuidance) Set(val *PrepGuidance) {
	v.value = val
	v.isSet = true
}

func (v NullablePrepGuidance) IsSet() bool {
	return v.isSet
}

func (v *NullablePrepGuidance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrepGuidance(val *PrepGuidance) *NullablePrepGuidance {
	return &NullablePrepGuidance{value: val, isSet: true}
}

func (v NullablePrepGuidance) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePrepGuidance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
