package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateShipmentSourceAddressRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateShipmentSourceAddressRequest{}

// UpdateShipmentSourceAddressRequest The `UpdateShipmentSourceAddress` request.
type UpdateShipmentSourceAddressRequest struct {
	Address AddressInput `json:"address"`
}

type _UpdateShipmentSourceAddressRequest UpdateShipmentSourceAddressRequest

// NewUpdateShipmentSourceAddressRequest instantiates a new UpdateShipmentSourceAddressRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateShipmentSourceAddressRequest(address AddressInput) *UpdateShipmentSourceAddressRequest {
	this := UpdateShipmentSourceAddressRequest{}
	this.Address = address
	return &this
}

// NewUpdateShipmentSourceAddressRequestWithDefaults instantiates a new UpdateShipmentSourceAddressRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateShipmentSourceAddressRequestWithDefaults() *UpdateShipmentSourceAddressRequest {
	this := UpdateShipmentSourceAddressRequest{}
	return &this
}

// GetAddress returns the Address field value
func (o *UpdateShipmentSourceAddressRequest) GetAddress() AddressInput {
	if o == nil {
		var ret AddressInput
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *UpdateShipmentSourceAddressRequest) GetAddressOk() (*AddressInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *UpdateShipmentSourceAddressRequest) SetAddress(v AddressInput) {
	o.Address = v
}

func (o UpdateShipmentSourceAddressRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

type NullableUpdateShipmentSourceAddressRequest struct {
	value *UpdateShipmentSourceAddressRequest
	isSet bool
}

func (v NullableUpdateShipmentSourceAddressRequest) Get() *UpdateShipmentSourceAddressRequest {
	return v.value
}

func (v *NullableUpdateShipmentSourceAddressRequest) Set(val *UpdateShipmentSourceAddressRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateShipmentSourceAddressRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateShipmentSourceAddressRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateShipmentSourceAddressRequest(val *UpdateShipmentSourceAddressRequest) *NullableUpdateShipmentSourceAddressRequest {
	return &NullableUpdateShipmentSourceAddressRequest{value: val, isSet: true}
}

func (v NullableUpdateShipmentSourceAddressRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateShipmentSourceAddressRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
