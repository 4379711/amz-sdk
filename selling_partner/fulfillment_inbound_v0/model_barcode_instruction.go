package fulfillment_inbound_v0

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// BarcodeInstruction Labeling requirements for the item. For more information about FBA labeling requirements, see the Seller Central Help for your marketplace.
type BarcodeInstruction string

// List of BarcodeInstruction
const (
	BARCODEINSTRUCTION_REQUIRES_FNSKU_LABEL     BarcodeInstruction = "RequiresFNSKULabel"
	BARCODEINSTRUCTION_CAN_USE_ORIGINAL_BARCODE BarcodeInstruction = "CanUseOriginalBarcode"
	BARCODEINSTRUCTION_MUST_PROVIDE_SELLER_SKU  BarcodeInstruction = "MustProvideSellerSKU"
)

// All allowed values of BarcodeInstruction enum
var AllowedBarcodeInstructionEnumValues = []BarcodeInstruction{
	BARCODEINSTRUCTION_REQUIRES_FNSKU_LABEL,
	BARCODEINSTRUCTION_CAN_USE_ORIGINAL_BARCODE,
	BARCODEINSTRUCTION_MUST_PROVIDE_SELLER_SKU,
}

func (v *BarcodeInstruction) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BarcodeInstruction(value)
	for _, existing := range AllowedBarcodeInstructionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BarcodeInstruction", value)
}

// NewBarcodeInstructionFromValue returns a pointer to a valid BarcodeInstruction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBarcodeInstructionFromValue(v string) (*BarcodeInstruction, error) {
	ev := BarcodeInstruction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BarcodeInstruction: valid values are %v", v, AllowedBarcodeInstructionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BarcodeInstruction) IsValid() bool {
	for _, existing := range AllowedBarcodeInstructionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BarcodeInstruction value
func (v BarcodeInstruction) Ptr() *BarcodeInstruction {
	return &v
}

type NullableBarcodeInstruction struct {
	value *BarcodeInstruction
	isSet bool
}

func (v NullableBarcodeInstruction) Get() *BarcodeInstruction {
	return v.value
}

func (v *NullableBarcodeInstruction) Set(val *BarcodeInstruction) {
	v.value = val
	v.isSet = true
}

func (v NullableBarcodeInstruction) IsSet() bool {
	return v.isSet
}

func (v *NullableBarcodeInstruction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBarcodeInstruction(val *BarcodeInstruction) *NullableBarcodeInstruction {
	return &NullableBarcodeInstruction{value: val, isSet: true}
}

func (v NullableBarcodeInstruction) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBarcodeInstruction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
