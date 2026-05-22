package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateShipmentTrackingDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateShipmentTrackingDetailsResponse{}

// UpdateShipmentTrackingDetailsResponse The `updateShipmentTrackingDetails` response.
type UpdateShipmentTrackingDetailsResponse struct {
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _UpdateShipmentTrackingDetailsResponse UpdateShipmentTrackingDetailsResponse

// NewUpdateShipmentTrackingDetailsResponse instantiates a new UpdateShipmentTrackingDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateShipmentTrackingDetailsResponse(operationId string) *UpdateShipmentTrackingDetailsResponse {
	this := UpdateShipmentTrackingDetailsResponse{}
	this.OperationId = operationId
	return &this
}

// NewUpdateShipmentTrackingDetailsResponseWithDefaults instantiates a new UpdateShipmentTrackingDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateShipmentTrackingDetailsResponseWithDefaults() *UpdateShipmentTrackingDetailsResponse {
	this := UpdateShipmentTrackingDetailsResponse{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *UpdateShipmentTrackingDetailsResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *UpdateShipmentTrackingDetailsResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *UpdateShipmentTrackingDetailsResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o UpdateShipmentTrackingDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableUpdateShipmentTrackingDetailsResponse struct {
	value *UpdateShipmentTrackingDetailsResponse
	isSet bool
}

func (v NullableUpdateShipmentTrackingDetailsResponse) Get() *UpdateShipmentTrackingDetailsResponse {
	return v.value
}

func (v *NullableUpdateShipmentTrackingDetailsResponse) Set(val *UpdateShipmentTrackingDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateShipmentTrackingDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateShipmentTrackingDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateShipmentTrackingDetailsResponse(val *UpdateShipmentTrackingDetailsResponse) *NullableUpdateShipmentTrackingDetailsResponse {
	return &NullableUpdateShipmentTrackingDetailsResponse{value: val, isSet: true}
}

func (v NullableUpdateShipmentTrackingDetailsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateShipmentTrackingDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
