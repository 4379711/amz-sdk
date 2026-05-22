package aplus_content_20201101

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// DecoratorType The type of rich text decorator.
type DecoratorType string

// List of DecoratorType
const (
	DECORATORTYPE_LIST_ITEM       DecoratorType = "LIST_ITEM"
	DECORATORTYPE_LIST_ORDERED    DecoratorType = "LIST_ORDERED"
	DECORATORTYPE_LIST_UNORDERED  DecoratorType = "LIST_UNORDERED"
	DECORATORTYPE_STYLE_BOLD      DecoratorType = "STYLE_BOLD"
	DECORATORTYPE_STYLE_ITALIC    DecoratorType = "STYLE_ITALIC"
	DECORATORTYPE_STYLE_LINEBREAK DecoratorType = "STYLE_LINEBREAK"
	DECORATORTYPE_STYLE_PARAGRAPH DecoratorType = "STYLE_PARAGRAPH"
	DECORATORTYPE_STYLE_UNDERLINE DecoratorType = "STYLE_UNDERLINE"
)

// All allowed values of DecoratorType enum
var AllowedDecoratorTypeEnumValues = []DecoratorType{
	DECORATORTYPE_LIST_ITEM,
	DECORATORTYPE_LIST_ORDERED,
	DECORATORTYPE_LIST_UNORDERED,
	DECORATORTYPE_STYLE_BOLD,
	DECORATORTYPE_STYLE_ITALIC,
	DECORATORTYPE_STYLE_LINEBREAK,
	DECORATORTYPE_STYLE_PARAGRAPH,
	DECORATORTYPE_STYLE_UNDERLINE,
}

func (v *DecoratorType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DecoratorType(value)
	for _, existing := range AllowedDecoratorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DecoratorType", value)
}

// NewDecoratorTypeFromValue returns a pointer to a valid DecoratorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDecoratorTypeFromValue(v string) (*DecoratorType, error) {
	ev := DecoratorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DecoratorType: valid values are %v", v, AllowedDecoratorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DecoratorType) IsValid() bool {
	for _, existing := range AllowedDecoratorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DecoratorType value
func (v DecoratorType) Ptr() *DecoratorType {
	return &v
}

type NullableDecoratorType struct {
	value *DecoratorType
	isSet bool
}

func (v NullableDecoratorType) Get() *DecoratorType {
	return v.value
}

func (v *NullableDecoratorType) Set(val *DecoratorType) {
	v.value = val
	v.isSet = true
}

func (v NullableDecoratorType) IsSet() bool {
	return v.isSet
}

func (v *NullableDecoratorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecoratorType(val *DecoratorType) *NullableDecoratorType {
	return &NullableDecoratorType{value: val, isSet: true}
}

func (v NullableDecoratorType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDecoratorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
