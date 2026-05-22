package easy_ship_20220323

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// PackageStatus The status of the package.
type PackageStatus string

// List of PackageStatus
const (
	PACKAGESTATUS_READY_FOR_PICKUP   PackageStatus = "ReadyForPickup"
	PACKAGESTATUS_PICKED_UP          PackageStatus = "PickedUp"
	PACKAGESTATUS_AT_ORIGIN_FC       PackageStatus = "AtOriginFC"
	PACKAGESTATUS_AT_DESTINATION_FC  PackageStatus = "AtDestinationFC"
	PACKAGESTATUS_DELIVERED          PackageStatus = "Delivered"
	PACKAGESTATUS_REJECTED           PackageStatus = "Rejected"
	PACKAGESTATUS_UNDELIVERABLE      PackageStatus = "Undeliverable"
	PACKAGESTATUS_RETURNED_TO_SELLER PackageStatus = "ReturnedToSeller"
	PACKAGESTATUS_LOST_IN_TRANSIT    PackageStatus = "LostInTransit"
	PACKAGESTATUS_LABEL_CANCELED     PackageStatus = "LabelCanceled"
	PACKAGESTATUS_DAMAGED_IN_TRANSIT PackageStatus = "DamagedInTransit"
	PACKAGESTATUS_OUT_FOR_DELIVERY   PackageStatus = "OutForDelivery"
)

// All allowed values of PackageStatus enum
var AllowedPackageStatusEnumValues = []PackageStatus{
	PACKAGESTATUS_READY_FOR_PICKUP,
	PACKAGESTATUS_PICKED_UP,
	PACKAGESTATUS_AT_ORIGIN_FC,
	PACKAGESTATUS_AT_DESTINATION_FC,
	PACKAGESTATUS_DELIVERED,
	PACKAGESTATUS_REJECTED,
	PACKAGESTATUS_UNDELIVERABLE,
	PACKAGESTATUS_RETURNED_TO_SELLER,
	PACKAGESTATUS_LOST_IN_TRANSIT,
	PACKAGESTATUS_LABEL_CANCELED,
	PACKAGESTATUS_DAMAGED_IN_TRANSIT,
	PACKAGESTATUS_OUT_FOR_DELIVERY,
}

func (v *PackageStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PackageStatus(value)
	for _, existing := range AllowedPackageStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PackageStatus", value)
}

// NewPackageStatusFromValue returns a pointer to a valid PackageStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPackageStatusFromValue(v string) (*PackageStatus, error) {
	ev := PackageStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PackageStatus: valid values are %v", v, AllowedPackageStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PackageStatus) IsValid() bool {
	for _, existing := range AllowedPackageStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PackageStatus value
func (v PackageStatus) Ptr() *PackageStatus {
	return &v
}

type NullablePackageStatus struct {
	value *PackageStatus
	isSet bool
}

func (v NullablePackageStatus) Get() *PackageStatus {
	return v.value
}

func (v *NullablePackageStatus) Set(val *PackageStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageStatus(val *PackageStatus) *NullablePackageStatus {
	return &NullablePackageStatus{value: val, isSet: true}
}

func (v NullablePackageStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackageStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
