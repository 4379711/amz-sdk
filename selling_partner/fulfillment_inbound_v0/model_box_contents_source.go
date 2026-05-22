package fulfillment_inbound_v0

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// BoxContentsSource Where the seller provided box contents information for a shipment.
type BoxContentsSource string

// List of BoxContentsSource
const (
	BOXCONTENTSSOURCE_NONE         BoxContentsSource = "NONE"
	BOXCONTENTSSOURCE_FEED         BoxContentsSource = "FEED"
	BOXCONTENTSSOURCE__2_D_BARCODE BoxContentsSource = "2D_BARCODE"
	BOXCONTENTSSOURCE_INTERACTIVE  BoxContentsSource = "INTERACTIVE"
)

// All allowed values of BoxContentsSource enum
var AllowedBoxContentsSourceEnumValues = []BoxContentsSource{
	BOXCONTENTSSOURCE_NONE,
	BOXCONTENTSSOURCE_FEED,
	BOXCONTENTSSOURCE__2_D_BARCODE,
	BOXCONTENTSSOURCE_INTERACTIVE,
}

func (v *BoxContentsSource) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BoxContentsSource(value)
	for _, existing := range AllowedBoxContentsSourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BoxContentsSource", value)
}

// NewBoxContentsSourceFromValue returns a pointer to a valid BoxContentsSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBoxContentsSourceFromValue(v string) (*BoxContentsSource, error) {
	ev := BoxContentsSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BoxContentsSource: valid values are %v", v, AllowedBoxContentsSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BoxContentsSource) IsValid() bool {
	for _, existing := range AllowedBoxContentsSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BoxContentsSource value
func (v BoxContentsSource) Ptr() *BoxContentsSource {
	return &v
}

type NullableBoxContentsSource struct {
	value *BoxContentsSource
	isSet bool
}

func (v NullableBoxContentsSource) Get() *BoxContentsSource {
	return v.value
}

func (v *NullableBoxContentsSource) Set(val *BoxContentsSource) {
	v.value = val
	v.isSet = true
}

func (v NullableBoxContentsSource) IsSet() bool {
	return v.isSet
}

func (v *NullableBoxContentsSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoxContentsSource(val *BoxContentsSource) *NullableBoxContentsSource {
	return &NullableBoxContentsSource{value: val, isSet: true}
}

func (v NullableBoxContentsSource) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBoxContentsSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
