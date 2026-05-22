package shipping_v2

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ClaimReason The reason for which shipper is filing the claim for a particular shipment.
type ClaimReason string

// List of ClaimReason
const (
	CLAIMREASON_LOST_IN_TRANSIT         ClaimReason = "LOST_IN_TRANSIT"
	CLAIMREASON_DAMAGED_IN_TRANSIT      ClaimReason = "DAMAGED_IN_TRANSIT"
	CLAIMREASON_DELIVERED_NOT_RECEIVED  ClaimReason = "DELIVERED_NOT_RECEIVED"
	CLAIMREASON_ITEM_MISSING_SWITCHEROO ClaimReason = "ITEM_MISSING_SWITCHEROO"
	CLAIMREASON_COD_ABUSE               ClaimReason = "COD_ABUSE"
)

// All allowed values of ClaimReason enum
var AllowedClaimReasonEnumValues = []ClaimReason{
	CLAIMREASON_LOST_IN_TRANSIT,
	CLAIMREASON_DAMAGED_IN_TRANSIT,
	CLAIMREASON_DELIVERED_NOT_RECEIVED,
	CLAIMREASON_ITEM_MISSING_SWITCHEROO,
	CLAIMREASON_COD_ABUSE,
}

func (v *ClaimReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClaimReason(value)
	for _, existing := range AllowedClaimReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClaimReason", value)
}

// NewClaimReasonFromValue returns a pointer to a valid ClaimReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClaimReasonFromValue(v string) (*ClaimReason, error) {
	ev := ClaimReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClaimReason: valid values are %v", v, AllowedClaimReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClaimReason) IsValid() bool {
	for _, existing := range AllowedClaimReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClaimReason value
func (v ClaimReason) Ptr() *ClaimReason {
	return &v
}

type NullableClaimReason struct {
	value *ClaimReason
	isSet bool
}

func (v NullableClaimReason) Get() *ClaimReason {
	return v.value
}

func (v *NullableClaimReason) Set(val *ClaimReason) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimReason) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimReason(val *ClaimReason) *NullableClaimReason {
	return &NullableClaimReason{value: val, isSet: true}
}

func (v NullableClaimReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableClaimReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
