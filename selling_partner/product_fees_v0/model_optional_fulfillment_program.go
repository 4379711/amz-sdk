package product_fees_v0

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// OptionalFulfillmentProgram An optional enrollment program to return the estimated fees when the offer is fulfilled by Amazon (IsAmazonFulfilled is set to true).
type OptionalFulfillmentProgram string

// List of OptionalFulfillmentProgram
const (
	OPTIONALFULFILLMENTPROGRAM_FBA_CORE OptionalFulfillmentProgram = "FBA_CORE"
	OPTIONALFULFILLMENTPROGRAM_FBA_SNL  OptionalFulfillmentProgram = "FBA_SNL"
	OPTIONALFULFILLMENTPROGRAM_FBA_EFN  OptionalFulfillmentProgram = "FBA_EFN"
)

// All allowed values of OptionalFulfillmentProgram enum
var AllowedOptionalFulfillmentProgramEnumValues = []OptionalFulfillmentProgram{
	OPTIONALFULFILLMENTPROGRAM_FBA_CORE,
	OPTIONALFULFILLMENTPROGRAM_FBA_SNL,
	OPTIONALFULFILLMENTPROGRAM_FBA_EFN,
}

func (v *OptionalFulfillmentProgram) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OptionalFulfillmentProgram(value)
	for _, existing := range AllowedOptionalFulfillmentProgramEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OptionalFulfillmentProgram", value)
}

// NewOptionalFulfillmentProgramFromValue returns a pointer to a valid OptionalFulfillmentProgram
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOptionalFulfillmentProgramFromValue(v string) (*OptionalFulfillmentProgram, error) {
	ev := OptionalFulfillmentProgram(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OptionalFulfillmentProgram: valid values are %v", v, AllowedOptionalFulfillmentProgramEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OptionalFulfillmentProgram) IsValid() bool {
	for _, existing := range AllowedOptionalFulfillmentProgramEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OptionalFulfillmentProgram value
func (v OptionalFulfillmentProgram) Ptr() *OptionalFulfillmentProgram {
	return &v
}

type NullableOptionalFulfillmentProgram struct {
	value *OptionalFulfillmentProgram
	isSet bool
}

func (v NullableOptionalFulfillmentProgram) Get() *OptionalFulfillmentProgram {
	return v.value
}

func (v *NullableOptionalFulfillmentProgram) Set(val *OptionalFulfillmentProgram) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionalFulfillmentProgram) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionalFulfillmentProgram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionalFulfillmentProgram(val *OptionalFulfillmentProgram) *NullableOptionalFulfillmentProgram {
	return &NullableOptionalFulfillmentProgram{value: val, isSet: true}
}

func (v NullableOptionalFulfillmentProgram) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptionalFulfillmentProgram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
