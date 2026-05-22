package fulfillment_inbound_20240320

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// LabelOwner Specifies who will label the items. Options include `AMAZON`, `SELLER` or `NONE`.
type LabelOwner string

// List of LabelOwner
const (
	LABELOWNER_AMAZON LabelOwner = "AMAZON"
	LABELOWNER_SELLER LabelOwner = "SELLER"
	LABELOWNER_NONE   LabelOwner = "NONE"
)

// All allowed values of LabelOwner enum
var AllowedLabelOwnerEnumValues = []LabelOwner{
	LABELOWNER_AMAZON,
	LABELOWNER_SELLER,
	LABELOWNER_NONE,
}

func (v *LabelOwner) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LabelOwner(value)
	for _, existing := range AllowedLabelOwnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LabelOwner", value)
}

// NewLabelOwnerFromValue returns a pointer to a valid LabelOwner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLabelOwnerFromValue(v string) (*LabelOwner, error) {
	ev := LabelOwner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LabelOwner: valid values are %v", v, AllowedLabelOwnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LabelOwner) IsValid() bool {
	for _, existing := range AllowedLabelOwnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LabelOwner value
func (v LabelOwner) Ptr() *LabelOwner {
	return &v
}

type NullableLabelOwner struct {
	value *LabelOwner
	isSet bool
}

func (v NullableLabelOwner) Get() *LabelOwner {
	return v.value
}

func (v *NullableLabelOwner) Set(val *LabelOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelOwner(val *LabelOwner) *NullableLabelOwner {
	return &NullableLabelOwner{value: val, isSet: true}
}

func (v NullableLabelOwner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLabelOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
