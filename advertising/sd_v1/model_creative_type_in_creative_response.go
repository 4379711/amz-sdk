package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// CreativeTypeInCreativeResponse  The type of the creative. |Name|Description| |----|-----------| |IMAGE |The creative will display static assets (e.g. headline, brandLogo or custom image).| |VIDEO |The creative will display video assets. This type of creative must have video assets provided.|
type CreativeTypeInCreativeResponse string

// List of CreativeTypeInCreativeResponse
const (
	CREATIVETYPEINCREATIVERESPONSE_IMAGE CreativeTypeInCreativeResponse = "IMAGE"
	CREATIVETYPEINCREATIVERESPONSE_VIDEO CreativeTypeInCreativeResponse = "VIDEO"
)

// All allowed values of CreativeTypeInCreativeResponse enum
var AllowedCreativeTypeInCreativeResponseEnumValues = []CreativeTypeInCreativeResponse{
	"IMAGE",
	"VIDEO",
}

func (v *CreativeTypeInCreativeResponse) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreativeTypeInCreativeResponse(value)
	for _, existing := range AllowedCreativeTypeInCreativeResponseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreativeTypeInCreativeResponse", value)
}

// NewCreativeTypeInCreativeResponseFromValue returns a pointer to a valid CreativeTypeInCreativeResponse
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeTypeInCreativeResponseFromValue(v string) (*CreativeTypeInCreativeResponse, error) {
	ev := CreativeTypeInCreativeResponse(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeTypeInCreativeResponse: valid values are %v", v, AllowedCreativeTypeInCreativeResponseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeTypeInCreativeResponse) IsValid() bool {
	for _, existing := range AllowedCreativeTypeInCreativeResponseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreativeTypeInCreativeResponse value
func (v CreativeTypeInCreativeResponse) Ptr() *CreativeTypeInCreativeResponse {
	return &v
}

type NullableCreativeTypeInCreativeResponse struct {
	value *CreativeTypeInCreativeResponse
	isSet bool
}

func (v NullableCreativeTypeInCreativeResponse) Get() *CreativeTypeInCreativeResponse {
	return v.value
}

func (v *NullableCreativeTypeInCreativeResponse) Set(val *CreativeTypeInCreativeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeTypeInCreativeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeTypeInCreativeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeTypeInCreativeResponse(val *CreativeTypeInCreativeResponse) *NullableCreativeTypeInCreativeResponse {
	return &NullableCreativeTypeInCreativeResponse{value: val, isSet: true}
}

func (v NullableCreativeTypeInCreativeResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeTypeInCreativeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
