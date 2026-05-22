package fulfillment_inbound_20240320

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// PrepType Preparation instructions for shipping an item to Amazon's fulfillment network. For more information about preparing items for shipment to Amazon's fulfillment network, refer to [Seller Central Help for your marketplace](https://developer-docs.amazon.com/sp-api/docs/seller-central-urls).
type PrepType string

// List of PrepType
const (
	PREPTYPE_ITEM_BLACK_SHRINKWRAP PrepType = "ITEM_BLACK_SHRINKWRAP"
	PREPTYPE_ITEM_BLANKSTK         PrepType = "ITEM_BLANKSTK"
	PREPTYPE_ITEM_BOXING           PrepType = "ITEM_BOXING"
	PREPTYPE_ITEM_BUBBLEWRAP       PrepType = "ITEM_BUBBLEWRAP"
	PREPTYPE_ITEM_CAP_SEALING      PrepType = "ITEM_CAP_SEALING"
	PREPTYPE_ITEM_DEBUNDLE         PrepType = "ITEM_DEBUNDLE"
	PREPTYPE_ITEM_HANG_GARMENT     PrepType = "ITEM_HANG_GARMENT"
	PREPTYPE_ITEM_LABELING         PrepType = "ITEM_LABELING"
	PREPTYPE_ITEM_NO_PREP          PrepType = "ITEM_NO_PREP"
	PREPTYPE_ITEM_POLYBAGGING      PrepType = "ITEM_POLYBAGGING"
	PREPTYPE_ITEM_RMOVHANG         PrepType = "ITEM_RMOVHANG"
	PREPTYPE_ITEM_SETCREAT         PrepType = "ITEM_SETCREAT"
	PREPTYPE_ITEM_SETSTK           PrepType = "ITEM_SETSTK"
	PREPTYPE_ITEM_SIOC             PrepType = "ITEM_SIOC"
	PREPTYPE_ITEM_SUFFOSTK         PrepType = "ITEM_SUFFOSTK"
	PREPTYPE_ITEM_TAPING           PrepType = "ITEM_TAPING"
)

// All allowed values of PrepType enum
var AllowedPrepTypeEnumValues = []PrepType{
	PREPTYPE_ITEM_BLACK_SHRINKWRAP,
	PREPTYPE_ITEM_BLANKSTK,
	PREPTYPE_ITEM_BOXING,
	PREPTYPE_ITEM_BUBBLEWRAP,
	PREPTYPE_ITEM_CAP_SEALING,
	PREPTYPE_ITEM_DEBUNDLE,
	PREPTYPE_ITEM_HANG_GARMENT,
	PREPTYPE_ITEM_LABELING,
	PREPTYPE_ITEM_NO_PREP,
	PREPTYPE_ITEM_POLYBAGGING,
	PREPTYPE_ITEM_RMOVHANG,
	PREPTYPE_ITEM_SETCREAT,
	PREPTYPE_ITEM_SETSTK,
	PREPTYPE_ITEM_SIOC,
	PREPTYPE_ITEM_SUFFOSTK,
	PREPTYPE_ITEM_TAPING,
}

func (v *PrepType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrepType(value)
	for _, existing := range AllowedPrepTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrepType", value)
}

// NewPrepTypeFromValue returns a pointer to a valid PrepType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrepTypeFromValue(v string) (*PrepType, error) {
	ev := PrepType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrepType: valid values are %v", v, AllowedPrepTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrepType) IsValid() bool {
	for _, existing := range AllowedPrepTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PrepType value
func (v PrepType) Ptr() *PrepType {
	return &v
}

type NullablePrepType struct {
	value *PrepType
	isSet bool
}

func (v NullablePrepType) Get() *PrepType {
	return v.value
}

func (v *NullablePrepType) Set(val *PrepType) {
	v.value = val
	v.isSet = true
}

func (v NullablePrepType) IsSet() bool {
	return v.isSet
}

func (v *NullablePrepType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrepType(val *PrepType) *NullablePrepType {
	return &NullablePrepType{value: val, isSet: true}
}

func (v NullablePrepType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePrepType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
