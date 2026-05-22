package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateShipmentSourceAddressResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateShipmentSourceAddressResponse{}

// UpdateShipmentSourceAddressResponse The `UpdateShipmentSourceAddress` response.
type UpdateShipmentSourceAddressResponse struct {
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _UpdateShipmentSourceAddressResponse UpdateShipmentSourceAddressResponse

// NewUpdateShipmentSourceAddressResponse instantiates a new UpdateShipmentSourceAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateShipmentSourceAddressResponse(operationId string) *UpdateShipmentSourceAddressResponse {
	this := UpdateShipmentSourceAddressResponse{}
	this.OperationId = operationId
	return &this
}

// NewUpdateShipmentSourceAddressResponseWithDefaults instantiates a new UpdateShipmentSourceAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateShipmentSourceAddressResponseWithDefaults() *UpdateShipmentSourceAddressResponse {
	this := UpdateShipmentSourceAddressResponse{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *UpdateShipmentSourceAddressResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *UpdateShipmentSourceAddressResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *UpdateShipmentSourceAddressResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o UpdateShipmentSourceAddressResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableUpdateShipmentSourceAddressResponse struct {
	value *UpdateShipmentSourceAddressResponse
	isSet bool
}

func (v NullableUpdateShipmentSourceAddressResponse) Get() *UpdateShipmentSourceAddressResponse {
	return v.value
}

func (v *NullableUpdateShipmentSourceAddressResponse) Set(val *UpdateShipmentSourceAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateShipmentSourceAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateShipmentSourceAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateShipmentSourceAddressResponse(val *UpdateShipmentSourceAddressResponse) *NullableUpdateShipmentSourceAddressResponse {
	return &NullableUpdateShipmentSourceAddressResponse{value: val, isSet: true}
}

func (v NullableUpdateShipmentSourceAddressResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateShipmentSourceAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
