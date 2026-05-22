package merchant_fulfillment_v0

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// FileType The file type for a label.
type FileType string

// List of FileType
const (
	FILETYPE_APPLICATION_PDF FileType = "application/pdf"
	FILETYPE_APPLICATION_ZPL FileType = "application/zpl"
	FILETYPE_IMAGE_PNG       FileType = "image/png"
)

// All allowed values of FileType enum
var AllowedFileTypeEnumValues = []FileType{
	FILETYPE_APPLICATION_PDF,
	FILETYPE_APPLICATION_ZPL,
	FILETYPE_IMAGE_PNG,
}

func (v *FileType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileType(value)
	for _, existing := range AllowedFileTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileType", value)
}

// NewFileTypeFromValue returns a pointer to a valid FileType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileTypeFromValue(v string) (*FileType, error) {
	ev := FileType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileType: valid values are %v", v, AllowedFileTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileType) IsValid() bool {
	for _, existing := range AllowedFileTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileType value
func (v FileType) Ptr() *FileType {
	return &v
}

type NullableFileType struct {
	value *FileType
	isSet bool
}

func (v NullableFileType) Get() *FileType {
	return v.value
}

func (v *NullableFileType) Set(val *FileType) {
	v.value = val
	v.isSet = true
}

func (v NullableFileType) IsSet() bool {
	return v.isSet
}

func (v *NullableFileType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileType(val *FileType) *NullableFileType {
	return &NullableFileType{value: val, isSet: true}
}

func (v NullableFileType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFileType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
