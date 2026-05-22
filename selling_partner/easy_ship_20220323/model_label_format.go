package easy_ship_20220323

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// LabelFormat The file format in which the shipping label will be created.
type LabelFormat string

// List of LabelFormat
const (
	LABELFORMAT_PDF LabelFormat = "PDF"
	LABELFORMAT_ZPL LabelFormat = "ZPL"
)

// All allowed values of LabelFormat enum
var AllowedLabelFormatEnumValues = []LabelFormat{
	LABELFORMAT_PDF,
	LABELFORMAT_ZPL,
}

func (v *LabelFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LabelFormat(value)
	for _, existing := range AllowedLabelFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LabelFormat", value)
}

// NewLabelFormatFromValue returns a pointer to a valid LabelFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLabelFormatFromValue(v string) (*LabelFormat, error) {
	ev := LabelFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LabelFormat: valid values are %v", v, AllowedLabelFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LabelFormat) IsValid() bool {
	for _, existing := range AllowedLabelFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LabelFormat value
func (v LabelFormat) Ptr() *LabelFormat {
	return &v
}

type NullableLabelFormat struct {
	value *LabelFormat
	isSet bool
}

func (v NullableLabelFormat) Get() *LabelFormat {
	return v.value
}

func (v *NullableLabelFormat) Set(val *LabelFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelFormat(val *LabelFormat) *NullableLabelFormat {
	return &NullableLabelFormat{value: val, isSet: true}
}

func (v NullableLabelFormat) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLabelFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
