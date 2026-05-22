package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsCreateOrUpdateEntityState Entity state for create or update operation
type SponsoredProductsCreateOrUpdateEntityState string

// List of SponsoredProductsCreateOrUpdateEntityState
const (
	SPONSOREDPRODUCTSCREATEORUPDATEENTITYSTATE_ENABLED SponsoredProductsCreateOrUpdateEntityState = "ENABLED"
	SPONSOREDPRODUCTSCREATEORUPDATEENTITYSTATE_PAUSED  SponsoredProductsCreateOrUpdateEntityState = "PAUSED"
)

// All allowed values of SponsoredProductsCreateOrUpdateEntityState enum
var AllowedSponsoredProductsCreateOrUpdateEntityStateEnumValues = []SponsoredProductsCreateOrUpdateEntityState{
	"ENABLED",
	"PAUSED",
}

func (v *SponsoredProductsCreateOrUpdateEntityState) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsCreateOrUpdateEntityState(value)
	for _, existing := range AllowedSponsoredProductsCreateOrUpdateEntityStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsCreateOrUpdateEntityState", value)
}

// NewSponsoredProductsCreateOrUpdateEntityStateFromValue returns a pointer to a valid SponsoredProductsCreateOrUpdateEntityState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsCreateOrUpdateEntityStateFromValue(v string) (*SponsoredProductsCreateOrUpdateEntityState, error) {
	ev := SponsoredProductsCreateOrUpdateEntityState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsCreateOrUpdateEntityState: valid values are %v", v, AllowedSponsoredProductsCreateOrUpdateEntityStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsCreateOrUpdateEntityState) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsCreateOrUpdateEntityStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsCreateOrUpdateEntityState value
func (v SponsoredProductsCreateOrUpdateEntityState) Ptr() *SponsoredProductsCreateOrUpdateEntityState {
	return &v
}

type NullableSponsoredProductsCreateOrUpdateEntityState struct {
	value *SponsoredProductsCreateOrUpdateEntityState
	isSet bool
}

func (v NullableSponsoredProductsCreateOrUpdateEntityState) Get() *SponsoredProductsCreateOrUpdateEntityState {
	return v.value
}

func (v *NullableSponsoredProductsCreateOrUpdateEntityState) Set(val *SponsoredProductsCreateOrUpdateEntityState) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateOrUpdateEntityState) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateOrUpdateEntityState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateOrUpdateEntityState(val *SponsoredProductsCreateOrUpdateEntityState) *NullableSponsoredProductsCreateOrUpdateEntityState {
	return &NullableSponsoredProductsCreateOrUpdateEntityState{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateOrUpdateEntityState) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateOrUpdateEntityState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
