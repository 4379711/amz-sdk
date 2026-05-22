package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the CancelSelfShipAppointmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelSelfShipAppointmentResponse{}

// CancelSelfShipAppointmentResponse The `CancelSelfShipAppointment` response.
type CancelSelfShipAppointmentResponse struct {
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _CancelSelfShipAppointmentResponse CancelSelfShipAppointmentResponse

// NewCancelSelfShipAppointmentResponse instantiates a new CancelSelfShipAppointmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelSelfShipAppointmentResponse(operationId string) *CancelSelfShipAppointmentResponse {
	this := CancelSelfShipAppointmentResponse{}
	this.OperationId = operationId
	return &this
}

// NewCancelSelfShipAppointmentResponseWithDefaults instantiates a new CancelSelfShipAppointmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelSelfShipAppointmentResponseWithDefaults() *CancelSelfShipAppointmentResponse {
	this := CancelSelfShipAppointmentResponse{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *CancelSelfShipAppointmentResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *CancelSelfShipAppointmentResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *CancelSelfShipAppointmentResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o CancelSelfShipAppointmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableCancelSelfShipAppointmentResponse struct {
	value *CancelSelfShipAppointmentResponse
	isSet bool
}

func (v NullableCancelSelfShipAppointmentResponse) Get() *CancelSelfShipAppointmentResponse {
	return v.value
}

func (v *NullableCancelSelfShipAppointmentResponse) Set(val *CancelSelfShipAppointmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelSelfShipAppointmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelSelfShipAppointmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelSelfShipAppointmentResponse(val *CancelSelfShipAppointmentResponse) *NullableCancelSelfShipAppointmentResponse {
	return &NullableCancelSelfShipAppointmentResponse{value: val, isSet: true}
}

func (v NullableCancelSelfShipAppointmentResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCancelSelfShipAppointmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
