package shipping_v2

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// DetailCodes A list of codes used to provide additional shipment information.
type DetailCodes string

// List of DetailCodes
const (
	DETAILCODES_BUSINESS_CLOSED                         DetailCodes = "BusinessClosed"
	DETAILCODES_CUSTOMER_UNAVAILABLE                    DetailCodes = "CustomerUnavailable"
	DETAILCODES_PAYMENT_NOT_READY                       DetailCodes = "PaymentNotReady"
	DETAILCODES_OTP_NOT_AVAILABLE                       DetailCodes = "OtpNotAvailable"
	DETAILCODES_DELIVERY_ATTEMPTED                      DetailCodes = "DeliveryAttempted"
	DETAILCODES_UNABLE_TO_ACCESS                        DetailCodes = "UnableToAccess"
	DETAILCODES_UNABLE_TO_CONTACT_RECIPIENT             DetailCodes = "UnableToContactRecipient"
	DETAILCODES_DELIVERED_TO_BEHIND_WHEELIE_BIN         DetailCodes = "DeliveredToBehindWheelieBin"
	DETAILCODES_DELIVERED_TO_PORCH                      DetailCodes = "DeliveredToPorch"
	DETAILCODES_DELIVERED_TO_GARAGE                     DetailCodes = "DeliveredToGarage"
	DETAILCODES_DELIVERED_TO_GARDEN                     DetailCodes = "DeliveredToGarden"
	DETAILCODES_DELIVERED_TO_GREENHOUSE                 DetailCodes = "DeliveredToGreenhouse"
	DETAILCODES_DELIVERED_TO_MAIL_SLOT                  DetailCodes = "DeliveredToMailSlot"
	DETAILCODES_DELIVERED_TO_MAIL_ROOM                  DetailCodes = "DeliveredToMailRoom"
	DETAILCODES_DELIVERED_TO_NEIGHBOR                   DetailCodes = "DeliveredToNeighbor"
	DETAILCODES_DELIVERED_TO_REAR_DOOR                  DetailCodes = "DeliveredToRearDoor"
	DETAILCODES_DELIVERED_TO_RECEPTIONIST               DetailCodes = "DeliveredToReceptionist"
	DETAILCODES_DELIVERED_TO_SHED                       DetailCodes = "DeliveredToShed"
	DETAILCODES_DELIVERED_WITH_OTP                      DetailCodes = "DeliveredWithOTP"
	DETAILCODES_SIGNED                                  DetailCodes = "Signed"
	DETAILCODES_DAMAGED                                 DetailCodes = "Damaged"
	DETAILCODES_INCORRECT_ITEMS                         DetailCodes = "IncorrectItems"
	DETAILCODES_NOT_REQUIRED                            DetailCodes = "NotRequired"
	DETAILCODES_REJECTED                                DetailCodes = "Rejected"
	DETAILCODES_REJECTED_BY_RECIPIENT_WITH_VERIFICATION DetailCodes = "RejectedByRecipientWithVerification"
	DETAILCODES_CANCELLED_BY_RECIPIENT                  DetailCodes = "CancelledByRecipient"
	DETAILCODES_ADDRESS_NOT_FOUND                       DetailCodes = "AddressNotFound"
	DETAILCODES_HAZMAT_SHIPMENT                         DetailCodes = "HazmatShipment"
	DETAILCODES_UNDELIVERABLE                           DetailCodes = "Undeliverable"
	DETAILCODES_ARRIVED_AT_LOCAL_FACILITY               DetailCodes = "ArrivedAtLocalFacility"
)

// All allowed values of DetailCodes enum
var AllowedDetailCodesEnumValues = []DetailCodes{
	DETAILCODES_BUSINESS_CLOSED,
	DETAILCODES_CUSTOMER_UNAVAILABLE,
	DETAILCODES_PAYMENT_NOT_READY,
	DETAILCODES_OTP_NOT_AVAILABLE,
	DETAILCODES_DELIVERY_ATTEMPTED,
	DETAILCODES_UNABLE_TO_ACCESS,
	DETAILCODES_UNABLE_TO_CONTACT_RECIPIENT,
	DETAILCODES_DELIVERED_TO_BEHIND_WHEELIE_BIN,
	DETAILCODES_DELIVERED_TO_PORCH,
	DETAILCODES_DELIVERED_TO_GARAGE,
	DETAILCODES_DELIVERED_TO_GARDEN,
	DETAILCODES_DELIVERED_TO_GREENHOUSE,
	DETAILCODES_DELIVERED_TO_MAIL_SLOT,
	DETAILCODES_DELIVERED_TO_MAIL_ROOM,
	DETAILCODES_DELIVERED_TO_NEIGHBOR,
	DETAILCODES_DELIVERED_TO_REAR_DOOR,
	DETAILCODES_DELIVERED_TO_RECEPTIONIST,
	DETAILCODES_DELIVERED_TO_SHED,
	DETAILCODES_DELIVERED_WITH_OTP,
	DETAILCODES_SIGNED,
	DETAILCODES_DAMAGED,
	DETAILCODES_INCORRECT_ITEMS,
	DETAILCODES_NOT_REQUIRED,
	DETAILCODES_REJECTED,
	DETAILCODES_REJECTED_BY_RECIPIENT_WITH_VERIFICATION,
	DETAILCODES_CANCELLED_BY_RECIPIENT,
	DETAILCODES_ADDRESS_NOT_FOUND,
	DETAILCODES_HAZMAT_SHIPMENT,
	DETAILCODES_UNDELIVERABLE,
	DETAILCODES_ARRIVED_AT_LOCAL_FACILITY,
}

func (v *DetailCodes) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DetailCodes(value)
	for _, existing := range AllowedDetailCodesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DetailCodes", value)
}

// NewDetailCodesFromValue returns a pointer to a valid DetailCodes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDetailCodesFromValue(v string) (*DetailCodes, error) {
	ev := DetailCodes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DetailCodes: valid values are %v", v, AllowedDetailCodesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DetailCodes) IsValid() bool {
	for _, existing := range AllowedDetailCodesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DetailCodes value
func (v DetailCodes) Ptr() *DetailCodes {
	return &v
}

type NullableDetailCodes struct {
	value *DetailCodes
	isSet bool
}

func (v NullableDetailCodes) Get() *DetailCodes {
	return v.value
}

func (v *NullableDetailCodes) Set(val *DetailCodes) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailCodes) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailCodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailCodes(val *DetailCodes) *NullableDetailCodes {
	return &NullableDetailCodes{value: val, isSet: true}
}

func (v NullableDetailCodes) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDetailCodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
